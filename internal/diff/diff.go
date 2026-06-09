// Package diff implements computing diff between two objects.
package diff

import (
	"fmt"
	"maps"
	"reflect"
	"slices"
)

// A Group marks whether a Node was added, deleted, modified, etc.
type Group string

type UnknownGroupError struct {
	Group Group
}

func (e UnknownGroupError) Error() string {
	return fmt.Sprintf(`unknown node group: "%s"`, e.Group)
}

const (
	Deleted    Group = "deleted"
	Added      Group = "added"
	Unmodified Group = "unmodified"
	Modified   Group = "modified"
	Nested     Group = "nested"
)

// A Node contains object key-value modification data.
type Node struct {
	Key       string `json:"key"`
	PrevValue any    `json:"prev_value,omitempty"`
	Value     any    `json:"value"`
	Group     Group  `json:"group"`
	Children  []Node `json:"children,omitempty"`
}

func getSortedKeys(objA, objB map[string]any) []string {
	set := make(map[string]struct{}, len(objA)+len(objB))

	for key := range objA {
		set[key] = struct{}{}
	}

	for key := range objB {
		set[key] = struct{}{}
	}

	keys := maps.Keys(set)

	return slices.Sorted(keys)
}

// Build returns a diff between two objects.
// The diff is represented by a slice of Nodes.
// Each Node contains object key-value modification data.
func Build(objA, objB map[string]any) []Node {
	sortedKeys := getSortedKeys(objA, objB)
	nodes := make([]Node, len(sortedKeys))

	for i, key := range sortedKeys {
		if _, ok := objA[key]; !ok {
			nodes[i] = Node{Key: key, Value: objB[key], Group: Added}
			continue
		}

		if _, ok := objB[key]; !ok {
			nodes[i] = Node{Key: key, Value: objA[key], Group: Deleted}
			continue
		}

		childA, isObjChildA := objA[key].(map[string]any)
		childB, isObjChildB := objB[key].(map[string]any)

		if isObjChildA && isObjChildB {
			nodes[i] = Node{Key: key, Children: Build(childA, childB), Group: Nested}
			continue
		}

		if !reflect.DeepEqual(objA[key], objB[key]) {
			nodes[i] = Node{Key: key, PrevValue: objA[key], Value: objB[key], Group: Modified}
			continue
		}

		nodes[i] = Node{Key: key, Value: objA[key], Group: Unmodified}
	}

	return nodes
}
