package utils

import (
	"encoding/json"
	"fmt"
)

// ShowOne is cool
func ShowOne() {
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)

	fmt.Println(string(mapB))
}

// ParseJSON is cool
func ParseJSON(s string) {
	byt := []byte(s)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	fmt.Println("num is ", dat["num"])
}
