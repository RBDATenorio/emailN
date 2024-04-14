package campaign

import (
	"emailN/internal/domain/campaign/contracts"
	internalerrors "emailN/internal/internal-errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

var (
	newCampaign = contracts.NewCampaign{
		Name:    "Test Y",
		Content: "Test content",
		Emails:  []string{"email@test.com", "test@email.com"},
	}

	newCampaignInvalid = contracts.NewCampaign{
		Name:    "",
		Content: "Test content",
		Emails:  []string{"email@test.com", "test@email.com"},
	}
)

func Test_Create_ShoulCallRepository_WhenNewCampaignIsValid(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	repository := new(repositoryMock)
	service := Service{Repository: repository}
	repository.On("Save", mock.Anything).Return(nil)

	// Act
	campaignId, err := service.Create(newCampaign)

	// Assert
	assert.Nil(err)
	assert.NotEmpty(campaignId)
	repository.AssertExpectations(t)
}

func Test_Create_ShouldNotCallRepository_WhenNewCampaignIsInvalid(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	repository := new(repositoryMock)
	service := Service{Repository: repository}

	// Act
	_, err := service.Create(newCampaignInvalid)

	// Assert
	assert.Equal("name is required", err[0].Error())
	repository.AssertNotCalled(t, "Save", mock.Anything)

}

func Test_Create_ShouldReturnErr_WhenRepositoryCouldNotSaveCampaign(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	repository := new(repositoryMock)
	repository.On("Save", mock.Anything).Return(internalerrors.ErrInternal)
	service := Service{Repository: repository}

	// Act
	campaingId, err := service.Create(newCampaign)

	// Assert
	assert.Equal(internalerrors.ErrInternal.Error(), err[0].Error())
	assert.Empty(campaingId)
	repository.AssertExpectations(t)
}
