go.strava
=========

Go.strava provides a complete, **including upload,** wrapper library for the [Strava V3 API](http://strava.github.com/api).
Structs are defined for all the basic types such as athletes, activities and leaderboards. Functions are
provided to fetch all these values.

**Please note:** Every effort is made to keep the [API Documentation](http://strava.github.com/api)
up to date, however this package may lag behind. If this is impacting your development,
please file an [issue](https://github.com/strava/go.strava/issues) and it will be addressed as soon as possible.

#### To install
	
	go get github.com/strava/go.strava

#### To use, imports as package name `strava`:

	import "github.com/strava/go.strava"

<br />
[![Build Status](https://travis-ci.org/strava/go.strava.png?branch=master)](https://travis-ci.org/strava/go.strava)
&nbsp; &nbsp;
[![Coverage Status](https://coveralls.io/repos/strava/go.strava/badge.png?branch=master)](https://coveralls.io/r/strava/go.strava?branch=master)
&nbsp; &nbsp;
[![Godoc Reference](https://godoc.org/github.com/strava/go.strava?status.png)](https://godoc.org/github.com/strava/go.strava)
&nbsp; &nbsp;
[Official Strava API Documentation](http://strava.github.io/api)

<br />
### This document is separated into three parts

* [Examples](#examples) - working code
* [Service Documentation](#services) - what you can do and how to do it
* [Testing](#testing)

<a name="examples"></a>Examples
-------------------------------
Make sure to check out the source code as it provides many helpful comments.

* #### [segment_example.go](examples/segment_example.go)
	This example shows how to pull segment descriptions and leaderboards. To run:

		cd $GOPATH/src/github.com/strava/go.strava/examples
		go run segment_example.go -token=<your-access-token>

	A sample access token can be found on the [API settings page](https://strava.com/settings/api).

* #### [oauth_example.go](examples/oauth_example.go) 
	This example shows how to use `OAuthCallbackHandler` to simplify the OAuth2 token exchange, 
	as well as how to handle the errors that might occur. To run:

		cd $GOPATH/src/github.com/strava/go.strava/examples
		go run oauth_example.go -id=<your-client-id> -secret=<your-client-secret>

		Visit http://localhost:8080/ in your favorite web browser

	Your client id and client secret can be found on the [API settings page](https://strava.com/settings/api).

* #### [upload.go](examples/upload.go) 
	This example shows how to upload data to Strava. It will upload a random GPX file. To run:

		cd $GOPATH/src/github.com/strava/go.strava/examples
		go run upload.go -token=<your-access-token>

	The access token must have 'write' permissions which must be created using the OAuth flow, see the above example.


<a name="services"></a>Service Documentation
--------------------------------------------
For full documentation, see the [Godoc documentation](https://godoc.org/github.com/strava/go.strava).
Below is a overview of how the library works, followed by examples for all the different calls.

1. All requests should start by creating a client that defines the access token to be used:

		client := strava.NewClient("<an-access-token>", optionalHttpClient)

	The library will use the http.DefaultClient by default. If the default client is unavailable, 
	like in the app engine environment for example, you can pass one in as the second parameter.

2. Then a service must be defined that represents a given API request endpoint, for example:

		service := strava.NewClubsService(client)

3. Required parameters are passed on call creation, optional parameters are added after, for example:

		call := service.ListMembers(clubId).Page(2).PerPage(50)

4. To actually execute the call, run `Do()` on it:

		members, err := call.Do()
		if e, ok := err.(*strava.Error); ok {
			// this is a strava provided error
		} else {
			// regular error, could be internet connectivity problems
		}

	This will return members 50-100 of the given clubs. All of these things can be chained together like so:

		members, err := strava.NewClubsService(strava.NewClient(token)).
			ListMembers(clubId).
			PerPage(100).
			Do()

**Polyline decoding**  
Activities and segments come with summary polylines encoded using the
[Google Polyline Format](https://developers.google.com/maps/documentation/utilities/polylinealgorithm). 
These can be decoded into a slice of [2]float64 using `Decode()`, for example: 
`activity.Map.Polyline.Decode()`, `segment.Map.Polyline.Decode()`, or `segmentExplorerSegment.Polyline.Decode()`.

### Examples for all the possible calls can be found below:

* [Authentication](#Authentication)
* [Athletes](#Athletes)
* [Activities](#Activities)
* [Comments](#Comments)
* [Kudos](#Kudos)
* [Clubs](#Clubs)
* [Gear](#Gear)
* [Segments](#Segments)
* [Segment Efforts](#SegmentEfforts)
* [Streams](#Streams)
* [Uploads](#Uploads)

### <a name="Authentication"></a>Authentication

Related objects: 
[OAuthAuthenticator](https://godoc.org/github.com/strava/go.strava#OAuthAuthenticator),
[AuthorizationResponse](https://godoc.org/github.com/strava/go.strava#AuthorizationResponse).

	authenticator := &strava.OAuthAuthenticator{
		CallbackURL: "http://yourdomain/strava/authorize",
	}

	path, err := authenticator.CallbackPath()
	http.HandleFunc(path, authenticator.HandlerFunc(oAuthSuccess, oAuthFailure))

	func oAuthSuccess(auth *strava.AuthorizationResponse, w http.ResponseWriter, r *http.Request) {
		// Success
	}

	func oAuthFailure(err error, w http.ResponseWriter, r *http.Request) {
		// Failure, or access was denied
	}

For a more detailed example of how to handle OAuth authorization see [oauth_example.go](examples/oauth_example.go) 

### <a name="Athletes"></a>Athletes

Related objects: 
[AthleteDetailed](https://godoc.org/github.com/strava/go.strava#AthleteDetailed),
[AthleteSummary](https://godoc.org/github.com/strava/go.strava#AthleteSummary),
[AthleteMeta](https://godoc.org/github.com/strava/go.strava#AthleteSummary).
[AthleteStats](https://godoc.org/github.com/strava/go.strava#AthleteStats).
[PersonalSegmentSummary](https://godoc.org/github.com/strava/go.strava#PersonalSegmentSummary).

For the athlete associated with the access token, aka current athlete:

	service := strava.NewCurrentAthleteService(client)

	// returns a AthleteDetailed object
	athlete, err := service.Get().Do()

	// returns a AthleteDetailed object
	athlete, err := service.Update().
		City(city).
		State(state).
		Country(country).
		Gender(gender).
		Weight(weight).
		Do()


	// returns a slice of ActivitySummary objects
	activities, err := service.ListActivities(athleteId).
		Page(page).
		PerPage(perPage).
		Before(before).
		After(after).
		Do()

	// returns a slice of ActivitySummary objects
	activities, err := service.ListFriendsActivities(athleteId).
		Page(page).
		PerPage(perPage).
		Before(before).
		Do()

	// returns a slice of AthleteSummary objects
	friends, err := service.ListFriends().Page(page).PerPage(perPage).Do()
	
	// returns a slice of AthleteSummary objects
	followers, err := service.ListFollowers().Page(page).PerPage(perPage).Do()

	// returns a slice of ClubSummary objects
	clubs, err := service.ListClubs().Do()

	// returns a slice of PersonalSegmentSummary objects
	segments, err := service.ListStarredSegments().Do()

For other athletes:

	service := strava.NewAthletesService(client)

	// returns a AthleteSummary object
	athlete, err := service.Get(athleteId).Do()

	// returns a slice of AthleteSummary objects
	friends, err := service.ListFriends(athleteId).Page(page).PerPage(perPage).Do()
	
	// returns a slice of AthleteSummary objects
	followers, err := service.ListFollowers(athleteId).Page(page).PerPage(perPage).Do()

	// returns a slice of AthleteSummary objects
	bothFollowing, err := service.ListBothFollowing(athleteId).Do()

	// returns an AthleteStats objects
	stats, err := service.Stats(athleteId).Do()

	// returns a slice of SegmentEffortSummary objects
	efforts, err := service.ListKOMs(athleteId).Do()

	// returns a slice of ActivitySummary objects
	// athleteId must match authenticated athlete's
	activities, err := service.ListActivities(athleteId).Do()

	// returns a slice of PersonalSegmentSummary objects
	segments, err := service.ListStarredSegments(athleteId).Do()


### <a name="Activities"></a>Activities

Related objects: 
[ActivityDetailed](https://godoc.org/github.com/strava/go.strava#ActivityDetailed),
[ActivitySummary](https://godoc.org/github.com/strava/go.strava#ActivitySummary),
[PhotoSummary](https://godoc.org/github.com/strava/go.strava#PhotoSummary),
[ZonesSummary](https://godoc.org/github.com/strava/go.strava#ZonesSummary),
[LapEffortSummary](https://godoc.org/github.com/strava/go.strava#LapEffortSummary),
[Location](https://godoc.org/github.com/strava/go.strava#Location).
<br />
Related constants:
[ActivityTypes](https://godoc.org/github.com/strava/go.strava#ActivityTypes).

	service := strava.NewActivitiesService(client)

	// returns a AthleteDetailed if the activity is owned by the requesting user
	// or an ActivitySummary object otherwise.
	// The Type is defined by Activity.ResourceState, 3 for detailed, 2 for summary.
	activity, err := service.Get(activityId).
		IncludeAllEfforts().
		Do()

	// create a manual activity entry. To upload a file see Upload below.
	activity, err := service.Create(name, type, startDateLocal, elapsedTime).
		Description(description).
		Distance(distance).
		Do()

	activity, err := service.Update(activityId).
		Name(name).
		Description(description).
		Type(ActivityTypes.Ride).
		Private(true).
		Communte(true).
		Trainer(false).
		Gear(gearId).
		Do()

	// returns a slice of PhotoSummary objects
	photos, err := service.ListPhotos(activityId).Do()

	// returns a slice of ZonesSummary objects
	zones, err := service.ListZones(activityId).Do()

	// returns a slice of LapEffortSummary objects
	laps, err := service.ListLaps(activityId).Do()

### <a name="Comments"></a>Comments

Related objects: 
[CommentDetailed](https://godoc.org/github.com/strava/go.strava#CommentDetailed),
[CommentSummary](https://godoc.org/github.com/strava/go.strava#CommentSummary).

	service := strava.NewActivityCommentsService(client, activityId)

	// returns a slice of CommentSummary objects
	comments, err := service.List().
		Page(page).
		PerPage(perPage).
		IncludeMarkdown().
		Do()

	// create a comment if your application has permission
	comment, err := service.Create("funny comment").Do()

	// delete a comment if your application has permission
	comment, err := service.Delete(commentId).Do()

### <a name="Kudos"></a>Kudos

Related objects: 
[AthleteSummary](https://godoc.org/github.com/strava/go.strava#AthleteSummary).

	service := strava.NewActivityKudosService(client, activityId)

	// returns a slice of AthleteSummary objects
	comments, err := service.List().
		Page(page).
		PerPage(perPage).
		Do()

	// create/give a kudo
	comment, err := service.Create().Do()

	// delete/take back a kudo
	comment, err := service.Delete().Do()

### <a name="Clubs"></a>Clubs

Related objects: 
[ClubDetailed](https://godoc.org/github.com/strava/go.strava#ClubDetailed),
[ClubSummary](https://godoc.org/github.com/strava/go.strava#ClubSummary).

	service := strava.NewClubService(client)

	// returns a ClubDetailed object
	club, err := service.Get(clubId).Do()

	// returns a slice of AthleteSummary objects
	members, err := service.ListMembers(clubId).
		Page(page).
		PerPage(perPage).
		Do()

	// returns a slice of ActivitySummary objects
	activities, err := service.ListActivities(clubId).
		Page(page).
		PerPage(perPage).
		Do()

### <a name="Gear"></a>Gear

Related objects:
[GearDetailed](https://godoc.org/github.com/strava/go.strava#GearDetailed),
[GearSummary](https://godoc.org/github.com/strava/go.strava#GearSummary).
<br />
Related constants:
[FrameTypes](https://godoc.org/github.com/strava/go.strava#FrameTypes).

	// returns a GearDetailed object
	gear, err := strava.NewGearService(client).Get(gearId).Do()

### <a name="Segments"></a>Segments

Related objects:
[SegmentDetailed](https://godoc.org/github.com/strava/go.strava#SegmentDetailed),
[SegmentSummary](https://godoc.org/github.com/strava/go.strava#SegmentSummary),
[SegmentLeaderboard](https://godoc.org/github.com/strava/go.strava#SegmentLeaderboard),
[SegmentLeaderboardEntry](https://godoc.org/github.com/strava/go.strava#SegmentLeaderboardEntry),
[SegmentExplorer](https://godoc.org/github.com/strava/go.strava#SegmentExplorer),
[SegmentExplorerSegment](https://godoc.org/github.com/strava/go.strava#SegmentExplorerSegment).
<br />
Related constants:
[AgeGroups](https://godoc.org/github.com/strava/go.strava#AgeGroups),
[ActivityTypes](https://godoc.org/github.com/strava/go.strava#ActivityTypes),
[ClimbCategories](https://godoc.org/github.com/strava/go.strava#ClimbCategories),
[DateRanges](https://godoc.org/github.com/strava/go.strava#DateRanges),
[WeightClasses](https://godoc.org/github.com/strava/go.strava#WeightClasses).

	service := strava.NewSegmentsService(client)
	
	// returns a SegmentDetailed object
	segment, err := service.Get(segmentId).Do()

	// return list of segment efforts
	efforts, err := service.ListEfforts(segmentId).
		Page(page).
		PerPage(perPage).
		AthleteId(athleteId).
		DateRange(startDateLocal, endDateLocal).
		Do()

	// returns a SegmentLeaderboard object
	leaderboard, err := service.GetLeaderboard(segmentId).
		Page(page).
		PerPage(perPage).
		Gender(gender).
		AgeGroup(ageGroup).
		WeightClass(weightClass).
		Following().
		ClubId(clubId).
		DateRange(dateRange).
		ContextEntries(count).
		Do()

	// returns a slice of SegmentExplorerSegment 
	segments, err := service.Explore(south, west, north, east).
		ActivityType(activityType).
		MinimumCategory(cat).
		MaximumCategory(cat).
		Do()

### <a name="SegmentEfforts"></a>Segment Efforts

Related objects:
[SegmentEffortDetailed](https://godoc.org/github.com/strava/go.strava#SegmentEffortDetailed),
[SegmentEffortSummary](https://godoc.org/github.com/strava/go.strava#SegmentEffortSummary).

	// returns a SegmentEffortDetailed object
	segmentEffort, err := strava.NewSegmentEffortsService(client).Get(effortId).Do() 


### <a name="Streams"></a>Streams

Related objects:
[StreamSet](https://godoc.org/github.com/strava/go.strava#StreamSet),
[LocationStream](https://godoc.org/github.com/strava/go.strava#StreamSet),
[IntegerStream](https://godoc.org/github.com/strava/go.strava#StreamSet),
[DecimalStream](https://godoc.org/github.com/strava/go.strava#StreamSet),
[BooleanStream](https://godoc.org/github.com/strava/go.strava#StreamSet),
[Stream](https://godoc.org/github.com/strava/go.strava#Stream).
<br />
Related constants:
[StreamTypes](https://godoc.org/github.com/strava/go.strava#StreamTypes).

	// Activity Streams
	// returns a StreamSet object
	strava.NewActivityStreamsService(client).
		Get(activityId, types).
		Resolution(resolution).
		SeriesType(seriesType).
		Do()

	// Segment Streams
	// returns a StreamSet object
	strava.NewSegmentStreamsService(client).
		Get(segmentId, types).
		Resolution(resolution).
		SeriesType(seriesType).
		Do()

	// SegmentEffort Streams
	// returns a StreamSet object
	strava.NewSegmentEffortStreamsService(client).
		Get(effortId, types).
		Resolution(resolution).
		SeriesType(seriesType).
		Do()


### <a name="Uploads"></a>Uploads

Related objects:
[UploadDetailed](https://godoc.org/github.com/strava/go.strava#UploadDetailed),
[UploadSummary](https://godoc.org/github.com/strava/go.strava#UploadSummary).
<br />
Related constants:
[FileDataTypes](https://godoc.org/github.com/strava/go.strava#FileDataTypes),

	service := strava.NewUploadsService(client)

	// returns a UploadDetailed object
	upload, err := service.Get(uploadId).
		Do()

	// returns a UploadSummary object
	// gzips the payload if not already done so,
	// use .Get(uploadId) to check progress on the upload and eventually get the activity id
	upload, err := service.Create(FileDataType, "filename", io.Reader).
		ActivityType(ActivityTypes.Ride).
		Name("name").
		Description("description").
		Private().
		Trainer().
		ExternalId("id").
		Do()

	// typical ways to create an io.Reader in Go
	fileReader, err := os.Open("file.go")
	byteReader, err := bytes.NewReader(binarydata)
	stringReader := strings.NewReader("stringdata")

<a name="testing"></a>Testing
-----------------------------
To test code using this package try the `StubResponseClient`.
This will stub the JSON returned by the Strava API. You'll need to 
be familiar with raw JSON responses, so see the [Documentation](http://strava.github.io/api)

	client := strava.NewStubResponseClient(`[{"id": 1,"name": "Team Strava Cycling"}`, http.StatusOK)
	clubs, err := strava.NewClubsService(client).Get(1000)

	clubs[0].Id == 1

This will return the club provided when creating the client even though you actually wanted club 1000.
