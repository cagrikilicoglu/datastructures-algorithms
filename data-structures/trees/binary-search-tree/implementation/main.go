package main

import "fmt"

type BinarySearchTree struct {
	root *Node
}
type Node struct {
	value int
	right *Node
	left  *Node
}

type NodeShow struct {
	value int
	right Node
	left  Node
}

func newBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}
func createNode(value int) *Node {
	return &Node{value: value}
}

func main() {
	myBinarySearchTree := newBinarySearchTree()
	myBinarySearchTree.insert(9)
	myBinarySearchTree.insert(4)
	myBinarySearchTree.insert(6)
	myBinarySearchTree.insert(20)
	myBinarySearchTree.insert(170)
	myBinarySearchTree.insert(15)
	myBinarySearchTree.insert(1)

	fmt.Println(*myBinarySearchTree)
	fmt.Println(*myBinarySearchTree.root)
	fmt.Println(*myBinarySearchTree.root.right)
	fmt.Println(*myBinarySearchTree.root.left)
	fmt.Println(*myBinarySearchTree.lookup(4).right)
	fmt.Println(*myBinarySearchTree.lookup(4).left)
	fmt.Println(myBinarySearchTree.lookup(20).right)
	fmt.Println(myBinarySearchTree.lookup(20).left)
	fmt.Println(myBinarySearchTree.lookup(6))
	fmt.Println(myBinarySearchTree.lookup(170))
	fmt.Println(myBinarySearchTree.lookup(15))
	fmt.Println(myBinarySearchTree.lookup(1))

}

func (b *BinarySearchTree) insert(value int) {
	newNode := createNode(value)
	if b.root == nil {
		b.root = newNode
		return
	}

	currentNode := b.root
	for {
		if newNode.value > currentNode.value {
			if currentNode.right == nil {
				currentNode.right = newNode
				break
			}
			currentNode = currentNode.right

		} else {
			if currentNode.left == nil {
				currentNode.left = newNode
				break
			}
			currentNode = currentNode.left

		}
	}

}

func (b *BinarySearchTree) lookup(value int) *Node {
	currentNode := b.root
	for currentNode != nil {
		if value > currentNode.value {
			currentNode = currentNode.right
		} else if value < currentNode.value {
			currentNode = currentNode.left
		} else {
			return currentNode
		}
	}
	return nil
}

func (b *BinarySearchTree) remove(value int) {

	currentNode := b.root

	// for currentNode.right.value == value || currentNode.left.value == value {
	for currentNode != nil {
		if value > currentNode.value {
			if currentNode.right.value == value {
				if currentNode.right.right == nil && currentNode.right.left == nil {
					currentNode.right = nil
				} else if currentNode.right.left != nil && currentNode.right.right == nil {
					currentNode.right = currentNode.right.left
				} else if currentNode.right.right != nil && currentNode.right.left == nil {
					currentNode.right = currentNode.right.right
				} else {
					// TODO
				}
				return

			} else {
				currentNode = currentNode.right
			}
		} else if value < currentNode.value {
			if currentNode.left.value == value {
				if currentNode.left.right == nil && currentNode.left.left == nil {
					currentNode.left = nil
				} else if currentNode.left.left != nil && currentNode.left.right == nil {
					currentNode.left = currentNode.left.left
				} else if currentNode.left.right != nil && currentNode.left.left == nil {
					currentNode.left = currentNode.left.right
				} else {
					// TODO
				}
				return
			} else {
				currentNode = currentNode.left
			}

		}
	}
	return

}
