package routes

import (
	"github.com/alwaysaashutosh/leaderboard-service/pkg/database"
	"github.com/alwaysaashutosh/leaderboard-service/pkg/server/controller"
	"github.com/alwaysaashutosh/leaderboard-service/pkg/server/repository"
	"github.com/alwaysaashutosh/leaderboard-service/pkg/server/service"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine, basePath string) {
	app := r.Group(basePath)
	app.Use(logger.SetLogger())

	// configuration for Swagger docs
	// docs.SwaggerInfo.BasePath = basePath
	// app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	db := database.Handler()
	repo := repository.NewLeaderBoardRepository(db)
	service := service.NewLeaderboardService(repo)
	controller := controller.NewLeaderBoardController(service)

	v1 := app.Group("/v1")
	v1.POST("/submit", controller.SubmitData)
	v1.GET("/get_rank", controller.GetData)
	v1.GET("/list_top_n", controller.GetTopRank)
}
