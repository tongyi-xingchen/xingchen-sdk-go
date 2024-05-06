package xingchen

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

type GroupChatApiSubService service

type ApiGroupChatRequest struct {
	ctx            context.Context
	ApiService     *GroupChatApiSubService
	chatReqParams  *BaseChatRequest
	dataInspection *bool
}

func (r ApiGroupChatRequest) ChatReqParams(chatReqParams BaseChatRequest) ApiGroupChatRequest {
	r.chatReqParams = &chatReqParams
	return r
}

func (r ApiGroupChatRequest) DataInspection(dataInspection *bool) ApiGroupChatRequest {
	r.dataInspection = dataInspection
	return r
}

func (r ApiGroupChatRequest) Execute() (*ResultDTOChatResultDTO, *http.Response, error) {
	return r.ApiService.ChatExecute(r)
}

func (r ApiGroupChatRequest) StreamExecute() (*ChatResultStream, error) {
	return r.ApiService.ChatStreamExecute(r)
}

/*
Chat 用户对话

	发起角色对话：


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiChatRequest
*/
func (a *GroupChatApiSubService) Chat(ctx context.Context) ApiGroupChatRequest {
	return ApiGroupChatRequest{
		ApiService:     a,
		ctx:            ctx,
		dataInspection: PtrBool(false),
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *GroupChatApiSubService) ChatExecute(r ApiGroupChatRequest) (*ResultDTOChatResultDTO, *http.Response, error) {
	var (
		localVarReturnValue *ResultDTOChatResultDTO
	)

	localVarHTTPResponse, err := a.call(r)
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

func (a *GroupChatApiSubService) ChatStreamExecute(r ApiGroupChatRequest) (*ChatResultStream, error) {
	localVarHTTPResponse, err := a.call(r)
	if err != nil {
		return nil, err
	}
	chatResultStream := NewChatResultStream(localVarHTTPResponse)
	return chatResultStream, nil

}

func (a *GroupChatApiSubService) call(r ApiGroupChatRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupChatApiSubService.Chat")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/groupchat/send"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.chatReqParams == nil {
		return nil, reportError("chatReqParams is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/event-stream"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.chatReqParams.Streaming != nil && *r.chatReqParams.Streaming {
		localVarHeaderParams["X-AcA-SSE"] = "enable"
		localVarHeaderParams["Accept"] = "text/event-stream"
	}
	if *r.dataInspection {
		localVarHeaderParams["X-AcA-DataInspection"] = "enable"
	}

	localVarPostBody = r.chatReqParams
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)

	return localVarHTTPResponse, err
}
