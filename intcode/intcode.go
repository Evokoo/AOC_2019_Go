package intcode

import (
	"fmt"
)

// ========================
// CPU
// ========================
type CPU struct {
	memory  map[int]int
	address int
	input   []int
	output  []int
	active  bool
	base    int
}

func initialiseMemory(program []int) map[int]int {
	memory := make(map[int]int)
	for i, value := range program {
		memory[i] = value
	}
	return memory
}

func NewCPU(program []int) *CPU {
	return &CPU{
		memory: initialiseMemory(program),
		input:  []int{},
		output: []int{},
		active: true,
	}
}
func (c *CPU) Reset(program []int) {
	*c = *NewCPU(program)
}
func (c *CPU) Clone() *CPU {
	memCopy := make(map[int]int, len(c.memory))
	for k, v := range c.memory {
		memCopy[k] = v
	}

	return &CPU{
		memory:  memCopy,
		address: c.address,
		active:  c.active,
		base:    c.base,
	}
}

// ========================
// OPERATION
// ========================
func (c *CPU) readInstruction() (int, [3]int) {
	instruction := c.ReadFromMemory(c.getCurrentAddress())
	opcode := instruction % 100
	modes := [3]int{
		(instruction / 100) % 10,
		(instruction / 1000) % 10,
		(instruction / 10000) % 10,
	}
	return opcode, modes
}

func (c *CPU) getOpInput(mode, offset int) int {
	address := c.getCurrentAddress() + offset

	switch mode {
	case 0:
		return c.ReadFromMemory(c.ReadFromMemory(address))
	case 1:
		return c.ReadFromMemory(address)
	case 2:
		return c.ReadFromMemory(c.ReadFromMemory(address) + c.currentBase())
	default:
		panic("Invalid Mode")
	}
}

func (c *CPU) getOpOutput(mode, offset int) int {
	address := c.getCurrentAddress() + offset
	switch mode {
	case 0, 1:
		return c.ReadFromMemory(address)
	case 2:
		return c.ReadFromMemory(address) + c.currentBase()
	default:
		panic("Invalid Mode")
	}
}

func (c *CPU) Run() {
	for {
		opcode, modes := c.readInstruction()

		switch opcode {
		case 1: // ADD
			x := c.getOpInput(modes[0], 1)
			y := c.getOpInput(modes[1], 2)
			z := c.getOpOutput(modes[2], 3)

			c.WriteToMemory(x+y, z)
			c.incrementAddress(4)

		case 2: // MULTIPLY
			x := c.getOpInput(modes[0], 1)
			y := c.getOpInput(modes[1], 2)
			z := c.getOpOutput(modes[2], 3)

			c.WriteToMemory(x*y, z)
			c.incrementAddress(4)

		case 3: // READ INPUT
			if len(c.input) == 0 {
				// fmt.Println("No Inputs found")
				return
			}
			x := c.writeInput()
			y := c.getOpOutput(modes[0], 1)
			c.WriteToMemory(x, y)
			c.incrementAddress(2)

		case 4: // PRINT TO OUTPUT
			x := c.getOpInput(modes[0], 1)

			c.writeOutput(x)
			c.incrementAddress(2)

		case 5: // JUMP IF NOT ZERO
			x := c.getOpInput(modes[0], 1)
			y := c.getOpInput(modes[1], 2)

			if x != 0 {
				c.setAddress(y)
			} else {
				c.incrementAddress(3)
			}

		case 6: // JUMP IF ZERO
			x := c.getOpInput(modes[0], 1)
			y := c.getOpInput(modes[1], 2)

			if x == 0 {
				c.setAddress(y)
			} else {
				c.incrementAddress(3)
			}

		case 7: // LESS THAN
			x := c.getOpInput(modes[0], 1)
			y := c.getOpInput(modes[1], 2)
			z := c.getOpOutput(modes[2], 3)
			v := 0

			if x < y {
				v = 1
			}
			c.WriteToMemory(v, z)
			c.incrementAddress(4)

		case 8: // EQUAL
			x := c.getOpInput(modes[0], 1)
			y := c.getOpInput(modes[1], 2)
			z := c.getOpOutput(modes[2], 3)
			v := 0

			if x == y {
				v = 1
			}
			c.WriteToMemory(v, z)
			c.incrementAddress(4)

		case 9: //ADJUST BASE
			x := c.getOpInput(modes[0], 1)
			c.adjustBase(x)
			c.incrementAddress(2)
		case 99: // HALT
			c.active = false
			return
		default:
			panic(fmt.Sprintf("Invalid opcode: %d", opcode))
		}
	}
}

// ========================
// GETTERS & SETTERS
// ========================

// Memory
func (c *CPU) WriteToMemory(value, address int) {
	c.memory[address] = value
}
func (c *CPU) ReadFromMemory(address int) int {
	return c.memory[address]
}

// Address
func (c *CPU) getCurrentAddress() int {
	return c.address
}
func (c *CPU) setAddress(value int) {
	c.address = value
}
func (c *CPU) incrementAddress(value int) {
	c.address += value
}

// Base
func (c *CPU) currentBase() int {
	return c.base
}
func (c *CPU) adjustBase(amount int) {
	c.base += amount
}

// Activity
func (c *CPU) IsActive() bool {
	return c.active
}

// ========================
// I/O
// ========================

// Input
func (c *CPU) ReadInput(value int) {
	c.input = append(c.input, value)
}
func (c *CPU) writeInput() int {
	value := c.input[0]
	c.input = c.input[1:]
	return value
}

// Output
func (c *CPU) writeOutput(value int) {
	c.output = append(c.output, value)
}

func (c *CPU) DumpOutput() []int {
	return c.output
}

func (c *CPU) ClearOutput() {
	c.output = c.output[:0]
}

func (c *CPU) ReadOutput() int {
	length := len(c.output)
	removed := c.output[length-1]
	c.output = c.output[:length-1]
	return removed
}

func (c *CPU) HasOutput() bool {
	return len(c.output) > 0
}

// ========================
// DEBUG
// ========================

func (c *CPU) PrintMemory() {
	if len(c.memory) == 0 {
		fmt.Println("Memory in empty")
	} else {
		fmt.Println("Current memory contents:")
		for i, value := range c.memory {
			fmt.Printf("%d at index %d\n", value, i)
		}
	}
}
func (c *CPU) PrintOutput() {
	if len(c.output) == 0 {
		fmt.Println("Output in empty")
	} else {
		fmt.Println("Current output contents:")
		for i, value := range c.output {
			fmt.Printf("%d at index %d\n", value, i)
		}
	}
}
func (c *CPU) PrintInput() {
	if len(c.input) == 0 {
		fmt.Println("Input in empty")
	} else {
		fmt.Println("Current input contents:")
		for i, value := range c.input {
			fmt.Printf("%d at index %d\n", value, i)
		}
	}
}
