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

func parseRules(rules []string) map[int][]int {
    parsedRules := make(map[int][]int)
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

        parsedRules[yInt] = append(parsedRules[yInt], xInt)
    }
    return parsedRules
}

func commaSeparatedStringToIntSlice(str string) []int {
    splitted := strings.Split(str, ",")
    slice := make([]int, len(splitted))
    for index, entry := range splitted {
        asInt, err := strconv.Atoi(entry)
        if err != nil {
            panic(err)
        }

        slice[index] = asInt
    }
    return slice
}

func splitUpdatedRulesOrder(updates []string, rules map[int][]int) ([][]int,[][]int) {
    correctlyOrderedUpdates := make([][]int, 0)
    incorrectlyOrderedUpdates := make([][]int, 0)
    for _, update := range updates {
        pages := commaSeparatedStringToIntSlice(update)

        isUpdateCorrectlyOrdered := true
        for pageIndex, page := range pages {
            if ruleEntries, exists := rules[page]; exists {
                for _, entry := range ruleEntries {
                    if entryIndex := slices.Index(pages, entry); entryIndex != -1 {
                        if entryIndex > pageIndex {
                            isUpdateCorrectlyOrdered = false
                            break
                        }
                    }
                }
            }
            if !isUpdateCorrectlyOrdered {
                break
            }
        }
        if isUpdateCorrectlyOrdered {
            correctlyOrderedUpdates = append(correctlyOrderedUpdates, pages)
            continue
        } else {
            incorrectlyOrderedUpdates = append(incorrectlyOrderedUpdates, pages)
        }
    }

    return correctlyOrderedUpdates, incorrectlyOrderedUpdates
}

func fixPagesOrder(updates [][]int, rules map[int][]int) [][]int {
    fixedUpdates := make([][]int, 0)
    for _, update := range updates {
        isUpdateFixed := false
        for !isUpdateFixed {
 //           updateCpy := make([]int, len(update))
            for pageIndex, page := range update {
                isPageFixed := true
                if ruleEntries, exists := rules[page]; exists {
                    for _, entry := range ruleEntries {
                        if entryIndex := slices.Index(update, entry); entryIndex != -1 {
                            if entryIndex > pageIndex {
                                log.Println("Swapping",update[entryIndex], "with", update[pageIndex], "in", update)
                                tmp := update[entryIndex]
                                update[entryIndex] = update[pageIndex]
                                update[pageIndex] = tmp
                                log.Println("Result", update)
                                isPageFixed = false
                                isUpdateFixed = false
                                break
                            }
                        }
                    }
                }
                isUpdateFixed = isPageFixed
                if !isUpdateFixed {
                    break
                }
            }
        }
        fixedUpdates = append(fixedUpdates, update)
    }
    return fixedUpdates
}

func Main() {
    lines := strings.Split(input, "\n")
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
    correctlyOrderedUpdates, incorrectlyOrderedUpdates := splitUpdatedRulesOrder(updates, rules)
    log.Println("rules:", rules)
    log.Println("Correct:",correctlyOrderedUpdates)
    log.Println("Incorrect:", incorrectlyOrderedUpdates)

    partOneResult = 0
    for _, update := range correctlyOrderedUpdates {
        mid := update[len(update)/2]
        partOneResult += mid
    }
    fixedUpdates := fixPagesOrder(incorrectlyOrderedUpdates, rules)
    log.Println("Fixed:", fixedUpdates)
    partTwoResult = 0
    for _, update := range fixedUpdates {
        mid := update[len(update)/2]
        partTwoResult += mid
    }
    log.Println("=====================\nPart One:", partOneResult, "\nPart Two:", partTwoResult)
}
