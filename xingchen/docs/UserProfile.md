# UserProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** |  | 
**UserName** | Pointer to **string** |  | [optional] 
**BasicInfo** | Pointer to **string** |  | [optional] 

## Methods

### NewUserProfile

`func NewUserProfile(userId string, ) *UserProfile`

NewUserProfile instantiates a new UserProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserProfileWithDefaults

`func NewUserProfileWithDefaults() *UserProfile`

NewUserProfileWithDefaults instantiates a new UserProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *UserProfile) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserProfile) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserProfile) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetUserName

`func (o *UserProfile) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UserProfile) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UserProfile) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *UserProfile) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetBasicInfo

`func (o *UserProfile) GetBasicInfo() string`

GetBasicInfo returns the BasicInfo field if non-nil, zero value otherwise.

### GetBasicInfoOk

`func (o *UserProfile) GetBasicInfoOk() (*string, bool)`

GetBasicInfoOk returns a tuple with the BasicInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicInfo

`func (o *UserProfile) SetBasicInfo(v string)`

SetBasicInfo sets BasicInfo field to given value.

### HasBasicInfo

`func (o *UserProfile) HasBasicInfo() bool`

HasBasicInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


