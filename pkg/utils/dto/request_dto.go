package dto

type RequestSubmitData struct {
	UserName string  `json:"user_name" validate:"required"`
	Country  string  `json:"country" validate:"required"`
	State    string  `json:"state" validate:"required"`
	Score    float64 `json:"score" validate:"required"`
}

type RequestGetRank struct {
	UserID int64  `form:"user_id" validate:"required"`
	Scope  string `form:"scope" validate:"required,oneof=state country global"`
}

type RequestGetTopNRank struct {
	N          int    `form:"n"`
	Scope      string `form:"scope" validate:"required,oneof=state country global"`
	ScopeValue string `form:"scope_value"`
}
