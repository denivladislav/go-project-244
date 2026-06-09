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

func writePlainNodes(pBuilder *strings.Builder, nodes []diff.Node, path string) error {
	for _, node := range nodes {
		newPath := makePlainPath(path, node.Key)

		strValue := stringifyPlain(node.Value)

		switch node.Group {
		case diff.Deleted:
			deletedLine := fmt.Sprintf("Property '%s' was removed\n", newPath)
			pBuilder.WriteString(deletedLine)

			continue
		case diff.Added:
			addedLine := fmt.Sprintf(
				"Property '%s' was added with value: %s\n",
				newPath,
				strValue,
			)
			pBuilder.WriteString(addedLine)

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
			pBuilder.WriteString(modifiedLine)

			continue
		case diff.Nested:
			err := writePlainNodes(pBuilder, node.Children, newPath)
			if err != nil {
				return err
			}

			continue
		default:
			return fmt.Errorf(`%w "%s"`, diff.ErrUnknownGroup, node.Group)
		}
	}

	return nil
}

// MakePlain transforms diff nodes to a string with plain format.
func MakePlain(nodes []diff.Node) (string, error) {
	var builder strings.Builder

	err := writePlainNodes(&builder, nodes, "")
	if err != nil {
		return "", fmt.Errorf("make plain failed: %w", err)
	}

	res := builder.String()

	return strings.TrimSpace(res), nil
}
