package intcode_test

import (
	. "github.com/Evokoo/AOC_2019_Go/intcode"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// type Test struct {
// 	// Test structure
// }

// var tests = []Test{
// 	// Test Table
// }

var _ = Describe("CPU", func() {
	It("Simple addition test", func() {
		program := []int{1, 0, 0, 0, 99}
		cpu := NewCPU(program)
		cpu.Run()
		Expect(cpu.ReadMemory(0)).To(Equal(2))
	})

	It("Simple multiplication test", func() {
		program := []int{2, 4, 4, 5, 99, 0}
		cpu := NewCPU(program)
		cpu.Run()
		Expect(cpu.ReadMemory(5)).To(Equal(9801))
	})
})
