package diff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {
	tests := map[string]struct {
		objA, objB map[string]any
		want       []Node
	}{
		"builds correct diff if one of the objects is empty": {
			objA: map[string]any{"key": "value"},
			objB: map[string]any{},
			want: []Node{{Key: "key", Value: "value", Group: Deleted}},
		},

		"builds correct alphabetically sorted diff for complex objects": {
			objA: map[string]any{
				"b": true,
				"a": "old",
				"c": map[string]any{
					"d": []int{1, 2, 3},
					"e": 1,
				},
			},
			objB: map[string]any{
				"a": "new",
				"f": nil,
				"c": map[string]any{
					"d": []int{1, 2},
					"e": 2,
				},
			},
			want: []Node{
				{Key: "a", PrevValue: "old", Value: "new", Group: Modified},
				{Key: "b", Value: true, Group: Deleted},
				{Key: "c", Group: Nested, Children: []Node{
					{Key: "d", PrevValue: []int{1, 2, 3}, Value: []int{1, 2}, Group: Modified},
					{Key: "e", PrevValue: 1, Value: 2, Group: Modified},
				}},
				{Key: "f", Value: nil, Group: Added},
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := Build(tt.objA, tt.objB)

			assert.Equal(t, tt.want, got)
		})
	}
}
