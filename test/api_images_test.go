/*
Corellium API

Testing ImagesApiService

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

func Test_corellium_ImagesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ImagesApiService V1CreateImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ImagesApi.V1CreateImage(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesApiService V1DeleteImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var imageId string

		httpRes, err := apiClient.ImagesApi.V1DeleteImage(context.Background(), imageId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesApiService V1GetImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var imageId string

		resp, httpRes, err := apiClient.ImagesApi.V1GetImage(context.Background(), imageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesApiService V1GetImages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ImagesApi.V1GetImages(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesApiService V1UploadImageData", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var imageId string

		resp, httpRes, err := apiClient.ImagesApi.V1UploadImageData(context.Background(), imageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
