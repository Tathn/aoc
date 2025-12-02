package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Range struct {
	Low int
	High int
}

func isValid(num string) bool {
	if len(num) % 2 != 0 {
		return true
	}

	iHigher := len(num) / 2
	iLowerEnd := iHigher - 1

	for iLower, iHigherEnd := 0, len(num) ; iHigher < iHigherEnd && iLower <= iLowerEnd; iHigher, iLower = iHigher + 1, iLower + 1 {
		if num[iHigher] != num[iLower] {
			return true
		}
	}

	return false
}

func isValid2(num string) bool {
	seq := []byte{num[0]}
	for _, chr := range num[1:] {
		if byte(chr) == seq[0] {
			break
		}
		seq = append(seq, byte(chr))
	}
	c := strings.Count(num, string(seq))
	if c < 2 {
		return true
	}
	return len(num) != len(seq) * c
}

func getInvalidIds(r Range) []int {
	if r.High == r.Low {
		return []int{}
	}

	result := []int{}
	for i := r.Low; i <= r.High; i++ {
		if !isValid2(strconv.Itoa(i)) || !isValid(strconv.Itoa(i)) {
			fmt.Println(i, " is invalid")
			result = append(result, i)
		}
	}
	return result
}

func main() {
	content, err := os.ReadFile("input_test.txt")
	if err != nil {
		fmt.Println("Could not read file")
		return
	}
	content = bytes.TrimSpace(content)
	contentStr := string(content)
	entries := []Range{}
	for entry := range strings.SplitSeq(contentStr, ",") {
		splitted := strings.Split(entry, "-")
		low, err := strconv.Atoi(splitted[0])
		if err != nil {
			fmt.Printf("Could not convert %v to int", splitted[0])
			return
		}
		high, err := strconv.Atoi(splitted[1])
		if err != nil {
			fmt.Printf("Could not convert %v to int", splitted[1])
			return
		}
		entries = append(entries, Range{Low: low, High: high})
	}

	result := []int{}
	for _, entry := range entries {
		result = append(result, getInvalidIds(entry)...)
	}

	sum := 0
	for _, r := range result {
		sum += r
	}

	fmt.Println("Sum: ", sum)
}
