package day23_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test23(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "23 Suite")
}