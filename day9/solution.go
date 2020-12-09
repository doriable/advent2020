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
	raw, err := ioutil.ReadFile("input")
	if err != nil {
		os.Exit(1)
	}

	data := strings.Split(string(raw), "\n")

	last := map[int]struct{}{}
	preamble := []int{}
	all := []int{}
	attack := 0

	for _, line := range data[:25] {
		num, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			os.Exit(1)
		}
		last[int(num)] = struct{}{}
		preamble = append(preamble, int(num))
	}

	all = append(all, preamble...)

	for _, line := range data[25:] {
		num, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			os.Exit(1)
		}

		all = append(all, int(num))

		if !valid(int(num), preamble, last) {
			attack = int(num)
			break
		}

		delete(last, preamble[0])
		last[int(num)] = struct{}{}
		preamble = append(preamble[1:], int(num))
	}

	fmt.Printf("first number without a paired sum after the preamble: %v\n", attack)
	fmt.Printf("the weakness is: %v\n", findWeakness(attack, all))

	os.Exit(0)
}

func valid(n int, preamble []int, set map[int]struct{}) bool {
	for _, m := range preamble {
		if _, ok := set[n-m]; ok {
			return true
		}
	}
	return false
}

func findWeakness(attack int, all []int) int {
	for i := 0; i < len(all); i++ {
		total := 0
		r := []int{}
		for _, num := range all[i:] {
			total += num
			r = append(r, num)
			if total == attack {
				sort.Ints(r)
				return r[0] + r[len(r)-1]
			}
		}
	}
	return -1
}
