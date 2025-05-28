package tests

import (
	"testing"

	"github.com/PyMarcus/FreelaIF/users-api/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCustomerRepositoryCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockICustomerRepository(ctrl)
	customer, _ := createUsersToTest()

	mockRepo.EXPECT().Create(customer).Return(nil)

	err := mockRepo.Create(customer)
	assert.NoError(t, err)

}
