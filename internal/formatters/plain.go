package formatters

import (
	"fmt"
	"strings"

	"code/internal/ast"
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

func MakePlain(nodes ast.AST) (string, error) {
	var b strings.Builder

	var iter func(nodes ast.AST, path string) error

	iter = func(nodes ast.AST, path string) error {
		for _, node := range nodes {
			newPath := makePlainPath(path, node.Key)

			strValue := stringifyPlain(node.Value)

			switch node.Group {
			case ast.Deleted:
				deletedLine := fmt.Sprintf("Property '%s' was removed\n", newPath)
				b.WriteString(deletedLine)

				continue
			case ast.Added:
				addedLine := fmt.Sprintf(
					"Property '%s' was added with value: %s\n",
					newPath,
					strValue,
				)
				b.WriteString(addedLine)

				continue
			case ast.Unmodified:
				continue
			case ast.Modified:
				strPrevValue := stringifyPlain(node.PrevValue)
				modifiedLine := fmt.Sprintf(
					"Property '%s' was updated. From %s to %s\n",
					newPath,
					strPrevValue,
					strValue,
				)
				b.WriteString(modifiedLine)

				continue
			case ast.Nested:
				err := iter(node.Children, newPath)
				if err != nil {
					return fmt.Errorf("make plain failed: %w", err)
				}

				continue
			default:
				return ast.UnknownGroupError{Group: node.Group}
			}
		}

		return nil
	}

	err := iter(nodes, "")
	if err != nil {
		return "", fmt.Errorf("make plain failed: %w", err)
	}

	res := b.String()

	return strings.TrimSpace(res), nil
}
