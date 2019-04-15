package spec

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestE2ESpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "E2E Spec Suite")
}
