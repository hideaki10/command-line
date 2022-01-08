package main

import (
	"fmt"

	"github.com/hideaki10/command-line/cmd"
)

var (
	GitTag    string
	Timestamp string
)

func main() {
	if GitTag != "" {
		fmt.Printf("Git tag : %s\nBuilt at: %s\n\n", GitTag, Timestamp)
	}
	cmd.Execute()
}
