package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var lineList [][]int
var partOneSum int
var partTwoSum int

func create_lists() {
	input, err := os.Open("./part_1.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	inputScanner := bufio.NewScanner(input)
	inputScanner.Split(bufio.ScanLines)
	for inputScanner.Scan() {
		var intSplit []int
		line := inputScanner.Text()
		lineSplit := strings.Split(line, " ")
		for i := 0; i < len(lineSplit); i++ {
			num, err := strconv.Atoi(lineSplit[i])
			if err != nil {
				panic(err)
			}
			intSplit = append(intSplit, num)
		}
		lineList = append(lineList, intSplit)
	}
}

func isValidSequence(numbers []int) bool {
	if len(numbers) < 2 {
		return true
	}

	increasing := numbers[1] > numbers[0]

	for i := 0; i < len(numbers)-1; i++ {
		diff := numbers[i+1] - numbers[i]
		absDiff := math.Abs(float64(diff))

		if absDiff < 1 || absDiff > 3 {
			return false
		}

		if increasing && diff <= 0 {
			return false
		}
		if !increasing && diff >= 0 {
			return false
		}
	}

	return true
}

func part_1() {
	for _, line := range lineList {
		if isValidSequence(line) {
			partOneSum++
		}
	}
	fmt.Println("The final answer to part 1 is:", partOneSum)
}

func create_deleted_lists(slice []int) [][]int {
	var deletedLists [][]int
	for i := 0; i < len(slice); i++ {
		newList := make([]int, 0, len(slice)-1)
		newList = append(newList, slice[:i]...)
		newList = append(newList, slice[i+1:]...)
		deletedLists = append(deletedLists, newList)
	}
	return deletedLists
}

func part_2() {
	for _, currentLine := range lineList {
		if isValidSequence(currentLine) {
			partTwoSum++
			continue
		}

		deletedLists := create_deleted_lists(currentLine)
		for _, testList := range deletedLists {
			if isValidSequence(testList) {
				partTwoSum++
				break
			}
		}
	}
	fmt.Println("The final answer to part 2 is:", partTwoSum)
}

func main() {
	create_lists()
	part_1()
	part_2()
}
