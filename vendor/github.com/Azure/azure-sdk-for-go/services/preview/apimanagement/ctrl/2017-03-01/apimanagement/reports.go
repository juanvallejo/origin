package apimanagement

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
	"net/http"
)

// ReportsClient is the client for the Reports methods of the Apimanagement service.
type ReportsClient struct {
	BaseClient
}

// NewReportsClient creates an instance of the ReportsClient client.
func NewReportsClient() ReportsClient {
	return ReportsClient{New()}
}

// ListByAPI lists report records by API.
// Parameters:
// apimBaseURL - the management endpoint of the API Management service, for example
// https://myapimservice.management.azure-api.net.
// filter - the filter to apply on the operation.
// top - number of records to return.
// skip - number of records to skip.
func (client ReportsClient) ListByAPI(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result ReportCollectionPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("apimanagement.ReportsClient", "ListByAPI", err.Error())
	}

	result.fn = client.listByAPINextResults
	req, err := client.ListByAPIPreparer(ctx, apimBaseURL, filter, top, skip)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "ListByAPI", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByAPISender(req)
	if err != nil {
		result.rc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "ListByAPI", resp, "Failure sending request")
		return
	}

	result.rc, err = client.ListByAPIResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "ListByAPI", resp, "Failure responding to request")
	}

	return
}

// ListByAPIPreparer prepares the ListByAPI request.
func (client ReportsClient) ListByAPIPreparer(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"apimBaseUrl": apimBaseURL,
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"$filter":     autorest.Encode("query", filter),
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{apimBaseUrl}", urlParameters),
		autorest.WithPath("/reports/byApi"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByAPISender sends the ListByAPI request. The method will close the
// http.Response Body if it receives an error.
func (client ReportsClient) ListByAPISender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByAPIResponder handles the response to the ListByAPI request. The method always
// closes the http.Response Body.
func (client ReportsClient) ListByAPIResponder(resp *http.Response) (result ReportCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByAPINextResults retrieves the next set of results, if any.
func (client ReportsClient) listByAPINextResults(lastResults ReportCollection) (result ReportCollection, err error) {
	req, err := lastResults.reportCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "listByAPINextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByAPISender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "listByAPINextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByAPIResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "listByAPINextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByAPIComplete enumerates all values, automatically crossing page boundaries as required.
func (client ReportsClient) ListByAPIComplete(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result ReportCollectionIterator, err error) {
	result.page, err = client.ListByAPI(ctx, apimBaseURL, filter, top, skip)
	return
}

// ListByGeo lists report records by GeoGraphy.
// Parameters:
// apimBaseURL - the management endpoint of the API Management service, for example
// https://myapimservice.management.azure-api.net.
// filter - the filter to apply on the operation.
// top - number of records to return.
// skip - number of records to skip.
func (client ReportsClient) ListByGeo(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result ReportCollectionPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("apimanagement.ReportsClient", "ListByGeo", err.Error())
	}

	result.fn = client.listByGeoNextResults
	req, err := client.ListByGeoPreparer(ctx, apimBaseURL, filter, top, skip)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "ListByGeo", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByGeoSender(req)
	if err != nil {
		result.rc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "ListByGeo", resp, "Failure sending request")
		return
	}

	result.rc, err = client.ListByGeoResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "ListByGeo", resp, "Failure responding to request")
	}

	return
}

