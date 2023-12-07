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

func main() {

	lines := utils.Setup()

	individualCardsStrength := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}

	handTypeStrength := map[string]int{
		highCard:     100,
		onePair:      200,
		twoPair:      300,
		threeOfAKind: 400,
		fullHouse:    500,
		fourOfAKind:  600,
		fiveOfAKind:  700,
	}

	type Hand struct {
		typeStrength       int
		cards              string
		bid                int
		individualStrength []int
	}

	sortedList := []Hand{}

	for _, line := range lines {
		info := strings.Split(line, " ")
		hand := info[0]
		bid, _ := strconv.Atoi(info[1])

		weight := make([]int, len(hand))

		for i, card := range hand {
			cardStr := string(card)
			cardIndex := indexOf(cardStr, individualCardsStrength)
			weight[i] = cardIndex * (len(hand) - i)
		}

		sortedList = append(sortedList, Hand{
			typeStrength:       getHandStrength(hand, handTypeStrength),
			cards:              hand,
			bid:                bid,
			individualStrength: weight,
		})

	}

	sort.Slice(sortedList, func(i, j int) bool {

		if sortedList[i].typeStrength == sortedList[j].typeStrength {

			for k := 0; k < len(sortedList[i].individualStrength); k++ {
				if sortedList[i].individualStrength[k] == sortedList[j].individualStrength[k] {
					continue
				}

				return sortedList[i].individualStrength[k] < sortedList[j].individualStrength[k]
			}

		}

		return sortedList[i].typeStrength < sortedList[j].typeStrength
	})

	val := utils.Reduce(sortedList, func(acc int, hand Hand, idx int) int {
		return acc + (hand.bid * (idx + 1))
	}, 0)

	fmt.Println(val)
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
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
