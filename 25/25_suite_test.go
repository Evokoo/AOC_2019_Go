package day25_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test25(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "25 Suite")
}