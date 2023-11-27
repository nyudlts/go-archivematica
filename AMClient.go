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
	AMHost   string
	AMAPIKey string
	SSHost   string
	SSAPIKey string
	Client   *http.Client
}

func (a *AMClient) String() string {
	return fmt.Sprintf("Username\t%s\nAMHost\t%s", a.Username, a.AMHost)
}

type AMEnvironment struct {
	AMURL    string `yaml:"am_url"`
	SSURL    string `yaml:"ss_url"`
	Username string `yaml:"username"`
	AMAPIKey string `yaml:"am_api_key"`
	SSAPIKey string `yaml:"ss_api_key"`
}

func NewAMClient(config string, timeout int) (*AMClient, error) {
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

	var amEnv AMEnvironment
	if err := yaml.Unmarshal(configBytes, &amEnv); err != nil {
		return nil, err
	}

	return &AMClient{amEnv.Username, amEnv.AMURL, amEnv.SSURL, amEnv.AMAPIKey, amEnv.SSAPIKey, nclient}, nil
}
