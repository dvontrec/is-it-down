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

	for _, link := range links {
		checkLink(link)
	}

}

// Function to check if the link is responding to http traffic
func checkLink(link string) {
	_, err := http.Get(link)
	// If there is an error
	if err != nil {
		// Print that the link may be down
		fmt.Println(link, "Might be down")
		// return
		return
	}
	// print that the link os up
	fmt.Println(link, "Is up")
}
