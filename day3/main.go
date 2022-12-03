package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Rucksack struct {
	comp1         *Compartment
	comp2         *Compartment
	misplacedItem string
}

func NewRucksack(s string) (*Rucksack, error) {
	split := strings.SplitAfter(s, "")

	c1 := NewCompartment(split[0 : len(split)/2])
	c2 := NewCompartment(split[len(split)/2:])

	m, err := findMisplaced(c1, c2)
	if err != nil {
		return nil, err
	}

	return &Rucksack{
		comp1:         c1,
		comp2:         c2,
		misplacedItem: *m,
	}, nil
}

// use rune value against ascii table to determine value
func (r *Rucksack) ValueOfMisplacedItem() (int, error) {
	val := int([]rune(r.misplacedItem)[0])

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

func keysFromMap(m map[string]int) (ans []string) {
	ans = make([]string, len(m))
	i := 0
	for k := range m {
		ans[i] = k
		i++
	}
	return ans
}

func findMisplaced(c1, c2 *Compartment) (*string, error) {
	c1keys := keysFromMap(c1.Histogram)

	// exactly one misplaced, so we only have to
	// look for matching keys in c2 (and not vice versa)
	for _, k := range c1keys {
		if _, ok := c2.Histogram[k]; ok {
			return &k, nil
		}
	}

	return nil, fmt.Errorf("no matching keys in compartments")
}

type Compartment struct {
	Histogram map[string]int
}

func NewCompartment(s []string) *Compartment {
	h := make(map[string]int)

	for _, value := range s {
		h[value] += 1
	}

	return &Compartment{
		Histogram: h,
	}
}

func main() {
	//read from stdin, so we can pipe it in
	scanner := bufio.NewScanner(os.Stdin)

	var total int

	for scanner.Scan() {
		t := strings.TrimSpace(scanner.Text())

		r, err := NewRucksack(t)
		if err != nil {
			log.Fatal(err)
		}

		v, err := r.ValueOfMisplacedItem()
		if err != nil {
			log.Fatal(err)
		}

		total += v
	}
	fmt.Printf("%d\n", total)
}
