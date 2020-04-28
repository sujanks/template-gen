package task

import (
	"context"
	"fmt"
	"github.com/kube-sailmaker/template-gen/functions"
	"github.com/kube-sailmaker/template-gen/model"
	templates "github.com/kube-sailmaker/template-gen/template"
	"strings"
)

const (
	appPath = "/Users/ss/workspace/manifest/user/apps/%s.yaml"
	cpu     = "cpu"
	memory  = "memory"
	replica = "replica"
	sep     = ":"
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

func ProcessApplication(app *model.App, args *model.Args) {
	appFile := fmt.Sprintf(appPath, app.Name)
	application := &model.Application{}
	functions.UnmarshalFile(appFile, &application)

	envVars := GenerateEnvVars(application, args)
	limits := GetResourceLimit(application, args)

	tmlValues := templates.TemplateValues{
		Application: templates.Application{
			Name:           app.Name,
			Tag:            app.Version,
			LivenessProbe:  application.LivenessProbe,
			ReadinessProbe: application.ReadinessProbe,
			Replicas:       limits[replica],
		},
		ReleaseName: args.ReleaseName,
		Namespace:   args.Namespace,
		Environment: args.Env,
		EnvVars:     envVars,
		Limits:      limits,
	}
	templates.Run(&tmlValues)
	fmt.Printf("Template generated for %s", app.Name)
}

func GetMixins(application *model.Application, args *model.Args) {
	context.TODO()
	for _, mxin := range application.Mixins {
		mixin := &model.Mixin{}
		mixinType := strings.Split(mxin, sep)
		name := mixinType[0]
		mType := mixinType[1]
		GetMixin(name, &mixin, args)
		for _, m := range mixin.Attributes {
			if mType == m["name"] {

			}
		}
	}
}

func GetResourceLimit(application *model.Application, args *model.Args) map[string]string {
	resourceLimit := make(map[string]string, 0)

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
			if val, ok := configs[replica]; ok {
				replicaLimit = val
			}
			resourceLimit[cpu] = cpuLimit
			resourceLimit[memory] = memLimit
			resourceLimit[replica] = replicaLimit
		}
	}
	return resourceLimit
}

func GenerateEnvVars(application *model.Application, args *model.Args) map[string]string {
	appEnvVars := make(map[string]string, 0)

	//Process app resources
	for _, appRes := range application.Resources {
		//elasticsearch-user:sit
		resDetails := strings.Split(appRes, sep)
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
					infra := strings.Split(resTemplate.Infra, sep)
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
	return appEnvVars
}

func addToEnvVars(name string, appEnvVars map[string]string, items map[string]string) {
	infraName := strings.ReplaceAll(name, "-", "_")
	for k, v := range items {
		key := strings.ToUpper(fmt.Sprintf("%s_%s", infraName, k))
		appEnvVars[key] = v
	}
}
