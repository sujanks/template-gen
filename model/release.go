package model

type Release struct {
	App []App `yaml:"apps"`
}

type App struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}
