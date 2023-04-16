/*
Corellium API

Testing CoreTraceApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package corellium

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_corellium_CoreTraceApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CoreTraceApiService V1ClearCoreTrace", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var instanceId string

		httpRes, err := apiClient.CoreTraceApi.V1ClearCoreTrace(context.Background(), instanceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CoreTraceApiService V1ListThreads", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var instanceId string

		resp, httpRes, err := apiClient.CoreTraceApi.V1ListThreads(context.Background(), instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CoreTraceApiService V1StartCoreTrace", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var instanceId string

		httpRes, err := apiClient.CoreTraceApi.V1StartCoreTrace(context.Background(), instanceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CoreTraceApiService V1StopCoreTrace", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var instanceId string

		httpRes, err := apiClient.CoreTraceApi.V1StopCoreTrace(context.Background(), instanceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
