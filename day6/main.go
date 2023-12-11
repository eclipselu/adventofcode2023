package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	scanner.Scan()
	line1 := scanner.Text()
	times := parseLine(line1)
	scanner.Scan()
	line2 := scanner.Text()
	dists := parseLine(line2)

	if len(times) != len(dists) {
		panic("times and dists should have the same length")
	}

	t, d := parseLine2(line1), parseLine2(line2)
	fmt.Println("part1:", calc(times, dists))
	fmt.Println("part2:", numOfWays(t, d))
}

func calc(times, dists []int64) int64 {
	ans := int64(1)
	for i := 0; i < len(times); i++ {
		ways := numOfWays(times[i], dists[i])
		fmt.Println("ways:", ways)
		ans = ans * ways
	}
	return ans
}

func numOfWays(t, d int64) int64 {
	count := int64(0)
	for i := int64(0); i <= t; i++ {
		if i*(t-i) > d {
			count++
		}
	}
	return count
}

func parseLine(line string) []int64 {
	var nums []int64
	for _, x := range strings.Fields(line) {
		if num, err := strconv.Atoi(x); err == nil {
			nums = append(nums, int64(num))
		}
	}

	return nums
}

func parseLine2(line string) int64 {
	splits := strings.Split(line, ":")
	numStr := strings.Join(strings.Fields(splits[1]), "")
	num, _ := strconv.ParseInt(numStr, 10, 64)
	return num
}
