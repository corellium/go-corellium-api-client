/*
Corellium API

Testing TeamsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package corellium

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/corellium/go-corellium-api-client"
)

func Test_corellium_TeamsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TeamsApiService V1AddUserToTeam", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var teamId string
		var userId string

		httpRes, err := apiClient.TeamsApi.V1AddUserToTeam(context.Background(), teamId, userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TeamsApiService V1RemoveUserFromTeam", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var teamId string
		var userId string

		httpRes, err := apiClient.TeamsApi.V1RemoveUserFromTeam(context.Background(), teamId, userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TeamsApiService V1TeamChange", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var teamId string

		httpRes, err := apiClient.TeamsApi.V1TeamChange(context.Background(), teamId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TeamsApiService V1TeamCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TeamsApi.V1TeamCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TeamsApiService V1TeamDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var teamId string

		httpRes, err := apiClient.TeamsApi.V1TeamDelete(context.Background(), teamId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TeamsApiService V1Teams", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TeamsApi.V1Teams(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
