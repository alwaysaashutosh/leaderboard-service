package response

import (
	"net/http"

	"github.com/alwaysaashutosh/leaderboard-service/pkg/utils/constants"
	"github.com/alwaysaashutosh/leaderboard-service/pkg/utils/dto"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// Return Error After service function called
func Send500(c *gin.Context, err error) {
	log.Error().Msgf("server error: %v\n", err.Error())
	c.JSON(http.StatusBadGateway, dto.ErrorDTO{
		ErrorCode: constants.Err500,
		ResponseDTO: dto.ResponseDTO{
			Status:  constants.STATUSFAILED,
			Message: err.Error(),
		},
	})
}

// Return Success After service function called
func SuccessMsg(c *gin.Context, msg interface{}) {
	c.JSON(http.StatusOK, msg)
}

// / checkjsonbind and validation
func Send400(c *gin.Context, msg string, err string) {
	log.Error().Msgf("%v : %v\n", msg, err)
	c.JSON(http.StatusBadRequest, dto.ErrorDTO{
		ErrorCode: constants.Err400,
		ResponseDTO: dto.ResponseDTO{
			Status:  constants.STATUSFAILED,
			Message: msg,
		},
	})
}
