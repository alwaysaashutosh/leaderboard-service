package response

import (
	"net/http"

	"github.com/alwaysaashutosh/leaderboard-service/pkg/utils/constants"
	"github.com/alwaysaashutosh/leaderboard-service/pkg/utils/dto"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// Return Error After service function called
func Send500(c *gin.Context, ErrMsg string, err error) {
	log.Error().Msgf(ErrMsg+": %s", err.Error())
	c.JSON(http.StatusBadGateway, dto.ErrorDTO{
		ErrorCode:    constants.Err500,
		ErrorMessage: ErrMsg,
	})
}

// Return Success After service function called
func Success(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, dto.SuccessDTO{
		Message: msg,
	})
}

// Return Success After service function called
func SuccessMsg(c *gin.Context, msg interface{}) {
	c.JSON(http.StatusOK, msg)

}

// / checkjsonbind and validation
func Send400(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, dto.ErrorDTO{
		ErrorCode:    constants.Err400,
		ErrorMessage: msg,
	})
}

// Return Success After service function called
func SuccessResp(c *gin.Context, resp interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"result": resp,
	})

}
