package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/ianchildress/platebot/throttle"

	"github.com/fatih/color"
	"golang.org/x/net/html"
)

func checkBars(bars []Product, throttler throttle.Throttler) []Product {
	var products []Product

	for i := range bars {
		if !bars[i].Track {
			continue
		}
		r, err := throttler.Get(bars[i].URL)
		if err != nil {
			fmt.Println(bars[i].URL, err)
			continue
		}

		if checkAddToCart(r) {
			color.Green("Found %s %s", bars[i].Name, bars[i].URL)
			products = append(products, Product{
				URL:  bars[i].URL,
				Name: bars[i].Name,
			})
		}

		r.Close()
	}
	return products
}

func checkAddToCart(r io.Reader) bool {
	z := html.NewTokenizer(r)
	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return false
		case tt == html.StartTagToken:
			t := z.Token()
			// check if there is a button that allows us to add to cart
			if t.Data == "button" {
				for _, a := range t.Attr {
					if a.Key == "class" && strings.Contains(a.Val, "btn-add-to-cart") {
						return true
					}
				}
			}
		}
	}
}
