package xingchen

import (
	"context"
	openapiclient "github.com/tongyi-xingchen/xingchen-sdk-go/xingchen"
	"testing"
)

func Test_xingchen_KnowledgeBaseApiSubService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	bearer := "xxx"
	ctx := context.WithValue(context.Background(), openapiclient.ContextAccessToken, bearer)

	t.Run("Test KnowledgeBaseApiSubService Create", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		createDTO := openapiclient.KnowledgeBaseCreateDTO{
			Name:        openapiclient.PtrString("test"),
			Description: openapiclient.PtrString("test"),
		}
		resp, httpRes, err := apiClient.KnowledgeBaseApiSubService.Create(ctx).CreateDTO(createDTO).Execute()

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

	t.Run("Test KnowledgeBaseApiSubService Update", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		updateDTO := openapiclient.KnowledgeBaseUpdateDTO{
			KnowledgeBaseId: openapiclient.PtrString("xxx"),
			Name:            openapiclient.PtrString("go tests"),
			Description:     openapiclient.PtrString("test"),
		}
		resp, httpRes, err := apiClient.KnowledgeBaseApiSubService.Update(ctx).UpdateDTO(updateDTO).Execute()

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

	t.Run("Test KnowledgeBaseApiSubService Search", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		searchDTO := openapiclient.KnowledgeBaseQueryDTO{
			PageSize: openapiclient.PtrInt32(10),
			PageNum:  openapiclient.PtrInt32(1),
		}
		resp, httpRes, err := apiClient.KnowledgeBaseApiSubService.Search(ctx).SearchDTO(searchDTO).Execute()

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

	t.Run("Test KnowledgeBaseApiSubService Delete", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		deleteDTO := openapiclient.KnowledgeBaseDeleteDTO{
			KnowledgeBaseId: openapiclient.PtrString("xxx"),
		}
		resp, httpRes, err := apiClient.KnowledgeBaseApiSubService.Delete(ctx).DeleteDTO(deleteDTO).Execute()

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

	t.Run("Test KnowledgeBaseApiSubService DetailUpload", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		uploadDTO := openapiclient.KnowledgeBaseDetailUploadDTO{
			KnowledgeBaseId: openapiclient.PtrString("xxx"),
			Type:            openapiclient.PtrString("text"),
			FileInfos: []openapiclient.FileInfoVO{
				{
					Filename: openapiclient.PtrString("测试.txt"),
					FileUrl:  openapiclient.PtrString("xxx"),
				},
			},
		}
		resp, httpRes, err := apiClient.KnowledgeBaseApiSubService.DetailUpload(ctx).DetailUploadDTO(uploadDTO).Execute()

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

	t.Run("Test KnowledgeBaseApiSubService DetailUpdate", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		updateDTO := openapiclient.KnowledgeBaseDetailUpdateDTO{
			KnowledgeBaseId: openapiclient.PtrString("xxx"),
			Name:            openapiclient.PtrString("测试.txt"),
			NewName:         openapiclient.PtrString("test.txt"),
		}
		resp, httpRes, err := apiClient.KnowledgeBaseApiSubService.DetailUpdate(ctx).DetailUpdateDTO(updateDTO).Execute()

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

	t.Run("Test KnowledgeBaseApiSubService DetailSearch", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		queryDTO := openapiclient.KnowledgeBaseDetailQueryDTO{
			KnowledgeBaseId: openapiclient.PtrString("xxx"),
		}
		resp, httpRes, err := apiClient.KnowledgeBaseApiSubService.DetailSearch(ctx).DetailSearchDTO(queryDTO).Execute()

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

	t.Run("Test KnowledgeBaseApiSubService DetailDelete", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		deleteDTO := openapiclient.KnowledgeBaseDetailDeleteDTO{
			KnowledgeBaseId: openapiclient.PtrString("xxx"),
			Name:            openapiclient.PtrString("test.txt"),
		}
		resp, httpRes, err := apiClient.KnowledgeBaseApiSubService.DetailDelete(ctx).DetailDeleteDTO(deleteDTO).Execute()

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
