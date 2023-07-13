package main

import (
	"bufio"
	"errors"
	"log"
	"os"
)

// evaluateScore takes a match string in the form of "A X" and returns the points
// earned from the match by the player in the second column
func evaluateScore(match string) (int, error) {
    if len(match) != 3 || string(match[1]) != " " {
        return 0, errors.New("Incorrect match string format")
    }

    win, loss, tie := 6, 0, 3
    rock, paper, scissors := 1, 2, 3
    
    opponent := string(match[0])
    player := string(match[2])

    if opponent == "A" {
        switch player {
        case "X":
            return scissors + loss, nil
        case "Y":
            return rock + tie, nil
        case "Z":
            return paper +  win, nil
        }
    } else if opponent == "B" {
        switch player {
        case "X":
            return rock + loss, nil
        case "Y":
            return paper + tie, nil
        case "Z":
            return scissors + win, nil
        }
    } else {
        switch player {
        case "X":
            return paper + loss, nil
        case "Y":
            return scissors + tie, nil
        case "Z":
            return rock + win, nil
        }
    }

    return 0, errors.New("Incorrect match string format")


}

func main() {
    args := os.Args[1:]

    if len(args) == 0 {
        log.Fatal("Please provide a filename")
    }

    file, err := os.Open(args[0])
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)

    points := 0

    for scanner.Scan() {
        match := scanner.Text()

        res, err := evaluateScore(match)
        if err != nil {
            log.Fatal(err)
        }

        points += res
    }

    log.Printf("The total number of points you can get is %d", points)

}
