package main

import (
	"fmt"
	"os"

	"github.com/v.godlevskiy/tdd-app/internal/countdown"
)

const (
	englishHelloPrefix = "Hello"
	spanishHelloPrefix = "Hola"
)

func DomainHelloWorld(val string, lng string) string {
	prefix := englishHelloPrefix

	if val == "" {
		val = "World"
	}

	if lng == "Spanish" {
		prefix = spanishHelloPrefix
	}

	return fmt.Sprintf("%s, %s", prefix, val)
}

func main() {
	countdown.CountDown(os.Stdout, &countdown.DefaultSleeper{})
	// fmt.Println(DomainHelloWorld("world", "Spanish"))
}
