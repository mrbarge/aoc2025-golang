package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/mrbarge/aoc2025-golang/helper"
)

type IngredientRange struct {
	Lower int
	Upper int
}

type Ingredient int

func readData(lines []string) (r []IngredientRange, ingredients []int) {
	r = make([]IngredientRange, 0)
	ingredients = make([]int, 0)

	for _, line := range lines {
		if strings.Index(line, "-") >= 0 {
			ranges := strings.Split(line, "-")
			lower, _ := strconv.Atoi(ranges[0])
			upper, _ := strconv.Atoi(ranges[1])
			r = append(r, IngredientRange{lower, upper})
		} else {
			i, _ := strconv.Atoi(line)
			ingredients = append(ingredients, i)
		}
	}
	return r, ingredients
}

func expandRanges(ranges []IngredientRange) []IngredientRange {
	if len(ranges) == 0 {
		return ranges
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Lower < ranges[j].Lower
	})

	expandedRanges := []IngredientRange{ranges[0]}

	for i := 1; i < len(ranges); i++ {
		current := ranges[i]
		last := &expandedRanges[len(expandedRanges)-1]

		if current.Lower <= last.Upper+1 {
			if current.Upper > last.Upper {
				last.Upper = current.Upper
			}
		} else {
			expandedRanges = append(expandedRanges, current)
		}
	}

	return expandedRanges
}

func partOne(lines []string) (r int, err error) {
	ranges, ingredients := readData(lines)
	for _, ingredient := range ingredients {
		for _, rng := range ranges {
			if ingredient >= rng.Lower && ingredient <= rng.Upper {
				r++
				break
			}
		}
	}
	return r, err
}

func partTwo(lines []string) (r int, err error) {
	ranges, _ := readData(lines)

	expandedRanges := expandRanges(ranges)
	for _, erng := range expandedRanges {
		r += erng.Upper - erng.Lower + 1
	}
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
