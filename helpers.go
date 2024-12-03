package main

import (
	"os"
	"strconv"
	"strings"
)

func linesToInts(lines []string, sep string) [][]int {
	if sep == "" {
		sep = " "
	}

	out := make([][]int, len(lines))

	for i, line := range lines {
		intStrings := strings.Split(line, sep)
		for _, intStr := range intStrings {
			num, err := strconv.Atoi(intStr)
			if err != nil {
				panic(err)
			}
			out[i] = append(out[i], num)
		}
	}

	return out
}

func readLines(fileName string) []string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")

	return lines
}
