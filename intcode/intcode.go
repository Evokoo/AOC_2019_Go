package intcode

import "fmt"

// ========================
// CPU
// ========================
type CPU struct {
	memory  []int
	address int
	input   Queue
	output  Queue
}

func NewCPU(program []int) *CPU {
	memory := make([]int, len(program))
	copy(memory, program)

	return &CPU{
		memory:  memory,
		address: 0,
		input:   NewQueue(),
		output:  NewQueue(),
	}
}
func (c *CPU) Reset(program []int) {
	*c = *NewCPU(program)
}

// ========================
// OPERATION
// ========================
type Param struct {
	value int
	mode  int
}

func (c *CPU) Run() {
	fmt.Println("Program Start...")
	for {
		opcode, params, step := c.readInstruction()

		switch opcode {
		case 1:
			c.add(params, step)
		case 2:
			c.multiply(params, step)
		case 3:
			c.write_from_input(params, step)
		case 4:
			c.write_to_output(params, step)
		case 5:
			c.jump_true(params, step)
		case 6:
			c.jump_false(params, step)
		case 7:
			c.less_than(params, step)
		case 8:
			c.equal_to(params, step)
		case 99:
			fmt.Println("Program Halt...")
			return
		default:
			fmt.Printf("Invalid OP Code: %d\n", opcode)
			panic("Invalid OP code")
		}
	}
}
func (c *CPU) parseInstruction() (int, int) {
	instruction := c.GetMemoryValue(c.address)
	code := instruction % 100
	modeData := instruction / 100
	return code, modeData
}
func (c *CPU) readInstruction() (int, [3]Param, int) {
	opcode, modeData := c.parseInstruction()
	count := 3

	switch opcode {
	case 3, 4:
		count = 1
	case 5, 6:
		count = 2
	case 99:
		count = 0
	}

	var params [3]Param
	for i := 0; i < count; i++ {
		params[i].value = c.GetMemoryValue(c.address + i + 1)
		params[i].mode = modeData % 10
		modeData /= 10
	}

	if opcode == 1 || opcode == 2 || opcode == 3 || opcode == 7 || opcode == 8 {
		last := count - 1
		params[last].mode = 0
	}

	return opcode, params, count + 1
}
func (c *CPU) getParamValue(p Param) int {
	if p.mode == 0 {
		return c.GetMemoryValue(p.value)
	}
	return p.value
}

// Opcode #1 Add
func (c *CPU) add(params [3]Param, step int) {
	c.SetMemoryValue(c.getParamValue(params[0])+c.getParamValue(params[1]), params[2].value)
	c.updateAddress(step)
}

// Opcode #2 Multiply
func (c *CPU) multiply(params [3]Param, step int) {
	c.SetMemoryValue(c.getParamValue(params[0])*c.getParamValue(params[1]), params[2].value)
	c.updateAddress(step)
}

// Opcode #3 Set Value from Input
func (c *CPU) write_from_input(params [3]Param, step int) {
	c.SetMemoryValue(c.popFromInput(), params[0].value)
	c.updateAddress(step)
}

// Opcode #4 Write Value to Output
func (c *CPU) write_to_output(params [3]Param, step int) {
	c.PushToOutput(c.getParamValue(params[0]))
	c.updateAddress(step)
}

// Opcode #5 Jump if p[0] != 0
func (c *CPU) jump_true(params [3]Param, step int) {
	if c.getParamValue(params[0]) != 0 {
		c.setAddress(c.getParamValue(params[1]))
	} else {
		c.updateAddress(step)
	}
}

// Opcode #6 Jump if p[0] == 0
func (c *CPU) jump_false(params [3]Param, step int) {
	if c.getParamValue(params[0]) == 0 {
		c.setAddress(c.getParamValue(params[1]))
	} else {
		c.updateAddress(step)
	}
}

// Opcode #7 Set p[2] to 1 if p[0] < p[1] else 0
func (c *CPU) less_than(params [3]Param, step int) {
	value := 0
	if c.getParamValue(params[0]) < c.getParamValue(params[1]) {
		value = 1
	}
	c.SetMemoryValue(value, params[2].value)
	c.updateAddress(step)
}

// Opcode #8 Set p[2] to 1 if p[0] == p[1] else 0
func (c *CPU) equal_to(params [3]Param, step int) {
	value := 0
	if c.getParamValue(params[0]) == c.getParamValue(params[1]) {
		value = 1
	}
	c.SetMemoryValue(value, params[2].value)
	c.updateAddress(step)
}

// ========================
// GETTERS & SETTERS
// ========================

// Memory
func (c *CPU) GetMemory() []int {
	return c.memory
}
func (c *CPU) SetMemoryValue(newValue, address int) {
	c.memory[address] = newValue
}
func (c *CPU) GetMemoryValue(address int) int {
	return c.memory[address]
}

// Address
func (c *CPU) GetAddres() int {
	return c.address
}
func (c *CPU) setAddress(value int) {
	c.address = value
}
func (c *CPU) updateAddress(value int) {
	c.address += value
}

// ========================
// I/O
// ========================

// Input
func (c *CPU) PushToInput(value int) {
	c.input.Push(value)
}
func (c *CPU) popFromInput() int {
	return c.input.Pop()
}
func (c *CPU) PrintInput() {
	for i, value := range c.input {
		fmt.Printf("%d at index %d\n", value, i)
	}
}

// Output
func (c *CPU) PushToOutput(value int) {
	c.output.Push(value)
}
func (c *CPU) popFromOutput() int {
	return c.output.Pop()
}
func (c *CPU) PrintOutput() {
	for i, value := range c.output {
		fmt.Printf("%d at index %d\n", value, i)
	}
}
func (c *CPU) GetOutput() Queue {
	return c.output
}
