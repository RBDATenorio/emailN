package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign X"
	content := "Body"
	contacts := []string{"email1@example", "email2@example"}

	campaign := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.Id)
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
	assert.NotNil(campaign.CreatedOn)
}
