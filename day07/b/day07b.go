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

type ints []int

func main() {
	lines := readInput()

	numbers := strings.Split(lines[0], ",")
	nums := ints{}
	for _, v := range numbers {
		num, _ := strconv.Atoi(v)
		nums = append(nums, num)
	}
	sort.Ints(nums)

	fuelUsed := ints{}
	for i := nums[0]; i < nums[len(nums)-1]; i++ {
		fuelUsed = append(fuelUsed, int(nums.fuelUsed(i)))
	}
	sort.Ints(fuelUsed)

	fmt.Printf("%v\n", fuelUsed[0])

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

	n := len(nums)
	if n%2 != 0 {
		return nums[n-1] + nums[n]/2
	}
	return nums[n/2]
}

func (nums ints) fuelUsed(m int) float64 {
	sum := 0.0
	for _, v := range nums {
		if v > m {
			sum += float64(v-m) / 2 * (1 + float64(v-m))
		}
		if v < m {
			sum += float64(m-v) / 2 * (1 + float64(m-v))
		}
	}
	return sum
}
