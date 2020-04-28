package templates

import (
	"log"
	"os"
)

func Run(values *TemplateValues) {
	tmplArray := []string{"ChartTemplate", "ServiceTemplate", "DeploymentTemplate", "ServiceAccountTemplate"}

	dir := "tmp/templates"
	os.MkdirAll(dir, os.ModePerm)
	os.Chdir(dir)

	for _, tName := range tmplArray {
		tmpl := LoadTemplates(tName)

		file, er := os.Create(tmpl.Name())
		if er != nil {
			log.Fatal("error ", er)
		}

		err := tmpl.Execute(file, values)
		if err != nil {
			log.Fatal("error ", err)
		}
	}

}
