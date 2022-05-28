package main

import (
	"fmt"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jimxshaw/samples-go/romanserver/data"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPath := strings.Split(r.URL.Path, "/")

		if urlPath[1] == "roman-numeral" {
			num, _ := strconv.Atoi(strings.TrimSpace(urlPath[2]))

			if num >= 1 && num <= 10 {
				fmt.Fprintf(w, "%q", html.EscapeString(data.Numberals[num]))
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

	s.ListenAndServer()
}
