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

	if (node.leftNode != nil && node.rightNode != nil) {

		for {
			scanner.Scan()

			answer := scanner.Text()

			if answer == "yes" {
				node.leftNode.playStory()
				break
			} else if answer == "no" {
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

	root := storyPage{"This is root type yes/no",nil,nil}
	
	leftPath := storyPage{"Left node", nil,nil}

	rightPath := storyPage{"Right path", nil,nil}

	root.leftNode = &leftPath
	root.rightNode = &rightPath


	root.playStory()


}