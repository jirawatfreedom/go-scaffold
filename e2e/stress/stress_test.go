package stress

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestE2EStress(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping stress test in short mode")
		return
	}
	RegisterFailHandler(Fail)
	RunSpecs(t, "E2E Stress Suite")
}
