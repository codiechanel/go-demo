package utils


import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}


// StartThread is cool
func StartThread() {
	go say("world")
	say("hello")

}