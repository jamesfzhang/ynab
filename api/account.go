package api

import "github.com/jamesfzhang/ynab/model"

// List active (not closed or deleted) accounts for specified budget.
// https://api.youneedabudget.com/v1#/Accounts/getAccounts
func (service *AccountService) List(budgetId string) (accounts []model.Account, err error) {

	var result model.AccountsResponse
	err = service.Client.get("/budgets/"+budgetId+"/accounts", &result)
	if err != nil {
		return
	}

	accounts = model.FilterActive(&result.Data.Accounts)
	return
}

// Get specified account.
// https://api.youneedabudget.com/v1#/Accounts/getAccountById
func (service *AccountService) Get(
	budgetId string,
	accountId string,
) (account model.Account, err error) {

	var result model.AccountResponse
	err = service.Client.get("/budgets/"+budgetId+"/accounts/"+accountId, &result)
	if err != nil {
		return
	}

	account = result.Data.Account
	return
}
