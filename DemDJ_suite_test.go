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

// P<Verb> means the spec is pending
var _ = PDescribe("server", func() {
  var (
    // server server
    // error error
  )

  PDescribe("Starting server", func() {
    PContext("When port 33635 is open", func() {
      PIt("should start successfully, occupying port 33635", func() {

      })
    })
  })
})
