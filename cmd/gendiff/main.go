package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

var ErrRequiredPath = errors.New("path is required")

func main() {
	cmd := &cli.Command{
		Name:  "gendiff",
		Usage: "Compares two configuration files and shows the difference",
		// Flags: cmdFlags,
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := cmd.Args().Get(0)

			if path == "" {
				return ErrRequiredPath
			}

			res := "hello, world"

			fmt.Printf("%s", res)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