// ListByGeoPreparer prepares the ListByGeo request.
func (client ReportsClient) ListByGeoPreparer(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"apimBaseUrl": apimBaseURL,
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{apimBaseUrl}", urlParameters),
		autorest.WithPath("/reports/byGeo"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByGeoSender sends the ListByGeo request. The method will close the
// http.Response Body if it receives an error.
func (client ReportsClient) ListByGeoSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByGeoResponder handles the response to the ListByGeo request. The method always
// closes the http.Response Body.
func (client ReportsClient) ListByGeoResponder(resp *http.Response) (result ReportCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByGeoNextResults retrieves the next set of results, if any.
func (client ReportsClient) listByGeoNextResults(lastResults ReportCollection) (result ReportCollection, err error) {
	req, err := lastResults.reportCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "listByGeoNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByGeoSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "listByGeoNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByGeoResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "listByGeoNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByGeoComplete enumerates all values, automatically crossing page boundaries as required.
func (client ReportsClient) ListByGeoComplete(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result ReportCollectionIterator, err error) {
	result.page, err = client.ListByGeo(ctx, apimBaseURL, filter, top, skip)
	return
}

// ListByOperation lists report records by API Operations.
// Parameters:
// apimBaseURL - the management endpoint of the API Management service, for example
// https://myapimservice.management.azure-api.net.
// filter - the filter to apply on the operation.
// top - number of records to return.
// skip - number of records to skip.
func (client ReportsClient) ListByOperation(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result ReportCollectionPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("apimanagement.ReportsClient", "ListByOperation", err.Error())
	}

	result.fn = client.listByOperationNextResults
	req, err := client.ListByOperationPreparer(ctx, apimBaseURL, filter, top, skip)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "ListByOperation", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByOperationSender(req)
	if err != nil {
		result.rc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "ListByOperation", resp, "Failure sending request")
		return
	}

	result.rc, err = client.ListByOperationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "ListByOperation", resp, "Failure responding to request")
	}

	return
}

// ListByOperationPreparer prepares the ListByOperation request.
func (client ReportsClient) ListByOperationPreparer(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"apimBaseUrl": apimBaseURL,
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"$filter":     autorest.Encode("query", filter),
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{apimBaseUrl}", urlParameters),
		autorest.WithPath("/reports/byOperation"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByOperationSender sends the ListByOperation request. The method will close the
// http.Response Body if it receives an error.
func (client ReportsClient) ListByOperationSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByOperationResponder handles the response to the ListByOperation request. The method always
// closes the http.Response Body.
func (client ReportsClient) ListByOperationResponder(resp *http.Response) (result ReportCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByOperationNextResults retrieves the next set of results, if any.
func (client ReportsClient) listByOperationNextResults(lastResults ReportCollection) (result ReportCollection, err error) {
	req, err := lastResults.reportCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "listByOperationNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByOperationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "listByOperationNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByOperationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "listByOperationNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByOperationComplete enumerates all values, automatically crossing page boundaries as required.
func (client ReportsClient) ListByOperationComplete(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result ReportCollectionIterator, err error) {
	result.page, err = client.ListByOperation(ctx, apimBaseURL, filter, top, skip)
	return
}

// ListByProduct lists report records by Product.
// Parameters:
// apimBaseURL - the management endpoint of the API Management service, for example
// https://myapimservice.management.azure-api.net.
// filter - the filter to apply on the operation.
// top - number of records to return.
// skip - number of records to skip.
func (client ReportsClient) ListByProduct(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result ReportCollectionPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("apimanagement.ReportsClient", "ListByProduct", err.Error())
	}

	result.fn = client.listByProductNextResults
	req, err := client.ListByProductPreparer(ctx, apimBaseURL, filter, top, skip)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "ListByProduct", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByProductSender(req)
	if err != nil {
		result.rc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "ListByProduct", resp, "Failure sending request")
		return
	}

	result.rc, err = client.ListByProductResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "ListByProduct", resp, "Failure responding to request")
	}

	return
}

