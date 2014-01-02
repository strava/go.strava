package strava

import (
	"time"
)

type CommentSummary struct {
	Id              int64          `json:"id"`
	ActivityId      int64          `json:"activity_id"`
	Text            string         `json:"text"`
	Athlete         AthleteSummary `json:"athlete"`
	CreatedAt       time.Time      `json:"-"`
	CreatedAtString string         `json:"created_at"` // the ISO 8601 encoding of when the comment was posted
}

func (c *CommentSummary) postProcessSummary() {
	c.Athlete.postProcessSummary()

	c.CreatedAt, _ = time.Parse(timeFormat, c.CreatedAtString)
}
