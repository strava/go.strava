package strava

import (
	"encoding/json"
	"fmt"
	"time"
)

type ActivityDetailed struct {
	ActivitySummary
	Calories       float64                 `json:"calories"`
	Description    string                  `json:"description"`
	Gear           GearSummary             `json:"gear"`
	SegmentEfforts []*SegmentEffortSummary `json:"segment_efforts"`
	SplitsMetric   []*Split                `json:"splits_metric"`
	SplitsStandard []*Split                `json:"splits_standard"`
	BestEfforts    []*BestEffort           `json:"best_efforts"`
}

type ActivitySummary struct {
	Id                 int64          `json:"id"`
	ResourceState      int            `json:"resource_state"`
	ExternalId         string         `json:"external_id"`
	UploadId           int64          `json:"upload_id"`
	Athlete            AthleteSummary `json:"athlete"`
	Name               string         `json:"name"`
	Distance           float64        `json:"distance"`
	MovingTime         int            `json:"moving_time"`
	ElapsedTime        int            `json:"elapsed_time"`
	TotalElevationGain float64        `json:"total_elevation_gain"`
	Type               ActivityType   `json:"type"`

	StartDate            time.Time `json:"-"`
	StartDateLocal       time.Time `json:"-"`
	StartDateString      string    `json:"start_date"`       // the ISO 8601 encoding of when the effort started
	StartDateLocalString string    `json:"start_date_local"` // the ISO 8601 encoding of the UTC version of the local time when the effort started, see: http://strava.github.io/api/#dates

	TimeZone         string   `json:"time_zone"`
	StartLocation    Location `json:"start_latlng"`
	EndLocation      Location `json:"end_latlng"`
	City             string   `json:"location_city"`
	State            string   `json:"location_state"`
	Country          string   `json:"location_country"`
	AchievementCount int      `json:"achievement_count"`
	KudosCount       int      `json:"kudos_count"`
	CommentCount     int      `json:"comment_count"`
	AthleteCount     int      `json:"athlete_count"`
	PhotoCount       int      `json:"photo_count"`
	Map              struct {
		Id              string   `json:"id"`
		Polyline        Polyline `json:"polyline"`
		SummaryPolyline Polyline `json:"summary_polyline"`
	} `json:"map"`
	Trainer            bool    `json:"trainer"`
	Commute            bool    `json:"commute"`
	Manual             bool    `json:"manual"`
	Private            bool    `json:"private"`
	Flagged            bool    `json:"flagged"`
	GearId             string  `json:"gear_id"` // bike or pair of shoes
	AverageSpeed       float64 `json:"average_speed"`
	MaximunSpeed       float64 `json:"max_speed"`
	AverageCadence     float64 `json:"average_cadence"`
	AverageTemperature float64 `json:"average_temp"`
	AveragePower       float64 `json:"average_watts"`
	Kilojoules         float64 `json:"kilojoules"`
	AverageHeartrate   float64 `json:"average_heartrate"`
	MaximumHeartrate   float64 `json:"max_heartrate"`
	Truncated          int     `json:"truncated"` // only present if activity is owned by authenticated athlete, returns 0 if not truncated by privacy zones
	HasKudoed          bool    `json:"has_kudoed"`
}

type BestEffort struct {
	EffortSummary
	PRRank int `json:"pr_rank"` // 1-3 personal record on segment at time of upload
}

type ActivityType string

var ActivityTypes = struct {
	Ride           ActivityType
	AlpineSki      ActivityType
	BackcountrySki ActivityType
	Hike           ActivityType
	IceSkate       ActivityType
	InlineSkate    ActivityType
	NordicSki      ActivityType
	RollerSki      ActivityType
	Run            ActivityType
	Walk           ActivityType
	Workout        ActivityType
	Snowboard      ActivityType
	Snowshoe       ActivityType
	Kitesurf       ActivityType
	Windsurf       ActivityType
	Swim           ActivityType
}{"Ride", "AlpineSki", "BackcountrySki", "Hike", "IceSkate", "InlineSkate", "NordicSki", "RollerSki",
	"Run", "Walk", "Workout", "Snowboard", "Snowshoe", "Kitesurf", "Windsurf", "Swim"}

