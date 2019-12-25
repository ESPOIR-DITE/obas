package users

import (
	"errors"
	"obas/api"
	domain "obas/domain/users"
)

const userRoleUrl = api.BASE_URL + "/users"

type UserRole domain.UserRole

func GetUserRoles() ([]UserRole, error) {
	entites := []UserRole{}
	resp, _ := api.Rest().Get(userRoleUrl + "/role/all")

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetUserRole(id string) (UserRole, error) {
	entity := UserRole{}
	resp, _ := api.Rest().Get(userRoleUrl + "/role/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateUserRole(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userRoleUrl + "/role/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateUserRole(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userRoleUrl + "/role/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteUserRole(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userRoleUrl + "/role/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
func GetUserRoleWithUserId(userId string) (domain.UserRole, error) {
	entity := domain.UserRole{}
	resp, _ := api.Rest().Get(userRoleUrl + "/role/getforuser/" + userId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
