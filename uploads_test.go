package strava

import (
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"reflect"
	"strings"
	"time"

	"testing"
)

func TestUploadsGet(t *testing.T) {
	client := newCassetteClient(testToken, "upload_get")
	upload, err := NewUploadsService(client).Get(46440854).Do()

	expected := &UploadDetailed{}
	expected.Id = 46440854
	expected.ExternalId = "25FA60D8-15CF-472E-8C86-228B16320F41"
	expected.Error = ""
	expected.Status = "The created activity has been deleted."
	expected.ActivityId = 0

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if !reflect.DeepEqual(upload, expected) {
		t.Errorf("should match\n%v\n%v", upload, expected)
	}

	// from here on out just check the request parameters
	s := NewUploadsService(newStoreRequestClient())

	// path
	s.Get(123).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/uploads/123" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestUploadsCreate(t *testing.T) {
	data := strings.NewReader(rawGPXDataForTesting())

	client := newCassetteClient("special token with write permissions here", "upload_create")
	upload, err := NewUploadsService(client).Create(FileDataTypes.GPX, "", data).
		Private().
		Do()

	expected := &UploadSummary{}
	expected.Id = 141032026
	expected.ExternalId = "golibraryupload.gpx"
	expected.Error = ""
	expected.Status = "Your activity is still being processed."
	expected.ActivityId = 0

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if !reflect.DeepEqual(upload, expected) {
		t.Errorf("should match\n%v\n%v", upload, expected)
	}

	// upload already gzipped data, first gzip our test data
	gzDataBuffer := &bytes.Buffer{}
	gzWriter := gzip.NewWriter(gzDataBuffer)

	io.Copy(gzWriter, strings.NewReader(rawGPXDataForTesting()))
	gzWriter.Close()

	client = newCassetteClient("special token with write permissions here", "upload_create_gz")
	upload, err = NewUploadsService(client).Create(FileDataTypes.GPXGZ, "upload", gzDataBuffer).
		Private().
		Do()

	expected = &UploadSummary{}
	expected.Id = 141038217
	expected.ExternalId = "upload.gpx"
	expected.Error = ""
	expected.Status = "Your activity is still being processed."
	expected.ActivityId = 0

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if !reflect.DeepEqual(upload, expected) {
		t.Errorf("should match\n%v\n%v", upload, expected)
	}

	// bad reader
	var r badReader
	s := NewUploadsService(newStoreRequestClient())
	upload, err = NewUploadsService(client).Create(FileDataTypes.GPX, "", r).
		Private().
		Do()

	if err == nil {
		t.Error("should return error for bad reader")
	}

	s = NewUploadsService(newStoreRequestClient())
	upload, err = NewUploadsService(client).Create(FileDataTypes.GPXGZ, "", r).
		Private().
		Do()

	if err == nil {
		t.Error("should return error for bad reader")
	}

	// test unauthorized if no write permissions
	data2 := strings.NewReader(rawGPXDataForTesting())

	client = newCassetteClient(testToken, "upload_create_unauthorized")
	upload, err = NewUploadsService(client).Create(FileDataTypes.GPX, "", data2).
		Private().
		Do()

	if upload != nil {
		t.Error("should return nil upload on error")
	}

	if err == nil {
		t.Error("should return error when using unauthorized token")
	}

	e, ok := err.(Error)
	if !ok {
		t.Fatal("should return strava error type")
	}

	if e.Message != "Authorization Error" {
		t.Error("should return authorization error")
	}

	// path
	s = NewUploadsService(newStoreRequestClient())
	s.Create(FileDataTypes.GPX, "", strings.NewReader(rawGPXDataForTesting())).Do()
	transport := s.client.httpClient.Transport.(*storeRequestTransport)

	if transport.request.URL.Path != "/api/v3/uploads" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	// from here on out just check the request parameters
	// data type
	s = NewUploadsService(newStoreRequestClient())
	s.Create(FileDataTypes.GPXGZ, "", strings.NewReader(rawGPXDataForTesting())).
		Do()

	body := s.client.httpClient.Transport.(*storeRequestTransport).request.Body
	content, _ := ioutil.ReadAll(body)

	if !strings.Contains(string(content), "\"data_type\"\r\n\r\ngpx.gz") {
		t.Errorf("should include data type in request")
	}

	// activity type
	s = NewUploadsService(newStoreRequestClient())
	s.Create(FileDataTypes.GPX, "", strings.NewReader(rawGPXDataForTesting())).
		ActivityType(ActivityTypes.AlpineSki).
		Do()

	body = s.client.httpClient.Transport.(*storeRequestTransport).request.Body
	content, _ = ioutil.ReadAll(body)

	if !strings.Contains(string(content), "\"activity_type\"\r\n\r\nAlpineSki") {
		t.Errorf("should include activity type in request")
	}

	// name
	s = NewUploadsService(newStoreRequestClient())
	s.Create(FileDataTypes.GPX, "", strings.NewReader(rawGPXDataForTesting())).
		Name("foo").
		Do()

	body = s.client.httpClient.Transport.(*storeRequestTransport).request.Body
	content, _ = ioutil.ReadAll(body)

	if !strings.Contains(string(content), "\"name\"\r\n\r\nfoo") {
		t.Errorf("should include name in request")
	}

	// description
	s = NewUploadsService(newStoreRequestClient())
	s.Create(FileDataTypes.GPX, "", strings.NewReader(rawGPXDataForTesting())).
		Description("foo").
		Do()

	body = s.client.httpClient.Transport.(*storeRequestTransport).request.Body
	content, _ = ioutil.ReadAll(body)

	if !strings.Contains(string(content), "\"description\"\r\n\r\nfoo") {
		t.Errorf("should include description value in request")
	}

	// private
	s = NewUploadsService(newStoreRequestClient())
	s.Create(FileDataTypes.GPX, "", strings.NewReader(rawGPXDataForTesting())).
		Private().
		Do()

	body = s.client.httpClient.Transport.(*storeRequestTransport).request.Body
	content, _ = ioutil.ReadAll(body)

	if !strings.Contains(string(content), "\"private\"\r\n\r\n1") {
		t.Errorf("should include private in request")
	}

	// trainer
	s = NewUploadsService(newStoreRequestClient())
	s.Create(FileDataTypes.GPX, "", strings.NewReader(rawGPXDataForTesting())).
		Trainer().
		Do()

	body = s.client.httpClient.Transport.(*storeRequestTransport).request.Body
	content, _ = ioutil.ReadAll(body)

	if !strings.Contains(string(content), "\"trainer\"\r\n\r\n1") {
		t.Errorf("should include trainer in request")
	}

	// external id
	s = NewUploadsService(newStoreRequestClient())
	s.Create(FileDataTypes.GPX, "", strings.NewReader(rawGPXDataForTesting())).
		ExternalId("foo").
		Do()

	body = s.client.httpClient.Transport.(*storeRequestTransport).request.Body
	content, _ = ioutil.ReadAll(body)

	if !strings.Contains(string(content), "\"external_id\"\r\n\r\nfoo") {
		t.Errorf("should include external id in request")
	}
}

