package format

import (
	"fmt"

	"code/internal/ast"
)

type UnknownGroupError struct {
	group string
}

func (e UnknownGroupError) Error() string {
	return fmt.Sprintf(`unknown node group: "%s"`, e.group)
}

type Options struct {
	leftIndent string
	marker     string
	key        string
	value      any
}

type UnsupportedFormatError struct {
	format string
}

func (e UnsupportedFormatError) Error() string {
	return fmt.Sprintf(`unsupported format: "%s"`, e.format)
}

func prettify(nodes ast.Ast, format string) (string, error) {
	switch format {
	case "stylish":
		return MakeStylish(nodes)
	default:
		err := UnsupportedFormatError{format: format}
		return "", err
	}
}

func Prettify(nodes ast.Ast, format string) (string, error) {
	prettified, err := prettify(nodes, format)
	if err != nil {
		return "", fmt.Errorf("prettify failed: %w", err)
	}

	return prettified, nil
}
