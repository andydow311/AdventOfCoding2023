package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var spacesRegex = regexp.MustCompile(`\s+`)

func main() {
	lines := lines("input.txt")
	fmt.Println("Part One: ", partOne(lines))
	fmt.Println("Part Two: ", partTwo(lines))
}

func partOne(lines []string) int {
	times := getInts(strings.Trim(strings.ReplaceAll(spacesRegex.ReplaceAllString(lines[0], " "), "Time:", ""), " "))
	distances := getInts(strings.Trim(strings.ReplaceAll(spacesRegex.ReplaceAllString(lines[1], " "), "Distance:", ""), " "))
	waysOfWinning := 1
	for i, time := range times {
		thisWayofWinning := 0
		distance := distances[i]
		for k := 0; k <= time; k++ {
			if k*(time-k) > distance {
				thisWayofWinning++
			}
		}
		waysOfWinning = waysOfWinning * thisWayofWinning
	}
	return waysOfWinning
}

func partTwo(lines []string) int {
	time, distance := convertToSingleInteger(lines)
	waysofWinning := 0
	for k := 0; k <= time; k++ {
		if k*(time-k) > distance {
			waysofWinning++
		}
	}
	return waysofWinning
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

func getInts(line string) []int {
	ints := []int{}
	for _, s := range strings.Split(line, " ") {
		i, _ := strconv.Atoi(s)
		ints = append(ints, i)
	}
	return ints
}

func convertToSingleInteger(lines []string) (int, int) {
	time, _ := strconv.Atoi(spacesRegex.ReplaceAllString(strings.Trim(strings.ReplaceAll(spacesRegex.ReplaceAllString(lines[0], " "), "Time:", ""), " "), ""))
	distance, _ := strconv.Atoi(spacesRegex.ReplaceAllString(strings.Trim(strings.ReplaceAll(spacesRegex.ReplaceAllString(lines[1], " "), "Distance:", ""), " "), ""))
	return time, distance
}
