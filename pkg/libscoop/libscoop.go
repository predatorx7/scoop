package libscoop

import (
	"errors"
	"fmt"
)

func Build(option *ScoopOption) error {
	if option == nil {
		fmt.Print("No options provided")
		return errors.New("No options provided")
	}

	fmt.Printf("inputDirectoryPath: %s\n", option.InputDirectoryPath)
	fmt.Printf("outputDirectoryPath: %s\n", option.OutputDirectoryPath)

	return nil
}
