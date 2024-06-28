package cli

import (
	"flag"
)

var inputDirectory string
var outputDirectory string

func Init() {
	const (
		defaultInputDirectory  = "./"
		usageInputDirectory    = "the path to the working directory for scoop project"
		defaultOutputDirectory = "./build"
		usageOutputDirectory   = "the path where build files will be created"
	)
	flag.StringVar(&inputDirectory, "input", defaultInputDirectory, usageInputDirectory)
	flag.StringVar(&inputDirectory, "i", defaultInputDirectory, usageInputDirectory+" (shorthand)")
	flag.StringVar(&outputDirectory, "output", defaultOutputDirectory, usageOutputDirectory)
	flag.StringVar(&outputDirectory, "o", defaultOutputDirectory, usageOutputDirectory+" (shorthand)")
}

func GetInputDirectoryPath() string {
	return inputDirectory
}

func GetOutputDirectoryPath() string {
	return outputDirectory
}

func Cli() {
	Init()

	flag.Parse()
}
