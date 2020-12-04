package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	raw, err := ioutil.ReadFile("input")
	if err != nil {
		os.Exit(1)
	}

	data := strings.Split(string(raw), "\n")
	passport := map[string]string{}
	valid := 0

	for _, line := range data {
		if line == "" {
			if validPassport(passport) {
				if validData(passport) == nil {
					valid++
				}
			}
			passport = map[string]string{}
			continue
		}
		tokens := strings.Split(line, " ")
		for _, token := range tokens {
			kv := strings.Split(token, ":")
			if len(kv) != 2 {
				os.Exit(2)
			}
			passport[kv[0]] = kv[1]
		}
	}

	fmt.Printf("there are %v valid passports/north pole credentials\n", valid)
	os.Exit(0)
}

func validPassport(passport map[string]string) bool {
	if len(passport) == 8 {
		return true
	}
	if _, ok := passport["cid"]; !ok && len(passport) == 7 {
		return true
	}
	return false
}

func validData(passport map[string]string) error {
	byr, err := strconv.ParseInt(passport["byr"], 10, 64)
	if err != nil {
		return err
	}
	if !validRange(byr, int64(1920), int64(2002)) {
		return fmt.Errorf("invalid byr: %v", byr)
	}

	iyr, err := strconv.ParseInt(passport["iyr"], 10, 64)
	if err != nil {
		return err
	}
	if !validRange(iyr, int64(2010), int64(2020)) {
		return fmt.Errorf("invalid iyr: %v", iyr)
	}

	eyr, err := strconv.ParseInt(passport["eyr"], 10, 64)
	if err != nil {
		return err
	}
	if !validRange(eyr, int64(2020), int64(2030)) {
		return fmt.Errorf("invalid eyr: %v", eyr)
	}

	if !validHeight(passport["hgt"]) {
		return fmt.Errorf("invalid height: %v", passport["hgt"])
	}

	if !validHCL(passport["hcl"]) {
		return fmt.Errorf("invalid hair colour: %v", passport["hcl"])
	}

	if !validECL(passport["ecl"]) {
		return fmt.Errorf("invalid eye colour: %v", passport["ecl"])
	}

	if len(passport["pid"]) != 9 {
		return fmt.Errorf("invalid pid: %v", passport["pid"])
	}

	return nil
}

func validRange(d, min, max int64) bool {
	if d < min || d > max {
		return false
	}
	return true
}

func validHeight(height string) bool {
	unit := height[len(height)-2 : len(height)]
	if unit == "cm" {
		n, err := strconv.ParseInt(height[:len(height)-2], 10, 64)
		if err != nil {
			return false
		}
		return validRange(n, int64(150), int64(193))
	}
	if unit == "in" {
		n, err := strconv.ParseInt(height[:len(height)-2], 10, 64)
		if err != nil {
			return false
		}
		return validRange(n, int64(59), int64(76))
	}
	return false
}

func validHCL(hcl string) bool {
	if hcl[0] != '#' {
		return false
	}

	if m, err := regexp.Match("[0-9a-f]{6}", []byte(hcl[1:])); err != nil || !m {
		return false
	}

	return true
}

func validECL(ecl string) bool {
	if ecl == "amb" || ecl == "blu" || ecl == "brn" || ecl == "gry" || ecl == "grn" || ecl == "hzl" || ecl == "oth" {
		return true
	}
	return false
}
