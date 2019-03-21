# Go library for accessing the Namesilo API

[![GoDoc](https://godoc.org/github.com/nrdcg/namesilo?status.svg)](https://godoc.org/github.com/nrdcg/namesilo)
[![Build Status](https://travis-ci.org/nrdcg/namesilo.svg?branch=master)](https://travis-ci.org/nrdcg/namesilo)
[![Go Report Card](https://goreportcard.com/badge/github.com/nrdcg/namesilo)](https://goreportcard.com/report/github.com/nrdcg/namesilo)

An Namesilo API client written in Go.

namesilo is a Go client library for accessing the Namesilo API.

## Example


```go
package main

import (
	"fmt"
	"log"

	"github.com/nrdcg/namesilo"
)

func main() {
	transport, err := namesilo.NewTokenTransport("1234")
	if err != nil {
		log.Fatal(err)
	}

	client := namesilo.NewClient(transport.Client())

	params := &namesilo.AddAccountFundsParams{
		Amount:    "1000000",
		PaymentID: "acbd",
	}

	funds, err := client.AddAccountFunds(params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(funds)
}
```

## API Documentation

- [API docs](https://www.namesilo.com/api_reference.php)
