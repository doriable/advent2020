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
	matrix := [][]string{}

	for _, line := range data {
		matrix = append(matrix, strings.Split(line, ""))
	}

	fmt.Printf("number of occupied seats: %v\n", findConwaySteadyState(matrix))
	os.Exit(0)
}

func findConwaySteadyState(matrix [][]string) int {
	steady := false
	for !steady {
		matrix, steady = transition(matrix)
	}

	count := 0
	for _, row := range matrix {
		for _, seat := range row {
			if seat == "#" {
				count++
			}
		}
	}
	return count
}

func transition(matrix [][]string) ([][]string, bool) {
	final := make([][]string, len(matrix))
	steady := true

	for i := range matrix {
		final[i] = make([]string, len(matrix[0]))
	}

	for i, row := range matrix {
		for j, seat := range row {
			switch seat {
			case "L":
				final[i][j] = "L"
				if countVisibleAdjOcc(matrix, i, j) == 0 {
					final[i][j] = "#"
					steady = false
				}

			case "#":
				final[i][j] = "#"
				if countVisibleAdjOcc(matrix, i, j) > 4 {
					final[i][j] = "L"
					steady = false
				}

			default:
				final[i][j] = seat
			}
		}
	}

	return final, steady
}

func countImmediateAdjacentOccupied(matrix [][]string, i, j int) int {
	c := 0
	if i != 0 {
		if matrix[i-1][j] == "#" {
			c++
		}

		if j != len(matrix[0])-1 {
			if matrix[i-1][j+1] == "#" {
				c++
			}
		}

		if j != 0 {
			if matrix[i-1][j-1] == "#" {
				c++
			}
		}
	}

	if i != len(matrix)-1 {
		if matrix[i+1][j] == "#" {
			c++
		}

		if j != len(matrix[0])-1 {
			if matrix[i+1][j+1] == "#" {
				c++
			}
		}

		if j != 0 {
			if matrix[i+1][j-1] == "#" {
				c++
			}
		}
	}

	if j != len(matrix[0])-1 {
		if matrix[i][j+1] == "#" {
			c++
		}
	}

	if j != 0 {
		if matrix[i][j-1] == "#" {
			c++
		}
	}

	return c
}

func countVisibleAdjOcc(matrix [][]string, i, j int) int {
	c := 0
	if i != 0 {
		if findVisibleUp(matrix, i, j) == "#" {
			c++
		}

		if j != len(matrix[0])-1 {
			if findVisibleUpRight(matrix, i, j) == "#" {
				c++
			}
		}

		if j != 0 {
			if findVisibleUpLeft(matrix, i, j) == "#" {
				c++
			}
		}
	}

	if i != len(matrix)-1 {
		if findVisibleDown(matrix, i, j) == "#" {
			c++
		}

		if j != len(matrix[0])-1 {
			if findVisibleDownRight(matrix, i, j) == "#" {
				c++
			}
		}

		if j != 0 {
			if findVisibleDownLeft(matrix, i, j) == "#" {
				c++
			}
		}
	}

	if j != len(matrix[0])-1 {
		if findVisibleRight(matrix, i, j) == "#" {
			c++
		}
	}

	if j != 0 {
		if findVisibleLeft(matrix, i, j) == "#" {
			c++
		}
	}

	return c
}

func findVisibleUp(matrix [][]string, i, j int) string {
	i--

	if i == 0 {
		return matrix[i][j]
	}

	if matrix[i][j] == "." {
		return findVisibleUp(matrix, i, j)
	}

	return matrix[i][j]
}

func findVisibleUpRight(matrix [][]string, i, j int) string {
	i--
	j++

	if i == 0 || j == len(matrix[0])-1 {
		return matrix[i][j]
	}

	if matrix[i][j] == "." {
		return findVisibleUpRight(matrix, i, j)
	}

	return matrix[i][j]
}

func findVisibleUpLeft(matrix [][]string, i, j int) string {
	i--
	j--

	if i == 0 || j == 0 {
		return matrix[i][j]
	}

	if matrix[i][j] == "." {
		return findVisibleUpLeft(matrix, i, j)
	}

	return matrix[i][j]
}

func findVisibleDown(matrix [][]string, i, j int) string {
	i++

	if i == len(matrix)-1 {
		return matrix[i][j]
	}

	if matrix[i][j] == "." {
		return findVisibleDown(matrix, i, j)
	}

	return matrix[i][j]
}

func findVisibleDownRight(matrix [][]string, i, j int) string {
	i++
	j++

	if i == len(matrix)-1 || j == len(matrix[0])-1 {
		return matrix[i][j]
	}

	if matrix[i][j] == "." {
		return findVisibleDownRight(matrix, i, j)
	}

	return matrix[i][j]
}

func findVisibleDownLeft(matrix [][]string, i, j int) string {
	i++
	j--

	if i == len(matrix)-1 || j == 0 {
		return matrix[i][j]
	}

	if matrix[i][j] == "." {
		return findVisibleDownLeft(matrix, i, j)
	}

	return matrix[i][j]
}

func findVisibleLeft(matrix [][]string, i, j int) string {
	j--

	if j == 0 {
		return matrix[i][j]
	}

	if matrix[i][j] == "." {
		return findVisibleLeft(matrix, i, j)
	}

	return matrix[i][j]
}

func findVisibleRight(matrix [][]string, i, j int) string {
	j++

	if j == len(matrix[0])-1 {
		return matrix[i][j]
	}

	if matrix[i][j] == "." {
		return findVisibleRight(matrix, i, j)
	}

	return matrix[i][j]
}
