package main

import (
	"fmt"
	"io/ioutil"
)

var infile = "day6.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func importfile() string {
	f, err := ioutil.ReadFile(infile)
	check(err)
	return string(f)
}

func part1(input string) {

	for i, v := range input {
		if i < 3 {
			continue
		}
		if input[i] != input[i-1] && input[i] != input[i-2] && input[i] != input[i-3] {
			if input[i-1] != input[i-2] && input[i-1] != input[i-3] {
				if input[i-2] != input[i-3] {
					fmt.Printf("%d:  %s\n", i+1, string(v))
					break
				}
			}
		}
	}

}

func isMarker(arr string) bool {
	m := make(map[rune]bool)
	for _, i := range arr {
		_, ok := m[i]
		if ok {
			return false
		}
		m[i] = true
	}
	return true
}

func part2(input string) {
	for i := 0; i < len(input)-14; i++ {
		if isMarker(input[i : i+14]) {
			fmt.Printf("%d:  %s\n", i+14, string(input[i+14]))
			break
		}
	}
}

func main() {
	data := importfile()
	part1(data)
	part2(data)
}
