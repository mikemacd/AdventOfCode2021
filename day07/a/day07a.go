package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	debug = false
)

func main() {
	lines := readInput()

	numbers := strings.Split(lines[0], ",")
	nums := []int{}
	for _, v := range numbers {
		num, _ := strconv.Atoi(v)
		nums = append(nums, num)
	}

	m := median(nums)
	sum := 0
	for _, v := range nums {
		if v > m {
			sum += (v - m)
		}
		if v < m {
			sum += (m - v)
		}
	}

	fmt.Printf("%v\n", sum)

	os.Exit(0)
}

func readInput() []string {

	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return []string{}
	}
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	output := []string{}
	for _, line := range bytes.Split(data, []byte("\n")) {
		output = append(output, strings.TrimSpace(string(line)))
	}

	return output
}

func median(nums []int) int {
	sort.Ints(nums)

	n := len(nums)
	if n%2 != 0 {
		return nums[n-1] + nums[n]/2
	}
	return nums[n/2]
}
