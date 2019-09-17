package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	r, err := http.Get("https://twitter.com/Todd_McLeod/status/1169751640926146560")
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", b)
}
