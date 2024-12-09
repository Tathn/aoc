package aoc

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

//go:embed input_example.txt
var inputExample string

func nextCombination(n int, c string) func() string {
    r := []rune(c)
    p := make([]rune, n)
    x := make([]int, len(p))
    return func() string {
        p := p[:len(x)]
        for i, xi := range x {
            p[i] = r[xi]
        }
        for i := len(x) - 1; i >= 0; i-- {
            x[i]++
            if x[i] < len(r) {
                break
            }
            x[i] = 0
            if i <= 0 {
                x = x[0:0]
                break
            }
        }
        return string(p)
    }
}

type Calculation struct {
    result int
    symbols []int
}

func timer() func() {
    start := time.Now()
    return func() {
        fmt.Printf("Program took %v\n", time.Since(start))
    }
}

func doOperation(x int, y int, op string) int {
    switch {
    case op == "*":
        return x * y
    case op == "+":
        return x + y
    case op == "|":
        res := fmt.Sprintf("%d%d", x, y)
        resInt, err := strconv.Atoi(res)
        if err != nil {
            panic(err)
        }
        return resInt
    }
    panic(fmt.Sprintf("Operation %s not permitted", op))
}

func Main() {
    defer timer()()
    lines := strings.Split(input, "\n")
    lines = lines[:len(lines)-1]
    partOneResult := 0
    partTwoResult := 0

    calculations := make([]Calculation, 0)
    for _, line := range lines {
        fields := strings.Fields(line)
        res := fields[0][:len(fields[0]) - 1]
        result, err := strconv.Atoi(res)
        if err != nil {
            panic(err)
        }

        fields = fields[1:]
        symbols := make([]int, 0)
        for _, field := range fields {
            fieldInt, err := strconv.Atoi(field)
            if err != nil {
                panic(err)
            }
            symbols = append(symbols, fieldInt)
        }

        calculations = append(calculations, Calculation{result, symbols})
    }

    for _, calc := range calculations {
        generate := nextCombination(len(calc.symbols) - 1, "*+|")
        for {
            combinations := generate()
            if len(combinations) == 0 {
                break
            }
            result := calc.symbols[0]
            for index, op := range combinations {
                opResult := doOperation(result, calc.symbols[index + 1], string(op))
                result = opResult
            }
            if result == calc.result {
                partOneResult += calc.result
                break
            }
        }
    }

    log.Println("\nPart One:", partOneResult)
    log.Println("\nPart Two:", partTwoResult)
}
