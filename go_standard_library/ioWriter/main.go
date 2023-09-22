/*
Writer interface
---
type Writer interface {
        Write(p []byte) (n int, err error)
}
    Writer is the interface that wraps the basic Write method.

    Write writes len(p) bytes from p to the underlying data stream. It returns
    the number of bytes written from p (0 <= n <= len(p)) and any error
    encountered that caused the write to stop early. Write must return a non-nil
    error if it returns n < len(p). Write must not modify the slice data,
    even temporarily.

    Implementations must not retain p.

var Discard Writer = discard{}
func MultiWriter(writers ...Writer) Writer

os.OpenFile flags
---
        // Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
        O_RDONLY int = syscall.O_RDONLY // open the file read-only.
        O_WRONLY int = syscall.O_WRONLY // open the file write-only.
        O_RDWR   int = syscall.O_RDWR   // open the file read-write.
        // The remaining values may be or'ed in to control behavior.
        O_APPEND int = syscall.O_APPEND // append data to the file when writing.
        O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
        O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
        O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
        O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
*/

package main

import (
	"bytes"
	"encoding/json"
	"os"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	f, err := os.OpenFile("file.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}

	// Write example
	defer f.Close()
	// f.Write([]byte("Writing some string into the file using write method\n"))

	// json encoder example
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	u := user{Name: "bob", Age: 10}

	if err = enc.Encode(u); err != nil {
		panic(err)
	}

	f.Write(buf.Bytes())
}
