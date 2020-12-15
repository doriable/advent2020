package main

import (
	"fmt"
	"os"
)

type turn struct {
	beforeLast int
	last       int
}

func main() {
	game := map[int][]int{
		0:  []int{1},
		1:  []int{2},
		4:  []int{3},
		13: []int{4},
		15: []int{5},
		12: []int{6},
		16: []int{7},
	}

	last := 16
	for i := 8; i < 30000001; i++ {
		if turns, ok := game[last]; ok {
			if len(turns) <= 1 {
				last = 0
				game[0] = append(game[0], i)
				fmt.Printf("turn: %v, number: %v\n", i, last)
			} else {
				last = turns[len(turns)-1] - turns[len(turns)-2]
				fmt.Printf("turn: %v, number: %v\n", i, last)
				if v, ok := game[last]; ok {
					game[last] = append(v, i)
				} else {
					game[last] = []int{i}
				}
			}
		}
	}
	os.Exit(0)
}
