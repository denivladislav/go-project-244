package parsers

import (
	"encoding/json"
	"errors"
	"fmt"
)

var ErrInvalidJSON = errors.New("invalid JSON")

// ParseJSON parses JSON-encoded data to key-value map.
func ParseJSON(data []byte) (map[string]any, error) {
	var parsedContent map[string]any

	err := json.Unmarshal(data, &parsedContent)
	if err != nil {
		return nil, fmt.Errorf(`%w: %w`, ErrInvalidJSON, err)
	}

	return parsedContent, nil
}
