# \CharacterApiSub

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CharacterDetails**](CharacterApiSub.md#CharacterDetails) | **Get** /v1/api/character/details | 角色详情
[**Create**](CharacterApiSub.md#Create) | **Post** /v1/api/character/create | 创建角色
[**CreateOrUpdateVersion**](CharacterApiSub.md#CreateOrUpdateVersion) | **Put** /v1/api/character/createOrUpdateVersion | 创建或更新角色版本
[**Delete**](CharacterApiSub.md#Delete) | **Delete** /v1/api/character/delete | 删除角色
[**ListCharacterVersions**](CharacterApiSub.md#ListCharacterVersions) | **Get** /v1/api/character/versions/{characterId} | 角色版本列表
[**RecommendCharacterVersion**](CharacterApiSub.md#RecommendCharacterVersion) | **Get** /v1/api/character/newversion/recommend/{characterId} | 角色版本列表
[**Search**](CharacterApiSub.md#Search) | **Post** /v1/api/character/search | 查询角色
[**Update**](CharacterApiSub.md#Update) | **Put** /v1/api/character/update | 更新角色信息



## CharacterDetails

> ResultDTOCharacterDTO CharacterDetails(ctx).CharacterId(characterId).Version(version).Execute()

角色详情



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
    characterId := "characterId_example" // string | 角色ID
    version := int32(56) // int32 | 角色版本

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CharacterApiSub.CharacterDetails(context.Background()).CharacterId(characterId).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CharacterApiSub.CharacterDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CharacterDetails`: ResultDTOCharacterDTO
    fmt.Fprintf(os.Stdout, "Response from `CharacterApiSub.CharacterDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCharacterDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **string** | 角色ID | 
 **version** | **int32** | 角色版本 | 

### Return type

[**ResultDTOCharacterDTO**](ResultDTOCharacterDTO.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create

> ResultDTOCharacterKey Create(ctx).CharacterCreateDTO(characterCreateDTO).Execute()

创建角色



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
    characterCreateDTO := *openapiclient.NewCharacterCreateDTO("Name_example", "Introduction_example", "BasicInformation_example", "OpeningLine_example") // CharacterCreateDTO | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CharacterApiSub.Create(context.Background()).CharacterCreateDTO(characterCreateDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CharacterApiSub.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: ResultDTOCharacterKey
    fmt.Fprintf(os.Stdout, "Response from `CharacterApiSub.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterCreateDTO** | [**CharacterCreateDTO**](CharacterCreateDTO.md) |  | 

### Return type

[**ResultDTOCharacterKey**](ResultDTOCharacterKey.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdateVersion

> ResultDTOCharacterDTO CreateOrUpdateVersion(ctx).CharacterVersionCreateOrUpdateDTO(characterVersionCreateOrUpdateDTO).Execute()

创建或更新角色版本



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
    characterVersionCreateOrUpdateDTO := *openapiclient.NewCharacterVersionCreateOrUpdateDTO("Name_example", "Introduction_example", "BasicInformation_example", "OpeningLine_example", "CharacterId_example", int32(123)) // CharacterVersionCreateOrUpdateDTO | 待更新的角色信息

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CharacterApiSub.CreateOrUpdateVersion(context.Background()).CharacterVersionCreateOrUpdateDTO(characterVersionCreateOrUpdateDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CharacterApiSub.CreateOrUpdateVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateVersion`: ResultDTOCharacterDTO
    fmt.Fprintf(os.Stdout, "Response from `CharacterApiSub.CreateOrUpdateVersion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterVersionCreateOrUpdateDTO** | [**CharacterVersionCreateOrUpdateDTO**](CharacterVersionCreateOrUpdateDTO.md) | 待更新的角色信息 | 

### Return type

[**ResultDTOCharacterDTO**](ResultDTOCharacterDTO.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> ResultDTOBoolean Delete(ctx).CharacterId(characterId).Version(version).Execute()

删除角色



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
    characterId := "characterId_example" // string | 待删除的 characterId
    version := int32(56) // int32 | 待删除的版本

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CharacterApiSub.Delete(context.Background()).CharacterId(characterId).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CharacterApiSub.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: ResultDTOBoolean
    fmt.Fprintf(os.Stdout, "Response from `CharacterApiSub.Delete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **string** | 待删除的 characterId | 
 **version** | **int32** | 待删除的版本 | 

### Return type

[**ResultDTOBoolean**](ResultDTOBoolean.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCharacterVersions

> ResultDTOListCharacterDTO ListCharacterVersions(ctx, characterId).Execute()

角色版本列表



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
    characterId := "characterId_example" // string | 角色 characterId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CharacterApiSub.ListCharacterVersions(context.Background(), characterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CharacterApiSub.ListCharacterVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCharacterVersions`: ResultDTOListCharacterDTO
    fmt.Fprintf(os.Stdout, "Response from `CharacterApiSub.ListCharacterVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **string** | 角色 characterId | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCharacterVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResultDTOListCharacterDTO**](ResultDTOListCharacterDTO.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecommendCharacterVersion

> ResultDTOCharacterDTO RecommendCharacterVersion(ctx, characterId).Execute()

角色版本列表



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
    characterId := "characterId_example" // string | 角色 characterId

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CharacterApiSub.RecommendCharacterVersion(context.Background(), characterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CharacterApiSub.RecommendCharacterVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecommendCharacterVersion`: ResultDTOCharacterDTO
    fmt.Fprintf(os.Stdout, "Response from `CharacterApiSub.RecommendCharacterVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **string** | 角色 characterId | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecommendCharacterVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResultDTOCharacterDTO**](ResultDTOCharacterDTO.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Search

> ResultDTOPageResultCharacterDTO Search(ctx).CharacterQueryDTO(characterQueryDTO).Execute()

查询角色



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
    characterQueryDTO := *openapiclient.NewCharacterQueryDTO() // CharacterQueryDTO | 查询条件

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CharacterApiSub.Search(context.Background()).CharacterQueryDTO(characterQueryDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CharacterApiSub.Search``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Search`: ResultDTOPageResultCharacterDTO
    fmt.Fprintf(os.Stdout, "Response from `CharacterApiSub.Search`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterQueryDTO** | [**CharacterQueryDTO**](CharacterQueryDTO.md) | 查询条件 | 

### Return type

[**ResultDTOPageResultCharacterDTO**](ResultDTOPageResultCharacterDTO.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> ResultDTOBoolean Update(ctx).CharacterUpdateDTO(characterUpdateDTO).Execute()

更新角色信息



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
    characterUpdateDTO := *openapiclient.NewCharacterUpdateDTO("Name_example", "Introduction_example", "BasicInformation_example", "OpeningLine_example", "CharacterId_example", int32(123)) // CharacterUpdateDTO | 待更新的角色信息

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CharacterApiSub.Update(context.Background()).CharacterUpdateDTO(characterUpdateDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CharacterApiSub.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: ResultDTOBoolean
    fmt.Fprintf(os.Stdout, "Response from `CharacterApiSub.Update`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterUpdateDTO** | [**CharacterUpdateDTO**](CharacterUpdateDTO.md) | 待更新的角色信息 | 

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

