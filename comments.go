package strava

import (
	"encoding/json"
	"fmt"
	"time"
)

type CommentDetailed struct {
	CommentSummary
}

type CommentSummary struct {
	Id         int64          `json:"id"`
	ActivityId int64          `json:"activity_id"`
	Text       string         `json:"text"`
	Athlete    AthleteSummary `json:"athlete"`
	CreatedAt  time.Time      `json:"created_at"`
}

type ActivityCommentsService struct {
	client     *Client
	activityId int64
}

func NewActivityCommentsService(client *Client, activityId int64) *ActivityCommentsService {
	return &ActivityCommentsService{client: client, activityId: activityId}
}

/*********************************************************/

type ActivitiesCommentsListCall struct {
	service *ActivityCommentsService
	ops     map[string]interface{}
}

func (s *ActivityCommentsService) List() *ActivitiesCommentsListCall {
	return &ActivitiesCommentsListCall{
		service: s,
		ops:     make(map[string]interface{}),
	}
}

func (c *ActivitiesCommentsListCall) IncludeMarkdown() *ActivitiesCommentsListCall {
	c.ops["markdown"] = true
	return c
}

func (c *ActivitiesCommentsListCall) Page(page int) *ActivitiesCommentsListCall {
	c.ops["page"] = page
	return c
}

func (c *ActivitiesCommentsListCall) PerPage(perPage int) *ActivitiesCommentsListCall {
	c.ops["per_page"] = perPage
	return c
}

func (c *ActivitiesCommentsListCall) Do() ([]*CommentSummary, error) {
	data, err := c.service.client.run("GET", fmt.Sprintf("/activities/%d/comments", c.service.activityId), c.ops)
	if err != nil {
		return nil, err
	}

	comments := make([]*CommentSummary, 0)
	err = json.Unmarshal(data, &comments)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

/*********************************************************/

type ActivityCommentsPostCall struct {
	service *ActivityCommentsService
	text    string
}

func (s *ActivityCommentsService) Create(text string) *ActivityCommentsPostCall {
	return &ActivityCommentsPostCall{
		service: s,
		text:    text,
	}
}

func (c *ActivityCommentsPostCall) Do() (*CommentDetailed, error) {
	data, err := c.service.client.run(
		"POST",
		fmt.Sprintf("/activities/%d/comments", c.service.activityId),
		map[string]interface{}{"text": c.text},
	)
	if err != nil {
		return nil, err
	}

	var comment CommentDetailed
	err = json.Unmarshal(data, &comment)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

/*********************************************************/

type ActivityCommentsDeleteCall struct {
	service    *ActivityCommentsService
	activityId int64
	commentId  int64
}

func (s *ActivityCommentsService) Delete(commentId int64) *ActivityCommentsDeleteCall {
	return &ActivityCommentsDeleteCall{
		service:   s,
		commentId: commentId,
	}
}

func (c *ActivityCommentsDeleteCall) Do() error {
	_, err := c.service.client.run(
		"DELETE",
		fmt.Sprintf("/activities/%d/comments/%d", c.service.activityId, c.commentId),
		nil,
	)
	return err
}
