package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type field struct {
	min int
	max int
}

func main() {
	raw, err := ioutil.ReadFile("input")
	if err != nil {
		os.Exit(1)
	}

	data := strings.Split(strings.TrimSpace(string(raw)), "\n")

	rules := map[string][]field{}
	myTicket := ""
	errorRate := 0
	validTickets := []string{}
	tickets := false
	mine := false
	nearby := false

	for _, line := range data {
		if !tickets {
			if line == "" {
				tickets = true
				continue
			}
			s, f, err := parseRule(line)
			if err != nil {
				os.Exit(1)
			}
			rules[s] = f
			continue
		}

		if !mine {
			if line == "your ticket:" {
				continue
			}
			if line == "" {
				mine = true
				continue
			}
			myTicket = line
			continue
		}

		if !nearby {
			if line == "nearby tickets:" {
				continue
			}
			if line == "" {
				nearby = true
			}
			n, err := ticketErrorRate(line, rules)
			if err != nil {
				os.Exit(2)
			}
			errorRate += n
			if n == 0 {
				validTickets = append(validTickets, line)
			}
		}
	}

	n, err := getDepartureProduct(myTicket, validTickets, rules)
	if err != nil {
		os.Exit(3)
	}

	fmt.Printf("the error rate: %v\n", errorRate)
	fmt.Printf("the departure fields product: %v\n", n)

	os.Exit(0)
}

func parseRule(raw string) (string, []field, error) {
	tokens := strings.Split(raw, ":")
	if len(tokens) != 2 {
		return "", []field{}, fmt.Errorf("invalid rule string: %v", raw)
	}

	fields := []field{}
	ranges := strings.Split(tokens[1], "or")
	for _, r := range ranges {
		nums := strings.Split(strings.TrimSpace(r), "-")

		if len(nums) != 2 {
			return "", []field{}, fmt.Errorf("invalid range: %v", r)
		}

		min, err := strconv.ParseInt(nums[0], 10, 64)
		if err != nil {
			return "", []field{}, err
		}
		max, err := strconv.ParseInt(nums[1], 10, 64)
		if err != nil {
			return "", []field{}, err
		}

		fields = append(fields, field{
			min: int(min),
			max: int(max),
		})
	}

	return tokens[0], fields, nil
}

func ticketErrorRate(raw string, rules map[string][]field) (int, error) {
	values := strings.Split(raw, ",")
	ret := 0

	for _, v := range values {
		num, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, err
		}

		if !validField(int(num), rules) {
			ret += int(num)
		}
	}

	return ret, nil
}

func validField(n int, rules map[string][]field) bool {
	for _, fields := range rules {
		for _, f := range fields {
			if n >= f.min && n <= f.max {
				return true
			}
		}
	}
	return false
}

func getDepartureProduct(myTicket string, validTickets []string, rules map[string][]field) (int, error) {
	ticketMatrix := [][]int{}

	for _, ticket := range validTickets {
		row := []int{}
		tokens := strings.Split(ticket, ",")
		for _, token := range tokens {
			n, err := strconv.ParseInt(token, 10, 64)
			if err != nil {
				return 0, err
			}
			row = append(row, int(n))
		}
		ticketMatrix = append(ticketMatrix, row)
	}

	ticketMap := map[string]int{}
	ticketMapCheck := map[string]struct{}{}
	for j := 0; j < len(ticketMatrix[0]); j++ {
		validClasses := map[string]int{}
		for i := 0; i < len(ticketMatrix); i++ {
			fmt.Println(ticketMatrix[i][j])
			for class, fields := range rules {
				for _, f := range fields {
					if ticketMatrix[i][j] >= f.min && ticketMatrix[i][j] <= f.max {
						validClasses[class]++
						break
					}
				}
			}
		}

		fmt.Println(validClasses)

		for class, count := range validClasses {
			if count == len(ticketMatrix) {
				if _, ok := ticketMapCheck[class]; ok {
					continue
				}
				ticketMapCheck[class] = struct{}{}
				ticketMap[class] = j
			}
		}
	}

	fmt.Println(ticketMap)

	myTicketArray := strings.Split(myTicket, ",")
	myTicketNums := []int{}
	for _, t := range myTicketArray {
		n, err := strconv.ParseInt(t, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid field in my ticket: %v", t)
		}
		myTicketNums = append(myTicketNums, int(n))
	}

	ret := 1
	for class, i := range ticketMap {
		if strings.Contains(class, "departure") {
			ret *= myTicketNums[i]
		}
	}

	return ret, nil
}
