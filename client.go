package amatica

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type AMClient struct {
	Username    string
	AMHost      string
	AMAPIKey    string
	SSHost      string
	SSAPIKey    string
	Client      *http.Client
	StagingLoc  string
	SSUserID    int
	SSUserEmail string
}

func (a *AMClient) String() string {
	return fmt.Sprintf("\nUsername\t%s\nAMHost\t\t%s\nSSHost\t\t%s", a.Username, a.AMHost, a.SSHost)
}

type AMEnvironment struct {
	AMURL       string `yaml:"am_url"`
	SSURL       string `yaml:"ss_url"`
	Username    string `yaml:"username"`
	AMAPIKey    string `yaml:"am_api_key"`
	SSAPIKey    string `yaml:"ss_api_key"`
	StagingLoc  string `yaml:"staging_location"`
	SSUserID    int    `yaml:"ss_user_id"`
	SSUserEmail string `yaml:"ss_user_email"`
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

	return &AMClient{
		Username:    amEnv.Username,
		AMHost:      amEnv.AMURL,
		SSHost:      amEnv.SSURL,
		AMAPIKey:    amEnv.AMAPIKey,
		SSAPIKey:    amEnv.SSAPIKey,
		Client:      nclient,
		StagingLoc:  amEnv.StagingLoc,
		SSUserID:    amEnv.SSUserID,
		SSUserEmail: amEnv.SSUserEmail}, nil
}
