package cmd

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/hideaki10/command-line/pkg/repo_manager"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ignoreErrors bool

var rootCmd = &cobra.Command{
	Use:   "mg",
	Short: "Manage git repositories",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		root := os.Getenv("MG_ROOT")
		if root[len(root)-1] != '/' {
			root += "/"
		}

		repoNames := []string{}

		if len(os.Getenv("ME_REPOS")) > 0 {
			repoNames = strings.Split(os.Getenv("MG_REPOS"), ",")
		}

		repoManager, err := repo_manager.NewRepoManager(root, repoNames, ignoreErrors)
		if err != nil {
			log.Fatal(err)
		}

		command := strings.Join(args, " ")

		output, err := repoManager.Exec(command)
		if err != nil {
			fmt.Printf("command '%s'failed with error ", err)
		}

		for repo, out := range output {
			fmt.Printf("[%s]: git %s\n", path.Base(repo), command)
			fmt.Println(out)
		}

	},
}

func init() {
	//rootCmd.Flags().BoolVar(&ignoreErrors, "ignore-errors", false, `will continue executing the command for all repositories if ignore-errors is true otherwise it will stop execution when an error occurs`)
	rootCmd.Flags().Bool(
		"ignore-errors",
		false,
		`will continue executing the command for all repos if ignore-errors is true otherwise it will stop execution when an error occurs`)
	err := viper.BindPFlag("ignore-errors", rootCmd.Flags().Lookup("ignore-errors"))
	if err != nil {
		panic("Unable to bind flag")
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
