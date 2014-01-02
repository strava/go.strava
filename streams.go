package strava

import (
	"encoding/json"
	"errors"
	"fmt"
)

type StreamSet struct {
	Time        *IntegerStream
	Location    *LocationStream
	Distance    *DecimalStream
	Altitude    *DecimalStream
	Speed       *DecimalStream
	HeartRate   *IntegerStream
	Cadence     *IntegerStream
	Power       *IntegerStream
	Temperature *IntegerStream
	Moving      *BooleanStream
	Grade       *DecimalStream
}

type LocationStream struct {
	Stream
	Data [][2]float64
}

type IntegerStream struct {
	Stream
	Data []int
}

type DecimalStream struct {
	Stream
	Data []float64
}

type BooleanStream struct {
	Stream
	Data []bool
}

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
	Altitude    StreamType
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
		return nil, errors.New("Invalid stream parent type")
	}

	if len(c.types) == 0 {
		return nil, errors.New("No streamtypes requested")
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
			set.Time = &IntegerStream{s, nil}
			base = set.Time
		case StreamTypes.Location:
			set.Location = &LocationStream{s, nil}
			base = set.Location

		case StreamTypes.Distance:
			set.Distance = &DecimalStream{s, nil}
			base = set.Distance

		case StreamTypes.Altitude:
			set.Altitude = &DecimalStream{s, nil}
			base = set.Altitude

		case StreamTypes.Speed:
			set.Speed = &DecimalStream{s, nil}
			base = set.Speed

		case StreamTypes.HeartRate:
			set.HeartRate = &IntegerStream{s, nil}
			base = set.HeartRate

		case StreamTypes.Cadence:
			set.Cadence = &IntegerStream{s, nil}
			base = set.Cadence

		case StreamTypes.Power:
			set.Power = &IntegerStream{s, nil}
			base = set.Power

		case StreamTypes.Temperature:
			set.Temperature = &IntegerStream{s, nil}
			base = set.Temperature

		case StreamTypes.Moving:
			set.Moving = &BooleanStream{s, nil}
			base = set.Moving

		case StreamTypes.Grade:
			set.Grade = &DecimalStream{s, nil}
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
			s.Data[i][0] = l[0].(float64)
			s.Data[i][1] = l[1].(float64)
		}
	}
}

func (s *IntegerStream) fill(data []interface{}) {
	s.Data = make([]int, len(data))
	for i, v := range data {
		s.Data[i] = int(v.(float64))
	}
}

func (s *DecimalStream) fill(data []interface{}) {
	s.Data = make([]float64, len(data))
	for i, v := range data {
		s.Data[i] = v.(float64)
	}
}

func (s *BooleanStream) fill(data []interface{}) {
	s.Data = make([]bool, len(data))
	for i, v := range data {
		s.Data[i] = v.(bool)
	}
}
