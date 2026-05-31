package formatters

import (
	"encoding/json"
	"strings"

	"code/internal/ast"
)

type Root struct {
	Key      string  `json:"key"`
	Children ast.AST `json:"children"`
}

func MakeJson(nodes ast.AST) (string, error) {
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
