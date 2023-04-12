package leaseweb

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbuseList(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 2}, "reports": [
			{
				"id": "000000",
				"subject": "Report description 1",
				"status": "OPEN",
				"reportedAt": "2022-01-01T00:00:00+00:00",
				"updatedAt": "2022-02-01T00:00:00+00:00",
				"notifier": "notifier1@email.com",
				"customerId": "10000001",
				"legalEntityId": "2000",
				"deadline": "2022-03-01T00:00:00+00:00"
			},
			{
				"id": "000001",
				"subject": "Report description 2",
				"status": "CLOSED",
				"reportedAt": "2022-03-01T00:00:00+00:00",
				"updatedAt": "2022-04-01T00:00:00+00:00",
				"notifier": "notifier2@email.com",
				"customerId": "10000001",
				"legalEntityId": "2600",
				"deadline": "2022-05-01T00:00:00+00:00"
			}
		]}`)
	})
	defer teardown()

	abuseApi := AbuseApi{}
	ctx := context.Background()
	response, err := abuseApi.List(ctx, AbuseListOptions{})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Reports), 2)

	abuseReport1 := response.Reports[0]
	assert.Equal(abuseReport1.Id, "000000")
	assert.Equal(abuseReport1.Subject, "Report description 1")
	assert.Equal(abuseReport1.Status, "OPEN")
	assert.Equal(abuseReport1.ReportedAt, "2022-01-01T00:00:00+00:00")
	assert.Equal(abuseReport1.UpdatedAt, "2022-02-01T00:00:00+00:00")
	assert.Equal(abuseReport1.Notifier, "notifier1@email.com")
	assert.Equal(abuseReport1.CustomerId, "10000001")
	assert.Equal(abuseReport1.LegalEntityId, "2000")
	assert.Equal(abuseReport1.Deadline, "2022-03-01T00:00:00+00:00")

	abuseReport2 := response.Reports[1]
	assert.Equal(abuseReport2.Id, "000001")
	assert.Equal(abuseReport2.Subject, "Report description 2")
	assert.Equal(abuseReport2.Status, "CLOSED")
	assert.Equal(abuseReport2.ReportedAt, "2022-03-01T00:00:00+00:00")
	assert.Equal(abuseReport2.UpdatedAt, "2022-04-01T00:00:00+00:00")
	assert.Equal(abuseReport2.Notifier, "notifier2@email.com")
	assert.Equal(abuseReport2.CustomerId, "10000001")
	assert.Equal(abuseReport2.LegalEntityId, "2600")
	assert.Equal(abuseReport2.Deadline, "2022-05-01T00:00:00+00:00")
}

func TestAbuseListPaginateAndPassStatuses(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 1, "totalCount": 11}, "reports": [
			{
				"id": "000000",
				"subject": "Report description 1",
				"status": "OPEN",
				"reportedAt": "2022-01-01T00:00:00+00:00",
				"updatedAt": "2022-02-01T00:00:00+00:00",
				"notifier": "notifier1@email.com",
				"customerId": "10000001",
				"legalEntityId": "2000",
				"deadline": "2022-03-01T00:00:00+00:00"
			}
		]}`)
	})
	defer teardown()

	abuseApi := AbuseApi{}

	ctx := context.Background()
	opts := AbuseListOptions{
		PaginationOptions: PaginationOptions{
			Limit: Int(1),
		},
		Status: []string{"OPEN", "WAITING"},
	}
	response, err := abuseApi.List(ctx, opts)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Reports), 1)

	abuseReport1 := response.Reports[0]
	assert.Equal(abuseReport1.Id, "000000")
	assert.Equal(abuseReport1.Subject, "Report description 1")
	assert.Equal(abuseReport1.Status, "OPEN")
	assert.Equal(abuseReport1.ReportedAt, "2022-01-01T00:00:00+00:00")
	assert.Equal(abuseReport1.UpdatedAt, "2022-02-01T00:00:00+00:00")
	assert.Equal(abuseReport1.Notifier, "notifier1@email.com")
	assert.Equal(abuseReport1.CustomerId, "10000001")
	assert.Equal(abuseReport1.LegalEntityId, "2000")
	assert.Equal(abuseReport1.Deadline, "2022-03-01T00:00:00+00:00")
}

