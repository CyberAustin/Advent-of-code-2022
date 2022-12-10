package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var infile = "day8_example.txt"

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

func part1(input []string) {
	for i, v := range input {
		fmt.Printf("%d, %d: %s\n", i, len(v), v)
	}
}

func main() {
	in := importfile()
	part1(in)
}
