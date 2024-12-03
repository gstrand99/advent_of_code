package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var listOne []int
var listTwo []int
var partOneSum int
var partTwoSum int

func create_lists() {
	input, err := os.Open("./part_1.txt")
	if err != nil {
		panic(err)
	}

	inputScanner := bufio.NewScanner(input)

	inputScanner.Split(bufio.ScanLines)

	for inputScanner.Scan() {
		line := inputScanner.Text()
		test := strings.Split(line, "   ")

		numOne, err := strconv.Atoi(test[0])
		if err != nil {
			panic(err)
		}
		listOne = append(listOne, numOne)

		numTwo, err := strconv.Atoi(test[1])
		if err != nil {
			panic(err)
		}
		listTwo = append(listTwo, numTwo)
	}

	input.Close()
}

func part_1() {
	slices.Sort(listOne)
	slices.Sort(listTwo)

	for i := 0; i < len(listOne); i++ {
		if listOne[i] < listTwo[i] {
			partOneSum += listTwo[i] - listOne[i]
		} else if listOne[i] > listTwo[i] {
			partOneSum += listOne[i] - listTwo[i]
		}
	}

	fmt.Println("The final answer to part 1 is: ", partOneSum)
}

func part_2() {
	intMap := make(map[int]int)
	for i := 0; i < len(listTwo); i++ {
		if _, ok := intMap[listTwo[i]]; ok {
			intMap[listTwo[i]] += 1
		} else {
			intMap[listTwo[i]] = 1
		}
	}

	for i := 0; i < len(listOne); i++ {
		if val, ok := intMap[listOne[i]]; ok {
			partTwoSum += val * listOne[i]
		}
	}

	fmt.Println("The final answer to part 2 is: ", partTwoSum)
}

func main() {
	create_lists()
	part_1()
	part_2()
}
