package day15

import "fmt"

func Solve(file string) int {
	program := ParseInput(file)
	zone := NewZone()

	ScanArea(program, &zone)

	fmt.Println(zone.distance)

	return 0
}
