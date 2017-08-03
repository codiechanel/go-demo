package main

import (
	"fmt"
	"io/ioutil"

	"github.com/codiechanel/go-demo/utils"

	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	data, err := ioutil.ReadFile("index.html")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Length", fmt.Sprint(len(data)))
	fmt.Fprint(w, string(data))
}

func vals() (int, int) {
	return 3, 7
}

func main() {

	fmt.Println(utils.WorkingDir())
	fmt.Println(utils.CurrentDir("nice"))
	utils.ShowOne()

	utils.ParseJSON(`{"num":6.13,"strs":["a","b"]}`)

	utils.ParseRSS()

	utils.StartThread()

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

}
