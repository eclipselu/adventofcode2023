package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func parseLine(line string) (string, []string) {
	fields := strings.FieldsFunc(line, func(r rune) bool { return !unicode.IsUpper(r) && !unicode.IsDigit(r) })
	return fields[0], fields[1:]
}

func part1(inst string, g map[string][]string) int {
	steps := 0
	node := "AAA"
	target := "ZZZ"

	for i := 0; ; i = (i + 1) % len(inst) {
		if node == target {
			break
		}

		idx := 0
		if inst[i] == 'R' {
			idx = 1
		}

		node = g[node][idx]
		steps++
	}

	return steps
}

func lcm(nums []int) int64 {
	if len(nums) == 1 {
		return int64(nums[0])
	}

	d := gcd(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		d = gcd(d, nums[i])
	}

	ans := int64(d)
	for _, num := range nums {
		ans *= int64(num / d)
	}
	return ans
}

func gcd(a, b int) int {
	for b != 0 {
		c := b
		b = a % b
		a = c
	}
	return a
}

func part2Single(inst string, g map[string][]string, node string) int {
	steps := 0
	for i := 0; ; i = (i + 1) % len(inst) {
		if node[2] == 'Z' {
			break
		}

		idx := 0
		if inst[i] == 'R' {
			idx = 1
		}

		node = g[node][idx]
		steps++
	}

	return steps
}

func part2(inst string, g map[string][]string) int64 {
	nodes := make([]string, 0)
	for k := range g {
		if k[2] == 'A' {
			nodes = append(nodes, k)
		}
	}

	stepList := make([]int, 0)
	for _, node := range nodes {
		stepList = append(stepList, part2Single(inst, g, node))
	}

	return lcm(stepList)
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

	graph := make(map[string][]string)
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	inst := scanner.Text()
	scanner.Scan()

	for scanner.Scan() {
		node, dst := parseLine(scanner.Text())
		graph[node] = dst
	}

	// fmt.Println("part1:", part1(inst, graph))
	fmt.Println("part2:", part2(inst, graph))
}
