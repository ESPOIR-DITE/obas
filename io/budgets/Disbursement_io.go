package budgets

import (
	"errors"
	"obas/api"
	"obas/domain/budget"
)

const disbursementURL = api.BASE_URL + "/disbursement/"

func CreateDisbursement(disbursement budget.Disbursement) (budget.Disbursement, error) {
	entity := budget.Disbursement{}

	resp, _ := api.Rest().SetBody(disbursement).Post(disbursementURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteDisbursement(disbursement budget.Disbursement) (budget.Disbursement, error) {
	entity := budget.Disbursement{}

	resp, _ := api.Rest().SetBody(disbursement).Post(disbursementURL + "delete")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateDisbursement(disbursement budget.Disbursement) (budget.Disbursement, error) {
	entity := budget.Disbursement{}

	resp, _ := api.Rest().SetBody(disbursement).Post(disbursementURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetDisbursement(disbursementId string) (budget.Disbursement, error) {
	entity := budget.Disbursement{}
	resp, _ := api.Rest().Get(disbursementURL + "get/" + disbursementId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetDisbursements(disbursementId string) ([]budget.Disbursement, error) {
	entity := []budget.Disbursement{}
	resp, _ := api.Rest().Get(disbursementURL + "get/" + disbursementId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
