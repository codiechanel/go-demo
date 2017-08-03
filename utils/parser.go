package utils

import (
	"fmt"
	// "log"
	"io/ioutil"
	"os"
	"encoding/xml"
	// "path/filepath"
)


type Query struct {
	Series Show
	// Have to specify where to find episodes since this
	// doesn't match the xml tags of the data that needs to go into it
	EpisodeList []Episode `xml:"Episode"`
}

type Show struct {
	// Have to specify where to find the series title since
	// the field of this struct doesn't match the xml tag
	Title    string `xml:"SeriesName"`
	SeriesID int
	Keywords map[string]bool
}

type Episode struct {
	SeasonNumber  int
	EpisodeNumber int
	EpisodeName   string
	FirstAired    string
}

// this two functions has something to do with the above struct 
func (s Show) String() string {
	return fmt.Sprintf("%s - %d", s.Title, s.SeriesID)
}

func (e Episode) String() string {
	return fmt.Sprintf("S%02dE%02d - %s - %s", e.SeasonNumber, e.EpisodeNumber, e.EpisodeName, e.FirstAired)
}

// ParseXml is cool
func ParseXml(s string) {
	xmlFile, err := os.Open(s)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	b, _ := ioutil.ReadAll(xmlFile)

	var q Query
	xml.Unmarshal(b, &q)

	fmt.Println(q.Series)
	for _, episode := range q.EpisodeList {
		fmt.Printf("\t%s\n", episode)
	}
}