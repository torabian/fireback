package fireback

import "gopkg.in/yaml.v2"

func GetConfig() Config {
	return config
}

func (x *Config) Yaml() string {
	if x != nil {
		str, _ := yaml.Marshal(x)
		return (string(str))
	}
	return ""
}