type Location [2]float64

type ActivitiesService struct {
	client *Client
}

func NewActivitiesService(client *Client) *ActivitiesService {
	return &ActivitiesService{client}
}

/*********************************************************/

type ActivitiesGetCall struct {
	service *ActivitiesService
	id      int64
	ops     map[string]interface{}
}

func (s *ActivitiesService) Get(activityId int64) *ActivitiesGetCall {
	return &ActivitiesGetCall{
		service: s,
		id:      activityId,
		ops:     make(map[string]interface{}),
	}
}

func (c *ActivitiesGetCall) IncludeAllEfforts() *ActivitiesGetCall {
	c.ops["include_all_efforts"] = true
	return c
}

func (c *ActivitiesGetCall) Do() (*ActivityDetailed, error) {
	data, err := c.service.client.run("GET", fmt.Sprintf("/activities/%d", c.id), c.ops)
	if err != nil {
		return nil, err
	}

	var activity ActivityDetailed
	err = json.Unmarshal(data, &activity)
	if err != nil {
		return nil, err
	}

	activity.postProcessDetailed()

	return &activity, nil
}

/*********************************************************/

type ActivitiesPutCall struct {
	service *ActivitiesService
	id      int64
	ops     map[string]interface{}
}

func (s *ActivitiesService) Update(activityId int64) *ActivitiesPutCall {
	return &ActivitiesPutCall{
		service: s,
		id:      activityId,
		ops:     make(map[string]interface{}),
	}
}

func (c *ActivitiesPutCall) Name(name string) *ActivitiesPutCall {
	c.ops["name"] = name
	return c
}

func (c *ActivitiesPutCall) Description(description string) *ActivitiesPutCall {
	c.ops["description"] = description
	return c
}

func (c *ActivitiesPutCall) Type(activityType ActivityType) *ActivitiesPutCall {
	c.ops["type"] = string(activityType)
	return c
}

func (c *ActivitiesPutCall) Private(isPrivate bool) *ActivitiesPutCall {
	c.ops["private"] = isPrivate
	return c
}

func (c *ActivitiesPutCall) Commute(isCommute bool) *ActivitiesPutCall {
	c.ops["commute"] = isCommute
	return c
}

func (c *ActivitiesPutCall) Trainer(isTrainer bool) *ActivitiesPutCall {
	c.ops["trainer"] = isTrainer
	return c
}

func (c *ActivitiesPutCall) Gear(gearId string) *ActivitiesPutCall {
	c.ops["gear_id"] = gearId
	return c
}

func (c *ActivitiesPutCall) Do() (*ActivityDetailed, error) {
	data, err := c.service.client.run("PUT", fmt.Sprintf("/activities/%d", c.id), c.ops)
	if err != nil {
		return nil, err
	}

	var activity ActivityDetailed
	err = json.Unmarshal(data, &activity)
	if err != nil {
		return nil, err
	}

	activity.postProcessDetailed()

	return &activity, nil
}

/*********************************************************/

type ActivitiesListPhotosCall struct {
	service *ActivitiesService
	id      int64
}

func (s *ActivitiesService) ListPhotos(activityId int64) *ActivitiesListPhotosCall {
	return &ActivitiesListPhotosCall{
		service: s,
		id:      activityId,
	}
}

func (c *ActivitiesListPhotosCall) Do() ([]*PhotoSummary, error) {
	data, err := c.service.client.run("GET", fmt.Sprintf("/activities/%d/photos", c.id), nil)
	if err != nil {
		return nil, err
	}

	photos := make([]*PhotoSummary, 0)
	err = json.Unmarshal(data, &photos)
	if err != nil {
		return nil, err
	}

	for _, p := range photos {
		p.postProcessSummary()
	}

	return photos, nil
}

/*********************************************************/

type ActivitiesListZonesCall struct {
	service *ActivitiesService
	id      int64
}

func (s *ActivitiesService) ListZones(activityId int64) *ActivitiesListZonesCall {
	return &ActivitiesListZonesCall{
		service: s,
		id:      activityId,
	}
}

