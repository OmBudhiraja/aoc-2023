package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/ombudhiraja/aoc-2023/utils"
)

const (
	highCard     = "highCard"
	onePair      = "onePair"
	twoPair      = "twoPair"
	threeOfAKind = "threeOfAKind"
	fullHouse    = "fullHouse"
	fourOfAKind  = "fourOfAKind"
	fiveOfAKind  = "fiveOfAKind"
)

type Hand struct {
	typeStrength       int
	cards              string
	bid                int
	individualStrength []int
}

func main() {

	lines := utils.Setup()

	individualCardsStrengthPart1 := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
	individualCardsStrengthPart2 := []string{"J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}

	handTypeStrength := map[string]int{
		highCard:     1,
		onePair:      2,
		twoPair:      3,
		threeOfAKind: 4,
		fullHouse:    5,
		fourOfAKind:  6,
		fiveOfAKind:  7,
	}

	sortedListPart1 := []Hand{}
	sortedListPart2 := []Hand{}

	for _, line := range lines {
		info := strings.Split(line, " ")
		hand := info[0]
		bid, _ := strconv.Atoi(info[1])

		weightPart1 := make([]int, len(hand))
		weightPart2 := make([]int, len(hand))

		for i, card := range hand {
			cardStr := string(card)

			weightPart1[i] = indexOf(cardStr, individualCardsStrengthPart1)
			weightPart2[i] = indexOf(cardStr, individualCardsStrengthPart2)
		}

		sortedListPart1 = append(sortedListPart1, Hand{
			typeStrength:       getHandStrength(hand, handTypeStrength),
			cards:              hand,
			bid:                bid,
			individualStrength: weightPart1,
		})

		sortedListPart2 = append(sortedListPart2, Hand{
			typeStrength:       getHandStrengthWithJokerRules(hand, handTypeStrength),
			cards:              hand,
			bid:                bid,
			individualStrength: weightPart2,
		})

	}

	sortCards(sortedListPart1)
	sortCards(sortedListPart2)

	resultPart1 := utils.Reduce(sortedListPart1, func(acc int, hand Hand, idx int) int {
		return acc + (hand.bid * (idx + 1))
	}, 0)

	resultPart2 := utils.Reduce(sortedListPart2, func(acc int, hand Hand, idx int) int {
		return acc + (hand.bid * (idx + 1))
	}, 0)

	fmt.Println("Part 1 ->", resultPart1)
	fmt.Println("Part 2 ->", resultPart2)
}

func sortCards(list []Hand) {
	sort.Slice(list, func(i, j int) bool {

		if list[i].typeStrength == list[j].typeStrength {

			for k := 0; k < len(list[i].individualStrength); k++ {
				if list[i].individualStrength[k] == list[j].individualStrength[k] {
					continue
				}

				return list[i].individualStrength[k] < list[j].individualStrength[k]
			}

		}

		return list[i].typeStrength < list[j].typeStrength
	})
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}

	panic("Element not found")
}

func getHandStrength(str string, strengthMap map[string]int) int {

	set := map[string]int{}

	for _, card := range str {
		set[string(card)]++
	}

	if len(set) == 1 {
		return strengthMap[fiveOfAKind]
	}

	if len(set) == 2 {
		for _, v := range set {
			if v == 4 || v == 1 {
				return strengthMap[fourOfAKind]
			}

			if v == 3 || v == 2 {
				return strengthMap[fullHouse]
			}
		}
	}

	if len(set) == 3 {
		for _, v := range set {
			if v == 3 {
				return strengthMap[threeOfAKind]
			}

			if v == 2 {
				return strengthMap[twoPair]
			}
		}
	}

	if len(set) == 4 {
		return strengthMap[onePair]
	}

	return strengthMap[highCard]

}

func getHandStrengthWithJokerRules(str string, strengthMap map[string]int) int {

	set := map[string]int{}

	for _, card := range str {
		set[string(card)]++
	}

	if _, ok := set["J"]; !ok {
		return getHandStrength(str, strengthMap)
	}

	if len(set) == 1 || len(set) == 2 {
		return strengthMap[fiveOfAKind]
	}

	if len(set) == 3 {
		if set["J"] == 1 {
			for _, v := range set {
				if v == 3 {
					return strengthMap[fourOfAKind]
				}
			}
			return strengthMap[fullHouse]
		} else {
			return strengthMap[fourOfAKind]
		}

	}

	if len(set) == 4 {
		return strengthMap[threeOfAKind]
	}

	return strengthMap[onePair]

}
