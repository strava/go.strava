package strava

import (
	"encoding/json"
	"fmt"
)

type ActivityKudosService struct {
	client     *Client
	activityId int64
}

func NewActivityKudosService(client *Client, activityId int64) *ActivityKudosService {
	return &ActivityKudosService{client: client, activityId: activityId}
}

/*********************************************************/

type ActivityKudosListCall struct {
	service *ActivityKudosService
	ops     map[string]interface{}
}

func (s *ActivityKudosService) List() *ActivityKudosListCall {
	return &ActivityKudosListCall{
		service: s,
		ops:     make(map[string]interface{}),
	}
}

func (c *ActivityKudosListCall) Page(page int) *ActivityKudosListCall {
	c.ops["page"] = page
	return c
}

func (c *ActivityKudosListCall) PerPage(perPage int) *ActivityKudosListCall {
	c.ops["per_page"] = perPage
	return c
}

func (c *ActivityKudosListCall) Do() ([]*AthleteSummary, error) {
	data, err := c.service.client.run("GET", fmt.Sprintf("/activities/%d/kudos", c.service.activityId), c.ops)
	if err != nil {
		return nil, err
	}

	kudoers := make([]*AthleteSummary, 0)
	err = json.Unmarshal(data, &kudoers)
	if err != nil {
		return nil, err
	}

	return kudoers, nil
}

/*********************************************************/

type ActivityKudosPostCall struct {
	service *ActivityKudosService
}

func (s *ActivityKudosService) Create() *ActivityKudosPostCall {
	return &ActivityKudosPostCall{
		service: s,
	}
}

func (c *ActivityKudosPostCall) Do() error {
	_, err := c.service.client.run("POST", fmt.Sprintf("/activities/%d/kudos", c.service.activityId), nil)
	return err
}

/*********************************************************/

type ActivityKudosDeleteCall struct {
	service *ActivityKudosService
}

func (s *ActivityKudosService) Delete() *ActivityKudosDeleteCall {
	return &ActivityKudosDeleteCall{
		service: s,
	}
}

func (c *ActivityKudosDeleteCall) Do() error {
	_, err := c.service.client.run("DELETE", fmt.Sprintf("/activities/%d/kudos", c.service.activityId), nil)
	return err
}
