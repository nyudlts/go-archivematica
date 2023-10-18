package go_am

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type AMClient struct {
	Username string
	Host     string
	APIKey   string
	Client   *http.Client
}

type AMEnvironment struct {
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	APIKey   string `yaml:"apikey"`
}

func NewAMClient(config string, env string, timeout int) (*AMClient, error) {
	transport := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    time.Duration(timeout) * time.Second,
		DisableCompression: true,
	}

	nclient := &http.Client{
		Transport: transport,
	}

	configBytes, err := os.ReadFile(config)
	if err != nil {
		return nil, err
	}

	envMap := map[string]AMEnvironment{}

	if err := yaml.Unmarshal(configBytes, &envMap); err != nil {
		return nil, err
	}

	var amEnv AMEnvironment
	for k, v := range envMap {
		if k == env {
			amEnv = v
			break
		}
		return nil, fmt.Errorf("environemnt %s does not exist", env)
	}

	return &AMClient{amEnv.Username, amEnv.URL, amEnv.APIKey, nclient}, nil
}
