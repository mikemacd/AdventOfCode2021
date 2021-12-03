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

	gamma, epsilon := runDiagnostic(lines)

	fmt.Printf("%d * %d = %d\n", gamma, epsilon, gamma*epsilon)

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

func runDiagnostic(lines []string) (gamma int, epsilon int) {
	diagnostics := map[int]map[int]int{}
	for _, line := range lines {
		for j, b := range line {
			value, _ := strconv.Atoi(string(b))

			if _, ok := diagnostics[value]; !ok {
				diagnostics[value] = map[int]int{}
			}

			diagnostics[value][j]++
		}
	}

	// fmt.Printf("%+v\n", diagnostics)
	width := len(lines[0])
	gs := ""
	es := ""
	for i := width; i >= 0; i-- {
		if diagnostics[1][width-i] > diagnostics[0][width-i] {
			gs += "1"
			es += "0"
		}
		if diagnostics[1][width-i] < diagnostics[0][width-i] {
			gs += "0"
			es += "1"
		}
	}
	g, _ := strconv.ParseInt(gs, 2, 32)
	e, _ := strconv.ParseInt(es, 2, 32)
	gamma = int(g)
	epsilon = int(e)
	//fmt.Printf("V:%+v\n", diagnostics)
	//fmt.Printf("gamma:%d epsilon:%d gs:%s es:%s\n", gamma, epsilon, gs, es)
	return gamma, epsilon
}
