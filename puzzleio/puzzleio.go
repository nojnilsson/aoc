package puzzleio

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(fname string) *os.File {
	f, err := os.Open(fname)
	if err != nil {
		fmt.Println("Failed to open file for reading:", fname)
		os.Exit(1)
	}
	return f
}

func fileToStrSlice(f *os.File) []string {
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func ReadLines(fname string) []string {
	return fileToStrSlice(readInput(fname))
}

func ReadLineOfNumbers(fname string) []int {
	lines := ReadLines(fname)
	var numbers []int
	for _, a := range strings.Fields(lines[0]) {
		i, _ := strconv.Atoi(a)
		numbers = append(numbers, i)
	}
	return numbers
}
