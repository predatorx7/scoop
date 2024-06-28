package main

import (
	"os"

	"github.com/predatorx7/scoop/internal/pkg/cli"
	"github.com/predatorx7/scoop/pkg/libscoop"
)

func main() {
	cli.Cli()

	result := libscoop.Build(&libscoop.ScoopOption{
		InputDirectoryPath:  cli.GetInputDirectoryPath(),
		OutputDirectoryPath: cli.GetOutputDirectoryPath(),
	})

	if result != nil {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
