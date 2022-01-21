package handler

import (
	"bwa-project/campaign"
	"bwa-project/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaign.Service
}

func NewCampaignHanlder(campaignService campaign.Service) *campaignHandler {
	return &campaignHandler{campaignService}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.campaignService.GetCampaigns(userId)
	if err != nil {
		response := helper.APIResponse("Failed to get campaigns", http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := campaign.FormatCampaigns(campaigns)
	response := helper.APIResponse("List of campaings", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of campaign", http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusOK, response)
		return
	}

	campaignDetail, err := h.campaignService.GetCampaign(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of campaign", http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusOK, response)
		return
	}

	formatter := campaign.FormatDetailCampaign(campaignDetail)
	response := helper.APIResponse("Campaign detail", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
