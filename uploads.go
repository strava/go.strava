package strava

import (
	"encoding/json"
	"fmt"
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

func NewUploadsService(client *Client) *UploadsService {
	return &UploadsService{client}
}

/*********************************************************/

type UploadsGetCall struct {
	service *UploadsService
	id      int
}

func (s *UploadsService) Get(uploadId int) *UploadsGetCall {
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

	upload.postProcessDetailed()

	return &upload, nil
}

func (u *UploadDetailed) postProcessDetailed() {
	u.postProcessSummary()
}

func (u *UploadSummary) postProcessSummary() {

}
