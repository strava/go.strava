package strava

type ZonesSummary struct {
	Score       int           `json:"score"`
	Buckets     []*ZoneBucket `json:"distribution_buckets"`
	Type        string        `json:"type"`         // power or heartrate
	SensorBased bool          `json:"sensor_based"` // estimated power?
	Points      int           `json:"points"`       // heartrate only
	CustonZones bool          `json:"custom_zones"` // heartrate only
}

type ZoneBucket struct {
	// legacy, left intact for compatibility
	Min int
	Max int
	// precision fields
	MinFloat float64 `json:"min"`
	MaxFloat float64 `json:"max"`
	Time     int     `json:"time"`
}