func (c *ActivitiesListZonesCall) Do() ([]*ZonesSummary, error) {
	data, err := c.service.client.run("GET", fmt.Sprintf("/activities/%d/zones", c.id), nil)
	if err != nil {
		return nil, err
	}

	zones := make([]*ZonesSummary, 0)
	err = json.Unmarshal(data, &zones)
	if err != nil {
		return nil, err
	}

	for _, z := range zones {
		z.postProcessSummary()
	}

	return zones, nil
}

/*********************************************************/

type ActivitiesListLapsCall struct {
	service *ActivitiesService
	id      int64
}

func (s *ActivitiesService) ListLaps(activityId int64) *ActivitiesListLapsCall {
	return &ActivitiesListLapsCall{
		service: s,
		id:      activityId,
	}
}

func (c *ActivitiesListLapsCall) Do() ([]*LapEffortSummary, error) {
	data, err := c.service.client.run("GET", fmt.Sprintf("/activities/%d/laps", c.id), nil)
	if err != nil {
		return nil, err
	}

	laps := make([]*LapEffortSummary, 0)
	err = json.Unmarshal(data, &laps)
	if err != nil {
		return nil, err
	}

	for _, l := range laps {
		l.postProcessSummary()
	}

	return laps, nil
}

/*********************************************************/

func (a *ActivityDetailed) postProcessDetailed() {
	for i := range a.SegmentEfforts {
		a.SegmentEfforts[i].postProcessSummary()
	}

	for i := range a.BestEfforts {
		a.BestEfforts[i].postProcessSummary()
	}

	a.postProcessSummary()
}

func (a *ActivitySummary) postProcessSummary() {
	a.Athlete.postProcessSummary()

	a.StartDate, _ = time.Parse(timeFormat, a.StartDateString)
	a.StartDateLocal, _ = time.Parse(timeFormat, a.StartDateLocalString)
}

/*********************************************************/

func (t ActivityType) Id() int {
	switch t {
	case ActivityTypes.Ride:
		return 1
	case ActivityTypes.AlpineSki:
		return 2
	case ActivityTypes.BackcountrySki:
		return 3
	case ActivityTypes.Hike:
		return 4
	case ActivityTypes.IceSkate:
		return 5
	case ActivityTypes.InlineSkate:
		return 6
	case ActivityTypes.NordicSki:
		return 7
	case ActivityTypes.RollerSki:
		return 8
	case ActivityTypes.Run:
		return 9
	case ActivityTypes.Walk:
		return 10
	case ActivityTypes.Workout:
		return 11
	case ActivityTypes.Snowboard:
		return 12
	case ActivityTypes.Snowshoe:
		return 13
	case ActivityTypes.Kitesurf:
		return 14
	case ActivityTypes.Windsurf:
		return 15
	case ActivityTypes.Swim:
		return 16
	}

	return 0
}

func (t ActivityType) String() string {
	switch t {
	case ActivityTypes.Ride:
		return "Ride"
	case ActivityTypes.AlpineSki:
		return "Alpine Ski"
	case ActivityTypes.BackcountrySki:
		return "Backcountry Ski"
	case ActivityTypes.Hike:
		return "Hike"
	case ActivityTypes.IceSkate:
		return "Ice Skate"
	case ActivityTypes.InlineSkate:
		return "Inline Skate"
	case ActivityTypes.NordicSki:
		return "Nordic Ski"
	case ActivityTypes.RollerSki:
		return "Roller Ski"
	case ActivityTypes.Run:
		return "Run"
	case ActivityTypes.Walk:
		return "Walk"
	case ActivityTypes.Workout:
		return "Workout"
	case ActivityTypes.Snowboard:
		return "Snowboard"
	case ActivityTypes.Snowshoe:
		return "Snowshoe"
	case ActivityTypes.Kitesurf:
		return "Kitesurf"
	case ActivityTypes.Windsurf:
		return "Windsurf"
	case ActivityTypes.Swim:
		return "Swim"
	}

	return "Activity"
}

func (l Location) String() string {
	return fmt.Sprintf("[%f, %f]", l[0], l[1])
}
