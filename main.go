package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
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

	if len(msg) < 1 {

		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {

			reader := bufio.NewReader(os.Stdin)

			if b, err := ioutil.ReadAll(reader); err != nil {
				panic(err)
			} else {
				msg = strings.TrimSpace(string(b))
			}

		}

		if len(msg) < 1 {
			fmt.Println("No Message!")
			return
		}

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
	msg = sanitizeInput(msg)

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

func sanitizeInput(msg string) string {

	// Strip ANSI Color Codes
	const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"
	ansiReg := regexp.MustCompile(ansi)

	// Strip non-alpha characters
	alphaReg := regexp.MustCompile("[^a-zA-Z]+")

	msg = ansiReg.ReplaceAllString(msg, "")
	return alphaReg.ReplaceAllString(msg, "")

}
