package strava

import (
	"encoding/json"
	"fmt"
	"time"
)

type AthleteDetailed struct {
	AthleteSummary
	Email                 string         `json:"email"`
	FollowerCount         int            `json:"follower_count"`
	FriendCount           int            `json:"friend_count"`
	MutualFriendCount     int            `json:"mutual_friend_count"`
	DatePreference        string         `json:"date_preference"`
	MeasurementPreference string         `json:"measurement_preference"`
	FTP                   int            `json:"ftp"`
	Clubs                 []*ClubSummary `json:"clubs"`
	Bikes                 []*GearSummary `json:"bikes"`
	Shoes                 []*GearSummary `json:"shoes"`
}

type AthleteSummary struct {
	AthleteMeta
	FirstName        string    `json:"firstname"`
	LastName         string    `json:"lastname"`
	ProfileMedium    string    `json:"profile_medium"` // URL to a 62x62 pixel profile picture
	Profile          string    `json:"profile"`        // URL to a 124x124 pixel profile picture
	City             string    `json:"city"`
	State            string    `json:"state"`
	Country          string    `json:"country"`
	Gender           Gender    `json:"sex"`
	Friend           string    `json:"friend"`   // ‘pending’, ‘accepted’, ‘blocked’ or ‘null’, the authenticated athlete’s following status of this athlete
	Follower         string    `json:"follower"` // this athlete’s following status of the authenticated athlete
	Premium          bool      `json:"premium"`
	CreatedAt        time.Time `json:"-"`
	UpdatedAt        time.Time `json:"-"`
	CreatedAtString  string    `json:"created_at"`        // the ISO 8601 encoding of when the athlete joined
	UpdatedAtString  string    `json:"updated_at"`        // the ISO 8601 encoding of when the athlete updated something
	ApproveFollowers bool      `json:"approve_followers"` // if has enhanced privacy enabled
}

type AthleteMeta struct {
	Id int64 `json:"id"`
}

type Gender string

var Genders = struct {
	Unspecified Gender
	Male        Gender
	Female      Gender
}{"", "M", "F"}

type AthletesService struct {
	client *Client
}

func NewAthletesService(client *Client) *AthletesService {
	return &AthletesService{client}
}

/*********************************************************/

type AthletesGetCall struct {
	service *AthletesService
	id      int64
}

func (s *AthletesService) Get(athleteId int64) *AthletesGetCall {
	return &AthletesGetCall{
		service: s,
		id:      athleteId,
	}
}

func (c *AthletesGetCall) Do() (*AthleteSummary, error) {
	data, err := c.service.client.run("GET", fmt.Sprintf("/athletes/%d", c.id), nil)
	if err != nil {
		return nil, err
	}

	var athlete AthleteSummary
	err = json.Unmarshal(data, &athlete)
	if err != nil {
		return nil, err
	}

	athlete.postProcessSummary()

	return &athlete, nil
}

/*********************************************************/

type AthletesListStarredSegmentsCall struct {
	service *AthletesService
	id      int64
	ops     map[string]interface{}
}

func (s *AthletesService) ListStarredSegments(athleteId int64) *AthletesListStarredSegmentsCall {
	return &AthletesListStarredSegmentsCall{
		service: s,
		id:      athleteId,
		ops:     make(map[string]interface{}),
	}
}

func (c *AthletesListStarredSegmentsCall) Page(page int) *AthletesListStarredSegmentsCall {
	c.ops["page"] = page
	return c
}

func (c *AthletesListStarredSegmentsCall) PerPage(perPage int) *AthletesListStarredSegmentsCall {
	c.ops["per_page"] = perPage
	return c
}

func (c *AthletesListStarredSegmentsCall) Do() ([]*PersonalSegmentSummary, error) {
	data, err := c.service.client.run("GET", fmt.Sprintf("/athletes/%d/segments/starred", c.id), c.ops)
	if err != nil {
		return nil, err
	}

	segments := make([]*PersonalSegmentSummary, 0)
	err = json.Unmarshal(data, &segments)
	if err != nil {
		return nil, err
	}

	for _, s := range segments {
		s.postProcess()
	}

	return segments, nil
}

/*********************************************************/

type AthletesListFriendsCall struct {
	service *AthletesService
	id      int64
	ops     map[string]interface{}
}

func (s *AthletesService) ListFriends(athleteId int64) *AthletesListFriendsCall {
	return &AthletesListFriendsCall{
		service: s,
		id:      athleteId,
		ops:     make(map[string]interface{}),
	}
}

func (c *AthletesListFriendsCall) Page(page int) *AthletesListFriendsCall {
	c.ops["page"] = page
	return c
}

func (c *AthletesListFriendsCall) PerPage(perPage int) *AthletesListFriendsCall {
	c.ops["per_page"] = perPage
	return c
}

func (c *AthletesListFriendsCall) Do() ([]*AthleteSummary, error) {
	data, err := c.service.client.run("GET", fmt.Sprintf("/athletes/%d/friends", c.id), c.ops)
	if err != nil {
		return nil, err
	}

	friends := make([]*AthleteSummary, 0)
	err = json.Unmarshal(data, &friends)
	if err != nil {
		return nil, err
	}

	for _, a := range friends {
		a.postProcessSummary()
	}

	return friends, nil
}

/*********************************************************/

