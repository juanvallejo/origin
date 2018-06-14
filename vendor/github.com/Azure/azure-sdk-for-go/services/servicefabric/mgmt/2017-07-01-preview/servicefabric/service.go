package servicefabric

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
	"net/http"
)

// ServiceClient is the azure Service Fabric Resource Provider API Client
type ServiceClient struct {
	BaseClient
}

// NewServiceClient creates an instance of the ServiceClient client.
func NewServiceClient() ServiceClient {
	return NewServiceClientWithBaseURI(DefaultBaseURI)
}

// NewServiceClientWithBaseURI creates an instance of the ServiceClient client.
func NewServiceClientWithBaseURI(baseURI string) ServiceClient {
	return ServiceClient{NewWithBaseURI(baseURI)}
}

// Delete deletes a service resource with the specified name.
//
// subscriptionID is the customer subscription identifier resourceGroupName is the name of the resource group.
// clusterName is the name of the cluster resource applicationName is the name of the application resource.
// serviceName is the name of the service resource in the format of {applicationName}~{serviceName}.
func (client ServiceClient) Delete(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string, applicationName string, serviceName string) (result ServiceDeleteFuture, err error) {
	req, err := client.DeletePreparer(ctx, subscriptionID, resourceGroupName, clusterName, applicationName, serviceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ServiceClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ServiceClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ServiceClient) DeletePreparer(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string, applicationName string, serviceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName":   autorest.Encode("path", applicationName),
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2017-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications/{applicationName}/services/{serviceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceClient) DeleteSender(req *http.Request) (future ServiceDeleteFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ServiceClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get returns a service resource with the specified name.
//
// subscriptionID is the customer subscription identifier resourceGroupName is the name of the resource group.
// clusterName is the name of the cluster resource applicationName is the name of the application resource.
// serviceName is the name of the service resource in the format of {applicationName}~{serviceName}.
func (client ServiceClient) Get(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string, applicationName string, serviceName string) (result ServiceResource, err error) {
	req, err := client.GetPreparer(ctx, subscriptionID, resourceGroupName, clusterName, applicationName, serviceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ServiceClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ServiceClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ServiceClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ServiceClient) GetPreparer(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string, applicationName string, serviceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName":   autorest.Encode("path", applicationName),
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2017-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications/{applicationName}/services/{serviceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ServiceClient) GetResponder(resp *http.Response) (result ServiceResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List returns all service resources in the specified application.
//
// subscriptionID is the customer subscription identifier resourceGroupName is the name of the resource group.
// clusterName is the name of the cluster resource applicationName is the name of the application resource.
func (client ServiceClient) List(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string, applicationName string) (result ServiceResourceList, err error) {
	req, err := client.ListPreparer(ctx, subscriptionID, resourceGroupName, clusterName, applicationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ServiceClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ServiceClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ServiceClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ServiceClient) ListPreparer(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string, applicationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName":   autorest.Encode("path", applicationName),
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2017-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications/{applicationName}/services", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ServiceClient) ListResponder(resp *http.Response) (result ServiceResourceList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Patch updates a service resource with the specified name.
//
// subscriptionID is the customer subscription identifier resourceGroupName is the name of the resource group.
// clusterName is the name of the cluster resource applicationName is the name of the application resource.
// serviceName is the name of the service resource in the format of {applicationName}~{serviceName}. parameters is
// the service resource for patch operations.
func (client ServiceClient) Patch(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string, applicationName string, serviceName string, parameters ServiceResourceUpdate) (result ServicePatchFuture, err error) {
	req, err := client.PatchPreparer(ctx, subscriptionID, resourceGroupName, clusterName, applicationName, serviceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ServiceClient", "Patch", nil, "Failure preparing request")
		return
	}

	result, err = client.PatchSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ServiceClient", "Patch", result.Response(), "Failure sending request")
		return
	}

	return
}

// PatchPreparer prepares the Patch request.
func (client ServiceClient) PatchPreparer(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string, applicationName string, serviceName string, parameters ServiceResourceUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName":   autorest.Encode("path", applicationName),
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2017-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications/{applicationName}/services/{serviceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PatchSender sends the Patch request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceClient) PatchSender(req *http.Request) (future ServicePatchFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// PatchResponder handles the response to the Patch request. The method always
// closes the http.Response Body.
func (client ServiceClient) PatchResponder(resp *http.Response) (result ServiceResourceUpdate, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Put creates or updates a service resource with the specified name.
//
// subscriptionID is the customer subscription identifier resourceGroupName is the name of the resource group.
// clusterName is the name of the cluster resource applicationName is the name of the application resource.
// serviceName is the name of the service resource in the format of {applicationName}~{serviceName}. parameters is
// the service resource.
func (client ServiceClient) Put(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string, applicationName string, serviceName string, parameters ServiceResource) (result ServicePutFuture, err error) {
	req, err := client.PutPreparer(ctx, subscriptionID, resourceGroupName, clusterName, applicationName, serviceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ServiceClient", "Put", nil, "Failure preparing request")
		return
	}

	result, err = client.PutSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ServiceClient", "Put", result.Response(), "Failure sending request")
		return
	}

	return
}

// PutPreparer prepares the Put request.
func (client ServiceClient) PutPreparer(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string, applicationName string, serviceName string, parameters ServiceResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName":   autorest.Encode("path", applicationName),
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2017-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications/{applicationName}/services/{serviceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutSender sends the Put request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceClient) PutSender(req *http.Request) (future ServicePutFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// PutResponder handles the response to the Put request. The method always
// closes the http.Response Body.
func (client ServiceClient) PutResponder(resp *http.Response) (result ServiceResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
