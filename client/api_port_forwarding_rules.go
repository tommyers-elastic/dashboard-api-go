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


// PortForwardingRulesApiService PortForwardingRulesApi service
type PortForwardingRulesApiService service

type PortForwardingRulesApiGetDeviceCellularGatewayPortForwardingRulesRequest struct {
	ctx context.Context
	ApiService *PortForwardingRulesApiService
	serial string
}

func (r PortForwardingRulesApiGetDeviceCellularGatewayPortForwardingRulesRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDeviceCellularGatewayPortForwardingRulesExecute(r)
}

/*
GetDeviceCellularGatewayPortForwardingRules Returns the port forwarding rules for a single MG.

Returns the port forwarding rules for a single MG.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial Serial
 @return PortForwardingRulesApiGetDeviceCellularGatewayPortForwardingRulesRequest
*/
func (a *PortForwardingRulesApiService) GetDeviceCellularGatewayPortForwardingRules(ctx context.Context, serial string) PortForwardingRulesApiGetDeviceCellularGatewayPortForwardingRulesRequest {
	return PortForwardingRulesApiGetDeviceCellularGatewayPortForwardingRulesRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *PortForwardingRulesApiService) GetDeviceCellularGatewayPortForwardingRulesExecute(r PortForwardingRulesApiGetDeviceCellularGatewayPortForwardingRulesRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortForwardingRulesApiService.GetDeviceCellularGatewayPortForwardingRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/cellularGateway/portForwardingRules"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterValueToString(r.serial, "serial")), -1)

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

type PortForwardingRulesApiGetNetworkApplianceFirewallPortForwardingRulesRequest struct {
	ctx context.Context
	ApiService *PortForwardingRulesApiService
	networkId string
}

func (r PortForwardingRulesApiGetNetworkApplianceFirewallPortForwardingRulesRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetNetworkApplianceFirewallPortForwardingRulesExecute(r)
}

/*
GetNetworkApplianceFirewallPortForwardingRules Return the port forwarding rules for an MX network

Return the port forwarding rules for an MX network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return PortForwardingRulesApiGetNetworkApplianceFirewallPortForwardingRulesRequest
*/
func (a *PortForwardingRulesApiService) GetNetworkApplianceFirewallPortForwardingRules(ctx context.Context, networkId string) PortForwardingRulesApiGetNetworkApplianceFirewallPortForwardingRulesRequest {
	return PortForwardingRulesApiGetNetworkApplianceFirewallPortForwardingRulesRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *PortForwardingRulesApiService) GetNetworkApplianceFirewallPortForwardingRulesExecute(r PortForwardingRulesApiGetNetworkApplianceFirewallPortForwardingRulesRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortForwardingRulesApiService.GetNetworkApplianceFirewallPortForwardingRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/firewall/portForwardingRules"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

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

type PortForwardingRulesApiUpdateDeviceCellularGatewayPortForwardingRulesRequest struct {
	ctx context.Context
	ApiService *PortForwardingRulesApiService
	serial string
	updateDeviceCellularGatewayPortForwardingRulesRequest *UpdateDeviceCellularGatewayPortForwardingRulesRequest
}

func (r PortForwardingRulesApiUpdateDeviceCellularGatewayPortForwardingRulesRequest) UpdateDeviceCellularGatewayPortForwardingRulesRequest(updateDeviceCellularGatewayPortForwardingRulesRequest UpdateDeviceCellularGatewayPortForwardingRulesRequest) PortForwardingRulesApiUpdateDeviceCellularGatewayPortForwardingRulesRequest {
	r.updateDeviceCellularGatewayPortForwardingRulesRequest = &updateDeviceCellularGatewayPortForwardingRulesRequest
	return r
}

func (r PortForwardingRulesApiUpdateDeviceCellularGatewayPortForwardingRulesRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.UpdateDeviceCellularGatewayPortForwardingRulesExecute(r)
}

/*
UpdateDeviceCellularGatewayPortForwardingRules Updates the port forwarding rules for a single MG.

Updates the port forwarding rules for a single MG.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial Serial
 @return PortForwardingRulesApiUpdateDeviceCellularGatewayPortForwardingRulesRequest
*/
func (a *PortForwardingRulesApiService) UpdateDeviceCellularGatewayPortForwardingRules(ctx context.Context, serial string) PortForwardingRulesApiUpdateDeviceCellularGatewayPortForwardingRulesRequest {
	return PortForwardingRulesApiUpdateDeviceCellularGatewayPortForwardingRulesRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *PortForwardingRulesApiService) UpdateDeviceCellularGatewayPortForwardingRulesExecute(r PortForwardingRulesApiUpdateDeviceCellularGatewayPortForwardingRulesRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortForwardingRulesApiService.UpdateDeviceCellularGatewayPortForwardingRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/cellularGateway/portForwardingRules"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterValueToString(r.serial, "serial")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.updateDeviceCellularGatewayPortForwardingRulesRequest
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

type PortForwardingRulesApiUpdateNetworkApplianceFirewallPortForwardingRulesRequest struct {
	ctx context.Context
	ApiService *PortForwardingRulesApiService
	networkId string
	updateNetworkApplianceFirewallPortForwardingRulesRequest *UpdateNetworkApplianceFirewallPortForwardingRulesRequest
}

func (r PortForwardingRulesApiUpdateNetworkApplianceFirewallPortForwardingRulesRequest) UpdateNetworkApplianceFirewallPortForwardingRulesRequest(updateNetworkApplianceFirewallPortForwardingRulesRequest UpdateNetworkApplianceFirewallPortForwardingRulesRequest) PortForwardingRulesApiUpdateNetworkApplianceFirewallPortForwardingRulesRequest {
	r.updateNetworkApplianceFirewallPortForwardingRulesRequest = &updateNetworkApplianceFirewallPortForwardingRulesRequest
	return r
}

func (r PortForwardingRulesApiUpdateNetworkApplianceFirewallPortForwardingRulesRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.UpdateNetworkApplianceFirewallPortForwardingRulesExecute(r)
}

/*
UpdateNetworkApplianceFirewallPortForwardingRules Update the port forwarding rules for an MX network

Update the port forwarding rules for an MX network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return PortForwardingRulesApiUpdateNetworkApplianceFirewallPortForwardingRulesRequest
*/
func (a *PortForwardingRulesApiService) UpdateNetworkApplianceFirewallPortForwardingRules(ctx context.Context, networkId string) PortForwardingRulesApiUpdateNetworkApplianceFirewallPortForwardingRulesRequest {
	return PortForwardingRulesApiUpdateNetworkApplianceFirewallPortForwardingRulesRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *PortForwardingRulesApiService) UpdateNetworkApplianceFirewallPortForwardingRulesExecute(r PortForwardingRulesApiUpdateNetworkApplianceFirewallPortForwardingRulesRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortForwardingRulesApiService.UpdateNetworkApplianceFirewallPortForwardingRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/firewall/portForwardingRules"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateNetworkApplianceFirewallPortForwardingRulesRequest == nil {
		return localVarReturnValue, nil, reportError("updateNetworkApplianceFirewallPortForwardingRulesRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.updateNetworkApplianceFirewallPortForwardingRulesRequest
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
