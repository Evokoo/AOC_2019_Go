package day08

import (
	"fmt"
	"strings"

	"github.com/Evokoo/AOC_2019_Go/utils"
)

// ========================
// PARSER
// ========================
type Layer [3]int

func CheckLayers(file string, width, height int) int {
	data := utils.ReadFile(file)
	size := width * height
	layers := make([]Layer, len(data)/size)

	for i, r := range data {
		index := i / size

		switch r {
		case '0':
			layers[index][0]++
		case '1':
			layers[index][1]++
		case '2':
			layers[index][2]++
		}
	}

	result := [2]int{size, 0}

	for _, layer := range layers {
		if layer[0] < result[0] {
			result[0] = layer[0]
			result[1] = layer[1] * layer[2]
		}
	}

	return result[1]
}

func CombineLayers(file string, width, height int) {
	data := utils.ReadFile(file)
	size := width * height
	pixels := []byte(data[len(data)-size:])

	for i := len(data) - size - 1; i >= 0; i-- {
		current := data[i]
		if current != '2' {
			pixels[i%size] = current
		}
	}

	var row strings.Builder

	for i := 0; i < len(pixels); i = i + width {
		for _, r := range pixels[i : i+width] {
			if r == '1' {
				row.WriteRune('▓')
			} else {
				row.WriteRune('░')
			}
		}
		fmt.Println(row.String())
		row.Reset()
	}
}
