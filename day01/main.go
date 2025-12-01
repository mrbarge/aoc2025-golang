package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mrbarge/aoc2025-golang/helper"
)

// euclidean mod
func mod(n, m int) int {
	return ((n % m) + m) % m
}

func zerocross(pos int, spin int) int {
	if spin >= 0 {
		return (pos + spin) / 100
	} else {
		pspin := -spin
		if pos == 0 {
			return pspin / 100
		} else if pspin >= pos {
			return (pspin-pos)/100 + 1
		} else {
			return 0
		}
	}
}

func rotate(pos int, cmd string) (r int, zeroCrosses int, err error) {
	spin, _ := strconv.Atoi(cmd[1:])
	if cmd[0] == 'L' {
		spin = -spin
	}
	zeroCrosses = zerocross(pos, spin)
	return mod(pos+spin, 100), zeroCrosses, nil
}

func partOne(lines []string) (r int, err error) {
	pos := 50
	for _, line := range lines {
		pos, _, err = rotate(pos, line)
		if pos == 0 {
			r++
		}
	}
	return r, nil
}

func partTwo(lines []string) (r int, err error) {
	pos := 50
	zc := 0
	for _, line := range lines {
		pos, zc, _ = rotate(pos, line)
		r += zc
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

	ans, err := partOne(lines)
	fmt.Printf("Part one: %d\n", ans)

	ans, err = partTwo(lines)
	fmt.Printf("Part two: %d\n", ans)

}
