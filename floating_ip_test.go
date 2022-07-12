package leaseweb

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListRanges(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 2}, "ranges": [
			{
				"id": "85.17.0.0_17",
				"range": "85.17.0.0/17",
				"customerId": "10001234",
				"salesOrgId": "2000",
				"location": "AMS-01",
				"type": "SITE"
			},
			{
				"id": "86.17.0.1_17",
				"range": "86.17.0.1/17",
				"customerId": "10001234",
				"salesOrgId": "2000",
				"location": "AMS",
				"type": "METRO"
			}
		]}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	response, err := floatingIpApi.ListRanges()

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Ranges), 2)

	range1 := response.Ranges[0]
	assert.Equal(range1.Id, "85.17.0.0_17")
	assert.Equal(range1.Range, "85.17.0.0/17")
	assert.Equal(range1.CustomerId, "10001234")
	assert.Equal(range1.SalesOrgId, "2000")
	assert.Equal(range1.Location, "AMS-01")
	assert.Equal(range1.Type, "SITE")

	range2 := response.Ranges[1]
	assert.Equal(range2.Id, "86.17.0.1_17")
	assert.Equal(range2.Range, "86.17.0.1/17")
	assert.Equal(range2.CustomerId, "10001234")
	assert.Equal(range2.SalesOrgId, "2000")
	assert.Equal(range2.Location, "AMS")
	assert.Equal(range2.Type, "METRO")
}

func TestListRangesPaginateAndFilter(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 1, "totalCount": 11}, "ranges": [
			{
				"id": "85.17.0.0_17",
				"range": "85.17.0.0/17",
				"customerId": "10001234",
				"salesOrgId": "2000",
				"location": "AMS-01",
				"type": "SITE"
			}
		]}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	response, err := floatingIpApi.ListRanges(1, 10, []string{"SITE", "METRO"}, "AMS-01")

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Ranges), 1)

	range1 := response.Ranges[0]
	assert.Equal(range1.Id, "85.17.0.0_17")
	assert.Equal(range1.Range, "85.17.0.0/17")
	assert.Equal(range1.CustomerId, "10001234")
	assert.Equal(range1.SalesOrgId, "2000")
	assert.Equal(range1.Location, "AMS-01")
	assert.Equal(range1.Type, "SITE")
}

func TestListRangesError403(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	resp, err := floatingIpApi.ListRanges()

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The access token is expired or invalid.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "ACCESS_DENIED")
}

func TestListRangesError500(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	resp, err := floatingIpApi.ListRanges()

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server encountered an unexpected condition that prevented it from fulfilling the request.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "500")
}

func TestListRangesError503(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	resp, err := floatingIpApi.ListRanges()

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "503")
}

