package main

import (
	"bufio"
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

type Game struct {
	opponent int
	me int  
}

// convert []string into Game
func NewGame(p []string) (*Game, error) {
	if len(p) != 2 {
		return nil, fmt.Errorf("too many players")
	}	

	var o int
	var m int

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
		m = rock
	} else if p[1] == "Y" {
		m = paper
	} else if p[1] == "Z" {
		m = scissors
	} else {
		return nil, fmt.Errorf("unknown entry for me: %v", p)
        }
        
	return &Game{
		opponent: o,
		me: m,
	}, nil
}

func (g *Game) Winner() string {
	// r p s
 	if g.me == g.opponent {
		return "tie"
	}	
	if g.me - g.opponent == 1 || (g.me == rock && g.opponent == scissors) {
		return "me"
	}
	return "opponent"
}

func (g *Game) Score() int {
	ans := g.me
	if g.Winner() == "me" {
		ans += 6
	} else if g.Winner() == "tie" {
		ans += 3
	}
	return ans
}

func main() {
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
