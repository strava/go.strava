package strava

import (
	"encoding/json"
	"fmt"
	"time"
)

type SegmentEffortDetailed struct {
	SegmentEffortSummary
}

type SegmentEffortSummary struct {
	Id       int64          `json:"id"`
	Name     string         `json:"name"`
	Segment  SegmentSummary `json:"segment"`
	Activity struct {
		Id int64 `json:"id"`
	} `json:"activity"`
	Athlete struct {
		Id int64 `json:"id"`
	} `json:"athlete"`

	KOMRank     int     `json:"kom_rank"` // 1-10 rank on segment at time of upload
	PRRank      int     `json:"pr_rank"`  // 1-3 personal record on segment at time of upload
	Distance    float64 `json:"distance"`
	MovingTime  int     `json:"moving_time"`
	ElapsedTime int     `json:"elapsed_time"`
	StartIndex  int     `json:"start_index"`
	EndIndex    int     `json:"end_index"`

	StartDate            time.Time `json:"-"`
	StartDateLocal       time.Time `json:"-"`
	StartDateString      string    `json:"start_date"`
	StartDateLocalString string    `json:"start_date_local"`
}

type SegmentEffortsService struct {
	client *Client
}

func NewSegmentEffortsService(client *Client) *SegmentEffortsService {
	return &SegmentEffortsService{client}
}

/*********************************************************/

type SegmentEffortsGetCall struct {
	service *SegmentEffortsService
	id      int64
}

func (s *SegmentEffortsService) Get(segmentEffortId int64) *SegmentEffortsGetCall {
	return &SegmentEffortsGetCall{
		service: s,
		id:      segmentEffortId,
	}
}

func (c *SegmentEffortsGetCall) Do() (*SegmentEffortDetailed, error) {
	data, err := c.service.client.run("GET", fmt.Sprintf("/segment_efforts/%d", c.id), nil)
	if err != nil {
		return nil, err
	}

	var effort SegmentEffortDetailed
	err = json.Unmarshal(data, &effort)
	if err != nil {
		return nil, err
	}

	effort.postProcessDetailed()

	return &effort, nil
}

/*********************************************************/

func (e *SegmentEffortDetailed) postProcessDetailed() {
	e.postProcessSummary()
}

func (e *SegmentEffortSummary) postProcessSummary() {
	e.Segment.postProcessSummary()

	e.StartDate, _ = time.Parse(timeFormat, e.StartDateString)
	e.StartDateLocal, _ = time.Parse(timeFormat, e.StartDateLocalString)
}
