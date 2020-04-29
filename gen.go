package main

import (
	"flag"
	"fmt"
	"github.com/kube-sailmaker/template-gen/model"
	"github.com/kube-sailmaker/template-gen/task"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("arguments missing")
		os.Exit(1)
	}

	dirPtr := flag.String("d", "", "Manifest director. (Required)")
	nsPrt := flag.String("n", "", "Namespace. (Required)")
	envPrt := flag.String("e", "", "Environment. (Required)")
	relPrt := flag.String("r", "", "Release name. (Required)")
	flag.Parse()

	args := &model.Args{
		ManifestDir: *dirPtr,
		Namespace:   *nsPrt,
		Env:         *envPrt,
		ReleaseName: *relPrt,
	}

	task.GenerateTemplates(args)
}
