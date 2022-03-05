package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

type Publication struct {
	Id          int
	Title       string
	Description string
	Content     string
	WiterUserId int
	WrittenAt   time.Time
	Type        PublicationTypes
}

func NewPublication(title, description, content string, witerUserId int, publicationType PublicationTypes) (*Publication, error) {
	return &Publication{
		Title:       title,
		Description: description,
		Content:     content,
		WiterUserId: witerUserId,
		WrittenAt:   time.Now(),
		Type:        publicationType,
	}, nil
}

type PublicationTypes string

const (
	Configuration PublicationTypes = "Configuration"
	Device                         = "Device"
	Handbook                       = "Handbook"
)

func (s PublicationTypes) String() string {
	return toString[s]
}

var toString = map[PublicationTypes]string{
	Configuration: "Configuration",
	Device:        "Device",
	Handbook:      "Handbook",
}
var toID = map[string]PublicationTypes{
	"Configuration": Configuration,
	"Device":        Device,
	"Handbook":      Handbook,
}

func (s PublicationTypes) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *PublicationTypes) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}

	value, ok := toID[j]
	if !ok {
		return fmt.Errorf("Error value %s is out of enum", j)
	}
	*s = value
	return nil
}
