package ast

import (
	"maps"
	"slices"
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
	Key       string
	PrevValue any
	Value     any
	Group     Group
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
			nodes = append(nodes, Node{Key: key, Value: objB[key], Group: Added})
			continue
		}

		if _, ok := objB[key]; !ok {
			nodes = append(nodes, Node{Key: key, Value: objA[key], Group: Deleted})
			continue
		}

		if objA[key] == objB[key] {
			nodes = append(nodes, Node{Key: key, Value: objA[key], Group: Unmodified})
			continue
		}

		nodes = append(
			nodes,
			Node{Key: key, PrevValue: objA[key], Value: objB[key], Group: Modified},
		)
	}

	return nodes
}
