package models

import "time"

type Likes struct {
	LikeId     int       `json:"likeId" gorm:"primaryKey"`
	UserIdFrom int       `json:"user_id_from"`
	UserIdTo   int       `json:"user_id_to"`
	CreatedAt  time.Time `json:"createdAt"`
}
