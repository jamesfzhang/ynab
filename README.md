# YNAB

This project implements an API client for [YNAB](https://api.youneedabudget.com/) in Go.

*This is a work in progress, and was mainly created to learn the language.*

## Installation

```
go get github.com/jamesfzhang/ynab
```

```go
import "github.com/jamesfzhang/ynab"
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/jamesfzhang/ynab"
)

// To use this client, create an access token for your YNAB account.
// Instructions: https://api.youneedabudget.com/#authentication

token  := "my-access-token"
client := ynab.NewClient(token)

// List budget summaries
budgets, err := client.BudgetService.List()
if err != nil {
	fmt.Println(err)
	return
}

for _, b := range budgets {
	// Get budget details
	budget, err := client.BudgetService.Get(b.Id)
	if err != nil {
		continue
	}

	// Print budget & net worth
	fmt.Printf("Budget: %v, Net Worth: %v\n\n", budget.Name, budget.FormattedNetWorth())

	for _, a := range budget.ActiveAccounts() {
		// Print balance of each account
		fmt.Printf("  %v: %v\n", a.Name, a.FormattedBalance(&budget.CurrencyFormat))
	}
}
```
