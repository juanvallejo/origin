package eventhub

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
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// Cluster single Event Hubs Cluster resource in List or Get operations.
type Cluster struct {
	autorest.Response `json:"-"`
	// Sku - Properties of the cluster SKU.
	Sku *ClusterSku `json:"sku,omitempty"`
	// ClusterProperties - Event Hubs Cluster properties supplied in responses in List or Get operations.
	*ClusterProperties `json:"properties,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
	// ID - Resource Id
	ID *string `json:"id,omitempty"`
	// Name - Resource name
	Name *string `json:"name,omitempty"`
	// Type - Resource type
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Cluster.
func (c Cluster) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if c.Sku != nil {
		objectMap["sku"] = c.Sku
	}
	if c.ClusterProperties != nil {
		objectMap["properties"] = c.ClusterProperties
	}
	if c.Location != nil {
		objectMap["location"] = c.Location
	}
	if c.Tags != nil {
		objectMap["tags"] = c.Tags
	}
	if c.ID != nil {
		objectMap["id"] = c.ID
	}
	if c.Name != nil {
		objectMap["name"] = c.Name
	}
	if c.Type != nil {
		objectMap["type"] = c.Type
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Cluster struct.
func (c *Cluster) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "sku":
			if v != nil {
				var sku ClusterSku
				err = json.Unmarshal(*v, &sku)
				if err != nil {
					return err
				}
				c.Sku = &sku
			}
		case "properties":
			if v != nil {
				var clusterProperties ClusterProperties
				err = json.Unmarshal(*v, &clusterProperties)
				if err != nil {
					return err
				}
				c.ClusterProperties = &clusterProperties
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				c.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				c.Tags = tags
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				c.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				c.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				c.Type = &typeVar
			}
		}
	}

	return nil
}

// ClusterListResult the response of the List Event Hubs Clusters operation.
type ClusterListResult struct {
	autorest.Response `json:"-"`
	// Value - The Event Hubs Clusters present in the List Event Hubs operation results.
	Value *[]Cluster `json:"value,omitempty"`
	// NextLink - Link to the next set of results. Empty unless the value parameter contains an incomplete list of Event Hubs Clusters.
	NextLink *string `json:"nextLink,omitempty"`
}

// ClusterListResultIterator provides access to a complete listing of Cluster values.
type ClusterListResultIterator struct {
	i    int
	page ClusterListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ClusterListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ClusterListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ClusterListResultIterator) Response() ClusterListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ClusterListResultIterator) Value() Cluster {
	if !iter.page.NotDone() {
		return Cluster{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (clr ClusterListResult) IsEmpty() bool {
	return clr.Value == nil || len(*clr.Value) == 0
}

// clusterListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (clr ClusterListResult) clusterListResultPreparer() (*http.Request, error) {
	if clr.NextLink == nil || len(to.String(clr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(clr.NextLink)))
}

// ClusterListResultPage contains a page of Cluster values.
type ClusterListResultPage struct {
	fn  func(ClusterListResult) (ClusterListResult, error)
	clr ClusterListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ClusterListResultPage) Next() error {
	next, err := page.fn(page.clr)
	if err != nil {
		return err
	}
	page.clr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ClusterListResultPage) NotDone() bool {
	return !page.clr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ClusterListResultPage) Response() ClusterListResult {
	return page.clr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ClusterListResultPage) Values() []Cluster {
	if page.clr.IsEmpty() {
		return nil
	}
	return *page.clr.Value
}

// ClusterProperties event Hubs Cluster properties supplied in responses in List or Get operations.
type ClusterProperties struct {
	// Created - The UTC time when the Event Hubs Cluster was created.
	Created *string `json:"created,omitempty"`
	// Updated - The UTC time when the Event Hubs Cluster was last updated.
	Updated *string `json:"updated,omitempty"`
	// MetricID - The metric ID of the cluster resource. Provided by the service and not modifiable by the user.
	MetricID *string `json:"metricId,omitempty"`
}

// ClusterQuotaConfigurationProperties contains all settings for the cluster.
type ClusterQuotaConfigurationProperties struct {
	autorest.Response `json:"-"`
	// Settings - All possible Cluster settings - a collection of key/value paired settings which apply to quotas and configurations imposed on the cluster.
	Settings map[string]*string `json:"settings"`
}

// MarshalJSON is the custom marshaler for ClusterQuotaConfigurationProperties.
func (cqcp ClusterQuotaConfigurationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if cqcp.Settings != nil {
		objectMap["settings"] = cqcp.Settings
	}
	return json.Marshal(objectMap)
}

// ClusterSku SKU parameters particular to a cluster instance.
type ClusterSku struct {
	// Name - Name of this SKU.
	Name *string `json:"name,omitempty"`
	// Capacity - The quantity of Event Hubs Cluster Capacity Units contained in this cluster.
	Capacity *int32 `json:"capacity,omitempty"`
}

// ClustersPatchFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type ClustersPatchFuture struct {
	azure.Future
	req *http.Request
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future ClustersPatchFuture) Result(client ClustersClient) (c Cluster, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.ClustersPatchFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		return c, azure.NewAsyncOpIncompleteError("eventhub.ClustersPatchFuture")
	}
	if future.PollingMethod() == azure.PollingLocation {
		c, err = client.PatchResponder(future.Response())
		if err != nil {
			err = autorest.NewErrorWithError(err, "eventhub.ClustersPatchFuture", "Result", future.Response(), "Failure responding to request")
		}
		return
	}
	var req *http.Request
	var resp *http.Response
	if future.PollingURL() != "" {
		req, err = http.NewRequest(http.MethodGet, future.PollingURL(), nil)
		if err != nil {
			return
		}
	} else {
		req = autorest.ChangeToGet(future.req)
	}
	resp, err = autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.ClustersPatchFuture", "Result", resp, "Failure sending request")
		return
	}
	c, err = client.PatchResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.ClustersPatchFuture", "Result", resp, "Failure responding to request")
	}
	return
}

// ErrorResponse error response that indicates the service is not able to process the incoming request. The reason
// is provided in the error message.
type ErrorResponse struct {
	// Code - Error code.
	Code *string `json:"code,omitempty"`
	// Message - Error message indicating why the operation failed.
	Message *string `json:"message,omitempty"`
}

// Operation a Event Hub REST API operation
type Operation struct {
	// Name - Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay the object that represents the operation.
type OperationDisplay struct {
	// Provider - Service provider: Microsoft.EventHub
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed: Invoice, etc.
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`
}

