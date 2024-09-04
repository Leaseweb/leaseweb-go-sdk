/*
Leaseweb API for aggregation packs

Testing AggregationPackAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package aggregationPack

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/aggregationPack"
)

func Test_aggregationPack_AggregationPackAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AggregationPackAPIService GetAggregationPack", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var aggregationPackId string

		resp, httpRes, err := apiClient.AggregationPackAPI.GetAggregationPack(context.Background(), aggregationPackId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AggregationPackAPIService GetAggregationPackList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AggregationPackAPI.GetAggregationPackList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
