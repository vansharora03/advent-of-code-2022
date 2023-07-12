package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
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

    defer file.Close()

    scanner := bufio.NewScanner(file)
    
    currTotal := 0
    currMaxes := [3]int{}

    for scanner.Scan() {
        line := scanner.Text()

        if line == "" {
            currMaxes[0] = max(currMaxes[0], currTotal)
            sort.Ints(currMaxes[:])
            currTotal = 0 
        } else {
            cals, err := strconv.ParseInt(line, 10, 64)
            if err != nil {
                log.Fatal(err)
            }

            currTotal += int(cals)
        }
            
    }

    res := 0

    for _, max := range currMaxes {
        res += max
    }

    log.Printf("The most calories held by an elf is : %d", currMaxes[2])
    log.Printf("The calories held by the top 3 elves is : %d", res)

}
