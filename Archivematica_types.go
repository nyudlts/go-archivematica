package go_am

import (
	"fmt"

	"github.com/google/uuid"
)

// types
type IngestStatus struct {
	Status       string    `json:"status"`
	Name         string    `json:"name"`
	Microservice string    `json:"microservice"`
	Directory    string    `json:"directory"`
	Path         string    `json:"path"`
	Message      string    `json:"message"`
	Type         string    `json:"type"`
	UUID         uuid.UUID `json:"uuid"`
}

type TransferStatus struct {
	Status       string    `json:"status"`
	Name         string    `json:"name"`
	SIPUUID      string    `json:"sip_uuid"`
	Microservice string    `json:"microservice"`
	Directory    string    `json:"directory"`
	Path         string    `json:"path"`
	Message      string    `json:"message"`
	Type         string    `json:"type"`
	UUID         uuid.UUID `json:"uuid"`
}

type UnapprovedTransfers struct {
	Message string               `json:"message"`
	Results []UnapprovedTransfer `json:"Results"`
}

type UnapprovedTransfer struct {
	Directory string    `json:"directory"`
	Type      string    `json:"type"`
	UUID      uuid.UUID `json:"uuid"`
}

type UUIDList struct {
	Message string   `json:"message"`
	Results []string `json:"results"`
}

type WaitingIngests struct {
	Message string          `json:"message"`
	Results []WaitingIngest `json:"results"`
}

type WaitingIngest struct {
	Microservice string    `json:"microservice"`
	SIPName      string    `json:"sip_name"`
	SIPDirectory string    `json:"sip_directory"`
	SIPUUID      uuid.UUID `json:"sip_uuid"`
}

// string formatters
func (u UUIDList) String() string {
	f := fmt.Sprintf("Message: %s\nCompleted Transfers:\n", u.Message)
	for _, s := range u.Results {
		f = f + fmt.Sprintf("\t%s\n", s)
	}
	return f
}

func (is IngestStatus) String() string {
	out := fmt.Sprintf("Message: %s\n", is.Message)
	out = out + "Transfer Status:\n"
	out = out + fmt.Sprintf("\tname: %s\n", is.Name)
	out = out + fmt.Sprintf("\tmicroservice: %s\n", is.Microservice)
	out = out + fmt.Sprintf("\tdirectory: %s\n", is.Directory)
	out = out + fmt.Sprintf("\tpath: %s\n", is.Path)
	out = out + fmt.Sprintf("\ttype: %s\n", is.Type)
	out = out + fmt.Sprintf("\tuuid: %s\n", is.UUID)
	return out
}

func (ts TransferStatus) String() string {
	out := fmt.Sprintf("Message: %s\n", ts.Message)
	out = out + "Transfer Status:\n"
	out = out + fmt.Sprintf("\tname: %s\n", ts.Name)
	out = out + fmt.Sprintf("\tsip_uuid: %s\n", ts.SIPUUID)
	out = out + fmt.Sprintf("\tmicroservice: %s\n", ts.Microservice)
	out = out + fmt.Sprintf("\tdirectory: %s\n", ts.Directory)
	out = out + fmt.Sprintf("\tpath: %s\n", ts.Path)
	out = out + fmt.Sprintf("\ttype: %s\n", ts.Type)
	out = out + fmt.Sprintf("\tuuid: %s\n", ts.UUID)
	return out
}
