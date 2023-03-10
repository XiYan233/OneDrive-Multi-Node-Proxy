package utils

import (
	"encoding/json"
	"os"
)

type UrlConfig struct {
	URL string `json:"url"`
}

type Config struct {
	URLConfig struct {
		CnMobile UrlConfig `json:"cn_mobile"`
		CnTele   UrlConfig `json:"cn_tele"`
		CnUni    UrlConfig `json:"cn_uni"`
		Other    UrlConfig `json:"other"`
	} `json:"url_config"`
}

func LoadConfig(configPath string) (*Config, error) {
	configFile, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	var config Config
	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
