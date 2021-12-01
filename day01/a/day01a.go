package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	numbers := readInput()

	n := countIncreses(numbers)
	fmt.Printf("%d\n", n)
	os.Exit(0)
}

func readInput() []int {
	var numbers []int

	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return []int{}
	}
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		panic(err)
	}

	lines := bytes.Split(data, []byte("\n"))
	for i, line := range lines {
		num, err := strconv.Atoi(string(line))
		if err != nil {
			log.Fatalf("Can't parse number on line %d: %v\n", i, line)
		}
		numbers = append(numbers, num)
	}

	return numbers
}

func countIncreses(numbers []int) int {
	increases := 0
	previousNumber := numbers[0]

	for _, num := range numbers {
		if num > previousNumber {
			increases++
		}
		previousNumber = num
	}
	return increases
}
