### Deprecated: script no longer works because the FIXER API has changed

## Golang command line script that fetches currency exchange rate data from the http://fixer.io/ JSON API

- Go has to be properly installed on your machine for this to work (https://golang.org/doc/install)

- Run the currency.go file from the /src folder of your $GOPATH like this:

	go run currency.go USD

- USD would be US Dollar as the base for the exchange rate, it can also be one of the following: AUD, CAD, CHF, EUR, NZD, RUB, JPY

The script command line output looks like this:

![Alt text](/screenshot.png?raw=true "Script output")