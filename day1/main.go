package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var digits = []struct {
	literal string
	val     int
}{
	{literal: "zero", val: 0},
	{literal: "one", val: 1},
	{literal: "two", val: 2},
	{literal: "three", val: 3},
	{literal: "four", val: 4},
	{literal: "five", val: 5},
	{literal: "six", val: 6},
	{literal: "seven", val: 7},
	{literal: "eight", val: 8},
	{literal: "nine", val: 9},
	{literal: "0", val: 0},
	{literal: "1", val: 1},
	{literal: "2", val: 2},
	{literal: "3", val: 3},
	{literal: "4", val: 4},
	{literal: "5", val: 5},
	{literal: "6", val: 6},
	{literal: "7", val: 7},
	{literal: "8", val: 8},
	{literal: "9", val: 9},
}

func main() {
	if len(os.Args) < 2 {
		panic("need an input file")
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		// num := procLine1(line)
		num := procLine2(line)
		fmt.Println(line, num)
		total += num
	}

	fmt.Println(total)
}

func procLine1(line string) int {
	first, last := 0, 0
	for _, ch := range line {
		if unicode.IsDigit(ch) {
			if first == 0 {
				first = int(ch - '0')
			}
			last = int(ch - '0')
		}
	}

	return first*10 + last
}

func procLine2(line string) int {
	first, last := 0, 0

	// find first
	index := 0
out1:
	for index < len(line) {
		for _, d := range digits {
			if index+len(d.literal) <= len(line) && line[index:index+len(d.literal)] == d.literal {
				first = d.val
				break out1
			}
		}
		index++
	}

	index = len(line)
out2:
	for index >= 0 {
		for _, d := range digits {
			if index-len(d.literal) >= 0 && line[index-len(d.literal):index] == d.literal {
				last = d.val
				break out2
			}
		}
		index--
	}

	return first*10 + last
}
