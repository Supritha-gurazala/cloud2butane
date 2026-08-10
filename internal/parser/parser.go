package parser

import (
	"gopkg.in/yaml.v3"

	"github.com/Supritha-gurazala/cloud2butane/internal/cloudconfig"
)

// Parse converts cloud-config YAML bytes into a Config struct.
func Parse(data []byte) (cloudconfig.Config, error) {
	var cfg cloudconfig.Config

	err := yaml.Unmarshal(data, &cfg)
	if err != nil {
		return cloudconfig.Config{}, err
	}
	return cfg, nil
}
