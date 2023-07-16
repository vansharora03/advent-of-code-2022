package main

import (
	"bufio"
	"errors"
	"log"
	"os"
)

// findCommonCharacter finds the common char in string slice
func findCommonCharacter(s []string) (rune, error) {
    if len(s) < 1 {
        return 'A', errors.New("Empty slice was passed")
    }

    sets := []map[rune]bool{}

    for _, str := range s {
        currSet := make(map[rune]bool)
        for _, char := range str {
            currSet[char] = true
        }
        sets = append(sets, currSet)
    }

    for _, char := range s[0] {
        count := 0
        for _, set := range sets {
            _, ok := set[char]
            if ok {
                count += 1
            }
        }
        
        if count == len(s) {
            return char, nil
        }
    }

    return 'A', errors.New("No common letter")
}

// priorityFromChar returns the number of points the given char nets
func priorityFromChar(char rune) int {
    if 'A' <= char && char <= 'Z' {
        return 27 + int(char - 'A')
    } else if 'a' <= char && char <= 'z' {
        return 1 + int(char - 'a')
    }
    return 0
}


func main() {
    args := os.Args[1:]
    if len(args) == 0 {
        log.Fatal("Please provide a file name")
    }

    file, err := os.Open(args[0])
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)

    totalPriority := 0

    currentGroup := []string{}

    for scanner.Scan() {
        rucksack := scanner.Text()

        currentGroup = append(currentGroup, rucksack)

        if len(currentGroup) == 3 {
            char, err := findCommonCharacter(currentGroup)
            if err != nil {
                log.Fatal(err)
            }

            totalPriority += priorityFromChar(char)
            currentGroup = []string{}
        }
    }

    log.Printf("The total priority is: %d", totalPriority)
}

