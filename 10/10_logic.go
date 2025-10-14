package day10

import (
	"math"
	"sort"
	"strings"

	"github.com/Evokoo/AOC_2019_Go/utils"
)

// ========================
// POINT
// ========================
type Point struct{ x, y int }

func (a *Point) ProjectFromPoint(b Point, grid *Grid) {
	vector := a.VectorTo(b)
	pos := Point{a.x + vector.x, a.y + vector.y}

	for pos.x >= 0 && pos.x < grid.dimensions.x && pos.y >= 0 && pos.y < grid.dimensions.y {
		if _, found := grid.astroids[pos]; found {
			grid.astroids[*a][vector] = append(grid.astroids[*a][vector], pos)
		}
		pos.x += vector.x
		pos.y += vector.y
	}
}

func (a *Point) VectorTo(b Point) Point {
	dx := b.x - a.x
	dy := b.y - a.y
	g := GCD(Abs(dx), Abs(dy))
	return Point{dx / g, dy / g}
}

// ========================
// GRID
// ========================

type Grid struct {
	astroids   map[Point]map[Point][]Point
	dimensions Point
	optimal    map[Point][]Point
	visibility int
}

func (g *Grid) ScanAstroids() {
	for origin := range g.astroids {
		for target := range g.astroids {
			if origin == target {
				continue
			}
			origin.ProjectFromPoint(target, g)
		}

		count := len(g.astroids[origin])
		if count > g.visibility {
			g.visibility = count
			g.optimal = g.astroids[origin]
		}
	}
}

func (g *Grid) DestoryAstroids(n int) int {
	sortedVectors := SortVectors(g.optimal)
	count := 0
	for {
		for _, vector := range sortedVectors {
			astroids := g.optimal[vector]

			if len(astroids) > 0 {
				astroid := astroids[0]
				g.optimal[vector] = astroids[1:]
				count++

				if count == n {
					return astroid.x*100 + astroid.y
				}
			}
		}
	}

}

func SortVectors(vectorMap map[Point][]Point) []Point {
	var vectors []Point
	for vector := range vectorMap {
		vectors = append(vectors, vector)
	}
	sort.Slice(vectors, func(i, j int) bool {
		return ConvertToAngle(vectors[i]) < ConvertToAngle(vectors[j])
	})
	return vectors
}

func ConvertToAngle(vector Point) float64 {
	angle := math.Atan2(float64(vector.x), float64(-vector.y)) * 180 / math.Pi
	if angle < 0 {
		angle += 360
	}
	return angle
}

// ========================
// PARSER
// ========================

func ParseInput(file string) Grid {
	data := utils.ReadFile(file)
	rows := strings.Split(data, "\n")

	grid := Grid{
		astroids:   make(map[Point]map[Point][]Point),
		dimensions: Point{len(rows), len(rows[0])},
	}

	for y, row := range rows {
		for x, r := range row {
			if r == '#' {
				grid.astroids[Point{x, y}] = make(map[Point][]Point)
			}
		}
	}
	return grid
}

// ========================
// MATH
// ========================

func GCD(a, b int) int {
	if b == 0 {
		if a < 0 {
			return -a
		}
		return a
	}
	return GCD(b, a%b)
}
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
