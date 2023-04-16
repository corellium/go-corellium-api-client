/*
Corellium API

Testing StatusApiService

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

func Test_corellium_StatusApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test StatusApiService V1Ready", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.StatusApi.V1Ready(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
