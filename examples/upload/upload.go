// upload.go provides a simple example of uploading a GPX file to Strava
//
// usage:
//   > go get github.com/strava/go.strava
//   > cd $GOPATH/github.com/strava/go.strava/examples
//   > go run upload.go -token=access_token
//
//   You will need an access token with 'write' permissions. You'll
//   need to complete the oauth flow to get one of those.
//
//   Note this will upload a dummy file to your account, so make sure you delete.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/strava/go.strava"
)

func main() {
	var accessToken string

	// Provide an access token, with write permissions.
	// You'll need to complete the oauth flow to get one.
	flag.StringVar(&accessToken, "token", "", "Access Token")
	flag.Parse()

	if accessToken == "" {
		fmt.Println("\nPlease provide an access_token")

		flag.PrintDefaults()
		os.Exit(1)
	}

	client := strava.NewClient(accessToken)
	service := strava.NewUploadsService(client)

	fmt.Printf("Uploading data...\n")

	upload, err := service.
		Create(strava.FileDataTypes.GPX, "test_file.gpx", strings.NewReader(rawGPXData())).
		Private().
		Do()
	if err != nil {
		if e, ok := err.(strava.Error); ok && e.Message == "Authorization Error" {
			log.Printf("Make sure your token has 'write' permissions. You'll need implement the oauth process to get one")
		}

		log.Fatal(err)
	}

	log.Printf("Upload Complete...")
	jsonForDisplay, _ := json.Marshal(upload)
	log.Printf(string(jsonForDisplay))

	log.Printf("Waiting a 5 seconds so the upload will finish (might not)")
	time.Sleep(5 * time.Second)

	uploadSummary, err := service.Get(upload.Id).Do()
	jsonForDisplay, _ = json.Marshal(uploadSummary)
	log.Printf(string(jsonForDisplay))

	log.Printf("Your new activity is id %d", uploadSummary.ActivityId)
	log.Printf("You can view it at http://www.strava.com/activities/%d", uploadSummary.ActivityId)
}

func rawGPXData() string {
	format := "2006-01-02T15:04:05Z"
	now := time.Now()

	return fmt.Sprintf(`
		<?xml version="1.0" encoding="UTF-8"?>
		<gpx creator="strava.com iPhone" version="1.1" xmlns="http://www.topografix.com/GPX/1/1" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.topografix.com/GPX/1/1 http://www.topografix.com/GPX/1/1/gpx.xsd">
			<metadata>
				<time>%s</time>
			</metadata>
			<trk>
				<name>Morning Ride</name>
				<trkseg>
					<trkpt lat="37.7737810" lon="-122.4669790">
						<ele>72.2</ele>
						<time>%s</time>
					</trkpt>
					<trkpt lat="37.7737580" lon="-122.4669550">
						<ele>72.2</ele>
						<time>%s</time>
					</trkpt>
				</trkseg>
			</trk>
		</gpx>`, now.Format(format), now.Format(format), now.Add(2*time.Second).Format(format))
}
