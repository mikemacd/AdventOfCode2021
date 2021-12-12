package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

var (
	debug = false
)

type point struct {
	x int
	y int
}

type bingoBoard map[point]string
type haveNumber map[string]point
type hasNumbers map[int]haveNumber

func main() {
	lines := readInput()

	callOrder := strings.Split(lines[0], ",")

	boards, boardsHave := parseBoards(lines[1:])
	// fmt.Printf("bh: %v\n\n", boardsHave)
	calledBoards := make(map[int]bingoBoard, len(boards))

	// for i, v := range boardsHave {
	// 	fmt.Printf("i:%d %v\n", i, v)
	// }
	completedOrder := []struct {
		i    int
		call int
		sum  int
	}{}

	for _, call := range callOrder {
		// fmt.Printf("Calling: %s\n", call)
		// mark boards
		for j := range boards {
			if _, ok := calledBoards[j]; !ok {
				bb := make(bingoBoard, 1)
				calledBoards[j] = bb
			}
			if p, ok := boardsHave[j][call]; ok {
				calledBoards[j][p] = "X"
			}
		}
		// check for bingo
		for i, b := range boards {
			if hasBingo(calledBoards[i]) {
				sum := b.calculateSumUnmarked(calledBoards[i])
				callValue, _ := strconv.Atoi(call)
				// fmt.Printf("sum: %d call:w %d = %d\n", sum, callValue, sum*callValue)
				found := false
				for _, v := range completedOrder {
					if v.i == i {
						found = true
					}
				}
				if !found {
					completedOrder = append(completedOrder, struct {
						i    int
						call int
						sum  int
					}{i: i, call: callValue, sum: sum})
				}
			}
		}
	}

	sum := completedOrder[len(boards)-1].sum * completedOrder[len(boards)-1].call
	spew.Dump(sum)
	// spew.Dump(completedOrder, sum)
	// spew.Dump(callOrder, boards, boardsHave)

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

func parseBoards(lines []string) (boards map[int]bingoBoard, boardsHave hasNumbers) {
	boards = map[int]bingoBoard{}
	boardsHave = hasNumbers{}

	if len(lines)%6 != 0 {
		fmt.Printf("Wrong number of lines %d in:\n%v\n", len(lines), lines)
		return nil, nil
	}
	for board := 0; board < len(lines)/6; board++ {
		boards[board] = map[point]string{}
		boardsHave[board] = map[string]point{}
		for y := 0; y < 5; y++ {
			for x, v := range strings.Fields(lines[board*6+1+y]) {
				p := point{x, y}
				boards[board][p] = string(v)
				boardsHave[board][v] = point{x, y}
			}
		}
	}

	return
}

func hasBingo(b bingoBoard) bool {
	for y := 0; y < 5; y++ {
		line := true
		for x := 0; x < 5; x++ {
			if _, ok := b[point{x, y}]; !ok {
				line = false
			}
		}
		if line {
			return true
		}
	}
	for x := 0; x < 5; x++ {
		col := true
		for y := 0; y < 5; y++ {
			if _, ok := b[point{x, y}]; !ok {
				col = false
			}
		}
		if col {
			return true
		}
	}

	// diagonals apparently don't count
	// diagA := true
	// diagB := true
	// for xy := 0; xy < 5; xy++ {
	// 	if _, ok := b[point{xy, xy}]; !ok {
	// 		diagA = false
	// 	}
	// 	if _, ok := b[point{xy, 4 - xy}]; !ok {
	// 		diagB = false
	// 	}
	// }
	// if diagA || diagB {

	// 	return true
	// }

	return false

}
func (b bingoBoard) calculateSumUnmarked(bb bingoBoard) int {
	sum := 0

	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if _, ok := bb[point{x, y}]; !ok {
				v, _ := strconv.Atoi(b[point{x, y}])
				sum += v
			}
		}
	}
	return sum
}
func (b bingoBoard) print(bb bingoBoard) {
	for y := 0; y < 5; y++ {
		fmt.Printf("\t")
		for x := 0; x < 5; x++ {
			fmt.Printf(" %v", b[point{x, y}])
		}
		fmt.Printf("\t")

		for x := 0; x < 5; x++ {
			if n, ok := bb[point{x, y}]; ok {
				fmt.Printf(" %v", n)
			} else {
				fmt.Printf(" __")
			}
		}
		fmt.Println("")
	}
	fmt.Println("\n")
}

func (b hasNumbers) Len() int {
	return len(b)
}
func (b hasNumbers) Less(i int, j int) bool {
	return i < j
}
func (b hasNumbers) Swap(i int, j int) {
	b[i], b[j] = b[j], b[i]
}
