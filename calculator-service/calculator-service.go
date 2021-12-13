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

	http.HandleFunc("/calculator", func(rw http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {

			err := r.ParseForm()
			if err != nil {
				rw.WriteHeader(http.StatusBadRequest)
				return
			}
			op1, err_op1 := strconv.Atoi(r.FormValue("operand1"))

			if err_op1 != nil {
				log.Fatal(err_op1)
				return
			}

			op2, err_op2 := strconv.Atoi(r.FormValue("operand2"))
			if err_op2 != nil {
				log.Fatal(err_op2)
				return
			}
			result, err := Calculate(r.FormValue("operation"), op1, op2)

			if err == nil {
				fmt.Fprintf(rw, "The result is: %d", result)
			}

			return

		}

		http.ServeFile(rw, r, "calculator.html")

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
