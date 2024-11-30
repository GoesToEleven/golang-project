package hello

import (
	"rsc.io/quote/v3"
)

func Hello() string {
	return "Hello from a fork"
}

func Proverb() string {
	return quote.Concurrency()
}
