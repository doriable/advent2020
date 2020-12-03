package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	raw, err := ioutil.ReadFile("input")
	if err != nil {
		os.Exit(1)
	}

	data := strings.Split(strings.TrimSpace(string(raw)), "\n")
	matrix := [][]rune{}
	for _, line := range data {
		row := []rune{}
		for _, r := range line {
			row = append(row, r)
		}
		matrix = append(matrix, row)
	}

	r3d1 := checkPath(matrix, 3, 1)
	r1d1 := checkPath(matrix, 1, 1)
	r5d1 := checkPath(matrix, 5, 1)
	r7d1 := checkPath(matrix, 7, 1)
	r1d2 := checkPath(matrix, 1, 2)

	fmt.Printf("right 3 down 1: %v\n", r3d1)
	fmt.Printf("product of all: %v\n", r3d1*r1d1*r5d1*r7d1*r1d2)

	os.Exit(0)
}

func checkPath(matrix [][]rune, x, y int) int {
	trees := 0
	j := 0
	for i := 0; i < len(matrix)-y; i += y {
		if i+y > len(matrix) {
			break
		}
		j += x
		next := matrix[i+y][j%len(matrix[i+y])]
		if next == '#' {
			trees++
		}
	}
	return trees
}
