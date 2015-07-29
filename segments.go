package strava

import (
	"encoding/json"
	"fmt"
	"time"
)

type SegmentDetailed struct {
	SegmentSummary

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	TotalElevationGain float64 `json:"total_elevation_gain"`
	Map                struct {
		Id       string   `json:"id"`
		Polyline Polyline `json:"polyline"`
	} `json:"map"`
	EffortCount  int  `json:"effort_count"`
	AthleteCount int  `json:"athlete_count"`
	StarCount    int  `json:"star_count"`
	Hazardous    bool `json:"hazardous"`
}

type SegmentSummary struct {
	Id            int64         `json:"id"`
	Name          string        `json:"name"`
	ActivityType  ActivityType  `json:"activity_type"`
	Distance      float64       `json:"distance"`
	AverageGrade  float64       `json:"average_grade"`
	MaximumGrade  float64       `json:"maximum_grade"`
	ElevationHigh float64       `json:"elevation_high"`
	ElevationLow  float64       `json:"elevation_low"`
	ClimbCategory ClimbCategory `json:"climb_category"`
	StartLocation Location      `json:"start_latlng"`
	EndLocation   Location      `json:"end_latlng"`
	City          string        `json:"city"`
	State         string        `json:"state"`
	Country       string        `json:"country"`
	Private       bool          `json:"private"`
	Starred       bool          `json:"starred"`
}

type PersonalSegmentSummary struct {
	SegmentSummary
	AthletePR struct {
		Id             int64     `json:"id"`
		ElapsedTime    int       `json:"elapsed_time"`
		Distance       float64   `json:"distance"`
		StartDate      time.Time `json:"start_date"`
		StartDateLocal time.Time `json:"start_date_local"`
		IsKOM          bool      `json:"is_kom"`
	} `json:"athlete_pr_effort"`
	StarredDate time.Time `json:"starred_date"`
}

type SegmentLeaderboard struct {
	EntryCount int                        `json:"entry_count"`
	Entries    []*SegmentLeaderboardEntry `json:"entries"`
}

type SegmentLeaderboardEntry struct {
	AthleteName      string    `json:"athlete_name"`
	AthleteId        int64     `json:"athlete_id"`
	AthleteGender    Gender    `json:"athlete_gender"`
	AverageHeartrate float64   `json:"average_hr"`
	AveragePower     float64   `json:"average_watts"`
	Distance         float64   `json:"distance"`
	ElapsedTime      int       `json:"elapsed_time"`
	MovingTime       int       `json:"moving_time"`
	StartDate        time.Time `json:"start_date"`
	StartDateLocal   time.Time `json:"start_date_local"`
	ActivityId       int64     `json:"activity_id"`
	EffortId         int64     `json:"effort_id"`
	Rank             int       `json:"rank"`
	AthleteProfile   string    `json:"athlete_profile"`
}

type segmentExplorer struct {
	Segments []*SegmentExplorerSegment `json:"segments"`
}

type SegmentExplorerSegment struct {
	Id                  int64         `json:"id"`
	Name                string        `json:"name"`
	ClimbCategory       ClimbCategory `json:"climb_category"`
	AverageGrade        float64       `json:"avg_grade"`
	StartLocation       Location      `json:"start_latlng"`
	EndLocation         Location      `json:"end_latlng"`
	ElevationDifference float64       `json:"elev_difference"`
	Distance            float64       `json:"distance"`
	Polyline            Polyline      `json:"points"`
}

type ClimbCategory int

var ClimbCategories = struct {
	NotCategorized ClimbCategory
	Category4      ClimbCategory
	Category3      ClimbCategory
	Category2      ClimbCategory
	Category1      ClimbCategory
	HorsCategorie  ClimbCategory
}{0, 1, 2, 3, 4, 5}

type AgeGroup string

