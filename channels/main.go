package main

import (
	"fmt"
	"net/http"
	"time"
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

	for l := range c { //iterating thru string
		time.Sleep(5 * time.Second) //Pause 5 seconds
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) { //Declare
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link //Channel message
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}