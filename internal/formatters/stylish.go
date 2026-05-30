package formatters

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"code/internal/ast"
)

func makeStylishIndent(depth, shift int) string {
	return strings.Repeat(" ", depth*4-shift)
}

func makeStylishLine(options LineOptions) string {
	shift := len(options.marker) + 1
	indent := makeStylishIndent(options.depth, shift)

	return fmt.Sprintf("%s%s %s: %v\n",
		indent,
		options.marker,
		options.key,
		options.value)
}

func stringifyStylish(value any, depth int) string {
	switch v := value.(type) {
	case nil:
		return "null"
	case map[string]any:
		var b strings.Builder
		b.WriteString("{\n")

		keys := maps.Keys(v)
		sortedKeys := slices.Sorted(keys)

		for _, key := range sortedKeys {
			strValue := stringifyStylish(v[key], depth+1)

			newLine := makeStylishLine(LineOptions{
				key:    key,
				value:  strValue,
				marker: " ",
				depth:  depth,
			})

			b.WriteString(newLine)
		}

		bracketIndent := makeStylishIndent(depth-1, 0)
		fmt.Fprintf(&b, "%s}", bracketIndent)

		return b.String()
	default:
		return fmt.Sprintf("%v", value)
	}
}

type LineOptions struct {
	key, value, marker string
	depth              int
}

func MakeStylish(nodes ast.AST) (string, error) {
	var b strings.Builder
	b.WriteString("{\n")

	var iter func(nodes ast.AST, depth int) error

	iter = func(nodes ast.AST, depth int) error {
		for _, node := range nodes {
			strValue := stringifyStylish(node.Value, depth+1)

			switch node.Group {
			case ast.Deleted:
				deletedLine := makeStylishLine(
					LineOptions{
						key:    node.Key,
						value:  strValue,
						marker: "-",
						depth:  depth,
					},
				)
				b.WriteString(deletedLine)

				continue
			case ast.Added:
				addedLine := makeStylishLine(
					LineOptions{
						key:    node.Key,
						value:  strValue,
						marker: "+",
						depth:  depth,
					},
				)
				b.WriteString(addedLine)

				continue
			case ast.Unmodified:
				unmodifiedLine := makeStylishLine(
					LineOptions{
						key:    node.Key,
						value:  strValue,
						marker: " ",
						depth:  depth,
					},
				)
				b.WriteString(unmodifiedLine)

				continue
			case ast.Modified:
				strPrevValue := stringifyStylish(node.PrevValue, depth+1)
				deletedLine := makeStylishLine(
					LineOptions{
						key:    node.Key,
						value:  strPrevValue,
						marker: "-",
						depth:  depth,
					},
				)
				b.WriteString(deletedLine)

				addedLine := makeStylishLine(
					LineOptions{
						key:    node.Key,
						value:  strValue,
						marker: "+",
						depth:  depth,
					},
				)
				b.WriteString(addedLine)

				continue
			case ast.Nested:
				keyLine := makeStylishLine(LineOptions{
					key:    node.Key,
					value:  "{",
					marker: " ",
					depth:  depth,
				})
				b.WriteString(keyLine)

				err := iter(node.Children, depth+1)
				if err != nil {
					return fmt.Errorf("make stylish failed: %w", err)
				}

				bracketIndent := makeStylishIndent(depth, 0)
				fmt.Fprintf(&b, "%s}\n", bracketIndent)

				continue
			default:
				return ast.UnknownGroupError{Group: node.Group.String()}
			}
		}

		return nil
	}

	err := iter(nodes, 1)
	if err != nil {
		return "", fmt.Errorf("make stylish failed: %w", err)
	}

	b.WriteString("}")

	return b.String(), nil
}
