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
// To use this client, create an access token for your YNAB account.
// Instructions: https://api.youneedabudget.com/#authentication

token  := "my-access-token"
client := ynab.NewClient(token)
```