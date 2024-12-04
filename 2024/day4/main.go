package aoc

import (
	_ "embed"
	"log"
	"strings"
)

//go:embed input.txt
var input string

//go:embed input_example.txt
var inputExample string

type Letter struct {
    val string
    x int
    y int
}

type WordSearch = []string
type Word = []Letter

func isInBounds(letter Letter, wordSearch WordSearch) bool {
    return letter.y < len(wordSearch) && letter.y > -1 && letter.x > -1 && letter.x < len(wordSearch[0])
}

func getWordsPart1(letter Letter) []Word {
    horizontalWords := []Word{
        {{"M", letter.x - 1, letter.y}, {"A", letter.x - 2, letter.y}, {"S", letter.x - 3, letter.y}},
        {{"M", letter.x + 1, letter.y}, {"A", letter.x + 2, letter.y}, {"S", letter.x + 3, letter.y}},
    }

    verticalWords := []Word{
        {{"M", letter.x, letter.y + 1}, {"A", letter.x, letter.y + 2}, {"S", letter.x, letter.y + 3}},
        {{"M", letter.x, letter.y - 1}, {"A", letter.x, letter.y - 2}, {"S", letter.x, letter.y - 3}},
    }

    diagonalWords := []Word{
        {{"M", letter.x - 1, letter.y - 1}, {"A", letter.x - 2, letter.y - 2}, {"S", letter.x - 3, letter.y - 3}},
        {{"M", letter.x + 1, letter.y - 1}, {"A", letter.x + 2, letter.y - 2}, {"S", letter.x + 3, letter.y - 3}},
        {{"M", letter.x - 1, letter.y + 1}, {"A", letter.x - 2, letter.y + 2}, {"S", letter.x - 3, letter.y + 3}},
        {{"M", letter.x + 1, letter.y + 1}, {"A", letter.x + 2, letter.y + 2}, {"S", letter.x + 3, letter.y + 3}},
    }

    checkedWords := []Word{}
    checkedWords = append(checkedWords, horizontalWords...)
    checkedWords = append(checkedWords, verticalWords...)
    checkedWords = append(checkedWords, diagonalWords...)

    return checkedWords
}

func getWordsPart2(letter Letter) []Word {
    checkedWords := []Word{
        {{"M", letter.x - 1, letter.y - 1}, {"S", letter.x + 1, letter.y + 1}, {"M", letter.x + 1, letter.y - 1}, {"S", letter.x - 1, letter.y + 1}},
        {{"M", letter.x - 1, letter.y - 1}, {"S", letter.x + 1, letter.y + 1}, {"M", letter.x - 1, letter.y + 1}, {"S", letter.x + 1, letter.y - 1}},
        {{"M", letter.x - 1, letter.y + 1}, {"S", letter.x + 1, letter.y - 1}, {"M", letter.x + 1, letter.y + 1}, {"S", letter.x - 1, letter.y - 1}},
        {{"M", letter.x + 1, letter.y - 1}, {"S", letter.x - 1, letter.y + 1}, {"M", letter.x + 1, letter.y + 1}, {"S", letter.x - 1, letter.y - 1}},
    }
    return checkedWords
}

func countWords(wordSearch WordSearch, words []Word) int {
    count := 0
    for _, word := range words {
        foundWord := true
        for _, checkedLetter := range word {
            if !isInBounds(checkedLetter, wordSearch) {
                foundWord = false
                break
            }
            if checkedLetter.val != string(wordSearch[checkedLetter.y][checkedLetter.x]) {
                foundWord = false
            }
        }
        if foundWord {
            count++
        }
    }
    return count
}

func Main() {
    lines := strings.Split(input, "\n")
    var wordSearch WordSearch = lines[:len(lines)-1]
    partOneResult := 0
    partTwoResult := 0
    for y, line := range wordSearch {
        for x, val := range line {
            currentLetter := Letter{string(val), x, y}
            if currentLetter.val == "X" {
                partOneResult += countWords(wordSearch, getWordsPart1(currentLetter))
            }
            if currentLetter.val == "A" {
                partTwoResult += countWords(wordSearch, getWordsPart2(currentLetter))
            }
        }
    }

    log.Println("\nPart One:", partOneResult, "\nPart Two:", partTwoResult)
}
