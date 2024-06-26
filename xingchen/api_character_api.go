/*
XingChen 开放接口定义

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xingchen

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// CharacterApiSubService CharacterApiSub service
type CharacterApiSubService service

type ApiCharacterDetailsRequest struct {
	ctx         context.Context
	ApiService  *CharacterApiSubService
	characterId *string
	version     *int32
}

// 角色ID
func (r ApiCharacterDetailsRequest) CharacterId(characterId string) ApiCharacterDetailsRequest {
	r.characterId = &characterId
	return r
}

// 角色版本
func (r ApiCharacterDetailsRequest) Version(version int32) ApiCharacterDetailsRequest {
	r.version = &version
	return r
}

func (r ApiCharacterDetailsRequest) Execute() (*ResultDTOCharacterDTO, *http.Response, error) {
	return r.ApiService.CharacterDetailsExecute(r)
}

/*
CharacterDetails 角色详情

获取角色详细信息

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCharacterDetailsRequest
*/
func (a *CharacterApiSubService) CharacterDetails(ctx context.Context) ApiCharacterDetailsRequest {
	return ApiCharacterDetailsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ResultDTOCharacterDTO
func (a *CharacterApiSubService) CharacterDetailsExecute(r ApiCharacterDetailsRequest) (*ResultDTOCharacterDTO, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ResultDTOCharacterDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CharacterApiSubService.CharacterDetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/character/details"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.characterId == nil {
		return localVarReturnValue, nil, reportError("characterId is required and must be specified")
	}
	if r.version == nil {
		return localVarReturnValue, nil, reportError("version is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "characterId", r.characterId, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreateRequest struct {
	ctx                context.Context
	ApiService         *CharacterApiSubService
	characterCreateDTO *CharacterCreateDTO
}

func (r ApiCreateRequest) CharacterCreateDTO(characterCreateDTO CharacterCreateDTO) ApiCreateRequest {
	r.characterCreateDTO = &characterCreateDTO
	return r
}

func (r ApiCreateRequest) Execute() (*ResultDTOCharacterKey, *http.Response, error) {
	return r.ApiService.CreateExecute(r)
}

/*
Create 创建角色

基于角色名称、人设和对话示例定义角色，并返回角色ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateRequest
*/
func (a *CharacterApiSubService) Create(ctx context.Context) ApiCreateRequest {
	return ApiCreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ResultDTOCharacterKey
func (a *CharacterApiSubService) CreateExecute(r ApiCreateRequest) (*ResultDTOCharacterKey, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ResultDTOCharacterKey
	)

	err := checkPlugins(r.characterCreateDTO.AdvancedConfig.PlatformPlugins)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: "unsupport plugin"}
	}

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CharacterApiSubService.Create")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/character/create"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.characterCreateDTO == nil {
		return localVarReturnValue, nil, reportError("characterCreateDTO is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.characterCreateDTO
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreateOrUpdateVersionRequest struct {
	ctx                               context.Context
	ApiService                        *CharacterApiSubService
	characterVersionCreateOrUpdateDTO *CharacterVersionCreateOrUpdateDTO
}

// 待更新的角色信息
func (r ApiCreateOrUpdateVersionRequest) CharacterVersionCreateOrUpdateDTO(characterVersionCreateOrUpdateDTO CharacterVersionCreateOrUpdateDTO) ApiCreateOrUpdateVersionRequest {
	r.characterVersionCreateOrUpdateDTO = &characterVersionCreateOrUpdateDTO
	return r
}

func (r ApiCreateOrUpdateVersionRequest) Execute() (*ResultDTOCharacterDTO, *http.Response, error) {
	return r.ApiService.CreateOrUpdateVersionExecute(r)
}

/*
CreateOrUpdateVersion 创建或更新角色版本

必须字段：characterId 和 version。返回角色版本详情。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateOrUpdateVersionRequest
*/
func (a *CharacterApiSubService) CreateOrUpdateVersion(ctx context.Context) ApiCreateOrUpdateVersionRequest {
	return ApiCreateOrUpdateVersionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ResultDTOCharacterDTO
func (a *CharacterApiSubService) CreateOrUpdateVersionExecute(r ApiCreateOrUpdateVersionRequest) (*ResultDTOCharacterDTO, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ResultDTOCharacterDTO
	)

	err := checkPlugins(r.characterVersionCreateOrUpdateDTO.AdvancedConfig.PlatformPlugins)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: "unsupport plugin"}
	}

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CharacterApiSubService.CreateOrUpdateVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/character/createOrUpdateVersion"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.characterVersionCreateOrUpdateDTO == nil {
		return localVarReturnValue, nil, reportError("characterVersionCreateOrUpdateDTO is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.characterVersionCreateOrUpdateDTO
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteRequest struct {
	ctx         context.Context
	ApiService  *CharacterApiSubService
	characterId *string
	version     *int32
}

// 待删除的 characterId
func (r ApiDeleteRequest) CharacterId(characterId string) ApiDeleteRequest {
	r.characterId = &characterId
	return r
}

// 待删除的版本
func (r ApiDeleteRequest) Version(version int32) ApiDeleteRequest {
	r.version = &version
	return r
}

func (r ApiDeleteRequest) Execute() (*ResultDTOBoolean, *http.Response, error) {
	return r.ApiService.DeleteExecute(r)
}

/*
Delete 删除角色

（逻辑）删除 character。返回成功或失败。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiDeleteRequest
*/
func (a *CharacterApiSubService) Delete(ctx context.Context) ApiDeleteRequest {
	return ApiDeleteRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ResultDTOBoolean
func (a *CharacterApiSubService) DeleteExecute(r ApiDeleteRequest) (*ResultDTOBoolean, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ResultDTOBoolean
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CharacterApiSubService.Delete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/character/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.characterId == nil {
		return localVarReturnValue, nil, reportError("characterId is required and must be specified")
	}
	if r.version == nil {
		return localVarReturnValue, nil, reportError("version is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "characterId", r.characterId, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarPostBody = CharacterKey{
		CharacterId: r.characterId,
		Version:     r.version,
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListCharacterVersionsRequest struct {
	ctx         context.Context
	ApiService  *CharacterApiSubService
	characterId string
}

func (r ApiListCharacterVersionsRequest) Execute() (*ResultDTOListCharacterDTO, *http.Response, error) {
	return r.ApiService.ListCharacterVersionsExecute(r)
}

/*
ListCharacterVersions 角色版本列表

获取角色版本列表

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param characterId 角色 characterId
	@return ApiListCharacterVersionsRequest
*/
func (a *CharacterApiSubService) ListCharacterVersions(ctx context.Context, characterId string) ApiListCharacterVersionsRequest {
	return ApiListCharacterVersionsRequest{
		ApiService:  a,
		ctx:         ctx,
		characterId: characterId,
	}
}

// Execute executes the request
//
//	@return ResultDTOListCharacterDTO
func (a *CharacterApiSubService) ListCharacterVersionsExecute(r ApiListCharacterVersionsRequest) (*ResultDTOListCharacterDTO, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ResultDTOListCharacterDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CharacterApiSubService.ListCharacterVersions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/character/versions/{characterId}"
	localVarPath = strings.Replace(localVarPath, "{"+"characterId"+"}", url.PathEscape(parameterValueToString(r.characterId, "characterId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRecommendCharacterVersionRequest struct {
	ctx         context.Context
	ApiService  *CharacterApiSubService
	characterId string
}

func (r ApiRecommendCharacterVersionRequest) Execute() (*ResultDTOCharacterDTO, *http.Response, error) {
	return r.ApiService.RecommendCharacterVersionExecute(r)
}

/*
RecommendCharacterVersion 角色版本列表

获取角色版本列表

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param characterId 角色 characterId
	@return ApiRecommendCharacterVersionRequest
*/
func (a *CharacterApiSubService) RecommendCharacterVersion(ctx context.Context, characterId string) ApiRecommendCharacterVersionRequest {
	return ApiRecommendCharacterVersionRequest{
		ApiService:  a,
		ctx:         ctx,
		characterId: characterId,
	}
}

// Execute executes the request
//
//	@return ResultDTOCharacterDTO
func (a *CharacterApiSubService) RecommendCharacterVersionExecute(r ApiRecommendCharacterVersionRequest) (*ResultDTOCharacterDTO, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ResultDTOCharacterDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CharacterApiSubService.RecommendCharacterVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/character/newversion/recommend/{characterId}"
	localVarPath = strings.Replace(localVarPath, "{"+"characterId"+"}", url.PathEscape(parameterValueToString(r.characterId, "characterId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSearchRequest struct {
	ctx               context.Context
	ApiService        *CharacterApiSubService
	characterQueryDTO *CharacterQueryDTO
}

// 查询条件
func (r ApiSearchRequest) CharacterQueryDTO(characterQueryDTO CharacterQueryDTO) ApiSearchRequest {
	r.characterQueryDTO = &characterQueryDTO
	return r
}

func (r ApiSearchRequest) Execute() (*ResultDTOPageResultCharacterDTO, *http.Response, error) {
	return r.ApiService.SearchExecute(r)
}

/*
Search 查询角色

查询角色：
- 可以指定的查询字段，and 关系：
  - 角色名称：左匹配
  - 查询范围
  - my - 只查询我创建的角色
  - public - 查询平台开放的角色
  - pre_configured - 预制角色

- 可以指定一定的排序规则，如更新时间倒排。
- 搜索结果为 角色详情。
- 支持分页。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSearchRequest
*/
func (a *CharacterApiSubService) Search(ctx context.Context) ApiSearchRequest {
	return ApiSearchRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ResultDTOPageResultCharacterDTO
func (a *CharacterApiSubService) SearchExecute(r ApiSearchRequest) (*ResultDTOPageResultCharacterDTO, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ResultDTOPageResultCharacterDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CharacterApiSubService.Search")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/character/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.characterQueryDTO == nil {
		return localVarReturnValue, nil, reportError("characterQueryDTO is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.characterQueryDTO
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateRequest struct {
	ctx                context.Context
	ApiService         *CharacterApiSubService
	characterUpdateDTO *CharacterUpdateDTO
}

// 待更新的角色信息
func (r ApiUpdateRequest) CharacterUpdateDTO(characterUpdateDTO CharacterUpdateDTO) ApiUpdateRequest {
	r.characterUpdateDTO = &characterUpdateDTO
	return r
}

func (r ApiUpdateRequest) Execute() (*ResultDTOBoolean, *http.Response, error) {
	return r.ApiService.UpdateExecute(r)
}

/*
Update 更新角色信息

必须字段：characterId。返回成功或失败。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiUpdateRequest
*/
func (a *CharacterApiSubService) Update(ctx context.Context) ApiUpdateRequest {
	return ApiUpdateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ResultDTOBoolean
func (a *CharacterApiSubService) UpdateExecute(r ApiUpdateRequest) (*ResultDTOBoolean, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ResultDTOBoolean
	)

	err := checkPlugins(r.characterUpdateDTO.AdvancedConfig.PlatformPlugins)
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: "unsupport plugin"}
	}

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CharacterApiSubService.Update")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/character/update"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.characterUpdateDTO == nil {
		return localVarReturnValue, nil, reportError("characterUpdateDTO is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.characterUpdateDTO
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCharacterDescGenerateRequest struct {
	ctx              context.Context
	ApiService       *CharacterApiSubService
	characterDescReq *CharacterDescGenerateRequest
}

func (r ApiCharacterDescGenerateRequest) CharacterDescGenerate(characterDescReq CharacterDescGenerateRequest) ApiCharacterDescGenerateRequest {
	r.characterDescReq = &characterDescReq
	return r
}

func (r ApiCharacterDescGenerateRequest) Execute() (*ResultDTOCharacterDescDTO, *http.Response, error) {
	return r.ApiService.CharacterDescGenerateExecute(r)
}

func (a *CharacterApiSubService) CharacterDescGenerate(ctx context.Context) ApiCharacterDescGenerateRequest {
	return ApiCharacterDescGenerateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

func (a *CharacterApiSubService) CharacterDescGenerateExecute(r ApiCharacterDescGenerateRequest) (*ResultDTOCharacterDescDTO, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ResultDTOCharacterDescDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CharacterApiSubService.CharacterDescGenerate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/character/auto/desc"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.characterDescReq == nil {
		return localVarReturnValue, nil, reportError("characterDescReq is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	localVarPostBody = r.characterDescReq

	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

func checkPlugins(plugins []interface{}) error {
	for _, plugin := range plugins {
		switch interface{}(plugin).(type) {
		case TextToImagePlugin, RejectAnswerPlugin:
			continue
		default:
			return errors.New("unsupport plugin")
		}
	}
	return nil
}
