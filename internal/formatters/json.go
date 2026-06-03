package formatters

import (
	"encoding/json"
	"strings"

	"code/internal/diff"
)

type Root struct {
	Key      string    `json:"key"`
	Children diff.Diff `json:"children"`
}

func MakeJSON(nodes diff.Diff) (string, error) {
	indent := strings.Repeat(" ", 2)

	root := Root{
		Key:      "root",
		Children: nodes,
	}

	b, err := json.MarshalIndent(root, "", indent)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
