package main

import (
	"regexp"
	"strconv"
)

func day3() (int, int) {
	rawData := readLines("data/day3.txt")

	return day3_part1(rawData), day3_part2(rawData)
}

func day3_part1(data []string) int {
	instructions := []string{}
	for _, line := range data {
		re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
		matches := re.FindAllString(line, -1)
		instructions = append(instructions, matches...)
	}

	result := 0
	for _, inst := range instructions {
		result += mul(inst)
	}
	return result
}

func mul(data string) int {
	re := regexp.MustCompile(`\d{1,3}`)
	matches := re.FindAllString(data, -1)
	a, err := strconv.Atoi(matches[0])
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(matches[1])
	if err != nil {
		panic(err)
	}

	return a * b
}

func day3_part2(data []string) int {
	return 0
}
