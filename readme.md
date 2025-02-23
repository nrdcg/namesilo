# Go library for accessing the Namesilo API

[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/mrehanabbasi/namesilo-go.svg)](https://github.com/mrehanabbasi/namesilo-go/releases)
[![Build Status](https://github.com/mrehanabbasi/namesilo-go/workflows/Main/badge.svg?branch=master)](https://github.com/mrehanabbasi/namesilo-go/actions)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/mrehanabbasi/namesilo-go)](https://pkg.go.dev/github.com/mrehanabbasi/namesilo-go)

A Namesilo API client written in Go.

namesilo is a Go client library for accessing the Namesilo API.

## Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/mrehanabbasi/namesilo-go"
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
