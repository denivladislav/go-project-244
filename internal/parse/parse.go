package parse

import (
	"encoding/json"
	"fmt"
)

var parseDict = map[string]func([]byte) (map[string]any, error){
	".json": ParseJson,
}

type UnsupportedExtError struct {
	ext string
}

func (e UnsupportedExtError) Error() string {
	return fmt.Sprintf(`unsupported file extension: "%s"`, e.ext)
}

func ParseJson(data []byte) (map[string]any, error) {
	var m map[string]any

	err := json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func getParseFn(ext string) (func([]byte) (map[string]any, error), error) {
	parseFn, ok := parseDict[ext]
	if !ok {
		err := UnsupportedExtError{ext}
		return nil, err
	}

	return parseFn, nil
}

func ParseContent(content []byte, ext string) (map[string]any, error) {
	parseFn, err := getParseFn(ext)
	if err != nil {
		return nil, fmt.Errorf("get parse fn failed: %w", err)
	}

	parsedContent, err := parseFn(content)
	if err != nil {
		return nil, fmt.Errorf("parse fn failed: %w", err)
	}

	return parsedContent, nil
}
