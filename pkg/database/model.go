package database

type Leaderboard struct {
	ID       int64   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserName string  `gorm:"column:user_name;not null" json:"user_name"`
	Country  string  `gorm:"column:country;not null" json:"country"`
	State    string  `gorm:"column:state;not null" json:"state"`
	Score    float64 `gorm:"column:score;not null" json:"score"`
}

// type Leaderboard struct {
// 	ID         uint   `gorm:"primaryKey"`
// 	UserName   string `gorm:"size:255;index"`              // Creates an index on UserName
// 	Score      int    `gorm:"index:idx_score_desc,sort:desc"` // Creates a descending index on Score
// 	Country    string `gorm:"index:idx_country_scope"`     // Part of a composite index
// 	State      string `gorm:"index:idx_country_scope"`     // Part of a composite index with Country
// 	CreatedAt  time.Time
// 	UpdatedAt  time.Time
// }
