package config

import "github.com/caarlos0/env/v11"

func Load() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(&cfg.Database); err != nil {
		return nil, err
	}
	if err := env.Parse(&cfg.JWT); err != nil {
		return nil, err
	}
	if err := env.Parse(&cfg.Server); err != nil {
		return nil, err
	}
	return cfg, nil
}
