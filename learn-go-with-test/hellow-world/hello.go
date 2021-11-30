package main

import "fmt"

func main() {
	fmt.Println(Hello("world"))
}

const englishHelloPrefix = "Hello, "

func Hello(user string) string {
	if user == "" {
		user = "World"
	}
	return englishHelloPrefix + user + "!"
}
