package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var infile = "day4_hms.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func importfile() []string {
	return_me := []string{}
	file, err := os.Open(infile)
	check(err)
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		return_me = append(return_me, string(line))
	}
	return return_me
}

func part1(input []string) {
	var count = 0
	for _, i := range input {
		line := strings.Split(i, ",")

		left, right := line[0], line[1]

		left_l, err := strconv.Atoi(strings.Split(left, "-")[0])
		check(err)
		left_r, err := strconv.Atoi(strings.Split(left, "-")[1])
		check(err)

		right_l, err := strconv.Atoi(strings.Split(right, "-")[0])
		check(err)
		right_r, err := strconv.Atoi(strings.Split(right, "-")[1])
		check(err)

		if (left_l >= right_l && left_r <= right_r) || (right_l >= left_l && right_r <= left_r) {
			fmt.Println(left_l, left_r, right_l, right_r)
			count++
		}

		//fmt.Println(left_l, left_r, right_l, right_r)
	}
	fmt.Println(count)
}

func part2(input []string) {
	var count = 0
	for _, i := range input {
		line := strings.Split(i, ",")

		left, right := line[0], line[1]

		left_l, err := strconv.Atoi(strings.Split(left, "-")[0])
		check(err)
		left_r, err := strconv.Atoi(strings.Split(left, "-")[1])
		check(err)

		right_l, err := strconv.Atoi(strings.Split(right, "-")[0])
		check(err)
		right_r, err := strconv.Atoi(strings.Split(right, "-")[1])
		check(err)

		if (left_l >= right_l || left_r <= right_r) || (right_l >= left_l || right_r <= left_r) {
			fmt.Println(left_l, left_r, right_l, right_r)
			count++
		}

	}
	fmt.Println(count)
}

func main() {
	input := importfile()
	part1(input)
	part2(input)
}
