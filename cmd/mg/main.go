package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	command := flag.String("command", "", "The git command")

	ignoreErrors := flag.Bool(
		"ignore-errors",
		false,
		"Keep running after error if true")
	flag.Parse()

	root := os.Getenv("MG_ROOT")
	if root[len(root)-1] != '/' {
		root += "/"
	}

	// get the repository name from the environment
	repo_names := strings.Split(os.Getenv("MG_REPOS"), ",")

	var repos []string
	for _, r := range repo_names {
		path := root + r
		_, err := os.Stat(path + "/.git")
		if err != nil {
			log.Fatal(err)
		}
		repos = append(repos, path)
	}

	// "status"
	var git_components []string
	for _, component := range strings.Split(*command, " ") {
		git_components = append(git_components, component)
	}

	command_string := "git " + *command

	// execute the command for each repository
	for _, r := range repos {
		os.Chdir(r)
		fmt.Printf("[%s] %s\n", r, command_string)
		out, err := exec.Command("git", git_components...).CombinedOutput()
		fmt.Println(string(out))
		if err != nil && !*ignoreErrors {
			os.Exit(1)
		}
	}
	fmt.Println("Done .")
}
