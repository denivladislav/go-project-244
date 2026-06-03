package formatters

import (
	"fmt"

	"code/internal/ast"
)

var formatDict = map[string]func(ast.AST) (string, error){
	"stylish": MakeStylish,
	"plain":   MakePlain,
	"json":    MakeJSON,
}

type UnsupportedFormatError struct {
	format string
}

func (e UnsupportedFormatError) Error() string {
	return fmt.Sprintf(`unsupported format: "%s"`, e.format)
}

func getFormatFn(format string) (func(ast.AST) (string, error), error) {
	formatFn, ok := formatDict[format]
	if !ok {
		err := UnsupportedFormatError{format}
		return nil, err
	}

	return formatFn, nil
}

func FormatAst(nodes ast.AST, format string) (string, error) {
	formatFn, err := getFormatFn(format)
	if err != nil {
		return "", fmt.Errorf("get format fn failed: %w", err)
	}

	formattedStr, err := formatFn(nodes)
	if err != nil {
		return "", fmt.Errorf("format fn failed: %w", err)
	}

	return formattedStr, nil
}
