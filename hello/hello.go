package main

import (
	"fmt"

	"github.com/codiechanel/go-demo/utils"
)

func main() {

	fmt.Println(utils.WorkingDir())
	fmt.Println(utils.CurrentDir("nice"))
	// utils.ParseXml("Castle.xml")

	utils.ShowOne()

	utils.ParseJSON(`{"num":6.13,"strs":["a","b"]}`)


}
