package main

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
			PIt("server should start successfully", func() {

			})
			PIt("port 33635 should be occupied", func() {

			})
			PIt("log file should be created", func() {

			})
			PIt("party file should be created", func() {

			})
			PIt("vote file should be created", func() {

			})
		})

		PContext("When port 33635 is not open", func() {
			PIt("should show a warning", func() {

			})
			PIt("should scan for the next open port", func() {

			})
			PIt("should print the open port", func() {

			})
			PIt("server should start successfully", func() {

			})
			PIt("occupying the open port", func() {

			})
			PIt("log file should be created", func() {

			})
			PIt("party file should be created", func() {

			})
			PIt("vote file should be created", func() {

			})
		})
	})
})
