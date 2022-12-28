package main

import (
	"fmt"
	cregex "github.com/mingrammer/commonregex"
)

func main() {
	test()
}

func test() {
	text := `John, please get that article on www.linkedin.com to me by 5:00PM on Jan 9th 2012. 4:00 would be ideal, actually. If you have any questions, You can reach me at (519)-236-2723x341 or get in touch with my associate at harold.smith@gmail.com`
	emails := cregex.Emails(text)
	fmt.Println("emails", emails)
}
