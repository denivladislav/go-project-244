package format

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"code/internal/ast"
)

func makeIndent(depth, shift int) string {
	return strings.Repeat(" ", depth*4-shift)
}

func makeLine(options LineOptions) string {
	shift := len(options.marker) + 1
	indent := makeIndent(options.depth, shift)

	return fmt.Sprintf("%s%s %s: %v\n",
		indent,
		options.marker,
		options.key,
		options.value)
}

func stringify(value any, depth int) string {
	switch v := value.(type) {
	case nil:
		return "null"
	case map[string]any:
		var b strings.Builder
		b.WriteString("{\n")

		keys := maps.Keys(v)
		sortedKeys := slices.Sorted(keys)

		for _, key := range sortedKeys {
			stringifiedValue := stringify(v[key], depth+1)

			newLine := makeLine(LineOptions{
				key:    key,
				value:  stringifiedValue,
				marker: " ",
				depth:  depth,
			})

			b.WriteString(newLine)
		}

		bracketIndent := makeIndent(depth-1, 0)
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
			strValue := stringify(node.Value, depth+1)

			switch node.Group {
			case ast.Deleted:
				deletedLine := makeLine(
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
				addedLine := makeLine(
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
				unmodifiedLine := makeLine(
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
				strPrevValue := stringify(node.PrevValue, depth+1)
				deletedLine := makeLine(
					LineOptions{
						key:    node.Key,
						value:  strPrevValue,
						marker: "-",
						depth:  depth,
					},
				)
				b.WriteString(deletedLine)

				addedLine := makeLine(
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
				keyLine := makeLine(LineOptions{
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

				bracketIndent := makeIndent(depth, 0)
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
