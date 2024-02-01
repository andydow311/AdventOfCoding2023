package main

import (
	"bufio"
	"fmt"

	"os"

	"strings"
)

type Tile struct {
	column int
	row    int
}

func main() {
	lines := lines("input.txt")
	//fmt.Println("1: ", partOne(lines))
	fmt.Println("2: ", partTwo(lines))
}

func partTwo(lines []string) int {
	grid := getGrid(lines)
	startingTile := getStartingPoint(grid)
	neighbours, froms := getNeighbours(startingTile, grid)
	for i, n := range neighbours {
		loop, found := findLoop(grid, n, froms[i])
		if found {
			fmt.Println(loop)
			gridDis:= displayMap(grid, loop)
			fmt.Println(gridDis)
			count:=0
			for col := 0; col < len(gridDis); col++ {
				flag:= false
				for row := 0; row < len(gridDis[col])-1; row++ {
					if gridDis[col][row] == "X" && !flag{
						flag = true
					}else if gridDis[col][row] == "X" && flag{
						flag = false
					}

					if gridDis[col][row] != "X" && flag{
						count++
					}
				}
			}

			for i:=0; i< len(gridDis); i++{
				fmt.Println(gridDis[i])
			}

			return count
		}
	}
	panic("no solution found")
}

func displayMap(grid [][]string, paths []Tile) [][] string {
	thisGrid := grid
	for col := 0; col < len(grid); col++ {
		for row := 0; row < len(grid[col]); row++ {
			if contains(paths, Tile{column: col, row: row}){
				thisGrid[col][row] = "X"
			}else{
				thisGrid[col][row] = grid[col][row] 
			}
		}
	}
	for col := 0; col < len(thisGrid); col++ {
		x:= ""
		for row := 0; row < len(thisGrid[col]); row++ {
			x = x+thisGrid[col][row]
		}
		fmt.Println(x)
	}
	return thisGrid
}

func contains(paths []Tile, tile Tile) bool {
	for _, thisTile := range paths{
		if tile.column == thisTile.column && thisTile.row == tile.row{
			return true
		}
	}
	return false
}

func partOne(lines []string) int {
	grid := getGrid(lines)
	startingTile := getStartingPoint(grid)
	neighbours, froms := getNeighbours(startingTile, grid)
	for i, n := range neighbours {
		loop, found := findLoop(grid, n, froms[i])
		if found {
			var l = len(loop)
			if l%2 == 0 {
				return l / 2
			}
			return l/2 + 1
		}
	}
	panic("no solution found")
}

func getNeighbours(startingTile Tile, grid [][]string) ([]Tile, []string) {
	neighbours := []Tile{}
	froms := []string{}
	if startingTile.column+1 < len(grid)-1 {
		neighbours = append(neighbours, Tile{startingTile.column + 1, startingTile.row})
		froms = append(froms, "NORTH")
	}
	if startingTile.column-1 > 0 {
		neighbours = append(neighbours, Tile{startingTile.column - 1, startingTile.row})
		froms = append(froms, "SOUTH")
	}
	if startingTile.row+1 < len(grid[startingTile.column])-1 {
		neighbours = append(neighbours, Tile{startingTile.column, startingTile.row + 1})
		froms = append(froms, "WEST")
	}
	if startingTile.row-1 > 0 {
		neighbours = append(neighbours, Tile{startingTile.column, startingTile.row - 1})
		froms = append(froms, "EAST")
	}
	return neighbours, froms
}

func findLoop(grid [][]string, tile Tile, from string) (path []Tile, ok bool) {
	path = []Tile{}
	path = append(path, tile)
	for {
		nextTile, newfrom, found := step(grid, tile, from)
		if !found {
			return path, false
		}
		path = append(path, nextTile)
		tileLabel := grid[nextTile.column][nextTile.row]
		if tileLabel == "S" {
			return path, true
		}
		tile = nextTile
		from = newfrom
	}
}

func step(grid [][]string, tile Tile, from string) (nextTile Tile, newFrom string, ok bool) {
	if grid[tile.column][tile.row] == "." {
		return tile, from, false
	}
	if grid[tile.column][tile.row] == "S" {
		return tile, from, true
	}
	switch grid[tile.column][tile.row] {
	case "|":
		if from == "NORTH" {
			return Tile{column: tile.column + 1, row: tile.row}, from, true
		} else if from == "SOUTH" {
			return Tile{column: tile.column - 1, row: tile.row}, from, true
		}
	case "-":
		if from == "EAST" {
			return Tile{column: tile.column, row: tile.row - 1}, from, true
		} else if from == "WEST" {
			return Tile{column: tile.column, row: tile.row + 1}, from, true
		}
	case "F":
		if from == "SOUTH" {
			return Tile{column: tile.column, row: tile.row + 1}, "WEST", true
		} else if from == "EAST" {
			return Tile{column: tile.column + 1, row: tile.row}, "NORTH", true
		}
	case "7":
		if from == "SOUTH" {
			return Tile{column: tile.column, row: tile.row - 1}, "EAST", true
		} else if from == "WEST" {
			return Tile{column: tile.column + 1, row: tile.row}, "NORTH", true
		}
	case "L":
		if from == "NORTH" {
			return Tile{column: tile.column, row: tile.row + 1}, "WEST", true
		} else if from == "EAST" {
			return Tile{column: tile.column - 1, row: tile.row}, "SOUTH", true
		}
	case "J":
		if from == "NORTH" {
			return Tile{column: tile.column, row: tile.row - 1}, "EAST", true
		} else if from == "WEST" {
			return Tile{column: tile.column - 1, row: tile.row}, "SOUTH", true
		}
	}
	return tile, from, false
}

func getStartingPoint(grid [][]string) Tile {
	for col := 0; col < len(grid); col++ {
		for row := 0; row < len(grid[col]); row++ {
			if grid[col][row] == "S" {
				return Tile{column: col, row: row}
			}
		}
	}
	panic("No start pos found")
}

func getGrid(lines []string) [][]string {
	grid := make([][]string, len(lines))
	for i := 0; i < len(lines); i++ {
		grid[i] = strings.Split(lines[i], "")
	}
	return grid
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
