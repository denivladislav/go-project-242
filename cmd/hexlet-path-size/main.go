package main

import (
	"code"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
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

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Flags: cmdFlags,
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := cmd.Args().Get(0)

			if path == "" {
				return fmt.Errorf("error: path is required")
			}

			config := code.Config{
				Human:     cmd.Bool(humanFlag.Name),
				All:       cmd.Bool(allFlag.Name),
				Recursive: cmd.Bool(recursiveFlag.Name),
			}

			size, err := code.GetPathSize(path, config)
			if err != nil {
				return fmt.Errorf("error: %w", err)
			}

			fmt.Printf("%s\t%s\n", size, path)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
