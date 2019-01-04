package model

// https://api.youneedabudget.com/v1#/Budgets
type BudgetSummaryResponse struct {
	Data BudgetSummaryWrapper `json:"data"`
}

type BudgetSummaryWrapper struct {
	Budgets []BudgetSummary `json:"budgets"`
}

type BudgetSummary struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	LastModifiedDate string `json:"last_modified_date`

	DateFormat     DateFormat     `json:"date_format"`
	CurrencyFormat CurrencyFormat `json:"currency_format"`
}

type BudgetDetailResponse struct {
	Data BudgetDetailWrapper `json:"data"`
}

type BudgetDetailWrapper struct {
	Budget BudgetDetail `json:"budget"`
}

type BudgetDetail struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	LastModifiedDate string `json:"last_modified_date`

	DateFormat     DateFormat     `json:"date_format"`
	CurrencyFormat CurrencyFormat `json:"currency_format"`

	Accounts []Account `json:"accounts"`
}

type BudgetSettingsResponse struct {
	Data BudgetSettingsWrapper `json:"data"`
}

type BudgetSettingsWrapper struct {
	Settings BudgetSettings `json:"settings"`
}

type BudgetSettings struct {
	DateFormat     DateFormat     `json:"date_format"`
	CurrencyFormat CurrencyFormat `json:"currency_format"`
}

// ActiveAccounts returns the budget's active (not closed or deleted) accounts.
func (budget *BudgetDetail) ActiveAccounts() []Account {
	return FilterActive(&budget.Accounts)
}

// NetWorth returns the net worth of the budget (sum of balances across active accounts).
func (budget *BudgetDetail) NetWorth() (val int64) {
	for _, account := range budget.ActiveAccounts() {
		val += account.Balance
	}
	return
}

func (budget *BudgetDetail) FormattedNetWorth() string {
	return budget.CurrencyFormat.Render(budget.NetWorth())
}
