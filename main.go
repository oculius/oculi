package main

import "fmt"

func main() {
	const (
		green  = "\033[97;42m"
		yellow = "\033[97;33m"
		reset  = "\033[0m"
	)
	fmt.Printf("%s %s %s %s %s\n", green, "haiii", reset, yellow, "hello")
}
