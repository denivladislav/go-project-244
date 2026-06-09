// Package formatters implements specific formatting of diff nodes.
package formatters

import (
	"errors"
	"fmt"

	"code/internal/diff"
)

var ErrUnsupportedFormat = errors.New("unsupported format")

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
		return "", fmt.Errorf(`%w "%s"`, ErrUnsupportedFormat, format)
	}
}
