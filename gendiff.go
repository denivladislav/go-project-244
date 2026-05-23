package code

import (
	"fmt"
	"os"
	"path/filepath"

	"code/internal/ast"
	"code/internal/format"
	"code/internal/parse"
)

func GenDiff(pathA, pathB, formatName string) (string, error) {
	bytesA, err := os.ReadFile(pathA)
	if err != nil {
		return "", fmt.Errorf("read file failed: %w", err)
	}

	bytesB, err := os.ReadFile(pathB)
	if err != nil {
		return "", fmt.Errorf("read file failed: %w", err)
	}

	extA := filepath.Ext(pathA)
	extB := filepath.Ext(pathB)

	parsedA, err := parse.FileContent(bytesA, extA)
	if err != nil {
		return "", fmt.Errorf("parse content failed: %w", err)
	}

	parsedB, err := parse.FileContent(bytesB, extB)
	if err != nil {
		return "", fmt.Errorf("parse content failed: %w", err)
	}

	newAst := ast.Build(parsedA, parsedB)

	return format.Prettify(newAst, formatName)
}
