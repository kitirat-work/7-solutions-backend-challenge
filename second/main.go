package main

import "strconv"

func main() {

}

func Decode(msg string) string {
	result := make([]int, len(msg)+1)
	nodes := make([]node, len(msg)+1)
	nodes[0] = node{
		nodeType: string(msg[0]),
		value:    &result[0],
		left:     nil,
	}
	for i := 1; i < len(result); i++ {
		if len(msg) != i {
			nodes[i].nodeType = string(msg[i])
		}
		nodes[i].value = &result[i]
		nodes[i].left = &nodes[i-1]
		nodes[i-1].right = &nodes[i]
	}
	for i := 0; i < len(nodes); i++ {
		nodes[i].Calculate(LEFT)
	}

	str := ""
	for _, v := range result {
		str += strconv.Itoa(v)
	}
	return str
}

type node struct {
	nodeType string
	value    *int
	left     *node
	right    *node
}
type Direction = string

const (
	LEFT  Direction = "LEFT"
	RIGHT Direction = "RIGHT"
)

func (n *node) Calculate(from Direction) {
	if n.nodeType == "L" {
		if *n.value > *n.right.value {
			return
		}
		*n.value = *n.right.value + 1
	} else if n.nodeType == "R" {
		if *n.value < *n.right.value {
			return
		}
		*n.right.value = *n.value + 1
	} else if n.nodeType == "=" {
		if *n.value == *n.right.value {
			return
		}
		if from == LEFT {
			*n.right.value = *n.value
		} else if from == RIGHT {
			*n.value = *n.right.value
		}
	}
	n.trigger()
}

func (n *node) trigger() {
	if n.left != nil {
		n.left.Calculate(RIGHT)
	}
	if n.right != nil {
		n.right.Calculate(LEFT)
	}
}
