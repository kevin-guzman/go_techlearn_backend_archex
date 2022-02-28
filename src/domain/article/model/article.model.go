package model

import (
	"bytes"
	"encoding/json"
	"time"
)

type Article struct {
	Id          int
	Title       string
	Description string
	Content     string
	WiterUserId int
	WrittenAt   time.Time
	Type        ArticleTypes
}

func NewArticle(title, description, content string, witerUserId int) (*Article, error) {
	return &Article{
		Title:       title,
		Description: description,
		Content:     content,
		WiterUserId: witerUserId,
		WrittenAt:   time.Now(),
	}, nil
}

type ArticleTypes string

const (
	Configuration ArticleTypes = "Configuration"
	Device                     = "Device"
	Handbook                   = "Handbook"
)

func (s ArticleTypes) String() string {
	return toString[s]
}

var toString = map[ArticleTypes]string{
	Configuration: "Configuration",
	Device:        "Device",
	Handbook:      "Handbook",
}
var toID = map[string]ArticleTypes{
	"Configuration": Configuration,
	"Device":        Device,
	"Handbook":      Handbook,
}

func (s ArticleTypes) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *ArticleTypes) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Configuration' in this case.
	*s = toID[j]
	return nil
}
