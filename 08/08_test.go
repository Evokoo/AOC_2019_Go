package day08_test

import (
	"fmt"

	. "github.com/Evokoo/AOC_2019_Go/08"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Test struct {
	file   string
	part   int
	target int
}

var tests = []Test{
	// {part: 1, file: "example.txt", target: -1},
	{part: 1, file: "input.txt", target: 1560},
	// {part: 2, file: "example.txt", target: -1},
	{part: 2, file: "input.txt", target: -1},
}

var _ = Describe("AOC 2019 - Day 08", func() {
	for _, test := range tests {
		msg := fmt.Sprintf("Testing Part %d with %s", test.part, test.file)
		It(msg, func() {
			result := Solve(test.file, test.part)
			Expect(result).To(Equal((test.target)))
		})
	}
})
