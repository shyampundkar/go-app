package main

import (
	"errors"
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	http.HandleFunc("/calculate/", func(rw http.ResponseWriter, r *http.Request) {
		operation := r.URL.Query().Get("operation")
		operand1, _ := strconv.Atoi(r.URL.Query().Get("operand1"))
		operand2, _ := strconv.Atoi(r.URL.Query().Get("operand2"))
		result, err := Calculate(operation, operand1, operand2)

		if err == nil {
			rw.Header().Set("Content-Type", "text/html")
			fmt.Fprintf(rw, "<h1> The result is:%d</h1> ", result)
		}
	})
	log.Fatal(http.ListenAndServe("localhost:8010", nil))
}

func Calculate(operation string, arg1, arg2 int) (int, error) {
	switch operation {
	case "+":
		{
			return arg1 + arg2, nil
		}

	case "-":
		{
			return arg1 - arg2, nil
		}

	case "*":
		{
			return arg1 * arg2, nil
		}

	case "/":
		{
			return arg1 / arg2, nil
		}
	default:
		return 0, errors.New("invalid operator")
	}
}
