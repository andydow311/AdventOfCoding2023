package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	lines := lines("input.txt")
	directions := lines[0]
	format := make(map[string][]string)
	for i := 2; i < len(lines); i++ {
		node := strings.Trim(strings.Split(lines[i], " = ")[0]," ")
		left:= strings.Trim(strings.Split(strings.ReplaceAll(strings.ReplaceAll(strings.Split(lines[i], " = ")[1],"(",""),")",""),",")[0]," ")
		right:= strings.Trim(strings.Split(strings.ReplaceAll(strings.ReplaceAll(strings.Split(lines[i], " = ")[1],"(",""),")",""),",")[1]," ")
		format[node]= []string{left,right}	
	}
	fmt.Println("Part TWO: ", partTwo(directions, format))
}

func partTwo(directions string, format map[string][]string) int {
	nodesEndingInA := getNodesEndingInA(format)
	nodesEndingInZ := getNodesEndingInZ(format)
	tmpArray := []string{}
	numberOfStep:= 0
	stepsTofirstHitZ := make(map[string]int)
	for i:=0; i < len(directions); i++{
		for _,node := range nodesEndingInA{
			if string(directions[i]) == "R"{
				tmpArray = append(tmpArray, format[node][1])
			}else{
				tmpArray = append(tmpArray, format[node][0])
			} 
		}
		numberOfStep++
		containsZ(tmpArray, stepsTofirstHitZ, numberOfStep) 
		if len(stepsTofirstHitZ) == len(nodesEndingInZ){
			break
		}else{
			nodesEndingInA = tmpArray
			tmpArray = []string{}
		}

		if i ==len(directions)-1{
			directions = directions +directions
		}
	
	}
    answers := []int{}
	for _,v := range stepsTofirstHitZ{
		answers = append(answers, v)
	}
	return lowestCommonMultiple(answers[0], answers[1], answers[2:]...)
}

func containsZ(tmpArray []string, stepsTofirstHitZ map[string]int , numberOfStepsint int) {
	for _, n := range tmpArray{
		if string(n[len(n)-1]) == "Z"{
			stepsTofirstHitZ[n] = numberOfStepsint
		}
	}	
}

func greeastCommonDivisor(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lowestCommonMultiple(a, b int, integers ...int) int {
	result := a * b / greeastCommonDivisor(a, b)
	for i := 0; i < len(integers); i++ {
		result = lowestCommonMultiple(result, integers[i])
	}
	return result
}

func equalArrays(a []string, b []string){
	for _, value := range a{
		for _, value2 := range b{
			if value == value2{

			}
		
		}
	}
}

func anyNodesEndInZ(nodes []string, numberOfSteps int) {
	for _,k:= range nodes {	
		if string(k) == "ZZZ"{
			fmt.Println("node: ", k, " ====> ", numberOfSteps)
		}
	}
}

func getNodesEndingInA(format map[string][]string) []string {
	output := []string{}
	for k := range format {	
		if string(k[len(k)-1]) == "A"{
			output = append(output,k)
		}
	}
	return output
}

func getNodesEndingInZ(format map[string][]string) []string {
	output := []string{}
	for k := range format {	
		if string(k[len(k)-1]) == "Z"{
			output = append(output,k)
		}
	}
	return output
}

func partOne(directions string, format map[string][]string) int {
	numberOfStep:=0
	node:= "AAA"
	end:="ZZZ"
	for i:=0; i < len(directions); i++{
		if string(directions[i]) == "R"{
			node = format[node][1]
		}else{
			node = format[node][0]
		} 
		numberOfStep++
		if node == end{
			break
		}
		if i ==len(directions)-1{
			directions = directions +directions
		}
	}
	return numberOfStep
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



