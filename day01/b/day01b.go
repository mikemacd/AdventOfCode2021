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
	debug := false
	increases := 0
	previousNumber := numbers[0]

	for i := range numbers {
		sum := 0
		for n := range []int{0, 1, 2} {
			if i+n < len(numbers) {
				sum += numbers[i+n]
			}
		}

		if debug {
			fmt.Printf("%d %d", i, sum)
		}
		if sum > previousNumber && i > 0 {
			increases++
			if debug {
				fmt.Print(" increases")
			}
		}

		if debug {
			fmt.Println("")
		}

		previousNumber = sum

	}
	return increases
}
