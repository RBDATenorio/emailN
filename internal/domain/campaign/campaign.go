package campaign

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

type Contacts struct {
	Email string
}

type Campaign struct {
	Id        string
	Name      string
	CreatedOn time.Time
	Content   string
	Contacts  []Contacts
}

func NewCampaign(name string, content string, emails []string) (*Campaign, []error) {
	var validationErrors []error

	if name == "" {
		validationErrors = append(validationErrors, errors.New("name is required"))
	}

	if content == "" {
		validationErrors = append(validationErrors, errors.New("content is required"))
	}

	if emails == nil {
		validationErrors = append(validationErrors, errors.New("emails are required"))
	}

	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	contacts := make([]Contacts, len(emails))
	for index, email := range emails {
		contacts[index].Email = email
	}

	return &Campaign{
		Id:        xid.New().String(),
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contacts:  contacts,
	}, nil
}
