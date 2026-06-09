package formatters

import (
	"encoding/json"
	"strings"

	"code/internal/diff"
)

type Root struct {
	Key      string      `json:"key"`
	Children []diff.Node `json:"children"`
}

// MakeJSON transforms diff nodes to a string with JSON format.
func MakeJSON(nodes []diff.Node) (string, error) {
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
