// segment_example.go provides a simple example to fetch a segment details
// and list the top 10 on the leaderboard.
//
// usage:
//   > go get github.com/strava/go.strava
//   > cd $GOPATH/github.com/strava/go.strava/examples
//   > go run segment_example.go -id=segment_id -token=access_token
//
//   You can find an access_token for your app at https://www.strava.com/settings/api
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/strava/go.strava"
)

func main() {
	var segmentId int64
	var accessToken string

	// Provide an access token, with write permissions.
	// You'll need to complete the oauth flow to get one.
	flag.Int64Var(&segmentId, "id", 229781, "Strava Segment Id")
	flag.StringVar(&accessToken, "token", "", "Access Token")

	flag.Parse()

	if accessToken == "" {
		fmt.Println("\nPlease provide an access_token, one can be found at https://www.strava.com/settings/api")

		flag.PrintDefaults()
		os.Exit(1)
	}

	client := strava.NewClient(accessToken)

	fmt.Printf("Fetching segment %d info...\n", segmentId)
	segment, err := strava.NewSegmentsService(client).Get(segmentId).Do()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	verb := "ridden"
	if segment.ActivityType == strava.ActivityTypes.Run {
		verb = "run"
	}
	fmt.Printf("%s, %s %d times by %d athletes\n\n", segment.Name, verb, segment.EffortCount, segment.AthleteCount)

	fmt.Printf("Fetching leaderboard...\n")
	results, err := strava.NewSegmentsService(client).GetLeaderboard(segmentId).Do()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, e := range results.Entries {
		fmt.Printf("%5d: %5d %s\n", e.Rank, e.ElapsedTime, e.AthleteName)
	}
}
