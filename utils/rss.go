package utils

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

// ParseRSS is cool
func ParseRSS() {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("http://feeds.twit.tv/twit.xml")
	fmt.Println(feed.Title)

	for index, element := range feed.Items {
		// index is the index where we are
		// element is the element from someSlice for where we are
		fmt.Println(index, element.Title)
	}

}
