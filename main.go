package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/ianchildress/platebot/throttle"

	"github.com/ianchildress/platebot/alert"
)

const (
	sleepDuration = time.Millisecond * 500
)

func main() {
	rand.Seed(time.Now().Unix())

	var sid, token, sendTo, sendFrom string
	flag.StringVar(&sid, "sid", "", "twilio sid")
	flag.StringVar(&token, "token", "", "twilio token")
	flag.StringVar(&sendTo, "to", "", "phone number to send alert message to")
	flag.StringVar(&sendFrom, "from", "", "phone number to send alert message from")
	flag.Parse()

	alerter, err := alert.NewTwilioAlerter(sid, token, sendTo, sendFrom)
	if err != nil {
		exit(err)
	}

	throttler := throttle.NewSimpleThrottler(throttle.DefaultThrottleDuration)

	for {
		start := time.Now()

		var products []Product
		products = append(products, checkPlates(plateProducts, throttler)...)
		products = append(products, checkBars(barProducts, throttler)...)

		if len(products) > 0 {
			if err := sendMsg(products, alerter); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("alert successfully sent")
			}
		}

		fmt.Println("completed search in", time.Since(start))
	}
}

func sendMsg(products []Product, alerter alert.Alerter) error {
	var msg string
	for i := range products {
		msg += fmt.Sprintf("%s %s\n", products[i].Name, products[i].URL)
	}
	return alerter.Alert(msg)
}

func exit(err error) {
	fmt.Println(err)
	os.Exit(1)
}
