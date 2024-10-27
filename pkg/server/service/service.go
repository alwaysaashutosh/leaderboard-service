package service

import (
	"github.com/alwaysaashutosh/leaderboard-service/pkg/server/repository"
	"github.com/alwaysaashutosh/leaderboard-service/pkg/utils/dto"
)

type LeaderboardService interface {
	SubmitData(data dto.RequestSubmitData) (dto.ResponseSubmitData, error)
	GetRank(data dto.RequestGetRank) (dto.ResponseGetRank, error)
	GetTopNRank(data dto.RequestGetTopNRank) (dto.ResponseGetTopNRank, error)
}

type leaderboardService struct {
	repo repository.LeaderboardRepository
}

func NewLeaderboardService(repo repository.LeaderboardRepository) LeaderboardService {
	return &leaderboardService{repo: repo}
}

// GetReports - Get reports
func (service *leaderboardService) SubmitData(data dto.RequestSubmitData) (dto.ResponseSubmitData, error) {
	ack, err := service.repo.SubmitData(data)
	if err != nil {
		return dto.ResponseSubmitData{}, err
	}

	return ack, nil
}

// GetRank - Get reports
func (service *leaderboardService) GetRank(data dto.RequestGetRank) (dto.ResponseGetRank, error) {

	resp, err := service.repo.GetRank(data)
	if err != nil {
		return dto.ResponseGetRank{}, err
	}

	return resp, nil
}

// GetTopNRank - Get reports
func (service *leaderboardService) GetTopNRank(data dto.RequestGetTopNRank) (dto.ResponseGetTopNRank, error) {

	resp, err := service.repo.GetTopRank(data)
	if err != nil {
		return dto.ResponseGetTopNRank{}, err
	}

	return resp, nil
}
