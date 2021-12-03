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

func main() {
	lines := readInput()

	o2rating := getRating("o2", 0, lines)
	co2rating := getRating("co2", 0, lines)

	fmt.Printf("%d * %d = %d\n", o2rating, co2rating, o2rating*co2rating)

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

func getDiagnostics(lines []string) (diagnostics map[int]map[int]int) {
	diagnostics = map[int]map[int]int{}
	for _, line := range lines {
		for j, b := range line {
			value, _ := strconv.Atoi(string(b))

			if _, ok := diagnostics[value]; !ok {
				diagnostics[value] = map[int]int{}
			}

			diagnostics[value][j]++
		}
	}
	return
}

func getRating(deviceType string, step int, lines []string) (rating int) {
	// fmt.Printf("%s %d %d\n", deviceType, step, len(lines) )

	filtered := []string{}

	diagnostics := getDiagnostics(lines)
	// fmt.Printf("diag: %v\n",diagnostics )
	for _, line := range lines {
		switch {
		case deviceType == "o2":
			// fmt.Printf("%d %d -- %d >= %d = %v\n",'1',line[step] , diagnostics[1][step] , diagnostics[0][step], diagnostics[1][step] >= diagnostics[0][step] )
			if diagnostics[1][step] >= diagnostics[0][step] && line[step] == '1' {
				filtered = append(filtered, line)
			}
			if diagnostics[1][step] < diagnostics[0][step] && line[step] == '0' {
				filtered = append(filtered, line)
			}
		case deviceType == "co2":
			if diagnostics[0][step] <= diagnostics[1][step] && line[step] == '0' {
				filtered = append(filtered, line)
			}
			if diagnostics[0][step] > diagnostics[1][step] && line[step] == '1' {
				filtered = append(filtered, line)
			}
		}
	}

	if len(filtered) == 1 {
		r, _ := strconv.ParseInt(filtered[0], 2, 32)
		return int(r)
	}
	if len(filtered) == 0 {
		return 0
	}

	return getRating(deviceType, step+1, filtered)
}
