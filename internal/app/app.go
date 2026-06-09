// Package app implements entry point to gendiff CLI-utility
package app

import (
	"context"
	"errors"
	"fmt"

	"github.com/urfave/cli/v3"

	gendiff "code"
)

var ErrMissingPath = errors.New("two paths are required")

// Run reads filepaths from CLI, invokes GenDiff and prints the diff.
func Run(ctx context.Context, cmd *cli.Command) error {
	pathA := cmd.Args().Get(0)
	pathB := cmd.Args().Get(1)

	if pathA == "" || pathB == "" {
		return ErrMissingPath
	}

	diff, err := gendiff.GenDiff(
		pathA,
		pathB,
		cmd.String("format"),
	)
	if err != nil {
		return fmt.Errorf("gen diff failed: %w", err)
	}

	fmt.Printf("\n%s\n", diff)

	return nil
}

// New returns new configured cli.Command.
func New() *cli.Command {
	return &cli.Command{
		Name:  "gendiff",
		Usage: "Compares two configuration files and shows the difference",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Value:   "stylish",
				Usage:   "output format",
				Aliases: []string{"f"},
			},
		},
		Action: Run,
	}
}
