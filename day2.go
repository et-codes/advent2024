package main

import (
	"fmt"
)

func day2() (int, int) {
	lines := readLines("data/day2.txt")
	reports := linesToInts(lines, " ")

	return day2_part1(reports), day2_part2(reports) // 598
}

func day2_part1(reports [][]int) int {
	safeReports := 0

	for _, report := range reports {
		increasing := isIncreasing(report)
		if isSafe(report, increasing) {
			safeReports++
		}
	}

	return safeReports
}

func day2_part2(reports [][]int) int {
	safeReports := 0

	for _, report := range reports {
		increasing := isIncreasing(report)
		if isSafeDampened(report, increasing) {
			safeReports++
		}
	}

	return safeReports
}

func isIncreasing(report []int) bool {
	return report[0] < report[1]
}

func isSafeDampened(report []int, increasing bool) bool {
	min := 1
	max := 3

	fmt.Print(report)
	for i := range len(report) - 1 {
		var dampened []int
		diff := report[i+1] - report[i]
		if increasing && (diff < min || diff > max) {
			dampened = append(report[:i], report[i+1:]...)
		}
		if !increasing && (diff > min*-1 || diff < max*-1) {
			dampened = append(report[:i], report[i+1:]...)
		}
		if len(dampened) > 0 {
			safe := isSafe(dampened, isIncreasing(dampened))
			fmt.Println(" --> ", i, " --> ", dampened, " --> ", safe)
			return safe
		}
	}
	fmt.Println(" --> ", true)
	return true
}

func isSafe(report []int, increasing bool) bool {
	min := 1
	max := 3

	for i := range len(report) - 1 {
		diff := report[i+1] - report[i]
		if increasing && (diff < min || diff > max) {
			return false
		}
		if !increasing && (diff > min*-1 || diff < max*-1) {
			return false
		}
	}
	return true
}
