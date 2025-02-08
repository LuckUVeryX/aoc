package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(4)

	keys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	passports := []map[string]string{}

	blocks := strings.Split(input, "\n\n")
	for _, block := range blocks {
		block = strings.ReplaceAll(block, "\n", " ")
		fields := strings.Fields(block)
		passport := map[string]string{}
		for _, field := range fields {
			parts := strings.Split(field, ":")
			passport[parts[0]] = parts[1]
		}
		passports = append(passports, passport)
	}

	count1 := 0
	for _, passport := range passports {
		valid := true
		for _, key := range keys {
			if _, ok := passport[key]; !ok {
				valid = false
				break
			}
		}
		if valid {
			count1++
		}
	}
	fmt.Println("part1:", count1)

	count2 := 0
	for _, passport := range passports {
		valid := true
	outer:
		for _, key := range keys {

			if _, ok := passport[key]; !ok {
				valid = false
				break
			}
			switch key {
			case "byr":
				if val, _ := strconv.Atoi(passport[key]); val < 1920 || val > 2002 {
					valid = false
				}
			case "iyr":
				if val, _ := strconv.Atoi(passport[key]); val < 2010 || val > 2020 {
					valid = false
				}
			case "eyr":
				if val, _ := strconv.Atoi(passport[key]); val < 2020 || val > 2030 {
					valid = false
				}
			case "hgt":
				var hgt int
				var unit string
				_, err := fmt.Sscanf(passport[key], "%d%s", &hgt, &unit)
				if err != nil {
					valid = false
					break outer
				}
				switch unit {
				case "cm":
					if hgt < 150 || hgt > 193 {
						valid = false
					}
				case "in":
					if hgt < 59 || hgt > 76 {
						valid = false
					}
				default:
					valid = false
				}
			case "hcl":
				if len(passport[key]) != 7 {
					valid = false
					break outer
				}
				if passport[key][0] != '#' {
					valid = false
					break outer
				}
				for _, c := range passport[key][1:] {
					if (c < '0' || c > '9') && (c < 'a' || c > 'f') {
						valid = false
						break outer
					}
				}
			case "ecl":
				ecl := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
				if !slices.Contains(ecl, passport[key]) {
					valid = false
				}
			case "pid":
				if len(passport[key]) != 9 {
					valid = false
					break outer
				}
				for _, c := range passport[key] {
					if c < '0' || c > '9' {
						valid = false
						break outer
					}
				}
			}
		}
		if valid {
			count2++
		}
	}
	fmt.Println("part2:", count2)
}

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
// hgt (Height) - a number followed by either cm or in:
// If cm, the number must be at least 150 and at most 193.
// If in, the number must be at least 59 and at most 76.
// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
// pid (Passport ID) - a nine-digit number, including leading zeroes.
// cid (Country ID) - ignored, missing or not.
