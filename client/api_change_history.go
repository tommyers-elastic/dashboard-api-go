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


// ChangeHistoryApiService ChangeHistoryApi service
type ChangeHistoryApiService service

type ChangeHistoryApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest struct {
	ctx context.Context
	ApiService *ChangeHistoryApiService
	organizationId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	t0 *string
	t1 *string
	timespan *float32
	serials *[]string
	productTypes *[]string
	networkIds *[]string
	statuses *[]string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r ChangeHistoryApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest) PerPage(perPage int32) ChangeHistoryApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ChangeHistoryApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest) StartingAfter(startingAfter string) ChangeHistoryApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ChangeHistoryApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest) EndingBefore(endingBefore string) ChangeHistoryApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest {
	r.endingBefore = &endingBefore
	return r
}

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r ChangeHistoryApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest) T0(t0 string) ChangeHistoryApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r ChangeHistoryApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest) T1(t1 string) ChangeHistoryApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
func (r ChangeHistoryApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest) Timespan(timespan float32) ChangeHistoryApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest {
	r.timespan = &timespan
	return r
}

// Optional parameter to filter device availabilities history by device serial numbers
func (r ChangeHistoryApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest) Serials(serials []string) ChangeHistoryApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest {
	r.serials = &serials
	return r
}

// Optional parameter to filter device availabilities history by device product types
func (r ChangeHistoryApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest) ProductTypes(productTypes []string) ChangeHistoryApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest {
	r.productTypes = &productTypes
	return r
}

// Optional parameter to filter device availabilities history by network IDs
func (r ChangeHistoryApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest) NetworkIds(networkIds []string) ChangeHistoryApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest {
	r.networkIds = &networkIds
	return r
}

// Optional parameter to filter device availabilities history by device statuses
func (r ChangeHistoryApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest) Statuses(statuses []string) ChangeHistoryApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest {
	r.statuses = &statuses
	return r
}

func (r ChangeHistoryApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest) Execute() ([]GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationDevicesAvailabilitiesChangeHistoryExecute(r)
}

/*
GetOrganizationDevicesAvailabilitiesChangeHistory List the availability history information for devices in an organization.

List the availability history information for devices in an organization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ChangeHistoryApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest
*/
func (a *ChangeHistoryApiService) GetOrganizationDevicesAvailabilitiesChangeHistory(ctx context.Context, organizationId string) ChangeHistoryApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest {
	return ChangeHistoryApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner
func (a *ChangeHistoryApiService) GetOrganizationDevicesAvailabilitiesChangeHistoryExecute(r ChangeHistoryApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest) ([]GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChangeHistoryApiService.GetOrganizationDevicesAvailabilitiesChangeHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/devices/availabilities/changeHistory"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "perPage", r.perPage, "")
	}
	if r.startingAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startingAfter", r.startingAfter, "")
	}
	if r.endingBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endingBefore", r.endingBefore, "")
	}
	if r.t0 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t0", r.t0, "")
	}
	if r.t1 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t1", r.t1, "")
	}
	if r.timespan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timespan", r.timespan, "")
	}
	if r.serials != nil {
		t := *r.serials
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "serials", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "serials", t, "multi")
		}
	}
	if r.productTypes != nil {
		t := *r.productTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "productTypes", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "productTypes", t, "multi")
		}
	}
	if r.networkIds != nil {
		t := *r.networkIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "networkIds", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "networkIds", t, "multi")
		}
	}
	if r.statuses != nil {
		t := *r.statuses
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "statuses", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "statuses", t, "multi")
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
