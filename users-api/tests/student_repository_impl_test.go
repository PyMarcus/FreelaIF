package tests

import (
	"testing"

	"github.com/PyMarcus/FreelaIF/users-api/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestStudentRepositoryCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIStudentRepository(ctrl)
	_, Student := createUsersToTest()

	mockRepo.EXPECT().Create(Student).Return(nil)

	err := mockRepo.Create(Student)
	assert.NoError(t, err)

}
