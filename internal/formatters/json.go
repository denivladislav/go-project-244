package formatters

import (
	"encoding/json"
	"strings"

	"code/internal/ast"
)

func MakeJson(nodes ast.AST) (string, error) {
	indent := strings.Repeat(" ", 2)

	b, err := json.MarshalIndent(nodes, "", indent)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
