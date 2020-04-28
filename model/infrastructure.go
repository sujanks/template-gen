package model

type Infrastructure struct {
	ApiVersion string            `yaml:"apiVersion"`
	Kind       string            `yaml:"kind"`
	Metadata   map[string]string `yaml:"metadata"`
	Spec       InfraSpec         `yaml:"spec"`
}

type InfraSpec struct {
	Template []InfraTemplate `yaml:"template"`
}

type InfraTemplate struct {
	Name       string            `yaml:"name"`
	Attributes map[string]string `yaml:"attributes"`
}
