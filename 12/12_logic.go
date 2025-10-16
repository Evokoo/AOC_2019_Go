package day12

import (
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2019_Go/utils"
)

// ========================
// XYZ
// ========================
type XYZ struct{ x, y, z int }

// ========================
// MOON
// ========================
type Moon struct {
	position XYZ
	velocity XYZ
}

func (a *Moon) ApplyGravity(b *Moon) {
	if a.position.x < b.position.x {
		a.velocity.x++
	}
	if a.position.x > b.position.x {
		a.velocity.x--
	}
	if a.position.y < b.position.y {
		a.velocity.y++
	}
	if a.position.y > b.position.y {
		a.velocity.y--
	}
	if a.position.z < b.position.z {
		a.velocity.z++
	}
	if a.position.z > b.position.z {
		a.velocity.z--
	}
}
func (m *Moon) ApplyVelocity() {
	m.position.x += m.velocity.x
	m.position.y += m.velocity.y
	m.position.z += m.velocity.z
}
func (m *Moon) CalculateEnergy() int {
	p := Abs(m.position.x) + Abs(m.position.y) + Abs(m.position.z)
	k := Abs(m.velocity.x) + Abs(m.velocity.y) + Abs(m.velocity.z)
	return p * k
}

// ========================
// MOONS
// ========================
type Moons []*Moon

func (m *Moons) Simulate(rounds int) {
	for range rounds {
		for i, a := range *m {
			for j, b := range *m {
				if i != j {
					a.ApplyGravity(b)
				}
			}
		}

		for _, moon := range *m {
			moon.ApplyVelocity()
		}
	}
}

func (m *Moons) CalcuateTotalEnergy() int {
	energy := 0
	for _, moon := range *m {
		energy += moon.CalculateEnergy()
	}
	return energy
}

// ========================
// PARSER
// ========================

func ParseInput(file string) Moons {
	data := utils.ReadFile(file)

	var moons Moons
	for _, line := range strings.Split(data, "\n") {
		values := utils.QuickMatch(line, `-*\d+`)
		x, _ := strconv.Atoi(values[0])
		y, _ := strconv.Atoi(values[1])
		z, _ := strconv.Atoi(values[2])

		moons = append(moons, &Moon{
			position: XYZ{x, y, z},
			velocity: XYZ{},
		})
	}

	return moons
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
