# GatewayIssuedParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserInfo** | Pointer to **map[string]string** |  | [optional] 
**Context** | Pointer to [**GatewayContext**](GatewayContext.md) |  | [optional] 

## Methods

### NewGatewayIssuedParams

`func NewGatewayIssuedParams() *GatewayIssuedParams`

NewGatewayIssuedParams instantiates a new GatewayIssuedParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayIssuedParamsWithDefaults

`func NewGatewayIssuedParamsWithDefaults() *GatewayIssuedParams`

NewGatewayIssuedParamsWithDefaults instantiates a new GatewayIssuedParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserInfo

`func (o *GatewayIssuedParams) GetUserInfo() map[string]string`

GetUserInfo returns the UserInfo field if non-nil, zero value otherwise.

### GetUserInfoOk

`func (o *GatewayIssuedParams) GetUserInfoOk() (*map[string]string, bool)`

GetUserInfoOk returns a tuple with the UserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfo

`func (o *GatewayIssuedParams) SetUserInfo(v map[string]string)`

SetUserInfo sets UserInfo field to given value.

### HasUserInfo

`func (o *GatewayIssuedParams) HasUserInfo() bool`

HasUserInfo returns a boolean if a field has been set.

### GetContext

`func (o *GatewayIssuedParams) GetContext() GatewayContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *GatewayIssuedParams) GetContextOk() (*GatewayContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *GatewayIssuedParams) SetContext(v GatewayContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *GatewayIssuedParams) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


