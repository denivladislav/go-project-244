package code

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"code/internal/format"
	"code/internal/parsers"
)

var testDataPath = filepath.Join(".", "testdata", "fixture")

func validateError(t *testing.T, err error, checkErr func(err error) bool) {
	t.Helper()

	if !checkErr(err) {
		t.Errorf("error check failed for error: %v", err)
	}
}

func TestGenDiff_Errors(t *testing.T) {
	tests := map[string]struct {
		pathA, pathB, formatName string
		checkErr                 func(err error) bool
	}{
		"unreachable path causes error": {
			pathA:      "",
			pathB:      filepath.Join(testDataPath, "file2.json"),
			formatName: "stylish",
			checkErr: func(err error) bool {
				return errors.Is(err, os.ErrNotExist)
			},
		},
		"path to directory causes error": {
			pathA:      testDataPath,
			pathB:      filepath.Join(testDataPath, "file2.json"),
			formatName: "stylish",
			checkErr: func(err error) bool {
				var pathErr *fs.PathError
				return errors.As(err, &pathErr)
			},
		},
		"unsupported file extension causes error": {
			pathA:      filepath.Join(testDataPath, "unsupported.diff"),
			pathB:      filepath.Join(testDataPath, "file2.json"),
			formatName: "stylish",
			checkErr: func(err error) bool {
				var extErr parsers.UnsupportedExtError
				return errors.As(err, &extErr)
			},
		},
		"invalid file causes error": {
			pathA:      filepath.Join(testDataPath, "invalid.json"),
			pathB:      filepath.Join(testDataPath, "file2.json"),
			formatName: "stylish",
			checkErr: func(err error) bool {
				return err != nil
			},
		},
		"unsupported output format causes error": {
			pathA:      filepath.Join(testDataPath, "file1.json"),
			pathB:      filepath.Join(testDataPath, "file2.json"),
			formatName: "unknown",
			checkErr: func(err error) bool {
				var formatError format.UnsupportedFormatError
				return errors.As(err, &formatError)
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := GenDiff(tt.pathA, tt.pathB, tt.formatName)

			validateError(t, err, tt.checkErr)
		})
	}
}

func TestGenDiff(t *testing.T) {
	tests := map[string]struct {
		pathA, pathB, formatName, pathWant string
	}{
		"generates correct diff for json with stylish format": {
			pathA:      filepath.Join(testDataPath, "file1.json"),
			pathB:      filepath.Join(testDataPath, "file2.json"),
			formatName: "stylish",
			pathWant:   filepath.Join(testDataPath, "result_flat_stylish.txt"),
		},
		"generates correct diff for yml and yaml with stylish format": {
			pathA:      filepath.Join(testDataPath, "file1.yml"),
			pathB:      filepath.Join(testDataPath, "file2.yaml"),
			formatName: "stylish",
			pathWant:   filepath.Join(testDataPath, "result_flat_stylish.txt"),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			bytesWant, err := os.ReadFile(tt.pathWant)
			require.NoError(t, err)

			got, err := GenDiff(tt.pathA, tt.pathB, tt.formatName)
			require.NoError(t, err)

			assert.Equal(t, string(bytesWant), got)
		})
	}
}
