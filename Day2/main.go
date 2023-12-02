package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"strings"
)

var maximumValue = make(map[string]int)
var amountAndColourRegex = regexp.MustCompile(`[0-9]+ [a-z]+\s*`)
var amountRegex = regexp.MustCompile(`[0-9]+\s*`)
var colourRegex = regexp.MustCompile(`[a-z]+\s*`)

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

func main() {
	lines := lines("input.txt")
	maximumValue["red"] = 12
	maximumValue["green"] = 13
	maximumValue["blue"] = 14

	fmt.Println("Part One: ", partOne(lines))
	fmt.Println("Part Two: ", partTwo(lines))
}

func partOne(lines []string) int {
	output := 0
	for _, line := range lines {
		gameId, _ := strconv.Atoi(strings.Split(strings.Split(line, ":")[0], " ")[1])
		if score(line) {
			output = output + gameId
		}
	}
	return output
}

func partTwo(lines []string) int {
	output := 0
	for _, line := range lines {
		output = output + power(line)
	}
	return output
}

func power(line string) int {
	blue := 0
	green := 0
	red := 0
	for _, cs := range amountAndColourRegex.FindAllString(line, -1){
		amount, _ := strconv.Atoi(strings.Trim(amountRegex.FindString(cs)," "))
		switch colour := colourRegex.FindString(cs); colour {
		case "green":
			if green < amount{
				green = amount
			}
		case "red":
			if red < amount{
				red = amount
			}
		case "blue":
			if blue < amount{
				blue = amount
			}
		}
	} 
	return blue * red * green
}

func score(line string) bool {
	for _, amountcolour := range amountAndColourRegex.FindAllString(line, -1){
		amount, _ := strconv.Atoi(strings.Trim(amountRegex.FindString(amountcolour)," "))
		if maximumValue[colourRegex.FindString(amountcolour)]  < amount{
			return false
		}
	}
	return true
}

