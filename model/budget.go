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

type BudgetResponse struct {
  Data BudgetWrapper `json:"data"`
}

type BudgetWrapper struct {
  Budget Budget `json:"budget"`
}

type Budget struct {
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
