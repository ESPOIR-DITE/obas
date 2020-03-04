package budgets

import (
	"errors"
	"obas/api"
	"obas/domain/budget"
)

const disbursementTypeURL = api.BASE_URL + "/disbursement_type/"

func CreateDisbursementType(disbursementType budget.DisbursementType) (budget.DisbursementType, error) {
	entity := budget.DisbursementType{}

	resp, _ := api.Rest().SetBody(disbursementType).Post(disbursementURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteDisbursementType(disbursementType budget.DisbursementType) (budget.DisbursementType, error) {
	entity := budget.DisbursementType{}

	resp, _ := api.Rest().SetBody(disbursementType).Post(disbursementURL + "delete")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateDisbursementType(disbursementType budget.DisbursementType) (budget.DisbursementType, error) {
	entity := budget.DisbursementType{}

	resp, _ := api.Rest().SetBody(disbursementType).Post(disbursementURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetDisbursementType(disbursementType string) (budget.DisbursementType, error) {
	entity := budget.DisbursementType{}
	resp, _ := api.Rest().Get(disbursementURL + "get/" + disbursementType)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetDisbursementTypes() ([]budget.DisbursementType, error) {
	entity := []budget.DisbursementType{}
	resp, _ := api.Rest().Get(disbursementURL + "all")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
