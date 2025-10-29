package day21_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test21(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "21 Suite")
}