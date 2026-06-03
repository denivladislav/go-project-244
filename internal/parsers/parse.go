package parsers

import "fmt"

type UnsupportedExtError struct {
	ext string
}

func (e UnsupportedExtError) Error() string {
	return fmt.Sprintf(`unsupported file extension: "%s"`, e.ext)
}

func ParseFileContent(content []byte, ext string) (map[string]any, error) {
	switch ext {
	case ".json":
		return ParseJSON(content)
	case ".yml", ".yaml":
		return ParseYAML(content)
	default:
		return nil, UnsupportedExtError{ext}
	}
}
