package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type RoShamBo struct {
	player1, player2 string
}

func main() {
	d, err := importfile()
	check(err)
	points := fillPoints()
	weapon := selectRPS()
	var score1, score2 int
	for _, v := range d {
		score1 += points[v.player2] + points[v.player1+v.player2]
		score2 += points[weapon[v.player1+v.player2]] + points[v.player1+weapon[v.player1+v.player2]]
	}
	fmt.Printf("Part 1: %d\n", score1)
	fmt.Printf("Part 2: %d\n", score2)
}

func fillPoints() map[string]int {
	p := make(map[string]int)
	// points on selected 'weapon'
	p["A"] = 1 // p1 rock
	p["B"] = 2 // p1 paper
	p["C"] = 3 // p1 scissor
	p["X"] = 1 // p2 rock
	p["Y"] = 2 // p2 paper
	p["Z"] = 3 // p2 scissor
	// win, draw, lost player 2 truth table
	p["AX"] = 3 // r+r draw
	p["AY"] = 6 // r+p win
	p["AZ"] = 0 // r+s lost
	p["BX"] = 0 // p+r lost
	p["BY"] = 3 // p+p draw
	p["BZ"] = 6 // p+s win
	p["CX"] = 6 // s+r win
	p["CY"] = 0 // s+p lost
	p["CZ"] = 3 // s+s draw
	return p
}

func selectRPS() map[string]string {
	// X = loose, Y = Draw, Z = Win
	s := make(map[string]string)
	s["AX"] = "Z" // loss, r+s
	s["AY"] = "X" // draw, r+r
	s["AZ"] = "Y" // win, r+p
	s["BX"] = "X" // loss, p+r
	s["BY"] = "Y" // draw, p+p
	s["BZ"] = "Z" // win, p+s
	s["CX"] = "Y" // loss, s+p
	s["CY"] = "Z" // draw, s+s
	s["CZ"] = "X" // win, s+r
	return s
}

func importfile() ([]RoShamBo, error) {
	var d []RoShamBo
	f, err := ioutil.ReadFile("./day2_hms.txt")
	if err != nil {
		return d, err
	}
	arr := strings.Split(string(f), "\n")
	for _, v := range arr {
		if v != "" {
			rps := strings.Split(v, " ")
			dT := RoShamBo{
				player1: rps[0],
				player2: rps[1],
			}
			d = append(d, dT)
		}
	}
	fmt.Printf("loaded %d data-points\n", len(d))
	return d, nil
}
