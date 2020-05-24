package main

import "fmt"

func Hello() string {
	return "Hello, world."
}

func main() {
	greeting := Hello()
	fmt.Println(greeting)
}
