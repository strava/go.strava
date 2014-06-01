package strava

import (
	"encoding/json"
	"errors"
	"fmt"
)

// A StreamSet is a collection of possible streams for an Activity, Segment or SegmentEffort.
// Some types may be nil if they weren't requested or do not exist.
// Time is the only required stream for Uploaded activities.
// Manually created activties have no streams.
type StreamSet struct {
	Time        *IntegerStream  // Seconds from start
	Location    *LocationStream // [lat, lng] tuples
	Distance    *DecimalStream  // Distance in meters from start
	Elevation   *DecimalStream  // Elevation in meters
	Speed       *DecimalStream  // Speed in meters per second
	HeartRate   *IntegerStream  // Heart Rate in beats per minute if available
	Cadence     *IntegerStream  // Running or Cycling cadence in revolutions per minute
	Power       *IntegerStream  // Cycling power output in watts
	Temperature *IntegerStream  // Temperature in Celsius
	Moving      *BooleanStream  // Derived from speed and time to give some idea of when moving
	Grade       *DecimalStream  // Grade or pitch in the road in percent
}

// A LocationStream represents [lat, lng] data.
// Values of [0, 0], aka. Null Island, should be considered as nil or unavailable
type LocationStream struct {
	Stream
	Data [][2]float64
}

// An IntegerStream represents time series data that is integer based
// such as Time, HeartRate, Cadence, Power or Temperature.
// Any nil/unavailable values in the original data will be 0 in the Data slice.
// To determine if a value is nil, check for it in RawData slice
type IntegerStream struct {
	Stream
	Data    []int
	RawData []*int
}

// A DecimalStream represents time series data that is decimal based
// such as Distance, Elevation, Speed or Grade.
// Any nil/unavailable values in the original data will be 0 in the Data slice.
// To determine if a value is nil, check for it in RawData slice
type DecimalStream struct {
	Stream
	Data    []float64
	RawData []*float64
}

// A BooleanStream represents time series data that is binary, or yes/no in nature.
// An example is the Moving stream
type BooleanStream struct {
	Stream
	Data []bool
}

// A Stream represents time series data of a given type.
// A streams for a given object are the same length. For every time in the Time stream
// there will be corresponding information in all the other available streams.
type Stream struct {
	Type         StreamType `json:"type"`
	SeriesType   string     `json:"series_type"`
	OriginalSize int        `json:"original_size"`
	Resolution   string     `json:"resolution"`
}

type StreamType string

var StreamTypes = struct {
	Time        StreamType
	Location    StreamType
	Distance    StreamType
	Elevation   StreamType
	Speed       StreamType
	HeartRate   StreamType
	Cadence     StreamType
	Power       StreamType
	Temperature StreamType
	Moving      StreamType
	Grade       StreamType
}{"time", "latlng", "distance", "altitude", "velocity_smooth", "heartrate",
	"cadence", "watts", "temp", "moving", "grade_smooth"}

type filler interface {
	fill([]interface{})
}

type ActivityStreamsService streamsService

type SegmentStreamsService streamsService

type SegmentEffortStreamsService streamsService

type streamsService struct {
	client     *Client
	parentType int
}

var types = struct {
	Activity      int
	Segment       int
	SegmentEffort int
}{1, 2, 3}

func NewActivityStreamsService(client *Client) *ActivityStreamsService {
	return &ActivityStreamsService{client, types.Activity}
}

func NewSegmentStreamsService(client *Client) *SegmentStreamsService {
	return &SegmentStreamsService{client, types.Segment}
}

func NewSegmentEffortStreamsService(client *Client) *SegmentEffortStreamsService {
	return &SegmentEffortStreamsService{client, types.SegmentEffort}
}

/*********************************************************/

type ActivityStreamsGetCall struct {
	streamsGetCall
}

type SegmentStreamsGetCall struct {
	streamsGetCall
}

type SegmentEffortStreamsGetCall struct {
	streamsGetCall
}

type streamsGetCall struct {
	service streamsService
	id      int64
	types   []StreamType
	ops     map[string]interface{}
}

/*********************************************************/

func (s *ActivityStreamsService) Get(activityId int64, types []StreamType) *ActivityStreamsGetCall {
	call := &ActivityStreamsGetCall{}

	call.service = streamsService(*s)
	call.id = activityId
	call.ops = make(map[string]interface{})
	call.types = make([]StreamType, len(types))

	copy(call.types, types)

	return call
}

func (c *ActivityStreamsGetCall) Resolution(resolution string) *ActivityStreamsGetCall {
	c.ops["resolution"] = resolution
	return c
}

func (c *ActivityStreamsGetCall) SeriesType(seriesType string) *ActivityStreamsGetCall {
	c.ops["series_type"] = seriesType
	return c
}

/*********************************************************/

func (s *SegmentStreamsService) Get(segmentId int64, types []StreamType) *SegmentStreamsGetCall {
	call := &SegmentStreamsGetCall{}

	call.service = streamsService(*s)
	call.id = segmentId
	call.ops = make(map[string]interface{})
	call.types = make([]StreamType, len(types))

	copy(call.types, types)
	return call
}

