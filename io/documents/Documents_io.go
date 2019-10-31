package documents

import (
	"errors"
	"obas/api"
	domain "obas/domain/documents"
)

const documentUrl = api.BASE_URL + "/documents"

type Document domain.Document

func GetDocuments() ([]Document, error) {
	entites := []Document{}
	resp, _ := api.Rest().Get(documentUrl + "/all")
	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetDocument(id string) (Document, error) {
	entity := Document{}
	resp, _ := api.Rest().Get(documentUrl + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateDocument(entity Document, token string) (bool, error) {
	resp, _ := api.Rest().
		SetAuthToken(token).
		SetBody(entity).
		Post(documentUrl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateDocument(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(documentUrl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteDocument(entity interface{}) (bool, error) {

	resp, _ := api.Rest().
		SetBody(entity).
		Post(documentUrl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
