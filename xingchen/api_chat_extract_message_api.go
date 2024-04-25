package xingchen

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

type ChatExtractMessageApiSubService service

type ApiExtractMemorySummaryRequest struct {
	ctx                         context.Context
	ApiService                  *ChatExtractMessageApiSubService
	extractMemorySummaryRequest *ExtractMemorySummaryRequest
}

func (r ApiExtractMemorySummaryRequest) ExtractMemorySummaryRequest(extractMemorySummaryRequest ExtractMemorySummaryRequest) ApiExtractMemorySummaryRequest {
	r.extractMemorySummaryRequest = &extractMemorySummaryRequest
	return r
}

func (r ApiExtractMemorySummaryRequest) Execute() (*ResultDTOMemorySummaryDTO, *http.Response, error) {
	return r.ApiService.ExtractMemorySummaryExecute(r)
}

func (a *ChatExtractMessageApiSubService) ExtractMemorySummary(ctx context.Context) ApiExtractMemorySummaryRequest {
	return ApiExtractMemorySummaryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

func (a *ChatExtractMessageApiSubService) ExtractMemorySummaryExecute(r ApiExtractMemorySummaryRequest) (*ResultDTOMemorySummaryDTO, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ResultDTOMemorySummaryDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatExtractMessageApiSubService.ExtractMemorySummary")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/extract/summary"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.extractMemorySummaryRequest == nil {
		return localVarReturnValue, nil, reportError("extractMemorySummaryRequest is required and must be specified")
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

	localVarPostBody = r.extractMemorySummaryRequest

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

type ApiExtractMemoryKVRequest struct {
	ctx                    context.Context
	ApiService             *ChatExtractMessageApiSubService
	extractMemoryKVRequest *ExtractMemoryKVRequest
}

func (r ApiExtractMemoryKVRequest) ExtractMemoryKVRequest(extractMemoryKVRequest ExtractMemoryKVRequest) ApiExtractMemoryKVRequest {
	r.extractMemoryKVRequest = &extractMemoryKVRequest
	return r
}

func (r ApiExtractMemoryKVRequest) Execute() (*ResultRTOMemoryKVDTO, *http.Response, error) {
	return r.ApiService.ExtractMemoryKVExecute(r)
}

func (a *ChatExtractMessageApiSubService) ExtractMemoryKV(ctx context.Context) ApiExtractMemoryKVRequest {
	return ApiExtractMemoryKVRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

func (a *ChatExtractMessageApiSubService) ExtractMemoryKVExecute(r ApiExtractMemoryKVRequest) (*ResultRTOMemoryKVDTO, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ResultRTOMemoryKVDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatExtractMessageApiSubService.ExtractMemoryKV")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/extract/kv"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.extractMemoryKVRequest == nil {
		return localVarReturnValue, nil, reportError("extractMemoryKVRequest is required and must be specified")
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

	localVarPostBody = r.extractMemoryKVRequest

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
