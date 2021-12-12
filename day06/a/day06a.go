package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	debug = false
)

type fish struct {
	timer int
}

type timers map[int]int

func main() {
	lines := readInput()

	school := createSchool(lines)

	for i := 0; i < 80; i++ {
		newSchool := timers{}
		newSchool[8] = school[0]
		newSchool[7] = school[8]
		newSchool[6] = school[7]
		newSchool[5] = school[6]
		newSchool[4] = school[5]
		newSchool[3] = school[4]
		newSchool[2] = school[3]
		newSchool[1] = school[2]
		newSchool[0] = school[1]
		newSchool[6] += school[0]

		school = newSchool
	}

	size := 0
	for i := range school {
		size += school[i]
	}

	fmt.Printf("school size : %v\n", size)

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

func createSchool(lines []string) timers {
	school := timers{}

	fishes := strings.Split(lines[0], ",")
	for _, v := range fishes {
		t, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		school[t]++
	}
	return school
}