func TestGetRange(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w, `{
			"id": "85.17.0.0_17",
			"range": "88.17.0.0/17",
			"customerId": "10001234",
			"salesOrgId": "2000",
			"location": "AMS-01",
			"type": "SITE"
		}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	response, err := floatingIpApi.GetRange("123456789")

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Id, "85.17.0.0_17")
	assert.Equal(response.Range, "88.17.0.0/17")
	assert.Equal(response.CustomerId, "10001234")
	assert.Equal(response.SalesOrgId, "2000")
	assert.Equal(response.Location, "AMS-01")
	assert.Equal(response.Type, "SITE")
}

func TestGetRangeError403(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	resp, err := floatingIpApi.GetRange("123456789")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The access token is expired or invalid.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "ACCESS_DENIED")
}

func TestGetRangeError404(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"correlationId": "39e010ed-0e93-42c3-c28f-3ffc373553d5", "errorCode": "404", "errorMessage": "Range with id 88.17.0.0_17 does not exist"}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	resp, err := floatingIpApi.GetRange("wrong range id")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Range with id 88.17.0.0_17 does not exist")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "39e010ed-0e93-42c3-c28f-3ffc373553d5")
	assert.Equal(lswErr.ErrorCode, "404")
}

func TestGetRangeError500(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	resp, err := floatingIpApi.GetRange("123456789")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server encountered an unexpected condition that prevented it from fulfilling the request.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "500")
}

func TestGetRangeError503(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	resp, err := floatingIpApi.GetRange("123456789")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "503")
}

func TestListRangeDefinitions(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 2}, "floatingIpDefinitions": [
			{
				"id": "88.17.34.108_32",
				"rangeId": "88.17.0.0_17",
				"location": "AMS-01",
				"type": "SITE",
				"customerId": "10001234",
				"salesOrgId": "2000",
				"floatingIp": "88.17.34.108/32",
				"anchorIp": "95.10.126.1",
				"status": "ACTIVE",
				"createdAt": "2019-03-13T09:10:02+0000",
				"updatedAt": "2019-03-13T09:10:02+0000"
			},
			{
				"id": "88.17.34.109_32",
				"rangeId": "88.17.0.0_17",
				"location": "AMS-01",
				"type": "SITE",
				"customerId": "10001234",
				"salesOrgId": "2000",
				"floatingIp": "88.17.34.109/32",
				"anchorIp": "95.10.126.12",
				"status": "ACTIVE",
				"createdAt": "2019-03-13T09:10:02+0000",
				"updatedAt": "2019-03-13T09:10:02+0000"
			}
		]}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	response, err := floatingIpApi.ListRangeDefinitions("123456789")

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.FloatingIpDefinitions), 2)

	floatingIpDefinition1 := response.FloatingIpDefinitions[0]
	assert.Equal(floatingIpDefinition1.Id, "88.17.34.108_32")
	assert.Equal(floatingIpDefinition1.RangeId, "88.17.0.0_17")
	assert.Equal(floatingIpDefinition1.CustomerId, "10001234")
	assert.Equal(floatingIpDefinition1.SalesOrgId, "2000")
	assert.Equal(floatingIpDefinition1.Location, "AMS-01")
	assert.Equal(floatingIpDefinition1.Type, "SITE")
	assert.Equal(floatingIpDefinition1.FloatingIp, "88.17.34.108/32")
	assert.Equal(floatingIpDefinition1.AnchorIp, "95.10.126.1")
	assert.Equal(floatingIpDefinition1.Status, "ACTIVE")
	assert.Equal(floatingIpDefinition1.CreatedAt, "2019-03-13T09:10:02+0000")
	assert.Equal(floatingIpDefinition1.UpdatedAt, "2019-03-13T09:10:02+0000")

	floatingIpDefinition2 := response.FloatingIpDefinitions[1]
	assert.Equal(floatingIpDefinition2.Id, "88.17.34.109_32")
	assert.Equal(floatingIpDefinition2.RangeId, "88.17.0.0_17")
	assert.Equal(floatingIpDefinition2.CustomerId, "10001234")
	assert.Equal(floatingIpDefinition2.SalesOrgId, "2000")
	assert.Equal(floatingIpDefinition2.Location, "AMS-01")
	assert.Equal(floatingIpDefinition2.Type, "SITE")
	assert.Equal(floatingIpDefinition2.FloatingIp, "88.17.34.109/32")
	assert.Equal(floatingIpDefinition2.AnchorIp, "95.10.126.12")
	assert.Equal(floatingIpDefinition2.Status, "ACTIVE")
	assert.Equal(floatingIpDefinition2.CreatedAt, "2019-03-13T09:10:02+0000")
	assert.Equal(floatingIpDefinition2.UpdatedAt, "2019-03-13T09:10:02+0000")
}

func TestListRangeDefinitionsPaginateAndFilter(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 1, "totalCount": 11}, "floatingIpDefinitions": [
			{
				"id": "88.17.34.108_32",
				"rangeId": "88.17.0.0_17",
				"location": "AMS-01",
				"type": "SITE",
				"customerId": "10001234",
				"salesOrgId": "2000",
				"floatingIp": "88.17.34.108/32",
				"anchorIp": "95.10.126.1",
				"status": "ACTIVE",
				"createdAt": "2019-03-13T09:10:02+0000",
				"updatedAt": "2019-03-13T09:10:02+0000"
			}
		]}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	response, err := floatingIpApi.ListRangeDefinitions("123456789", 1, 10, []string{"SITE", "METRO"}, "AMS-01")

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)

	floatingIpDefinition1 := response.FloatingIpDefinitions[0]
	assert.Equal(floatingIpDefinition1.Id, "88.17.34.108_32")
	assert.Equal(floatingIpDefinition1.RangeId, "88.17.0.0_17")
	assert.Equal(floatingIpDefinition1.CustomerId, "10001234")
	assert.Equal(floatingIpDefinition1.SalesOrgId, "2000")
	assert.Equal(floatingIpDefinition1.Location, "AMS-01")
	assert.Equal(floatingIpDefinition1.Type, "SITE")
	assert.Equal(floatingIpDefinition1.FloatingIp, "88.17.34.108/32")
	assert.Equal(floatingIpDefinition1.AnchorIp, "95.10.126.1")
	assert.Equal(floatingIpDefinition1.Status, "ACTIVE")
	assert.Equal(floatingIpDefinition1.CreatedAt, "2019-03-13T09:10:02+0000")
	assert.Equal(floatingIpDefinition1.UpdatedAt, "2019-03-13T09:10:02+0000")
}

func TestListRangeDefinitionsError403(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	resp, err := floatingIpApi.ListRangeDefinitions("123456789")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The access token is expired or invalid.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "ACCESS_DENIED")
}

func TestListRangeDefinitionsError500(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	resp, err := floatingIpApi.ListRangeDefinitions("123456789")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server encountered an unexpected condition that prevented it from fulfilling the request.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "500")
}

func TestListRangeDefinitionsError503(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	resp, err := floatingIpApi.ListRangeDefinitions("123456789")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "503")
}

func TestCreateRangeDefinition(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w, `{
			"id": "88.17.34.108_32",
			"rangeId": "88.17.0.0_17",
			"location": "AMS-01",
			"type": "SITE",
			"customerId": "10001234",
			"salesOrgId": "2000",
			"floatingIp": "88.17.34.108/32",
			"anchorIp": "95.10.126.1",
			"status": "ACTIVE",
			"createdAt": "2019-03-13T09:10:02+0000",
			"updatedAt": "2019-03-13T09:10:02+0000"
		}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	response, err := floatingIpApi.CreateRangeDefinition("10.0.0.0_29", "88.17.0.5/32", "95.10.126.1")

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Id, "88.17.34.108_32")
	assert.Equal(response.RangeId, "88.17.0.0_17")
	assert.Equal(response.CustomerId, "10001234")
	assert.Equal(response.SalesOrgId, "2000")
	assert.Equal(response.Location, "AMS-01")
	assert.Equal(response.Type, "SITE")
	assert.Equal(response.FloatingIp, "88.17.34.108/32")
	assert.Equal(response.AnchorIp, "95.10.126.1")
	assert.Equal(response.Status, "ACTIVE")
	assert.Equal(response.CreatedAt, "2019-03-13T09:10:02+0000")
	assert.Equal(response.UpdatedAt, "2019-03-13T09:10:02+0000")
}

