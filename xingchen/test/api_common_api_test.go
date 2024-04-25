package xingchen

import (
	"context"
	"fmt"
	openapiclient "github.com/tongyi-xingchen/xingchen-sdk-go/xingchen"
	"testing"
	"time"
)

func Test_xingchen_CommonApiSubService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	configuration.Version = openapiclient.V2
	apiClient := openapiclient.NewAPIClient(configuration)

	bearer := "xxx"
	ctx := context.WithValue(context.Background(), openapiclient.ContextAccessToken, bearer)

	t.Run("Test CommonApiSubService DocConvert", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		request := openapiclient.DocConvertSubmitRequest{
			Type:      openapiclient.PtrString("excelToDocx"),
			SourceUrl: openapiclient.PtrString("https://lang.alicdn.com/xingchen/%E6%B5%8B%E8%AF%95.xlsx"),
		}
		resp, httpRes, err := apiClient.CommonApiSubService.DocConvertSubmit(ctx).DocConvertSubmit(request).Execute()

		if err != nil {
			t.Error()
		}
		if resp == nil {
			t.Error()
		}
		if httpRes.StatusCode != 200 {
			t.Error()
		}
		taskId := resp.Data.Id
		for i := 0; i < 10; i++ {
			resp, httpRes, err := apiClient.CommonApiSubService.DocConvertResult(ctx).TaskId(taskId).Execute()
			if err != nil {
				t.Error()
			}
			if resp == nil {
				t.Error()
			}
			if httpRes.StatusCode != 200 {
				t.Error()
			}
			if *resp.Data.Status == "success" {
				fmt.Println(*resp.Data.Url)
				break
			}
			if *resp.Data.Status == "failed" {
				break
			}
			time.Sleep(10 * time.Second)
		}
	})
}
