package main

import (
	"math"
	"slices"
)

func day1() (int, int) {
	rawData := readLines("data/day1.txt")
	lists := linesToInts(rawData, "   ")

	return day1_part1(lists), day1_part2(lists)
}

func day1_part1(lists [][]int) int {

	left, right := day1_SplitLists(lists)

	// Sort the lists in ascending order
	slices.Sort(left)
	slices.Sort(right)

	// Calculate the total distance for Part 1
	dist := 0
	for i := range len(left) {
		dist += int(math.Abs(float64(left[i] - right[i])))
	}

	return dist
}

func day1_part2(lists [][]int) int {

	left, right := day1_SplitLists(lists)

	similarity := 0
	counts := make(map[int]int)

	// Calculate repeats in the right list
	for _, num := range right {
		counts[num] += 1
	}

	// Calculate similarity
	for _, num := range left {
		count, found := counts[num]
		if !found {
			continue
		}
		similarity += num * count
	}

	return similarity
}

func day1_SplitLists(lists [][]int) ([]int, []int) {
	var left, right []int

	for _, pair := range lists {
		left = append(left, pair[0])
		right = append(right, pair[1])
	}

	return left, right
}
