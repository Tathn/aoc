package main

import (
	"bytes"
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Panicln("Could not read input")
	}
	papers := bytes.Split(bytes.TrimSpace(content), []byte{'\n'})
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
				}
			}
		}
	}

	log.Println("forkliftAccesableCount", forkliftAccesableCount)
}
