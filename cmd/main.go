package main

import "fmt"

func DomainHelloWorld(val string) string {
	return fmt.Sprintf("Hello, %s", val)
}

func main() {
	fmt.Println(DomainHelloWorld("world"))
}
