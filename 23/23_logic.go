package day23

import (
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2019_Go/intcode"
	"github.com/Evokoo/AOC_2019_Go/utils"
)

// ========================
// PACKET
// ========================
type Packet [2]int
type Queue []Packet

func NewQueue() Queue {
	return make(Queue, 0)
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
func (q *Queue) Pop() Packet {
	packet := (*q)[0]
	*q = (*q)[1:]
	return packet
}
func (q *Queue) Push(packet Packet) {
	*q = append(*q, packet)
}

// ========================
// NETWORK
// ========================
type Computer struct {
	cpu   *intcode.CPU
	queue Queue
}

func (c *Computer) ProcessPacket() {
	if c.queue.IsEmpty() {
		c.cpu.ReadInput(-1)
	} else {
		packet := c.queue.Pop()
		c.cpu.ReadInput(packet[0])
		c.cpu.ReadInput(packet[1])
	}
	c.cpu.Run()
}

func (c *Computer) SendPackets(network Network) (bool, Packet) {
	recieved := c.cpu.DumpOutput()
	for i := 0; i < len(recieved); i += 3 {
		data := recieved[i : i+3]

		// fmt.Println(data)

		if data[0] == 255 {
			return true, Packet{data[1], data[2]}
		}

		computer := network[data[0]]
		computer.queue.Push(Packet{data[1], data[2]})
	}

	return false, Packet{}
}

// ========================
// NETWORK
// ========================
type Network map[int]*Computer

func InitNetwork(program []int, size int) Network {
	network := make(Network)

	for id := range size {
		cpu := intcode.NewCPU(program)
		cpu.ReadInput(id)
		network[id] = &Computer{cpu, NewQueue()}
	}
	return network
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
