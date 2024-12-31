package main

import "fmt"

// type storyPage struct {
// 	text string
// 	nextPage *storyPage
// }


// func playStory(page *storyPage) {

// 	if page == nil { // checking for null pointer
// 		return
// 	}

// 	fmt.Println(page.text)

// 	playStory(page.nextPage)
// }

type storyPage struct {
	text string
	nextPage *storyPage
}


func (page *storyPage) playStory() {
	// //with recursion
	// if page == nil {
	// 	return
	// }

	// fmt.Println(page.text)
	// page.nextPage.playStory()

	// with loop 
	for page != nil {
		fmt.Println(page.text)
		page = page.nextPage
	}
}

func (page *storyPage) addToEnd(text string) {

	pagePointerToAdd := &storyPage{text, nil}

	for page.nextPage != nil {
		page = page.nextPage
	}

	page.nextPage = pagePointerToAdd

}




func main() {
	page1 := storyPage{"1st page", nil}
	page1.addToEnd("2nd Page")
	page1.addToEnd("3rd page")


	page1.playStory() // playStory(&page1) // in the first approach

	// Functions - has return value - 
	// Procedures - has no return value just executes commands
	// Methods - functions that are attached to struct/obj/etc
}