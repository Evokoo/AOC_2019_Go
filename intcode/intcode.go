package intcode

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
	for {
		opcode, params := c.ReadInstruction()

		switch opcode {
		case 1:
			c.Add(params)
		case 2:
			c.Multi(params)
		case 3:
			c.Read(params)
		case 4:
			c.Write(params)
		case 99:
			return
		default:
			panic("Invalid OP code")
		}
	}
}

func (c *CPU) Add(params [3]Param) {
	a := c.GetParamValue(params[0])
	b := c.GetParamValue(params[1])
	c.SetValue(a+b, params[2].value)
	c.UpdateAddress(4)
}
func (c *CPU) Multi(params [3]Param) {
	a := c.GetParamValue(params[0])
	b := c.GetParamValue(params[1])
	c.SetValue(a*b, params[2].value)
	c.UpdateAddress(4)
}
func (c *CPU) Read(params [3]Param) {
	c.SetValue(c.ReadFromInput(), params[0].value)
	c.UpdateAddress(2)
}
func (c *CPU) Write(params [3]Param) {
	c.WriteToOutput(c.GetValue(params[0].value))
	c.UpdateAddress(2)
}

// ========================
// GETTERS & SETTERS
// ========================
func (c *CPU) SetValue(value, index int) {
	c.memory[index] = value
}
func (c *CPU) GetValue(index int) int {
	return c.memory[index]
}
func (c *CPU) UpdateAddress(amount int) {
	c.address += amount
}
func (c *CPU) ReadInstruction() (int, [3]Param) {
	instruction := c.GetValue(c.address)

	opcode := instruction % 100
	paramModes := instruction / 100

	var params [3]Param
	for i := 0; i < 3; i++ {
		params[i].value = c.GetValue(c.address + i + 1)
		params[i].mode = paramModes % 10
		paramModes /= 10
	}

	return opcode, params
}
func (c *CPU) GetParamValue(p Param) int {
	if p.mode == 0 {
		return c.GetValue(p.value)
	}
	return p.value
}

// ========================
// I/O
// ========================
func (c *CPU) WriteToInput(value int) {
	c.input.Push(value)
}
func (c *CPU) WriteToOutput(value int) {
	c.output.Push(value)
}
func (c *CPU) ReadFromInput() int {
	return (*c).input.Pop()
}
func (c *CPU) ReadFromOutput() int {
	return (*c).output.Pop()
}
func (c *CPU) GetOutput() Queue {
	return c.output
}

// ========================
// QUEUE
// ========================
type Queue []int

func NewQueue() Queue {
	return []int{}
}
func (q *Queue) Push(value int) {
	*q = append(*q, value)
}
func (q *Queue) Pop() int {
	removed := (*q)[0]
	*q = (*q)[1:]
	return removed
}
