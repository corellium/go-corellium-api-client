/*
Corellium API

Testing AgentApiService

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

func Test_corellium_AgentApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AgentApiService V1AgentAppReady", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var instanceId string

		resp, httpRes, err := apiClient.AgentApi.V1AgentAppReady(context.Background(), instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentApiService V1AgentDeleteFile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var instanceId string
		var filePath string

		httpRes, err := apiClient.AgentApi.V1AgentDeleteFile(context.Background(), instanceId, filePath).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentApiService V1AgentGetFile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var instanceId string
		var filePath string

		resp, httpRes, err := apiClient.AgentApi.V1AgentGetFile(context.Background(), instanceId, filePath).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentApiService V1AgentGetTempFilename", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var instanceId string

		resp, httpRes, err := apiClient.AgentApi.V1AgentGetTempFilename(context.Background(), instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentApiService V1AgentInstallApp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var instanceId string

		httpRes, err := apiClient.AgentApi.V1AgentInstallApp(context.Background(), instanceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentApiService V1AgentInstallProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var instanceId string

		httpRes, err := apiClient.AgentApi.V1AgentInstallProfile(context.Background(), instanceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentApiService V1AgentKillApp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var instanceId string
		var bundleId string

		resp, httpRes, err := apiClient.AgentApi.V1AgentKillApp(context.Background(), instanceId, bundleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentApiService V1AgentListAppIcons", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var instanceId string

		resp, httpRes, err := apiClient.AgentApi.V1AgentListAppIcons(context.Background(), instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentApiService V1AgentListApps", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var instanceId string

		resp, httpRes, err := apiClient.AgentApi.V1AgentListApps(context.Background(), instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentApiService V1AgentListAppsStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var instanceId string

		resp, httpRes, err := apiClient.AgentApi.V1AgentListAppsStatus(context.Background(), instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentApiService V1AgentListProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var instanceId string

		resp, httpRes, err := apiClient.AgentApi.V1AgentListProfiles(context.Background(), instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentApiService V1AgentRunApp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var instanceId string
		var bundleId string

		resp, httpRes, err := apiClient.AgentApi.V1AgentRunApp(context.Background(), instanceId, bundleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentApiService V1AgentSetFileAttributes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var instanceId string
		var filePath string

		httpRes, err := apiClient.AgentApi.V1AgentSetFileAttributes(context.Background(), instanceId, filePath).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentApiService V1AgentSystemGetAdbAuth", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var instanceId string

		resp, httpRes, err := apiClient.AgentApi.V1AgentSystemGetAdbAuth(context.Background(), instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentApiService V1AgentSystemGetNetwork", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var instanceId string

		resp, httpRes, err := apiClient.AgentApi.V1AgentSystemGetNetwork(context.Background(), instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentApiService V1AgentSystemGetProp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var instanceId string

		resp, httpRes, err := apiClient.AgentApi.V1AgentSystemGetProp(context.Background(), instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentApiService V1AgentSystemInstallOpenGApps", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var instanceId string

		httpRes, err := apiClient.AgentApi.V1AgentSystemInstallOpenGApps(context.Background(), instanceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentApiService V1AgentSystemLock", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var instanceId string

		httpRes, err := apiClient.AgentApi.V1AgentSystemLock(context.Background(), instanceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentApiService V1AgentSystemSetAdbAuth", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var instanceId string

		httpRes, err := apiClient.AgentApi.V1AgentSystemSetAdbAuth(context.Background(), instanceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentApiService V1AgentSystemSetHostname", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var instanceId string

		httpRes, err := apiClient.AgentApi.V1AgentSystemSetHostname(context.Background(), instanceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentApiService V1AgentSystemShutdown", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var instanceId string

		httpRes, err := apiClient.AgentApi.V1AgentSystemShutdown(context.Background(), instanceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentApiService V1AgentSystemUnlock", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var instanceId string

		httpRes, err := apiClient.AgentApi.V1AgentSystemUnlock(context.Background(), instanceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentApiService V1AgentUninstallApp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var instanceId string
		var bundleId string

		resp, httpRes, err := apiClient.AgentApi.V1AgentUninstallApp(context.Background(), instanceId, bundleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentApiService V1AgentUninstallProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var instanceId string
		var profileId string

		httpRes, err := apiClient.AgentApi.V1AgentUninstallProfile(context.Background(), instanceId, profileId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentApiService V1AgentUploadFile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var instanceId string
		var filePath string

		httpRes, err := apiClient.AgentApi.V1AgentUploadFile(context.Background(), instanceId, filePath).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
