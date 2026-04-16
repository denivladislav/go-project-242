package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"

	pathsize "code"
)

var humanFlag = &cli.BoolFlag{
	Name:    "human",
	Value:   true,
	Usage:   "human-readable sizes (auto-select unit)",
	Aliases: []string{"H"},
}

var allFlag = &cli.BoolFlag{
	Name:    "all",
	Value:   false,
	Usage:   "include hidden files and directories",
	Aliases: []string{"a"},
}

var recursiveFlag = &cli.BoolFlag{
	Name:    "recursive",
	Value:   false,
	Usage:   "recursive size of directories",
	Aliases: []string{"r"},
}

var cmdFlags = []cli.Flag{
	humanFlag,
	allFlag,
	recursiveFlag,
}

var ErrRequiredPath = errors.New("path is required")

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Flags: cmdFlags,
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := cmd.Args().Get(0)

			if path == "" {
				return ErrRequiredPath
			}

			size, err := pathsize.GetPathSize(
				path,
				cmd.Bool(recursiveFlag.Name),
				cmd.Bool(humanFlag.Name),
				cmd.Bool(allFlag.Name),
			)
			if err != nil {
				return fmt.Errorf("get path size failed: %w", err)
			}

			fmt.Printf("%s\t%s\n", size, path)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
