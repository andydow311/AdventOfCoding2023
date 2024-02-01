package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := lines("input.txt")
		for part:= 1; part <=2; part++{
			output:=0
			for _, line := range lines {
				ints := getInts(line)
				levels := getLevels(ints)
				keys := sortKeys(levels)
				if part == 1{
					output = output + getNextValuePartOne(keys, levels)
				}else if part ==2{
					output = output + getNextValuePartTwo(keys, levels)
				}
			}
			if part == 1{
				fmt.Println("Part One: ", output)
			}else{
				fmt.Println("Part Two: ",output)
			}
		}
}

func sortKeys(levels map[int][]int) []int {
	keys := make([]int, 0)
	for k := range levels {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	return keys
}

func getNextValuePartTwo(keys []int, levels map[int][]int) int {
	value:=0
	for _, key := range keys{
		values:= levels[key]
		value = values[0]-value 
	}
	return value
}

func getNextValuePartOne(keys []int, levels map[int][]int) int {
	value:=0
	for _, key := range keys{
		values:= levels[key]
		value = value +  values[len(values)-1] 
	}
	return value
}

func getLevels(ints []int) map[int][]int {
	levels := make(map[int][]int)
	nextLevel := []int{}
	levels[0] = ints
	level := 1
	for {
		for i := 0; i < len(ints)-1; i++ {
			nextLevel = append(nextLevel, ints[i+1]-ints[i])
		}
		levels[level] = nextLevel
		if allZeros(nextLevel) {
			return levels
		} else {
			ints = nextLevel
			nextLevel = []int{}
		}
		level++
	}
}

func allZeros(ints []int) bool {
	for _, int := range ints {
		if int != 0 {
			return false
		}
	}
	return true
}

func getInts(line string) []int {
	ints := []int{}
	for _, s := range strings.Split(line, " ") {
		i, _ := strconv.Atoi(s)
		ints = append(ints, i)
	}
	return ints
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

