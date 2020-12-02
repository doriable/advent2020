package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type policy struct {
	min  int
	max  int
	char string
}

func main() {
	raw, err := ioutil.ReadFile("input")
	if err != nil {
		os.Exit(1)
	}

	data := strings.TrimSpace(string(raw))
	tokens := strings.Split(data, "\n")
	validOld := 0
	validNew := 0

	for _, token := range tokens {
		sub := strings.Split(token, " ")
		if len(sub) != 3 {
			os.Exit(1)
		}

		p, err := parsePolicy(sub[0], sub[1])
		if err != nil {
			os.Exit(2)
		}

		if oldPolicy(p, sub[2]) == nil {
			validOld++
		}

		if newPolicy(p, sub[2]) == nil {
			validNew++
		}
	}

	fmt.Printf("number of valid passwords, based on the old policy is: %v\n", validOld)
	fmt.Printf("number of valid passwords, based on the new policy is: %v\n", validNew)

	os.Exit(0)
}

func newPolicy(p policy, password string) error {
	if string(password[p.min-1]) != p.char && string(password[p.max-1]) != p.char {
		return fmt.Errorf("invalid password: %v for policy: %v", password, p)
	}

	if string(password[p.min-1]) == p.char && string(password[p.max-1]) == p.char {
		return fmt.Errorf("invalid password: %v for policy: %v", password, p)
	}

	return nil
}

func oldPolicy(p policy, password string) error {
	charMap := charCount(password)

	if charMap[p.char] < p.min || charMap[p.char] > p.max {
		return fmt.Errorf("invalid password: %v for policy: %v", password, p)
	}

	return nil

}

func parsePolicy(a, b string) (policy, error) {
	r := strings.Split(a, "-")
	if len(r) != 2 {
		return policy{}, fmt.Errorf("range substring did not parse, %v", r)
	}

	min, err := strconv.ParseInt(r[0], 10, 64)
	if err != nil {
		return policy{}, err
	}
	max, err := strconv.ParseInt(r[1], 10, 64)
	if err != nil {
		return policy{}, err
	}

	return policy{
		min:  int(min),
		max:  int(max),
		char: strings.TrimRight(b, ":"),
	}, nil
}

func charCount(password string) map[string]int {
	ret := map[string]int{}
	for _, char := range password {
		key := string(char)
		if _, ok := ret[key]; ok {
			ret[key]++
		} else {
			ret[key] = 1
		}
	}
	return ret
}
