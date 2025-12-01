package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("D1")
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File read failed")
		return
	}

	contentStr := string(bytes.TrimSpace(content))
	lines := strings.Split(contentStr, "\n")

	dial := 50
	zeros_count := 0
	for _, line := range lines {
		direction := line[0]
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Printf("Line not convertible %v to int", err)
		}

		if direction == 'L' {
			for range steps {
				dial -= 1
				dial %= 100
				if dial == 0 {
					zeros_count++
				}
			}
		} else {
			for range steps {
				dial += 1
				dial %= 100
				if dial == 0 {
					zeros_count++
				}
			}
		}
		fmt.Printf("Line: %v, dial: %v, zeros_count: %v\n", line, dial, zeros_count)
	}
	fmt.Println(zeros_count)
}
