package configs

import (
	"encoding/json"
	"os"
)

type Config struct {
	HOST string `json:"HOST"`
	PORT string `json:"PORT"`

	ADMIN_LOGIN string `json:"ADMIN_LOGIN"`

	TIME_COOKIE int `json:"TIME_COOKIE"`

	MONGO_DB string `json:"MONGO_DB"`

	STAT_PATH string `json:"STAT_PATH"`
}

func GetConfig(path string) (*Config, error) {
	cfg := &Config{}
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
