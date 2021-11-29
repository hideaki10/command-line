package repo_manager

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//directory
type RepoManager struct {
	repos        []string
	ingoreErrors bool
}

//
func NewRepoManager(baseDir string, repoNames []string, ingoreErrors bool) (repoManager *RepoManager, err error) {
	_, err = os.Stat(baseDir)
	if err != nil {
		if os.IsNotExist(err) {
			err = errors.New(fmt.Sprintf("base dir: '%s' doesn't exist", baseDir))
		}
		return
	}

	if baseDir[len(baseDir)-1] != '/' {
		baseDir += "/"
	}

	if len(repoNames) == 0 {
		err = errors.New("repo list can't be empty")
		return
	}
	return nil, nil
}

// get repository list
func (m *RepoManager) GetRepos() []string {
	return m.repos
}

// execute the command
func (m *RepoManager) Exec(cmd string) (output map[string]string, err error) {
	output = map[string]string{}
	var components []string
	var multiWord []string

	for _, component := range strings.Split(cmd, " ") {
		if strings.HasPrefix(component, "\"") {
			multiWord = append(multiWord, component[1:])
			continue
		}

		if len(multiWord) > 0 {
			if !strings.HasSuffix(component, "\"") {
				multiWord = append(multiWord, component)
			}

			multiWord = append(multiWord, component[:len(component)-1])

			component = strings.Join(multiWord, " ")

			multiWord = []string{}
		}

		components = append(components, component)
	}

	// Restore working directory after executing the

	wd, _ := os.Getwd()
	defer os.Chdir(wd)

	var out []byte
	for _, r := range m.repos {
		os.Chdir(r)
		out, err = exec.Command("git ", components...).CombinedOutput()
		output[r] = string(out)

		if err != nil && !m.ingoreErrors {
			return
		}

	}

	return

}
