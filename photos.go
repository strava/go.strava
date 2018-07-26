package strava

import "time"

type StravaPhotoSummary struct {
	UniqueId string `json:"unique_id"`
}

type InstagramPhotoSummary struct {
	Reference string `json:"ref"`
	UID       string `json:"uid"`
}

type PhotoSummary struct {
	Id            int64             `json:"id"`
	ActivityId    int64             `json:"activity_id"`
	ActivityName  string            `json:"activity_name"`
	ResourceState int               `json:"resource_state"`
	Caption       string            `json:"caption"`
	Type          string            `json:"type"`
	Source        int               `json:"source"`
	Urls          map[string]string `json:"urls"`
	Sizes         map[string][2]int `json:"sizes"`
	UploadedAt    time.Time         `json:"uploaded_at"`
	CreatedAt     time.Time         `json:"created_at"`
	Location      Location          `json:"location"`

	StravaPhotoSummary
	InstagramPhotoSummary
}
