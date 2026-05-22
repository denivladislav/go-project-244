package parse

import (
	"fmt"

	"encoding/json"
)

type UnsupportedExtError struct {
	ext string
}

func (e UnsupportedExtError) Error() string {
	return fmt.Sprintf(`unsupported file extension: "%s"`, e.ext)
}

func parseJson(data []byte) (map[string]any, error) {
	var m map[string]any

	err := json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func parse(data []byte, ext string) (map[string]any, error) {
	switch ext {
	case ".json":
		return parseJson(data)
	default:
		err := UnsupportedExtError{ext: ext}
		return nil, err
	}
}

func FileContent(content []byte, ext string) (map[string]any, error) {
	parsed, err := parse(content, ext)
	if err != nil {
		return nil, fmt.Errorf("parse failed: %w", err)
	}

	return parsed, nil
}
