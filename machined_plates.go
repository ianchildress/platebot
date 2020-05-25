package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/fatih/color"
)

const (
	machinedPlateURL = "https://www.roguefitness.com/rogue-machined-olympic-plates"
)

var machinedProducts = map[string]string{
	//"47205": "2.5LB Machined Olympic",
	"47211": "5LB Machined Olympic",
	"47213": "10LB Machined Olympic",
	"47227": "25LB Machined Olympic",
	"47229": "35LB Machined Olympic",
	"47219": "45LB Machined Olympic",
}

func checkMachinedPlates() []Product {
	var products []Product
	resp, err := http.Get(machinedPlateURL)
	if err != nil {
		exit(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		exit(err)
	}

	for k, v := range machinedProducts {
		product := fmt.Sprintf("grouped-product-item-%s", k)
		if bytes.Contains(b, []byte(product)) {
			color.Green("%v %s", v, time.Now())
			products = append(products, Product{
				URL:    steelPlateURL,
				Name:   v,
				Number: k,
			})
		}
	}
	return products
}