func TestAbuseListBeEmpty(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 0}, "reports": []}`)
	})
	defer teardown()

	abuseApi := AbuseApi{}
	ctx := context.Background()
	response, err := abuseApi.List(ctx, AbuseListOptions{})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 0)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Reports), 0)
}

func TestAbuseListServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return AbuseApi{}.List(ctx, AbuseListOptions{})
			},
			ExpectedError: ApiError{
				Code:    "ACCESS_DENIED",
				Message: "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return AbuseApi{}.List(ctx, AbuseListOptions{})
			},
			ExpectedError: ApiError{
				Code:    "SERVER_ERROR",
				Message: "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return AbuseApi{}.List(ctx, AbuseListOptions{})
			},
			ExpectedError: ApiError{
				Code:    "TEMPORARILY_UNAVAILABLE",
				Message: "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestAbuseGet(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{
			"id": "000005",
			"subject": "Report description",
			"status": "CLOSED",
			"reopened": false,
			"reportedAt": "2015-01-01T00:00:00+0100",
			"updatedAt": "2015-02-01T00:00:00+0100",
			"notifier": "notifier@email.com",
			"customerId": "10000001",
			"legalEntityId": "2000",
			"body": "string with content",
			"deadline": "2015-01-01T00:00:00+0100",
			"detectedIpAddresses": [
				"127.0.0.1"
			],
			"detectedDomainNames": [
				{
					"name": "example.com",
					"ipAddresses": [
						"93.184.216.34"
					]
				}
			],
			"attachments": [
				{
					"id": "1abd8e7f-0fdf-453c-b1f5-8fef436acbbe",
					"mimeType": "part/xml",
					"filename": "000001.xml"
				}
			],
			"totalMessagesCount": 2,
			"latestMessages": [
				{
					"postedBy": "CUSTOMER",
					"postedAt": "2015-09-30T06:23:40+00:00",
					"body": "Hello, this is my first message!"
				},
				{
					"postedBy": "ABUSE_AGENT",
					"postedAt": "2015-10-08T08:25:29+00:00",
					"body": "Hi, this is our first reply.",
					"attachment": {
						"id": "436acbbe-0fdf-453c-b1f5-1abd8e7f8fef",
						"mimeType": "image/png",
						"filename": "notification.png"
					}
				}
			]
		}`)
	})
	defer teardown()

	abuseApi := AbuseApi{}
	ctx := context.Background()
	response, err := abuseApi.Get(ctx, "000005")

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Id, "000005")
	assert.Equal(response.Subject, "Report description")
	assert.Equal(response.Status, "CLOSED")
	assert.Equal(response.Reopened, false)
	assert.Equal(response.ReportedAt, "2015-01-01T00:00:00+0100")
	assert.Equal(response.UpdatedAt, "2015-02-01T00:00:00+0100")
	assert.Equal(response.Notifier, "notifier@email.com")
	assert.Equal(response.CustomerId, "10000001")
	assert.Equal(response.LegalEntityId, "2000")
	assert.Equal(response.Body, "string with content")
	assert.Equal(response.Deadline, "2015-01-01T00:00:00+0100")
	assert.Equal(response.DetectedIpAddresses[0], "127.0.0.1")
	assert.Equal(response.DetectedDomainNames[0].Name, "example.com")
	assert.Equal(response.DetectedDomainNames[0].IpAddresses[0], "93.184.216.34")
	assert.Equal(response.Attachments[0].Id, "1abd8e7f-0fdf-453c-b1f5-8fef436acbbe")
	assert.Equal(response.Attachments[0].MimeType, "part/xml")
	assert.Equal(response.Attachments[0].Filename, "000001.xml")
	assert.Equal(response.TotalMessagesCount.String(), "2")
	assert.Equal(len(response.LatestMessages), 2)

	message1 := response.LatestMessages[0]
	assert.Equal(message1.PostedBy, "CUSTOMER")
	assert.Equal(message1.PostedAt, "2015-09-30T06:23:40+00:00")
	assert.Equal(message1.Body, "Hello, this is my first message!")

	message2 := response.LatestMessages[1]
	assert.Equal(message2.PostedBy, "ABUSE_AGENT")
	assert.Equal(message2.PostedAt, "2015-10-08T08:25:29+00:00")
	assert.Equal(message2.Body, "Hi, this is our first reply.")

	assert.Equal(message2.Attachment.Id, "436acbbe-0fdf-453c-b1f5-1abd8e7f8fef")
	assert.Equal(message2.Attachment.MimeType, "image/png")
	assert.Equal(message2.Attachment.Filename, "notification.png")
}

func TestAbuseGetServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return AbuseApi{}.Get(ctx, "wrong-id")
			},
			ExpectedError: ApiError{
				Code:    "ACCESS_DENIED",
				Message: "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, `{}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return AbuseApi{}.Get(ctx, "123456")
			},
			ExpectedError: ApiError{Code: "404", Message: "Not Found"},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return AbuseApi{}.Get(ctx, "123456")
			},
			ExpectedError: ApiError{
				Code:    "SERVER_ERROR",
				Message: "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return AbuseApi{}.Get(ctx, "123456")
			},
			ExpectedError: ApiError{
				Code:    "TEMPORARILY_UNAVAILABLE",
				Message: "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestAbuseListMessages(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 2}, "messages": [
			{
				"postedBy": "CUSTOMER",
				"postedAt": "2015-09-30T06:23:40+00:00",
				"body": "Hello, this is my first message!"
			},
			{
				"postedBy": "ABUSE_AGENT",
				"postedAt": "2015-10-08T08:25:29+00:00",
				"body": "Hi, this is our first reply.",
				"attachment": {
					"id": "436acbbe-0fdf-453c-b1f5-1abd8e7f8fef",
					"mimeType": "image/png",
					"filename": "notification.png"
				}
			}
		]}`)
	})
	defer teardown()

	abuseApi := AbuseApi{}
	ctx := context.Background()
	response, err := abuseApi.ListMessages(ctx, "123456789", PaginationOptions{})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Messages), 2)

	message1 := response.Messages[0]
	assert.Equal(message1.PostedBy, "CUSTOMER")
	assert.Equal(message1.PostedAt, "2015-09-30T06:23:40+00:00")
	assert.Equal(message1.Body, "Hello, this is my first message!")

	message2 := response.Messages[1]
	assert.Equal(message2.PostedBy, "ABUSE_AGENT")
	assert.Equal(message2.PostedAt, "2015-10-08T08:25:29+00:00")
	assert.Equal(message2.Body, "Hi, this is our first reply.")

	assert.Equal(message2.Attachment.Id, "436acbbe-0fdf-453c-b1f5-1abd8e7f8fef")
	assert.Equal(message2.Attachment.MimeType, "image/png")
	assert.Equal(message2.Attachment.Filename, "notification.png")
}

func TestAbuseListMessagesPaginate(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 1, "totalCount": 11}, "messages": [
			{
				"postedBy": "ABUSE_AGENT",
				"postedAt": "2015-10-08T08:25:29+00:00",
				"body": "Hi, this is our first reply.",
				"attachment": {
					"id": "436acbbe-0fdf-453c-b1f5-1abd8e7f8fef",
					"mimeType": "image/png",
					"filename": "notification.png"
				}
			}
		]}`)
	})
	defer teardown()

	abuseApi := AbuseApi{}
	ctx := context.Background()
	opts := PaginationOptions{
		Limit: Int(1),
	}
	response, err := abuseApi.ListMessages(ctx, "123456789", opts)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Messages), 1)

	message1 := response.Messages[0]
	assert.Equal(message1.PostedBy, "ABUSE_AGENT")
	assert.Equal(message1.PostedAt, "2015-10-08T08:25:29+00:00")
	assert.Equal(message1.Body, "Hi, this is our first reply.")

	assert.Equal(message1.Attachment.Id, "436acbbe-0fdf-453c-b1f5-1abd8e7f8fef")
	assert.Equal(message1.Attachment.MimeType, "image/png")
	assert.Equal(message1.Attachment.Filename, "notification.png")
}

func TestAbuseListMessagesBeEmpty(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 0}, "messages": []}`)
	})
	defer teardown()

	abuseApi := AbuseApi{}
	ctx := context.Background()
	response, err := abuseApi.ListMessages(ctx, "123456789", PaginationOptions{})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Empty(response.Messages)
	assert.Equal(response.Metadata.TotalCount, 0)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
}

func TestAbuseListMessagesServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return AbuseApi{}.ListMessages(ctx, "123456789", PaginationOptions{})
			},
			ExpectedError: ApiError{
				Code:    "ACCESS_DENIED",
				Message: "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 404",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, `{}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return AbuseApi{}.ListMessages(ctx, "123456789", PaginationOptions{})
			},
			ExpectedError: ApiError{Code: "404", Message: "Not Found"},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return AbuseApi{}.ListMessages(ctx, "123456789", PaginationOptions{})
			},
			ExpectedError: ApiError{
				Code:    "SERVER_ERROR",
				Message: "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return AbuseApi{}.ListMessages(ctx, "123456789", PaginationOptions{})
			},
			ExpectedError: ApiError{
				Code:    "TEMPORARILY_UNAVAILABLE",
				Message: "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestAbuseCreateMessage(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, `["To make sure the request has been processed please see if the message is added to the list."]`)
	})
	defer teardown()

	abuseApi := AbuseApi{}
	ctx := context.Background()
	resp, err := abuseApi.CreateMessage(ctx, "123456789", "message body...")

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(resp[0], "To make sure the request has been processed please see if the message is added to the list.")
}

func TestAbuseCreateMessageServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return AbuseApi{}.CreateMessage(ctx, "123456789", "message body...")
			},
			ExpectedError: ApiError{
				Code:    "ACCESS_DENIED",
				Message: "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return AbuseApi{}.CreateMessage(ctx, "123456789", "message body...")
			},
			ExpectedError: ApiError{
				Code:    "SERVER_ERROR",
				Message: "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return AbuseApi{}.CreateMessage(ctx, "123456789", "message body...")
			},
			ExpectedError: ApiError{
				Code:    "TEMPORARILY_UNAVAILABLE",
				Message: "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestAbuseListResolutionOptions(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"resolutions": [
			{
				"id": "CONTENT_REMOVED",
				"description": "The mentioned content has been removed."
			},
			{
				"id": "DOMAINS_REMOVED",
				"description": "The mentioned domain(s) has/have been removed from the LeaseWeb network."
			},
			{
				"id": "SUSPENDED",
				"description": "The end customer (or responsible user) has been suspended."
			},
			{
				"id": "DUPLICATE",
				"description": "This is either a duplicate or old notification and has already been resolved."
			}
		]}`)
	})
	defer teardown()

	abuseApi := AbuseApi{}
	ctx := context.Background()
	response, err := abuseApi.ListResolutionOptions(ctx, "123456789")

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Resolutions[0].Id, "CONTENT_REMOVED")
	assert.Equal(response.Resolutions[1].Id, "DOMAINS_REMOVED")
	assert.Equal(response.Resolutions[2].Id, "SUSPENDED")
	assert.Equal(response.Resolutions[3].Id, "DUPLICATE")
}

func TestAbuseListResolutionOptionsBeEmpty(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		fmt.Fprintf(w, `{"resolutions": []}`)
	})
	defer teardown()

	abuseApi := AbuseApi{}
	ctx := context.Background()
	response, err := abuseApi.ListResolutionOptions(ctx, "123456789")

	assert := assert.New(t)
	assert.Nil(err)
	assert.Empty(response.Resolutions)
}

func TestAbuseListResolutionOptionsServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return AbuseApi{}.ListResolutionOptions(ctx, "123456789")
			},
			ExpectedError: ApiError{
				Code:    "ACCESS_DENIED",
				Message: "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return AbuseApi{}.ListResolutionOptions(ctx, "123456789")
			},
			ExpectedError: ApiError{
				Code:    "SERVER_ERROR",
				Message: "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return AbuseApi{}.ListResolutionOptions(ctx, "123456789")
			},
			ExpectedError: ApiError{
				Code:    "TEMPORARILY_UNAVAILABLE",
				Message: "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}

func TestAbuseResolve(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
		w.WriteHeader(http.StatusNoContent)
	})
	defer teardown()

	abuseApi := AbuseApi{}
	ctx := context.Background()
	err := abuseApi.Resolve(ctx, "123456789", []string{"CONTENT_REMOVED", "SUSPENDED"})

	assert := assert.New(t)
	assert.Nil(err)
}

func TestAbuseResolveServerErrors(t *testing.T) {
	serverErrorTests := []serverErrorTest{
		{
			Title: "error 403",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, AbuseApi{}.Resolve(ctx, "123456789", []string{"CONTENT_REMOVED", "SUSPENDED"})
			},
			ExpectedError: ApiError{
				Code:    "ACCESS_DENIED",
				Message: "The access token is expired or invalid.",
			},
		},
		{
			Title: "error 500",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, AbuseApi{}.Resolve(ctx, "123456789", []string{"CONTENT_REMOVED", "SUSPENDED"})
			},
			ExpectedError: ApiError{
				Code:    "SERVER_ERROR",
				Message: "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			},
		},
		{
			Title: "error 503",
			MockServer: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, testApiKey, r.Header.Get("x-lsw-auth"))
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, `{"errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
			},
			FunctionCall: func() (interface{}, error) {
				ctx := context.Background()
				return nil, AbuseApi{}.Resolve(ctx, "123456789", []string{"CONTENT_REMOVED", "SUSPENDED"})
			},
			ExpectedError: ApiError{
				Code:    "TEMPORARILY_UNAVAILABLE",
				Message: "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
			},
		},
	}
	assertServerErrorTests(t, serverErrorTests)
}
