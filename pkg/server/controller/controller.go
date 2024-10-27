package controller

import (
	"github.com/alwaysaashutosh/leaderboard-service/pkg/server/service"

	"github.com/alwaysaashutosh/leaderboard-service/pkg/utils/dto"
	"github.com/alwaysaashutosh/leaderboard-service/pkg/utils/request"
	"github.com/alwaysaashutosh/leaderboard-service/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type LeaderboardController struct {
	service service.LeaderboardService
}

func NewLeaderBoardController(service service.LeaderboardService) *LeaderboardController {
	return &LeaderboardController{service: service}
}

// SubmitData - Submits a user’s data to the leaderboard service.
//
//	@Summary		Submit User Data
//	@Description	Submit specific details related to a user.
//	@BasePath		/leaderboard-service
//	@Accept			json
//	@Produce		json
//
//	@Param			        {object}	body	dto.RequestSubmitData	false	"Request Body"
//	@Success		200		{object}	dto.ResponseSubmitData
//	@Failure		400		{object}	dto.ErrorDTO
//	@Failure		500		{object}	dto.ErrorDTO
//	@Router			/api/v1/submit [post]
func (controller *LeaderboardController) SubmitData(c *gin.Context) {
	var data dto.RequestSubmitData

	// function to bind the request body.
	if request.CheckJSON(c, &data) {
		return // exit
	}

	// function Validator to validate the request body on validation tags .
	if request.CheckValidator(c, data) {
		return // exit
	}

	resp, err := controller.service.SubmitData(data)
	if err != nil {
		response.Send500(c, err)
		return
	}

	response.SuccessMsg(c, resp)
}

// GetRank - Retrieves a specific user's rank across various scopes.
//
//	@Summary		Get Rank for a Specific User
//	@Description	Fetch the rank of a user within multiple defined scopes.
//	@BasePath		/leaderboard-service
//	@Accept			json
//	@Produce		json
//
//	@Param			        {object}    query	dto.RequestGetRank	false  "Request Body"
//	@Success		200		{object}	dto.ResponseGetRank
//	@Failure		400		{object}	dto.ErrorDTO
//	@Failure		500		{object}	dto.ErrorDTO
//	@Router			/api/v1/get_rank [get]
func (controller *LeaderboardController) GetRank(c *gin.Context) {
	var data dto.RequestGetRank

	if request.CheckQueryParams(c, &data) {
		return // exit
	}

	if request.CheckValidator(c, data) {
		return // exit
	}

	resp, err := controller.service.GetRank(data)
	if err != nil {
		response.Send500(c, err)
		return
	}

	response.SuccessMsg(c, resp)

}

// GetTopRank - Retrieves the top-ranking users across various scopes.
//
//	@Summary		Get Top Rankings Dashboard
//	@Description	Fetch the top ranks of users across multiple defined scopes.
//	@BasePath		/leaderboard-service
//	@Accept			json
//	@Produce		json
//
//	@Param			            {object}    query  dto.RequestGetTopNRank	false   "Request Body"
//	@Success		200			{object}	dto.ResponseGetTopNRank
//	@Failure		400			{object}	dto.ErrorDTO
//	@Failure		500			{object}	dto.ErrorDTO
//	@Router			/api/v1/list_top_n [get]
func (controller *LeaderboardController) GetTopRank(c *gin.Context) {

	var data dto.RequestGetTopNRank

	if request.CheckQueryParams(c, &data) {
		return // exit
	}

	if request.CheckValidator(c, data) {
		return // exit
	}

	if data.Scope != "global" && data.ScopeValue == "" {
		response.Send400(c, "scope value is needed if the scope isn't global", "")
		return
	}

	resp, err := controller.service.GetTopNRank(data)
	if err != nil {
		response.Send500(c, err)
		return
	}

	response.SuccessMsg(c, resp)

}
