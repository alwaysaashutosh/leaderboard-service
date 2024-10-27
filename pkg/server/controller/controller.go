package controller

import (
	"github.com/alwaysaashutosh/leaderboard-service/pkg/server/service"

	"github.com/alwaysaashutosh/leaderboard-service/pkg/utils/dto"
	"github.com/alwaysaashutosh/leaderboard-service/pkg/utils/request"
	"github.com/alwaysaashutosh/leaderboard-service/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

// type LeaderboardController struct {
// 	service *service.LeaderboardService
// }

//	func NewLeaderBoardController(service *service.LeaderboardService) *LeaderboardController {
//		return &LeaderboardController{service: service}
//	}
type LeaderboardController struct {
	service service.LeaderboardService
}

func NewLeaderBoardController(service service.LeaderboardService) *LeaderboardController {
	return &LeaderboardController{service: service}
}

func (controller *LeaderboardController) SubmitData(c *gin.Context) {
	var data dto.RequestSubmitData

	if request.CheckJSON(c, &data) {
		return // exit
	}

	// function Validator to Check Request Body is correct or not
	if request.CheckValidator(c, data) {
		return // exit
	}

	resp, err := controller.service.SubmitData(data)
	if err != nil {
		response.Send500(c, err.Error(), err)
		return
	}

	response.SuccessMsg(c, resp)
}

func (controller *LeaderboardController) GetData(c *gin.Context) {
	var data dto.RequestGetRank

	if request.CheckQueryParams(c, &data) {
		return // exit
	}

	if request.CheckValidator(c, data) {
		return // exit
	}

	resp, err := controller.service.GetRank(data)
	if err != nil {
		response.Send500(c, err.Error(), err)
		return
	}

	response.SuccessMsg(c, resp)

}

func (controller *LeaderboardController) GetTopRank(c *gin.Context) {

	var data dto.RequestGetTopNRank

	if request.CheckQueryParams(c, &data) {
		return // exit
	}

	if request.CheckValidator(c, data) {
		return // exit
	}

	resp, err := controller.service.GetTopNRank(data)
	if err != nil {
		response.Send500(c, err.Error(), err)
		return
	}

	response.SuccessMsg(c, resp)

}
