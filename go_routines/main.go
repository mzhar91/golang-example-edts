package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	// routinesWait()
	channel()
}

func routinesWait() {
	urls := []string{
		"https://sg-edts.com/",
		"https://indomaret.co.id/",
		"https://www.klikindomaret.com/",
		"https://id.wikipedia.org/wiki/Salim_Group",
		"https://indogrosir.co.id/",
		"https://www.klikindogrosir.com/",
	}
	var wg sync.WaitGroup
	
	for _, u := range urls {
		// Increment the wait group counter
		wg.Add(1)
		
		go func(url string) {
			// Decrement the counter when the go routine completes
			defer wg.Done()
			// Call the function check
			checkUrl(url)
		}(u)
		
	}
	
	// Wait for all the checkWebsite calls to finish
	wg.Wait()
}

func channel() {
	urls := []string{
		"https://sg-edts.com/",
		"https://indomaret.co.id/",
		"https://www.klikindomaret.com/",
		"https://id.wikipedia.org/wiki/Salim_Group",
		"https://indogrosir.co.id/",
		"https://www.klikindogrosir.com/",
	}
	
	c := make(chan urlStatus)
	for _, url := range urls {
		go checkUrlStatus(url, c)
	}
	
	result := make([]urlStatus, len(urls))
	for i, _ := range result {
		result[i] = <-c
		if result[i].status {
			fmt.Println(result[i].url, "is up and running.")
		} else {
			fmt.Println(result[i].url, "is down !!!")
		}
	}
}

// checks and prints a message if a website is up or down
func checkUrl(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "is down !!!")
		return
	}
	fmt.Println(url, "is up and running.")
}

// checks and prints a message if a website is up or down
func checkUrlStatus(url string, c chan urlStatus) {
	_, err := http.Get(url)
	if err != nil {
		// The website is down
		c <- urlStatus{url, false}
	} else {
		// The website is up
		c <- urlStatus{url, true}
	}
}

type urlStatus struct {
	url    string
	status bool
}