func TestCreateRangeDefinitionError400(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"correlationId": "945bef2e-1caf-4027-bd0a-8976848f3dee", "errorCode": "400", "errorMessage": "Validation Failed"}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	resp, err := floatingIpApi.CreateRangeDefinition("10.0.0.0_29", "wrong 1", "wrong 2")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Validation Failed")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "945bef2e-1caf-4027-bd0a-8976848f3dee")
	assert.Equal(lswErr.ErrorCode, "400")
}

func TestCreateRangeDefinitionError403(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	resp, err := floatingIpApi.CreateRangeDefinition("10.0.0.0_29", "88.17.0.5/32", "95.10.126.1")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The access token is expired or invalid.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "ACCESS_DENIED")
}

func TestCreateRangeDefinitionError500(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	resp, err := floatingIpApi.CreateRangeDefinition("10.0.0.0_29", "88.17.0.5/32", "95.10.126.1")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server encountered an unexpected condition that prevented it from fulfilling the request.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "500")
}

func TestCreateRangeDefinitionError503(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	resp, err := floatingIpApi.CreateRangeDefinition("10.0.0.0_29", "88.17.0.5/32", "95.10.126.1")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "503")
}

func TestGetRangeDefinition(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w, `{
			"id": "88.17.34.108_32",
			"rangeId": "88.17.0.0_17",
			"location": "AMS-01",
			"type": "SITE",
			"customerId": "10001234",
			"salesOrgId": "2000",
			"floatingIp": "88.17.34.108/32",
			"anchorIp": "95.10.126.1",
			"status": "ACTIVE",
			"createdAt": "2019-03-13T09:10:02+0000",
			"updatedAt": "2019-03-13T09:10:02+0000"
		}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	response, err := floatingIpApi.GetRangeDefinition("88.17.0.0_17", "88.17.34.108_32")

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Id, "88.17.34.108_32")
	assert.Equal(response.RangeId, "88.17.0.0_17")
	assert.Equal(response.CustomerId, "10001234")
	assert.Equal(response.SalesOrgId, "2000")
	assert.Equal(response.Location, "AMS-01")
	assert.Equal(response.Type, "SITE")
	assert.Equal(response.FloatingIp, "88.17.34.108/32")
	assert.Equal(response.AnchorIp, "95.10.126.1")
	assert.Equal(response.Status, "ACTIVE")
	assert.Equal(response.CreatedAt, "2019-03-13T09:10:02+0000")
	assert.Equal(response.UpdatedAt, "2019-03-13T09:10:02+0000")
}

func TestGetRangeDefinitionError403(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	resp, err := floatingIpApi.GetRangeDefinition("88.17.0.0_17", "88.17.34.108_32")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The access token is expired or invalid.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "ACCESS_DENIED")
}

func TestGetRangeDefinitionError500(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	resp, err := floatingIpApi.GetRangeDefinition("88.17.0.0_17", "88.17.34.108_32")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server encountered an unexpected condition that prevented it from fulfilling the request.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "500")
}

func TestGetRangeDefinitionError503(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	resp, err := floatingIpApi.GetRangeDefinition("88.17.0.0_17", "88.17.34.108_32")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "503")
}

func TestUpdateRangeDefinition(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w, `{
			"id": "88.17.34.108_32",
			"rangeId": "88.17.0.0_17",
			"location": "AMS-01",
			"type": "SITE",
			"customerId": "10001234",
			"salesOrgId": "2000",
			"floatingIp": "88.17.34.108/32",
			"anchorIp": "95.10.126.1",
			"status": "ACTIVE",
			"createdAt": "2019-03-13T09:10:02+0000",
			"updatedAt": "2019-03-13T09:10:02+0000"
		}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	response, err := floatingIpApi.UpdateRangeDefinition("88.17.0.0_17", "88.17.34.108_32", "95.10.126.1")

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Id, "88.17.34.108_32")
	assert.Equal(response.RangeId, "88.17.0.0_17")
	assert.Equal(response.CustomerId, "10001234")
	assert.Equal(response.SalesOrgId, "2000")
	assert.Equal(response.Location, "AMS-01")
	assert.Equal(response.Type, "SITE")
	assert.Equal(response.FloatingIp, "88.17.34.108/32")
	assert.Equal(response.AnchorIp, "95.10.126.1")
	assert.Equal(response.Status, "ACTIVE")
	assert.Equal(response.CreatedAt, "2019-03-13T09:10:02+0000")
	assert.Equal(response.UpdatedAt, "2019-03-13T09:10:02+0000")
}

