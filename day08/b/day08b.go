package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/mikemacd/AdventOfCode2021/util"
)

var (
	debug = false

	// given segments labeled as:
	// 	segments:
	//  0000
	// 1    2
	// 1    2
	//  3333
	// 4    5
	// 4    5
	//  6666

	digits = map[int][]int{
		0: {0, 1, 2, 4, 5, 6},
		1: {2, 5},
		2: {0, 2, 3, 4, 6},
		3: {0, 2, 3, 5, 6},
		4: {1, 2, 3, 5},
		5: {0, 1, 3, 5, 6},
		6: {0, 1, 3, 4, 5, 6},
		7: {0, 2, 5},
		8: {0, 1, 2, 3, 4, 5, 6},
		9: {0, 1, 2, 3, 5, 6},
	}
)

func main() {
	_ = spew.Sdump("")
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
		wrds := map[int][]string{}
		outputs := []string{}
		for hn, h := range halfs {
			h = strings.TrimSpace(h)
			// fmt.Printf("hn:%v h: %v\n", hn, h)
			words := strings.Split(h, " ")

			for wn, w := range words {
				_ = wn
				w = strings.TrimSpace(w)

				s := strings.Split(w, "")
				sort.Strings(s)
				w = strings.Join(s, "")

				if hn == 0 {
					wrds[len(w)] = append(wrds[len(w)], w)
				}

				if hn == 1 {
					outputs = append(outputs, w)
				}
			}
		}

		d, rd := decode(wrds)
		_ = rd

		// fmt.Printf("outputs:%v\n", outputs)
		// fmt.Printf("Digits: %+v\n", digits)
		// fmt.Printf("Decode:  %v\n", d)
		// fmt.Printf("RDecode: %v\n", rd)

		answer := 0
		for _, v := range outputs {
			//			fmt.Printf("i:%v ", i)
			output := []int{}
			for j, c := range v {
				_ = j
				// fmt.Printf("%v:%v ",string(c),d[string(c)])

				output = append(output, d[string(c)])
			}
			sort.Ints(output)

			//fmt.Printf("output:%v -- %v \n", output, reverseDigits(output))
			r := reverseDigits(output)

			answer = answer*10 + r
		}
		fmt.Printf("sum: %11v outputs: %8v ---- %v\n", sum, outputs, answer)

		sum += answer

	}

	fmt.Printf("%v\n", sum)

	os.Exit(0)
}

// -The words of length 6 (representing numbers 0,9,6) have letters which occur twice for segments (2,3,4 not in order ),
// of those one of them was previously identified as segment 4 which means that the 6lw that does not have that letter is number "9"
// which means that the number "0 and 6" have the other two leters (for segments 2,3). One of those two letters is in number
// "1" making it segment 2 (and the other letter in the 2lw is segment 5) and meaning that that was number "0". Which in turn means
// that number "6" has the other segment 3. The last unaccounted for letter is for segment 6

func decode(wrds map[int][]string) (map[string]int, map[int]string) {
	rv := map[string]int{}
	vr := map[int]string{}

	// -the seg1 that is not in the length 3 word is segment 0
	seg0 := regexp.MustCompile("["+wrds[2][0]+"]").ReplaceAllString(wrds[3][0], "")
	rv[seg0] = 0
	vr[0] = seg0

	// -of the two letters which occurs once in each of the 3 length 5 words ("2","5","3") the one of those which occurs in the length 4 word
	//  is segment 1 the other is segment 4 which means that the letter which does not appear in the length four word is segment 3
	occurrences := countOccur(wrds[5][0] + wrds[5][1] + wrds[5][2])
	seg1or4 := occurrences[1][0] + occurrences[1][1]
	seg1 := string(regexp.MustCompile("[" + seg1or4 + "]").Find([]byte(wrds[4][0])))
	seg4 := string(regexp.MustCompile("["+seg1+"]").ReplaceAllString(seg1or4, ""))
	// fmt.Printf("oif: %v from: %v = %v\n", seg1or4, wrds[4][0], seg1)
	rv[seg1] = 1
	vr[1] = seg1
	rv[seg4] = 4
	vr[4] = seg4

	// 	which means that the letter which does not appear in the length two or once In Five   is segment 3

	seg1 = regexp.MustCompile("["+seg1or4+wrds[2][0]+"]").ReplaceAllString(wrds[4][0], "")
	rv[seg1] = 3
	vr[3] = seg1

	//	segment 6 is the one that doesnt have wl2 and the known segments
	knownLetters := ""
	knownSegments := ""
	for i, k := range rv {
		knownSegments += strconv.Itoa(k)
		knownLetters += string(i)
	}

	seg6 := (regexp.MustCompile("["+knownLetters+wrds[2][0]+"]").ReplaceAllString(wrds[6][0]+wrds[6][1]+wrds[6][2], ""))
	seg6 = string(seg6[0])

	rv[seg6] = 6
	vr[6] = seg6

	knownLetters += seg6
	knownSegments += "6"

	// we know segments 0,1,3,4,6 at this point

	// Segment 5 is the six letter word that only has one of the segments in "1" in it
	seg5 := ""
	for i := 0; i < 3; i++ {
		if seg5matches := regexp.MustCompile("["+wrds[2][0]+"]").FindAll([]byte(wrds[6][i]), -1); len(seg5matches) == 1 {
			seg5 = string(seg5matches[0])
		}

	}
	rv[seg5] = 5
	vr[5] = seg5

	// segment2 is the 'other' segment in "1" not counting segment5
	seg2 := string(regexp.MustCompile("["+seg5+"]").ReplaceAllString(wrds[2][0], ""))
	seg2 = string(seg2[0])

	rv[seg2] = 2
	vr[2] = seg2

	return rv, vr
}

func countOccur(s string) map[int][]string {
	rv := map[int][]string{}
	for i := 0; i < 7; i++ {
		letter := string(int([]byte("a")[0]) + i)
		occur := len(regexp.MustCompile("["+letter+"]").FindAllIndex([]byte(s), -1))
		rv[occur] = append(rv[occur], letter)
	}

	return rv

}

func reverseDigits(d []int) int {
	for i, v := range digits {

		if Equal(v, d) {
			return i
		}
	}
	return -1
}

// from https://yourbasic.org/golang/compare-slices/
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
