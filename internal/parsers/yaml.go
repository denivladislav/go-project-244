package parsers

import (
	"errors"
	"fmt"

	"gopkg.in/yaml.v3"
)

var ErrInvalidYAML = errors.New("invalid YAML")

// ParseYAML parses YAML-encoded data to key-value map.
func ParseYAML(data []byte) (map[string]any, error) {
	var parsedContent map[string]any

	err := yaml.Unmarshal(data, &parsedContent)
	if err != nil {
		return nil, fmt.Errorf(`%w: %w`, ErrInvalidYAML, err)
	}

	return parsedContent, nil
}
