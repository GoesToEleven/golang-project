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
		if n.Type == html.TextNode {
			fmt.Println(n.Data)
			// for _, a := range n.Attr {
			// 	if a.Key == "class" {
			// 		xs := strings.Split(a.Val, " ")
			// 		for _, s := range xs {
			// 			switch s {
			// 				// fullname show-popup-with-id u-textTruncate
			// 			case "fullname":
			// 				fmt.Println(a.Data)
			// 				// username u-dir u-textTruncate
			// 			case "username":
			// 				fmt.Println(a.Data)
			// 				// TweetTextSize TweetTextSize--jumbo js-tweet-text tweet-text
			// 			case "tweet-text":
			// 				fmt.Println(a.Data)
			// 			}
			// 			break
			// 		}
			// 	}
			// }
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}
