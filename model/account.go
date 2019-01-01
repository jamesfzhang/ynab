package model

// https://api.youneedabudget.com/v1#/Accounts
type AccountResponse struct {
  Data AccountWrapper `json:"data"`
}

type AccountWrapper struct {
  Account  Account   `json:"account"`
  Accounts []Account `json:"accounts"`
}

type Account struct {
  Id              string `json:"id"`
  Name            string `json:"name"`
  Type            string `json:"type"`
  Note            string `json:"note"`
  TransferPayeeId string `json:"transfer_payee_id"`

  Balance          int `json:"balance"`
  ClearedBalance   int `json:"cleared_balance"`
  UnclearedBalance int `json:"uncleared_balance"`

  OnBudget bool `json:"on_budget"`
  Closed   bool `json:"closed"`
  Deleted  bool `json:"deleted"`
}
