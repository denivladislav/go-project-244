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

func FormatDiff(nodes diff.Diff, format string) (string, error) {
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
