package solutions

import (
	"fmt"
	"log"
	"math/rand"
	"slices"
	"strconv"
)

// type board [][]int
type Direction string

func (d Direction) IsOpposite(op Direction) bool {
	switch d {
	case DirectionUP:
		return op == DirectionDown
	case DirectionDown:
		return op == DirectionUP
	case DirectionLeft:
		return op == DirectionRight
	case DirectionRight:
		return op == DirectionLeft
	}
	return false
}

const (
	DirectionUP    Direction = "UP"
	DirectionDown  Direction = "DOWN"
	DirectionLeft  Direction = "LEFT"
	DirectionRight Direction = "RIGHT"
)

func (b board) Print() {
	for i := 0; i <= 2; i++ {
		fmt.Println(b[i])
	}
}

func (b board) Copy() board {
	return [][]int{
		slices.Clone(b[0]),
		slices.Clone(b[1]),
		slices.Clone(b[2]),
	}
}

type board [][]int

func (b board) Serialize() string {
	list := flatten(b)

	res := ""
	for _, r := range list {
		res += fmt.Sprintf("%d", r)
	}
	return res
}

func NewBoardDeserialized(serialized string) board {
	l := make([]int, len(serialized))

	for i, a := range serialized {
		d, _ := strconv.ParseInt(string(a), 10, 64)
		l[i] = int(d)
	}

	res := board([][]int{
		l[0:3],
		l[3:6],
		l[6:9],
	})

	return res

}

func (b board) EmptyIndex() (x, y int) {
	list := flatten(b)
	idx := slices.IndexFunc(list, func(v int) bool {
		return v == 0
	})

	return int(idx / 3), idx % 3
}

func (b board) IsSolved() bool {
	l := flatten(b)

	side := append([]int{}, l[0:8]...)
	log.Println(side)

	return slices.IsSorted(side) && l[8] == 0
}

func nextMove(i, j int) []Direction {

	/*
	   (0, 0)    (0, 1)    (0, 2)
	   (1, 0)    (1, 1)    (1, 2)
	   (2, 0)    (2, 1)    (2, 2)
	*/
	res := []Direction{}
	if j > 0 {
		res = append(res, DirectionRight)
	}

	// down
	if i < 2 {
		res = append(res, DirectionUP)
	}

	// up
	if i > 0 {
		res = append(res, DirectionDown)
	}

	// right
	if j < 2 {
		res = append(res, DirectionLeft)
	}

	return res

}

func printSteps(history map[int][]Node, lastNode Node) {
	level := lastNode.Level - 1
	curIndex := lastNode.PrevIndex

	for i := level; i >= 1; i-- {
		nodes := history[i]
		fmt.Printf("Step %d: Out of %d \n", i, len(nodes))

		for _, n := range nodes {
			if n.Index == curIndex {
				b := NewBoardDeserialized(n.Value)
				b.Print()
				fmt.Println("-------")
				fmt.Println("")
				// Update curIndex only when a match is found
				curIndex = n.PrevIndex
				break // Exit the loop since we've found the correct node
			}
		}
	}
}

func flatten(b board) []int {

	result := []int{}

	for _, r := range b {
		result = append(result, r...)
	}
	return result
}

func shuffle() board {
	result := board{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 0},
	}

	log.Println("in the beginning", result.IsSolved())

	lists := flatten(result)

	rand.Shuffle(len(lists), func(i, j int) {
		lists[i], lists[j] = lists[j], lists[i]
	})

	return [][]int{
		lists[0:3],
		lists[3:6],
		lists[6:9],
	}

}

type Node struct {
	Value     string
	Level     int
	Index     int64
	PrevIndex int64
}
type Queue struct {
	items []Node
}

func (q *Queue) Enqueue(value Node) {
	q.items = append(q.items, value)
}

func (q *Queue) Dequeue() Node {
	if len(q.items) == 0 {
		panic("Queue is empty")
	}
	value := q.items[0]
	q.items = q.items[1:] // Avoids reallocation by slicing
	return value
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func Solution6() {
	b := shuffle()
	b.Print()
	fmt.Println("")

	steps := map[int][]Node{}

	visitedState := map[string]bool{}
	solved := false
	queue := Queue{
		items: []Node{},
	}

	node := Node{b.Serialize(), 1, 1, 0}
	queue.Enqueue(node)
	steps[1] = []Node{node}

	lastNode := node

	for !queue.IsEmpty() {
		n := queue.Dequeue()
		b = NewBoardDeserialized(n.Value)
		level := n.Level

		if b.IsSolved() {
			solved = true
			lastNode = n
			fmt.Println("step on:", level)
			b.Print()
			break
		}

		if visitedState[n.Value] {
			continue
		}

		visitedState[n.Value] = true
		i, j := b.EmptyIndex()
		directions := nextMove(i, j)

		for _, m := range directions {
			newState := b.Copy()
			switch m {
			case DirectionUP:
				newState[i][j] = newState[i+1][j]
				newState[i+1][j] = 0

			case DirectionLeft:
				newState[i][j] = newState[i][j+1]
				newState[i][j+1] = 0

			case DirectionDown:
				newState[i][j] = newState[i-1][j]
				newState[i-1][j] = 0

			case DirectionRight:
				newState[i][j] = newState[i][j-1]
				newState[i][j-1] = 0

			}

			nodeState := Node{
				Value:     newState.Serialize(),
				Level:     level + 1,
				Index:     int64(rand.Int()),
				PrevIndex: n.Index,
			}

			queue.Enqueue(nodeState)

			if len(steps[level+1]) == 0 {
				steps[level+1] = []Node{}
			}

			steps[level+1] = append(steps[level+1], nodeState)
		}

	}

	if solved {
		printSteps(steps, lastNode)
	}

	fmt.Println("finished", solved, b.Serialize(), b.IsSolved())
	b.Print()
}
