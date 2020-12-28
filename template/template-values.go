package templates

type ReleaseTemplate struct {
	Namespace   string
	Environment string
	Application []Application
}

type Application struct {
	ReleaseName    string
	Name           string
	Tag            string
	Kind           string
	Annotations    map[string]string
	Replicas       string
	LivenessProbe  string
	ReadinessProbe string
	EnvVars        map[string]string
	Limits         map[string]string
	Command        []string
	Entrypoint     []string
	NodeSelector   map[string]string
	VolumeMounts   []Mount
	Volumes        []map[string]interface{}
	ConfigMaps     []map[string]interface{}
	Service        Service
}

type Service struct {
	Name       string
	Type       string
	Port       int
	TargetPort int
}

type KeyValue struct {
	Key   string
	Value string
}

type Mount struct {
	Name      string
	MountPath string
}
