package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	example = "FBFBBFFRLR"
)

func main() {
	raw, err := ioutil.ReadFile("input")
	if err != nil {
		os.Exit(1)
	}

	data := strings.Split(strings.TrimSpace(string(raw)), "\n")
	max := 0

	rowArray := []int{}
	colArray := []int{}
	seats := []int{}

	for i := 0; i < 128; i++ {
		rowArray = append(rowArray, i)
	}

	for i := 0; i < 8; i++ {
		colArray = append(colArray, i)
	}

	for _, line := range data {
		row := getBinaryPartition(line[:7], rowArray, 'F', 'B')
		col := getBinaryPartition(line[7:], colArray, 'L', 'R')

		id := row*8 + col
		seats = append(seats, id)
		if id > max {
			max = id
		}
	}

	fmt.Printf("the highest seat id: %v\n", max)

	idSet := map[int]struct{}{}

	for i := 0; i < 128; i++ {
		for j := 0; j < 8; j++ {
			idSet[i*8+j] = struct{}{}
		}
	}

	for _, seat := range seats {
		if _, ok := idSet[seat]; ok {
			delete(idSet, seat)
		}
	}

	fmt.Println(idSet)
	os.Exit(0)
}

func getBinaryPartition(token string, seatArray []int, lower, upper rune) int {
	if len(seatArray) == 1 {
		return seatArray[0]
	}

	if rune(token[0]) == lower {
		return getBinaryPartition(token[1:], seatArray[:len(seatArray)/2], lower, upper)
	}

	if rune(token[0]) == upper {
		return getBinaryPartition(token[1:], seatArray[len(seatArray)/2:], lower, upper)
	}

	return -1
}
