package repository

import (
	"fmt"
	"strings"

	"github.com/alwaysaashutosh/leaderboard-service/pkg/database"
	"github.com/alwaysaashutosh/leaderboard-service/pkg/utils/constants"
	"github.com/alwaysaashutosh/leaderboard-service/pkg/utils/dto"
	"gorm.io/gorm"
)

type LeaderboardRepository interface {
	SubmitData(data dto.RequestSubmitData) (dto.ResponseSubmitData, error)
	GetRank(data dto.RequestGetRank) (dto.ResponseGetRank, error)
	GetTopRank(data dto.RequestGetTopNRank) (dto.ResponseGetTopNRank, error)
}

type leaderboardRepository struct {
	db *gorm.DB
}

func NewLeaderBoardRepository(db *gorm.DB) LeaderboardRepository {
	return &leaderboardRepository{db: db}
}

// List Report
func (repository *leaderboardRepository) SubmitData(data dto.RequestSubmitData) (dto.ResponseSubmitData, error) {

	entry := database.Leaderboard{
		UserName: data.UserName,
		Country:  strings.ToUpper(data.Country),
		State:    strings.ToUpper(data.State),
		Score:    data.Score,
	}
	if err := repository.db.Create(&entry).Error; err != nil {
		return dto.ResponseSubmitData{}, fmt.Errorf("error while saving the data to leaderboard, error: %v", err)
	}

	return dto.ResponseSubmitData{
		Status: "Data submission successful",
		ID:     entry.ID,
	}, nil
}

// GetRank - this is used to get the rank of a particular user in his state,country or globally
func (repository *leaderboardRepository) GetRank(data dto.RequestGetRank) (dto.ResponseGetRank, error) {
	var (
		user database.Leaderboard
		rank int64
	)

	// fetching user's details
	if err := repository.db.Where("id = ?", data.UserID).First(&user).Error; err != nil {
		return dto.ResponseGetRank{}, fmt.Errorf("error while fetching the existing user details, userID: %d | error: %v", data.UserID, err)
	}

	countquery := repository.db.Model(&database.Leaderboard{}).Where("score > ?", user.Score)

	switch data.Scope {
	case "country":
		countquery = countquery.Where("country = ?", strings.ToUpper(user.Country))
	case "state":
		countquery = countquery.Where("state = ?", strings.ToUpper(user.State))
	}

	if err := countquery.Count(&rank).Error; err != nil {
		return dto.ResponseGetRank{}, fmt.Errorf("error while fetching user's rank, userID: %d | error: %v", data.UserID, err)
	}

	return dto.ResponseGetRank{
		ResponseDTO: dto.ResponseDTO{
			Status:  constants.STATUSSUCCESS,
			Message: "Rank retrieved successfully",
		},
		Data: &dto.RankData{
			UserID: user.ID,
			Rank:   rank + 1,
			Score:  user.Score,
			Scope:  data.Scope,
		},
	}, nil
}

// List Top n ranks from a state, country, globally .
func (repository *leaderboardRepository) GetTopRank(data dto.RequestGetTopNRank) (dto.ResponseGetTopNRank, error) {
	var result []dto.GetTopNRank

	// Set default value for N if it's zero)
	if data.N == 0 {
		data.N = 50
	}

	query := repository.db.Model(&database.Leaderboard{}).Order("score desc").Limit(data.N)

	switch data.Scope {
	case "country":
		query = query.Where("country = ?", strings.ToUpper(data.ScopeValue))
	case "state":
		query = query.Where("state = ?", strings.ToUpper(data.ScopeValue))
	}

	if err := query.Select("user_name,score").Scan(&result).Error; err != nil {
		return dto.ResponseGetTopNRank{}, nil
	}

	for i := range result {
		result[i].Rank = i + 1
	}

	return dto.ResponseGetTopNRank{
		ResponseDTO: dto.ResponseDTO{
			Status:  constants.STATUSSUCCESS,
			Message: "Rank's retrieved successfully",
		},
		Data: &result,
	}, nil

}
