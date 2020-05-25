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
	olympicPlateURL = "https://www.roguefitness.com/rogue-olympic-plates"
)

var olympicProducts = map[string]string{
	//"38873": "1LB Calibrated Plate",
	//"7185":  "2.5LB Olympic Plate - Pair",
	"7183": "5LB Olympic Plate - Pair",
	"7181": "10LB Olympic Plate - Pair",
	"7179": "25LB Olympic Plate - Pair",
	"7177": "35LB Olympic Plate - Pair",
	"7175": "45LB Olympic Plate - Pair",
}

func checkOlympicPlates() []Product {
	var products []Product
	resp, err := http.Get(steelPlateURL)
	if err != nil {
		exit(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		exit(err)
	}

	for k, v := range olympicProducts {
		product := fmt.Sprintf("grouped-product-item-%s", k)
		if bytes.Contains(b, []byte(product)) {
			color.Green("%v %s %s", v, olympicPlateURL, time.Now())
			products = append(products, Product{
				URL:    steelPlateURL,
				Name:   v,
				Number: k,
			})
		}
	}
	return products
}
