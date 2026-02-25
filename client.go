package go_archivematica

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type AMEnvironments map[string]AMEnvironment

type AMClient struct {
	Client               http.Client
	StagingLoc           string
	AIPStoreLocation     string
	DashboardConfig      AMConfig `json:"dashboard"`
	StorageServiceConfig AMConfig `json:"storage_service"`
}

type AMConfig struct {
	URL       string `json:"url"`
	Key       string `json:"api_key"`
	UserID    int    `json:"user_id"`
	UserEmail string `json:"user_email"`
	Username  string `json:"username"`
}

type StorageServiceConfig struct {
	URL string `json:"url"`
}

type AMEnvironment struct {
	StagingLoc       string   `json:"staging_location"`
	AIPStoreLocation string   `json:"aipsstore_location"`
	Dashboard        AMConfig `json:"dashboard"`
	StorageService   AMConfig `json:"storage_service"`
}

func NewAMClient(config string, environment string, timeout int) (*AMClient, error) {

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

	var amEnvs AMEnvironments
	if err := json.Unmarshal(configBytes, &amEnvs); err != nil {
		fmt.Println("Error parsing config file:", err)
		return nil, err
	}

	fmt.Println(amEnvs)

	for envName, env := range amEnvs {
		if envName == environment {
			return &AMClient{
				Client:               *nclient,
				StagingLoc:           env.StagingLoc,
				AIPStoreLocation:     env.AIPStoreLocation,
				DashboardConfig:      env.Dashboard,
				StorageServiceConfig: env.StorageService,
			}, nil
		}
	}

	return nil, fmt.Errorf("environment '%s' not found in config", environment)

}
