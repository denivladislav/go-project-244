// Package parsers implements parsing of structured files like JSON and YAML.
package parsers

import (
	"errors"
	"fmt"
)

var ErrUnsupportedExt = errors.New("unsupported file extension")

// ParseFileContent parses contents of a file depending on its extension.
//
// If file extension is not supported returns UnsupportedExtError.
func ParseFileContent(content []byte, ext string) (map[string]any, error) {
	switch ext {
	case ".json":
		return ParseJSON(content)
	case ".yml", ".yaml":
		return ParseYAML(content)
	default:
		return nil, fmt.Errorf(`%w "%s"`, ErrUnsupportedExt, ext)
	}
}
