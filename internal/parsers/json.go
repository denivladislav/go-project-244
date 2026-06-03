package parsers

import "encoding/json"

func ParseJSON(data []byte) (map[string]any, error) {
	var parsedContent map[string]any

	err := json.Unmarshal(data, &parsedContent)
	if err != nil {
		return nil, err
	}

	return parsedContent, nil
}
