package model

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func NewArticle(title, description, content string, witerUserId int, articleType ArticleTypes) (*Article, error) {
	return &Article{
		Title:       title,
		Description: description,
		Content:     content,
		WiterUserId: witerUserId,
		WrittenAt:   time.Now(),
		Type:        articleType,
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

	value, ok := toID[j]
	if !ok {
		return fmt.Errorf("Error value %s is out of enum", j)
	}
	*s = value
	return nil
}
