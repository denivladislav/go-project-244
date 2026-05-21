package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"

	gendiff "code"
)

var formats = []string{"stylish"}

var formatFlag = &cli.StringFlag{
	Name:    "format",
	Value:   formats[0],
	Usage:   "output format",
	Aliases: []string{"f"},
}

var cmdFlags = []cli.Flag{
	formatFlag,
}

var ErrMissingPaths = errors.New("two paths are required")

func main() {
	cmd := &cli.Command{
		Name:  "gendiff",
		Usage: "Compares two configuration files and shows the difference",
		Flags: cmdFlags,
		Action: func(ctx context.Context, cmd *cli.Command) error {
			pathA := cmd.Args().Get(0)
			pathB := cmd.Args().Get(1)

			if pathA == "" || pathB == "" {
				return ErrMissingPaths
			}

			diff, err := gendiff.GenDiff(
				pathA,
				pathB,
				cmd.String(formatFlag.Name),
			)
			if err != nil {
				return fmt.Errorf("gen diff failed: %w", err)
			}

			fmt.Printf("%s\n", diff)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
