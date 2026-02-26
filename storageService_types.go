package go_archivematica

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type CreateSpaceRequest struct {
	AccessProtocol string `json:"access_protocol"`
	Path           string `json:"path"`
	Size           int    `json:"size"`
	StagingPath    string `json:"staging_path,omitempty"`
}

type CreateSpaceResponse map[string]Space

type CreatePipelineRequest struct {
	UUID                   string `json:"uuid"`
	Description            string `json:"description"`
	APIKey                 string `json:"api_key"`
	APIUsername            string `json:"api_username"`
	CreateDefaultLocations bool   `json:"create_default_locations"`
	SharedPath             string `json:"shared_path,omitempty"`
	RemoteName             string `json:"remote_name,omitempty"`
}

type DeletionRequest struct {
	EventReason string    `json:"event_reason"`
	Pipeline    uuid.UUID `json:"pipeline"`
	UserID      string    `json:"user_id"`
	UserEmail   string    `json:"user_email"`
}

type Endpoints struct {
	Async    EndpointValues `json:"async"`
	File     EndpointValues `json:"file"`
	Location EndpointValues `json:"location"`
	Pipeline EndpointValues `json:"pipeline"`
	Space    EndpointValues `json:"space"`
}

type EndpointValues struct {
	ListEndpoint string `json:"list_endpoint"`
	Schema       string `json:"schema"`
}

func (e EndpointValues) String() string {
	return fmt.Sprintf("  ListEndpoint: %s\n  Schema: %s\n", e.ListEndpoint, e.Schema)
}

func (e Endpoints) String() string {
	return fmt.Sprintf("Async:\n%v\nFile:\n%v\nLocation:\n%v\nPipeline:\n%v\nSpace:\n%v\n", e.Async, e.File, e.Location, e.Pipeline, e.Space)
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

func (m Meta) String() string {
	return fmt.Sprintf("  Limit: %d\n  Next: %d\n  Offset: %d\n  Previous: %d\n  TotalCount: %d", m.Limit, m.Next, m.Offset, m.Previous, m.TotalCount)
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

func (p Pipeline) String() string {
	return fmt.Sprintf("%s\n  RemoteName: %s\n  ResourceURI: %s\n  UUID: %s\n", p.Description, p.RemoteName, p.ResourceURI, p.UUID)
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

func (s Space) String() string {
	return fmt.Sprintf("  ID: %d\n    AccessProtocol: %s\n    Path: %s\n    ResourceURI: %s\n    Size: %v\n    Used: %d\n    UUID: %s\n    Verified: %t\n", s.ID, s.AccessProtocol, s.Path, s.ResourceURI, s.Size, s.Used, s.UUID, s.Verified)
}

type Spaces struct {
	Meta    Meta    `json:"meta"`
	Objects []Space `json:"objects"`
}

func (s Spaces) String() string {
	objects := ""
	for _, space := range s.Objects {
		objects += fmt.Sprintf("%s\n", space)
	}
	return fmt.Sprintf("Meta:\n%s\nObjects:\n%s", s.Meta, objects)
}

type SpaceBrowser struct {
	Directories []string            `json:"directories"`
	Entries     []string            `json:"entries"`
	Properties  map[string]Property `json:"properties"`
}

type Property struct {
	ObjectCount FlexibleValue `json:"object count,omitempty"`
	Size        *int64        `json:"size,omitempty"`
}

type FlexibleValue struct {
	IntValue    *int64
	StringValue *string
}

func (f *FlexibleValue) UnmarshalJSON(data []byte) error {
	// Try integer
	var i int64
	if err := json.Unmarshal(data, &i); err == nil {
		f.IntValue = &i
		return nil
	}

	// Try string
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		f.StringValue = &s
		return nil
	}

	return fmt.Errorf("FlexibleValue: unsupported type: %s", string(data))
}

func (s SpaceBrowser) String() string {
	directories := ""
	for _, directory := range s.Directories {
		directories += fmt.Sprintf("  %s\n", directory)
	}

	entries := ""
	for _, entry := range s.Entries {
		entries += fmt.Sprintf("  %s\n", entry)
	}

	properties := ""
	for key, value := range s.Properties {
		properties += fmt.Sprintf("  %s:\n", key)
		if value.ObjectCount.IntValue != nil {
			properties += fmt.Sprintf("    ObjectCount: %d\n", *value.ObjectCount.IntValue)
		} else if value.ObjectCount.StringValue != nil {
			properties += fmt.Sprintf("    ObjectCount: %s\n", *value.ObjectCount.StringValue)
		}
		if value.Size != nil {
			properties += fmt.Sprintf("    Size: %d\n", *value.Size)
		}
	}

	return fmt.Sprintf("directories:\n%sentries:\n%sproperties:\n%s", directories, entries, properties)
}

type SchemaField struct {
	Blank       bool        `json:"blank"`
	Default     interface{} `json:"default"`
	HelpText    string      `json:"help_text"`
	Nullable    bool        `json:"nullable"`
	PrimaryKey  bool        `json:"primary_key"`
	ReadOnly    bool        `json:"readonly"`
	Type        string      `json:"type"`
	Unique      bool        `json:"unique"`
	VerboseName string      `json:"verbose_name"`
}

type Schema struct {
	AllowedDetailHTTPMethods []string               `json:"allowed_detail_http_methods"`
	AllowedListHTTPMethods   []string               `json:"allowed_list_http_methods"`
	DefaultFormat            string                 `json:"default_format"`
	DefaultLimit             int                    `json:"default_limit"`
	Fields                   map[string]SchemaField `json:"fields"`
	Filtering                map[string]int         `json:"filtering"`
}
