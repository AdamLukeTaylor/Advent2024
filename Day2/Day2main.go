package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// run the following:
// go mod init dummy.test/create
// go mod tidy
// go run . -example=true -part1=true -part2=false

func main() {
	log.SetPrefix("Day2: ")
	example := flag.Bool("example", true, "")
	runPart1 := flag.Bool("part1", false, "")
	runPart2 := flag.Bool("part2", false, "")
	flag.Parse()
	var lines = readToLines(*example)

	if *runPart1 {
		part1(lines)
	}
	if *runPart2 {
		part2(lines)
	}
}

func part1(lines []string) {
	println("---------Part 1--------------")
	var result []bool
	for _, line := range lines {
		split := strings.Split(line, " ")
		// fmt.Println("Result:", split)
		var report []int
		for _, element := range split {
			elemNum, _ := strconv.Atoi(element)
			report = append(report, elemNum)
		}
		outcome := validDecreasing(report) || validIncreasing(report)
		if outcome {
			fmt.Println("report:", report)
			fmt.Println("outcome:", outcome)
		}
		result = append(result, outcome)
	}
	fmt.Println("result:", result)
	var resultCount = 0
	for _, element := range result {
		if element {
			resultCount++
		}
	}
	fmt.Println("resultCount:", resultCount)
	// 380 high
}

func part2(lines []string) {
	println("---------Part 2--------------")
	var result []bool
	for i, line := range lines {
		split := strings.Split(line, " ")
		// fmt.Println("Result:", split)
		var report []int
		for _, element := range split {
			elemNum, _ := strconv.Atoi(element)
			report = append(report, elemNum)
		}
		outcome := validDecreasing(report) || validIncreasing(report) || validDecreasingDampened(report) || validIncreasingDampened(report)
		//if outcome {
		fmt.Println("++++report:", report, i)
		fmt.Println("++++outcome:", outcome)
		//}
		result = append(result, outcome)
	}
	//fmt.Println("result:", result)
	var resultCount = 0
	for _, element := range result {
		if element {
			resultCount++
		}
	}
	fmt.Println("resultCount:", resultCount)
	// 391,396,417 low

	// 495 high
}

func validIncreasing(report []int) bool {	
	var output = true
	if len(report) > 2 {
		for index := 0; index < len(report)-1; index++ {
			diff := report[index+1] - report[index]
			if diff < 1 || diff > 3 {
				output = false
			}
		}
	} else {
		output = false
	}
	fmt.Println("validIncreasing:", report, output)
	return output
}

func validDecreasing(report []int) bool {	
	var output = true
	if len(report) > 2 {
		for index := 0; index < len(report)-1; index++ {
			diff := report[index] - report[index+1]
			if diff < 1 || diff > 3 {
				output = false
			}
		}
	} else {
		output = false
	}
	fmt.Println("validDecreasing:", report, output)
	return output
}
func validIncreasingDampened(report []int) bool {
	fmt.Println("validIncreasingDampened:", report)
	var outcomes []bool
	for index := 0; index < len(report); index++ {
		removed := validIncreasing(remove(report,index))
		outcomes = append(outcomes, removed)
		if(removed){
			return true
		}
	}
	return false
}

func validDecreasingDampened(report []int) bool {
	println("validDecreasingDampened")
	var outcomes []bool
	for index := 0; index < len(report); index++ {
		removed := validDecreasing(remove(report,index))
		outcomes = append(outcomes, removed)
		if(removed){
			return true
		}
	}
	return false
}

func remove(slice []int, index int) []int {
    ret := make([]int, 0)
    ret = append(ret, slice[:index]...)
    return append(ret, slice[index+1:]...)
}



func readToLines(example bool) []string {
	var filename = ""
	if example {
		println("---------Example--------------")
		filename = "Day2-example.txt"
	} else {
		println("---------Real data--------------")
		filename = "Day2-data.txt"
	}
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")
	return lines
}
