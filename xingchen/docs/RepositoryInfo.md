# RepositoryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepositoryFiles** | Pointer to [**[]Repository**](Repository.md) | 知识库文件信息 | [optional] 
**Status** | Pointer to **string** | 知识库同步到向量数据库的状态 | [optional] 

## Methods

### NewRepositoryInfo

`func NewRepositoryInfo() *RepositoryInfo`

NewRepositoryInfo instantiates a new RepositoryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryInfoWithDefaults

`func NewRepositoryInfoWithDefaults() *RepositoryInfo`

NewRepositoryInfoWithDefaults instantiates a new RepositoryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepositoryFiles

`func (o *RepositoryInfo) GetRepositoryFiles() []Repository`

GetRepositoryFiles returns the RepositoryFiles field if non-nil, zero value otherwise.

### GetRepositoryFilesOk

`func (o *RepositoryInfo) GetRepositoryFilesOk() (*[]Repository, bool)`

GetRepositoryFilesOk returns a tuple with the RepositoryFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryFiles

`func (o *RepositoryInfo) SetRepositoryFiles(v []Repository)`

SetRepositoryFiles sets RepositoryFiles field to given value.

### HasRepositoryFiles

`func (o *RepositoryInfo) HasRepositoryFiles() bool`

HasRepositoryFiles returns a boolean if a field has been set.

### GetStatus

`func (o *RepositoryInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RepositoryInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RepositoryInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RepositoryInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


