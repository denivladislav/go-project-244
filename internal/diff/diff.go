package diff

import (
	"fmt"
	"maps"
	"slices"

	"reflect"
)

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

type Node struct {
	Key       string `json:"key"`
	PrevValue any    `json:"prev_value,omitempty"`
	Value     any    `json:"value"`
	Group     Group  `json:"group"`
	Children  []Node `json:"children,omitempty"`
}

type Diff = []Node

func Build(objA, objB map[string]any) Diff {
	set := make(map[string]struct{}, len(objA)+len(objB))

	for key := range objA {
		set[key] = struct{}{}
	}

	for key := range objB {
		set[key] = struct{}{}
	}

	keys := maps.Keys(set)
	sortedKeys := slices.Sorted(keys)

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
