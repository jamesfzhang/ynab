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

  Balance          int64 `json:"balance"`
  ClearedBalance   int64 `json:"cleared_balance"`
  UnclearedBalance int64 `json:"uncleared_balance"`

  OnBudget bool `json:"on_budget"`
  Closed   bool `json:"closed"`
  Deleted  bool `json:"deleted"`
}

func (account Account) FormattedBalance(format CurrencyFormat) string {
  return format.Render(account.Balance)
}

// FilterActive filters out any closed or deleted accounts.
func FilterActive(accounts []Account) (result []Account) {
  for _, account := range accounts {
    if !account.Closed && !account.Deleted {
      result = append(result, account)
    }
  }
  return
}
