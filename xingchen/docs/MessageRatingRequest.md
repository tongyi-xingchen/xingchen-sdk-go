# MessageRatingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayAddContent** | Pointer to [**GatewayIssuedParams**](GatewayIssuedParams.md) |  | [optional] 
**MessageId** | **string** | 待评分消息ID | 
**Rating** | **int32** | 评分 | 
**Info** | Pointer to **string** | 差评的具体信息 | [optional] 

## Methods

### NewMessageRatingRequest

`func NewMessageRatingRequest(messageId string, rating int32, ) *MessageRatingRequest`

NewMessageRatingRequest instantiates a new MessageRatingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageRatingRequestWithDefaults

`func NewMessageRatingRequestWithDefaults() *MessageRatingRequest`

NewMessageRatingRequestWithDefaults instantiates a new MessageRatingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayAddContent

`func (o *MessageRatingRequest) GetGatewayAddContent() GatewayIssuedParams`

GetGatewayAddContent returns the GatewayAddContent field if non-nil, zero value otherwise.

### GetGatewayAddContentOk

`func (o *MessageRatingRequest) GetGatewayAddContentOk() (*GatewayIssuedParams, bool)`

GetGatewayAddContentOk returns a tuple with the GatewayAddContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayAddContent

`func (o *MessageRatingRequest) SetGatewayAddContent(v GatewayIssuedParams)`

SetGatewayAddContent sets GatewayAddContent field to given value.

### HasGatewayAddContent

`func (o *MessageRatingRequest) HasGatewayAddContent() bool`

HasGatewayAddContent returns a boolean if a field has been set.

### GetMessageId

`func (o *MessageRatingRequest) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *MessageRatingRequest) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *MessageRatingRequest) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.


### GetRating

`func (o *MessageRatingRequest) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *MessageRatingRequest) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *MessageRatingRequest) SetRating(v int32)`

SetRating sets Rating field to given value.


### GetInfo

`func (o *MessageRatingRequest) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *MessageRatingRequest) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *MessageRatingRequest) SetInfo(v string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *MessageRatingRequest) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


