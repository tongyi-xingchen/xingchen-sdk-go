/*
XingChen 开放接口定义

Testing CharacterApiSubService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package xingchen

import (
	openapiclient "alimt-character-ai-sdk-go/xingchen"
	"context"
	"testing"
)

func Test_xingchen_CharacterApiSubService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	configuration.Version = openapiclient.V2
	apiClient := openapiclient.NewAPIClient(configuration)

	bearer := "xxx"
	ctx := context.WithValue(context.Background(), openapiclient.ContextAccessToken, bearer)

	t.Run("Test CharacterApiSubService CharacterDetails", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CharacterApiSub.CharacterDetails(ctx).CharacterId("xxx").Version(1).Execute()

		if err != nil {
			t.Error()
		}
		if resp == nil {
			t.Error()
		}
		if httpRes.StatusCode != 200 {
			t.Error()
		}
	})

	t.Run("Test CharacterApiSubService Create", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		characterCreateDTO := openapiclient.CharacterCreateDTO{}
		resp, httpRes, err := apiClient.CharacterApiSub.Create(ctx).CharacterCreateDTO(characterCreateDTO).Execute()

		if err != nil {
			t.Error()
		}
		if resp == nil {
			t.Error()
		}
		if httpRes.StatusCode != 200 {
			t.Error()
		}
	})

	t.Run("Test CharacterApiSubService CreateOrUpdateVersion", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		characterVersionCreateOrUpdateDTO := openapiclient.CharacterVersionCreateOrUpdateDTO{}
		resp, httpRes, err := apiClient.CharacterApiSub.CreateOrUpdateVersion(ctx).CharacterVersionCreateOrUpdateDTO(characterVersionCreateOrUpdateDTO).Execute()

		if err != nil {
			t.Error()
		}
		if resp == nil {
			t.Error()
		}
		if httpRes.StatusCode != 200 {
			t.Error()
		}
	})

	t.Run("Test CharacterApiSubService ListCharacterVersions", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CharacterApiSub.ListCharacterVersions(ctx, "xxx").Execute()

		if err != nil {
			t.Error()
		}
		if resp == nil {
			t.Error()
		}
		if httpRes.StatusCode != 200 {
			t.Error()
		}
	})

	t.Run("Test CharacterApiSubService RecommendCharacterVersion", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CharacterApiSub.RecommendCharacterVersion(ctx, "xxx").Execute()

		if err != nil {
			t.Error()
		}
		if resp == nil {
			t.Error()
		}
		if httpRes.StatusCode != 200 {
			t.Error()
		}
	})

	t.Run("Test CharacterApiSubService Search", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		characterQueryDTO := openapiclient.CharacterQueryDTO{
			Where: &openapiclient.CharacterQueryWhere{
				Scope: openapiclient.PtrString("my"),
			},
			PageNum:  openapiclient.PtrInt64(1),
			PageSize: openapiclient.PtrInt64(10),
		}

		resp, httpRes, err := apiClient.CharacterApiSub.Search(ctx).CharacterQueryDTO(characterQueryDTO).Execute()

		if err != nil {
			t.Error()
		}
		if resp == nil {
			t.Error()
		}
		if httpRes.StatusCode != 200 {
			t.Error()
		}
	})

	t.Run("Test CharacterApiSubService Update", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		characterUpdateDTO := openapiclient.CharacterUpdateDTO{}

		resp, httpRes, err := apiClient.CharacterApiSub.Update(ctx).CharacterUpdateDTO(characterUpdateDTO).Execute()

		if err != nil {
			t.Error()
		}
		if resp == nil {
			t.Error()
		}
		if httpRes.StatusCode != 200 {
			t.Error()
		}
	})

	t.Run("Test CharacterApiSubService Delete", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CharacterApiSub.Delete(ctx).
			CharacterId("xxx").Version(1).Execute()

		if err != nil {
			t.Error()
		}
		if resp == nil {
			t.Error()
		}
		if httpRes.StatusCode != 200 {
			t.Error()
		}
	})
}
