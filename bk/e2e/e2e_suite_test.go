package e2e_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRepoManger(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MultiGit End to End Test  Suite")
}
