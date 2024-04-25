package xingchen

import (
	"context"
	openapiclient "github.com/tongyi-xingchen/xingchen-sdk-go/xingchen"
	"testing"
)

func Test_xingchen_ChatExtractMessageApiSubService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	bearer := "xxx"
	ctx := context.WithValue(context.Background(), openapiclient.ContextAccessToken, bearer)

	t.Run("Test ChatExtractMessageApiSubService ExtractMemorySummary", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		request := openapiclient.ExtractMemorySummaryRequest{
			Messages: []openapiclient.Message{
				{
					Content: openapiclient.PtrString("你好"),
					Role:    openapiclient.PtrString("user"),
				},
				{
					Content: openapiclient.PtrString("你好呀，今天天气真好"),
					Role:    openapiclient.PtrString("assistant"),
				},
				{
					Content: openapiclient.PtrString("我们一起去打篮球把"),
					Role:    openapiclient.PtrString("user"),
				},
				{
					Content: openapiclient.PtrString("好的呀，那我们一起去打球吧，打完篮球我们去吃火锅"),
					Role:    openapiclient.PtrString("assistant"),
				},
			},
			ModelParameters: &openapiclient.ModelParameters{
				ModelName: openapiclient.PtrString("xingchen-base"),
			},
		}
		resp, httpRes, err := apiClient.ChatExtractMessageApiSubService.ExtractMemorySummary(ctx).
			ExtractMemorySummaryRequest(request).Execute()

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

	t.Run("Test ChatExtractMessageApiSubService ExtractMemoryKV", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		request := openapiclient.ExtractMemoryKVRequest{
			Messages: []openapiclient.Message{
				{
					Content: openapiclient.PtrString("你好"),
					Role:    openapiclient.PtrString("user"),
				},
				{
					Content: openapiclient.PtrString("你好呀"),
					Role:    openapiclient.PtrString("assistant"),
				},
				{
					Content: openapiclient.PtrString("我们喜欢打篮球，你呢"),
					Role:    openapiclient.PtrString("user"),
				},
				{
					Content: openapiclient.PtrString("我喜欢踢足球"),
					Role:    openapiclient.PtrString("assistant"),
				},
			},
			ModelParameters: &openapiclient.ModelParameters{
				ModelName: openapiclient.PtrString("xingchen-base"),
			},
			KVMemoryConfigs: []openapiclient.KVMemoryConfig{
				{
					Enabled:    openapiclient.PtrBool(true),
					MemoryText: openapiclient.PtrString("职业"),
				},
				{
					Enabled:    openapiclient.PtrBool(true),
					MemoryText: openapiclient.PtrString("喜好"),
				},
				{
					Enabled:    openapiclient.PtrBool(true),
					MemoryText: openapiclient.PtrString("性别"),
				},
			},
		}
		resp, httpRes, err := apiClient.ChatExtractMessageApiSubService.ExtractMemoryKV(ctx).
			ExtractMemoryKVRequest(request).Execute()

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
