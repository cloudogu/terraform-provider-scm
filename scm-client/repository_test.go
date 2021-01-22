package scm_client

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

var testConfig = Config{
	URL:      "http://localhost:8080/scm",
	Username: "scmadmin",
	Password: "scmadmin",
}

var testRepo = Repository{
	NameSpace:   "testspace",
	Name:        "testrepo",
	Type:        "git",
	Description: "desc",
}

func TestClient_CreateRepository(t *testing.T) {
	c := NewClient(testConfig)

	err := c.CreateRepository(testRepo)

	require.NoError(t, err)
}

func TestClient_GetRepository(t *testing.T) {
	c := NewClient(testConfig)

	r, err := c.GetRepository(testRepo.GetID())
	require.NoError(t, err)

	assert.Equal(t, testRepo.NameSpace, r.NameSpace)
	assert.Equal(t, testRepo.Name, r.Name)
	assert.Equal(t, testRepo.Type, r.Type)
	assert.Equal(t, testRepo.Description, r.Description)
}

func TestClient_UpdateRepository(t *testing.T) {
	c := NewClient(testConfig)
	oldRepo, err := c.GetRepository(testRepo.GetID())
	require.NoError(t, err)
	updatedRepo := oldRepo
	updatedRepo.Description = "updated desc"

	err = c.UpdateRepository(testRepo.GetID(), updatedRepo)
	require.NoError(t, err)

	newRepo, err := c.GetRepository(testRepo.GetID())
	require.NoError(t, err)
	require.Equal(t, updatedRepo, newRepo)
}

func TestClient_DeleteRepository(t *testing.T) {
	c := NewClient(testConfig)

	err := c.DeleteRepository(testRepo.GetID())
	require.NoError(t, err)

	_, err = c.GetRepository(testRepo.GetID())
	require.Error(t, err)
}