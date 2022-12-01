package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//read from stdin, so we can pipe it in
	scanner := bufio.NewScanner(os.Stdin)

	var mostCalories uint64
	var thisElf uint64

	for scanner.Scan() {
		t := strings.TrimSpace(scanner.Text())
		if t == "" {
			// thisElf has no more so if this was the largest, set mostCalories
			if thisElf > mostCalories {
				mostCalories = thisElf
			}
			thisElf = 0
		} else {
			itemCalories, err := strconv.ParseUint(t, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			thisElf += itemCalories
		}
	}

	fmt.Printf("Elf carrying most calories, is carrying %v calories\n", mostCalories)
}
