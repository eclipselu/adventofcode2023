package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	HIGH_CARD = iota
	ONE_PAIR
	TWO_PAIR
	THREE
	FULL_HOUSE
	FOUR
	FIVE
)

type handBid struct {
	hand     string
	bid      int
	handType int
}

func calcCardStrength(c byte) int {
	return strings.IndexByte("23456789TJQKA", c)
}

func calcCardStrength2(c byte) int {
	return strings.IndexByte("J23456789TQKA", c)
}

func calcHandType(hand string) int {
	counts := make(map[rune]int)
	for _, r := range hand {
		counts[r]++
	}

	freq := make(map[int]int)
	for _, c := range counts {
		freq[c]++
	}

	if freq[5] == 1 {
		return FIVE
	} else if freq[4] == 1 && freq[1] == 1 {
		return FOUR
	} else if freq[3] == 1 && freq[2] == 1 {
		return FULL_HOUSE
	} else if freq[3] == 1 && freq[1] == 2 {
		return THREE
	} else if freq[2] == 2 && freq[1] == 1 {
		return TWO_PAIR
	} else if freq[2] == 1 && freq[1] == 3 {
		return ONE_PAIR
	} else if freq[1] == 5 {
		return HIGH_CARD
	}

	return -1
}

func calcHandType2(hand string) int {
	counts := make(map[rune]int)
	jCount := 0
	for _, r := range hand {
		if r == 'J' {
			jCount++
		} else {
			counts[r]++
		}
	}

	maxKey := hand[0]
	maxCnt := 0
	for k, v := range counts {
		if v > maxCnt {
			maxKey = byte(k)
			maxCnt = v
		}
	}
	counts[rune(maxKey)] += jCount

	freq := make(map[int]int)
	for _, c := range counts {
		freq[c]++
	}

	if freq[5] == 1 {
		return FIVE
	} else if freq[4] == 1 && freq[1] == 1 {
		return FOUR
	} else if freq[3] == 1 && freq[2] == 1 {
		return FULL_HOUSE
	} else if freq[3] == 1 && freq[1] == 2 {
		return THREE
	} else if freq[2] == 2 && freq[1] == 1 {
		return TWO_PAIR
	} else if freq[2] == 1 && freq[1] == 3 {
		return ONE_PAIR
	} else if freq[1] == 5 {
		return HIGH_CARD
	}

	return -1
}

func parseLine(line string) handBid {
	fields := strings.Fields(line)
	hand := fields[0]
	bid, _ := strconv.Atoi(fields[1])

	return handBid{
		hand: hand,
		bid:  bid,
	}
}

func calcWinnings(hbs []handBid, handTypeFunc func(string) int, cardStrengthFunc func(byte) int) int {
	for i := 0; i < len(hbs); i++ {
		hbs[i].handType = handTypeFunc(hbs[i].hand)
	}

	slices.SortFunc(hbs, func(a, b handBid) int {
		if a.handType != b.handType {
			return a.handType - b.handType
		} else {
			for i := 0; i < len(a.hand); i++ {
				strengthA, strengthB := cardStrengthFunc(a.hand[i]), cardStrengthFunc(b.hand[i])
				if strengthA != strengthB {
					return strengthA - strengthB
				}
			}
			return 0
		}
	})

	total := 0
	for i, hb := range hbs {
		total += (i + 1) * hb.bid
	}

	return total
}

func part1(hbs []handBid) int {
	return calcWinnings(hbs, calcHandType, calcCardStrength)
}

func part2(hbs []handBid) int {
	return calcWinnings(hbs, calcHandType2, calcCardStrength2)
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
	hbs := make([]handBid, 0)

	for scanner.Scan() {
		line := scanner.Text()
		hbs = append(hbs, parseLine(line))
	}

	fmt.Println("part 1:", part1(hbs))
	fmt.Println("part 2:", part2(hbs))
}
