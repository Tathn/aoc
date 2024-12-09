package aoc

import (
	_ "embed"
	"fmt"
	"log"
	"strings"
	"time"
)

//go:embed input.txt
var input string

//go:embed input_example.txt
var inputExample string

type Position struct {
    x int
    y int
}

type Guard struct {
    vector Position
    current Position
}

func (g *Guard) Rotate() {
    switch {
    case g.vector.y == -1 && g.vector.x == 0:
        g.vector.y = 0
        g.vector.x = 1
    case g.vector.y == 0 && g.vector.x == 1:
        g.vector.y = 1
        g.vector.x = 0
    case g.vector.y == 1 && g.vector.x == 0:
        g.vector.y = 0
        g.vector.x = -1
    case g.vector.y == 0 && g.vector.x == -1:
        g.vector.y = -1
        g.vector.x = 0
    }
}

func (g *Guard) Move(lab []string) {
    newX := g.current.x + g.vector.x
    newY := g.current.y + g.vector.y
    if string(lab[newY][newX]) == "#" {
        g.Rotate()
        return
    }
    tmp := []rune(lab[g.current.y])
    tmp[g.current.x] = rune('X')
    lab[g.current.y] = string(tmp)
    g.current.x = newX
    g.current.y = newY
}

func (g *Guard) MoveLines(lab []string) {
    newX := g.current.x + g.vector.x
    newY := g.current.y + g.vector.y
    if string(lab[newY][newX]) == "#" || string(lab[newY][newX]) == "O" {
        g.Rotate()
        return
    }
    tmp := []rune(lab[g.current.y])
    var symbol rune
    switch {
    case g.vector.x != 0:
        symbol = rune('-')
    case g.vector.y != 0:
        symbol = rune('|')
    }
    tmp[g.current.x] = symbol
    lab[g.current.y] = string(tmp)
    g.current.x = newX
    g.current.y = newY
}


func (g *Guard) IsMovingOutOfBound(lab []string) bool{
    newX := g.current.x + g.vector.x
    newY := g.current.y + g.vector.y
    return newX > len(lab[0]) - 1 ||
           newY > len(lab) - 1 ||
           newX < 0 ||
           newY < 0
}

func findGuard(lab []string) Guard {
    for y, row := range lab {
        for x, column := range row {
            if string(column) == "^" {
                return Guard{Position{0,-1},Position{x,y}}
            }
        }
    }
    panic("Guard not found in lab")
}

func walkTheLab(lab []string) {
    guard := findGuard(lab)
    for !guard.IsMovingOutOfBound(lab) {
        guard.Move(lab)
    }
}

func timer() func() {
    start := time.Now()
    return func() {
        fmt.Printf("Program took %v\n", time.Since(start))
    }
}

func doesGuardLoop(lab []string, c chan int) {
    guard := findGuard(lab)
    history := make(map[Position] int)
    history[guard.current]++
    for {
        if guard.IsMovingOutOfBound(lab) {
            c <- 0
            break
        }
        guard.MoveLines(lab)
        history[guard.current]++
        if history[guard.current] == 100 {
            c <- 1
            break
        }
    }
}

func Main() {
    defer timer()()
    lines := strings.Split(input, "\n")
    lines = lines[:len(lines)-1]
    partOneResult := 0
    partTwoResult := 0

    lab := make([]string, len(lines))
    copy(lab, lines)
    fmt.Println(strings.Join(lab, "\n"))
    walkTheLab(lab)
    visited := 1 //including staring position
    for _, column := range lab {
        visited += strings.Count(column, "X")
    }
    partOneResult = visited
    log.Println("\nPart One:", partOneResult)
    fmt.Println(strings.Join(lab, "\n"))

    obstructionPositions := make([]Position, 0)
    for y, str := range lines {
        for x, entry := range str {
            if entry == '.' {
                obstructionPositions = append(obstructionPositions, Position{x,y})
            }
        }
    }
    fmt.Println("Obstructions will be placed at:", obstructionPositions)
    guardStartingPos := findGuard(lines).current
    c := make(chan int)
    for _, pos := range obstructionPositions {
        if pos == guardStartingPos {
            continue
        }
        obstructedLab := make([]string, len(lines))
        copy(obstructedLab, lines)
        tmp := []rune(obstructedLab[pos.y])
        tmp[pos.x] = rune('O')
        obstructedLab[pos.y] = string(tmp)

        go doesGuardLoop(obstructedLab, c)
    }
    for range len(obstructionPositions) {
        partTwoResult += <-c
    }
    log.Println("\nPart Two:", partTwoResult)
}