func TestUploadsBadJSON(t *testing.T) {
	var err error
	s := NewUploadsService(NewStubResponseClient("bad json"))

	_, err = s.Get(123).Do()
	if err == nil {
		t.Error("should return a bad json error")
	}

	_, err = s.Create(FileDataTypes.FIT, "", strings.NewReader("data")).Do()
	if err == nil {
		t.Error("should return a bad json error")
	}
}

func TestFileDataTypeIsGzipped(t *testing.T) {
	if FileDataTypes.FIT.isGzipped() {
		t.Error("should not be gzipped type")
	}

	if FileDataTypes.TCX.isGzipped() {
		t.Error("should not be gzipped type")
	}

	if FileDataTypes.GPX.isGzipped() {
		t.Error("should not be gzipped type")
	}

	if !FileDataTypes.FITGZ.isGzipped() {
		t.Error("should be gzipped type")
	}

	if !FileDataTypes.TCXGZ.isGzipped() {
		t.Error("should be gzipped type")
	}

	if !FileDataTypes.GPXGZ.isGzipped() {
		t.Error("should be gzipped type")
	}
}

func TestFileDataTypeToGzippedType(t *testing.T) {
	if FileDataTypes.FIT.toGzippedType() != FileDataTypes.FITGZ {
		t.Error("should convert to proper gzipped type")
	}

	if FileDataTypes.TCX.toGzippedType() != FileDataTypes.TCXGZ {
		t.Error("should convert to proper gzipped type")
	}

	if FileDataTypes.GPX.toGzippedType() != FileDataTypes.GPXGZ {
		t.Error("should convert to proper gzipped type")
	}

	if FileDataTypes.FITGZ.toGzippedType() != FileDataTypes.FITGZ {
		t.Error("should return self if already gzipped type")
	}

	if FileDataTypes.TCXGZ.toGzippedType() != FileDataTypes.TCXGZ {
		t.Error("should return self if already gzipped type")
	}

	if FileDataTypes.GPXGZ.toGzippedType() != FileDataTypes.GPXGZ {
		t.Error("should return self if already gzipped type")
	}

	rand := FileDataType("random")
	if rand.toGzippedType() != rand {
		t.Error("should return self if random file data type")
	}
}

func rawGPXDataForTesting() string {
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

type badReader int

func (badReader) Read(b []byte) (int, error) {
	return 0, errors.New("bad reader")
}
