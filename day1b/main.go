package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type TopThree struct {
	Elves []uint64
}

func NewTopThree() *TopThree {
	// store top three, and one more for ranking a new elf
	return &TopThree{Elves: make([]uint64, 4, 4)}
}

func (t *TopThree) RankNewElf(elf uint64) {
	// add/replace this elf as the last element of array
	t.Elves[len(t.Elves)-1] = elf

	sort.Slice(t.Elves[:], func(i, j int) bool { return t.Elves[i] > t.Elves[j] })
}

func (t *TopThree) AllCalories() uint64 {
	var sum uint64
	for _, elf := range t.Elves[0:3] {
		sum += elf
	}
	return sum
}

func main() {
	//read from stdin, so we can pipe it in
	scanner := bufio.NewScanner(os.Stdin)

	f := NewTopThree()
	var thisElf uint64

	for scanner.Scan() {
		t := strings.TrimSpace(scanner.Text())
		if t == "" {
			f.RankNewElf(thisElf)
			thisElf = 0
		} else {
			itemCalories, err := strconv.ParseUint(t, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			thisElf += itemCalories
		}
	}

	fmt.Printf("Top three elves are carrying %v calories\n", f.AllCalories())
}
