package main

import "fmt"



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

	for page.nextPage != nil {
		page = page.nextPage
	}

	page.nextPage = &storyPage{text, nil}

}

func (page *storyPage) addAfter(text string) {
	
	pagePtrToAdd := &storyPage{text, page.nextPage}

	page.nextPage = pagePtrToAdd
}


func main() {
	page1 := storyPage{"1st page", nil}
	page1.addToEnd("2nd Page")
	page1.addToEnd("3rd page")

	page1.addAfter("Hai john")

	page1.playStory() // playStory(&page1) // in the first approach

	// Functions - has return value - 
	// Procedures - has no return value just executes commands
	// Methods - functions that are attached to struct/obj/etc
}