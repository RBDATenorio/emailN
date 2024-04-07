package campaign

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	validName     = "Campaign X"
	validContent  = "Body"
	validContacts = []string{"email1@example", "email2@example"}
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(validName, validContent, validContacts)

	assert.NotNil(campaign.Id)
	assert.Equal(campaign.Name, validName)
	assert.Equal(campaign.Content, validContent)
	assert.Equal(len(campaign.Contacts), len(validContacts))
	assert.NotNil(campaign.CreatedOn)
}

func Test_NewCampaign_ShouldNotCreateCampaign_WhenNameIsEmpty(t *testing.T) {
	assert := assert.New(t)
	expectedErrorMessage := "name is required"
	emptyName := ""

	campaign, err := NewCampaign(emptyName, validContent, validContacts)

	assert.Nil(campaign)
	assert.Len(err, 1)
	assert.EqualError(err[0], expectedErrorMessage)
}

func Test_NewCampaign_ShouldNotCreateCampaign_WhenContentIsEmpty(t *testing.T) {
	assert := assert.New(t)
	expectedErrorMessage := "content is required"
	emptyContent := ""

	campaign, err := NewCampaign(validName, emptyContent, validContacts)

	assert.Nil(campaign)
	assert.Len(err, 1)
	assert.EqualError(err[0], expectedErrorMessage)
}

func Test_NewCampaign_ShouldNotCreateCampaign_WhenEmailsIsNil(t *testing.T) {
	assert := assert.New(t)
	expectedErrorMessage := "emails are required"

	campaign, err := NewCampaign(validName, validContent, nil)

	assert.Nil(campaign)
	assert.Len(err, 1)
	assert.EqualError(err[0], expectedErrorMessage)
}

func Test_NewCampaign_ShouldNotCreateCampaign_WhenNameContentAndEmailsAreInvalid(t *testing.T) {
	assert := assert.New(t)
	expectedErrorList := []error{errors.New("name is required"),
		errors.New("content is required"),
		errors.New("emails are required")}

	campaign, err := NewCampaign("", "", nil)

	assert.Nil(campaign)
	assert.Len(err, 3)
	assert.ElementsMatch(err, expectedErrorList)
}