type AthletesListFollowersCall struct {
	service *AthletesService
	id      int64
	ops     map[string]interface{}
}

func (s *AthletesService) ListFollowers(athleteId int64) *AthletesListFollowersCall {
	return &AthletesListFollowersCall{
		service: s,
		id:      athleteId,
		ops:     make(map[string]interface{}),
	}
}

func (c *AthletesListFollowersCall) Page(page int) *AthletesListFollowersCall {
	c.ops["page"] = page
	return c
}

func (c *AthletesListFollowersCall) PerPage(perPage int) *AthletesListFollowersCall {
	c.ops["per_page"] = perPage
	return c
}

func (c *AthletesListFollowersCall) Do() ([]*AthleteSummary, error) {
	data, err := c.service.client.run("GET", fmt.Sprintf("/athletes/%d/followers", c.id), c.ops)
	if err != nil {
		return nil, err
	}

	followers := make([]*AthleteSummary, 0)
	err = json.Unmarshal(data, &followers)
	if err != nil {
		return nil, err
	}

	for _, a := range followers {
		a.postProcessSummary()
	}

	return followers, nil
}

/*********************************************************/

type AthletesListBothFollowingCall struct {
	service *AthletesService
	id      int64
	ops     map[string]interface{}
}

func (s *AthletesService) ListBothFollowing(athleteId int64) *AthletesListBothFollowingCall {
	return &AthletesListBothFollowingCall{
		service: s,
		id:      athleteId,
		ops:     make(map[string]interface{}),
	}
}

func (c *AthletesListBothFollowingCall) Page(page int) *AthletesListBothFollowingCall {
	c.ops["page"] = page
	return c
}

func (c *AthletesListBothFollowingCall) PerPage(perPage int) *AthletesListBothFollowingCall {
	c.ops["per_page"] = perPage
	return c
}

func (c *AthletesListBothFollowingCall) Do() ([]*AthleteSummary, error) {
	data, err := c.service.client.run("GET", fmt.Sprintf("/athletes/%d/both-following", c.id), c.ops)
	if err != nil {
		return nil, err
	}

	athletes := make([]*AthleteSummary, 0)
	err = json.Unmarshal(data, &athletes)
	if err != nil {
		return nil, err
	}

	for _, a := range athletes {
		a.postProcessSummary()
	}

	return athletes, nil
}

/*********************************************************/

type AthletesListKOMsCall struct {
	service *AthletesService
	id      int64
	ops     map[string]interface{}
}

func (s *AthletesService) ListKOMs(athleteId int64) *AthletesListKOMsCall {
	return &AthletesListKOMsCall{
		service: s,
		id:      athleteId,
		ops:     make(map[string]interface{}),
	}
}

func (c *AthletesListKOMsCall) Page(page int) *AthletesListKOMsCall {
	c.ops["page"] = page
	return c
}

func (c *AthletesListKOMsCall) PerPage(perPage int) *AthletesListKOMsCall {
	c.ops["per_page"] = perPage
	return c
}

func (c *AthletesListKOMsCall) Do() ([]*SegmentEffortSummary, error) {
	data, err := c.service.client.run("GET", fmt.Sprintf("/athletes/%d/koms", c.id), c.ops)
	if err != nil {
		return nil, err
	}

	efforts := make([]*SegmentEffortSummary, 0)
	err = json.Unmarshal(data, &efforts)
	if err != nil {
		return nil, err
	}

	for _, e := range efforts {
		e.postProcessSummary()
	}

	return efforts, nil
}

/*********************************************************/

type AthletesListActivitiesCall struct {
	service *AthletesService
	id      int64
	ops     map[string]interface{}
}

func (s *AthletesService) ListActivities(athleteId int64) *AthletesListActivitiesCall {
	return &AthletesListActivitiesCall{
		service: s,
		id:      athleteId,
		ops:     make(map[string]interface{}),
	}
}

func (c *AthletesListActivitiesCall) Before(before int64) *AthletesListActivitiesCall {
	c.ops["before"] = before
	return c
}

func (c *AthletesListActivitiesCall) After(after int64) *AthletesListActivitiesCall {
	c.ops["after"] = after
	return c
}

func (c *AthletesListActivitiesCall) Page(page int) *AthletesListActivitiesCall {
	c.ops["page"] = page
	return c
}

func (c *AthletesListActivitiesCall) PerPage(perPage int) *AthletesListActivitiesCall {
	c.ops["per_page"] = perPage
	return c
}

func (c *AthletesListActivitiesCall) Do() ([]*ActivitySummary, error) {
	data, err := c.service.client.run("GET", fmt.Sprintf("/athletes/%d/activities", c.id), c.ops)
	if err != nil {
		return nil, err
	}

	activities := make([]*ActivitySummary, 0)
	err = json.Unmarshal(data, &activities)
	if err != nil {
		return nil, err
	}

	for _, e := range activities {
		e.postProcessSummary()
	}

	return activities, nil
}

/*********************************************************/

func (a *AthleteDetailed) postProcessDetailed() {
	a.postProcessSummary()
}

func (a *AthleteSummary) postProcessSummary() {
	a.CreatedAt, _ = time.Parse(timeFormat, a.CreatedAtString)
	a.UpdatedAt, _ = time.Parse(timeFormat, a.UpdatedAtString)
}
