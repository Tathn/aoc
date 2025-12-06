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
	m := slices.Max(bank)
	mIdx := bytes.IndexByte(bank, m)
	if mIdx == len(bank) - 1 {
		m = slices.Max(bank[:mIdx])
		mIdx = bytes.IndexByte(bank, m)
	}
	n := slices.Max(bank[mIdx + 1:])

	val, err := strconv.Atoi(string(m) + string(n))
	if err != nil {
		log.Println("Could not convert MaxJoltage")
	}
	return val

}

func getMaxJoltage12(bank []byte) int {
	m := slices.Max(bank[:len(bank) - 11])
	mIdx := bytes.IndexByte(bank, m)

	batteriesIdxs := []int{}
	rightBank := bank[mIdx + 1:]
	rightBankCopy := make([]byte, len(rightBank))
	copy(rightBankCopy, rightBank)
	for range 11 {
		var maxIdx int
		maxVal := byte(0)
		for idx := len(rightBankCopy) - 1; idx >= 0; idx-- {
			val := rightBankCopy[idx]
			if val > maxVal {
				maxVal = val
				maxIdx = idx
			}
		}
		batteriesIdxs = append(batteriesIdxs, maxIdx)
		rightBankCopy = slices.Replace(rightBankCopy, maxIdx, maxIdx + 1, strconv.Itoa(0)[0])
	}
	slices.Sort(batteriesIdxs)
	tail := []byte{}
	for _, idx := range batteriesIdxs {
		tail = append(tail, rightBank[idx])
	}

	val, err := strconv.Atoi(string(m) + string(tail))
	if err != nil {
		log.Println("Could not convert MaxJoltage")
	}
	log.Println("\nb", string(bank), "\nb", string(bank[:len(bank) - 11]))
	return val

}

func getMaxJoltage2(bank []byte) int {
	maxJolt := []byte{0}
	for idx, battery := range bank {
		if battery > maxJolt[len(maxJolt) - 1] {
			i := slices.IndexFunc(maxJolt, func(b byte) bool {
					return b < battery
			})
			if idx > len(bank) - 13 {
				i = min(i, idx % 12) // i = 1, idx = 5, lenbank = 13
			}
			maxJolt = maxJolt[i:]
			maxJolt = append(maxJolt, battery)
		}
	}

	val, err := strconv.Atoi(string(maxJolt))
	if err != nil {
		log.Println("Could not convert MaxJoltage")
	}
	return val
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
		sumJoltages += getMaxJoltage12(line)
	}
	log.Println("res:", sumJoltages)
	log.Printf("time: %v", time.Since(start))
}
