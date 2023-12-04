# \ChatMessageApiSub

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatHistories**](ChatMessageApiSub.md#ChatHistories) | **Post** /v1/api/chat/message/histories | 对话历史
[**RateMessage**](ChatMessageApiSub.md#RateMessage) | **Post** /v1/api/chat/rating | 消息评分
[**SysReminder**](ChatMessageApiSub.md#SysReminder) | **Post** /v1/api/chat/reminder | 



## ChatHistories

> ResultDTOPageResultChatMessageDTO ChatHistories(ctx).ChatHistoryQueryDTO(chatHistoryQueryDTO).Execute()

对话历史



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
    chatHistoryQueryDTO := *openapiclient.NewChatHistoryQueryDTO() // ChatHistoryQueryDTO | 对话请求

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChatMessageApiSub.ChatHistories(context.Background()).ChatHistoryQueryDTO(chatHistoryQueryDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatMessageApiSub.ChatHistories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatHistories`: ResultDTOPageResultChatMessageDTO
    fmt.Fprintf(os.Stdout, "Response from `ChatMessageApiSub.ChatHistories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatHistoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatHistoryQueryDTO** | [**ChatHistoryQueryDTO**](ChatHistoryQueryDTO.md) | 对话请求 | 

### Return type

[**ResultDTOPageResultChatMessageDTO**](ResultDTOPageResultChatMessageDTO.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RateMessage

> ResultDTOBoolean RateMessage(ctx).MessageRatingRequest(messageRatingRequest).Execute()

消息评分



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
    messageRatingRequest := *openapiclient.NewMessageRatingRequest("MessageId_example", int32(123)) // MessageRatingRequest | 对话请求

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChatMessageApiSub.RateMessage(context.Background()).MessageRatingRequest(messageRatingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatMessageApiSub.RateMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RateMessage`: ResultDTOBoolean
    fmt.Fprintf(os.Stdout, "Response from `ChatMessageApiSub.RateMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRateMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **messageRatingRequest** | [**MessageRatingRequest**](MessageRatingRequest.md) | 对话请求 | 

### Return type

[**ResultDTOBoolean**](ResultDTOBoolean.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SysReminder

> ResultDTOBoolean SysReminder(ctx).SysReminderRequest(sysReminderRequest).Execute()



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
    sysReminderRequest := *openapiclient.NewSysReminderRequest("CharacterId_example", "Content_example", "BizUserId_example") // SysReminderRequest | 对话请求

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChatMessageApiSub.SysReminder(context.Background()).SysReminderRequest(sysReminderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatMessageApiSub.SysReminder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SysReminder`: ResultDTOBoolean
    fmt.Fprintf(os.Stdout, "Response from `ChatMessageApiSub.SysReminder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSysReminderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sysReminderRequest** | [**SysReminderRequest**](SysReminderRequest.md) | 对话请求 | 

### Return type

[**ResultDTOBoolean**](ResultDTOBoolean.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

