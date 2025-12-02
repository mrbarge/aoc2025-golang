package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mrbarge/aoc2025-golang/helper"
)

func isInvalid(i int) bool {
	si := strconv.Itoa(i)
	if len(si)%2 != 0 {
		return false
	}
	return si[:len(si)/2] == si[len(si)/2:]
}

func isInvalidTwo(i int) bool {
	si := strconv.Itoa(i)
	for ix := 1; ix <= len(si)/2; ix++ {
		if len(si)%ix != 0 {
			continue
		}
		ss := si[:ix]
		rs := ""
		for ij := 0; ij < len(si)/ix; ij++ {
			rs += ss
		}
		if rs == si {
			return true
		}
	}
	return false
}

func getRange(s string) (l int, r int) {
	c := strings.Index(s, "-")
	l, _ = strconv.Atoi(s[:c])
	r, _ = strconv.Atoi(s[c+1:])
	return l, r
}

func partOne(line string) (r int, err error) {
	ranges := strings.Split(line, ",")
	for _, rn := range ranges {
		from, to := getRange(rn)
		for i := from; i <= to; i++ {
			if isInvalid(i) {
				r += i
			}
		}
	}
	return r, nil
}

func partTwo(line string) (r int, err error) {
	ranges := strings.Split(line, ",")
	for _, rn := range ranges {
		from, to := getRange(rn)
		for i := from; i <= to; i++ {
			if isInvalidTwo(i) {
				fmt.Printf("invalid: %d\n", i)
				r += i
			}
		}
	}
	return r, nil
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := partOne(lines[0])
	fmt.Printf("Part one: %d\n", ans)

	ans, err = partTwo(lines[0])
	fmt.Printf("Part two: %d\n", ans)

}
