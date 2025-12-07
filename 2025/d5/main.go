package main

import (
	"bytes"
	"log"
	"os"
	"slices"
	"strconv"
)

type Range struct {
	Min int
	Max int
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Panicln("Couldn't read file")
	}

	in := bytes.Split(bytes.TrimSpace(content), []byte{'\n'})
	freshRanges := []Range{}
	productIds := []int{}

	parsingRanges := true
	for _, item := range in {
		if len(item) == 0 {
			parsingRanges = false
			continue
		}

		if parsingRanges {
			r := bytes.Split(item, []byte{'-'})
			minRange, err := strconv.Atoi(string(r[0]))
			if err != nil {
				log.Panicln(err)
			}
			maxRange, err := strconv.Atoi(string(r[1]))
			if err != nil {
				log.Panicln(err)
			}
			freshRanges = append(freshRanges, Range{Min: minRange, Max: maxRange})
		} else {
			id, err := strconv.Atoi(string(item))
			if err != nil {
				log.Panicln(err)
			}
			productIds = append(productIds, id)
		}
	}

	slices.SortFunc(freshRanges, func(a Range, b Range) int {
		if a.Min < b.Min {
			return -1
		}
		if a.Min > b.Min {
			return 1
		}
		return 0
	})

	reducedFreshRanges := []Range{freshRanges[0]}
	for _, ran := range freshRanges[1:] {
		reducedFreshRangesLen := len(reducedFreshRanges)
		prevRange := reducedFreshRanges[reducedFreshRangesLen - 1]
		if prevRange.Min <= ran.Max && ran.Min <= prevRange.Max {
			reducedFreshRanges = slices.Delete(reducedFreshRanges, reducedFreshRangesLen - 1, reducedFreshRangesLen)
			reducedFreshRanges = append(reducedFreshRanges, Range{Min: min(prevRange.Min, ran.Min), Max: max(prevRange.Max, ran.Max)})
		} else {
			reducedFreshRanges = append(reducedFreshRanges, ran)
		}
	}

	sum := 0
	for _, id := range productIds {
		for _, ran := range reducedFreshRanges {
			if id >= ran.Min && id <= ran.Max {
				sum++
				break
			}
		}
	}
	log.Println("Sum p1", sum)

	sumFreshIds := 0
	for _, ran := range reducedFreshRanges {
		sumFreshIds += ran.Max - ran.Min + 1
	}
	log.Println("Sum p2", sumFreshIds)
}
