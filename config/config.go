package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadFromFile(path string) (*Config, error) {
	cfgFile, err := os.Open(path)
	if err != nil {
		fmt.Println("Failed to open config file: ", err)
		return nil, err
	}
	defer cfgFile.Close()

	var cfg *Config
	err = json.NewDecoder(cfgFile).Decode(&cfg)
	if err != nil {
		fmt.Println("Failed to read config file: ", err)
		return nil, err
	}

	return cfg, nil
}

type Config struct {
	Watchdog WatchdogConfig `json:"watchdog"`
}

type WatchdogConfig struct {
	InetChecker InetCheckerConfig `json:"inet_checker"`
	VPNChecker  VPNCheckerConfig  `json:"vpn_checker"`
}

type InetCheckerConfig struct {
	Enabled     bool   `json:"enabled"`
	URL         string `json:"url"`
	LongPeriod  int    `json:"long_period"`
	ShortPeriod int    `json:"short_period"`
}

type VPNCheckerConfig struct {
	Enabled    bool   `json:"enabled"`
	SvcName    string `json:"svc_name"`
	LongPeriod int    `json:"long_period"`
}
