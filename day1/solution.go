package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	raw, err := ioutil.ReadFile("input")
	if err != nil {
		os.Exit(1)
	}

	data := strings.TrimSpace(string(raw))
	input := []int64{}

	for _, line := range strings.Split(data, "\n") {
		num, err := strconv.ParseInt(strings.TrimSpace(line), 10, 64)
		if err != nil {
			os.Exit(2)
		}
		input = append(input, num)
	}

	for i, a := range input {
		ret := checkDifference(input, i, a)
		if ret != int64(-1) {
			fmt.Printf("the entries are: %v, %v. the product is %v\n", a, ret, a*ret)
			os.Exit(0)
		}
	}
}

// checkDifference takes the set of numbers, the number being checking, and the index of
// this number. It takes the difference of the number from 2020, then scans the rest of the
// list for the difference from 2020. It skips the duplicate entry based on the index.
// If there is no other entry that adds up to 2020 with the number, -1 is returned.
func checkDifference(set []int64, i int, a int64) int64 {
	ret := int64(-1)
	diff := 2020 - a
	for j, b := range set {
		if i == j {
			continue
		}
		if b == diff {
			ret = b
			break
		}
	}
	return ret
}
