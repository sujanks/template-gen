package templates

import (
	"fmt"
	"log"
	"os"
)

const (
	DEPLOY    = "DeploymentTemplate"
	CRON      = "CronJobTemplate"
	STATEFUL  = "StatefulSetTemplate"
	CONFIGMAP = "ConfigMapTemplate"
	SERVICE   = "ServiceTemplate"
)

func Run(releaseTemplate *ReleaseTemplate) {
	templateMap := map[string]string{"deployment": DEPLOY,
		"cronjob":     CRON,
		"statefulset": STATEFUL}

	dir := "tmp/templates"
	os.MkdirAll(dir, os.ModePerm)
	os.Chdir(dir)

	for _, application := range releaseTemplate.Application {
		fmt.Println("Generating template for: ", application.Name)

		templateArray := make([]string, 0)
		kind := application.Kind
		if application.Kind == "" {
			templateArray = append(templateArray, DEPLOY)
		} else {
			templateArray = append(templateArray, templateMap[kind])
		}
		//Check for configmap
		if len(application.ConfigMaps) > 0 {
			templateArray = append(templateArray, CONFIGMAP)
		}

		if application.Service.Name != "" {
			templateArray = append(templateArray, SERVICE)
		}

		for _, tName := range templateArray {
			tmpl := LoadTemplates(tName, &application)

			file, er := os.Create(tmpl.Name())
			if er != nil {
				log.Fatal("error ", er)
			}

			err := tmpl.Execute(file, &application)
			if err != nil {
				log.Fatal("error ", err)
			}
		}
	}
}
