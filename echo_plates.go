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
	echoPlateURL = "https://www.roguefitness.com/rogue-echo-bumper-plates-with-white-text"
)

var echoProducts = map[string]string{
	"65856": "10LB Echo Bumper Plate",
	"65858": "15LB Echo Bumper Plate",
	"65860": "25LB Echo Bumper Plate",
	"65862": "35LB Echo Bumper Plate",
	"65864": "45LB Echo Bumper Plate",
}

func checkEchoPlates() []Product {
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

	for k, v := range echoProducts {
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
