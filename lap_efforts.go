package strava

type LapEffortSummary struct {
	EffortSummary
	TotalElevationGain float64 `json:"total_elevation_gain"`
	AverageSpeed       float64 `json:"average_speed"`
	MaximunSpeed       float64 `json:"max_speed"`
	AverageCadence     float64 `json:"average_cadence"`
	AveragePower       float64 `json:"average_watts"`
	AverageHeartrate   float64 `json:"average_heartrate"`
	MaximumHeartrate   float64 `json:"max_heartrate"`
	LapIndex           int     `json:"lap_index"`
}
