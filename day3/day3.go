package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var infile = "./day3.txt"

var supplies = make(map[string]int)

var priority = map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26, "A": 27, "B": 28, "C": 29, "D": 30, "E": 31, "F": 32, "G": 33, "H": 34, "I": 35, "J": 36, "K": 37, "L": 38, "M": 39, "N": 40, "O": 41, "P": 42, "Q": 43, "R": 44, "S": 45, "T": 46, "U": 47, "V": 48, "W": 49, "X": 50, "Y": 51, "Z": 52}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func importfile() []string {
	f, err := ioutil.ReadFile(infile)
	check(err)
	arr := strings.Split(string(f), "\n")
	return arr
}

func part1(part1_in []string) {
	var count = 0

	for _, a := range part1_in {
		already_found := ""
		left := a[0 : len(a)/2]
		right := a[len(a)/2:]
		for _, i := range left {
			if strings.Contains(already_found, string(i)) {
				continue
			}
			if strings.Contains(right, string(i)) {
				fmt.Println(string(i))
				count += priority[string(i)]
				already_found += string(i)
			}
		}
	}
	fmt.Println(count)
	//Answer: 7428
}

func part2(part2_in []string) {
  var count2 = 0
	for i := 0; i < len(part2_in); i += 3 {
    one := strings.Trim(part2_in[i]," ")
    two := strings.Trim(part2_in[i+1]," ")
    three := strings.Trim(part2_in[i+2]," ")

    //fmt.Printf("One: %s\nTwo: %s\nThree: %s\n", one, two, three)
    
    for _, i := range one {
      fmt.Println(i)
      already_found2 := ""
      if strings.Contains(already_found2, string(i)) {
        //fmt.Println(string(i))
        continue
      }
  		if strings.Contains(two, string(i)) {
        if strings.Contains(three, string(i)) {
        count2 += priority[string(i)]
        already_found2 += string(i)
         // fmt.Println(rune(i))
         }
      }
      //fmt.Println(already_found2)
    }
  }
  
  fmt.Println(count2)
}
	


func main() {
	input := importfile()
	//part1(input)
	part2(input)
}
