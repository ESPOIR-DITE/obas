package budgets

import (
	"errors"
	"obas/api"
	"obas/domain/budget"
)

const budgetURL = api.BASE_URL + "/budget/"

func CreateBudget(budgetObj budget.Budget) (budget.Budget, error) {
	entity := budget.Budget{}
	resp, _ := api.Rest().SetBody(budgetObj).Post(budgetURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetBudget(budgetId string) (budget.Budget, error) {
	entity := budget.Budget{}
	resp, _ := api.Rest().Get(budgetURL + "get/" + budgetId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetBudgets() ([]budget.Budget, error) {
	entity := []budget.Budget{}
	resp, _ := api.Rest().Get(budgetURL + "all")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteBudget(budgetObj budget.Budget) (budget.Budget, error) {
	entity := budget.Budget{}
	resp, _ := api.Rest().SetBody(budgetObj).Post(budgetURL + "delete")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func UpdateBudget(budgetObj budget.Budget) (budget.Budget, error) {
	entity := budget.Budget{}
	resp, _ := api.Rest().SetBody(budgetObj).Post(budgetURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
