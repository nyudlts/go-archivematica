package go_archivematica

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type DeletionRequest struct {
	EventReason string    `json:"event_reason"`
	Pipeline    uuid.UUID `json:"pipeline"`
	UserID      string    `json:"user_id"`
	UserEmail   string    `json:"user_email"`
}

type Location struct {
	Description  string   `json:"description"`
	Enabled      bool     `json:"enabled"`
	Path         string   `json:"path"`
	Pipeline     []string `json:"pipeline"`
	Purpose      string
	Quota        interface{} `json:"quota"`
	RelativePath string      `json:"relative_path"`
	ResourceURI  string      `json:"resource_uri"`
	Space        string      `json:"space"`
	Used         int64       `json:"used"`
	UUID         uuid.UUID   `json:"uuid"`
}

type Locations struct {
	Meta    Meta       `json:"meta"`
	Objects []Location `json:"objects"`
}

type LocationBrowser struct {
	Directories []string `json:"directories"`
	Entries     []string `json:"entries"`
	Properties  map[string]map[string]int
}

type Meta struct {
	Limit      int `json:"limit"`
	Next       int `json:"next"`
	Offset     int `json:"offset"`
	Previous   int `json:"previous"`
	TotalCount int `json:"total_count"`
}

type Package struct {
	CurrentFullPath   string        `json:"current_full_path"`
	CurrentLocation   string        `json:"current_location"`
	CurrentPath       string        `json:"current_path"`
	Encrypted         bool          `json:"encrypted"`
	MiscAttributes    interface{}   `json:"misc_attributes"`
	OriginPipeline    string        `json:"origin_pipeline"`
	PackageType       string        `json:"package_type"`
	RelatedPackages   []string      `json:"related_packages"`
	Replicas          []interface{} `json:"replicas"`
	ReplicatedPackage []interface{} `json:"replicated_Package"`
	ResourceURI       string        `json:"resource_uri"`
	Size              int64         `json:"size"`
	Status            string        `json:"status"`
	UUID              uuid.UUID     `json:"uuid"`
}

func (p Package) GetPipelineUUID() (uuid.UUID, error) {
	urlSplit := strings.Split(p.OriginPipeline, "/")
	pipelineUUID, err := uuid.Parse(urlSplit[4])
	if err != nil {
		return uuid.UUID{}, err
	}

	return pipelineUUID, nil
}

func (p Package) GetPackageName() string {
	urisplit := strings.Split(p.CurrentPath, "/")
	return urisplit[len(urisplit)-1]
}

func (p Package) MarshalPack() (string, error) {
	packJson, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(packJson), nil
}

type Packages struct {
	Meta    PackageMeta `json:"meta"`
	Objects []Package   `json:"objects"`
}

type PackageMeta struct {
	Limit      int    `json:"limit"`
	Next       string `json:"next"`
	Offset     int    `json:"offset"`
	Previous   string `json:"previous"`
	TotalCount int    `json:"total_count"`
}

func (p PackageMeta) String() string {
	return fmt.Sprintf("Limit: %d, Next %s, Offset %d, Previous %s, TotalCount: %d", p.Limit, p.Next, p.Offset, p.Previous, p.TotalCount)
}

type Pipeline struct {
	Description string    `json:"description"`
	RemoteName  string    `json:"remote_name"`
	ResourceURI string    `json:"resource_uri"`
	UUID        uuid.UUID `json:"uuid"`
}

type Pipelines struct {
	Meta    Meta       `json:"meta"`
	Objects []Pipeline `json:"objects"`
}

type Space struct {
	AccessProtocol string      `json:"access_protocol"`
	ID             int         `json:"id"`
	LastVerified   interface{} `json:"last_verified"`
	Path           string      `json:"path"`
	ResourceURI    string      `json:"resource_uri"`
	Size           interface{} `json:"size"`
	Space          uuid.UUID   `json:"space"`
	Used           int         `json:"used"`
	UUID           uuid.UUID   `json:"uuid"`
	Verified       bool        `json:"verified"`
}

type Spaces struct {
	Meta    Meta    `json:"meta"`
	Objects []Space `json:"objects"`
}
