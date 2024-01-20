package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type (
	ServiceItem struct {
		Addr  string `json:"addr" yaml:"addr"`
		Proxy string `json:"proxy" yaml:"proxy"`
		Flag  bool   `json:"flag" yaml:"flag"`
	}

	ApiCnf struct {
		Addr string `json:"addr" yaml:"addr"`
		Key  string `json:"key" yaml:"key"`
	}
	ProxyCnf struct {
		Flag bool   `json:"flag" yaml:"flag"`
		Addr string `json:"addr" yaml:"addr"`
	}

	Config struct {
		Proxy    ProxyCnf               `json:"proxy" yaml:"proxy"`
		Services map[string]ServiceItem `json:"services" yaml:"services"`
		Api      ApiCnf                 `json:"api" yaml:"api"`
	}
)

func ReadConfig() (Config, error) {
	var config Config
	curDir, err := GetCurDir()
	if err != nil {
		return config, err
	}
	configFile := filepath.Join(curDir, "conf", "config.yaml")
	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(yamlFile, &config)
	fmt.Printf("service: %v, err: %v\n", config.Services, err)
	fmt.Printf("proxy: %v, err: %v\n", config.Proxy, err)
	fmt.Printf("api: %v, err: %v\n", config.Api, err)

	return config, err
}

func GetCurDir() (string, error) {
	dir, err := filepath.Abs(filepath.Dir("./"))
	if err != nil {
		return "", err
	}
	return dir, err
}
