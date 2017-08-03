package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/cheggaaa/pb"
	"time"

	"github.com/codiechanel/go-demo/utils"
)

func main() {

	fmt.Println(utils.WorkingDir())
	fmt.Println(utils.CurrentDir("nice"))

	// utils.ParseXml("Castle.xml")

	// var cliUrl     = "https://google.com"

	/*
		check status and CheckRedirect
	*/
	checkStatus := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	/*
		Get Response: 200 OK?
	*/
	// response, err := checkStatus.Get("http://www.colorado.edu/conflict/peace/download/peace.zip")
	response, err := checkStatus.Get("http://ipv4.download.thinkbroadband.com/5MB.zip")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer response.Body.Close()
	fmt.Printf("Request Status: %s\n\n", response.Status) // Example: 200 OK
	filesize := response.ContentLength
	fmt.Printf("filesize %d\n\n", filesize) // Example: 200 OK

	// go Copy(response, filesize,"cool.zip")

	var fileName = "cool.zip"

		file, err := os.Create("cool.zip")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
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
		bar := pb.StartNew(countSize) // start new progressbar
		var fi os.FileInfo // get file information from os
		for fi == nil || fi.Size() < filesize { // for like while
			fi, _ = file.Stat() // File status
			bar.Set(int(fi.Size() / 1000)) // File size / 1000
			time.Sleep(time.Millisecond) // wait millisecond
		}
		finishMessage := fmt.Sprintf("\n%s with %v bytes downloaded",
		 fileName, filesize)
		bar.FinishPrint(finishMessage) // finished messages

		if err != nil {
			panic(err)
		}

}

func Copy(response *http.Response, filesize int64, fileName string) {

	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()
	n, err := io.Copy(file, response.Body)
	if n != filesize {
		fmt.Println("Truncated")
	}
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	countSize := int(filesize / 1000)
		bar := pb.StartNew(countSize) // start new progressbar
		var fi os.FileInfo // get file information from os
		for fi == nil || fi.Size() < filesize { // for like while
			fi, _ = file.Stat() // File status
			bar.Set(int(fi.Size() / 1000)) // File size / 1000
			time.Sleep(time.Millisecond) // wait millisecond
		}
		finishMessage := fmt.Sprintf("\n%s with %v bytes downloaded",
		 fileName, filesize)
		bar.FinishPrint(finishMessage) // finished messages

		if err != nil {
			panic(err)
		}

}
