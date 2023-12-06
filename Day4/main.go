package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"

	"strings"
)

var digitRegex = regexp.MustCompile(`[0-9]+\s*`)
var cardNumberRegex = regexp.MustCompile("[0-9]+:")

func main() {
	lines := lines("input.txt")
	part1:= 0.0
	cards := getCards(lines)
	for index, line := range lines {
		line := cardNumberRegex.ReplaceAllString(line, "") 
		winningNumbers := digitRegex.FindAllString(strings.Split(line, "|")[0], -1)
		numbers := digitRegex.FindAllString(strings.Split(line, "|")[1], -1)
		matchingNumbers := getMatchingNumbers(winningNumbers, numbers)
		part1 = part1 + partOne(matchingNumbers)
		partTwo(index, cards, matchingNumbers)
	}
	fmt.Println("Part One: ", part1)
	fmt.Println("Part Two: ", sumArray(cards))
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

func partOne(matchingNumbers[]string) float64 {
	if len(matchingNumbers) > 0 {
		return (math.Exp2(float64(len(matchingNumbers) - 1)))
	}
	return 0.0
}

func partTwo(index int, cards []int, matchingNumbers[]string){
		cards[index]++	
	    for j:=0; j< cards[index]; j++{
			for i := index+1; i <= index+len(matchingNumbers); i++{
				cards[i]++
			}	
		}
}

func getNumbers(line string) ([]string,[]string){
	line = cardNumberRegex.ReplaceAllString(line, "")
	winningNumbers := digitRegex.FindAllString(strings.Split(line, "|")[0], -1)
	numbers := digitRegex.FindAllString(strings.Split(line, "|")[1], -1)
	return winningNumbers, numbers
}

func getMatchingNumbers(winning []string, numbers []string) []string{
	matchingNumbers := []string{}
	for _, x := range winning {
		for _, y := range numbers {
			if strings.Trim(x, " ") == strings.Trim(y, " ") {
				matchingNumbers = append(matchingNumbers, x)
			}
		}
	}
	return matchingNumbers
}

func getCards(lines []string) []int{
	cards := make([]int, len(lines))
	for j:=0; j< len(lines)-1;j++{
		cards[j] = 0
	}
	return cards
}


func sumArray(ints []int) int{
	output:= 0
	for _,y := range ints{
		output = output+y
	}
	return output
}