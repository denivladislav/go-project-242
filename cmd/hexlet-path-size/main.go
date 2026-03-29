package main

import (
	"context"
	"fmt"
	"log"
	"os"

	path_size "code"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := cmd.Args().Get(0)

			if path == "" {
				return fmt.Errorf("error: path is required")
			}

			size, err := path_size.GetPathSize(path)
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
