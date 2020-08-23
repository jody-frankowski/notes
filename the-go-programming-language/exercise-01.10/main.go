// Exercise 1.10: Find a web site that produces a large amount of data. Investigate caching by
// running fetchall twice in succession to see whether the reported time changes much. Do you get
// the same content each time? Modify fetchall to print its output to a file so it can be examined.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, arg := range os.Args[1:] {
		fmt.Println(url.PathEscape(arg))
		go fetch(arg, ch) // start a goroutine
	}
	for range os.Args[1:] {
		select {
		case err, ok := <-ch:
			if ok {
				fmt.Println(err)
			}
		}
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(uri string, ch chan<- string) {
	f, err := os.Create(url.PathEscape(uri))
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	defer f.Close()

	start := time.Now()
	resp, err := http.Get(uri)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", uri, err)
		return
	}
	secs := time.Since(start).Seconds()
	fmt.Fprintf(f, "%.2fs  %7d  %s", secs, nbytes, uri)

	f.Sync()
	close(ch)
}
