package webfetch

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

/*

You need to scrape data from multiple web pages concurrently. For simplicity, assume that you're only scraping the number of words from each page.
Requirements:
Use goroutines to fetch and process the pages concurrently.
Use channels to communicate the word count from each page.
Use a waitgroup to ensure all pages are processed before printing the final results.

urls := []string{"https://example.com", "https://www.golang.org", "https://www.github.com"}
*/

func CallConcurrently() {

	client := http.Client{}
	errSvc := &Err{errChn: make(chan error), done: make(chan bool)}
	go errSvc.start()

	urls := []string{"https://example.com", "https://www.golang.org", "https://www.github.com", "", "http://bad_request.co"}
	wg := &sync.WaitGroup{}	
	rspChan := make(chan int, len(urls))	
	for _, url := range urls {
		wg.Add(1)
		go fetchWorker(client, url, rspChan, errSvc.errChn, wg)
	}	
	
	wg.Wait()	
	errSvc.done <- true
	close(rspChan)

	ans := 0
	for r := range rspChan {		
		ans += r		
	}
	
	fmt.Println("Total: ", ans)
}

type Err struct {
	errChn chan error
	done  chan bool
}

func (e *Err) start() {
	for {
		select {
		case <-e.done:
			println("stopping err service")
			return
		case cErr := <-e.errChn:
			fmt.Println("err received:", cErr)
		}
	}
}

func fetchWorker(httpClient http.Client, url string, respLenOfStr chan<- int, errChan chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	rsp, err := httpClient.Get(url)
	if err != nil {
		errChan <- err
		return
	}
	defer rsp.Body.Close()
	if rsp.StatusCode != http.StatusOK {
		errChan <- fmt.Errorf("status: %v", rsp.Status)
	}

	buffer, err := io.ReadAll(rsp.Body)
	if err != nil {
		errChan <- err
		return
	}

	str := string(buffer)
	n := len(strings.Fields(str))
	fmt.Println("received...", n)
	respLenOfStr <- n	
}
