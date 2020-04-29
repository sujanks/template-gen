package templates

import (
	"fmt"
	"log"
	"os"
)

func Run(releaseTemplate *ReleaseTemplate) {
	tmplArray := []string{"ServiceTemplate", "DeploymentTemplate", "ServiceAccountTemplate"}

	dir := "tmp/templates"
	os.MkdirAll(dir, os.ModePerm)
	os.Chdir(dir)

	for _, application := range releaseTemplate.Application {
		fmt.Println("Generating template for: ", application.Name)
		for _, tName := range tmplArray {
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
