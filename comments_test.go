package strava

import (
	"testing"
)

func TestActivityCommentsList(t *testing.T) {
	client := newCassetteClient(testToken, "activity_comments_list")
	comments, err := NewActivityCommentsService(client, 103221154).List().Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if len(comments) == 0 {
		t.Fatal("comments not parsed")
	}

	if v := comments[0].Id; v != 19035182 {
		t.Errorf("value incorrect, got %v", v)
	}

	if v := comments[0].ActivityId; v != 103221154 {
		t.Errorf("value incorrect, got %v", v)
	}

	if v := comments[0].Text; v != "Testing!!!" {
		t.Errorf("value incorrect, got %v", v)
	}

	if comments[0].CreatedAt.IsZero() {
		t.Error("dates not parsed")
	}

	if comments[0].Athlete.CreatedAt.IsZero() || comments[0].Athlete.UpdatedAt.IsZero() {
		t.Error("athlete dates not parsed")
	}

	// from here on out just check the request parameters
	s := NewActivityCommentsService(newStoreRequestClient(), 123)

	// path
	s.List().Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/activities/123/comments" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}

	// parameters
	s.List().IncludeMarkdown().Do()

	transport = s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/activities/123/comments" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "markdown=true" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}

	// parameters2
	s.List().Page(1).PerPage(10).Do()

	transport = s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/activities/123/comments" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "page=1&per_page=10" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestActivityCommentsCreate(t *testing.T) {
	client := newCassetteClient(testToken, "activity_comments_post")
	comment, err := NewActivityCommentsService(client, 118293263).Create("test comment").Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	if v := comment.Id; v != 30043520 {
		t.Errorf("value incorrect, got %v", v)
	}

	if v := comment.ActivityId; v != 118293263 {
		t.Errorf("value incorrect, got %v", v)
	}

	if v := comment.Text; v != "test comment" {
		t.Errorf("value incorrect, got %v", v)
	}

	if comment.CreatedAt.IsZero() {
		t.Error("dates not parsed")
	}

	if comment.Athlete.CreatedAt.IsZero() || comment.Athlete.UpdatedAt.IsZero() {
		t.Error("athlete dates not parsed")
	}

	// from here on out just check the request parameters
	s := NewActivityCommentsService(newStoreRequestClient(), 123)

	// path
	s.Create("test comment").Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/activities/123/comments" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestActivityCommentsDelete(t *testing.T) {
	client := newCassetteClient(testToken, "activity_comments_delete")
	err := NewActivityCommentsService(client, 118293263).Delete(30043520).Do()

	if err != nil {
		t.Fatalf("service error: %v", err)
	}

	// from here on out just check the request parameters
	s := NewActivityCommentsService(newStoreRequestClient(), 123)

	// path
	s.Delete(321).Do()

	transport := s.client.httpClient.Transport.(*storeRequestTransport)
	if transport.request.URL.Path != "/api/v3/activities/123/comments/321" {
		t.Errorf("request path incorrect, got %v", transport.request.URL.Path)
	}

	if transport.request.URL.RawQuery != "" {
		t.Errorf("request query incorrect, got %v", transport.request.URL.RawQuery)
	}
}

func TestActivityCommentsBadJSON(t *testing.T) {
	var err error
	s := NewActivityCommentsService(NewStubResponseClient("bad json"), 123)

	_, err = s.List().Do()
	if err == nil {
		t.Error("should return a bad json error")
	}

	_, err = s.Create("abc").Do()
	if err == nil {
		t.Error("should return a bad json error")
	}
}
