package formatters

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"code/internal/diff"
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
		var builder strings.Builder
		builder.WriteString("{\n")

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

			builder.WriteString(newLine)
		}

		bracketIndent := makeStylishIndent(depth-1, 0)
		fmt.Fprintf(&builder, "%s}", bracketIndent)

		return builder.String()
	default:
		return fmt.Sprintf("%v", value)
	}
}

type LineOptions struct {
	key, value, marker string
	depth              int
}

func MakeStylish(nodes diff.Diff) (string, error) {
	var builder strings.Builder
	builder.WriteString("{\n")

	var iter func(nodes diff.Diff, depth int) error

	iter = func(nodes diff.Diff, depth int) error {
		for _, node := range nodes {
			strValue := stringifyStylish(node.Value, depth+1)

			switch node.Group {
			case diff.Deleted:
				deletedLine := makeStylishLine(
					LineOptions{
						key:    node.Key,
						value:  strValue,
						marker: "-",
						depth:  depth,
					},
				)
				builder.WriteString(deletedLine)

				continue
			case diff.Added:
				addedLine := makeStylishLine(
					LineOptions{
						key:    node.Key,
						value:  strValue,
						marker: "+",
						depth:  depth,
					},
				)
				builder.WriteString(addedLine)

				continue
			case diff.Unmodified:
				unmodifiedLine := makeStylishLine(
					LineOptions{
						key:    node.Key,
						value:  strValue,
						marker: " ",
						depth:  depth,
					},
				)
				builder.WriteString(unmodifiedLine)

				continue
			case diff.Modified:
				strPrevValue := stringifyStylish(node.PrevValue, depth+1)
				deletedLine := makeStylishLine(
					LineOptions{
						key:    node.Key,
						value:  strPrevValue,
						marker: "-",
						depth:  depth,
					},
				)
				builder.WriteString(deletedLine)

				addedLine := makeStylishLine(
					LineOptions{
						key:    node.Key,
						value:  strValue,
						marker: "+",
						depth:  depth,
					},
				)
				builder.WriteString(addedLine)

				continue
			case diff.Nested:
				keyLine := makeStylishLine(LineOptions{
					key:    node.Key,
					value:  "{",
					marker: " ",
					depth:  depth,
				})
				builder.WriteString(keyLine)

				err := iter(node.Children, depth+1)
				if err != nil {
					return fmt.Errorf("make stylish failed: %w", err)
				}

				bracketIndent := makeStylishIndent(depth, 0)
				fmt.Fprintf(&builder, "%s}\n", bracketIndent)

				continue
			default:
				return diff.UnknownGroupError{Group: node.Group}
			}
		}

		return nil
	}

	err := iter(nodes, 1)
	if err != nil {
		return "", fmt.Errorf("make stylish failed: %w", err)
	}

	builder.WriteString("}")

	return builder.String(), nil
}
