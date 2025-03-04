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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// LocalNetworkGatewaysClient is the network Client
type LocalNetworkGatewaysClient struct {
	BaseClient
}

// NewLocalNetworkGatewaysClient creates an instance of the LocalNetworkGatewaysClient client.
func NewLocalNetworkGatewaysClient(subscriptionID string) LocalNetworkGatewaysClient {
	return NewLocalNetworkGatewaysClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLocalNetworkGatewaysClientWithBaseURI creates an instance of the LocalNetworkGatewaysClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewLocalNetworkGatewaysClientWithBaseURI(baseURI string, subscriptionID string) LocalNetworkGatewaysClient {
	return LocalNetworkGatewaysClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a local network gateway in the specified resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// localNetworkGatewayName - the name of the local network gateway.
// parameters - parameters supplied to the create or update local network gateway operation.
func (client LocalNetworkGatewaysClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, parameters LocalNetworkGateway) (result LocalNetworkGatewaysCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LocalNetworkGatewaysClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: localNetworkGatewayName,
			Constraints: []validation.Constraint{{Target: "localNetworkGatewayName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.LocalNetworkGatewayPropertiesFormat", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("network.LocalNetworkGatewaysClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, localNetworkGatewayName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client LocalNetworkGatewaysClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, parameters LocalNetworkGateway) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"localNetworkGatewayName": autorest.Encode("path", localNetworkGatewayName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client LocalNetworkGatewaysClient) CreateOrUpdateSender(req *http.Request) (future LocalNetworkGatewaysCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client LocalNetworkGatewaysClient) (lng LocalNetworkGateway, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("network.LocalNetworkGatewaysCreateOrUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		lng.Response.Response, err = future.GetResult(sender)
		if lng.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysCreateOrUpdateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && lng.Response.Response.StatusCode != http.StatusNoContent {
			lng, err = client.CreateOrUpdateResponder(lng.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysCreateOrUpdateFuture", "Result", lng.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client LocalNetworkGatewaysClient) CreateOrUpdateResponder(resp *http.Response) (result LocalNetworkGateway, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified local network gateway.
// Parameters:
// resourceGroupName - the name of the resource group.
// localNetworkGatewayName - the name of the local network gateway.
func (client LocalNetworkGatewaysClient) Delete(ctx context.Context, resourceGroupName string, localNetworkGatewayName string) (result LocalNetworkGatewaysDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LocalNetworkGatewaysClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: localNetworkGatewayName,
			Constraints: []validation.Constraint{{Target: "localNetworkGatewayName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("network.LocalNetworkGatewaysClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, localNetworkGatewayName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client LocalNetworkGatewaysClient) DeletePreparer(ctx context.Context, resourceGroupName string, localNetworkGatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"localNetworkGatewayName": autorest.Encode("path", localNetworkGatewayName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client LocalNetworkGatewaysClient) DeleteSender(req *http.Request) (future LocalNetworkGatewaysDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client LocalNetworkGatewaysClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("network.LocalNetworkGatewaysDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client LocalNetworkGatewaysClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified local network gateway in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// localNetworkGatewayName - the name of the local network gateway.
func (client LocalNetworkGatewaysClient) Get(ctx context.Context, resourceGroupName string, localNetworkGatewayName string) (result LocalNetworkGateway, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LocalNetworkGatewaysClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: localNetworkGatewayName,
			Constraints: []validation.Constraint{{Target: "localNetworkGatewayName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("network.LocalNetworkGatewaysClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, localNetworkGatewayName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client LocalNetworkGatewaysClient) GetPreparer(ctx context.Context, resourceGroupName string, localNetworkGatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"localNetworkGatewayName": autorest.Encode("path", localNetworkGatewayName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client LocalNetworkGatewaysClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client LocalNetworkGatewaysClient) GetResponder(resp *http.Response) (result LocalNetworkGateway, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all the local network gateways in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
func (client LocalNetworkGatewaysClient) List(ctx context.Context, resourceGroupName string) (result LocalNetworkGatewayListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LocalNetworkGatewaysClient.List")
		defer func() {
			sc := -1
			if result.lnglr.Response.Response != nil {
				sc = result.lnglr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.lnglr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysClient", "List", resp, "Failure sending request")
		return
	}

	result.lnglr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysClient", "List", resp, "Failure responding to request")
		return
	}
	if result.lnglr.hasNextLink() && result.lnglr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client LocalNetworkGatewaysClient) ListPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client LocalNetworkGatewaysClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client LocalNetworkGatewaysClient) ListResponder(resp *http.Response) (result LocalNetworkGatewayListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client LocalNetworkGatewaysClient) listNextResults(ctx context.Context, lastResults LocalNetworkGatewayListResult) (result LocalNetworkGatewayListResult, err error) {
	req, err := lastResults.localNetworkGatewayListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client LocalNetworkGatewaysClient) ListComplete(ctx context.Context, resourceGroupName string) (result LocalNetworkGatewayListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LocalNetworkGatewaysClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName)
	return
}

// UpdateTags updates a local network gateway tags.
// Parameters:
// resourceGroupName - the name of the resource group.
// localNetworkGatewayName - the name of the local network gateway.
// parameters - parameters supplied to update local network gateway tags.
func (client LocalNetworkGatewaysClient) UpdateTags(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, parameters TagsObject) (result LocalNetworkGatewaysUpdateTagsFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LocalNetworkGatewaysClient.UpdateTags")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: localNetworkGatewayName,
			Constraints: []validation.Constraint{{Target: "localNetworkGatewayName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("network.LocalNetworkGatewaysClient", "UpdateTags", err.Error())
	}

	req, err := client.UpdateTagsPreparer(ctx, resourceGroupName, localNetworkGatewayName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysClient", "UpdateTags", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateTagsSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysClient", "UpdateTags", nil, "Failure sending request")
		return
	}

	return
}

// UpdateTagsPreparer prepares the UpdateTags request.
func (client LocalNetworkGatewaysClient) UpdateTagsPreparer(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, parameters TagsObject) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"localNetworkGatewayName": autorest.Encode("path", localNetworkGatewayName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateTagsSender sends the UpdateTags request. The method will close the
// http.Response Body if it receives an error.
func (client LocalNetworkGatewaysClient) UpdateTagsSender(req *http.Request) (future LocalNetworkGatewaysUpdateTagsFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client LocalNetworkGatewaysClient) (lng LocalNetworkGateway, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysUpdateTagsFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("network.LocalNetworkGatewaysUpdateTagsFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		lng.Response.Response, err = future.GetResult(sender)
		if lng.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysUpdateTagsFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && lng.Response.Response.StatusCode != http.StatusNoContent {
			lng, err = client.UpdateTagsResponder(lng.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysUpdateTagsFuture", "Result", lng.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// UpdateTagsResponder handles the response to the UpdateTags request. The method always
// closes the http.Response Body.
func (client LocalNetworkGatewaysClient) UpdateTagsResponder(resp *http.Response) (result LocalNetworkGateway, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
