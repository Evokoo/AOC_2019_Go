package day02

import (
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2019_Go/utils"
)

// ========================
// INT MACHINE
// ========================
type CPU struct {
	memory  []int
	address int
}

func NewCPU(program []int) CPU {
	memory := make([]int, len(program))
	copy(memory, program)

	return CPU{
		memory:  memory,
		address: 0,
	}
}

func (c *CPU) Run() {
	for {
		switch c.GetValue(c.address) {
		case 1:
			c.Add()
			c.UpdateAddress(4)
		case 2:
			c.Multi()
			c.UpdateAddress(4)
		case 99:
			return
		default:
			panic("Invalid OP code")
		}
	}
}
func (c *CPU) Reset(program []int) {
	*c = NewCPU(program)
}
func (c *CPU) SetValue(value, index int) {
	c.memory[index] = value
}
func (c *CPU) GetValue(index int) int {
	return c.memory[index]
}
func (c *CPU) UpdateAddress(amount int) {
	c.address += amount
}
func (c *CPU) Add() {
	a := c.GetValue(c.GetValue(c.address + 1))
	b := c.GetValue(c.GetValue(c.address + 2))
	c.SetValue(a+b, c.GetValue(c.address+3))
}
func (c *CPU) Multi() {
	a := c.GetValue(c.GetValue(c.address + 1))
	b := c.GetValue(c.GetValue(c.address + 2))
	c.SetValue(a*b, c.GetValue(c.address+3))
}

// ========================
// PARSER
// ========================
func ParseInput(file string) []int {
	data := utils.ReadFile(file)
	var program []int
	for value := range strings.SplitSeq(data, ",") {
		n, _ := strconv.Atoi(value)
		program = append(program, n)
	}
	return program
}
