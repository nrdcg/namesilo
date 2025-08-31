# Go library for accessing the Namesilo API

[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/nrdcg/namesilo.svg)](https://github.com/nrdcg/namesilo/releases)
[![Build Status](https://github.com/nrdcg/namesilo/workflows/Main/badge.svg?branch=master)](https://github.com/nrdcg/namesilo/actions)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/nrdcg/namesilo)](https://pkg.go.dev/github.com/nrdcg/namesilo)

A Namesilo API client written in Go.

namesilo is a Go client library for accessing the Namesilo API.

## Examples

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
	
	// Set the endpoint to use the OTE endpoint.
	endpoint, err := GetEndpoint(false, true)
	if err != nil {
		log.Fatal(err)
	}
	
	client.Endpoint = endpoint 

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
