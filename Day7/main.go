package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"strconv"
	"strings"
)

var cardStrength = make(map[string]int)

func populateCardStrength() {
	cardStrength["A"] = 13
	cardStrength["K"] = 12
	cardStrength["Q"] = 11
	cardStrength["J"] = 10
	cardStrength["T"] = 9
	cardStrength["9"] = 8
	cardStrength["8"] = 7
	cardStrength["7"] = 6
	cardStrength["6"] = 5
	cardStrength["5"] = 4
	cardStrength["4"] = 3
	cardStrength["3"] = 2
	cardStrength["2"] = 1
}

type Hand struct {
	cards      []string
	score      int
	rank       int
	stake      int
}

func main() {
	populateCardStrength()
	lines := lines("input.txt")
	hands := getHands(lines)
	fmt.Println("Part One:",partOne(hands))
	fmt.Println("Part Two:",partTwo(hands))
	
}

func partOne(hands []Hand) int {
	for index := range hands {
		hand := &hands[index]
		hand.score = scoreCards(hand.cards)
	}
	rankHands(hands)
	return getTotalWinnings(hands)
}

func partTwo(hands []Hand) int {
	cardStrength["J"] = 0
	for index := range hands {
		hand := &hands[index]
		hand.score = ScoreCardsWithJoker(hand.cards)
	}
	rankHands(hands)
	return getTotalWinnings(hands)
}

func rankHands(hands []Hand) {
	sort.Slice(hands[:], func(i, j int) bool {
		if hands[i].score == hands[j].score {
			for x:=0; x<len(hands[i].cards);x++{
				if cardStrength[hands[i].cards[x]] == cardStrength[hands[j].cards[x]]{
					continue
				}else{
					return cardStrength[hands[i].cards[x]] > cardStrength[hands[j].cards[x]]
				}
			}
		}		
		return hands[i].score > hands[j].score
	})
	for i := range hands {
		hand := &hands[i]
		hand.rank = len(hands) - i
	}
}

func getHands(lines []string) []Hand {
	output := []Hand{}
	for _, line := range lines {
		hand := strings.Split(line, " ")[0]
		stake, _ := strconv.Atoi(strings.Trim(strings.Split(line, " ")[1], " "))
		cards := strings.Split(hand, "")
		output = append(output, Hand{cards: cards, rank: 0, stake: stake})
	}
	return output
}

func ScoreCardsWithJoker(cards []string) int{
	cardsScore := make(map[string]int)
	for _, card := range cards {
		_, ok := cardsScore[card]
		if !ok {
			cardsScore[card] = 1
		} else {
			cardsScore[card]++
		}
	}	
	_, ok := cardsScore["J"]
	if ok{
		updateCards(cardsScore)
	}
	return classifyHand(cardsScore)
}

func updateCards(cardsScore map[string]int){
	winningValue := 0
	winningCards := []string{}
	for cards, value := range cardsScore {
		if cards != "J"{
			if value > winningValue {
				winningValue = value
				winningCards = []string{}
				winningCards = append(winningCards, cards)
			}else if value == winningValue{
				winningCards = append(winningCards, cards)
			}
		}
	}
	bestCard:=""
	bcValue := 0
	for _, card :=range winningCards{
		if cardStrength[card] > bcValue{
			bestCard = card
			bcValue = cardStrength[card]
		}
	}
	if len(winningCards) == 0{
		cardsScore["A"] = 5
		delete(cardsScore, "J");
	}else{
		cardsScore[bestCard] = winningValue + cardsScore["J"]
		delete(cardsScore, "J");
	}
}

func scoreCards(cards []string) int {
	cardsScore := make(map[string]int)
	for _, card := range cards {
		_, ok := cardsScore[card]
		if !ok {
			cardsScore[card] = 1
		} else {
			cardsScore[card]++
		}
	}
	return classifyHand(cardsScore)
}

func classifyHand(cardsScore map[string]int) int{
	fives := 0
	fours := 0
	threes := 0
	pairs := 0

	for _, value := range cardsScore {
		if value == 5 {
			fives++
		} else if value == 4 {
			fours++
		} else if value == 3 {
			threes++
		} else if value == 2 {
			pairs++
		}
	}
	if fives == 1 {
		return 8
	} else if fours == 1 {
		return 7
	} else if pairs == 1 && threes == 1 {
		return 6
	} else if threes == 1 {
		return 5
	} else if pairs == 2 {
		return 4
	} else if pairs == 1 {
		return 3
	} else {
		return 2
	}
}

func lines(filename string) []string {
	output := []string{}
	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		output = append(output, fileScanner.Text())
	}
	readFile.Close()
	return output
}

func getTotalWinnings(hands []Hand) int {
	output := 0
	for _, h := range hands {
		output = output + (h.rank * h.stake)
	}
	return output
}