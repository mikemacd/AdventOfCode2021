package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	debug = false
)

func main() {
	lines := readInput()

	x, y := navigateSub(lines)

	fmt.Printf("%d * %d = %d\n", x, y, x*y)

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

func navigateSub(lines []string) (x int, y int) {
	x, y = 0, 0

	for _, line := range lines {
		direction, distance := "", 0
		fmt.Sscanf(line, "%s %d\n", &direction, &distance)
		if debug {
			fmt.Printf("line: '%s' direction:%s distance:%d\n", line, direction, distance)
		}
		switch {
		case direction == "up":
			y -= distance
		case direction == "down":
			y += distance
		case direction == "back":
			x -= distance
		case direction == "forward":
			x += distance
		}
	}

	return
}
