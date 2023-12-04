# ModelParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelName** | Pointer to **string** |  | [optional] 
**TopP** | Pointer to **float64** |  | [optional] 
**TopK** | Pointer to **int32** |  | [optional] 
**Seed** | Pointer to **int64** |  | [optional] 
**MinLength** | Pointer to **int32** |  | [optional] 
**MaxLength** | Pointer to **int32** |  | [optional] 
**Temperature** | Pointer to **float64** |  | [optional] 

## Methods

### NewModelParameters

`func NewModelParameters() *ModelParameters`

NewModelParameters instantiates a new ModelParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelParametersWithDefaults

`func NewModelParametersWithDefaults() *ModelParameters`

NewModelParametersWithDefaults instantiates a new ModelParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelName

`func (o *ModelParameters) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *ModelParameters) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *ModelParameters) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *ModelParameters) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### GetTopP

`func (o *ModelParameters) GetTopP() float64`

GetTopP returns the TopP field if non-nil, zero value otherwise.

### GetTopPOk

`func (o *ModelParameters) GetTopPOk() (*float64, bool)`

GetTopPOk returns a tuple with the TopP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopP

`func (o *ModelParameters) SetTopP(v float64)`

SetTopP sets TopP field to given value.

### HasTopP

`func (o *ModelParameters) HasTopP() bool`

HasTopP returns a boolean if a field has been set.

### GetTopK

`func (o *ModelParameters) GetTopK() int32`

GetTopK returns the TopK field if non-nil, zero value otherwise.

### GetTopKOk

`func (o *ModelParameters) GetTopKOk() (*int32, bool)`

GetTopKOk returns a tuple with the TopK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopK

`func (o *ModelParameters) SetTopK(v int32)`

SetTopK sets TopK field to given value.

### HasTopK

`func (o *ModelParameters) HasTopK() bool`

HasTopK returns a boolean if a field has been set.

### GetSeed

`func (o *ModelParameters) GetSeed() int64`

GetSeed returns the Seed field if non-nil, zero value otherwise.

### GetSeedOk

`func (o *ModelParameters) GetSeedOk() (*int64, bool)`

GetSeedOk returns a tuple with the Seed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeed

`func (o *ModelParameters) SetSeed(v int64)`

SetSeed sets Seed field to given value.

### HasSeed

`func (o *ModelParameters) HasSeed() bool`

HasSeed returns a boolean if a field has been set.

### GetMinLength

`func (o *ModelParameters) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *ModelParameters) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *ModelParameters) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *ModelParameters) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetMaxLength

`func (o *ModelParameters) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *ModelParameters) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *ModelParameters) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *ModelParameters) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetTemperature

`func (o *ModelParameters) GetTemperature() float64`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *ModelParameters) GetTemperatureOk() (*float64, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *ModelParameters) SetTemperature(v float64)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *ModelParameters) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


