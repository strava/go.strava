package strava

import (
	"encoding/json"
	"fmt"
)

type ClubDetailed struct {
	ClubSummary
	Description string    `json:"description"`
	Type        ClubType  `json:"club_type"`
	SportType   SportType `json:"sport_type"`
	City        string    `json:"city"`
	State       string    `json:"state"`
	Country     string    `json:"country"`
	Private     bool      `json:"private"`
	MemberCount int       `json:"member_count"`
}

type ClubSummary struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	ProfileMedium string `json:"profile_medium"` // URL to a 62x62 pixel profile picture
	Profile       string `json:"profile"`        // URL to a 124x124 pixel profile picture
}

type ClubType string

var ClubTypes = struct {
	CasualClub ClubType
	RacingTeam ClubType
	Shop       ClubType
	Company    ClubType
	Other      ClubType
}{"casual_club", "racing_team", "shop", "company", "other"}

type SportType string

var SportTypes = struct {
	Cycling   SportType
	Running   SportType
	Triathlon SportType
	Other     SportType
}{"cycling", "running", "triathlon", "other"}

type ClubsService struct {
	client *Client
}

func NewClubsService(client *Client) *ClubsService {
	return &ClubsService{client}
}

/*********************************************************/

type ClubsGetCall struct {
	service *ClubsService
	id      int64
}

func (s *ClubsService) Get(clubId int64) *ClubsGetCall {
	return &ClubsGetCall{
		service: s,
		id:      clubId,
	}
}

func (c *ClubsGetCall) Do() (*ClubDetailed, error) {
	data, err := c.service.client.run("GET", fmt.Sprintf("/clubs/%d", c.id), nil)
	if err != nil {
		return nil, err
	}

	var club ClubDetailed
	err = json.Unmarshal(data, &club)
	if err != nil {
		return nil, err
	}

	return &club, nil
}

/*********************************************************/

type ClubListMembersCall struct {
	service *ClubsService
	id      int64
	ops     map[string]interface{}
}

func (s *ClubsService) ListMembers(clubId int64) *ClubListMembersCall {
	return &ClubListMembersCall{
		service: s,
		id:      clubId,
		ops:     make(map[string]interface{}),
	}
}

func (c *ClubListMembersCall) Page(page int) *ClubListMembersCall {
	c.ops["page"] = page
	return c
}

func (c *ClubListMembersCall) PerPage(perPage int) *ClubListMembersCall {
	c.ops["per_page"] = perPage
	return c
}

func (c *ClubListMembersCall) Do() ([]*AthleteSummary, error) {
	data, err := c.service.client.run("GET", fmt.Sprintf("/clubs/%d/members", c.id), c.ops)
	if err != nil {
		return nil, err
	}

	members := make([]*AthleteSummary, 0)
	err = json.Unmarshal(data, &members)
	if err != nil {
		return nil, err
	}

	return members, nil
}

/*********************************************************/

type ClubListActivitiesCall struct {
	service *ClubsService
	id      int64
	ops     map[string]interface{}
}

func (s *ClubsService) ListActivities(clubId int64) *ClubListActivitiesCall {
	return &ClubListActivitiesCall{
		service: s,
		id:      clubId,
		ops:     make(map[string]interface{}),
	}
}

func (c *ClubListActivitiesCall) Page(page int) *ClubListActivitiesCall {
	c.ops["page"] = page
	return c
}

func (c *ClubListActivitiesCall) PerPage(perPage int) *ClubListActivitiesCall {
	c.ops["per_page"] = perPage
	return c
}

func (c *ClubListActivitiesCall) Do() ([]*ActivitySummary, error) {
	data, err := c.service.client.run("GET", fmt.Sprintf("/clubs/%d/activities", c.id), c.ops)
	if err != nil {
		return nil, err
	}

	activities := make([]*ActivitySummary, 0)
	err = json.Unmarshal(data, &activities)
	if err != nil {
		return nil, err
	}

	return activities, nil
}
