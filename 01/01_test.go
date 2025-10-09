package day01_test

import (
	"fmt"

	. "github.com/Evokoo/AOC_2019_Go/01"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Test struct {
	file   string
	part   int
	target int
}

var tests = []Test{
	{part: 1, file: "example.txt", target: 33583},
	{part: 1, file: "input.txt", target: 3256794},
	{part: 2, file: "example.txt", target: 50346},
	{part: 2, file: "input.txt", target: 4882337},
}

var _ = Describe("AOC 2019 - Day 01", func() {
	for _, test := range tests {
		msg := fmt.Sprintf("Testing Part %d with %s", test.part, test.file)
		It(msg, func() {
			result := Solve(test.file, test.part)
			Expect(result).To(Equal((test.target)))
		})
	}
})
