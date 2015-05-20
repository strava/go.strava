package strava

import (
	"time"
)

// EffortSummary is the base object for BestEfforts, SegmentEfforts and LapEfforts
type EffortSummary struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Activity struct {
		Id int64 `json:"id"`
	} `json:"activity"`
	Athlete struct {
		Id int64 `json:"id"`
	} `json:"athlete"`
	Distance       float64   `json:"distance"`
	MovingTime     int       `json:"moving_time"`
	ElapsedTime    int       `json:"elapsed_time"`
	StartIndex     int       `json:"start_index"`
	EndIndex       int       `json:"end_index"`
	StartDate      time.Time `json:"start_date"`
	StartDateLocal time.Time `json:"start_date_local"`
}
