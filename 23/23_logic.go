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
// COMPUTER
// ========================
type Computer struct {
	cpu   *intcode.CPU
	queue Queue
	idle  bool
}

func (c *Computer) ProcessPacket() {
	if c.queue.IsEmpty() {
		c.cpu.ReadInput(-1)
		c.idle = true

	} else {
		packet := c.queue.Pop()
		c.cpu.ReadInput(packet[0])
		c.cpu.Run()
		c.cpu.ReadInput(packet[1])
		c.idle = false
	}

	c.cpu.Run()

}

func (c *Computer) SendPackets(network Network) (bool, Packet) {
	toSend := c.cpu.DumpOutput()
	is255 := false
	output := Packet{}

	for i := 0; i < len(toSend); i += 3 {
		data := toSend[i : i+3]

		if data[0] == 255 {
			is255 = true
			output = Packet{data[1], data[2]}
			continue
		}

		computer := network.computers[data[0]]
		computer.queue.Push(Packet{data[1], data[2]})
		// computer.idle = false
	}

	return is255, output
}

// ========================
// NETWORK
// ========================
type Network struct {
	computers map[int]*Computer
	NAT       Packet
	sent      int
}

func InitNetwork(program []int, size int) Network {
	computers := make(map[int]*Computer)

	for id := range size {
		cpu := intcode.NewCPU(program)
		cpu.ReadInput(id)
		computers[id] = &Computer{cpu, NewQueue(), false}
	}
	return Network{computers: computers, NAT: Packet{}}
}

func (n *Network) IsIdle() bool {
	for _, computer := range n.computers {
		if !computer.queue.IsEmpty() || !computer.idle {
			return false
		}
	}
	return true
}

func (n *Network) Reset() (bool, int) {
	if n.NAT[1] == n.sent {
		return true, n.sent
	}

	computer := n.computers[0]
	computer.queue.Push(n.NAT)
	n.sent = n.NAT[1]

	return false, 0
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
