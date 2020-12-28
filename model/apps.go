package model

type Args struct {
	ManifestDir string
	Namespace   string
	Env         string
	ReleaseName string
}

type Application struct {
	Name           string                   `yaml:"name"`
	Kind           string                   `yaml:"kind"`
	LivenessProbe  string                   `yaml:"livenessProbe"`
	ReadinessProbe string                   `yaml:"readinessProbe"`
	Annotations    map[string]string        `yaml:"annotations"`
	Resources      []string                 `yaml:"resources"`
	Capabilities   []string                 `yaml:"capabilities"`
	Mixins         []string                 `yaml:"mixins"`
	Template       []AppTemplate            `yaml:"template"`
	NodeSelector   map[string]string        `yaml:"nodeSelector"`
	VolumeMounts   []Mount                  `yaml:"volumeMounts"`
	Volumes        []map[string]interface{} `yaml:"volumes"`
	ConfigMaps     []map[string]string      `yaml:"configMaps"`
	Service        Service                  `yaml:"service"`
}

type Service struct {
	Name       string `yaml:"name"`
	Type       string `yaml:"type"`
	Port       int    `yaml:"port"`
	TargetPort int    `yaml:"targetPort"`
}

type KeyValue struct {
	Key   string `yaml:key`
	Value string `yaml:value`
}

type Mount struct {
	Name      string `yaml:"name"`
	MountPath string `yaml:"mountPath"`
}

type AppTemplate struct {
	Name    string            `yaml:"name"`
	Replica int               `yaml:"replica"`
	Config  map[string]string `yaml:"config"`
}
