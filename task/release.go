package task

import "github.com/kube-sailmaker/template-gen/model"

func ProcessRelease(apps *model.Release, args *model.Args) {
	for _, app := range apps.App {
		ProcessApplication(&app, args)
	}
}
