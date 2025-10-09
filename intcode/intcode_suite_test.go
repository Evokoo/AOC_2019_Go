package intcode_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestIntcode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Intcode Suite")
}
