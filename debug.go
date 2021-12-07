package main

import (
	"fmt"

	. "github.com/hideaki10/command-line/pkg/helpers"
	. "github.com/hideaki10/command-line/pkg/repo_manager"
)

const baseDir = "/Users/senshikou/Desktop/Development/gomodgithub/commandline/gitrepos"

var repoList = []string{
	"repo1",
	"repo2",
}

func main() {

	rm, err := NewRepoManager(baseDir, repoList, true)
	if err != nil {
		fmt.Print(err)
	}

	AddFiles(baseDir, repoList[0], true, "file_1.txt", "file_2.txt")

	output, err := rm.Exec("log --oneline")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(output[baseDir+"/repo1"])
}
