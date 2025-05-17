package webfetch

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

// Function to fetch and count words from a single URL
func scrapeURL(url string, wg *sync.WaitGroup, ch chan<- int) {
defer wg.Done() // Ensure that the WaitGroup counter is decremented once the goroutine finishes.
// Fetch the URL
resp, err := http.Get(url)
if err != nil {
log.Printf("Error fetching %s: %v\n", url, err)
ch <- 0 // Send 0 count to channel on error
return
}
defer resp.Body.Close()
// Read the response body
body, err := ioutil.ReadAll(resp.Body)
if err != nil {
log.Printf("Error reading body of %s: %v\n", url, err)
ch <- 0 // Send 0 count to channel on error
return
}
// Count the words in the body
wordCount := len(strings.Fields(string(body)))
// Send the word count back to the channel
ch <- wordCount
}
func Call() {
// List of URLs to scrape
urls := []string{
"https://www.golang.org",
"https://www.github.com",
"https://www.example.com",
}
// Create a WaitGroup to wait for all goroutines to finish
var wg sync.WaitGroup
// Create a channel to receive the word counts
ch := make(chan int, len(urls)) // Buffered channel to hold word counts
// Loop through the URLs and start a goroutine for each URL
for _, url := range urls {
wg.Add(1)
go scrapeURL(url, &wg, ch)
}
// Wait for all goroutines to finish
wg.Wait()
// Close the channel after all goroutines are done
close(ch)
// Aggregate the word counts
totalWordCount := 0
for count := range ch {
totalWordCount += count
}
// Output the total word count
fmt.Printf("Total word count across all URLs: %d\n", totalWordCount)
}