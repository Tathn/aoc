package aoc

import (
	_ "embed"
	"log"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func isSafe(levels []int) bool {
    isIncreasing := levels[0] < levels[1]
    for index, level := range levels[1:] {
        if levels[index] == level {
            return false
        }

        if (levels[index] < level) != isIncreasing {
            return false
        }

        diff := levels[index] - level
        abs := max(diff, -diff)
        if abs < 1 || abs > 3 {
            return false
        }
    }
    return true
}

func isSafePart2(levels []int) bool {
    isIncreasing := levels[0] < levels[1]
    for index, level := range levels[1:] {


        if levels[index] == level {
            return false
        }

        if (levels[index] < level) != isIncreasing {
            return false
        }

        diff := levels[index] - level
        abs := max(diff, -diff)
        if abs < 1 || abs > 3 {
            return false
        }
    }

    return true
}

func Main() {
    lines := strings.Split(input, "\n")
    reports := lines[:len(lines)-1]
    safeReports := 0
    for _, report := range reports {

        levels := make([]int, 0)
        for _, num := range strings.Fields(report) {
            inum, err := strconv.Atoi(num)
            if err != nil {
                panic(err)
            }
            levels = append(levels, inum)
        }

        safe := isSafePart2(levels)

        if !safe {
            for i := range len(levels) {
                sl := make([]int, len(levels))
                copy(sl, levels)
                sl = slices.Delete(sl, i, i + 1)
                safe = isSafePart2(sl)
                if safe {
                    break
                }
            }
        }

        if safe {
            safeReports++
        }
    }

    log.Print("Safe reports: ", safeReports)
}
