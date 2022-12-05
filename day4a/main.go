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

	var total int

	for scanner.Scan() {
		t := strings.TrimSpace(scanner.Text())

		r, err := NewSectionAssignment(t)
		if err != nil {
			log.Fatal(err)
		}

		if r.OneContainsTheOther() {
			total += 1
		}
	}
	fmt.Printf("%d\n", total)
}

type SectionAssignment struct {
	Afirst int
	Alast  int
	Bfirst int
	Blast  int
}

func NewSectionAssignment(s string) (*SectionAssignment, error) {
	sections := strings.Split(s, ",")

	a := strings.Split(sections[0], "-")
	b := strings.Split(sections[1], "-")

	aFirst, err := strconv.Atoi(a[0])
	if err != nil {
		return nil, err
	}

	aLast, err := strconv.Atoi(a[1])
	if err != nil {
		return nil, err
	}

	bFirst, err := strconv.Atoi(b[0])
	if err != nil {
		return nil, err
	}

	bLast, err := strconv.Atoi(b[1])
	if err != nil {
		return nil, err
	}

	return &SectionAssignment{
		Afirst: aFirst,
		Alast:  aLast,
		Bfirst: bFirst,
		Blast:  bLast,
	}, nil
}

func (s *SectionAssignment) OneContainsTheOther() bool {
	// check if B is within A
	if s.Bfirst >= s.Afirst && s.Blast <= s.Alast {
		return true
	}

	// check if A is within B
	if s.Afirst >= s.Bfirst && s.Alast <= s.Blast {
		return true
	}
	return false
}
