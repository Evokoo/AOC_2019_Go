package intcode

import "fmt"

// ========================
// CPU
// ========================
type CPU struct {
	memory  []int
	address int
	input   []int
	output  []int
	active  bool
}

func NewCPU(program []int) *CPU {
	memory := make([]int, len(program))
	copy(memory, program)

	return &CPU{
		memory:  memory,
		address: 0,
		input:   []int{},
		output:  []int{},
		active:  true,
	}
}
func (c *CPU) Reset(program []int) {
	*c = *NewCPU(program)
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
	default:
		panic("Invalid Mode")
	}
}

func (c *CPU) getOpOutput(mode, offset int) int {
	address := c.getCurrentAddress() + offset
	switch mode {
	case 0, 1:
		return c.ReadFromMemory(address)
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

			c.WriteOutput(x)
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
func (c *CPU) WriteOutput(value int) {
	c.output = append(c.output, value)
}
func (c *CPU) ReadOutput() int {
	return c.output[len(c.output)-1]
}

// ========================
// DEBUG
// ========================

func (c *CPU) PrintMemory() {
	fmt.Println("Current memory contents")
	for i, value := range c.memory {
		fmt.Printf("%d at index %d\n", value, i)
	}
}
func (c *CPU) PrintOutput() {
	fmt.Println("Current output contents")
	for i, value := range c.output {
		fmt.Printf("%d at index %d\n", value, i)
	}
}
func (c *CPU) PrintInput() {
	fmt.Println("Current input contents")
	for i, value := range c.input {
		fmt.Printf("%d at index %d\n", value, i)
	}
}
