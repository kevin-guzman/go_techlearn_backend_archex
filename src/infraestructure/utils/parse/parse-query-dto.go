package parse

import (
	"encoding/json"
	"net/url"
)

func ParseQueryToDto(query url.Values, instancedDto interface{}) error {
	if bytes, err := json.Marshal(query); err != nil {
		return err
	} else {
		if err = json.Unmarshal(bytes, &instancedDto); err != nil {
			return err
		}
		return nil
	}
}
