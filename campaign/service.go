package campaign

type Service interface {
	GetCampaings(userid int) ([]Campaign, error)
}

type service struct {
	repository Repository
}

func NewServices(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaings(userId int) ([]Campaign, error) {
	if userId != 0 {
		campaigns, err := s.repository.FindByUserId(userId)
		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}
