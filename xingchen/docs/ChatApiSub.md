# \ChatApiSub

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Chat**](ChatApiSub.md#Chat) | **Post** /v1/api/chat/send | 用户对话



## Chat

> map[string]interface{} Chat(ctx).ChatReqParams(chatReqParams).Execute()

用户对话



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/xingchen"
)

func main() {
    chatReqParams := *openapiclient.NewChatReqParams(*openapiclient.NewBotProfile(), []openapiclient.Message{*openapiclient.NewMessage()}, *openapiclient.NewUserProfile("UserId_example")) // ChatReqParams | 对话请求

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChatApiSub.Chat(context.Background()).ChatReqParams(chatReqParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatApiSub.Chat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Chat`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ChatApiSub.Chat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatReqParams** | [**ChatReqParams**](ChatReqParams.md) | 对话请求 | 

### Return type

**map[string]interface{}**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

