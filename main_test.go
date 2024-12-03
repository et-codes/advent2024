package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadLines(t *testing.T) {
	got := readLines("data/day1-test.txt")
	assert.Contains(t, got, "3   4")
	assert.Contains(t, got, "2   5")
	assert.Contains(t, got, "3   3")
	assert.Len(t, got, 6)
}

func TestLinesToInts(t *testing.T) {
	lines := readLines("data/day1-test.txt")
	got := linesToInts(lines, "   ")
	assert.Contains(t, got, []int{3, 4})
	assert.Contains(t, got, []int{2, 5})
	assert.Contains(t, got, []int{3, 3})
	assert.Len(t, got, 6)
}
