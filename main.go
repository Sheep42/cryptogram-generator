package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {

	var shift int
	var msg string

	flag.IntVar(&shift, "shift", 1, "An integer value to shift by, between 1 and 25")
	flag.StringVar(&msg, "msg", "", "The message you want to encrypt")
	flag.Parse()

	if shift < 1 || shift > 25 {
		fmt.Println("Incorrect shift value")
		return
	}

	if msg == "" {
		fmt.Println("No Message!")
		return
	}

	key := generateKey(shift)
	fmt.Println(encrypt(key, msg))

}

func generateKey(n int) map[rune]rune {

	var key = map[rune]rune{}
	alpha := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	count := 0

	for _, c := range alpha {

		key[c] = []rune(alpha)[((count + n) % len(alpha))]
		count += 1

	}

	return key

}

func encrypt(key map[rune]rune, msg string) string {

	msg = strings.ToUpper(msg)
	var newMsg []rune

	for _, c := range msg {

		// consciously strip chars that can't be encrypted
		if val, ok := key[c]; ok {
			newMsg = append(newMsg, val)
		}

	}

	encrypted := string(newMsg)

	return encrypted

}
