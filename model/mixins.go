package model

type MixinList struct {
	Mixin []Mixin `yaml:"mixin"`
}

type Mixin struct {
	Name             string            `yaml:"name"`
	Cpu              string            `yaml:"cpu"`
	Memory           string            `yaml:"memory"`
	Replicas         string            `yaml:"replicas"`
	ResourceStrategy string            `yaml:"resource-limit-strategy"`
	Env              map[string]string `yaml:env`
	Cmd              []string          `yaml:"cmd"`
	Entrypoint       []string          `yaml:"entrypoint"`
}
