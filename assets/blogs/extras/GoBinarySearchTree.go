package main

import "fmt"

const (
    COUNT int = 5
)

type Node struct {
    left *Node
    data int
    right *Node
}

type BinarySearchTree struct {
    root *Node
}

func (tree *BinarySearchTree) insert (data int) *BinarySearchTree {
    if tree.root == nil {
        tree.root = &Node{
            left: nil,
            data: data,
            right: nil,
        }
    } else {
        tree.root.insert(data)
    }
    return tree
}

func (node *Node) insert (data int) {
    if node == nil {
        return
    } else if data < node.data {
        if node.left == nil {
            node.left = &Node{
                left: nil,
                data: data,
                right: nil,
            }
        } else {
            node.left.insert(data)
        }
    } else {
        if node.right == nil {
            node.right = &Node{
                left: nil,
                data: data,
                right: nil,
            }
        } else {
            node.right.insert(data)
        }
    }
}

func preOrderTraversal(root *Node) {
    if root == nil {
        return
    }
    fmt.Print(root.data, " ")
    preOrderTraversal(root.left)
    preOrderTraversal(root.right)
}

func inOrderTraversal(root *Node) {
    if root == nil {
        return
    }
    inOrderTraversal(root.left)
    fmt.Print(root.data, " ")
    inOrderTraversal(root.right)
}

func postOrderTraversal(root *Node) {
    if root == nil {
        return
    }
    postOrderTraversal(root.left)
    postOrderTraversal(root.right)
    fmt.Print(root.data, " ")
}


func printTree(root *Node, space int) {
    if root == nil {
        return
    }
    space += COUNT
    printTree(root.right, space)
    for i := COUNT; i < space; i++ {
        fmt.Printf(":")
    }
    fmt.Printf("[")
    fmt.Printf("%v", root.data)
    fmt.Println("]")
    printTree(root.left, space)
}

func main(){
	tree := &BinarySearchTree{}
    fmt.Println("=====================================")
	tree.insert(10)
	tree.insert(5)
	tree.insert(15)
	tree.insert(4)
	tree.insert(16)
	tree.insert(6)
	tree.insert(13)
	tree.insert(1)
	tree.insert(12)
	tree.insert(2)
	tree.insert(20)
    printTree(tree.root, 0)
    fmt.Println("=====================================")
    fmt.Printf("Pre-Order: ")
    preOrderTraversal(tree.root)
    fmt.Printf("\nIn-Order: ")
    inOrderTraversal(tree.root)
    fmt.Printf("\nPost-Order: ")
    postOrderTraversal(tree.root)
    fmt.Println()
}
