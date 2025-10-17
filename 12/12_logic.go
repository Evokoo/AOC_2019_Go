package day12

import (
	"fmt"
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
	p := utils.Abs(m.position.x) + utils.Abs(m.position.y) + utils.Abs(m.position.z)
	k := utils.Abs(m.velocity.x) + utils.Abs(m.velocity.y) + utils.Abs(m.velocity.z)
	return p * k
}

// ========================
// MOONS
// ========================
type Moons []*Moon

func (m *Moons) Simulate(rounds int) {
	for range rounds {
		m.PhaseOne()
		m.PhaseTwo()
	}
}

func (m *Moons) PhaseOne() {
	for i, a := range *m {
		for j, b := range *m {
			if i != j {
				a.ApplyGravity(b)
			}
		}
	}
}
func (m *Moons) PhaseTwo() {
	for _, moon := range *m {
		moon.ApplyVelocity()
	}
}
func (m *Moons) CalcuateTotalEnergy() int {
	energy := 0
	for _, moon := range *m {
		energy += moon.CalculateEnergy()
	}
	return energy
}
func (m *Moons) GetCurrentState() [3]string {
	var x, y, z strings.Builder

	for _, moon := range *m {
		fmt.Fprintf(&x, "%d,%d,", moon.position.x, moon.velocity.x)
		fmt.Fprintf(&y, "%d,%d,", moon.position.y, moon.velocity.y)
		fmt.Fprintf(&z, "%d,%d,", moon.position.z, moon.velocity.z)
	}

	return [3]string{x.String(), y.String(), z.String()}
}
func (m *Moons) PredictLoop() int {
	cycles := []int{0, 0, 0}
	orgin := m.GetCurrentState()

	for step := 1; ; step++ {
		m.PhaseOne()
		m.PhaseTwo()

		current := m.GetCurrentState()

		if cycles[0] == 0 && current[0] == orgin[0] {
			cycles[0] = step
		}
		if cycles[1] == 0 && current[1] == orgin[1] {
			cycles[1] = step
		}
		if cycles[2] == 0 && current[2] == orgin[2] {
			cycles[2] = step
		}

		if cycles[0] != 0 && cycles[1] != 0 && cycles[2] != 0 {
			break
		}
	}

	return ArrLCM(cycles)
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
			velocity: XYZ{0, 0, 0},
		})
	}

	return moons
}

func ArrLCM(values []int) int {
	result := values[0]

	for _, value := range values[1:] {
		result = utils.LCM(value, result)
	}

	return result
}
