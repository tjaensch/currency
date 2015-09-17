# Golang command line script that fetches currency exchange rate data from the http://fixer.io/ JSON API

- Go has to be properly installed on your machine for this to work (https://golang.org/doc/install)

- Run the currency.go file from the /src folder of your $GOPATH like this:

	go run currency.go USD

- USD would be US Dollar as the base for the exchange rate, it can also be one of the following: AUD, CAD, CHF, EUR, NZD, RUB, JPY

The script output looks like this:

==== Currency Rates ====

> Base:	 AUD
> Date:	 2015-09-17
> US Dollar:	 0.71704
> Australian Dollar:	 0
> Canadian Dollar:	 0.94606
> Swiss Franc:	 0.69409
> Euro:	 0.63387
> Russian Ruble:	 47.237
> Japanese Yen:	 86.689
> New Zealand Dollar:	 1.1296
