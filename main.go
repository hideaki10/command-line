package main

import (
	"fmt"
	"log"

	"github.com/blang/semver"
	"github.com/hideaki10/command-line/cmd"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

const (
	version = "v0.9.8"
)

var (
	GitTag    string
	Timestamp string
)

func main() {
	fmt.Println("version : " + version)
	v := semver.MustParse(version[1:])
	latest, err := selfupdate.UpdateSelf(v, "hideaki10/command-line")
	if err != nil {
		log.Fatalf("Binary update failed: %v", err)
		return
	} else {
		fmt.Println("Current version is : " + latest.Version.String())
	}
	if GitTag != "" {
		fmt.Printf("Git tag : %s\nBuilt at: %s\n\n", GitTag, Timestamp)
	}
	cmd.Execute()
}
