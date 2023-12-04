package amatica

import "github.com/google/uuid"

type Meta struct {
	Limit      int `json:"limit"`
	Next       int `json:"next"`
	Offset     int `json:"offset"`
	Previous   int `json:"previous"`
	TotalCount int `json:"total_count"`
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

/*
{
    "directories": [
        "RUlDQVI=",
        "ZmFsZXNfbXNzNDIwX2VsZWN0cm9uaWMtcmVjb3Jkcy1iYXRjaC0x",
        "ZmFsZXNfbXNzNDIwX2VsZWN0cm9uaWMtcmVjb3Jkcy1iYXRjaC0y",
        "bnl1YXJjaGl2ZXNfcmczN181MV9lbGVjdHJvbmljLXJlY29yZHMtYmF0Y2gtMg=="
    ],
    "entries": [
        "QkFWQzEwMDQwMTVfUkQzX09zY2FySS5tb3Y=",
        "RUlDQVI=",
        "ZmFsZXNfbXNzNDIwX2VsZWN0cm9uaWMtcmVjb3Jkcy1iYXRjaC0x",
        "ZmFsZXNfbXNzNDIwX2VsZWN0cm9uaWMtcmVjb3Jkcy1iYXRjaC0y",
        "bnl1YXJjaGl2ZXNfcmczN181MV9lbGVjdHJvbmljLXJlY29yZHMtYmF0Y2gtMg=="
    ],
    "properties": {
        "QkFWQzEwMDQwMTVfUkQzX09zY2FySS5tb3Y=": {
            "size": 56349350023
        },
        "RUlDQVI=": {
            "object count": 1
        },
        "ZmFsZXNfbXNzNDIwX2VsZWN0cm9uaWMtcmVjb3Jkcy1iYXRjaC0x": {
            "object count": 129
        },
        "ZmFsZXNfbXNzNDIwX2VsZWN0cm9uaWMtcmVjb3Jkcy1iYXRjaC0y": {
            "object count": 1491
        },
        "bnl1YXJjaGl2ZXNfcmczN181MV9lbGVjdHJvbmljLXJlY29yZHMtYmF0Y2gtMg==": {
            "object count": 9
        }
    }
}
*/
