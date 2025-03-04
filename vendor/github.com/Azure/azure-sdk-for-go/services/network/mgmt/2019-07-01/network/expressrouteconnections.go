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

// ExpressRouteConnectionsClient is the network Client
type ExpressRouteConnectionsClient struct {
	BaseClient
}

// NewExpressRouteConnectionsClient creates an instance of the ExpressRouteConnectionsClient client.
func NewExpressRouteConnectionsClient(subscriptionID string) ExpressRouteConnectionsClient {
	return NewExpressRouteConnectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewExpressRouteConnectionsClientWithBaseURI creates an instance of the ExpressRouteConnectionsClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewExpressRouteConnectionsClientWithBaseURI(baseURI string, subscriptionID string) ExpressRouteConnectionsClient {
	return ExpressRouteConnectionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a connection between an ExpressRoute gateway and an ExpressRoute circuit.
// Parameters:
// resourceGroupName - the name of the resource group.
// expressRouteGatewayName - the name of the ExpressRoute gateway.
// connectionName - the name of the connection subresource.
// putExpressRouteConnectionParameters - parameters required in an ExpressRouteConnection PUT operation.
func (client ExpressRouteConnectionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string, putExpressRouteConnectionParameters ExpressRouteConnection) (result ExpressRouteConnectionsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteConnectionsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: putExpressRouteConnectionParameters,
			Constraints: []validation.Constraint{{Target: "putExpressRouteConnectionParameters.ExpressRouteConnectionProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "putExpressRouteConnectionParameters.ExpressRouteConnectionProperties.ExpressRouteCircuitPeering", Name: validation.Null, Rule: true, Chain: nil}}},
				{Target: "putExpressRouteConnectionParameters.Name", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("network.ExpressRouteConnectionsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, expressRouteGatewayName, connectionName, putExpressRouteConnectionParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteConnectionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteConnectionsClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ExpressRouteConnectionsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string, putExpressRouteConnectionParameters ExpressRouteConnection) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectionName":          autorest.Encode("path", connectionName),
		"expressRouteGatewayName": autorest.Encode("path", expressRouteGatewayName),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}/expressRouteConnections/{connectionName}", pathParameters),
		autorest.WithJSON(putExpressRouteConnectionParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteConnectionsClient) CreateOrUpdateSender(req *http.Request) (future ExpressRouteConnectionsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ExpressRouteConnectionsClient) (erc ExpressRouteConnection, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.ExpressRouteConnectionsCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("network.ExpressRouteConnectionsCreateOrUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		erc.Response.Response, err = future.GetResult(sender)
		if erc.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "network.ExpressRouteConnectionsCreateOrUpdateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && erc.Response.Response.StatusCode != http.StatusNoContent {
			erc, err = client.CreateOrUpdateResponder(erc.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "network.ExpressRouteConnectionsCreateOrUpdateFuture", "Result", erc.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ExpressRouteConnectionsClient) CreateOrUpdateResponder(resp *http.Response) (result ExpressRouteConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a connection to a ExpressRoute circuit.
// Parameters:
// resourceGroupName - the name of the resource group.
// expressRouteGatewayName - the name of the ExpressRoute gateway.
// connectionName - the name of the connection subresource.
func (client ExpressRouteConnectionsClient) Delete(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string) (result ExpressRouteConnectionsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteConnectionsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, expressRouteGatewayName, connectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteConnectionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteConnectionsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ExpressRouteConnectionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectionName":          autorest.Encode("path", connectionName),
		"expressRouteGatewayName": autorest.Encode("path", expressRouteGatewayName),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}/expressRouteConnections/{connectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteConnectionsClient) DeleteSender(req *http.Request) (future ExpressRouteConnectionsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ExpressRouteConnectionsClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.ExpressRouteConnectionsDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("network.ExpressRouteConnectionsDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ExpressRouteConnectionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified ExpressRouteConnection.
// Parameters:
// resourceGroupName - the name of the resource group.
// expressRouteGatewayName - the name of the ExpressRoute gateway.
// connectionName - the name of the ExpressRoute connection.
func (client ExpressRouteConnectionsClient) Get(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string) (result ExpressRouteConnection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteConnectionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, expressRouteGatewayName, connectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteConnectionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ExpressRouteConnectionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteConnectionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ExpressRouteConnectionsClient) GetPreparer(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, connectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectionName":          autorest.Encode("path", connectionName),
		"expressRouteGatewayName": autorest.Encode("path", expressRouteGatewayName),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}/expressRouteConnections/{connectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteConnectionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ExpressRouteConnectionsClient) GetResponder(resp *http.Response) (result ExpressRouteConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists ExpressRouteConnections.
// Parameters:
// resourceGroupName - the name of the resource group.
// expressRouteGatewayName - the name of the ExpressRoute gateway.
func (client ExpressRouteConnectionsClient) List(ctx context.Context, resourceGroupName string, expressRouteGatewayName string) (result ExpressRouteConnectionList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteConnectionsClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, resourceGroupName, expressRouteGatewayName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteConnectionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ExpressRouteConnectionsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteConnectionsClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ExpressRouteConnectionsClient) ListPreparer(ctx context.Context, resourceGroupName string, expressRouteGatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"expressRouteGatewayName": autorest.Encode("path", expressRouteGatewayName),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}/expressRouteConnections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteConnectionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ExpressRouteConnectionsClient) ListResponder(resp *http.Response) (result ExpressRouteConnectionList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
