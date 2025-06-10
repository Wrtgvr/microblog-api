package config

import (
	"fmt"
	"os"
	"time"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	Port    int           `mapstructure:"port" yaml:"port"`
	Timeout time.Duration `mapstructure:"timeout"`
	Debug   bool          `mapstructure:"debug"`
}

func loadServerConfig() (*ServerConfig, error) {
	b, err := os.ReadFile("config/config_dev.yaml")
	if err != nil {
		return nil, err
	}

	var rawCfg map[string]any
	if err := yaml.Unmarshal(b, &rawCfg); err != nil {
		return nil, err
	}

	serverRaw, ok := rawCfg["server"].(map[string]any)
	if !ok {
		return nil, fmt.Errorf("cannot parse server config: field 'server' not found or invalid")
	}

	var cfg ServerConfig
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: mapstructure.StringToTimeDurationHookFunc(),
		Result:     &cfg,
	})
	if err != nil {
		return nil, err
	}

	if err := decoder.Decode(&serverRaw); err != nil {
		return nil, err
	}

	fmt.Printf("%+v\n", cfg)

	return &cfg, nil
}
