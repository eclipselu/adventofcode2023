package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type gameInfo struct {
	id       uint
	maxCount map[string]uint
}

func parseLine(line string) gameInfo {
	splits := strings.Split(line, ":")
	var gid uint
	fmt.Sscanf(splits[0], "Game %d", &gid)

	game := gameInfo{id: gid, maxCount: make(map[string]uint)}
	for _, set := range strings.Split(splits[1], ";") {
		for _, split := range strings.Split(set, ",") {
			var cnt uint
			var color string
			fmt.Sscanf(split, "%d %s", &cnt, &color)

			if cnt > game.maxCount[color] {
				game.maxCount[color] = cnt
			}
		}
	}

	return game
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
	var games []gameInfo
	for scanner.Scan() {
		line := scanner.Text()
		game := parseLine(line)
		games = append(games, game)
	}

	// part 1
	sum1 := 0
	for _, game := range games {
		if game.maxCount["red"] <= 12 && game.maxCount["green"] <= 13 && game.maxCount["blue"] <= 14 {
			sum1 += int(game.id)
		}
	}
	fmt.Println(sum1)

	// part 2
	sum2 := 0
	for _, game := range games {
		sum2 += int(game.maxCount["red"]) * int(game.maxCount["green"]) * int(game.maxCount["blue"])
	}
	fmt.Println(sum2)
}
