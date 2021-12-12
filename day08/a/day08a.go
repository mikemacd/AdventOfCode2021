package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mikemacd/AdventOfCode2021/util"
)

var (
	debug = false
)

func main() {
	mapping := map[int]int{
		0: 6,
		1: 2, //
		2: 5,
		3: 5,
		4: 4, //
		5: 5,
		6: 6,
		7: 3, //
		8: 7, //
		9: 6,
	}
	_ = mapping
	reverseMapping := map[int][]int{
		2: {1}, //
		3: {7}, //
		4: {4}, //
		5: {2, 3, 5},
		6: {9, 0, 6},
		7: {8}, //
	}
	_ = reverseMapping
	lines := util.ReadInput()

	sum := 0
	for _, l := range lines {
		halfs := strings.Split(l, "|")
		for hn, h := range halfs {
			h = strings.TrimSpace(h)
			//fmt.Printf("hn:%v h: %v\n", hn, h)
			words := strings.Split(h, " ")

			for wn, w := range words {
				_ = wn
				w = strings.TrimSpace(w)
				if hn == 1 {
					// fmt.Printf("hn: %v wn: %v w: %v %v\n", hn, wn, len(w), w)
					switch reverseMapping[len(w)][0] {
					case 1, 4, 7, 8:
						sum++
					}
				}
			}
		}
	}

	fmt.Printf("%v\n", sum)

	os.Exit(0)
}
