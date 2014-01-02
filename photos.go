package strava

import (
	"time"
)

type PhotoSummary struct {
	Id               int64     `json:"id"`
	ActivityId       int64     `json:"activity_id"`
	Reference        string    `json:"ref"`
	UID              string    `json:"uid"`
	Caption          string    `json:"caption"`
	Type             string    `json:"type"`
	UploadedAt       time.Time `json:"-"`
	CreatedAt        time.Time `json:"-"`
	UploadedAtString string    `json:"uploaded_at"` // the ISO 8601 encoding of when the photo was uploaded to the external service
	CreatedAtString  string    `json:"created_at"`  // the ISO 8601 encoding of when the photo was linked with strava
	Location         Location  `json:"location"`
}

func (p *PhotoSummary) postProcessSummary() {
	p.UploadedAt, _ = time.Parse(timeFormat, p.UploadedAtString)
	p.CreatedAt, _ = time.Parse(timeFormat, p.CreatedAtString)
}
