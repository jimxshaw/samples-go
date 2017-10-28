package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Parse all arguments.
	flag.Parse()

	// flag.Args contains all non-flag arguments.
	sites := flag.Args()

	// Lets keep a reference to when we started.
	start := time.Now()

	processSites(sites)

	fmt.Println("Entire process took %s\n", time.Since(start))
}

func processSites(sites []string) {
	for i, site := range sites {
		// Start a timer for this request.
		begin := time.Now()

		// Retrieve the site
		resp, err := http.Get(site)

		if err != nil {
			fmt.Println(site, err)
			continue
		}

		fmt.Printf("%d) Site %q took %s to retrieve with status code of %d.\n", i, site, time.Since(begin), resp.StatusCode)
	}
}
