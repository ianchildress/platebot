package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/fatih/color"
	"golang.org/x/net/html"
)

type BarProduct struct {
	URL     string
	Name    string
	InStock bool
}

var barProducts = []BarProduct{
	//{
	//	URL:  "https://www.roguefitness.com/rogue-t-2-5kg-technique-bar",
	//	Name: "Rogue T-2.5KG Technique Bar",
	//},
	//{
	//	URL:  "https://www.roguefitness.com/rogue-bella-bar-2-0-black-zinc",
	//	Name: "The Bella Bar 2.0 - Black Zinc",
	//},
	{
		URL:  "https://www.roguefitness.com/rogue-45lb-ohio-power-bar-black-zinc",
		Name: "Rogue 45LB Ohio Power Bar - Black Zinc",
	},
	{
		URL:  "https://www.roguefitness.com/rogue-45lb-ohio-power-bar-bare-steel",
		Name: "Rogue 45LB Ohio Power Bar - Bare Steel",
	},
	{
		URL:  "https://www.roguefitness.com/rogue-45lb-ohio-powerlift-bar-cerakote",
		Name: "Rogue Ohio Power Bar - Cerakote",
	},
	{
		URL:  "https://www.roguefitness.com/rogue-45lb-ohio-power-bar-stainless",
		Name: "Rogue 45LB Ohio Power Bar - Stainless Steel",
	},
	{
		URL:  "https://www.roguefitness.com/rogue-ohio-power-bar-e-coat",
		Name: "Rogue 45LB Ohio Power Bar - E-Coat",
	},
	{
		URL:  "https://www.roguefitness.com/rogue-echo-bar",
		Name: "Rogue Echo Bar 2.0",
	},
	{
		URL:  "https://www.roguefitness.com/rogue-ohio-bar-black-oxide",
		Name: "The Ohio Bar - Black Oxide",
	},
	{
		URL:  "https://www.roguefitness.com/the-ohio-bar-cerakote",
		Name: "The Ohio Bar - Cerakote",
	},
	{
		URL:  "https://www.roguefitness.com/the-ohio-bar-black-zinc",
		Name: "The Ohio Bar - Black Zinc",
	},
	{
		URL:  "https://www.roguefitness.com/rogue-ohio-bar-black-oxide",
		Name: "The Ohio Bar - Black Oxide",
	},
	{
		URL:  "https://www.roguefitness.com/rogue-ohio-power-bar-e-coat",
		Name: "Rogue 45LB Ohio Power Bar - E-Coat",
	},
	{
		URL:  "https://www.roguefitness.com/the-rogue-bar-2-0",
		Name: "The Rogue Bar 2.0",
	},
	//{
	//	URL:  "",
	//	Name: "",
	//},
	//{
	//	URL:  "",
	//	Name: "",
	//},
	//{
	//	URL:  "",
	//	Name: "",
	//},
}

func checkBars() []Product {
	var products []Product

	for i := range barProducts {
		// throttle
		time.Sleep(time.Millisecond * 100)
		resp, err := http.Get(barProducts[i].URL)
		if err != nil {
			fmt.Println(barProducts[i].URL, err)
			continue
		}
		defer resp.Body.Close()

		if barInStock(resp.Body) {
			color.Green("Found %s %s", barProducts[i].Name, barProducts[i].URL)
			products = append(products, Product{
				URL:    barProducts[i].URL,
				Name:   barProducts[i].Name,
				Number: "",
			})
		}
	}
	return products
}

func barInStock(r io.Reader) bool {
	z := html.NewTokenizer(r)
	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return false
		case tt == html.StartTagToken:
			t := z.Token()
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
