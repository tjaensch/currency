package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Rates struct {
	Base       string `json:"base"`
	Date       string `json:"date"`
	Currencies struct {
		AUD float64 `json:"AUD"`
		CAD float64 `json:"CAD"`
		CHF float64 `json:"CHF"`
		EUR float64 `json:"EUR"`
		NZD float64 `json:"NZD"`
		RUB float64 `json:"RUB"`
		JPY float64 `json:"JPY"`
		USD float64 `json:"USD"`
	} `json:"rates"`
}

var (
	current  string
	err      error
	rates    Rates
	response *http.Response
	body     []byte
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Please specify start currency, e.g. 'go run currency.go USD' (Available values AUD/CAD/CHF/EUR/NZD/RUB/JPY/USD)")
		os.Exit(1)
	}
	// Provide base currency
	current = args[0]

	// Use api.fixer.io to get a JSON response
	response, err = http.Get("http://api.fixer.io/latest?base=" + current)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	// Read the data into a byte slice(string)
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	// Unmarshal the JSON byte slice to a predefined struct
	err = json.Unmarshal(body, &rates)
	if err != nil {
		fmt.Println(err)
	}

	// Everything accessible in struct now
	fmt.Println("\n==== Currency Rates ====\n")

	fmt.Println("Base:\t", rates.Base)
	fmt.Println("Date:\t", rates.Date)
	fmt.Println("US Dollar:\t", rates.Currencies.USD)
	fmt.Println("Australian Dollar:\t", rates.Currencies.AUD)
	fmt.Println("Canadian Dollar:\t", rates.Currencies.CAD)
	fmt.Println("Swiss Franc:\t", rates.Currencies.CHF)
	fmt.Println("Euro:\t", rates.Currencies.EUR)
	fmt.Println("Russian Ruble:\t", rates.Currencies.RUB)
	fmt.Println("Japanese Yen:\t", rates.Currencies.JPY)
	fmt.Println("New Zealand Dollar:\t", rates.Currencies.NZD)
}
