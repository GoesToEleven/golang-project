package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// CLIENTS make REQUESTS

// you can create a new client, and then have that client make GET a request and receive a RESPONSE
// var DefaultClient = &Client{}
// func (c *Client) Get(url string) (resp *Response, err error)

// if you want to add DETAILS ot the request
// create a new request first
// func NewRequest(method, url string, body io.Reader) (*Request, error)
// then create a client
// var DefaultClient = &Client{}
// then have the client DO that request
// func (c *Client) Do(req *Request) (*Response, error)

// if you want to keep it easy, use this function
// which is a wrapper function
/*
func Get(url string) (resp *Response, err error) {
	return DefaultClient.Get(url)
}
*/
// here is the func
// func Get(url string) (resp *Response, err error)
// https://golang.org/src/net/http/client.go#L369

// ONCE YOU GET YOUR URL, YOU WILL HAVE A RESPONSE
// https://godoc.org/net/http#Response

// this is the example from
// https://godoc.org/net/http#example-Get

func main() {
	r, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
