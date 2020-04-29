package model

type Args struct {
	ManifestDir string
	Namespace   string
	Env         string
	ReleaseName string
}

type Application struct {
	Name           string            `yaml:"name"`
	LivenessProbe  string            `yaml:"liveness_probe"`
	ReadinessProbe string            `yaml:"readiness_probe"`
	Labels         map[string]string `yaml:"labels"`
	Resources      []string          `yaml:"resources"`
	Capabilities   []string          `yaml:"capabilities"`
	Mixins         []string          `yaml:"mixins"`
	Template       []AppTemplate     `yaml:"template"`
}

type AppTemplate struct {
	Name    string            `yaml:"name"`
	Replica int               `yaml:"replica"`
	Config  map[string]string `yaml:"config"`
}
