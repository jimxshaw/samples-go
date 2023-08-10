package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// With .gz suffix
	sig, err := sha1Sum("http.log.gz")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Println(sig)

	// Without .gz suffix
	sig, err = sha1Sum("sha1.go")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Println(sig)
}

// if file name ends in .gz -> $ cat http.log.gz | gunzip | sha1sum
// else -> $ cat http.log.gz | sha1sum
func sha1Sum(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	// idiom: acquire a resource, check for error, defer release.
	// defer will only execute when the function exits,
	// regardless of how it exits, normally or with error.
	// If multiple defers exist in a function then they'll
	// be called in reverse order, like a stack.
	// defer should always be at the function level.
	defer file.Close()

	var r io.Reader = file

	if strings.HasSuffix(fileName, ".gz") {
		gzFile, err := gzip.NewReader(file)
		if err != nil {
			return "", err
		}
		defer gzFile.Close()
		r = gzFile
	}

	w := sha1.New()

	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}

	sig := w.Sum(nil)

	// io.CopyN(os.Stdout, r, 200)

	return fmt.Sprintf("%x", sig), nil
}
