package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Rucksack struct {
	Histogram map[rune]int
}

func NewRucksack(s string) (*Rucksack, error) {
	h := make(map[rune]int)

	for _, value := range s {
		h[value] += 1
	}

	return &Rucksack{
		Histogram: h,
	}, nil
}

type ElfGroup struct {
	badge rune
}

func NewElfGroup(r1, r2, r3 *Rucksack) (*ElfGroup, error) {
	b, err := findBadge(r1, r2, r3)
	if err != nil {
		return nil, err
	}

	return &ElfGroup{
		badge: *b,
	}, nil
}

// use rune value against ascii table to determine value
func (e *ElfGroup) ValueOfBadge() (int, error) {
	val := int(e.badge)

	// a = 1 ... z = 26
	if val >= 'a' && val <= 'z' {
		return val - 'a' + 1, nil
	}
	// a = 1 ... z = 27
	if val >= 'A' && val <= 'Z' {
		return val - 'A' + 27, nil
	}

	return 0, nil
}

func keysFromMap(m map[rune]int) (ans []rune) {
	ans = make([]rune, len(m))
	i := 0
	for k := range m {
		ans[i] = k
		i++
	}
	return ans
}

func findBadge(r1, r2, r3 *Rucksack) (*rune, error) {
	r1keys := keysFromMap(r1.Histogram)

	for _, k := range r1keys {
		_, r2ok := r2.Histogram[k]
		_, r3ok := r3.Histogram[k]
		if r2ok && r3ok {
			return &k, nil
		}
	}

	return nil, fmt.Errorf("no matching badge in this elf group")
}

func main() {
	//read from stdin, so we can pipe it in
	scanner := bufio.NewScanner(os.Stdin)

	var total int

	for scanner.Scan() {
		r1, err := NewRucksack(strings.TrimSpace(scanner.Text()))
		if err != nil {
			log.Fatal(err)
		}

		scanner.Scan()
		r2, err := NewRucksack(strings.TrimSpace(scanner.Text()))
		if err != nil {
			log.Fatal(err)
		}

		scanner.Scan()
		r3, err := NewRucksack(strings.TrimSpace(scanner.Text()))
		if err != nil {
			log.Fatal(err)
		}

		g, err := NewElfGroup(r1, r2, r3)
		if err != nil {
			log.Fatal(err)
		}

		v, err := g.ValueOfBadge()
		if err != nil {
			log.Fatal(err)
		}

		total += v
	}
	fmt.Printf("%d\n", total)
}
