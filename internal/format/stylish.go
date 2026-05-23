package format

import (
	"fmt"
	"strings"

	"code/internal/ast"
)

type Options struct {
	leftIndent string
	marker     string
	key        string
	value      any
}

func makeLine(options Options) string {
	return fmt.Sprintf(
		"%s%s %s: %v\n",
		options.leftIndent,
		options.marker,
		options.key,
		options.value,
	)
}

func MakeStylish(nodes ast.Ast) (string, error) {
	var b strings.Builder
	b.WriteString("{\n")

	leftIndent := strings.Repeat(" ", 2)

	for _, node := range nodes {
		switch node.Group {
		case ast.Deleted:
			newLine := makeLine(
				Options{leftIndent: leftIndent, marker: "-", key: node.Key, value: node.Value},
			)
			b.WriteString(newLine)

			continue
		case ast.Added:
			newLine := makeLine(
				Options{leftIndent: leftIndent, marker: "+", key: node.Key, value: node.Value},
			)
			b.WriteString(newLine)

			continue
		case ast.Unmodified:
			newLine := makeLine(
				Options{leftIndent: leftIndent, marker: " ", key: node.Key, value: node.Value},
			)
			b.WriteString(newLine)

			continue
		case ast.Modified:
			prevValueLine := makeLine(
				Options{
					leftIndent: leftIndent,
					marker:     "-",
					key:        node.Key,
					value:      node.PrevValue,
				},
			)
			b.WriteString(prevValueLine)

			newValueLine := makeLine(
				Options{leftIndent: leftIndent, marker: "+", key: node.Key, value: node.Value},
			)
			b.WriteString(newValueLine)

			continue
		default:
			return "", ast.UnknownGroupError{Group: node.Group.String()}
		}
	}

	b.WriteString("}")

	str := b.String()

	return str, nil
}
