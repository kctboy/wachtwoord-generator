package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	upCaseSet      = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowCaseSet     = "abcdedfghijklmnopqrstuvwxyz"
	numberSet      = "0123456789"
	specialCharSet = "!@#$%&*"
	allSet         = lowCaseSet + upCaseSet + specialCharSet + numberSet
)

var minUpCase int
var minNumber int
var minSpecial int
var passwordLenght int

func init() {
	flag.IntVar(&minUpCase, "u", 1, "The minimu. uppercase letters")
	flag.IntVar(&minSpecial, "s", 1, "The minimum special characters")
	flag.IntVar(&minNumber, "n", 1, "The minimum numbers")
	flag.IntVar(&passwordLenght, "l", 5, "The lenght of the password")
}
func main() {

	rand.Seed(time.Now().Unix())

	password := generatePassword(minNumber, minSpecial, minUpCase, passwordLenght)
	fmt.Println(password)
}

func generatePassword(passwordLenght, minSpecial, minUpCase, minNumber int) string {
	var password strings.Builder

	for i := 0; i < minNumber; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	for i := 0; i < minNumber; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))

	}

	for i := 0; i < minUpCase; i++ {
		random := rand.Intn(len(upCaseSet))
		password.WriteString(string(upCaseSet[random]))
	}

	remainingLengt := passwordLenght - minNumber - minSpecial - minUpCase
	for i := 0; i < remainingLengt; i++ {
		random := rand.Intn(len(allSet))
		password.WriteString(string(allSet[random]))
	}

	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)

}
