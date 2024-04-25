package xingchen

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

type CommonApiSubService service

type ApiDocConvertSubmitRequest struct {
	ctx           context.Context
	ApiService    *CommonApiSubService
	docConvertReq *DocConvertSubmitRequest
}

func (r ApiDocConvertSubmitRequest) DocConvertSubmit(docConvertReq DocConvertSubmitRequest) ApiDocConvertSubmitRequest {
	r.docConvertReq = &docConvertReq
	return r
}

func (r ApiDocConvertSubmitRequest) Execute() (*ResultDTODocConvertDTO, *http.Response, error) {
	return r.ApiService.DocConvertSubmitExecute(r)
}

func (a *CommonApiSubService) DocConvertSubmit(ctx context.Context) ApiDocConvertSubmitRequest {
	return ApiDocConvertSubmitRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

func (a *CommonApiSubService) DocConvertSubmitExecute(r ApiDocConvertSubmitRequest) (*ResultDTODocConvertDTO, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ResultDTODocConvertDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommonApiSubService.DocConvertSubmit")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/common/file/asyn/upload"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.docConvertReq == nil {
		return localVarReturnValue, nil, reportError("docConvertReq is required and must be specified")
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

	localVarPostBody = r.docConvertReq

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

type ApiDocConvertResultRequest struct {
	ctx        context.Context
	ApiService *CommonApiSubService
	taskId     *int64
}

func (r ApiDocConvertResultRequest) TaskId(taskId *int64) ApiDocConvertResultRequest {
	r.taskId = taskId
	return r
}

func (r ApiDocConvertResultRequest) Execute() (*ResultDTODocConvertDTO, *http.Response, error) {
	return r.ApiService.DocConvertResultExecute(r)
}

func (a *CommonApiSubService) DocConvertResult(ctx context.Context) ApiDocConvertResultRequest {
	return ApiDocConvertResultRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

func (a *CommonApiSubService) DocConvertResultExecute(r ApiDocConvertResultRequest) (*ResultDTODocConvertDTO, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ResultDTODocConvertDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommonApiSubService.DocConvertResult")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/common/file/asyn/download"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.taskId == nil {
		return localVarReturnValue, nil, reportError("taskId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "taskId", r.taskId, "")
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
