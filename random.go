package main

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func random() string {
	var output strings.Builder
	//Lowercase and Uppercase Both
	charSet := "abcdedfghijklmnopqrstABCDEFGHIJKLMNOP"
	length := 5
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}
	return output.String()
}
