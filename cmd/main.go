package main

import (
	"fmt"

	"github.com/predatorx7/scoop/internal/pkg/cli"
)

func main() {
	cli.Cli()

	fmt.Printf("input: %s, output: %s\n", cli.GetInputDirectoryPath(), cli.GetOutputDirectoryPath())
}
