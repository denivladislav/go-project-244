// Package code provides GenDiff function.
// GenDiff outputs the diff between two configuration files.
package code

import (
	"fmt"
	"os"
	"path/filepath"

	"code/internal/diff"
	"code/internal/formatters"
	"code/internal/parsers"
)

// GenDiff reads two filepaths, computes the diff between them and outputs as formatted string.
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

	parsedA, err := parsers.ParseFileContent(bytesA, extA)
	if err != nil {
		return "", fmt.Errorf(`parse file content for "%s" failed: %w`, pathA, err)
	}

	parsedB, err := parsers.ParseFileContent(bytesB, extB)
	if err != nil {
		return "", fmt.Errorf(`parse file content for "%s" failed: %w`, pathB, err)
	}

	diff := diff.Build(parsedA, parsedB)

	return formatters.FormatDiff(diff, formatName)
}
