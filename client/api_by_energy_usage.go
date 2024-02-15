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
)


// ByEnergyUsageApiService ByEnergyUsageApi service
type ByEnergyUsageApiService service

type ByEnergyUsageApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest struct {
	ctx context.Context
	ApiService *ByEnergyUsageApiService
	organizationId string
	t0 *string
	t1 *string
	timespan *float32
}

// The beginning of the timespan for the data.
func (r ByEnergyUsageApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest) T0(t0 string) ByEnergyUsageApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r ByEnergyUsageApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest) T1(t1 string) ByEnergyUsageApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 25 minutes and be less than or equal to 31 days. The default is 1 day.
func (r ByEnergyUsageApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest) Timespan(timespan float32) ByEnergyUsageApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest {
	r.timespan = &timespan
	return r
}

func (r ByEnergyUsageApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest) Execute() ([]GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationSummaryTopSwitchesByEnergyUsageExecute(r)
}

/*
GetOrganizationSummaryTopSwitchesByEnergyUsage Return metrics for organization's top 10 switches by energy usage over given time range

Return metrics for organization's top 10 switches by energy usage over given time range. Default unit is joules.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ByEnergyUsageApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest
*/
func (a *ByEnergyUsageApiService) GetOrganizationSummaryTopSwitchesByEnergyUsage(ctx context.Context, organizationId string) ByEnergyUsageApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest {
	return ByEnergyUsageApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner
func (a *ByEnergyUsageApiService) GetOrganizationSummaryTopSwitchesByEnergyUsageExecute(r ByEnergyUsageApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest) ([]GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationSummaryTopSwitchesByEnergyUsage200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ByEnergyUsageApiService.GetOrganizationSummaryTopSwitchesByEnergyUsage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/summary/top/switches/byEnergyUsage"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.t0 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t0", r.t0, "")
	}
	if r.t1 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t1", r.t1, "")
	}
	if r.timespan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timespan", r.timespan, "")
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
