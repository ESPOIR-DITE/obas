package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDocument(t *testing.T) {
	value, err := GetDocument("")
	assert.Nil(t, err)
	assert.Equal(t, value, "entity", "Return entity")
}

func TestGetDocuments(t *testing.T) {
	value, err := GetDocuments()
	assert.NotNil(t, err)
	assert.Equal(t, value, "Return entity")
}

func TestCreateDocument(t *testing.T) {
	value, err := CreateDocument("")
	assert.NotNil(t, err)
	assert.Equal(t, value, "Return entity")
}

func TestUpdateDocument(t *testing.T) {
	value, err := UpdateDocument("")
	assert.NotNil(t, err)
	assert.Equal(t, value, "Return entity")
}

func TestDeleteDocument(t *testing.T) {
	value, err := DeleteDocument("")
	assert.NotNil(t, err)
	assert.Equal(t, value, "Return entity")
}
