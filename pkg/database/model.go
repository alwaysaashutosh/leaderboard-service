package database

type Leaderboard struct {
	ID       int64   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserName string  `gorm:"column:user_name;not null" json:"user_name"`
	Country  string  `gorm:"column:country;not null;index:idx_leaderboard_country_score"`
	State    string  `gorm:"column:state;not null;index:idx_leaderboard_state_score"`
	Score    float64 `gorm:"column:score;not null;index:idx_leaderboard_country_score,sort:desc;index:idx_leaderboard_state_score,sort:desc;index:idx_leaderboard_score,sort:desc" json:"score"`
}
