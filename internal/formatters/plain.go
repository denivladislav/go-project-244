package formatters

import (
	"fmt"
	"strings"

	"code/internal/diff"
)

func makePlainPath(path, key string) string {
	var newPath string
	if path == "" {
		newPath = key
	} else {
		newPath = fmt.Sprintf("%s.%s", path, key)
	}

	return newPath
}

func stringifyPlain(value any) string {
	switch v := value.(type) {
	case nil:
		return "null"
	case map[string]any:
		return "[complex value]"
	case string:
		return fmt.Sprintf("'%v'", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

func MakePlain(nodes diff.Diff) (string, error) {
	var builder strings.Builder

	var iter func(nodes diff.Diff, path string) error

	iter = func(nodes diff.Diff, path string) error {
		for _, node := range nodes {
			newPath := makePlainPath(path, node.Key)

			strValue := stringifyPlain(node.Value)

			switch node.Group {
			case diff.Deleted:
				deletedLine := fmt.Sprintf("Property '%s' was removed\n", newPath)
				builder.WriteString(deletedLine)

				continue
			case diff.Added:
				addedLine := fmt.Sprintf(
					"Property '%s' was added with value: %s\n",
					newPath,
					strValue,
				)
				builder.WriteString(addedLine)

				continue
			case diff.Unmodified:
				continue
			case diff.Modified:
				strPrevValue := stringifyPlain(node.PrevValue)
				modifiedLine := fmt.Sprintf(
					"Property '%s' was updated. From %s to %s\n",
					newPath,
					strPrevValue,
					strValue,
				)
				builder.WriteString(modifiedLine)

				continue
			case diff.Nested:
				err := iter(node.Children, newPath)
				if err != nil {
					return fmt.Errorf("make plain failed: %w", err)
				}

				continue
			default:
				return diff.UnknownGroupError{Group: node.Group}
			}
		}

		return nil
	}

	err := iter(nodes, "")
	if err != nil {
		return "", fmt.Errorf("make plain failed: %w", err)
	}

	res := builder.String()

	return strings.TrimSpace(res), nil
}
