package demographics

import (
	"errors"
	"obas/api"
	domain "obas/domain/demographics"
)

const genderUrl = api.BASE_URL + "/demographics"

type Genders domain.Gender

func GetGenders() ([]Genders, error) {
	entites := []Genders{}
	resp, _ := api.Rest().Get(genderUrl + "/gender/all")

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetGender(id string) (Genders, error) {
	entity := Genders{}
	resp, _ := api.Rest().Get(genderUrl + "/gender/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateGender(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(genderUrl + "/gender/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateGender(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(genderUrl + "/gender/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteGender(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(genderUrl + "/gender/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
