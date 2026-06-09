package parsers

import "gopkg.in/yaml.v3"

// ParseYAML parses YAML-encoded data to key-value map.
func ParseYAML(data []byte) (map[string]any, error) {
	var parsedContent map[string]any

	err := yaml.Unmarshal(data, &parsedContent)
	if err != nil {
		return nil, err
	}

	return parsedContent, nil
}
