package request

import (
	"github.com/alwaysaashutosh/leaderboard-service/pkg/utils/constants"
	"github.com/alwaysaashutosh/leaderboard-service/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/rs/zerolog/log"
)

// Check JSON according to given structure
func CheckJSON(c *gin.Context, data interface{}) bool {
	if err := c.BindJSON(&data); err != nil {
		log.Error().Msgf(constants.WrongReqBody + err.Error())
		// jsonMsgError_400 Function Call
		response.Send400(c, constants.InvalidReqBody)
		return true
	}
	return false
}

// Validate Data According to Request Body
func CheckValidator(c *gin.Context, data interface{}) bool {
	validation := validator.New()
	if err := validation.Struct(data); err != nil {
		log.Error().Msgf(constants.WrongReqBody + err.Error())
		// jsonMsgError_400 Function Call
		response.Send400(c, constants.RequiredReqBody)
		return true
	}
	return false
}

// Check QueryParams according to given structure
func CheckQueryParams(c *gin.Context, data interface{}) bool {
	if err := c.BindQuery(data); err != nil {
		// jsonMsgError_400 Function Call
		response.Send400(c, constants.InvalidQUeryParams)
		return true
	}
	return false
}
