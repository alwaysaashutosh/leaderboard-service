package dto

// SuccessDTO - Success Message
type ResponseDTO struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ErrorDTO struct {
	ErrorCode string `json:"error_code"`
	ResponseDTO
}
type ResponseSubmitData struct {
	Status string `json:"status"`
	ID     int64  `json:"id"`
}
type ResponseGetRank struct {
	ResponseDTO
	Data *RankData `json:"data,omitempty"`
}
type RankData struct {
	UserID int64   `json:"user_id"`
	Rank   int64   `json:"rank"`
	Score  float64 `json:"score"`
	Scope  string  `json:"scope"` // "global", "country", or "state"
}

type ResponseGetTopNRank struct {
	ResponseDTO
	Data *[]GetTopNRank
}

type GetTopNRank struct {
	Rank     int     `json:"rank"`
	UserName string  `json:"user_name"`
	Score    float64 `json:"score"`
}
