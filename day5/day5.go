package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
  "math"
)

var infile = "day5_hms_clean.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func importfile() [][]int {
	file, err := os.Open(infile)
	check(err)
	ret := make([][]int, 0)
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		//fmt.Println(line)
		templine := strings.Split(string(line), ",")
		fmt.Println(templine)
		tempslice := make([]int, 0)
		for _, i := range templine {
			asdf, _ := strconv.Atoi(i)
			//fmt.Println(reflect.TypeOf(asdf))
			tempslice = append(tempslice, asdf)
		}
		//byteToInt, _ := strconv.Atoi(string(templine))
		ret = append(ret, tempslice)
	}
	return ret
}

type stack []string

func (s stack) Push(v string) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, string) {
	l := len(s)
	return s[:l-1], s[l-1]
}

/*
            [L] [M]         [M]
        [D] [R] [Z]         [C] [L]
        [C] [S] [T] [G]     [V] [M]
[R]     [L] [Q] [B] [B]     [D] [F]
[H] [B] [G] [D] [Q] [Z]     [T] [J]
[M] [J] [H] [M] [P] [S] [V] [L] [N]
[P] [C] [N] [T] [S] [F] [R] [G] [Q]
[Z] [P] [S] [F] [F] [T] [N] [P] [W]
*/

func main() {
	//instructions := importfile()
	ship := make([]stack, 0)
	ship = append(ship, []string{"R", "H", "M", "P", "Z"})
	ship = append(ship, []string{"B", "J", "C", "P"})
	ship = append(ship, []string{"D", "C", "L", "G", "H", "N", "S"})
	ship = append(ship, []string{"L", "R", "S", "Q", "D", "M", "T", "F"})
	ship = append(ship, []string{"M", "Z", "T", "B", "Q", "P", "S", "F"})
	ship = append(ship, []string{"G", "B", "Z", "S", "F", "T"})
	ship = append(ship, []string{"V", "R", "N"})
	ship = append(ship, []string{"M", "C", "V", "D", "T", "L", "G", "P"})
	ship = append(ship, []string{"L", "M", "F", "J", "N", "Q", "W"})
	//part1(ship, instructions)
  //partOne()
  partTwo()
}

//Seek inspiration from here I guess
//https://old.reddit.com/r/adventofcode/comments/zcxid5/2022_day_5_solutions/iz583oy/
func partOne() {
  input, _ := os.ReadFile("day5_hms.txt")
    split := strings.Split(string(input), "\n\n")
    layout := strings.Split(split[0], "\n")
    stacks := make(map[int][]rune)
    stackKeys := []int{}
    for i := len(layout) - 1; i >= 0; i-- {
        if i == len(layout)-1 {
            for _, k := range strings.Split(strings.TrimSpace(layout[i]), "   ") {
                key, _ := strconv.Atoi(k)
                stacks[key] = []rune{}
                stackKeys = append(stackKeys, key)
            }
        } else {
            for i, c := range layout[i] {
                if !strings.ContainsAny(string(c), " []") {
                    key := int(math.Ceil(float64(i) / 4))
                    stacks[key] = append(stacks[key], c)
                }
            }
        }

    }
    moves := strings.Split(split[1], "\n")
    
  
    for _, move := range moves {
        var amount int
        var from int
        var to int
        fmt.Sscanf(move, "move %d from %d to %d", &amount, &from, &to)
        source := stacks[from]
        destination := stacks[to]
        for i := 0; i < amount; i++ {
            n := len(source) - 1 // Top element
            destination = append(destination, source[n])
            source = source[:n] // Pop
        }
        stacks[from] = source
        stacks[to] = destination
    }
    partOne := ""
    for _, k := range stackKeys {
        partOne += string(stacks[k][len(stacks[k])-1])
    }
    fmt.Println(partOne)
  //VQZNJMWTR
}

func partTwo() {
  input, _ := os.ReadFile("day5_hms.txt")
    split := strings.Split(string(input), "\n\n")
    layout := strings.Split(split[0], "\n")
    stacks := make(map[int][]rune)
    stackKeys := []int{}
    for i := len(layout) - 1; i >= 0; i-- {
        if i == len(layout)-1 {
            for _, k := range strings.Split(strings.TrimSpace(layout[i]), "   ") {
                key, _ := strconv.Atoi(k)
                stacks[key] = []rune{}
                stackKeys = append(stackKeys, key)
            }
        } else {
            for i, c := range layout[i] {
                if !strings.ContainsAny(string(c), " []") {
                    key := int(math.Ceil(float64(i) / 4))
                    stacks[key] = append(stacks[key], c)
                }
            }
        }

    }
    moves := strings.Split(split[1], "\n")
    for _, move := range moves {
        var amount int
        var from int
        var to int
        fmt.Sscanf(move, "move %d from %d to %d", &amount, &from, &to)
        source := stacks[from]
        destination := stacks[to]
        tempstack := []rune{}
        for i := 0; i < amount; i++ {
            n := len(source) - 1 // Top element
            tempstack = append(tempstack, source[n])
            source = source[:n] // Pop
        }
      for i := 0; i < amount; i++ {
            n := len(tempstack) - 1 // Top element
            destination = append(destination, tempstack[n])
            tempstack = tempstack[:n] // Pop
        }

        stacks[from] = source
        stacks[to] = destination
    }
    partTwo := ""
    for _, k := range stackKeys {
        partTwo += string(stacks[k][len(stacks[k])-1])
    }
    fmt.Println(partTwo)
  //NLCDCLVMQ
}

/*This shit didn't work
func part1(p1_ship []stack, instructions [][]int) {
	for _, line := range instructions {
		fmt.Println(line)
		fmt.Println("Coming in:", p1_ship)
		for v := 0; v < line[0]; v++ {
			new_move_from, temp := p1_ship[line[1]-1].Pop()
			p1_ship[line[1]-1] = new_move_from
			p1_ship[line[2]-1] = p1_ship[line[2]-1].Push(temp)
			fmt.Println(p1_ship[line[2]-1])
		}
		fmt.Println("Going out:", p1_ship, "\n")
	}
	fmt.Println("Final stacks:", p1_ship)
	//answer?: NBSCJQLRT
	fmt.Printf("Final answer?: ")
	for _, v := range p1_ship {
		_, val := v.Pop()
		fmt.Printf("%s", val)
	}
	fmt.Println()
}*/
