package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(resp)
}

// reader type
// type Reader interface {
// 	Read(p []byte) (n int, err error)
// }

// *io.ReadCloser*

// io.Reader Interface
// Read([]byte)(int, error)

// io.Closer Interface
// Close() (error)