package aoc

import (
	_ "embed"
	"log"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed input_example.txt
var inputExample string

func Main() {
    lines := strings.Split(input, "\n")
    lines = lines[:len(lines)-1]

    operationMatcher := regexp.MustCompile("do\\(\\)|don't\\(\\)|mul\\(([0-9]{1,3}),([0-9]{1,3})\\)")
    result := 0

    shouldExecute := true
    for _, line := range lines {
        operations := operationMatcher.FindAllStringSubmatch(line, -1)
        for _, vals := range operations {
            if vals[0] == "do()" {
                shouldExecute = true
                continue
            } else if vals[0] == "don't()" {
                shouldExecute = false
                continue
            }

            if !shouldExecute {
                continue
            }

            v1, err1 := strconv.Atoi(vals[1])
            v2, err2 := strconv.Atoi(vals[2])
            if err1 != nil || err2 != nil {
                panic("Skibidi vals not convertable to sigma ints")
            }
            result += v1 * v2
        }
    }

    log.Print("Operations result: ", result)
}
