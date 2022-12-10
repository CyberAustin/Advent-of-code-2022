package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Node struct {
	children []*Node
	name     string
	size     int
}

var infile = "day7.txt"

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

func createStructure(input []string) []Node {
	tree := make([]Node, 0)
	fmt.Println(tree)

	for _, v := range input {
		fmt.Println(v)
	}
	return tree
	//do the things
}

func part1(input []string) {
	asdf := createStructure(input)
	fmt.Println(asdf)
}

func main() {
	part1(importfile())
}