var AgeGroups = struct {
	From0to24  AgeGroup
	From25to34 AgeGroup
	From35to44 AgeGroup
	From45to54 AgeGroup
	From55to64 AgeGroup
	From65plus AgeGroup
}{"0_24", "25_34", "35_44", "45_54", "55_64", "65_plus"}

type WeightClass string

var WeightClasses = struct {
	From0To125Pounds   WeightClass
	From125To149Pounds WeightClass
	From150To164Pounds WeightClass
	From165To179Pounds WeightClass
	From180To199Pounds WeightClass
	From200PlusPounds  WeightClass

	From0To54Kilograms  WeightClass
	From55To64Kilograms WeightClass
	From65To74Kilograms WeightClass
	From75To84Kilograms WeightClass
	From85To94Kilograms WeightClass
	From95PlusKilograms WeightClass
}{"0_124", "125_149", "150_164", "165_179", "180_199", "200_plus", "0_54", "55_64", "65_74", "75_84", "85_94", "95_plus"}

type DateRange string

var DateRanges = struct {
	ThisYear  DateRange
	ThisMonth DateRange
	ThisWeek  DateRange
	Today     DateRange
}{"this_year", "this_month", "this_week", "today"}

type SegmentsService struct {
	client *Client
}

func NewSegmentsService(client *Client) *SegmentsService {
	return &SegmentsService{client}
}

/*********************************************************/

type SegmentsGetCall struct {
	service *SegmentsService
	id      int64
}

func (s *SegmentsService) Get(segmentId int64) *SegmentsGetCall {
	return &SegmentsGetCall{
		service: s,
		id:      segmentId,
	}
}

func (s *SegmentsGetCall) Do() (*SegmentDetailed, error) {
	data, err := s.service.client.run("GET", fmt.Sprintf("/segments/%d", s.id), nil)
	if err != nil {
		return nil, err
	}

	var segment SegmentDetailed
	err = json.Unmarshal(data, &segment)
	if err != nil {
		return nil, err
	}

	return &segment, nil
}

/*********************************************************/

type SegmentsListEffortsCall struct {
	service *SegmentsService
	id      int64
	ops     map[string]interface{}
}

func (s *SegmentsService) ListEfforts(segmentId int64) *SegmentsListEffortsCall {
	return &SegmentsListEffortsCall{
		service: s,
		id:      segmentId,
		ops:     make(map[string]interface{}),
	}
}

func (c *SegmentsListEffortsCall) AthleteId(athleteId int64) *SegmentsListEffortsCall {
	c.ops["athlete_id"] = athleteId
	return c
}

func (c *SegmentsListEffortsCall) DateRange(startDateLocal, endDateLocal time.Time) *SegmentsListEffortsCall {
	c.ops["start_date_local"] = startDateLocal.UTC().Format(timeFormat)
	c.ops["end_date_local"] = endDateLocal.UTC().Format(timeFormat)
	return c
}

func (c *SegmentsListEffortsCall) Page(page int) *SegmentsListEffortsCall {
	c.ops["page"] = page
	return c
}

func (c *SegmentsListEffortsCall) PerPage(perPage int) *SegmentsListEffortsCall {
	c.ops["per_page"] = perPage
	return c
}

func (c *SegmentsListEffortsCall) Do() ([]*SegmentEffortSummary, error) {
	data, err := c.service.client.run("GET", fmt.Sprintf("/segments/%d/all_efforts", c.id), c.ops)
	if err != nil {
		return nil, err
	}

	efforts := make([]*SegmentEffortSummary, 0)
	err = json.Unmarshal(data, &efforts)
	if err != nil {
		return nil, err
	}

	return efforts, nil
}

/*********************************************************/

type SegmentsGetLeaderboardCall struct {
	service *SegmentsService
	id      int64
	ops     map[string]interface{}
}

func (s *SegmentsService) GetLeaderboard(segmentId int64) *SegmentsGetLeaderboardCall {
	return &SegmentsGetLeaderboardCall{
		service: s,
		id:      segmentId,
		ops:     make(map[string]interface{}),
	}
}

