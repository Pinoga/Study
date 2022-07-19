package tree

import "math"

type Node struct {
	value int
	left  *Node
	right *Node
}

var Tree *Node = &Node{
	left: &Node{
		value: 5,
		right: &Node{
			value: 30,
			right: &Node{
				value: 7,
				left:  nil,
				right: nil,
			},
			left: nil,
		},
		left: nil,
	},
	right: &Node{
		value: 5,
		left: &Node{
			value: 30,
			left: &Node{
				value: 7,
				left:  nil,
				right: nil,
			},
			right: nil,
		},
		right: nil,
	},
	value: 3,
}

func BinaryTreeSum(root *Node) int {

	if root == nil {
		return 0
	}

	sum := root.value

	sum += BinaryTreeSum(root.left)
	sum += BinaryTreeSum(root.right)

	return sum
}

func BinaryTreeMininum(root *Node) int {

	if root == nil {
		return int(math.NaN())
	}

	min := root.value

	leftMin := BinaryTreeMininum(root.left)
	if leftMin != int(math.NaN()) {
		min = int(math.Min(float64(min), float64(leftMin)))
	}
	rightMin := BinaryTreeMininum(root.right)
	if rightMin != int(math.NaN()) {
		min = int(math.Min(float64(min), float64(rightMin)))
	}

	return min
}

func MaxPathSum(root *Node) int {

	if root == nil {
		return int(math.Inf(-1))
	}

	sum := root.value

	if root.left == nil && root.right == nil {
		return sum
	}

	maxSumOfLeft := MaxPathSum(root.left)
	maxSumOfRight := MaxPathSum(root.right)

	return sum + int(math.Max(float64(maxSumOfLeft), float64(maxSumOfRight)))
}

// func isSymmetric(root *Node) bool {

// 	queue := make([]*Node, 0)
// 	queue = append(queue, root)

// 	for len(queue) > 0 {
// 		node := queue[0]
// 		queue = queue[1:]

// 		if node.left != nil

// 	}

// }

func IsSymmetric(root *Node) bool {

	if root == nil {
		return true
	}

	leftChild := root.left
	rightChild := root.right

	return areSymmetric(leftChild, rightChild)
}

func areSymmetric(first *Node, second *Node) bool {
	if first == nil && second == nil {
		return true
	}
	if first == nil || second == nil {
		return false
	}
	if first.value != second.value {
		return false
	}

	return areSymmetric(first.left, second.right) && areSymmetric(first.right, second.left)
}
