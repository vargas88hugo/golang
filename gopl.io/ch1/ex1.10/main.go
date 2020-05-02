package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string) // channel of strings

	for _, uri := range os.Args[1:] {
		go fetch(uri, ch) // start a goroutine and call fetch asynchronously
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(uri string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(uri)

	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch, Sprint return instead to send stdout
		return
	}

	f, err := os.Create(url.QueryEscape(uri))

	if err != nil {
		ch <- err.Error()
	}

	nbytes, err := io.Copy(f, resp.Body)
	resp.Body.Close()

	if closeErr := f.Close(); err == nil {
		err = closeErr
	}

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", uri, err) // return instead to send stdout
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, uri)
}
