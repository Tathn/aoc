package aoc

import (
	_ "embed"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed input_example.txt
var inputExample string

func parseRules(rules []string) map[int]int {
    parsedRules := make(map[int]int)
    for _, rule := range rules {
        x, y, found := strings.Cut(rule, "|")
        if !found {
            panic(fmt.Sprintf("Could not parse rule %s", rule))
        }

        xInt, errX := strconv.Atoi(x)
        yInt, errY := strconv.Atoi(y)
        if errX != nil || errY != nil {
            log.Println(errX)
            log.Println(errY)
            panic(fmt.Sprintf("Could not convert rules to int, x=%s, y=%s", x, y))
        }

        parsedRules[yInt] = xInt
    }
    return parsedRules
}

func Main() {
    //TODO rules parsed wrong
    lines := strings.Split(inputExample, "\n")
    lines = lines[:len(lines)-1]
    partOneResult := 0
    partTwoResult := 0

    delimiterIndex := slices.Index(lines, "")
    pageOrderingRules := lines[:delimiterIndex]
    updates := lines[delimiterIndex + 1:]

    log.Println(pageOrderingRules)
    log.Println("====")
    log.Println(updates)

    rules := parseRules(pageOrderingRules)
    correctlyOrderedUpdates := make([][]int, 0)
    for _, update := range updates {
        ps := strings.Split(update, ",")
        pages := make([]int, len(ps))
        for index, page := range ps {
            p, err := strconv.Atoi(page)
            if err != nil {
                panic(err)
            }

            pages[index] = p
        }

        isUpdateCorrectlyOrdered := true
        for pageIndex, page := range pages {
            if requiredPage, exists := rules[page]; exists {
                if requiredPageIndex := slices.Index(pages, requiredPage); requiredPageIndex != -1 {
                    if requiredPageIndex > pageIndex {
                        isUpdateCorrectlyOrdered = false
                    }
                }
            }
        }
        if isUpdateCorrectlyOrdered {
            correctlyOrderedUpdates = append(correctlyOrderedUpdates, pages)
        }
    }
    log.Println(rules)
    log.Println(correctlyOrderedUpdates)

    sum := 0
    for _, update := range correctlyOrderedUpdates {
        mid := update[len(update)/2]
        sum += mid
    }
    log.Println(sum)
    log.Println("=====================\nPart One:", partOneResult, "\nPart Two:", partTwoResult)
}
