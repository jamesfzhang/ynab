package api

import "github.com/jamesfzhang/ynab/model"

// List budget summaries.
// https://api.youneedabudget.com/v1#/Budgets/getBudgets
func (service *BudgetService) List() (budgets []model.BudgetSummary, err error) {

	var result model.BudgetSummaryResponse
	err = service.Client.get("/budgets", &result)
	if err != nil {
		return
	}

	budgets = result.Data.Budgets
	return
}

// Get specified budget's details.
// https://api.youneedabudget.com/v1#/Budgets/getBudgetById
func (service *BudgetService) Get(id string) (budget model.BudgetDetail, err error) {

	var result model.BudgetDetailResponse
	err = service.Client.get("/budgets/"+id, &result)
	if err != nil {
		return
	}

	budget = result.Data.Budget
	return
}

// Get specified budget's settings.
// https://api.youneedabudget.com/v1#/Budgets/getBudgetSettingsById
func (service *BudgetService) Settings(id string) (settings model.BudgetSettings, err error) {

	var result model.BudgetSettingsResponse
	err = service.Client.get("/budgets/"+id+"/settings", &result)
	if err != nil {
		return
	}

	settings = result.Data.Settings
	return
}
