package task

import (
	"fmt"
	"github.com/kube-sailmaker/template-gen/functions"
	"github.com/kube-sailmaker/template-gen/model"
	templates "github.com/kube-sailmaker/template-gen/template"
)

func GenerateTemplates(args *model.Args) {
	relManifest := fmt.Sprintf("%s/user/releases/%s/%s.yaml", args.ManifestDir, args.Env, args.Namespace)
	fmt.Println("release manifest from directory ", relManifest)
	release := model.Release{}
	functions.UnmarshalFile(relManifest, &release)
	appTemplate := make([]templates.Application, 0)
	for _, app := range release.App {
		application := ProcessApplication(&app, args)
		appTemplate = append(appTemplate, *application)
	}

	releaseTemplate := templates.ReleaseTemplate{Application: appTemplate}
	templates.Run(&releaseTemplate)
}
