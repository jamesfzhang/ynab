package model

import "sort"

// https://api.youneedabudget.com/v1#/Accounts
type AccountsResponse struct {
	Data AccountsWrapper `json:"data"`
}

type AccountsWrapper struct {
	Accounts []Account `json:"accounts"`
}

type AccountResponse struct {
	Data AccountWrapper `json:"data"`
}

type AccountWrapper struct {
	Account Account `json:"account"`
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

func (account *Account) FormattedBalance(format *CurrencyFormat) string {
	return format.Render(account.Balance)
}

// Sorting
type ByBalance []Account

func (a ByBalance) Len() int           { return len(a) }
func (a ByBalance) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByBalance) Less(i, j int) bool { return a[i].Balance > a[j].Balance }

// FilterActive filters out any closed or deleted accounts, and sorts by
// balance in descending order.
func FilterActive(accounts *[]Account) (result []Account) {
	for _, account := range *accounts {
		if !account.Closed && !account.Deleted {
			result = append(result, account)
		}
	}
	sort.Sort(ByBalance(result))
	return
}
