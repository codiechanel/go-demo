package utils

import (
	// "fmt"
	"log"
	"os"
	"path/filepath"
)

// CurrentDir is cool
func CurrentDir(s string) string {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(dir)
	return dir
}

// WorkingDir is cool
func WorkingDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(dir)
	return dir

}
