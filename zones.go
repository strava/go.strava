package strava

type ZonesSummary struct {
	Score         int           `json:"score"`
	Buckets       []*ZoneBucket `json:"distribution_buckets"`
	Type          string        `json:"type"`           // power or heartrate
	SensorBased   bool          `json:"sensor_based"`   // estimated power?
	Points        int           `json:"points"`         // heartrate only
	CustonZones   bool          `json:"custom_zones"`   // heartrate only
	Max           int           `json:"max"`            // heartrate only
	BikeWeight    float64       `json:"bike_weight"`    // power only
	AthleteWeight float64       `json:"athlete_weight"` // power only
}

type ZoneBucket struct {
	Min  int `json:"min"`
	Max  int `json:"max"`
	Time int `json:"time"`
}

func (a *ZonesSummary) postProcessSummary() {

}
