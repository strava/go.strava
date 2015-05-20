package strava

import (
	"encoding/json"
)

type CurrentAthleteService struct {
	client *Client
}

func NewCurrentAthleteService(client *Client) *CurrentAthleteService {
	return &CurrentAthleteService{client}
}

/*********************************************************/

type CurrentAthleteGetCall struct {
	service *CurrentAthleteService
}

func (s *CurrentAthleteService) Get() *CurrentAthleteGetCall {
	return &CurrentAthleteGetCall{
		service: s,
	}
}

func (c *CurrentAthleteGetCall) Do() (*AthleteDetailed, error) {
	data, err := c.service.client.run("GET", "/athlete", nil)
	if err != nil {
		return nil, err
	}

	var athlete AthleteDetailed
	err = json.Unmarshal(data, &athlete)
	if err != nil {
		return nil, err
	}

	return &athlete, nil
}

/*********************************************************/

type CurrentAthletePutCall struct {
	service *CurrentAthleteService
	ops     map[string]interface{}
}

func (s *CurrentAthleteService) Update() *CurrentAthletePutCall {
	return &CurrentAthletePutCall{
		service: s,
		ops:     make(map[string]interface{}),
	}
}

func (c *CurrentAthletePutCall) City(city string) *CurrentAthletePutCall {
	c.ops["city"] = city
	return c
}

func (c *CurrentAthletePutCall) State(state string) *CurrentAthletePutCall {
	c.ops["state"] = state
	return c
}

func (c *CurrentAthletePutCall) Country(country string) *CurrentAthletePutCall {
	c.ops["country"] = country
	return c
}

func (c *CurrentAthletePutCall) Gender(gender Gender) *CurrentAthletePutCall {
	c.ops["sex"] = gender
	return c
}

func (c *CurrentAthletePutCall) Weight(weight float64) *CurrentAthletePutCall {
	c.ops["weight"] = weight
	return c
}

func (c *CurrentAthletePutCall) Do() (*AthleteDetailed, error) {
	data, err := c.service.client.run("PUT", "/athlete", c.ops)
	if err != nil {
		return nil, err
	}

	var athlete AthleteDetailed
	err = json.Unmarshal(data, &athlete)
	if err != nil {
		return nil, err
	}

	return &athlete, nil
}

/*********************************************************/

type CurrentAthleteListActivitiesCall struct {
	service *CurrentAthleteService
	ops     map[string]interface{}
}

func (s *CurrentAthleteService) ListActivities() *CurrentAthleteListActivitiesCall {
	return &CurrentAthleteListActivitiesCall{
		service: s,
		ops:     make(map[string]interface{}),
	}
}

func (c *CurrentAthleteListActivitiesCall) Before(before int) *CurrentAthleteListActivitiesCall {
	c.ops["before"] = before
	return c
}

func (c *CurrentAthleteListActivitiesCall) After(after int) *CurrentAthleteListActivitiesCall {
	c.ops["after"] = after
	return c
}

func (c *CurrentAthleteListActivitiesCall) Page(page int) *CurrentAthleteListActivitiesCall {
	c.ops["page"] = page
	return c
}

func (c *CurrentAthleteListActivitiesCall) PerPage(perPage int) *CurrentAthleteListActivitiesCall {
	c.ops["per_page"] = perPage
	return c
}

