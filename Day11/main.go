package main

import (
	"bufio"
	"fmt"

	"os"
	"strings"
)

type Galaxy struct {
	label int
	row int
	col int
}

func main() {
	lines := lines("input.txt")
	space := getSpace(lines)
	fmt.Println("Part One: ",partOne(space))
	fmt.Println("Part Two: ",partTwo(space))
}

func partOne(space [][]string) int {
	expandedSpace := expandSpace(space)
	galaxies := createGalaxies(expandedSpace)
	return sumShortestDistances(galaxies)
}

func partTwo(space [][]string) int {
	colWithoutGalaxies:= colWithoutGalaxies(space)
	rowWithoutGalaxies := rowWithoutGalaxies(space)
	galaxies := createGalaxies(space)
	number:=1000000
	for i := range galaxies{
		g:= &galaxies[i]
		numberOfRowsBefore:= 0
		numberOfColsBefore:= 0
		for _,c:= range colWithoutGalaxies{
				if c < g.col{
					numberOfColsBefore++
				}
		}
		for _,r:= range rowWithoutGalaxies{
			if r < g.row{
				numberOfRowsBefore++
			}
		}
		g.col = g.col + (numberOfColsBefore*(number-1))
		g.row = g.row + (numberOfRowsBefore*(number-1))
	}
	return sumShortestDistances(galaxies)
}

func createGalaxies(space [][]string) []Galaxy{
	galaxies := []Galaxy{}
	label:= 0
	for row := 0; row < len(space); row++ {
		for col := 0; col < len(space[row]); col++ {
			if space[row][col] == "#" {
				label++
				galaxies = append(galaxies, Galaxy{label: label,row: row, col: col})
			}
		}
	}
	return galaxies
}

func sumShortestDistances(galaxies []Galaxy) int {
	output:=0
	for i:=0; i < len(galaxies)-1; i++{
		firstGalaxy:= galaxies[i]
		for j:=i+1; j<len(galaxies);j++{
			secondGalaxy:= galaxies[j]
			shortestDistance := shortestDistance(firstGalaxy, secondGalaxy)
			output = output +shortestDistance
		}
	}
	return output
}

func colWithoutGalaxies(space [][]string) []int {
	cols:=[]int{}
	for col := 0; col < len(space); col++ {
		flag := false
		for row := 0; row < len(space); row++ {
			if space[row][col] == "#" {
				flag = true
				break
			}
		}
		if !flag{
			cols= append(cols, col)
		}
		
	}
	return cols
}

func rowWithoutGalaxies(space [][]string) []int {
	rows:=[]int{}
	for row := 0; row < len(space); row++ {
		flag := false
		for col := 0; col < len(space); col++ {
			if space[row][col] == "#" {
				flag = true
				break
			}
		}
		if !flag{
			rows= append(rows, row)
		}
		
	}
	return rows
}

func expandSpace(space [][]string) [][]string {
	colWithoutGalaxies:= colWithoutGalaxies(space)
	expandedSpace:= [][]string{}
	for row := 0; row < len(space); row++ {
		flag := false
		thisRow := space[row]
		for col := 0; col < len(thisRow); col++ {
			if thisRow[col] == "#" {
				flag = true
			}
		}
		//insert expanded cols
		i:=0
		for _,j := range colWithoutGalaxies{
			thisRow = insert(thisRow,j+i,".")
			i++
		}
		//insert expanded rows
		if flag{		
			expandedSpace= append(expandedSpace, thisRow)
		}else{
			expandedSpace= append(expandedSpace, thisRow)
			expandedSpace= append(expandedSpace, thisRow)
		}
	}
	return expandedSpace
}

func insert(a []string, index int, value string) []string {
    a = append(a[:index+1], a[index:]...) 
    a[index] = value                     
    return a
}

func shortestDistance(firstGalaxy Galaxy, secondGalaxy Galaxy) int {
	if firstGalaxy.col == secondGalaxy.col{
		return absDiffInt(firstGalaxy.row, secondGalaxy.row)
	} else if firstGalaxy.row == secondGalaxy.row{
		return absDiffInt(firstGalaxy.col, secondGalaxy.col)
	}else{
		return absDiffInt(firstGalaxy.col, secondGalaxy.col) + absDiffInt(firstGalaxy.row, secondGalaxy.row)
	}
}

func absDiffInt(x, y int) int {
	if x < y {
	   return y - x
	}
	return x - y
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

func getSpace(lines []string) [][]string {
	space := make([][]string, len(lines))
	for i := 0; i < len(lines); i++ {
		space[i] = strings.Split(lines[i], "")
	}
	return space
}
