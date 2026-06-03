package parsers

import "fmt"

var parseDict = map[string]func([]byte) (map[string]any, error){
	".json": ParseJSON,
	".yml":  ParseYAML,
	".yaml": ParseYAML,
}

type UnsupportedExtError struct {
	ext string
}

func (e UnsupportedExtError) Error() string {
	return fmt.Sprintf(`unsupported file extension: "%s"`, e.ext)
}

func getParseFn(ext string) (func([]byte) (map[string]any, error), error) {
	parseFn, ok := parseDict[ext]
	if !ok {
		err := UnsupportedExtError{ext}
		return nil, err
	}

	return parseFn, nil
}

func ParseFileContent(content []byte, ext string) (map[string]any, error) {
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
