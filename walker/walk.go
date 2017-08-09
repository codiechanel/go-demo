package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func main() () {
   // searchDir := "C:\\Users\\admin\\go\\src\\github.com\\codiechanel\\go-demo\\walker\\cool"
searchDir := "cool"
    fileList := []string{}
    filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			
			fmt.Println("dir:", f.Name())
		} else {
			fmt.Println("file:", f.Name())
			fmt.Println("size:", f.Size())

		}
        fileList = append(fileList, path)
        return nil
	})
	
	fmt.Println("displaying ....")

    for _, file := range fileList {
        fmt.Println(file)
    }
}