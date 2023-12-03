package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"strings"
)

var digitRegex = regexp.MustCompile(`[0-9]+`)
var periodRegex = regexp.MustCompile(`"\.`)
var gearRegex = regexp.MustCompile(`\*`)

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
	output := 0
	for lineIndex, line := range lines {
		theseLines := getLinesAboveAndBelow(lines, lineIndex, line)
		digits := digitRegex.FindAllString(line, -1)
		digitsIndex := digitRegex.FindAllStringIndex(line, -1)
		for index := range digitsIndex {
			startIndex := maximum(0, digitsIndex[index][0]-1)
			endIndex := minimum(digitsIndex[index][1]+1, len(line)-1)
			if isPart(theseLines, startIndex, endIndex) {
				value, _ := strconv.Atoi(digits[index])
				output = output + value
			}
		}

	}
	return output
}

func partTwo(lines []string) int {
	output := 0
	mapSymbolToPortNumber := make(map[string][]string)
	
	for lineIndex, line := range lines {
		theseLines := getLinesAboveAndBelow(lines, lineIndex, line)
		digits := digitRegex.FindAllString(line, -1)
		digitsIndex := digitRegex.FindAllStringIndex(line, -1)
		for index := range digitsIndex {
			startIndex := maximum(0, digitsIndex[index][0]-1)
			endIndex := minimum(digitsIndex[index][1]+1, len(line)-1)
			digit := digits[index]
			if isPart(theseLines, startIndex, endIndex){
				keyOk,line := partAdjacentToAsterix(theseLines, startIndex, endIndex)
				if keyOk {
					_, ok := mapSymbolToPortNumber[line]
					if ok {
						array := mapSymbolToPortNumber[line]
						array = append(array, digit)
						mapSymbolToPortNumber[line] = array
					} else {
						array := []string{digit}
						mapSymbolToPortNumber[line] = array
					}
				}					
			}
		}
	}

	for key, value := range mapSymbolToPortNumber {
		if len(value) == 2 {
			firstRatio, _ := strconv.Atoi(mapSymbolToPortNumber[key][0])
			secondRatio, _ := strconv.Atoi(mapSymbolToPortNumber[key][1])
			output=output+(firstRatio*secondRatio)
		}
	}
	return output
}

func maximum(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func minimum(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func isPart(theseLines []string, startIndex int, endIndex int) bool {
	for _, line := range theseLines {
		for i := startIndex; i < endIndex; i++ {
			if string(line[i]) != "." && !digitRegex.MatchString(string(line[i])) {
				return true
			}
		}
	}
	return false
}

func partAdjacentToAsterix(theseLines []string, startIndex int, endIndex int) (bool, string) {
	for _, line := range theseLines {
		for i := startIndex; i < endIndex; i++ {
			if gearRegex.MatchString(string(line[i])) {
				line = line+string(rune(i))
				return true, line
			}
		}
	}
	return false, ""
}

func getLinesAboveAndBelow(lines []string, index int, line string) []string {
	lineBelow := ""
	line = strings.Trim(line, " ")
	lineAbove := ""

	if index == 0 {
		lineAbove = line
		lineBelow = lines[index+1]

	} else if index == len(lines)-1 {
		lineAbove = lines[index-1]
		lineBelow = line
	} else {
		lineAbove = lines[index-1]
		lineBelow = lines[index+1]
	}

	return []string{lineAbove, line, lineBelow}
}
