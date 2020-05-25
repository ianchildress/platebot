package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/ianchildress/platebot/alert"
)

const (
	testNumber = "2058318618"
)

func main() {
	rand.Seed(time.Now().Unix())

	sid := flag.String("sid", "", "twilio sid")
	token := flag.String("token", "", "twilio token")
	sendTo := flag.String("send-to", "", "phone number to send alert message to")
	flag.Parse()

	switch {
	case sid == nil:
		fmt.Println("missing flag: sid")
	case token == nil:
		fmt.Println("missing flag: token")
	case sendTo == nil:
		fmt.Println("missing flag: send-to")
	}

	alerter, err := alert.NewTwilioAlerter(*sid, *token, *sendTo, testNumber)
	if err != nil {
		exit(err)
	}
	_ = alerter

	for {
		var products []Product
		products = append(products, checkSteelPlates()...)
		products = append(products, checkMachinedPlates()...)
		products = append(products, checkOlympicPlates()...)
		products = append(products, checkEchoPlates()...)
		products = append(products, checkBars()...)
		if len(products) > 0 {
			var msg string
			for i := range products {
				msg += fmt.Sprintf("%s %s\n", products[i].Name, products[i].URL)
			}
			if err := alerter.Alert(msg); err != nil {
				fmt.Println("failed to send message", err)
			}
		}
	}

}

func exit(err error) {
	fmt.Println(err)
	os.Exit(1)
}
