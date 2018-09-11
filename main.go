package main

import (
	"fmt"
	"net/http"
)

func main() {
	// creates a slice of links
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://twitter.com",
		"http://amazon.com",
		"http://golang.org",
	}

	// Creates a new channel to communicate with new go routines
	c := make(chan string)

	for _, link := range links {
		// Sets up a new go routine to call the checkline function
		go checkLink(link, c)
	}
	// Loops through the links infinitly whenever a value is received
	for {
		// calls the check link function with the link again
		go checkLink(<-c, c)
	}
}

// Function to check if the link is responding to http traffic
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	// If there is an error
	if err != nil {
		// Print that the link may be down
		fmt.Println(link, "Might be down")
		// Sends message to channel
		c <- link
		// return
		return
	}
	// print that the link os up
	fmt.Println(link, "Is up")
	c <- link
}
