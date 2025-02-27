package extendedlocation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/extendedlocation/mgmt/2021-03-15-preview/extendedlocation"

// AzureEntityResource the resource model definition for an Azure Resource Manager resource with an etag.
type AzureEntityResource struct {
	// Etag - READ-ONLY; Resource Etag.
	Etag *string `json:"etag,omitempty"`
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// CustomLocation custom Locations definition.
type CustomLocation struct {
	autorest.Response `json:"-"`
	// CustomLocationProperties - The set of properties specific to a Custom Location
	*CustomLocationProperties `json:"properties,omitempty"`
	// SystemData - READ-ONLY; Metadata pertaining to creation and last modification of the resource
	SystemData *SystemData `json:"systemData,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for CustomLocation.
func (cl CustomLocation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if cl.CustomLocationProperties != nil {
		objectMap["properties"] = cl.CustomLocationProperties
	}
	if cl.Tags != nil {
		objectMap["tags"] = cl.Tags
	}
	if cl.Location != nil {
		objectMap["location"] = cl.Location
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for CustomLocation struct.
func (cl *CustomLocation) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var customLocationProperties CustomLocationProperties
				err = json.Unmarshal(*v, &customLocationProperties)
				if err != nil {
					return err
				}
				cl.CustomLocationProperties = &customLocationProperties
			}
		case "systemData":
			if v != nil {
				var systemData SystemData
				err = json.Unmarshal(*v, &systemData)
				if err != nil {
					return err
				}
				cl.SystemData = &systemData
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				cl.Tags = tags
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				cl.Location = &location
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				cl.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				cl.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				cl.Type = &typeVar
			}
		}
	}

	return nil
}

// CustomLocationListResult the List Custom Locations operation response.
type CustomLocationListResult struct {
	autorest.Response `json:"-"`
	// NextLink - READ-ONLY; The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
	// Value - READ-ONLY; The list of Custom Locations.
	Value *[]CustomLocation `json:"value,omitempty"`
}

// CustomLocationListResultIterator provides access to a complete listing of CustomLocation values.
type CustomLocationListResultIterator struct {
	i    int
	page CustomLocationListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *CustomLocationListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomLocationListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *CustomLocationListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter CustomLocationListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter CustomLocationListResultIterator) Response() CustomLocationListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter CustomLocationListResultIterator) Value() CustomLocation {
	if !iter.page.NotDone() {
		return CustomLocation{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the CustomLocationListResultIterator type.
func NewCustomLocationListResultIterator(page CustomLocationListResultPage) CustomLocationListResultIterator {
	return CustomLocationListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (cllr CustomLocationListResult) IsEmpty() bool {
	return cllr.Value == nil || len(*cllr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (cllr CustomLocationListResult) hasNextLink() bool {
	return cllr.NextLink != nil && len(*cllr.NextLink) != 0
}

// customLocationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (cllr CustomLocationListResult) customLocationListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !cllr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(cllr.NextLink)))
}

// CustomLocationListResultPage contains a page of CustomLocation values.
type CustomLocationListResultPage struct {
	fn   func(context.Context, CustomLocationListResult) (CustomLocationListResult, error)
	cllr CustomLocationListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *CustomLocationListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomLocationListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.cllr)
		if err != nil {
			return err
		}
		page.cllr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *CustomLocationListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page CustomLocationListResultPage) NotDone() bool {
	return !page.cllr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page CustomLocationListResultPage) Response() CustomLocationListResult {
	return page.cllr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page CustomLocationListResultPage) Values() []CustomLocation {
	if page.cllr.IsEmpty() {
		return nil
	}
	return *page.cllr.Value
}

// Creates a new instance of the CustomLocationListResultPage type.
func NewCustomLocationListResultPage(cur CustomLocationListResult, getNextPage func(context.Context, CustomLocationListResult) (CustomLocationListResult, error)) CustomLocationListResultPage {
	return CustomLocationListResultPage{
		fn:   getNextPage,
		cllr: cur,
	}
}

// CustomLocationOperation custom Locations operation.
type CustomLocationOperation struct {
	// CustomLocationOperationValueDisplay - Describes the properties of a Custom Locations Operation Value Display.
	*CustomLocationOperationValueDisplay `json:"display,omitempty"`
	// IsDataAction - READ-ONLY; Is this Operation a data plane operation
	IsDataAction *bool `json:"isDataAction,omitempty"`
	// Name - READ-ONLY; The name of the compute operation.
	Name *string `json:"name,omitempty"`
	// Origin - READ-ONLY; The origin of the compute operation.
	Origin *string `json:"origin,omitempty"`
}

// MarshalJSON is the custom marshaler for CustomLocationOperation.
func (clo CustomLocationOperation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if clo.CustomLocationOperationValueDisplay != nil {
		objectMap["display"] = clo.CustomLocationOperationValueDisplay
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for CustomLocationOperation struct.
func (clo *CustomLocationOperation) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "display":
			if v != nil {
				var customLocationOperationValueDisplay CustomLocationOperationValueDisplay
				err = json.Unmarshal(*v, &customLocationOperationValueDisplay)
				if err != nil {
					return err
				}
				clo.CustomLocationOperationValueDisplay = &customLocationOperationValueDisplay
			}
		case "isDataAction":
			if v != nil {
				var isDataAction bool
				err = json.Unmarshal(*v, &isDataAction)
				if err != nil {
					return err
				}
				clo.IsDataAction = &isDataAction
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				clo.Name = &name
			}
		case "origin":
			if v != nil {
				var origin string
				err = json.Unmarshal(*v, &origin)
				if err != nil {
					return err
				}
				clo.Origin = &origin
			}
		}
	}

	return nil
}

// CustomLocationOperationsList lists of Custom Locations operations.
type CustomLocationOperationsList struct {
	autorest.Response `json:"-"`
	// NextLink - Next page of operations.
	NextLink *string `json:"nextLink,omitempty"`
	// Value - Array of customLocationOperation
	Value *[]CustomLocationOperation `json:"value,omitempty"`
}

// CustomLocationOperationsListIterator provides access to a complete listing of CustomLocationOperation
// values.
type CustomLocationOperationsListIterator struct {
	i    int
	page CustomLocationOperationsListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *CustomLocationOperationsListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomLocationOperationsListIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *CustomLocationOperationsListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter CustomLocationOperationsListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter CustomLocationOperationsListIterator) Response() CustomLocationOperationsList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter CustomLocationOperationsListIterator) Value() CustomLocationOperation {
	if !iter.page.NotDone() {
		return CustomLocationOperation{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the CustomLocationOperationsListIterator type.
func NewCustomLocationOperationsListIterator(page CustomLocationOperationsListPage) CustomLocationOperationsListIterator {
	return CustomLocationOperationsListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (clol CustomLocationOperationsList) IsEmpty() bool {
	return clol.Value == nil || len(*clol.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (clol CustomLocationOperationsList) hasNextLink() bool {
	return clol.NextLink != nil && len(*clol.NextLink) != 0
}

// customLocationOperationsListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (clol CustomLocationOperationsList) customLocationOperationsListPreparer(ctx context.Context) (*http.Request, error) {
	if !clol.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(clol.NextLink)))
}

// CustomLocationOperationsListPage contains a page of CustomLocationOperation values.
type CustomLocationOperationsListPage struct {
	fn   func(context.Context, CustomLocationOperationsList) (CustomLocationOperationsList, error)
	clol CustomLocationOperationsList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *CustomLocationOperationsListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomLocationOperationsListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.clol)
		if err != nil {
			return err
		}
		page.clol = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *CustomLocationOperationsListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page CustomLocationOperationsListPage) NotDone() bool {
	return !page.clol.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page CustomLocationOperationsListPage) Response() CustomLocationOperationsList {
	return page.clol
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page CustomLocationOperationsListPage) Values() []CustomLocationOperation {
	if page.clol.IsEmpty() {
		return nil
	}
	return *page.clol.Value
}

// Creates a new instance of the CustomLocationOperationsListPage type.
func NewCustomLocationOperationsListPage(cur CustomLocationOperationsList, getNextPage func(context.Context, CustomLocationOperationsList) (CustomLocationOperationsList, error)) CustomLocationOperationsListPage {
	return CustomLocationOperationsListPage{
		fn:   getNextPage,
		clol: cur,
	}
}

// CustomLocationOperationValueDisplay describes the properties of a Custom Locations Operation Value
// Display.
type CustomLocationOperationValueDisplay struct {
	// Description - READ-ONLY; The description of the operation.
	Description *string `json:"description,omitempty"`
	// Operation - READ-ONLY; The display name of the compute operation.
	Operation *string `json:"operation,omitempty"`
	// Provider - READ-ONLY; The resource provider for the operation.
	Provider *string `json:"provider,omitempty"`
	// Resource - READ-ONLY; The display name of the resource the operation applies to.
	Resource *string `json:"resource,omitempty"`
}

// CustomLocationProperties properties for a custom location.
type CustomLocationProperties struct {
	// Authentication - This is optional input that contains the authentication that should be used to generate the namespace.
	Authentication *CustomLocationPropertiesAuthentication `json:"authentication,omitempty"`
	// ClusterExtensionIds - Contains the reference to the add-on that contains charts to deploy CRDs and operators.
	ClusterExtensionIds *[]string `json:"clusterExtensionIds,omitempty"`
	// DisplayName - Display name for the Custom Locations location.
	DisplayName *string `json:"displayName,omitempty"`
	// HostResourceID - Connected Cluster or AKS Cluster. The Custom Locations RP will perform a checkAccess API for listAdminCredentials permissions.
	HostResourceID *string `json:"hostResourceId,omitempty"`
	// HostType - Type of host the Custom Locations is referencing (Kubernetes, etc...). Possible values include: 'Kubernetes'
	HostType HostType `json:"hostType,omitempty"`
	// Namespace - Kubernetes namespace that will be created on the specified cluster.
	Namespace *string `json:"namespace,omitempty"`
	// ProvisioningState - Provisioning State for the Custom Location.
	ProvisioningState *string `json:"provisioningState,omitempty"`
}

// CustomLocationPropertiesAuthentication this is optional input that contains the authentication that
// should be used to generate the namespace.
type CustomLocationPropertiesAuthentication struct {
	// Type - The type of the Custom Locations authentication
	Type *string `json:"type,omitempty"`
	// Value - The kubeconfig value.
	Value *string `json:"value,omitempty"`
}

// CustomLocationsCreateOrUpdateFuture an abstraction for monitoring and retrieving the results of a
// long-running operation.
type CustomLocationsCreateOrUpdateFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(CustomLocationsClient) (CustomLocation, error)
}

// UnmarshalJSON is the custom unmarshaller for CreateFuture.
func (future *CustomLocationsCreateOrUpdateFuture) UnmarshalJSON(body []byte) error {
	var azFuture azure.Future
	if err := json.Unmarshal(body, &azFuture); err != nil {
		return err
	}
	future.FutureAPI = &azFuture
	future.Result = future.result
	return nil
}

// result is the default implementation for CustomLocationsCreateOrUpdateFuture.Result.
func (future *CustomLocationsCreateOrUpdateFuture) result(client CustomLocationsClient) (cl CustomLocation, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "extendedlocation.CustomLocationsCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		cl.Response.Response = future.Response()
		err = azure.NewAsyncOpIncompleteError("extendedlocation.CustomLocationsCreateOrUpdateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if cl.Response.Response, err = future.GetResult(sender); err == nil && cl.Response.Response.StatusCode != http.StatusNoContent {
		cl, err = client.CreateOrUpdateResponder(cl.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "extendedlocation.CustomLocationsCreateOrUpdateFuture", "Result", cl.Response.Response, "Failure responding to request")
		}
	}
	return
}

// CustomLocationsDeleteFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type CustomLocationsDeleteFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(CustomLocationsClient) (autorest.Response, error)
}

// UnmarshalJSON is the custom unmarshaller for CreateFuture.
func (future *CustomLocationsDeleteFuture) UnmarshalJSON(body []byte) error {
	var azFuture azure.Future
	if err := json.Unmarshal(body, &azFuture); err != nil {
		return err
	}
	future.FutureAPI = &azFuture
	future.Result = future.result
	return nil
}

// result is the default implementation for CustomLocationsDeleteFuture.Result.
func (future *CustomLocationsDeleteFuture) result(client CustomLocationsClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "extendedlocation.CustomLocationsDeleteFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		ar.Response = future.Response()
		err = azure.NewAsyncOpIncompleteError("extendedlocation.CustomLocationsDeleteFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// EnabledResourceType enabledResourceType definition.
type EnabledResourceType struct {
	// EnabledResourceTypeProperties - The set of properties for EnabledResourceType specific to a Custom Location
	*EnabledResourceTypeProperties `json:"properties,omitempty"`
	// SystemData - READ-ONLY; Metadata pertaining to creation and last modification of the resource
	SystemData *SystemData `json:"systemData,omitempty"`
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for EnabledResourceType.
func (ert EnabledResourceType) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ert.EnabledResourceTypeProperties != nil {
		objectMap["properties"] = ert.EnabledResourceTypeProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for EnabledResourceType struct.
func (ert *EnabledResourceType) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var enabledResourceTypeProperties EnabledResourceTypeProperties
				err = json.Unmarshal(*v, &enabledResourceTypeProperties)
				if err != nil {
					return err
				}
				ert.EnabledResourceTypeProperties = &enabledResourceTypeProperties
			}
		case "systemData":
			if v != nil {
				var systemData SystemData
				err = json.Unmarshal(*v, &systemData)
				if err != nil {
					return err
				}
				ert.SystemData = &systemData
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				ert.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				ert.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				ert.Type = &typeVar
			}
		}
	}

	return nil
}

// EnabledResourceTypeProperties properties for EnabledResourceType of a custom location.
type EnabledResourceTypeProperties struct {
	// ClusterExtensionID - Cluster Extension ID
	ClusterExtensionID *string `json:"clusterExtensionId,omitempty"`
	// ExtensionType - Cluster Extension Type
	ExtensionType *string `json:"extensionType,omitempty"`
	// TypesMetadata - Metadata of the Resource Type
	TypesMetadata *[]EnabledResourceTypePropertiesTypesMetadataItem `json:"typesMetadata,omitempty"`
}

// EnabledResourceTypePropertiesTypesMetadataItem metadata of the Resource Type.
type EnabledResourceTypePropertiesTypesMetadataItem struct {
	// APIVersion - Api Version of Resource Type
	APIVersion *string `json:"apiVersion,omitempty"`
	// ResourceProviderNamespace - Resource Provider Namespace of Resource Type
	ResourceProviderNamespace *string `json:"resourceProviderNamespace,omitempty"`
	// ResourceType - Resource Type
	ResourceType *string `json:"resourceType,omitempty"`
}

// EnabledResourceTypesListResult list of EnabledResourceTypes definition.
type EnabledResourceTypesListResult struct {
	autorest.Response `json:"-"`
	// NextLink - READ-ONLY; The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
	// Value - READ-ONLY; The list of EnabledResourceTypes available for a customLocation.
	Value *[]EnabledResourceType `json:"value,omitempty"`
}

// EnabledResourceTypesListResultIterator provides access to a complete listing of EnabledResourceType
// values.
type EnabledResourceTypesListResultIterator struct {
	i    int
	page EnabledResourceTypesListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *EnabledResourceTypesListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnabledResourceTypesListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *EnabledResourceTypesListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter EnabledResourceTypesListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter EnabledResourceTypesListResultIterator) Response() EnabledResourceTypesListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter EnabledResourceTypesListResultIterator) Value() EnabledResourceType {
	if !iter.page.NotDone() {
		return EnabledResourceType{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the EnabledResourceTypesListResultIterator type.
func NewEnabledResourceTypesListResultIterator(page EnabledResourceTypesListResultPage) EnabledResourceTypesListResultIterator {
	return EnabledResourceTypesListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (ertlr EnabledResourceTypesListResult) IsEmpty() bool {
	return ertlr.Value == nil || len(*ertlr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (ertlr EnabledResourceTypesListResult) hasNextLink() bool {
	return ertlr.NextLink != nil && len(*ertlr.NextLink) != 0
}

// enabledResourceTypesListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (ertlr EnabledResourceTypesListResult) enabledResourceTypesListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !ertlr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(ertlr.NextLink)))
}

// EnabledResourceTypesListResultPage contains a page of EnabledResourceType values.
type EnabledResourceTypesListResultPage struct {
	fn    func(context.Context, EnabledResourceTypesListResult) (EnabledResourceTypesListResult, error)
	ertlr EnabledResourceTypesListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *EnabledResourceTypesListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnabledResourceTypesListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.ertlr)
		if err != nil {
			return err
		}
		page.ertlr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *EnabledResourceTypesListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page EnabledResourceTypesListResultPage) NotDone() bool {
	return !page.ertlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page EnabledResourceTypesListResultPage) Response() EnabledResourceTypesListResult {
	return page.ertlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page EnabledResourceTypesListResultPage) Values() []EnabledResourceType {
	if page.ertlr.IsEmpty() {
		return nil
	}
	return *page.ertlr.Value
}

// Creates a new instance of the EnabledResourceTypesListResultPage type.
func NewEnabledResourceTypesListResultPage(cur EnabledResourceTypesListResult, getNextPage func(context.Context, EnabledResourceTypesListResult) (EnabledResourceTypesListResult, error)) EnabledResourceTypesListResultPage {
	return EnabledResourceTypesListResultPage{
		fn:    getNextPage,
		ertlr: cur,
	}
}

// ErrorAdditionalInfo the resource management error additional info.
type ErrorAdditionalInfo struct {
	// Type - READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty"`
	// Info - READ-ONLY; The additional info.
	Info interface{} `json:"info,omitempty"`
}

// ErrorDetail the error detail.
type ErrorDetail struct {
	// Code - READ-ONLY; The error code.
	Code *string `json:"code,omitempty"`
	// Message - READ-ONLY; The error message.
	Message *string `json:"message,omitempty"`
	// Target - READ-ONLY; The error target.
	Target *string `json:"target,omitempty"`
	// Details - READ-ONLY; The error details.
	Details *[]ErrorDetail `json:"details,omitempty"`
	// AdditionalInfo - READ-ONLY; The error additional info.
	AdditionalInfo *[]ErrorAdditionalInfo `json:"additionalInfo,omitempty"`
}

// ErrorResponse common error response for all Azure Resource Manager APIs to return error details for
// failed operations. (This also follows the OData error response format.).
type ErrorResponse struct {
	// Error - The error object.
	Error *ErrorDetail `json:"error,omitempty"`
}

// PatchableCustomLocations the Custom Locations patchable resource definition.
type PatchableCustomLocations struct {
	// CustomLocationProperties - The Custom Locations patchable properties.
	*CustomLocationProperties `json:"properties,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for PatchableCustomLocations.
func (pcl PatchableCustomLocations) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if pcl.CustomLocationProperties != nil {
		objectMap["properties"] = pcl.CustomLocationProperties
	}
	if pcl.Tags != nil {
		objectMap["tags"] = pcl.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for PatchableCustomLocations struct.
func (pcl *PatchableCustomLocations) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var customLocationProperties CustomLocationProperties
				err = json.Unmarshal(*v, &customLocationProperties)
				if err != nil {
					return err
				}
				pcl.CustomLocationProperties = &customLocationProperties
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				pcl.Tags = tags
			}
		}
	}

	return nil
}

// ProxyResource the resource model definition for a Azure Resource Manager proxy resource. It will not
// have tags and a location
type ProxyResource struct {
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// Resource common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// SystemData metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// CreatedBy - The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`
	// CreatedByType - The type of identity that created the resource. Possible values include: 'User', 'Application', 'ManagedIdentity', 'Key'
	CreatedByType CreatedByType `json:"createdByType,omitempty"`
	// CreatedAt - The timestamp of resource creation (UTC).
	CreatedAt *date.Time `json:"createdAt,omitempty"`
	// LastModifiedBy - The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`
	// LastModifiedByType - The type of identity that last modified the resource. Possible values include: 'User', 'Application', 'ManagedIdentity', 'Key'
	LastModifiedByType CreatedByType `json:"lastModifiedByType,omitempty"`
	// LastModifiedAt - The timestamp of resource last modification (UTC)
	LastModifiedAt *date.Time `json:"lastModifiedAt,omitempty"`
}

// TrackedResource the resource model definition for an Azure Resource Manager tracked top level resource
// which has 'tags' and a 'location'
type TrackedResource struct {
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for TrackedResource.
func (tr TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if tr.Tags != nil {
		objectMap["tags"] = tr.Tags
	}
	if tr.Location != nil {
		objectMap["location"] = tr.Location
	}
	return json.Marshal(objectMap)
}
