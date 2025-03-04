package network

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// P2sVpnServerConfigurationsClient is the network Client
type P2sVpnServerConfigurationsClient struct {
	BaseClient
}

// NewP2sVpnServerConfigurationsClient creates an instance of the P2sVpnServerConfigurationsClient client.
func NewP2sVpnServerConfigurationsClient(subscriptionID string) P2sVpnServerConfigurationsClient {
	return NewP2sVpnServerConfigurationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewP2sVpnServerConfigurationsClientWithBaseURI creates an instance of the P2sVpnServerConfigurationsClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewP2sVpnServerConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) P2sVpnServerConfigurationsClient {
	return P2sVpnServerConfigurationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a P2SVpnServerConfiguration to associate with a VirtualWan if it doesn't exist else updates
// the existing P2SVpnServerConfiguration.
// Parameters:
// resourceGroupName - the resource group name of the VirtualWan.
// virtualWanName - the name of the VirtualWan.
// p2SVpnServerConfigurationName - the name of the P2SVpnServerConfiguration.
// p2SVpnServerConfigurationParameters - parameters supplied to create or Update a P2SVpnServerConfiguration.
func (client P2sVpnServerConfigurationsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualWanName string, p2SVpnServerConfigurationName string, p2SVpnServerConfigurationParameters P2SVpnServerConfiguration) (result P2sVpnServerConfigurationsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/P2sVpnServerConfigurationsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, virtualWanName, p2SVpnServerConfigurationName, p2SVpnServerConfigurationParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.P2sVpnServerConfigurationsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.P2sVpnServerConfigurationsClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client P2sVpnServerConfigurationsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, virtualWanName string, p2SVpnServerConfigurationName string, p2SVpnServerConfigurationParameters P2SVpnServerConfiguration) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"p2SVpnServerConfigurationName": autorest.Encode("path", p2SVpnServerConfigurationName),
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
		"virtualWanName":                autorest.Encode("path", virtualWanName),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	p2SVpnServerConfigurationParameters.Etag = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWanName}/p2sVpnServerConfigurations/{p2SVpnServerConfigurationName}", pathParameters),
		autorest.WithJSON(p2SVpnServerConfigurationParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client P2sVpnServerConfigurationsClient) CreateOrUpdateSender(req *http.Request) (future P2sVpnServerConfigurationsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client P2sVpnServerConfigurationsClient) (pvsc P2SVpnServerConfiguration, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.P2sVpnServerConfigurationsCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("network.P2sVpnServerConfigurationsCreateOrUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		pvsc.Response.Response, err = future.GetResult(sender)
		if pvsc.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "network.P2sVpnServerConfigurationsCreateOrUpdateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && pvsc.Response.Response.StatusCode != http.StatusNoContent {
			pvsc, err = client.CreateOrUpdateResponder(pvsc.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "network.P2sVpnServerConfigurationsCreateOrUpdateFuture", "Result", pvsc.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client P2sVpnServerConfigurationsClient) CreateOrUpdateResponder(resp *http.Response) (result P2SVpnServerConfiguration, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a P2SVpnServerConfiguration.
// Parameters:
// resourceGroupName - the resource group name of the P2SVpnServerConfiguration.
// virtualWanName - the name of the VirtualWan.
// p2SVpnServerConfigurationName - the name of the P2SVpnServerConfiguration.
func (client P2sVpnServerConfigurationsClient) Delete(ctx context.Context, resourceGroupName string, virtualWanName string, p2SVpnServerConfigurationName string) (result P2sVpnServerConfigurationsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/P2sVpnServerConfigurationsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, virtualWanName, p2SVpnServerConfigurationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.P2sVpnServerConfigurationsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.P2sVpnServerConfigurationsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client P2sVpnServerConfigurationsClient) DeletePreparer(ctx context.Context, resourceGroupName string, virtualWanName string, p2SVpnServerConfigurationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"p2SVpnServerConfigurationName": autorest.Encode("path", p2SVpnServerConfigurationName),
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
		"virtualWanName":                autorest.Encode("path", virtualWanName),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWanName}/p2sVpnServerConfigurations/{p2SVpnServerConfigurationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client P2sVpnServerConfigurationsClient) DeleteSender(req *http.Request) (future P2sVpnServerConfigurationsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client P2sVpnServerConfigurationsClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.P2sVpnServerConfigurationsDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("network.P2sVpnServerConfigurationsDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client P2sVpnServerConfigurationsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieves the details of a P2SVpnServerConfiguration.
// Parameters:
// resourceGroupName - the resource group name of the P2SVpnServerConfiguration.
// virtualWanName - the name of the VirtualWan.
// p2SVpnServerConfigurationName - the name of the P2SVpnServerConfiguration.
func (client P2sVpnServerConfigurationsClient) Get(ctx context.Context, resourceGroupName string, virtualWanName string, p2SVpnServerConfigurationName string) (result P2SVpnServerConfiguration, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/P2sVpnServerConfigurationsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, virtualWanName, p2SVpnServerConfigurationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.P2sVpnServerConfigurationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.P2sVpnServerConfigurationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.P2sVpnServerConfigurationsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client P2sVpnServerConfigurationsClient) GetPreparer(ctx context.Context, resourceGroupName string, virtualWanName string, p2SVpnServerConfigurationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"p2SVpnServerConfigurationName": autorest.Encode("path", p2SVpnServerConfigurationName),
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
		"virtualWanName":                autorest.Encode("path", virtualWanName),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWanName}/p2sVpnServerConfigurations/{p2SVpnServerConfigurationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client P2sVpnServerConfigurationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client P2sVpnServerConfigurationsClient) GetResponder(resp *http.Response) (result P2SVpnServerConfiguration, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByVirtualWan retrieves all P2SVpnServerConfigurations for a particular VirtualWan.
// Parameters:
// resourceGroupName - the resource group name of the VirtualWan.
// virtualWanName - the name of the VirtualWan.
func (client P2sVpnServerConfigurationsClient) ListByVirtualWan(ctx context.Context, resourceGroupName string, virtualWanName string) (result ListP2SVpnServerConfigurationsResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/P2sVpnServerConfigurationsClient.ListByVirtualWan")
		defer func() {
			sc := -1
			if result.lpvscr.Response.Response != nil {
				sc = result.lpvscr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByVirtualWanNextResults
	req, err := client.ListByVirtualWanPreparer(ctx, resourceGroupName, virtualWanName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.P2sVpnServerConfigurationsClient", "ListByVirtualWan", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByVirtualWanSender(req)
	if err != nil {
		result.lpvscr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.P2sVpnServerConfigurationsClient", "ListByVirtualWan", resp, "Failure sending request")
		return
	}

	result.lpvscr, err = client.ListByVirtualWanResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.P2sVpnServerConfigurationsClient", "ListByVirtualWan", resp, "Failure responding to request")
		return
	}
	if result.lpvscr.hasNextLink() && result.lpvscr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByVirtualWanPreparer prepares the ListByVirtualWan request.
func (client P2sVpnServerConfigurationsClient) ListByVirtualWanPreparer(ctx context.Context, resourceGroupName string, virtualWanName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"virtualWanName":    autorest.Encode("path", virtualWanName),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWanName}/p2sVpnServerConfigurations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByVirtualWanSender sends the ListByVirtualWan request. The method will close the
// http.Response Body if it receives an error.
func (client P2sVpnServerConfigurationsClient) ListByVirtualWanSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByVirtualWanResponder handles the response to the ListByVirtualWan request. The method always
// closes the http.Response Body.
func (client P2sVpnServerConfigurationsClient) ListByVirtualWanResponder(resp *http.Response) (result ListP2SVpnServerConfigurationsResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByVirtualWanNextResults retrieves the next set of results, if any.
func (client P2sVpnServerConfigurationsClient) listByVirtualWanNextResults(ctx context.Context, lastResults ListP2SVpnServerConfigurationsResult) (result ListP2SVpnServerConfigurationsResult, err error) {
	req, err := lastResults.listP2SVpnServerConfigurationsResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.P2sVpnServerConfigurationsClient", "listByVirtualWanNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByVirtualWanSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.P2sVpnServerConfigurationsClient", "listByVirtualWanNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByVirtualWanResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.P2sVpnServerConfigurationsClient", "listByVirtualWanNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByVirtualWanComplete enumerates all values, automatically crossing page boundaries as required.
func (client P2sVpnServerConfigurationsClient) ListByVirtualWanComplete(ctx context.Context, resourceGroupName string, virtualWanName string) (result ListP2SVpnServerConfigurationsResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/P2sVpnServerConfigurationsClient.ListByVirtualWan")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByVirtualWan(ctx, resourceGroupName, virtualWanName)
	return
}
