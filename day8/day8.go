package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var infile = "day8_example.txt"

//var infile = "day8.txt"

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

func getForest(input []string) [][]int {
	forest := make([][]int, 0)
	for _, row := range input {
		tree_row := make([]int, len(row))
		for y, column := range row {
			column_int, _ := strconv.Atoi(string(column))
			tree_row[y] = column_int
		}
		forest = append(forest, tree_row)
	}
	return forest
}

func part1(forest [][]int) {
	var total_visible = 0
	for x := 0; x < len(forest[0]); x++ {
		for y := 0; y < len(forest); y++ {
			for _, v := range forest[x][0:y] {
				if v < forest[x][y] {
					total_visible++
					break
				}
			}
			for _, v := range forest[x][y : len(forest)-1] {
				if v < forest[x][y] {
					total_visible++
					break
				}
			}
			for _, v := range forest[0:][y] {
				fmt.Println(x, y)
				if v < forest[x][y] {
					total_visible++
				}
			}
			for _, v := range forest[x : len(forest[x])-1][y] {
				if v < forest[x][y] {
					total_visible++
					break
				}
			}
		}

	}
	fmt.Println(total_visible)
}

func main() {
	forest := getForest(importfile())
	part1(forest)
}
