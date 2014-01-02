package strava

import (
	"time"
)

type BestEffort struct {
	Id       int64          `json:"id"`
	Name     string         `json:"name"`
	Segment  SegmentSummary `json:"segment"`
	Activity struct {
		Id int `json:"id"`
	} `json:"activity"`
	Athlete struct {
		Id int `json:"id"`
	} `json:"athlete"`
	KOMRank     int     `json:"kom_rank"` // 1-10 rank on segment at time of upload
	PRRank      int     `json:"pr_rank"`  // 1-3 personal record on segment at time of upload
	Distance    float64 `json:"distance"`
	MovingTime  int     `json:"moving_time"`
	ElapsedTime int     `json:"elapsed_time"`

	StartDate            time.Time `json:"-"`
	StartDateLocal       time.Time `json:"-"`
	StartDateString      string    `json:"start_date"`       // the ISO 8601 encoding of when the effort started
	StartDateLocalString string    `json:"start_date_local"` // the ISO 8601 encoding of the UTC version of the local time when the effort started, see: http://strava.github.io/api/#dates

	StartIndex int `json:"start_index"`
	EndIndex   int `json:"end_index"`
}

func (be *BestEffort) postProcess() {
	be.Segment.postProcessSummary()

	be.StartDate, _ = time.Parse(timeFormat, be.StartDateString)
	be.StartDateLocal, _ = time.Parse(timeFormat, be.StartDateLocalString)
}
