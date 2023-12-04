# ModelConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxLength** | Pointer to **int32** | 角色说话的最长字数 | [optional] 
**MinLength** | Pointer to **int32** | 角色说话的最长字数 | [optional] 
**TopP** | Pointer to **float64** | 可选，默认切在“默认”选项，选择“其他”时可以自行填入数字（提示最大为1？）；默认（背后参数为0.8） | [optional] 
**Temperature** | Pointer to **float64** | 可选，默认切在“默认”选项，共4个选项（默认、风格稳定、自由发散、其他），选择“其他”时可以自行填入数字（提示最大为xx）；默认（背后参数为1），风格稳定（背后参数为0.92），自由发散（背后参数为1.02） | [optional] 

## Methods

### NewModelConfig

`func NewModelConfig() *ModelConfig`

NewModelConfig instantiates a new ModelConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelConfigWithDefaults

`func NewModelConfigWithDefaults() *ModelConfig`

NewModelConfigWithDefaults instantiates a new ModelConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxLength

`func (o *ModelConfig) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *ModelConfig) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *ModelConfig) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *ModelConfig) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetMinLength

`func (o *ModelConfig) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *ModelConfig) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *ModelConfig) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *ModelConfig) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetTopP

`func (o *ModelConfig) GetTopP() float64`

GetTopP returns the TopP field if non-nil, zero value otherwise.

### GetTopPOk

`func (o *ModelConfig) GetTopPOk() (*float64, bool)`

GetTopPOk returns a tuple with the TopP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopP

`func (o *ModelConfig) SetTopP(v float64)`

SetTopP sets TopP field to given value.

### HasTopP

`func (o *ModelConfig) HasTopP() bool`

HasTopP returns a boolean if a field has been set.

### GetTemperature

`func (o *ModelConfig) GetTemperature() float64`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *ModelConfig) GetTemperatureOk() (*float64, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *ModelConfig) SetTemperature(v float64)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *ModelConfig) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


