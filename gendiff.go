package code

import (
	"fmt"
	"os"
	"path/filepath"

	"code/internal/ast"
	"code/internal/formatters"
	"code/internal/parsers"
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

	parsedA, err := parsers.ParseContent(bytesA, extA)
	if err != nil {
		return "", fmt.Errorf("parse content failed: %w", err)
	}

	parsedB, err := parsers.ParseContent(bytesB, extB)
	if err != nil {
		return "", fmt.Errorf("parse content failed: %w", err)
	}

	newAst := ast.BuildAst(parsedA, parsedB)

	return formatters.FormatAst(newAst, formatName)
}
