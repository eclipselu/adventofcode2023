package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type partNum struct {
	idx   int
	start int
	end   int
	val   int
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

	m := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		m = append(m, line)
	}

	nums := make([]partNum, 0)
	gears := make(map[string][]partNum)
	for row, line := range m {
		for col := 0; col < len(line); {
			if unicode.IsDigit(rune(line[col])) {
				start := col
				end := col
				val := 0

				for end < len(line) && unicode.IsDigit(rune(line[end])) {
					val = val*10 + int(line[end]-'0')
					end++
				}

				num := partNum{idx: row, start: start, end: end, val: val}
				if isValidPartNum(m, num, gears) {
					nums = append(nums, num)
				}
				col = end
			} else {
				col++
			}
		}
	}

	// part 1
	total := 0
	for _, num := range nums {
		total += num.val
	}
	fmt.Println("part 1:", total)

	// part 2
	sumOfRatios := 0
	for _, gear := range gears {
		if len(gear) != 2 {
			continue
		}
		sumOfRatios += gear[0].val * gear[1].val
	}
	fmt.Println("part 2:", sumOfRatios)
}

func isValidPartNum(m []string, num partNum, gears map[string][]partNum) bool {
	valid := false
	for row := max(0, num.idx-1); row <= min(len(m)-1, num.idx+1); row++ {
		for col := max(0, num.start-1); col <= min(len(m[row])-1, num.end); col++ {
			if row == num.idx && col >= num.start && col < num.end {
				continue
			}
			ch := rune(m[row][col])
			if !unicode.IsDigit(ch) && !unicode.IsLetter(ch) && ch != '.' {
				valid = true
				if ch == '*' {
					key := coord2Str(row, col)
					gears[key] = append(gears[key], num)
				}
			}
		}
	}

	return valid
}

func coord2Str(r, c int) string {
	return fmt.Sprintf("%d_%d", r, c)
}
