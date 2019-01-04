// Package ynab provides an API client to the v1 YNAB API, https://api.youneedabudget.com/v1
package ynab

import (
	"github.com/jamesfzhang/ynab/api"
)

// NewClient returns an API client with the specified access token.
func NewClient(accessToken string) *api.Client {
	return api.NewClient(accessToken)
}
