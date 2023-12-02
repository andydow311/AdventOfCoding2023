package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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
	fmt.Println("Part One: ", partOne(lines))
	fmt.Println("Part Two: ", partTwo(lines))
}

func partOne(lines []string) int {
	nonNumericRegex := regexp.MustCompile(`[^0-9 ]+`)
	calibrationValue := 0
	for _, line := range lines {
		digits := nonNumericRegex.ReplaceAllString(line, "")
		if len(digits) == 1 {
			digit, _ := strconv.Atoi(string(digits) + string(digits))
			calibrationValue = calibrationValue + digit
		} else {
			digit, _ := strconv.Atoi(digits[0:1] + digits[len(digits)-1:])
			calibrationValue = calibrationValue + digit
		}
	}
	return calibrationValue
}

func partTwo(lines []string) int {
	cleanedLines := []string{}
	for _, line := range lines {
		tempString := ""
		for _, character := range line {
			tempString = tempString + string(character)
			tempString = cleanString(tempString)
		}
		cleanedLines = append(cleanedLines, tempString)
	}
	return partOne(cleanedLines)
}

func cleanString(line string) string {
	line = strings.ReplaceAll(line, "one", "1")
	line = strings.ReplaceAll(line, "two", "2")
	line = strings.ReplaceAll(line, "three", "3")
	line = strings.ReplaceAll(line, "four", "4")
	line = strings.ReplaceAll(line, "five", "5")
	line = strings.ReplaceAll(line, "six", "6")
	line = strings.ReplaceAll(line, "seven", "7")
	line = strings.ReplaceAll(line, "eight", "8")
	line = strings.ReplaceAll(line, "nine", "9")
	return line
}
