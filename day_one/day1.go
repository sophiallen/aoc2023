package day_one

import (
	"bufio"
	"fmt"
	"strconv"
)

func (s Solution) GetDataPath() string {
	return "day_one/data.txt"
}

type Solution struct{}

type Node struct {
	Letter   rune
	Value    int
	Children map[rune]*Node
}

func buildTree(word string, root *Node, value int) *Node {
	currentNode := root
	for _, ltr := range word {
		_, ok := currentNode.Children[ltr]
		if !ok {
			currentNode.Children[ltr] = &Node{
				Letter:   ltr,
				Children: map[rune]*Node{},
			}
		}
		currentNode = currentNode.Children[ltr]
	}
	currentNode.Value = value
	return root
}

func printTree(n *Node) {
	for k := range n.Children {
		fmt.Printf("%s > ", string(k))
		printTree(n.Children[k])
	}
}

func findValue(s string, start int, root *Node) (int, int) {
	curNode := root
	for i, ltr := range s[start:] {
		val, ok := curNode.Children[ltr]
		if !ok {
			return -1, -1
		}
		if val.Value != 0 {
			return val.Value, i
		}
		curNode = curNode.Children[ltr]
	}
	return -1, -1
}

func (s Solution) Solve(scanner *bufio.Scanner) {
	tree := &Node{
		Children: map[rune]*Node{},
	}
	words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for i, word := range words {
		tree = buildTree(word, tree, i+1)
	}
	// printTree(tree)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		accum := []int{}

		for i := range line {
			ltr := rune(line[i])
			if !isLetter(ltr) {
				val, _ := strconv.Atoi(string(ltr))
				accum = append(accum, val)
				continue
			}
			val, _ := findValue(line, i, tree)
			if val != -1 {
				accum = append(accum, val)
			}
		}
		strval := fmt.Sprintf("%+v%+v", accum[0], accum[len(accum)-1])
		v, _ := strconv.Atoi(strval)
		sum += v
	}
	fmt.Printf("Sum is %+v", sum)
}

func isLetter(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func (s Solution) SolvePtOne(scanner *bufio.Scanner) {
	sum := 0
	for scanner.Scan() {
		s := scanner.Text()
		accum := []rune{}
		for _, ltr := range s {
			if ltr >= 'a' && ltr <= 'z' {
				continue
			}
			accum = append(accum, ltr)
		}
		firstlast := string(accum[0]) + string(accum[len(accum)-1])
		fmt.Println(firstlast)
		num, _ := strconv.Atoi(firstlast)
		sum += num
	}
	fmt.Printf("SUM: %+v", sum)
}
