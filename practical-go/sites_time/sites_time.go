package main

import (
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://non-existent-site.org",
	}

	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		// wg.Add(1)
		go func(u string) {
			defer wg.Done()
			siteTime(u)
		}(url)
	}

	wg.Wait()
}

func siteTime(url string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("error: %s -> %s", url, err)
		return
	}

	defer resp.Body.Close()

	if _, err := io.Copy(io.Discard, resp.Body); err != nil {
		log.Printf("error: %s -> %s", url, err)
	}

	duration := time.Since(start)
	log.Printf("info: %s -> %s", url, duration)
}
