package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// LoadConfig liest die YAML-Konfigurationsdatei und parst sie in ein Config-Struct.
func loadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("fehler beim Lesen der Datei: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("fehler beim Parsen der YAML-Daten: %w", err)
	}
	return &cfg, nil
}

func GetConfig() (*Config, error) {
	config, err := loadConfig("misc/sample-config.yml")
	if err != nil {
		return nil, fmt.Errorf("unable to read config: %w", err)
	}

	return config, nil

}
