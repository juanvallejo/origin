package management

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
	"github.com/satori/go.uuid"
	"net/http"
)

// GroupsClient is the the Azure Management Groups API enables consolidation of multiple
// subscriptions/resources into an organizational hierarchy and centrally
// manage access control, policies, alerting and reporting for those resources.
type GroupsClient struct {
	BaseClient
}

// NewGroupsClient creates an instance of the GroupsClient client.
func NewGroupsClient(groupID uuid.UUID) GroupsClient {
	return NewGroupsClientWithBaseURI(DefaultBaseURI, groupID)
}

// NewGroupsClientWithBaseURI creates an instance of the GroupsClient client.
func NewGroupsClientWithBaseURI(baseURI string, groupID uuid.UUID) GroupsClient {
	return GroupsClient{NewWithBaseURI(baseURI, groupID)}
}

// Get get the details of the management group.
//
// expand is the $expand=children query string parameter allows clients to request inclusion of children in the
// response payload. recurse is the $recurse=true query string parameter allows clients to request inclusion of
// entire hierarchy in the response payload.
func (client GroupsClient) Get(ctx context.Context, expand string, recurse *bool) (result GroupWithHierarchy, err error) {
	req, err := client.GetPreparer(ctx, expand, recurse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "management.GroupsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "management.GroupsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "management.GroupsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client GroupsClient) GetPreparer(ctx context.Context, expand string, recurse *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId": autorest.Encode("path", client.GroupID),
	}

	const APIVersion = "2017-08-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(expand)) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if recurse != nil {
		queryParameters["$recurse"] = autorest.Encode("query", *recurse)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementGroups/{groupId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client GroupsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client GroupsClient) GetResponder(resp *http.Response) (result GroupWithHierarchy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list management groups for the authenticated user.
//
// skiptoken is page continuation token is only used if a previous operation returned a partial result.
// If a previous response contains a nextLink element, the value of the nextLink element will include a token
// parameter that specifies a starting point to use for subsequent calls.
func (client GroupsClient) List(ctx context.Context, skiptoken string) (result GroupListResultPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, skiptoken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "management.GroupsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.glr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "management.GroupsClient", "List", resp, "Failure sending request")
		return
	}

	result.glr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "management.GroupsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client GroupsClient) ListPreparer(ctx context.Context, skiptoken string) (*http.Request, error) {
	const APIVersion = "2017-08-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Management/managementGroups"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client GroupsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client GroupsClient) ListResponder(resp *http.Response) (result GroupListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client GroupsClient) listNextResults(lastResults GroupListResult) (result GroupListResult, err error) {
	req, err := lastResults.groupListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "management.GroupsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "management.GroupsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "management.GroupsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client GroupsClient) ListComplete(ctx context.Context, skiptoken string) (result GroupListResultIterator, err error) {
	result.page, err = client.List(ctx, skiptoken)
	return
}
