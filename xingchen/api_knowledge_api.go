package xingchen

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

type KnowledgeBaseApiSubService service

type ApiKnowledgeBaseCreateRequest struct {
	ctx                    context.Context
	ApiService             *KnowledgeBaseApiSubService
	knowledgeBaseCreateDTO *KnowledgeBaseCreateDTO
}

func (r ApiKnowledgeBaseCreateRequest) CreateDTO(createDTO KnowledgeBaseCreateDTO) ApiKnowledgeBaseCreateRequest {
	r.knowledgeBaseCreateDTO = &createDTO
	return r
}

func (r ApiKnowledgeBaseCreateRequest) Execute() (*ResultDTOKnowledgeBaseDTO, *http.Response, error) {
	return r.ApiService.CreateExecute(r)
}

func (a *KnowledgeBaseApiSubService) Create(ctx context.Context) ApiKnowledgeBaseCreateRequest {
	return ApiKnowledgeBaseCreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

func (a *KnowledgeBaseApiSubService) CreateExecute(r ApiKnowledgeBaseCreateRequest) (*ResultDTOKnowledgeBaseDTO, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ResultDTOKnowledgeBaseDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KnowledgeBaseApiSubService.Create")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/knowledge_base/create"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.knowledgeBaseCreateDTO == nil {
		return localVarReturnValue, nil, reportError("knowledgeBaseCreateDTO is required and must be specified")
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

	localVarPostBody = r.knowledgeBaseCreateDTO

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

type ApiKnowledgeBaseUpdateRequest struct {
	ctx                    context.Context
	ApiService             *KnowledgeBaseApiSubService
	knowledgeBaseUpdateDTO *KnowledgeBaseUpdateDTO
}

func (r ApiKnowledgeBaseUpdateRequest) UpdateDTO(createDTO KnowledgeBaseUpdateDTO) ApiKnowledgeBaseUpdateRequest {
	r.knowledgeBaseUpdateDTO = &createDTO
	return r
}

func (r ApiKnowledgeBaseUpdateRequest) Execute() (*ResultDTOKnowledgeBaseDTO, *http.Response, error) {
	return r.ApiService.UpdateExecute(r)
}

func (a *KnowledgeBaseApiSubService) Update(ctx context.Context) ApiKnowledgeBaseUpdateRequest {
	return ApiKnowledgeBaseUpdateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

func (a *KnowledgeBaseApiSubService) UpdateExecute(r ApiKnowledgeBaseUpdateRequest) (*ResultDTOKnowledgeBaseDTO, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ResultDTOKnowledgeBaseDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KnowledgeBaseApiSubService.Update")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/knowledge_base/update"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.knowledgeBaseUpdateDTO == nil {
		return localVarReturnValue, nil, reportError("knowledgeBaseUpdateDTO is required and must be specified")
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

	localVarPostBody = r.knowledgeBaseUpdateDTO

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

type ApiKnowledgeBaseSearchRequest struct {
	ctx                   context.Context
	ApiService            *KnowledgeBaseApiSubService
	knowledgeBaseQueryDTO *KnowledgeBaseQueryDTO
}

func (r ApiKnowledgeBaseSearchRequest) SearchDTO(createDTO KnowledgeBaseQueryDTO) ApiKnowledgeBaseSearchRequest {
	r.knowledgeBaseQueryDTO = &createDTO
	return r
}

func (r ApiKnowledgeBaseSearchRequest) Execute() (*ResultDTOPageResultKnowledgeBaseDetailDTO, *http.Response, error) {
	return r.ApiService.QueryExecute(r)
}

func (a *KnowledgeBaseApiSubService) Search(ctx context.Context) ApiKnowledgeBaseSearchRequest {
	return ApiKnowledgeBaseSearchRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

func (a *KnowledgeBaseApiSubService) QueryExecute(r ApiKnowledgeBaseSearchRequest) (*ResultDTOPageResultKnowledgeBaseDetailDTO, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ResultDTOPageResultKnowledgeBaseDetailDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KnowledgeBaseApiSubService.Search")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/knowledge_base/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.knowledgeBaseQueryDTO == nil {
		return localVarReturnValue, nil, reportError("knowledgeBaseQueryDTO is required and must be specified")
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

	localVarPostBody = r.knowledgeBaseQueryDTO

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

type ApiKnowledgeBaseDeleteRequest struct {
	ctx                    context.Context
	ApiService             *KnowledgeBaseApiSubService
	knowledgeBaseDeleteDTO *KnowledgeBaseDeleteDTO
}

func (r ApiKnowledgeBaseDeleteRequest) DeleteDTO(createDTO KnowledgeBaseDeleteDTO) ApiKnowledgeBaseDeleteRequest {
	r.knowledgeBaseDeleteDTO = &createDTO
	return r
}

func (r ApiKnowledgeBaseDeleteRequest) Execute() (*ResultDTOBoolean, *http.Response, error) {
	return r.ApiService.DeleteExecute(r)
}

func (a *KnowledgeBaseApiSubService) Delete(ctx context.Context) ApiKnowledgeBaseDeleteRequest {
	return ApiKnowledgeBaseDeleteRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

func (a *KnowledgeBaseApiSubService) DeleteExecute(r ApiKnowledgeBaseDeleteRequest) (*ResultDTOBoolean, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ResultDTOBoolean
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KnowledgeBaseApiSubService.Delete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/knowledge_base/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.knowledgeBaseDeleteDTO == nil {
		return localVarReturnValue, nil, reportError("knowledgeBaseDeleteDTO is required and must be specified")
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

	localVarPostBody = r.knowledgeBaseDeleteDTO

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

type ApiKnowledgeBaseDetailUploadRequest struct {
	ctx                          context.Context
	ApiService                   *KnowledgeBaseApiSubService
	knowledgeBaseDetailUploadDTO *KnowledgeBaseDetailUploadDTO
}

func (r ApiKnowledgeBaseDetailUploadRequest) DetailUploadDTO(createDTO KnowledgeBaseDetailUploadDTO) ApiKnowledgeBaseDetailUploadRequest {
	r.knowledgeBaseDetailUploadDTO = &createDTO
	return r
}

func (r ApiKnowledgeBaseDetailUploadRequest) Execute() (*ResultDTOBoolean, *http.Response, error) {
	return r.ApiService.DetailUploadExecute(r)
}

func (a *KnowledgeBaseApiSubService) DetailUpload(ctx context.Context) ApiKnowledgeBaseDetailUploadRequest {
	return ApiKnowledgeBaseDetailUploadRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

func (a *KnowledgeBaseApiSubService) DetailUploadExecute(r ApiKnowledgeBaseDetailUploadRequest) (*ResultDTOBoolean, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ResultDTOBoolean
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KnowledgeBaseApiSubService.DetailUpload")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/knowledge_base/detail/upload"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.knowledgeBaseDetailUploadDTO == nil {
		return localVarReturnValue, nil, reportError("knowledgeBaseDetailUploadDTO is required and must be specified")
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

	localVarPostBody = r.knowledgeBaseDetailUploadDTO

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

type ApiKnowledgeBaseDetailUpdateRequest struct {
	ctx                          context.Context
	ApiService                   *KnowledgeBaseApiSubService
	knowledgeBaseDetailUpdateDTO *KnowledgeBaseDetailUpdateDTO
}

func (r ApiKnowledgeBaseDetailUpdateRequest) DetailUpdateDTO(createDTO KnowledgeBaseDetailUpdateDTO) ApiKnowledgeBaseDetailUpdateRequest {
	r.knowledgeBaseDetailUpdateDTO = &createDTO
	return r
}

func (r ApiKnowledgeBaseDetailUpdateRequest) Execute() (*ResultDTOBoolean, *http.Response, error) {
	return r.ApiService.DetailUpdateExecute(r)
}

func (a *KnowledgeBaseApiSubService) DetailUpdate(ctx context.Context) ApiKnowledgeBaseDetailUpdateRequest {
	return ApiKnowledgeBaseDetailUpdateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

func (a *KnowledgeBaseApiSubService) DetailUpdateExecute(r ApiKnowledgeBaseDetailUpdateRequest) (*ResultDTOBoolean, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ResultDTOBoolean
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KnowledgeBaseApiSubService.DetailUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/knowledge_base/detail/update"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.knowledgeBaseDetailUpdateDTO == nil {
		return localVarReturnValue, nil, reportError("knowledgeBaseDetailUpdateDTO is required and must be specified")
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

	localVarPostBody = r.knowledgeBaseDetailUpdateDTO

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

type ApiKnowledgeBaseDetailSearchRequest struct {
	ctx                          context.Context
	ApiService                   *KnowledgeBaseApiSubService
	knowledgeBaseDetailSearchDTO *KnowledgeBaseDetailQueryDTO
}

func (r ApiKnowledgeBaseDetailSearchRequest) DetailSearchDTO(createDTO KnowledgeBaseDetailQueryDTO) ApiKnowledgeBaseDetailSearchRequest {
	r.knowledgeBaseDetailSearchDTO = &createDTO
	return r
}

func (r ApiKnowledgeBaseDetailSearchRequest) Execute() (*ResultDTOPageResultKnowledgeBaseDetailDTO, *http.Response, error) {
	return r.ApiService.DetailSearchExecute(r)
}

func (a *KnowledgeBaseApiSubService) DetailSearch(ctx context.Context) ApiKnowledgeBaseDetailSearchRequest {
	return ApiKnowledgeBaseDetailSearchRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

func (a *KnowledgeBaseApiSubService) DetailSearchExecute(r ApiKnowledgeBaseDetailSearchRequest) (*ResultDTOPageResultKnowledgeBaseDetailDTO, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ResultDTOPageResultKnowledgeBaseDetailDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KnowledgeBaseApiSubService.DetailSearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/knowledge_base/detail/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.knowledgeBaseDetailSearchDTO == nil {
		return localVarReturnValue, nil, reportError("knowledgeBaseDetailQueryDTO is required and must be specified")
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

	localVarPostBody = r.knowledgeBaseDetailSearchDTO

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

type ApiKnowledgeBaseDetailDeleteRequest struct {
	ctx                          context.Context
	ApiService                   *KnowledgeBaseApiSubService
	knowledgeBaseDetailDeleteDTO *KnowledgeBaseDetailDeleteDTO
}

func (r ApiKnowledgeBaseDetailDeleteRequest) DetailDeleteDTO(createDTO KnowledgeBaseDetailDeleteDTO) ApiKnowledgeBaseDetailDeleteRequest {
	r.knowledgeBaseDetailDeleteDTO = &createDTO
	return r
}

func (r ApiKnowledgeBaseDetailDeleteRequest) Execute() (*ResultDTOBoolean, *http.Response, error) {
	return r.ApiService.DetailDeleteExecute(r)
}

func (a *KnowledgeBaseApiSubService) DetailDelete(ctx context.Context) ApiKnowledgeBaseDetailDeleteRequest {
	return ApiKnowledgeBaseDetailDeleteRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

func (a *KnowledgeBaseApiSubService) DetailDeleteExecute(r ApiKnowledgeBaseDetailDeleteRequest) (*ResultDTOBoolean, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ResultDTOBoolean
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KnowledgeBaseApiSubService.DetailDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/knowledge_base/detail/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.knowledgeBaseDetailDeleteDTO == nil {
		return localVarReturnValue, nil, reportError("knowledgeBaseDetailDeleteDTO is required and must be specified")
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

	localVarPostBody = r.knowledgeBaseDetailDeleteDTO

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
