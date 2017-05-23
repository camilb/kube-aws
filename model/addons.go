package model

type Addons struct {
	Rescheduler       Rescheduler              `yaml:"rescheduler"`
	ClusterAutoscaler ClusterAutoscalerSupport `yaml:"clusterAutoscaler,omitempty"`
	UnknownKeys       `yaml:",inline"`
}

type ClusterAutoscalerSupport struct {
	Enabled     bool `yaml:"enabled"`
	UnknownKeys `yaml:",inline"`
}

type Rescheduler struct {
	Enabled     bool `yaml:"enabled"`
	UnknownKeys `yaml:",inline"`
}