func TestUpdateRangeDefinitionError400(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"correlationId": "945bef2e-1caf-4027-bd0a-8976848f3dee", "errorCode": "400", "errorMessage": "Validation Failed"}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	resp, err := floatingIpApi.UpdateRangeDefinition("wrong 1", "88.17.34.108_32", "95.10.126.1")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Validation Failed")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "945bef2e-1caf-4027-bd0a-8976848f3dee")
	assert.Equal(lswErr.ErrorCode, "400")
}

func TestUpdateRangeDefinitionError403(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	resp, err := floatingIpApi.UpdateRangeDefinition("88.17.0.0_17", "88.17.34.108_32", "95.10.126.1")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The access token is expired or invalid.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "ACCESS_DENIED")
}

func TestUpdateRangeDefinitionError500(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	resp, err := floatingIpApi.UpdateRangeDefinition("88.17.0.0_17", "88.17.34.108_32", "95.10.126.1")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server encountered an unexpected condition that prevented it from fulfilling the request.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "500")
}

func TestUpdateRangeDefinitionError503(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	resp, err := floatingIpApi.UpdateRangeDefinition("88.17.0.0_17", "88.17.34.108_32", "95.10.126.1")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "503")
}

func TestRemoveRangeDefinition(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w, `{
			"id": "88.17.34.108_32",
			"rangeId": "88.17.0.0_17",
			"location": "AMS-01",
			"type": "SITE",
			"customerId": "10001234",
			"salesOrgId": "2000",
			"floatingIp": "88.17.34.108/32",
			"anchorIp": "95.10.126.1",
			"status": "ACTIVE",
			"createdAt": "2019-03-13T09:10:02+0000",
			"updatedAt": "2019-03-13T09:10:02+0000"
		}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	response, err := floatingIpApi.RemoveRangeDefinition("88.17.0.0_17", "88.17.34.108_32")

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Id, "88.17.34.108_32")
	assert.Equal(response.RangeId, "88.17.0.0_17")
	assert.Equal(response.CustomerId, "10001234")
	assert.Equal(response.SalesOrgId, "2000")
	assert.Equal(response.Location, "AMS-01")
	assert.Equal(response.Type, "SITE")
	assert.Equal(response.FloatingIp, "88.17.34.108/32")
	assert.Equal(response.AnchorIp, "95.10.126.1")
	assert.Equal(response.Status, "ACTIVE")
	assert.Equal(response.CreatedAt, "2019-03-13T09:10:02+0000")
	assert.Equal(response.UpdatedAt, "2019-03-13T09:10:02+0000")
}

func TestRemoveRangeDefinitionError403(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	resp, err := floatingIpApi.RemoveRangeDefinition("88.17.0.0_17", "88.17.34.108_32")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The access token is expired or invalid.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "ACCESS_DENIED")
}

func TestRemoveRangeDefinitionError404(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"correlationId": "39e010ed-0e93-42c3-c28f-3ffc373553d5", "errorCode": "404", "errorMessage": "Floating Ip definition with id 88.17.34.108_32 does not exist."}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	resp, err := floatingIpApi.RemoveRangeDefinition("88.17.0.0_17", "88.17.34.108_32")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Floating Ip definition with id 88.17.34.108_32 does not exist.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "39e010ed-0e93-42c3-c28f-3ffc373553d5")
	assert.Equal(lswErr.ErrorCode, "404")
}

func TestRemoveRangeDefinitionError500(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "500", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	resp, err := floatingIpApi.RemoveRangeDefinition("88.17.0.0_17", "88.17.34.108_32")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server encountered an unexpected condition that prevented it from fulfilling the request.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "500")
}

func TestRemoveRangeDefinitionError503(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"correlationId": "289346a1-3eaf-4da4-b707-62ef12eb08be", "errorCode": "503", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
	})
	defer teardown()

	floatingIpApi := FloatingIpApi{}
	resp, err := floatingIpApi.RemoveRangeDefinition("88.17.0.0_17", "88.17.34.108_32")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.CorrelationId, "289346a1-3eaf-4da4-b707-62ef12eb08be")
	assert.Equal(lswErr.ErrorCode, "503")
}
