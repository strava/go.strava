package strava

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

type UploadDetailed struct {
	UploadSummary
}

type UploadSummary struct {
	Id         int64  `json:"id"`
	ExternalId string `json:"external_id"`
	Error      string `json:"error"`
	Status     string `json:"status"`
	ActivityId int64  `json:"activity_id"`
}

type UploadsService struct {
	client *Client
}

type FileDataType string

var FileDataTypes = struct {
	FIT   FileDataType
	FITGZ FileDataType
	TCX   FileDataType
	TCXGZ FileDataType
	GPX   FileDataType
	GPXGZ FileDataType
}{"fit", "fit.gz", "tcx", "tcx.gz", "gpx", "gpx.gz"}

var errorHandler ErrorHandler = func(response *http.Response) error {
	if response.StatusCode == 400 {
		contents, _ := ioutil.ReadAll(response.Body)
		var e UploadSummary
		json.Unmarshal(contents, &e)
		return Error{e.Error, nil}
	} else {
		return defaultErrorHandler(response)
	}
}

func NewUploadsService(client *Client) *UploadsService {
	return &UploadsService{client}
}

/*********************************************************/

type UploadsGetCall struct {
	service *UploadsService
	id      int64
}

func (s *UploadsService) Get(uploadId int64) *UploadsGetCall {
	return &UploadsGetCall{
		service: s,
		id:      uploadId,
	}
}

func (c *UploadsGetCall) Do() (*UploadDetailed, error) {
	data, err := c.service.client.run("GET", fmt.Sprintf("/uploads/%d", c.id), nil)
	if err != nil {
		return nil, err
	}

	var upload UploadDetailed
	err = json.Unmarshal(data, &upload)
	if err != nil {
		return nil, err
	}

	return &upload, nil
}

/*********************************************************/

type UploadsCreateCall struct {
	service    *UploadsService
	ops        map[string]interface{}
	filename   string
	fileReader io.Reader
}

// Create defines an upload call containing the contents of the reader.
// will gzip the file, if the the dataType indicates it's not already.
func (s *UploadsService) Create(dataType FileDataType, filename string, reader io.Reader) *UploadsCreateCall {
	call := &UploadsCreateCall{
		service:    s,
		ops:        make(map[string]interface{}),
		filename:   filename,
		fileReader: reader,
	}
	if call.filename == "" {
		call.filename = fmt.Sprintf("golibraryupload.%v", dataType)
	}

	call.ops["data_type"] = dataType

	return call
}

func (c *UploadsCreateCall) ActivityType(activityType ActivityType) *UploadsCreateCall {
	c.ops["activity_type"] = string(activityType)
	return c
}

func (c *UploadsCreateCall) Name(name string) *UploadsCreateCall {
	c.ops["name"] = name
	return c
}

func (c *UploadsCreateCall) Description(description string) *UploadsCreateCall {
	c.ops["description"] = description
	return c
}

func (c *UploadsCreateCall) Private() *UploadsCreateCall {
	c.ops["private"] = 1
	return c
}

func (c *UploadsCreateCall) Trainer() *UploadsCreateCall {
	c.ops["trainer"] = 1
	return c
}

func (c *UploadsCreateCall) ExternalId(id string) *UploadsCreateCall {
	c.ops["external_id"] = id
	return c
}

func (c *UploadsCreateCall) Do() (*UploadSummary, error) {
	var err error
	// since we're doing a multipart post, the request is custom built

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(c.filename))

	// gzip the file if it isn't already
	if c.ops["data_type"].(FileDataType).isGzipped() {
		_, err = io.Copy(part, c.fileReader)
		if err != nil {
			return nil, err
		}
	} else {
		// gzip here for the user

		gzBuffer := &bytes.Buffer{}
		gzWriter := gzip.NewWriter(gzBuffer)

		_, err = io.Copy(gzWriter, c.fileReader)
		gzWriter.Close()
		if err != nil {
			return nil, err
		}

		io.Copy(part, gzBuffer)

		c.ops["data_type"] = c.ops["data_type"].(FileDataType).toGzippedType()
	}

	for k, v := range c.ops {
		writer.WriteField(k, fmt.Sprintf("%v", v))
	}

	writer.Close() // so it finishes writing everything to the body buffer

	req, err := http.NewRequest("POST", basePath+"/uploads", body)
	req.Header.Add("Content-Type", "multipart/form-data; boundary="+writer.Boundary())

	data, err := c.service.client.runRequestWithErrorHandler(req, errorHandler)
	if err != nil {
		return nil, err
	}

	var upload UploadSummary
	err = json.Unmarshal(data, &upload)
	if err != nil {
		return nil, err
	}

	return &upload, nil
}

/*********************************************************/

func (f FileDataType) isGzipped() bool {
	return f == FileDataTypes.FITGZ || f == FileDataTypes.TCXGZ || f == FileDataTypes.GPXGZ
}

func (f FileDataType) toGzippedType() FileDataType {
	if f == FileDataTypes.FITGZ || f == FileDataTypes.TCXGZ || f == FileDataTypes.GPXGZ {
		return f
	}

	switch f {
	case FileDataTypes.FIT:
		return FileDataTypes.FITGZ
	case FileDataTypes.TCX:
		return FileDataTypes.TCXGZ
	case FileDataTypes.GPX:
		return FileDataTypes.GPXGZ
	}

	return f
}
