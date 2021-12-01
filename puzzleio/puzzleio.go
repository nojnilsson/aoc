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

func ParseNumbers(nums string) []int {
	var numbers []int
	for _, a := range strings.Fields(nums) {
		i, _ := strconv.Atoi(a)
		numbers = append(numbers, i)
	}
	return numbers
}

func ReadLineOfNumbers(fname string) []int {
	lines := ReadLines(fname)
	return ParseNumbers(lines[0])
}

func ReadNumbers(fname string) []int {
	lines := ReadLines(fname)
	var numbers []int
	for l, _ := range lines {
		n, _ := strconv.Atoi(l)
		numbers = append(numbers, n)
	}
	return numbers
}

func ReadCSVNumbers(fname string, separator string) [][]int {
	lines := ReadLines(fname)
	var numberLists [][]int
	for _, l := range lines {
		var numbers []int
		for _, a := range strings.Split(l, separator) {
			i, _ := strconv.Atoi(a)
			numbers = append(numbers, i)
		}
		numberLists = append(numberLists, numbers)
	}
	return numberLists
}
