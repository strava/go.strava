package strava

import (
	"time"
)

type PhotoSummary struct {
	Id         int64     `json:"id"`
	ActivityId int64     `json:"activity_id"`
	Reference  string    `json:"ref"`
	UID        string    `json:"uid"`
	Caption    string    `json:"caption"`
	Type       string    `json:"type"`
	UploadedAt time.Time `json:"uploaded_at"`
	CreatedAt  time.Time `json:"created_at"`
	Location   Location  `json:"location"`
}
