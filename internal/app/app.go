// Package app implements entry point to gendiff CLI-utility
package app

import (
	"context"
	"errors"
	"fmt"

	"github.com/urfave/cli/v3"

	gendiff "code"
)

var ErrRequiredPaths = errors.New("expected 2 paths")

// Run reads filepaths from CLI, invokes GenDiff and prints the diff.
func Run(_ context.Context, cmd *cli.Command) error {
	if cmd.Args().Len() != 2 {
		return fmt.Errorf(`%w, got: %d`, ErrRequiredPaths, cmd.Args().Len())
	}

	pathA := cmd.Args().Get(0)
	pathB := cmd.Args().Get(1)

	diff, err := gendiff.GenDiff(
		pathA,
		pathB,
		cmd.String("format"),
	)
	if err != nil {
		return fmt.Errorf("gen diff failed: %w", err)
	}

	fmt.Println(diff)

	return nil
}

// New returns new configured cli.Command.
func New() *cli.Command {
	return &cli.Command{
		Name:      "gendiff",
		Usage:     "Compares two configuration files and shows the difference",
		ArgsUsage: "[pathA] [pathB]",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Value:   "stylish",
				Usage:   "output format (available: stylish, plain, json)",
				Aliases: []string{"f"},
			},
		},
		Action: Run,
	}
}
