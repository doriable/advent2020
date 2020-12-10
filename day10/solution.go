package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	raw, err := ioutil.ReadFile("example")
	if err != nil {
		os.Exit(1)
	}

	data := strings.Split(strings.TrimSpace(string(raw)), "\n")

	adapters := []int{}

	for _, line := range data {
		n, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		adapters = append(adapters, int(n))
	}

	sort.Ints(adapters)
	// device := adapters[len(adapters)-1] + 3
	jolts := 0
	count := map[int]int{}

	for _, a := range adapters {
		count[a-jolts]++
		jolts = a
	}
	count[3]++

	fmt.Println(adapters)

	set := map[int]struct{}{}
	for _, a := range adapters {
		set[a] = struct{}{}
	}

	chains := getAllChains(adapters, set)

	fmt.Printf("product of 1 diff x 3 diff: %v\n", count[1]*count[3])
	fmt.Printf("chains: %v\n", chains)
	os.Exit(0)
}

// this function gets all the combinations, but cannot run on trillions of permutations, lol
// func getAllChains(adapters map[int]struct{}, jolts, device, chains int) int {
// 	if jolts+1 == device || jolts+2 == device || jolts+3 == device {
// 		return 1
// 	}
//
// 	if _, ok := adapters[jolts+1]; ok {
// 		chains += getAllChains(adapters, jolts+1, device, 0)
// 	}
//
// 	if _, ok := adapters[jolts+2]; ok {
// 		chains += getAllChains(adapters, jolts+2, device, 0)
// 	}
//
// 	if _, ok := adapters[jolts+3]; ok {
// 		chains += getAllChains(adapters, jolts+3, device, 0)
// 	}
//
// 	return chains
// }

func getAllChains(adapters []int, set map[int]struct{}) int {
	total := 0
	skip := map[int]struct{}{}
	for i := len(adapters) - 1; i > 0; i-- {
		a := adapters[i]
		fmt.Printf("adapter: %v\n", a)

		_, three := set[a-3]
		_, two := set[a-2]
		_, one := set[a-1]

		if three && two && one {
			skip[a-1] = struct{}{}
			skip[a-2] = struct{}{}
		}

		if three && two {
			skip[a-2] = struct{}{}
		}

		if two && one {
			skip[a-1] = struct{}{}
		}

		if three && one {
			skip[a-1] = struct{}{}
		}
	}
	fmt.Println(skip)
	return total
}
