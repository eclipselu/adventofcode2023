package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type rangeMapLine struct {
	dstStart int
	srcStart int
	len      int
}

type rangeMap struct {
	srcType string
	dstType string
	ranges  []rangeMapLine
}

type interval struct {
	start int
	end   int
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

	graph := make(map[string]*rangeMap)
	// get seeds
	var seeds []int
	scanner.Scan()
	seedLine := strings.FieldsFunc(scanner.Text(), func(r rune) bool { return r == ':' })
	for _, x := range strings.Fields(seedLine[1]) {
		num, _ := strconv.Atoi(x)
		seeds = append(seeds, num)
	}
	scanner.Scan()

	// build graph
	srcType := "seed"
	for {
		r := readRangeMap(scanner)
		if r == nil {
			break
		}

		graph[srcType] = r
		srcType = r.dstType
	}

	fmt.Println("part 1:", part1(seeds, graph))
	fmt.Println("part 2:", part2(seeds, graph))
}

func part1(seeds []int, graph map[string]*rangeMap) int {
	var locs []int
	for _, seed := range seeds {
		loc := getLocation(seed, graph)
		locs = append(locs, loc)
	}
	return slices.Min(locs)
}

// func part2Slow(seeds []int, graph map[string]*rangeMap) int {
// 	var sds []int
// 	for i := 0; i < len(seeds); i += 2 {
// 		for j := seeds[i]; j < seeds[i]+seeds[i+1]; j++ {
// 			sds = append(sds, j)
// 		}
// 	}

// 	return part1(sds, graph)
// }

func part2(seeds []int, graph map[string]*rangeMap) int {
	var itvls []interval
	for i := 0; i < len(seeds); i += 2 {
		itvls = append(itvls, interval{start: seeds[i], end: seeds[i] + seeds[i+1] - 1})
	}

	srcType := "seed"
	for srcType != "location" {
		m := graph[srcType]
		itvls = mapInterval(itvls, m)
		// fmt.Println(srcType, "->", m.dstType, itvls)
		srcType = m.dstType
	}

	return slices.MinFunc(itvls, func(a, b interval) int {
		return a.start - b.start
	}).start
}

func getTargetVal(srcVal int, m *rangeMap) int {
	for _, rg := range m.ranges {
		if srcVal >= rg.srcStart && srcVal < rg.srcStart+rg.len {
			return rg.dstStart + (srcVal - rg.srcStart)
		}
	}

	return srcVal
}

// mapIntervals
func mapInterval(itvls []interval, m *rangeMap) []interval {
	var result []interval

	remaining := make([]interval, len(itvls))
	copy(remaining, itvls)

	for _, rg := range m.ranges {
		var nextRemaining []interval

		for _, it := range remaining {
			// check if there's an intersection
			if overlap := intersect(it, interval{start: rg.srcStart, end: rg.srcStart + rg.len - 1}); overlap != nil {
				// map the intersected part
				offset := overlap.start - rg.srcStart
				start := rg.dstStart + offset
				end := rg.dstStart + offset + (overlap.end - overlap.start)
				result = append(result, interval{start: start, end: end})

				// match the other parts with other map ranges
				if overlap.start > it.start {
					nextRemaining = append(nextRemaining, interval{start: it.start, end: overlap.start - 1})
				}
				if overlap.end < it.end {
					nextRemaining = append(nextRemaining, interval{start: overlap.end + 1, end: it.end})
				}
			} else {
				nextRemaining = append(nextRemaining, it)
			}
		}

		remaining = nextRemaining
	}

	result = append(result, remaining...)

	return result
}

func intersect(i1, i2 interval) *interval {
	if i1.start > i2.start {
		return intersect(i2, i1)
	}

	if i1.end >= i2.start {
		return &interval{start: i2.start, end: min(i1.end, i2.end)}
	}

	return nil
}

func getLocation(seed int, g map[string]*rangeMap) int {
	srcType := "seed"
	val := seed
	for srcType != "location" {
		m := g[srcType]
		val = getTargetVal(val, m)
		srcType = m.dstType
	}

	return val
}

func readRangeMap(scanner *bufio.Scanner) *rangeMap {
	var header string
	var types string
	if ok := scanner.Scan(); !ok {
		return nil
	}

	header = scanner.Text()
	fmt.Sscanf(header, "%s map:", &types)
	typeList := strings.FieldsFunc(types, func(r rune) bool { return r == '-' })

	r := &rangeMap{
		srcType: typeList[0],
		dstType: typeList[2],
	}

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		// make rangeMapLine
		rLine := rangeMapLine{}
		fmt.Sscanf(line, "%d %d %d", &rLine.dstStart, &rLine.srcStart, &rLine.len)
		r.ranges = append(r.ranges, rLine)
	}

	return r
}