// OperationListResult result of the request to list Event Hub operations. It contains a list of operations and a
// URL link to get the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - List of Event Hub operations supported by the Microsoft.EventHub resource provider.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationListResultIterator provides access to a complete listing of Operation values.
type OperationListResultIterator struct {
	i    int
	page OperationListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationListResultIterator) Response() OperationListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListResultIterator) Value() Operation {
	if !iter.page.NotDone() {
		return Operation{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (olr OperationListResult) IsEmpty() bool {
	return olr.Value == nil || len(*olr.Value) == 0
}

// operationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (olr OperationListResult) operationListResultPreparer() (*http.Request, error) {
	if olr.NextLink == nil || len(to.String(olr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(olr.NextLink)))
}

// OperationListResultPage contains a page of Operation values.
type OperationListResultPage struct {
	fn  func(OperationListResult) (OperationListResult, error)
	olr OperationListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListResultPage) Next() error {
	next, err := page.fn(page.olr)
	if err != nil {
		return err
	}
	page.olr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListResultPage) NotDone() bool {
	return !page.olr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationListResultPage) Response() OperationListResult {
	return page.olr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListResultPage) Values() []Operation {
	if page.olr.IsEmpty() {
		return nil
	}
	return *page.olr.Value
}

// Resource the Resource definition
type Resource struct {
	// ID - Resource Id
	ID *string `json:"id,omitempty"`
	// Name - Resource name
	Name *string `json:"name,omitempty"`
	// Type - Resource type
	Type *string `json:"type,omitempty"`
}

// TrackedResource definition of an Azure resource.
type TrackedResource struct {
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
	// ID - Resource Id
	ID *string `json:"id,omitempty"`
	// Name - Resource name
	Name *string `json:"name,omitempty"`
	// Type - Resource type
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for TrackedResource.
func (tr TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if tr.Location != nil {
		objectMap["location"] = tr.Location
	}
	if tr.Tags != nil {
		objectMap["tags"] = tr.Tags
	}
	if tr.ID != nil {
		objectMap["id"] = tr.ID
	}
	if tr.Name != nil {
		objectMap["name"] = tr.Name
	}
	if tr.Type != nil {
		objectMap["type"] = tr.Type
	}
	return json.Marshal(objectMap)
}