// ListByProductPreparer prepares the ListByProduct request.
func (client ReportsClient) ListByProductPreparer(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"apimBaseUrl": apimBaseURL,
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"$filter":     autorest.Encode("query", filter),
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{apimBaseUrl}", urlParameters),
		autorest.WithPath("/reports/byProduct"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByProductSender sends the ListByProduct request. The method will close the
// http.Response Body if it receives an error.
func (client ReportsClient) ListByProductSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByProductResponder handles the response to the ListByProduct request. The method always
// closes the http.Response Body.
func (client ReportsClient) ListByProductResponder(resp *http.Response) (result ReportCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByProductNextResults retrieves the next set of results, if any.
func (client ReportsClient) listByProductNextResults(lastResults ReportCollection) (result ReportCollection, err error) {
	req, err := lastResults.reportCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "listByProductNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByProductSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "listByProductNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByProductResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "listByProductNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByProductComplete enumerates all values, automatically crossing page boundaries as required.
func (client ReportsClient) ListByProductComplete(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result ReportCollectionIterator, err error) {
	result.page, err = client.ListByProduct(ctx, apimBaseURL, filter, top, skip)
	return
}

// ListByRequest lists report records by Request.
// Parameters:
// apimBaseURL - the management endpoint of the API Management service, for example
// https://myapimservice.management.azure-api.net.
// filter - the filter to apply on the operation.
// top - number of records to return.
// skip - number of records to skip.
func (client ReportsClient) ListByRequest(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result RequestReportCollection, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("apimanagement.ReportsClient", "ListByRequest", err.Error())
	}

	req, err := client.ListByRequestPreparer(ctx, apimBaseURL, filter, top, skip)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "ListByRequest", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByRequestSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "ListByRequest", resp, "Failure sending request")
		return
	}

	result, err = client.ListByRequestResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "ListByRequest", resp, "Failure responding to request")
	}

	return
}

// ListByRequestPreparer prepares the ListByRequest request.
func (client ReportsClient) ListByRequestPreparer(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"apimBaseUrl": apimBaseURL,
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"$filter":     autorest.Encode("query", filter),
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{apimBaseUrl}", urlParameters),
		autorest.WithPath("/reports/byRequest"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByRequestSender sends the ListByRequest request. The method will close the
// http.Response Body if it receives an error.
func (client ReportsClient) ListByRequestSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByRequestResponder handles the response to the ListByRequest request. The method always
// closes the http.Response Body.
func (client ReportsClient) ListByRequestResponder(resp *http.Response) (result RequestReportCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBySubscription lists report records by subscription.
// Parameters:
// apimBaseURL - the management endpoint of the API Management service, for example
// https://myapimservice.management.azure-api.net.
// filter - the filter to apply on the operation.
// top - number of records to return.
// skip - number of records to skip.
func (client ReportsClient) ListBySubscription(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result ReportCollectionPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("apimanagement.ReportsClient", "ListBySubscription", err.Error())
	}

	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx, apimBaseURL, filter, top, skip)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.rc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.rc, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "ListBySubscription", resp, "Failure responding to request")
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client ReportsClient) ListBySubscriptionPreparer(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"apimBaseUrl": apimBaseURL,
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{apimBaseUrl}", urlParameters),
		autorest.WithPath("/reports/bySubscription"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client ReportsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client ReportsClient) ListBySubscriptionResponder(resp *http.Response) (result ReportCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client ReportsClient) listBySubscriptionNextResults(lastResults ReportCollection) (result ReportCollection, err error) {
	req, err := lastResults.reportCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client ReportsClient) ListBySubscriptionComplete(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result ReportCollectionIterator, err error) {
	result.page, err = client.ListBySubscription(ctx, apimBaseURL, filter, top, skip)
	return
}

// ListByTime lists report records by Time.
// Parameters:
// apimBaseURL - the management endpoint of the API Management service, for example
// https://myapimservice.management.azure-api.net.
// interval - by time interval. Interval must be multiple of 15 minutes and may not be zero. The value should
// be in ISO  8601 format (http://en.wikipedia.org/wiki/ISO_8601#Durations).This code can be used to convert
// TimeSpan to a valid interval string: XmlConvert.ToString(new TimeSpan(hours, minutes, secconds))
// filter - the filter to apply on the operation.
// top - number of records to return.
// skip - number of records to skip.
func (client ReportsClient) ListByTime(ctx context.Context, apimBaseURL string, interval string, filter string, top *int32, skip *int32) (result ReportCollectionPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("apimanagement.ReportsClient", "ListByTime", err.Error())
	}

	result.fn = client.listByTimeNextResults
	req, err := client.ListByTimePreparer(ctx, apimBaseURL, interval, filter, top, skip)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "ListByTime", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByTimeSender(req)
	if err != nil {
		result.rc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "ListByTime", resp, "Failure sending request")
		return
	}

	result.rc, err = client.ListByTimeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "ListByTime", resp, "Failure responding to request")
	}

	return
}

