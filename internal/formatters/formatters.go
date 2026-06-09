// Package formatters implements specific formatting of diff nodes.
package formatters

import (
	"fmt"

	"code/internal/diff"
)

type UnsupportedFormatError struct {
	format string
}

func (e UnsupportedFormatError) Error() string {
	return fmt.Sprintf(`unsupported format: "%s"`, e.format)
}

// FormatDiff returns a string representing diff nodes depending on format.
//
// If format is not supported returns UnsupportedFormatError.
func FormatDiff(nodes []diff.Node, format string) (string, error) {
	switch format {
	case "stylish":
		return MakeStylish(nodes)
	case "plain":
		return MakePlain(nodes)
	case "json":
		return MakeJSON(nodes)
	default:
		return "", UnsupportedFormatError{format}
	}
}