func (c *SegmentStreamsGetCall) Resolution(resolution string) *SegmentStreamsGetCall {
	c.ops["resolution"] = resolution
	return c
}

func (c *SegmentStreamsGetCall) SeriesType(seriesType string) *SegmentStreamsGetCall {
	c.ops["series_type"] = seriesType
	return c
}

/*********************************************************/

func (s *SegmentEffortStreamsService) Get(segmentEffortId int64, types []StreamType) *SegmentEffortStreamsGetCall {
	call := &SegmentEffortStreamsGetCall{}

	call.service = streamsService(*s)
	call.id = segmentEffortId
	call.ops = make(map[string]interface{})
	call.types = make([]StreamType, len(types))

	copy(call.types, types)
	return call
}

func (c *SegmentEffortStreamsGetCall) Resolution(resolution string) *SegmentEffortStreamsGetCall {
	c.ops["resolution"] = resolution
	return c
}

func (c *SegmentEffortStreamsGetCall) SeriesType(seriesType string) *SegmentEffortStreamsGetCall {
	c.ops["series_type"] = seriesType
	return c
}

/*********************************************************/

func (c *streamsGetCall) Do() (*StreamSet, error) {
	var source string
	switch c.service.parentType {
	case types.Activity:
		source = "activities"
	case types.Segment:
		source = "segments"
	case types.SegmentEffort:
		source = "segment_efforts"
	}

	if source == "" {
		return nil, errors.New("invalid stream parent type")
	}

	if len(c.types) == 0 {
		return nil, errors.New("no streamtypes requested")
	}

	types := string(c.types[0])
	for i := 1; i < len(c.types); i++ {
		types += "," + string(c.types[i])
	}

	path := fmt.Sprintf("/%s/%d/streams/%s", source, c.id, types)
	data, err := c.service.client.run("GET", path, c.ops)

	if err != nil {
		return nil, err
	}

	var set StreamSet

	streams := make([]interface{}, 0, 10)
	json.Unmarshal(data, &streams)

	for _, stream := range streams {
		m := stream.(map[string]interface{})

		var base interface{}
		var s Stream
		s.Type = StreamType(m["type"].(string))
		s.SeriesType = m["series_type"].(string)
		s.OriginalSize = int(m["original_size"].(float64))
		s.Resolution = m["resolution"].(string)

		switch StreamType(m["type"].(string)) {
		case StreamTypes.Time:
			set.Time = &IntegerStream{s, nil, nil}
			base = set.Time

		case StreamTypes.Location:
			set.Location = &LocationStream{s, nil}
			base = set.Location

		case StreamTypes.Distance:
			set.Distance = &DecimalStream{s, nil, nil}
			base = set.Distance

		case StreamTypes.Elevation:
			set.Elevation = &DecimalStream{s, nil, nil}
			base = set.Elevation

		case StreamTypes.Speed:
			set.Speed = &DecimalStream{s, nil, nil}
			base = set.Speed

		case StreamTypes.HeartRate:
			set.HeartRate = &IntegerStream{s, nil, nil}
			base = set.HeartRate

		case StreamTypes.Cadence:
			set.Cadence = &IntegerStream{s, nil, nil}
			base = set.Cadence

		case StreamTypes.Power:
			set.Power = &IntegerStream{s, nil, nil}
			base = set.Power

		case StreamTypes.Temperature:
			set.Temperature = &IntegerStream{s, nil, nil}
			base = set.Temperature

		case StreamTypes.Moving:
			set.Moving = &BooleanStream{s, nil}
			base = set.Moving

		case StreamTypes.Grade:
			set.Grade = &DecimalStream{s, nil, nil}
			base = set.Grade
		}

		f := base.(filler)
		f.fill(m["data"].([]interface{}))

	}

	return &set, nil
}

func (s *LocationStream) fill(data []interface{}) {
	s.Data = make([][2]float64, len(data))
	for i, v := range data {
		if l, ok := v.([]interface{}); ok {
			s.Data[i] = [2]float64{l[0].(float64), l[1].(float64)}
		}
	}
}

func (s *IntegerStream) fill(data []interface{}) {
	s.Data = make([]int, len(data))
	s.RawData = make([]*int, len(data))
	for i, v := range data {
		if f, ok := v.(float64); ok {
			s.Data[i] = int(f)
			s.RawData[i] = &s.Data[i]
		}
	}
}

func (s *DecimalStream) fill(data []interface{}) {
	var ok bool

	s.Data = make([]float64, len(data))
	s.RawData = make([]*float64, len(data))
	for i, v := range data {
		if s.Data[i], ok = v.(float64); ok {
			s.RawData[i] = &s.Data[i]
		}
	}
}

func (s *BooleanStream) fill(data []interface{}) {
	s.Data = make([]bool, len(data))
	for i, v := range data {
		if f, ok := v.(bool); ok {
			s.Data[i] = f
		}
	}
}
