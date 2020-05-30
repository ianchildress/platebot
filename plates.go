package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/fatih/color"
	"github.com/ianchildress/platebot/throttle"
)

func checkPlates(plates []Product, throttler throttle.Throttler) []Product {
	var products []Product

	exists := func(s string, l []string) bool {
		for i := range l {
			if l[i] == s {
				return true
			}
		}
		return false
	}

	// get a list of all product urls so we only need to request a page once
	var urls []string
	for i := range plates {
		if !exists(plates[i].URL, urls) {
			urls = append(urls, plates[i].URL)
		}
	}

	for i := range urls {
		//fmt.Println("checking", urls[i])
		r, err := throttler.Get(urls[i])
		if err != nil {
			fmt.Println(err)
			continue
		}

		b, err := ioutil.ReadAll(r)
		if err != nil {
			fmt.Println(err)
			continue
		}
		r.Close()

		for _, p := range plates {
			// skip untracked products
			if !p.Track {
				//fmt.Println("skipping", p.Name)
				continue
			}

			tag := fmt.Sprintf("grouped-product-item-%s", p.Number)
			if bytes.Contains(b, []byte(tag)) {
				color.Green("%v %s?=%s", p.Name, p.URL, random())
				products = append(products, p)
			}
		}

	}
	return products
}
