/*
Another approach towards solving this.

Given an array of [rock, paper, scissors], and a player
choses X (array index Y), this player will win against
a player chosing array element Y-1.  This player will lose
against a player choosing Y+1.

This requires ring logic, to ensure rock beats scissors,
and scissors loses to rock.
*/

package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	rock int = iota + 1
	paper
	scissors
)

const (
	lose = 0
	tie  = 3
	win  = 6
)

var decoderRing *ring.Ring

func initializeDecoderRing() {
	decoderRing = ring.New(3)
	for i := 0; i < decoderRing.Len(); i++ {
		decoderRing.Value = rock + i
		decoderRing = decoderRing.Next()
	}
}

type Game struct {
	opponent int
	me       int
}

// convert []string into Game
func NewGame(p []string) (*Game, error) {
	if len(p) != 2 {
		return nil, fmt.Errorf("too many players")
	}

	var o int
	var d int

	if p[0] == "A" {
		o = rock
	} else if p[0] == "B" {
		o = paper
	} else if p[0] == "C" {
		o = scissors
	} else {
		return nil, fmt.Errorf("unknown entry: %s", p[0])
	}

	if p[1] == "X" {
		d = lose
	} else if p[1] == "Y" {
		d = tie
	} else if p[1] == "Z" {
		d = win
	} else {
		return nil, fmt.Errorf("unknown entry: %v", p)
	}

	// calculate my required play
	m := MyRequiredPlay(o, d)

	return &Game{
		opponent: o,
		me:       m,
	}, nil
}

func MyRequiredPlay(o, d int) int {
	if d == tie {
		return o
	}

	// advance ring until decoderRing.Value == o
	for decoderRing.Value != o {
		decoderRing = decoderRing.Next()
	}

	// to lose, i must play x-1
	if d == lose {
		return decoderRing.Prev().Value.(int)
	}

	// to win, i must play x+1
	if d == win {
		return decoderRing.Next().Value.(int)
	}

	return 0
}

func (g *Game) OutcomeScore() int {
	if g.me == g.opponent {
		return tie
	}

	// advance ring until decoderRing.Value == g.me
	for decoderRing.Value.(int) != g.me {
		decoderRing = decoderRing.Next()
	}

	var advancedCount int

	// advance ring until decoderRing.Value == g.opponent
	for decoderRing.Value.(int) != g.opponent {
		advancedCount += 1
		decoderRing = decoderRing.Next()
	}

	// i played x, opponent played x+1, i lose
	if advancedCount == 1 {
		return lose
	}

	// i played x, opponent played x+2 (x-1), i win
	return win
}

func (g *Game) Score() int {
	return g.me + g.OutcomeScore()
}

func main() {
	initializeDecoderRing()

	//read from stdin, so we can pipe it in
	scanner := bufio.NewScanner(os.Stdin)

	var total int

	for scanner.Scan() {
		t := strings.TrimSpace(scanner.Text())

		play := strings.Split(t, " ")
		g, err := NewGame(play)
		if err != nil {
			log.Fatal(err)
		}
		total += g.Score()
	}
	fmt.Printf("%d\n", total)
}
