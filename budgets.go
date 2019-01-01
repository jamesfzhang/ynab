package ynab

type BudgetSummaryService Service

type BudgetSummaryResponse struct {
	Data BudgetSummaryWrapper `json:"data"`
}

type BudgetSummaryWrapper struct {
	Budget  BudgetSummary   `json:"budget"`
	Budgets []BudgetSummary `json:"budgets"`
}

type BudgetSummary struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// List budgets summaries.
func (service *BudgetSummaryService) List() (budgets []BudgetSummary, err error) {

	var result BudgetSummaryResponse
	err = service.Client.Get("/budgets", &result)

	if err != nil {
		return
	}

	budgets = result.Data.Budgets
	return
}

// Get the specified budget's summary.
func (service *BudgetSummaryService) Get(id string) (budget BudgetSummary, err error) {

	var result BudgetSummaryResponse
	err = service.Client.Get("/budgets/"+id, &result)

	if err != nil {
		return
	}

	budget = result.Data.Budget
	return
}
