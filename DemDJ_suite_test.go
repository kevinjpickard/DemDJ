package DemDJ_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDemDJ(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DemDJ Suite")
}
