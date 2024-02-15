/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// ByBoundaryApiService ByBoundaryApi service
type ByBoundaryApiService service

type ByBoundaryApiGetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRequest struct {
	ctx context.Context
	ApiService *ByBoundaryApiService
	organizationId string
	boundaryIds *[]string
	ranges *[]GetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRangesParameterInner
	duration *int32
	perPage *int32
	boundaryTypes *[]string
}

// A list of boundary ids. The returned cameras will be filtered to only include these ids.
func (r ByBoundaryApiGetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRequest) BoundaryIds(boundaryIds []string) ByBoundaryApiGetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRequest {
	r.boundaryIds = &boundaryIds
	return r
}

// A list of time ranges with intervals
func (r ByBoundaryApiGetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRequest) Ranges(ranges []GetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRangesParameterInner) ByBoundaryApiGetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRequest {
	r.ranges = &ranges
	return r
}

// The minimum time, in seconds, that the person or car remains in the area to be counted. Defaults to boundary configuration or 60.
func (r ByBoundaryApiGetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRequest) Duration(duration int32) ByBoundaryApiGetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRequest {
	r.duration = &duration
	return r
}

// The number of entries per page returned. Acceptable range is 1 - 1000. Defaults to 1000.
func (r ByBoundaryApiGetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRequest) PerPage(perPage int32) ByBoundaryApiGetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRequest {
	r.perPage = &perPage
	return r
}

// The detection types. Defaults to &#39;person&#39;.
func (r ByBoundaryApiGetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRequest) BoundaryTypes(boundaryTypes []string) ByBoundaryApiGetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRequest {
	r.boundaryTypes = &boundaryTypes
	return r
}

func (r ByBoundaryApiGetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRequest) Execute() ([]GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationCameraDetectionsHistoryByBoundaryByIntervalExecute(r)
}

/*
GetOrganizationCameraDetectionsHistoryByBoundaryByInterval Returns analytics data for timespans

Returns analytics data for timespans

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ByBoundaryApiGetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRequest
*/
func (a *ByBoundaryApiService) GetOrganizationCameraDetectionsHistoryByBoundaryByInterval(ctx context.Context, organizationId string) ByBoundaryApiGetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRequest {
	return ByBoundaryApiGetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInner
func (a *ByBoundaryApiService) GetOrganizationCameraDetectionsHistoryByBoundaryByIntervalExecute(r ByBoundaryApiGetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRequest) ([]GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ByBoundaryApiService.GetOrganizationCameraDetectionsHistoryByBoundaryByInterval")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/camera/detections/history/byBoundary/byInterval"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.boundaryIds == nil {
		return localVarReturnValue, nil, reportError("boundaryIds is required and must be specified")
	}
	if r.ranges == nil {
		return localVarReturnValue, nil, reportError("ranges is required and must be specified")
	}

	{
		t := *r.boundaryIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "boundaryIds", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "boundaryIds", t, "multi")
		}
	}
	{
		t := *r.ranges
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "ranges", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "ranges", t, "multi")
		}
	}
	if r.duration != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "duration", r.duration, "")
	}
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "perPage", r.perPage, "")
	}
	if r.boundaryTypes != nil {
		t := *r.boundaryTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "boundaryTypes", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "boundaryTypes", t, "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["meraki_api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Cisco-Meraki-API-Key"] = key
			}
		}
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
