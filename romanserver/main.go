package main

import (
	"fmt"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var numberals = map[int]string{
	1:  "I",
	2:  "II",
	3:  "III",
	4:  "IV",
	5:  "V",
	6:  "VI",
	7:  "VII",
	8:  "VIII",
	9:  "IX",
	10: "X",
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPath := strings.Split(r.URL.Path, "/")

		if urlPath[1] == "roman-numeral" {
			num, _ := strconv.Atoi(strings.TrimSpace(urlPath[2]))

			if num >= 1 && num <= 10 {
				fmt.Fprintf(w, "%q", html.EscapeString(numberals[num]))
			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - Not Found"))
			}

		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
		}
	})

	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
