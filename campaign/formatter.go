package campaign

import (
	"strings"
)

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

type CampaignDetailFormatter struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	Description      string `json:"description"`
	// mageUrl         interface{}               `json:"image_url"`
	GoalAmount    int                       `json:"goal_amount"`
	CurrentAmount int                       `json:"current_amount"`
	Slug          string                    `json:"slug"`
	UserId        int                       `json:"user_id"`
	Perks         []string                  `json:"perks"`
	User          CampaignUserFormatter     `json:"user"`
	Images        []CampaignImagesFormatter `json:"images"`
}

type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type CampaignImagesFormatter struct {
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatDetailCampaign(campaign Campaign) CampaignDetailFormatter {
	// var imageUrl interface{}
	// if len(campaign.CampaignImages) > 0 {
	// 	imageUrl = campaign.CampaignImages[0].FileName
	// }

	user := campaign.User
	campaignUserFormatter := CampaignUserFormatter{
		Name:     user.Name,
		ImageUrl: user.AvatarFilename,
	}

	images := campaign.CampaignImages
	campaignImages := []CampaignImagesFormatter{}
	for _, image := range images {
		campaignImageFormatter := CampaignImagesFormatter{}
		campaignImageFormatter.ImageUrl = image.FileName

		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
		}

		campaignImageFormatter.IsPrimary = isPrimary

		campaignImages = append(campaignImages, campaignImageFormatter)
	}

	campaignDetailFormatter := CampaignDetailFormatter{
		Id:               campaign.ID,
		UserId:           campaign.UserId,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		// ImageUrl:         imageUrl,
		GoalAmount:    campaign.GoalAmount,
		CurrentAmount: campaign.CurrentAmount,
		Slug:          campaign.Slug,
		// atau pake string,TrimSpace
		Perks:  strings.Split(campaign.Perks, ", "),
		User:   campaignUserFormatter,
		Images: campaignImages,
	}

	return campaignDetailFormatter
}
