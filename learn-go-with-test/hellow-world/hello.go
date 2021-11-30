package main

import "fmt"

func main() {
	fmt.Println(Hello("world"))
}

func Hello(user string) string {
	return fmt.Sprintf("Hello, %s!", user)
}
