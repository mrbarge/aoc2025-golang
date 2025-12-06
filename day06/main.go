package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mrbarge/aoc2025-golang/helper"
)

type Sum struct {
	Nums     []string
	Operator string
}

type OperatorRange struct {
	Lower    int
	Upper    int
	Operator string
}

func calcTwo(lines []string, operatorRanges []OperatorRange) (r int) {
	for _, o := range operatorRanges {
		nums := make([]string, 0)
		for i := o.Upper; i >= o.Lower; i-- {
			number := ""
			for j := 0; j < len(lines); j++ {
				if i < len(lines[j]) && lines[j][i] != ' ' {
					number += string(lines[j][i])
				}
			}
			nums = append(nums, number)
		}
		r += sumdigits(nums, o.Operator)
	}
	return r
}

func operatorRange(s string) (r []OperatorRange) {
	r = make([]OperatorRange, 0)
	start := 0
	end := 0
	operator := string(s[0])
	for i := 1; i < len(s); i++ {
		if s[i] != ' ' {
			end = i - 2
			r = append(r, OperatorRange{start, end, operator})
			start = i
			operator = string(s[i])
		} else {
			end = i
		}
	}
	r = append(r, OperatorRange{start, end, operator})
	return r
}

func sumdigits(nums []string, operator string) int {
	ans, _ := strconv.Atoi(nums[0])
	for i := 1; i < len(nums); i++ {
		inum, _ := strconv.Atoi(nums[i])
		if operator == "*" {
			ans *= inum
		} else if operator == "+" {
			ans += inum
		} else if operator == "-" {
			ans -= inum
		}
	}
	return ans
}

func (s Sum) Operate() int {
	return sumdigits(s.Nums, s.Operator)
}

func readData(lines []string) []Sum {
	sums := make([]Sum, 0)

	ilines := make([][]string, 0)
	for rowNum, line := range lines[:len(lines)-1] {
		ilines = append(ilines, make([]string, 0))
		for _, s := range strings.Fields(line) {
			ilines[rowNum] = append(ilines[rowNum], s)
		}
	}
	for rowNum, operator := range strings.Fields(lines[len(lines)-1]) {
		sum := Sum{Nums: make([]string, 0), Operator: operator}
		for _, iline := range ilines {
			sum.Nums = append(sum.Nums, iline[rowNum])
		}
		sums = append(sums, sum)
	}
	return sums
}

func partOne(lines []string) (r int, err error) {
	sums := readData(lines)
	for _, sum := range sums {
		r += sum.Operate()
	}
	return r, err
}

func partTwo(lines []string) (r int, err error) {
	operatorRanges := operatorRange(lines[len(lines)-1])
	r = calcTwo(lines[:len(lines)-1], operatorRanges)
	return r, err
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
