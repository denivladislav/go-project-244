/*
Gendiff compares two configuration files and shows the difference.

Without explicit format default "stylish" format is used.

Usage:

	bin/gendiff [pathA] [pathB] [flags]

Flags:

  - f
    Specific output format ("stylish", "plain", etc.)
*/
package main

import (
	"context"
	"fmt"
	"os"

	"code/internal/app"
)

func main() {
	newApp := app.New()

	if err := newApp.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
