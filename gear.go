package strava

import (
	"encoding/json"
)

type GearDetailed struct {
	GearSummary
	BrandName   string    `json:"brand_name"`
	ModelName   string    `json:"model_name"`
	FrameType   FrameType `json:"frame_type"`
	Description string    `json:"description"`
}

type GearSummary struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Primary  bool    `json:"primary"`
	Distance float64 `json:"distance"`
}

type FrameType int

var FrameTypes = struct {
	MountainBike FrameType
	Cyclocross   FrameType
	Road         FrameType
	TimeTrial    FrameType
}{1, 2, 3, 4}

type GearService struct {
	client *Client
}

func NewGearService(client *Client) *GearService {
	return &GearService{client}
}

/*********************************************************/

type GearGetCall struct {
	service *GearService
	id      string
}

func (s *GearService) Get(gearId string) *GearGetCall {
	return &GearGetCall{
		service: s,
		id:      gearId,
	}
}

func (c *GearGetCall) Do() (*GearDetailed, error) {
	data, err := c.service.client.run("GET", "/gear/"+c.id, nil)
	if err != nil {
		return nil, err
	}

	var gear GearDetailed
	err = json.Unmarshal(data, &gear)
	if err != nil {
		return nil, err
	}

	return &gear, nil
}

func (f FrameType) Id() int {
	return int(f)
}

func (f FrameType) String() string {
	switch f.Id() {
	case 1:
		return "Mountain Bike"
	case 2:
		return "Cyclocross"
	case 3:
		return "Road"
	case 4:
		return "Time Trial"
	}

	return "Bike Frame"
}
