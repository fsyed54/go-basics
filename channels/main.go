package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string) //Channel

	for _, link := range links {
		go checkLink(link, c) //Added channel
	}

	for i := 0; i < len(links); i++ { //iterating thru string
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) { //Declare
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think!" //Channel message
		return
	}
	fmt.Println(link, "is up!")
	c <- "Yes it is up!"
}