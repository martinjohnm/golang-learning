package main

import "fmt"

type storyPage struct {
	text string
	nextPage *storyPage
}


func playStory(page *storyPage) {

	if page == nil {
		return
	}

	fmt.Println(page.text)

	playStory(page.nextPage)
}



func main() {
	page1 := storyPage{"1st page", nil}
	page2 := storyPage{"2nd page", nil}
	page3 := storyPage{"3rd page", nil}

	page1.nextPage = &page2
	page2.nextPage = &page3

	playStory(&page1)
}