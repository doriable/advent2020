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

	data := strings.Split(strings.TrimSpace(string(raw)), "\n")
	total := 0

	for i, line := range data {
		tokens := strings.Split(line, " ")
		if len(tokens) != 2 {
			os.Exit(1)
		}

		switch tokens[0] {
		case "jmp":
			data[i] = strings.Join([]string{"nop", tokens[1]}, " ")

		case "nop":
			data[i] = strings.Join([]string{"jmp", tokens[1]}, " ")

		default:
			continue
		}

		if n, ok := valid(data); ok {
			total = n
			break
		}
		data = strings.Split(strings.TrimSpace(string(raw)), "\n")
	}

	fmt.Printf("total: %v\n", total)
	os.Exit(0)
}

func parseInstruction(i int, tokens []string) (int, int) {
	action, v := tokens[0], tokens[1]

	if action == "nop" {
		return 0, 1
	}

	num, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return -1, -1
	}

	if action == "acc" {
		return int(num), 1
	}

	if action == "jmp" {
		return 0, int(num)
	}

	return -1, -1
}

func valid(data []string) (int, bool) {
	seen := map[int]struct{}{}
	i := 0
	total := 0

	for {
		if _, ok := seen[i]; ok {
			return -1, false
		}

		tokens := strings.Split(data[i], " ")
		acc, n := parseInstruction(i, tokens)
		total += acc
		seen[i] = struct{}{}
		i += n
		if i >= len(data) {
			break
		}
	}

	return total, true
}