func (c *SegmentsGetLeaderboardCall) Gender(gender Gender) *SegmentsGetLeaderboardCall {
	c.ops["gender"] = gender
	return c
}

func (c *SegmentsGetLeaderboardCall) AgeGroup(ageGroup AgeGroup) *SegmentsGetLeaderboardCall {
	c.ops["age_group"] = ageGroup
	return c
}

func (c *SegmentsGetLeaderboardCall) WeightClass(class WeightClass) *SegmentsGetLeaderboardCall {
	c.ops["weight_class"] = class
	return c
}

func (c *SegmentsGetLeaderboardCall) Following() *SegmentsGetLeaderboardCall {
	c.ops["following"] = true
	return c
}

func (c *SegmentsGetLeaderboardCall) ClubId(clubId int64) *SegmentsGetLeaderboardCall {
	c.ops["club_id"] = clubId
	return c
}

func (c *SegmentsGetLeaderboardCall) DateRange(dateRange DateRange) *SegmentsGetLeaderboardCall {
	c.ops["date_range"] = dateRange
	return c
}

func (c *SegmentsGetLeaderboardCall) ContextEntries(count int) *SegmentsGetLeaderboardCall {
	c.ops["context_entries"] = count
	return c
}

func (c *SegmentsGetLeaderboardCall) Page(page int) *SegmentsGetLeaderboardCall {
	c.ops["page"] = page
	return c
}

func (c *SegmentsGetLeaderboardCall) PerPage(perPage int) *SegmentsGetLeaderboardCall {
	c.ops["per_page"] = perPage
	return c
}

func (c *SegmentsGetLeaderboardCall) Do() (*SegmentLeaderboard, error) {
	data, err := c.service.client.run("GET", fmt.Sprintf("/segments/%d/leaderboard", c.id), c.ops)
	if err != nil {
		return nil, err
	}

	var leaderboard SegmentLeaderboard
	err = json.Unmarshal(data, &leaderboard)
	if err != nil {
		return nil, err
	}

	return &leaderboard, nil
}

/*********************************************************/

type SegmentsExplorerCall struct {
	service *SegmentsService
	ops     map[string]interface{}
}

func (s *SegmentsService) Explore(south, west, north, east float64) *SegmentsExplorerCall {
	call := &SegmentsExplorerCall{
		service: s,
		ops:     make(map[string]interface{}),
	}

	call.ops["bounds"] = fmt.Sprintf("%f,%f,%f,%f", south, west, north, east)
	return call
}

func (c *SegmentsExplorerCall) ActivityType(activityType string) *SegmentsExplorerCall {
	c.ops["activity_type"] = activityType
	return c
}

func (c *SegmentsExplorerCall) MinimumCategory(cat int) *SegmentsExplorerCall {
	c.ops["min_cat"] = cat
	return c
}

func (c *SegmentsExplorerCall) MaximumCategory(cat int) *SegmentsExplorerCall {
	c.ops["max_cat"] = cat
	return c
}

func (c *SegmentsExplorerCall) Do() ([]*SegmentExplorerSegment, error) {
	data, err := c.service.client.run("GET", "/segments/explore", c.ops)
	if err != nil {
		return nil, err
	}

	var explorer segmentExplorer
	err = json.Unmarshal(data, &explorer)
	if err != nil {
		return nil, err
	}

	return explorer.Segments, nil
}

/*********************************************************/

func (c ClimbCategory) Id() int {
	return int(c)
}

func (c ClimbCategory) String() string {
	switch c.Id() {
	case 1:
		return "Category 4"
	case 2:
		return "Category 3"
	case 3:
		return "Category 2"
	case 4:
		return "Category 1"
	case 5:
		return "Hors Categorie"
	}

	return "Not Categorized"
}
