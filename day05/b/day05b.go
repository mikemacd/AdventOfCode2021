package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

var (
	debug = false
)

type point struct {
	x int
	y int
}
type field map[point][]int

func main() {
	lines := readInput()
	input, maxX, maxY, minX, minY := parseInput(lines)
	fmt.Printf("%d,%d, %d, %d \n", maxX, maxY, minX, minY)
	//input.print(maxX, maxY, minX, minY)

	crossings := input.calculateCrossings(maxX, maxY, minX, minY)

	fmt.Printf("crossings: %v\n", crossings)

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
func parseInput(lines []string) (f field, maxX, maxY, minX, minY int) {
	f = field{}
	maxX, maxY = 0, 0
	minX, minY = 10000, 10000

	for i, v := range lines {
		p1 := point{}
		p2 := point{}
		fmt.Sscanf(v, "%d,%d -> %d,%d\n", &p1.x, &p1.y, &p2.x, &p2.y)
		// fmt.Printf("i:%d p1:%v p2:%v\n", i, p1, p2)

		// slope := 0
		// if p2.x-p1.x == 0 {
		// 	slope = 8
		// } else {
		// 	slope = (p2.y - p1.y) / (p2.x - p1.x)
		// }
		// if slope != 0 && slope != 1 && slope != -1 && slope != 8 {
		// 	fmt.Printf("line %d has slope %d\n", i, slope)
		// }

		if true || p1.x == p2.x || p1.y == p2.y {
			x, y, xDelta, yDelta := p1.x, p1.y, 0, 0
			switch {
			case p1.x < p2.x:
				xDelta = 1
			case p1.x > p2.x:
				xDelta = -1
			}
			switch {
			case p1.y < p2.y:
				yDelta = 1
			case p1.y > p2.y:
				yDelta = -1
			}

			// 2,2 -> 2,1
			ll := int(math.Abs(float64(p2.x)-float64(p1.x)))  +1
			if math.Abs(float64(p2.x)-float64(p1.x)) == 0 {
				ll = int(math.Abs(float64(p2.y)-float64(p1.y)))+1
			}
			// fmt.Printf("line:%v has length: %v\n", i,ll)
			for k := 0; k < int(ll); k++ {
				if x > maxX {
					maxX = x
				}
				if x < minX {
					minX = x
				}
				if y > maxY {
					maxY = y
				}
				if y < minY {
					minY = y
				}

				if _, ok := f[point{x: x, y: y}]; !ok {
					f[point{x: x, y: y}] = []int{}
				}
				f[point{x: x, y: y}] = append(f[point{x: x, y: y}], i)
				x += xDelta
				y += yDelta
			}
		}

	}

	return
}

func (input field) print(maxX, maxY, minX, minY int) {
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if l, ok := input[point{x, y}]; ok {
				fmt.Printf(" %v", len(l))
				// fmt.Printf(" %v", (l))
			} else {
				fmt.Printf(" .")
			}
		}
		fmt.Println()
	}

}

func (input field) calculateCrossings(maxX, maxY, minX, minY int) (crossings int) {
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if l, ok := input[point{x, y}]; ok {
				if len(l) > 1 {
					// fmt.Printf("Crossing at %v %v : %d -- %v\n", x,y, len(l), l )
					crossings++
				}
			}
		}
	}

	return
}
