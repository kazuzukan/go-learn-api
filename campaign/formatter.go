package campaign

type CampaignFormatter struct {
	Id               int         `json:"id"`
	UserId           int         `json:"user_id"`
	Name             string      `json:"name"`
	ShortDescription string      `json:"short_description"`
	ImageUrl         interface{} `json:"image_url"`
	GoalAmount       int         `json:"goal_amount"`
	CurrentAmount    int         `json:"current_amount"`
	Slug             string      `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	var imageUrl interface{}
	if len(campaign.CampaignImages) > 0 {
		imageUrl = campaign.CampaignImages[0].FileName
	}

	formatter := CampaignFormatter{
		Id:               campaign.ID,
		UserId:           campaign.UserId,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		ImageUrl:         imageUrl,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
	}

	return formatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}
	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}