func (c *CurrentAthleteListActivitiesCall) Do() ([]*ActivitySummary, error) {
	data, err := c.service.client.run("GET", "/athlete/activities", c.ops)
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

/*********************************************************/

type CurrentAthleteListFriendsActivitiesCall struct {
	service *CurrentAthleteService
	ops     map[string]interface{}
}

func (s *CurrentAthleteService) ListFriendsActivities() *CurrentAthleteListFriendsActivitiesCall {
	return &CurrentAthleteListFriendsActivitiesCall{
		service: s,
		ops:     make(map[string]interface{}),
	}
}

func (c *CurrentAthleteListFriendsActivitiesCall) Before(before int) *CurrentAthleteListFriendsActivitiesCall {
	c.ops["before"] = before
	return c
}

func (c *CurrentAthleteListFriendsActivitiesCall) Page(page int) *CurrentAthleteListFriendsActivitiesCall {
	c.ops["page"] = page
	return c
}

func (c *CurrentAthleteListFriendsActivitiesCall) PerPage(perPage int) *CurrentAthleteListFriendsActivitiesCall {
	c.ops["per_page"] = perPage
	return c
}

func (c *CurrentAthleteListFriendsActivitiesCall) Do() ([]*ActivitySummary, error) {
	data, err := c.service.client.run("GET", "/activities/following", c.ops)
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

/*********************************************************/

type CurrentAthleteListFriendsCall struct {
	service *CurrentAthleteService
	ops     map[string]interface{}
}

func (s *CurrentAthleteService) ListFriends() *CurrentAthleteListFriendsCall {
	return &CurrentAthleteListFriendsCall{
		service: s,
		ops:     make(map[string]interface{}),
	}
}

func (c *CurrentAthleteListFriendsCall) Page(page int) *CurrentAthleteListFriendsCall {
	c.ops["page"] = page
	return c
}

func (c *CurrentAthleteListFriendsCall) PerPage(perPage int) *CurrentAthleteListFriendsCall {
	c.ops["per_page"] = perPage
	return c
}

func (c *CurrentAthleteListFriendsCall) Do() ([]*AthleteSummary, error) {
	data, err := c.service.client.run("GET", "/athlete/friends", c.ops)
	if err != nil {
		return nil, err
	}

	friends := make([]*AthleteSummary, 0)
	err = json.Unmarshal(data, &friends)
	if err != nil {
		return nil, err
	}

	return friends, nil
}

/*********************************************************/

type CurrentAthleteListFollowersCall struct {
	service *CurrentAthleteService
	ops     map[string]interface{}
}

func (s *CurrentAthleteService) ListFollowers() *CurrentAthleteListFollowersCall {
	return &CurrentAthleteListFollowersCall{
		service: s,
		ops:     make(map[string]interface{}),
	}
}

func (c *CurrentAthleteListFollowersCall) Page(page int) *CurrentAthleteListFollowersCall {
	c.ops["page"] = page
	return c
}

func (c *CurrentAthleteListFollowersCall) PerPage(perPage int) *CurrentAthleteListFollowersCall {
	c.ops["per_page"] = perPage
	return c
}

func (c *CurrentAthleteListFollowersCall) Do() ([]*AthleteSummary, error) {
	data, err := c.service.client.run("GET", "/athlete/followers", c.ops)
	if err != nil {
		return nil, err
	}

	followers := make([]*AthleteSummary, 0)
	err = json.Unmarshal(data, &followers)
	if err != nil {
		return nil, err
	}

	return followers, nil
}

/*********************************************************/

type CurrentAthleteListClubsCall struct {
	service *CurrentAthleteService
}

func (s *CurrentAthleteService) ListClubs() *CurrentAthleteListClubsCall {
	return &CurrentAthleteListClubsCall{
		service: s,
	}
}

func (c *CurrentAthleteListClubsCall) Do() ([]*ClubSummary, error) {
	data, err := c.service.client.run("GET", "/athlete/clubs", nil)
	if err != nil {
		return nil, err
	}

	clubs := make([]*ClubSummary, 0)
	err = json.Unmarshal(data, &clubs)
	if err != nil {
		return nil, err
	}

	return clubs, nil
}

/*********************************************************/

type CurrentAthleteListStarredSegmentsCall struct {
	service *CurrentAthleteService
	ops     map[string]interface{}
}

func (s *CurrentAthleteService) ListStarredSegments() *CurrentAthleteListStarredSegmentsCall {
	return &CurrentAthleteListStarredSegmentsCall{
		service: s,
		ops:     make(map[string]interface{}),
	}
}

func (c *CurrentAthleteListStarredSegmentsCall) Page(page int) *CurrentAthleteListStarredSegmentsCall {
	c.ops["page"] = page
	return c
}

func (c *CurrentAthleteListStarredSegmentsCall) PerPage(perPage int) *CurrentAthleteListStarredSegmentsCall {
	c.ops["per_page"] = perPage
	return c
}

func (c *CurrentAthleteListStarredSegmentsCall) Do() ([]*PersonalSegmentSummary, error) {
	data, err := c.service.client.run("GET", "/segments/starred", c.ops)
	if err != nil {
		return nil, err
	}

	segments := make([]*PersonalSegmentSummary, 0)
	err = json.Unmarshal(data, &segments)
	if err != nil {
		return nil, err
	}

	return segments, nil
}
