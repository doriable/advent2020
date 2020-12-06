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

	data := strings.Split(string(raw), "\n")

	answers := map[rune]struct{}{}
	groups := map[rune]int{}
	groupSize := 0
	total := 0
	part2 := 0

	for _, line := range data {
		if line == "" {
			total += len(answers)
			answers = map[rune]struct{}{}

			for _, answer := range groups {
				if answer == groupSize {
					part2++
				}
			}

			groupSize = 0
			groups = map[rune]int{}
			continue
		}

		groupSize++

		for _, r := range line {
			answers[rune(r)] = struct{}{}
			groups[rune(r)]++
		}
	}

	fmt.Printf("the total count across all groups: %v\n", total)
	fmt.Printf("the total number of questions where everyone answered yes : %v\n", part2)

	os.Exit(0)
}
