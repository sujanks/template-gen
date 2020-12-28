package task

import (
	"fmt"
	"github.com/kube-sailmaker/template-gen/functions"
	"github.com/kube-sailmaker/template-gen/model"
	templates "github.com/kube-sailmaker/template-gen/template"
	"strings"
)

const (
	appPath        = "%s/user/apps/%s.yaml"
	cpu            = "cpu"
	memory         = "memory"
	replicas       = "replicas"
	slashSeparator = "/"
	filePrefix     = "file:/"
)

//CPU value mapping
var CPU = map[string]string{
	"c05":     "0.5",
	"c1":      "1",
	"c2":      "2",
	"c3":      "3",
	"default": "0.5",
}

//Memory value mapping
var MEMORY = map[string]string{
	"m05":     "0.5Gi",
	"m1":      "1Gi",
	"m2":      "2Gi",
	"m3":      "3Gi",
	"default": "500Mi",
}

func ProcessApplication(app *model.App, args *model.Args) *templates.Application {
	appFile := fmt.Sprintf(appPath, args.ManifestDir, app.Name)
	application := &model.Application{}
	functions.UnmarshalFile(appFile, application)

	appValues := templates.Application{
		Name:           app.Name,
		Tag:            app.Version,
		ReleaseName:    args.ReleaseName,
		Kind:           application.Kind,
		Annotations:    application.Annotations,
		LivenessProbe:  application.LivenessProbe,
		ReadinessProbe: application.ReadinessProbe,
		NodeSelector:   application.NodeSelector,
		Volumes:        application.Volumes,
	}
	CopyService(application, &appValues)
	CopyConfigMap(application, &appValues)
	CopyVolumeMount(application, &appValues)
	GenerateEnvVars(application, args, &appValues)
	GenerateResourceLimit(application, args, &appValues)
	GenerateMixins(application, args, &appValues)

	return &appValues
}

func CopyService(application *model.Application, appValues *templates.Application) {
	if len(application.Service.Name) > 0 {
		appValues.Service = templates.Service{
			Name:       application.Service.Name,
			Type:       application.Service.Type,
			Port:       application.Service.Port,
			TargetPort: application.Service.TargetPort,
		}
	}
}

func CopyConfigMap(application *model.Application, appValues *templates.Application) {
	appValues.ConfigMaps = make([]map[string]interface{}, 0)
	for _, value := range application.ConfigMaps {
		config := make(map[string]interface{}, 0)
		for k, v := range value {
			if strings.HasPrefix(v, filePrefix) {
				filename := strings.Split(v, filePrefix)[1]
				content := functions.ReadFile(filename)
				stringContent := fmt.Sprintf("|\n%s", string(*content))
				config[filename] = stringContent
			} else {
				config[k] = v
			}
		}
		appValues.ConfigMaps = append(appValues.ConfigMaps, config)
	}
}

func CopyVolumeMount(application *model.Application, appValues *templates.Application) {
	appValues.VolumeMounts = make([]templates.Mount, 0)
	for _, value := range application.VolumeMounts {
		mount := templates.Mount{
			Name:      value.Name,
			MountPath: value.MountPath,
		}
		appValues.VolumeMounts = append(appValues.VolumeMounts, mount)
	}
}

func GenerateMixins(application *model.Application, args *model.Args, appValues *templates.Application) {
	appValues.Command = make([]string, 0)
	appValues.Entrypoint = make([]string, 0)
	for _, mxin := range application.Mixins {
		mixinType := strings.Split(mxin, slashSeparator)
		name := mixinType[0]
		mType := mixinType[1]
		mixinList := model.MixinList{}
		GetMixin(name, &mixinList, args)
		for _, m := range mixinList.Mixin {
			if mType == m.Name {
				for k, v := range m.Env {
					appValues.EnvVars[k] = v
				}
				appValues.Command = m.Cmd
				appValues.Entrypoint = m.Entrypoint
			}
		}
	}
}

func GenerateResourceLimit(application *model.Application, args *model.Args, appValues *templates.Application) {
	appValues.Limits = make(map[string]string, 0)
	//Process app template
	for _, tmpl := range application.Template {
		env := tmpl.Name
		if env == args.Env {
			configs := tmpl.Config
			cpuLimit := CPU["default"]
			if val, ok := configs[cpu]; ok {
				cpuLimit = CPU[val]
			}
			memLimit := MEMORY["default"]
			if val, ok := configs[memory]; ok {
				memLimit = MEMORY[val]
			}

			replicaLimit := "1"
			if val, ok := configs[replicas]; ok {
				replicaLimit = val
			}
			appValues.Limits[cpu] = cpuLimit
			appValues.Limits[memory] = memLimit
			appValues.Replicas = replicaLimit
		}
	}
}

func GenerateEnvVars(application *model.Application, args *model.Args, appValues *templates.Application) {
	appEnvVars := make(map[string]string, 0)

	//Process app resources
	for _, appRes := range application.Resources {
		//elasticsearch-user:sit
		resDetails := strings.Split(appRes, slashSeparator)
		//TODO: Error handle
		name := resDetails[0]
		envType := resDetails[1]
		resource := &model.Resource{}
		GetResource(name, &resource, args)
		for _, resTemplate := range resource.Spec.ResourceTemplate {
			//Only using the context
			if resTemplate.Name == envType {
				addToEnvVars(name, appEnvVars, resTemplate.Element)

				//cassandra-cluster-a:test
				if len(resTemplate.Infra) > 0 {
					infra := strings.Split(resTemplate.Infra, slashSeparator)
					//TODO: Error handle
					infraName := infra[0]
					infraEnv := infra[1]
					infrastructure := &model.Infrastructure{}
					GetInfrastructure(infraName, &infrastructure, args)
					for _, infraTemplate := range infrastructure.Spec.Template {
						if infraEnv == infraTemplate.Name {
							addToEnvVars(name, appEnvVars, infraTemplate.Attributes)
						}
					}
				}
			}
		}
	}
	appValues.EnvVars = appEnvVars
}

func addToEnvVars(name string, appEnvVars map[string]string, items map[string]string) {
	infraName := strings.ReplaceAll(name, "-", "_")
	for k, v := range items {
		key := strings.ToUpper(fmt.Sprintf("%s_%s", infraName, k))
		appEnvVars[key] = v
	}
}
