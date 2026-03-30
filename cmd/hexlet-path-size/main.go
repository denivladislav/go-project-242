package main

import (
	"code/pathsize"
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

var cmdFlags = []cli.Flag{
	humanFlag,
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

			human := cmd.Bool(humanFlag.Name)
			size, err := pathsize.GetPathSize(path, human)
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
