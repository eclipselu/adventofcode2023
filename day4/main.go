package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type scratchCard struct {
	idx         int
	winningNums map[int]bool
	nums        []int
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

	cards := make([]scratchCard, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		card := parseLine(line)
		cards = append(cards, card)
	}

	// part 1
	total := 0
	for _, card := range cards {
		total += calcScore(card)
	}
	fmt.Println("part 1:", total)

	// part 2
	total = 0
	cardCount := make([]int, len(cards))
	for i := 0; i < len(cardCount); i++ {
		cardCount[i] = 1
	}
	for i := 0; i < len(cardCount); i++ {
		total += cardCount[i]
		matches := calcMatches(cards[i])
		for idx := i + 1; idx <= min(len(cardCount)-1, i+matches); idx++ {
			cardCount[idx] += cardCount[i]
		}
	}

	fmt.Println("part 2:", total)
}

func parseLine(line string) scratchCard {
	splits := strings.Split(line, ":")
	var cid int
	fmt.Sscanf(splits[0], "Card %d", &cid)

	numsSplits := strings.Split(splits[1], "|")

	splitNums := func(numsStr string) []int {
		scanner := bufio.NewScanner(strings.NewReader(numsStr))
		scanner.Split(bufio.ScanWords)
		nums := make([]int, 0)
		for scanner.Scan() {
			num, _ := strconv.Atoi(scanner.Text())
			nums = append(nums, num)
		}

		return nums
	}

	winNums := splitNums(numsSplits[0])
	card := scratchCard{
		idx:         cid,
		winningNums: make(map[int]bool),
		nums:        splitNums(numsSplits[1]),
	}

	for _, x := range winNums {
		card.winningNums[x] = true
	}
	return card
}

func calcMatches(card scratchCard) int {
	count := 0
	for _, num := range card.nums {
		if card.winningNums[num] {
			count++
		}
	}
	return count
}

func calcScore(card scratchCard) int {
	score := 0
	count := calcMatches(card)

	for i := 0; i < count; i++ {
		if score == 0 {
			score = 1
		} else {
			score *= 2
		}
	}
	return score
}
