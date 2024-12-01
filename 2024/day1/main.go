package aoc

import (
	_ "embed"
	"log"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func splitColumns(lines []string) (sort.IntSlice, sort.IntSlice) {
    firstList := make([]int, len(lines)/2)
    secondList := make([]int, len(lines)/2)
    for _, line := range lines {
        columns := strings.Fields(line)

        firstVal, err := strconv.Atoi(columns[0])
        if err != nil {
            log.Fatal(err)
        }
        firstList = append(firstList, firstVal)

        secondVal, err := strconv.Atoi(columns[1])
        if err != nil {
            log.Fatal(err)
        }
        secondList = append(secondList, secondVal)
    }
    return firstList, secondList
}

func countInList(num int, list []int) int {
    numOccurences := 0
    for _, currentNum := range list {
        if currentNum == num {
            numOccurences++
        }
    }
    return numOccurences
}

func Main() {
    lines := strings.Split(input, "\n")
    lines = lines[:len(lines)-1]
    firstList, secondList := splitColumns(lines)
    sort.Sort(firstList)
    sort.Sort(secondList)

    //distances := make([]int, len(firstList))
    sum := 0
    for index, num := range firstList {
        diff := num - secondList[index]
        sum += max(diff, -diff)
    }
    log.Printf("Distance sum: %d", sum)

    //part 2
    numberToSScore := make(map[int]int)
    similiarityScore := 0
    for _, num := range firstList {
        if score, ok := numberToSScore[num]; ok {
            similiarityScore += score
            continue
        }

        score := num * countInList(num, secondList)
        numberToSScore[num] = score
        similiarityScore += score
    }
    log.Printf("Similiarity score: %d", similiarityScore)
}
