# Scenario

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Topics** | Pointer to **[]string** |  | [optional] 
**UserTags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewScenario

`func NewScenario() *Scenario`

NewScenario instantiates a new Scenario object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScenarioWithDefaults

`func NewScenarioWithDefaults() *Scenario`

NewScenarioWithDefaults instantiates a new Scenario object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Scenario) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Scenario) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Scenario) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Scenario) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTopics

`func (o *Scenario) GetTopics() []string`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *Scenario) GetTopicsOk() (*[]string, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *Scenario) SetTopics(v []string)`

SetTopics sets Topics field to given value.

### HasTopics

`func (o *Scenario) HasTopics() bool`

HasTopics returns a boolean if a field has been set.

### GetUserTags

`func (o *Scenario) GetUserTags() []string`

GetUserTags returns the UserTags field if non-nil, zero value otherwise.

### GetUserTagsOk

`func (o *Scenario) GetUserTagsOk() (*[]string, bool)`

GetUserTagsOk returns a tuple with the UserTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTags

`func (o *Scenario) SetUserTags(v []string)`

SetUserTags sets UserTags field to given value.

### HasUserTags

`func (o *Scenario) HasUserTags() bool`

HasUserTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


