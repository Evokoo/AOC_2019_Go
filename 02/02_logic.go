package day02

import (
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2019_Go/utils"
)

// ========================
// PROGRAM
// ========================
type Program []int

func (p Program) Run() {
	for i := 0; i < len(p); i = i + 4 {
		switch p[i] {
		case 1:
			p.Add(p[i+1], p[i+2], p[i+3])
		case 2:
			p.Multi(p[i+1], p[i+2], p[i+3])
		case 99:
			return
		default:
			panic("Invalid OP code")
		}
	}
}

func (p *Program) SetValue(value, index int) {
	(*p)[index] = value
}
func (p *Program) GetValue(index int) int {
	return (*p)[index]
}
func (p *Program) Add(a, b, target int) {
	a, b = (*p).GetValue(a), (*p).GetValue(b)
	(*p).SetValue(a+b, target)
}
func (p *Program) Multi(a, b, target int) {
	a, b = (*p).GetValue(a), (*p).GetValue(b)
	(*p).SetValue(a*b, target)
}

// ========================
// PARSER
// ========================
func ParseInput(file string) Program {
	data := utils.ReadFile(file)
	var program Program
	for value := range strings.SplitSeq(data, ",") {
		n, _ := strconv.Atoi(value)
		program = append(program, n)
	}
	return program
}
