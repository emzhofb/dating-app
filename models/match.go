package models

import "time"

type Matches struct {
	MatchId   int       `json:"matchId" gorm:"primaryKey"`
	UserIdA   int       `json:"user_id_a"`
	UserIdB   int       `json:"user_id_b"`
	CreatedAt time.Time `json:"createdAt"`
}
