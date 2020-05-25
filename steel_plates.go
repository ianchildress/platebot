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
	steelPlateURL = "https://www.roguefitness.com/rogue-calibrated-lb-steel-plates"
)

var steelProducts = map[string]string{
	//"38873": "1LB Calibrated Plate",
	//"38875": "2.5LB Calibrated Plate",
	"38877": "5LB Calibrated Plate",
	"38879": "10LB Calibrated Plate",
	"38881": "25LB Calibrated Plate",
	"38883": "35LB Calibrated Plate",
	"38885": "45LB Calibrated Plate",
}

func checkSteelPlates() []Product {
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

	for k, v := range steelProducts {
		product := fmt.Sprintf("grouped-product-item-%s", k)
		if bytes.Contains(b, []byte(product)) {
			color.Green("%v %s %s", v, steelPlateURL, time.Now())
			products = append(products, Product{
				URL:    steelPlateURL,
				Name:   v,
				Number: k,
			})
		}
	}
	return products
}
