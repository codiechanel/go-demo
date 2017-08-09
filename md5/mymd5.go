package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

func ComputeMd5(filePath string) ([]byte, error) {
	start := time.Now()
	var result []byte
	file, err := os.Open(filePath)
	if err != nil {
		return result, err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return result, err
	}
	elapsed := time.Since(start)
	fmt.Printf("elapsed %s\n", elapsed)

	return hash.Sum(result), nil
}

var str string

func main() {
	num_args := len(os.Args)

	// check if received any command line arguments
	if num_args < 2 {
		fmt.Println(">> No args passed in what?")
	}
	flag.StringVar(&str, "str", "Castle.xml", "text description")
	flag.Parse()
	fmt.Println("str passed", str)
	if b, err := ComputeMd5(str); err != nil {
		fmt.Printf("Err: %v", err)
	} else {
		fmt.Printf("main.go md5 checksum is: %x", b)
	}
}
