package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// run the following:
// go mod init dummy.test/create
// go mod tidy
// go run . -example=true -part1=true -part2=false

func main() {
	log.SetPrefix("Day1: ")
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
	re := regexp.MustCompile(`\d+`)
	var left []int
	var right []int
	var diff []int
	for _, line := range lines {
		println(line)
		matches := re.FindAllString(line,-1)
		intL,_:=strconv.Atoi(matches[0])
		intR,_:=strconv.Atoi(matches[1])
		fmt.Println(fmt.Sprintf("l %d",intL))
		fmt.Println(fmt.Sprintf("r %d",intR))
		left = append(left,intL)
		right = append(right,intR)
	}
	sort.Ints(left)
	sort.Ints(right)
	fmt.Println(left)
	fmt.Println(right)
	for i := range left {
		diff=append(diff,Abs(left[i]-right[i]))
	}
	fmt.Println(diff)
	fmt.Println(sum(diff))

}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func sum(nums []int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func part2(lines []string) {
	println("---------Part 2--------------")
}

func readToLines(example bool) []string {
	var filename = ""
	if example {
		println("---------Example--------------")
		filename = "Day1-example.txt"
	} else {
		println("---------Real data--------------")
		filename = "Day1-data.txt"
	}
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")
	return lines
}
