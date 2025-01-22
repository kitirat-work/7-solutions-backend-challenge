package main

func main() {

}

func MaxRoute(tree [][]int) int {
	startIdx := len(tree) - 2
	for i := startIdx; i >= 0; i-- {
		for nodeIdx := 0; nodeIdx < len(tree[i]); nodeIdx++ {
			left := tree[i+1][nodeIdx]
			right := tree[i+1][nodeIdx+1]
			newValue := maxValueOfNode(tree[i][nodeIdx], left, right)
			tree[i][nodeIdx] = newValue
		}
	}
	return tree[0][0]
}

func maxValueOfNode(node, left, right int) int {
	if left > right {
		return node + left
	}
	return node + right
}
