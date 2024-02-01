package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	lines := lines("input.txt")
	partOne(lines)
}

func partOne(lines []string){
	score:=0
	mirrors:= [][][]string{}
	mirror:= [][]string{}
	for _, line :=range lines{
		if strings.Trim(line, " ") == ""{
			mirrors = append(mirrors, mirror)
			mirror= [][]string{}
		}else{
			mirror = append(mirror, strings.Split(strings.Trim(line, " "),""))
		}
	}
	for _, mirror := range mirrors{
		t,a,_ := findReflection(mirror)
		switch t{
			case "col": score = score + a+1
			case "row": score = score + (100*(a+1))
			default: fmt.Println("no reflection found")
		}
		fmt.Println("M: ", mirror," score:",score)
	}
	

}

func findReflection(mirror [][]string) (string, int, int){
	
	for i:=0; i< len(mirror)-1; i++ {
		match:=true
		for j :=range mirror[i]{
			if mirror[i][j] != mirror[i+1][j]{
					match = false
			}
		}
		if match && validRowReflection(mirror,i,i+1){
			return "row", i, i+1
		}
	}

	for k:=0; k< len(mirror[0])-1; k++ {
		match:=true
		for j :=range mirror{
			if mirror[j][k] != mirror[j][k+1]{
					match = false
			}
		}
		if match && validColReflection(mirror,k,k+1){
			return "col", k, k+1
		}
	}


	return "None", -1, -1
}

func validRowReflection(mirror [][]string, above,below int) bool {
	checks:=minimum(above, len(mirror)-1-below)
	for i:=0; i<=checks;i++{
		for j :=range mirror[0]{
			if mirror[above-i][j] != mirror[below+i][j]{
					return false
			}
		}
	}
	return true
}

func validColReflection(mirror [][]string, above,below int) bool {
	checks:=minimum(above, len(mirror[0])-1-below)
	for i:=0; i<=checks;i++{
		for j :=range mirror{
			if mirror[j][above-i] != mirror[j][below+i]{
					return false
			}
		}
	}
	return true
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

func minimum(a,b int) int{
	if a>b{
		return b
	}else{
		return a
	}
}