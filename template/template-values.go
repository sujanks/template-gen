package templates

type TemplateValues struct {
	Application Application
	ReleaseName string
	Namespace   string
	Environment string
	EnvVars     map[string]string
	Limits      map[string]string
}

type Application struct {
	Name           string
	Tag            string
	Replicas       string
	LivenessProbe  string
	ReadinessProbe string
}
