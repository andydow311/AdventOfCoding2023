package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SpringType int

const (
	springTypeUnknown SpringType = iota
	springTypeOperational
	springTypeDamaged
)

type Entry struct {
	springTypesLength int
	numbersLength     int
	remaining         int
}

func main() {
	lines := lines("input.txt")
	fmt.Println("Part One:", partOne(lines))
	fmt.Println("Part Two:", partTwo(lines))
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
	fmt.Println("line ", line)
	for _, s := range strings.Split(line, ",") {
		i, _ := strconv.Atoi(s)
		ints = append(ints, i)
	}
	return ints
}

func partOne(lines []string) int {
	output:= 0
	for _, line := range lines {
		springTypes, numbers := parse(line)
		fmt.Println("SpringTypes ", springTypes, "numbers ", numbers)
		output += counter{cache: make(map[Entry]int)}.countArrangements(springTypes, numbers, -1)
	}
	return output
}

func parse(line string) ([]SpringType, []int) {
	var springs []SpringType
	pattern := strings.Split(strings.Split(line, " ")[0], "")
	for i := 0; i < len(pattern); i++ {
		item := pattern[i]
		switch item {
		case "?":
			springs = append(springs, springTypeUnknown)
		case ".":
			springs = append(springs, springTypeOperational)
		case "#":
			springs = append(springs, springTypeDamaged)
		default:
			panic(string(item))
		}
	}

	return springs, getInts(strings.Split(line, " ")[1])
}

type counter struct {
	cache map[Entry]int
}

func (c counter) countArrangements(springTypes []SpringType, numbers []int, remaining int) int {
	if len(springTypes) == 0 {
		if len(numbers) == 0 || (len(numbers) == 1 && remaining == 0) {
			return 1
		}
		return 0
	}

	entry := Entry{springTypesLength: len(springTypes), numbersLength: len(numbers), remaining: remaining}
	if v, exists := c.cache[entry]; exists {
		return v
	}

	springType := springTypes[0]
	res := 0
	switch springType {
	case springTypeUnknown:
		switch remaining {
		case -1:
			with := 0
			if len(numbers) != 0 {
				with = c.countArrangements(springTypes[1:], numbers, numbers[0]-1)
			}
			without := c.countArrangements(springTypes[1:], numbers, -1)
			res = with + without
		case 0:
			res = c.countArrangements(springTypes[1:], numbers[1:], -1)
		default:
			res = c.countArrangements(springTypes[1:], numbers, remaining-1)
		}
	case springTypeOperational:
		switch remaining {
		case -1:
			res = c.countArrangements(springTypes[1:], numbers, remaining)
		case 0:
			res = c.countArrangements(springTypes[1:], numbers[1:], -1)
		default:
		}
	case springTypeDamaged:
		switch remaining {
		case -1:
			if len(numbers) == 0 {
				break
			}
			res = c.countArrangements(springTypes[1:], numbers, numbers[0]-1)
		case 0:
		default:
			res = c.countArrangements(springTypes[1:], numbers, remaining-1)
		}
	default:
		panic(springType)
	}

	c.cache[entry] = res
	return res
}


func partTwo(lines []string) int {
	output := 0
	for _, line := range lines {
		springTypes, numbers := parse(line)
		springTypes, numbers = unfold(springTypes, numbers, 5)
		output += counter{cache: make(map[Entry]int)}.countArrangements(springTypes, numbers, -1)
	}

	return output
}

func unfold(springTypes []SpringType, numbers []int, count int) ([]SpringType, []int) {
	var resSpringTypes []SpringType
	var resNumbers []int
	for i := 0; i < count; i++ {
		resSpringTypes = append(resSpringTypes, springTypes...)
		if i != count-1 {
			resSpringTypes = append(resSpringTypes, springTypeUnknown)
		}
		resNumbers = append(resNumbers, numbers...)
	}
	return resSpringTypes, resNumbers
}
