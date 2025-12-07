package main

import (
	"bytes"
	"log"
	"os"
	"slices"
)

func p1(papers [][]byte) int {
	return getAccessiblePapers(papers, false)
}

func getAccessiblePapers(papers [][]byte, remove bool) int {
	forkliftAccesableCount := 0
	for y, row := range papers {
		for x, paper := range row {
			if paper == '@' {
				adjacentCount := 0
				if y != 0 {
					if papers[y-1][x] == '@' {
						adjacentCount++
					}
					if x != 0 && papers[y-1][x-1] == '@' {
						adjacentCount++
					}
					if x != len(row) - 1 && papers[y-1][x+1] == '@' {
						adjacentCount++
					}
				}

				if x != 0 && papers[y][x-1] == '@' {
					adjacentCount++
				}
				if x != len(row) - 1 && papers[y][x+1] == '@' {
					adjacentCount++
				}

				if y < len(papers) - 1 {
					if papers[y+1][x] == '@' {
						adjacentCount++
					}
					if x != 0 && papers[y+1][x-1] == '@' {
						adjacentCount++
					}
					if x != len(row) - 1 && papers[y+1][x+1] == '@' {
						adjacentCount++
					}
				}

				if adjacentCount < 4 {
					forkliftAccesableCount++
					if remove {
						papers[y][x] = '.'
					}
				}
			}
		}
	}
	return forkliftAccesableCount
}

func p2(papers [][]byte) int {
	removedPapersCount := 0
	lastRoundRemovedPapers := -1
	for lastRoundRemovedPapers != 0 {
		lastRoundRemovedPapers = getAccessiblePapers(papers, true)
		removedPapersCount += lastRoundRemovedPapers
	}

	return removedPapersCount
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Panicln("Could not read input")
	}
	papers := bytes.Split(bytes.TrimSpace(content), []byte{'\n'})

	paperscpy := make([][]byte, len(papers))
	for idx, row := range papers {
		paperscpy[idx] = slices.Clone(row)
	}

	log.Println("forkliftAccesableCount", p1(paperscpy))
	log.Println("removedPapers", p2(papers))
}
