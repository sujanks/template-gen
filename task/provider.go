package task

import (
	"fmt"
	"github.com/kube-sailmaker/template-gen/functions"
	"github.com/kube-sailmaker/template-gen/model"
)

const (
	infraManifest = "%s/provider/infrastructure/%s.yaml"
	mixinManifest = "%s/provider/mixins/%s.yaml"
	resourceManifest = "%s/provider/resources/%s.yaml"
)

func GetInfrastructure(name string, t interface{}, args *model.Args) {
	file := fmt.Sprintf(infraManifest,  args.ManifestDir, name)
	functions.UnmarshalFile(file, t)
}

func GetMixin(name string, t interface{}, args *model.Args) {
	file := fmt.Sprintf(mixinManifest,  args.ManifestDir, )
	functions.UnmarshalFile(file, t)
}

func GetResource(name string, t interface{}, args *model.Args) {
	file := fmt.Sprintf(resourceManifest, args.ManifestDir, name)
	functions.UnmarshalFile(file, t)
}