// ListByTimePreparer prepares the ListByTime request.
func (client ReportsClient) ListByTimePreparer(ctx context.Context, apimBaseURL string, interval string, filter string, top *int32, skip *int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"apimBaseUrl": apimBaseURL,
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
		"interval":    autorest.Encode("query", interval),
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{apimBaseUrl}", urlParameters),
		autorest.WithPath("/reports/byTime"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByTimeSender sends the ListByTime request. The method will close the
// http.Response Body if it receives an error.
func (client ReportsClient) ListByTimeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByTimeResponder handles the response to the ListByTime request. The method always
// closes the http.Response Body.
func (client ReportsClient) ListByTimeResponder(resp *http.Response) (result ReportCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByTimeNextResults retrieves the next set of results, if any.
func (client ReportsClient) listByTimeNextResults(lastResults ReportCollection) (result ReportCollection, err error) {
	req, err := lastResults.reportCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "listByTimeNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByTimeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "listByTimeNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByTimeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "listByTimeNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByTimeComplete enumerates all values, automatically crossing page boundaries as required.
func (client ReportsClient) ListByTimeComplete(ctx context.Context, apimBaseURL string, interval string, filter string, top *int32, skip *int32) (result ReportCollectionIterator, err error) {
	result.page, err = client.ListByTime(ctx, apimBaseURL, interval, filter, top, skip)
	return
}

// ListByUser lists report records by User.
// Parameters:
// apimBaseURL - the management endpoint of the API Management service, for example
// https://myapimservice.management.azure-api.net.
// filter - the filter to apply on the operation.
// top - number of records to return.
// skip - number of records to skip.
func (client ReportsClient) ListByUser(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result ReportCollectionPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("apimanagement.ReportsClient", "ListByUser", err.Error())
	}

	result.fn = client.listByUserNextResults
	req, err := client.ListByUserPreparer(ctx, apimBaseURL, filter, top, skip)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "ListByUser", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByUserSender(req)
	if err != nil {
		result.rc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "ListByUser", resp, "Failure sending request")
		return
	}

	result.rc, err = client.ListByUserResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "ListByUser", resp, "Failure responding to request")
	}

	return
}

// ListByUserPreparer prepares the ListByUser request.
func (client ReportsClient) ListByUserPreparer(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"apimBaseUrl": apimBaseURL,
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"$filter":     autorest.Encode("query", filter),
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{apimBaseUrl}", urlParameters),
		autorest.WithPath("/reports/byUser"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByUserSender sends the ListByUser request. The method will close the
// http.Response Body if it receives an error.
func (client ReportsClient) ListByUserSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByUserResponder handles the response to the ListByUser request. The method always
// closes the http.Response Body.
func (client ReportsClient) ListByUserResponder(resp *http.Response) (result ReportCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByUserNextResults retrieves the next set of results, if any.
func (client ReportsClient) listByUserNextResults(lastResults ReportCollection) (result ReportCollection, err error) {
	req, err := lastResults.reportCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "listByUserNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByUserSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "listByUserNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByUserResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "listByUserNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByUserComplete enumerates all values, automatically crossing page boundaries as required.
func (client ReportsClient) ListByUserComplete(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result ReportCollectionIterator, err error) {
	result.page, err = client.ListByUser(ctx, apimBaseURL, filter, top, skip)
	return
}
