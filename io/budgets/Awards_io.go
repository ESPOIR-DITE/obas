package budgets

import (
	"errors"
	"obas/api"
	"obas/domain/budget"
)

const awardsURL = api.BASE_URL + "/awards/"

func CreateAwards(awards budget.Awards) (budget.Awards, error) {
	entity := budget.Awards{}
	resp, _ := api.Rest().SetBody(awards).Post(awardsURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetAward(awardId string) (budget.Awards, error) {
	entity := budget.Awards{}
	resp, _ := api.Rest().Get(awardsURL + "read/" + awardId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetAwards() ([]budget.Awards, error) {
	entity := []budget.Awards{}
	resp, _ := api.Rest().Get(awardsURL + "all")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteAwards(awards budget.Awards) (budget.Awards, error) {
	entity := budget.Awards{}
	resp, _ := api.Rest().SetBody(awards).Post(awardsURL + "delete")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateAwards(awards budget.Awards) (budget.Awards, error) {
	entity := budget.Awards{}
	resp, _ := api.Rest().SetBody(awards).Post(awardsURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
