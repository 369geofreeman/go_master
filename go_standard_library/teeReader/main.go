package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
)

// http://storage.googleapis.com/books/ngrams/books/googlebooks-eng-all-5gram-20120701-0.gz

// Objective
// ---
// :reader:	 Download the file
// :writer:	 Progress counter: Measure in real time the number of mb downloaded
// :writer:  Write the file into our local file system
// :writer:  Write the file into our archive

type counter struct {
	total uint64
}

// write
func (c *counter) Write(b []byte) (int, error) {
	c.total += uint64(len(b)) //32 kb at a time
	progress := float64(c.total) / (1024 * 1024)
	fmt.Printf("\rDownloading %f MB... \n", progress)
	return len(b), nil
}

// TeeReader
// ---
// r1: <file from the internet>
// w1: <progress counter>
// r2 := TeeReader(r1, w1)
// r2.Read()
// 		-- does the following
// 			- r1.read()
// 			- w1.Write(b)
// 			= return

func main() {
	res, err := http.Get("http://storage.googleapis.com/books/ngrams/books/googlebooks-eng-all-5gram-20120701-0.gz")
	if err != nil {
		panic(err)
	}

	// Download the file into our local filesystem
	local, err := os.OpenFile("download-5gram.txt", os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer local.Close()

	dec, err := gzip.NewReader(res.Body)
	if err != nil {
		panic(err)
	}

	// Copy res.Body into local
	_, err = io.Copy(local, io.TeeReader(dec, &counter{}))
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
}
