package main

import (
	"bufio"
	"fmt"
	"os"
)



type storyPage struct {
	text string
	leftNode *storyPage
	rightNode *storyPage
}


func (node *storyPage) playStory() {

	fmt.Println(node.text)

	scanner := bufio.NewScanner(os.Stdin)

	if (node.leftNode != nil && node.rightNode != nil){


	for {
		scanner.Scan()
		answer := scanner.Text()

		if (answer == "yes") {
			node.leftNode.playStory()
			break
		} else if (answer == "no") {
			node.rightNode.playStory()
			break
		} else {
			fmt.Println("Incorrect input")
		}
	}
	}
}


// This is a binary tree implementation in golang


func main() {

	root := storyPage{"Choose yes/no", nil, nil}
	leftPath := storyPage{"You are at left node", nil,nil}
	rightPath := storyPage{"You are at right node",nil,nil}

	root.leftNode = &leftPath
	root.rightNode = &rightPath

	root.playStory()

}