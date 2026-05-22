package ast

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

type Group struct {
	slug string
}

func (g Group) String() string {
	return g.slug
}

var (
	Deleted    = Group{"deleted"}
	Added      = Group{"added"}
	Unmodified = Group{"unmodified"}
	Modified   = Group{"modified"}
)

type Node struct {
	key       string
	prevValue any
	value     any
	group     Group
	// children  []Node
}

type Ast = []Node

func Build(objA, objB map[string]any) Ast {
	set := make(map[string]struct{}, len(objA)+len(objB))

	for key := range objA {
		set[key] = struct{}{}
	}

	for key := range objB {
		set[key] = struct{}{}
	}

	keys := maps.Keys(set)
	sortedKeys := slices.Sorted(keys)

	nodes := make([]Node, 0, len(sortedKeys))

	for _, key := range sortedKeys {
		if _, ok := objA[key]; !ok {
			nodes = append(nodes, Node{key: key, value: objB[key], group: Added})
			continue
		}

		if _, ok := objB[key]; !ok {
			nodes = append(nodes, Node{key: key, value: objA[key], group: Deleted})
			continue
		}

		if objA[key] == objB[key] {
			nodes = append(nodes, Node{key: key, value: objA[key], group: Unmodified})
			continue
		}

		nodes = append(
			nodes,
			Node{key: key, prevValue: objA[key], value: objB[key], group: Modified},
		)
	}

	return nodes
}

type UnknownGroupError struct {
	group string
}

func (e UnknownGroupError) Error() string {
	return fmt.Sprintf(`unknown node group: "%s"`, e.group)
}

func Prettify(ast Ast) (string, error) {
	var b strings.Builder
	b.WriteString("{\n")

	for _, node := range ast {
		switch node.group {
		case Deleted:
			newLine := fmt.Sprintf("  - %s: %v\n", node.key, node.value)
			b.WriteString(newLine)

			continue
		case Added:
			newLine := fmt.Sprintf("  + %s: %v\n", node.key, node.value)
			b.WriteString(newLine)

			continue
		case Unmodified:
			newLine := fmt.Sprintf("    %s: %v\n", node.key, node.value)
			b.WriteString(newLine)

			continue
		case Modified:
			prevValueLine := fmt.Sprintf("  - %s: %v\n", node.key, node.prevValue)
			b.WriteString(prevValueLine)

			newValueLine := fmt.Sprintf("  + %s: %v\n", node.key, node.value)
			b.WriteString(newValueLine)

			continue
		default:
			return "", UnknownGroupError{group: node.group.String()}
		}
	}

	b.WriteString("}")

	str := b.String()

	return str, nil
}
