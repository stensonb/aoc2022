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

		// if r.OneContainsTheOther() {
		if r.Overlap() {
			total += 1
		}
	}
	fmt.Printf("%d\n", total)
}

type SectionRange struct {
	Begin int
	End   int
}

// 74-76
func NewSectionRange(s string) (*SectionRange, error) {
	a := strings.Split(s, "-")

	if len(a) != 2 {
		return nil, fmt.Errorf("invalid range: %s", a)
	}

	begin, err := strconv.Atoi(a[0])
	if err != nil {
		return nil, err
	}

	end, err := strconv.Atoi(a[1])
	if err != nil {
		return nil, err
	}

	return &SectionRange{
		Begin: begin,
		End:   end,
	}, nil
}

type SectionAssignment struct {
	SectionRanges []*SectionRange
}

// 74-76,34-39
func NewSectionAssignment(s string) (*SectionAssignment, error) {
	inputSplit := strings.Split(s, ",")

	if len(inputSplit) != 2 {
		return nil, fmt.Errorf("invalid section assignment: %s", inputSplit)
	}

	sectionRanges := make([]*SectionRange, len(inputSplit))

	for i := range inputSplit {
		sectionRange, err := NewSectionRange(inputSplit[i])
		if err != nil {
			return nil, err
		}
		sectionRanges[i] = sectionRange
	}

	// sort sectionRanges
	sort.Slice(sectionRanges, func(i, j int) bool { return sectionRanges[i].Begin < sectionRanges[j].Begin })

	return &SectionAssignment{
		SectionRanges: sectionRanges,
	}, nil
}

func (s *SectionAssignment) OneContainsTheOther() bool {
	// only two section ranges
	// sectionRanges sorted by beginning of each range

	// if they both start with the same, or end with the same, return true
	if s.SectionRanges[0].Begin == s.SectionRanges[1].Begin || s.SectionRanges[0].End == s.SectionRanges[1].End {
		return true
	}

	// if second range is within first
	if s.SectionRanges[1].Begin <= s.SectionRanges[0].End && s.SectionRanges[1].End <= s.SectionRanges[0].End {
		return true
	}

	return false
}

func (s *SectionAssignment) Overlap() bool {
	// only two section ranges
	// sectionRanges sorted by beginning of each range

	if s.SectionRanges[1].Begin <= s.SectionRanges[0].End {
		return true

	}
	return false
}
