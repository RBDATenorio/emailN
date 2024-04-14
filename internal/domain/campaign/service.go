package campaign

import (
	"emailN/internal/domain/campaign/contracts"
	internalerrors "emailN/internal/internal-errors"
)

type Service struct {
	Repository
}

func (s *Service) Create(newCampaign contracts.NewCampaign) (string, []error) {
	campaign, domainErrs := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	if domainErrs != nil {
		return "", domainErrs
	}

	err := s.Repository.Save(campaign)

	if err != nil {
		return "", []error{internalerrors.ErrInternal}
	}

	return campaign.Id, nil
}
