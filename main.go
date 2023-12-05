package main

import (
	"aoc2023/day_one"

	"bufio"
	"fmt"
	"os"
)

type Solvable interface {
	Solve(scanner *bufio.Scanner)
	GetDataPath() string
}

var solutions = map[int]Solvable{
	1: day_one.Solution{},
}

func main() {
	// todo: read from args
	const day = 1
	solution := solutions[day]
	fileScanner, reader := openFile(solution.GetDataPath())
	solution.Solve(fileScanner)
	reader.Close()
}

func openFile(file string) (*bufio.Scanner, *os.File) {

	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	return fileScanner, readFile
}
