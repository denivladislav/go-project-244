package formatters

import (
	"errors"
	"testing"

	"code/internal/diff"

	"github.com/stretchr/testify/assert"
)

func validateError(t *testing.T, err error, checkErr func(err error) bool) {
	t.Helper()

	if !checkErr(err) {
		t.Errorf("error check failed for error: %v", err)
	}
}

func TestFormatDiff_Errors(t *testing.T) {
	tests := map[string]struct {
		nodes    []diff.Node
		format   string
		checkErr func(err error) bool
	}{
		"unsupported format causes error": {
			nodes:  []diff.Node{{Key: "key", Value: "value", Group: diff.Deleted}},
			format: "unsupported",
			checkErr: func(err error) bool {
				return errors.Is(err, ErrUnsupportedFormat)
			},
		},

		"invalid group causes error": {
			nodes:  []diff.Node{{Key: "key", Value: "value", Group: "invalid"}},
			format: "stylish",
			checkErr: func(err error) bool {
				return errors.Is(err, diff.ErrUnknownGroup)
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			_, err := FormatDiff(tt.nodes, tt.format)

			validateError(t, err, tt.checkErr)
		})
	}
}

var emptyStringNodesAdded = []diff.Node{{Key: "foo", Value: "", Group: diff.Added}}
var nilNodesUnmodified = []diff.Node{{Key: "baz", Value: nil, Group: diff.Unmodified}}
var sliceNodesModified = []diff.Node{
	{Key: "s", PrevValue: []int{1, 2, 3}, Value: []int{1, 2}, Group: diff.Modified}}

func TestFormatDiff(t *testing.T) {
	tests := map[string]struct {
		nodes  []diff.Node
		format string
		want   string
	}{
		"formats empty nodes in stylish format correctly": {
			nodes:  []diff.Node{},
			format: "stylish",
			want:   "{\n}",
		},
		"formats empty string value in stylish format correctly": {
			nodes:  emptyStringNodesAdded,
			format: "stylish",
			want:   "{\n  + foo: \n}",
		},
		"formats slices in stylish format correctly": {
			nodes:  sliceNodesModified,
			format: "stylish",
			want:   "{\n  - s: [1 2 3]\n  + s: [1 2]\n}",
		},
		"formats nil in stylish format correctly": {
			nodes:  nilNodesUnmodified,
			format: "stylish",
			want:   "{\n    baz: null\n}",
		},

		"formats empty nodes in plain format correctly": {
			nodes:  []diff.Node{},
			format: "plain",
			want:   "",
		},
		"formats empty string value in plain format correctly": {
			nodes:  emptyStringNodesAdded,
			format: "plain",
			want:   "Property 'foo' was added with value: ''",
		},
		"formats slices in plain format correctly": {
			nodes:  sliceNodesModified,
			format: "plain",
			want:   "Property 's' was updated. From [1 2 3] to [1 2]",
		},
		"formats nil in plain format correctly": {
			nodes:  sliceNodesModified,
			format: "plain",
			want:   "Property 's' was updated. From [1 2 3] to [1 2]",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := FormatDiff(tt.nodes, tt.format)
			assert.NoError(t, err)

			assert.Equal(t, tt.want, got)
		})
	}
}
