package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/cheggaaa/pb"

	"github.com/codiechanel/go-demo/utils"
)

func main() {

	fmt.Println(utils.WorkingDir())
	fmt.Println(utils.CurrentDir("nice"))

	response, filesize := Request("http://ipv4.download.thinkbroadband.com/5MB.zip")
	// response, filesize := Request("http://www.colorado.edu/conflict/peace/download/peace.zip")
	defer response.Body.Close()

	var fileName = "cool.zip"
	file, err := Create(fileName)
	defer file.Close()

	go func() {
		n, err := io.Copy(file, response.Body)
		if n != filesize {
			fmt.Println("Truncated")
		}
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
	}()

	countSize := int(filesize / 1000)
	fmt.Println("countSize", countSize)
	bar := pb.StartNew(countSize)           // start new progressbar
	var fi os.FileInfo                      // get file information from os
	for fi == nil || fi.Size() < filesize { // for like while
		fi, _ = file.Stat()            // File status
		bar.Set(int(fi.Size() / 1000)) // File size / 1000
		time.Sleep(time.Millisecond)   // wait millisecond
	}
	finishMessage := fmt.Sprintf("\n%s with %v bytes downloaded",
		fileName, filesize)
	bar.FinishPrint(finishMessage) // finished messages

	if err != nil {
		panic(err)
	}

}

// Progress is cool
func Progress(s string) {

}

// Create is cool
func Create(s string) (*os.File, error) {
	file, err := os.Create(s)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	// defer file.Close()
	return file, err

}

// Request is cool
func Request(s string) (*http.Response, int64) {
	/*
		check status and CheckRedirect
	*/
	checkStatus := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	response, err := checkStatus.Get(s)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	// defer response.Body.Close()
	fmt.Printf("Request Status: %s\n\n", response.Status) // Example: 200 OK
	filesize := response.ContentLength
	fmt.Printf("filesize %d\n\n", filesize) // Example: 200 OK

	return response, filesize

}
