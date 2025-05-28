package tests

import (
	"testing"

	"github.com/PyMarcus/FreelaIF/users-api/internal/adapters/utils"
	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T){
	password := "test123456&"
	hash, err := utils.HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	match := utils.CheckPasswordHash(password, hash)
	assert.True(t, match)
}