package main

import (
	"bytes"
	"log"
	"os"
	"slices"
	"strconv"
	"time"
)

func getMaxJoltage(bank []byte) int {
	maxJoltages := []byte{0}
	for idx, battery := range bank {
		leftInBank := len(bank) - 1 - idx
		for i := len(maxJoltages) - 1; i >= 0; i-- {
			if len(maxJoltages) + leftInBank < 12 {
				break
			}

			if battery > maxJoltages[i] {
				maxJoltages = slices.Delete(maxJoltages, i, i + 1)
			}
		}

		if len(maxJoltages) < 12 {
			maxJoltages = append(maxJoltages, battery)
		}
	}

	res, err := strconv.Atoi(string(maxJoltages))
	if err != nil {
		log.Panicf("Failed to convert %v :<", maxJoltages)
	}
	return res
}

func main() {
	start := time.Now()
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Println("Could not open file")
		return
	}

	lines := bytes.SplitSeq(bytes.TrimSpace(content), []byte{'\n'})
	sumJoltages := 0
	for line := range lines {
		sumJoltages += getMaxJoltage(line)
	}
	log.Println("res:", sumJoltages)
	log.Printf("time: %v", time.Since(start))
}
