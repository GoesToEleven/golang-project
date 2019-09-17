package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

//  https://godoc.org/golang.org/x/net/html

// example from
// https://godoc.org/golang.org/x/net/html#Parse

func main() {
	r, err := http.Get("https://twitter.com/Todd_McLeod/status/1169751640926146560")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	doc, err := html.Parse(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println(a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}
