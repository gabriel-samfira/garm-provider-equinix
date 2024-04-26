/*
Metal API

# Introduction Equinix Metal provides a RESTful HTTP API which can be reached at <https://api.equinix.com/metal/v1>. This document describes the API and how to use it.  The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account. Every feature of the Equinix Metal web interface is accessible through the API.  The API docs are generated from the Equinix Metal OpenAPI specification and are officially hosted at <https://metal.equinix.com/developers/api>.  # Common Parameters  The Equinix Metal API uses a few methods to minimize network traffic and improve throughput. These parameters are not used in all API calls, but are used often enough to warrant their own section. Look for these parameters in the documentation for the API calls that support them.  ## Pagination  Pagination is used to limit the number of results returned in a single request. The API will return a maximum of 100 results per page. To retrieve additional results, you can use the `page` and `per_page` query parameters.  The `page` parameter is used to specify the page number. The first page is `1`. The `per_page` parameter is used to specify the number of results per page. The maximum number of results differs by resource type.  ## Sorting  Where offered, the API allows you to sort results by a specific field. To sort results use the `sort_by` query parameter with the root level field name as the value. The `sort_direction` parameter is used to specify the sort direction, either either `asc` (ascending) or `desc` (descending).  ## Filtering  Filtering is used to limit the results returned in a single request. The API supports filtering by certain fields in the response. To filter results, you can use the field as a query parameter.  For example, to filter the IP list to only return public IPv4 addresses, you can filter by the `type` field, as in the following request:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/projects/id/ips?type=public_ipv4 ```  Only IP addresses with the `type` field set to `public_ipv4` will be returned.  ## Searching  Searching is used to find matching resources using multiple field comparissons. The API supports searching in resources that define this behavior. Currently the search parameter is only available on devices, ssh_keys, api_keys and memberships endpoints.  To search resources you can use the `search` query parameter.  ## Include and Exclude  For resources that contain references to other resources, sucha as a Device that refers to the Project it resides in, the Equinix Metal API will returns `href` values (API links) to the associated resource.  ```json {   ...   \"project\": {     \"href\": \"/metal/v1/projects/f3f131c8-f302-49ef-8c44-9405022dc6dd\"   } } ```  If you're going need the project details, you can avoid a second API request.  Specify the contained `href` resources and collections that you'd like to have included in the response using the `include` query parameter.  For example:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=projects ```  The `include` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests where `href` resources are presented.  To have multiple resources include, use a comma-separated list (e.g. `?include=emails,projects,memberships`).  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=emails,projects,memberships ```  You may also include nested associations up to three levels deep using dot notation (`?include=memberships.projects`):  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=memberships.projects ```  To exclude resources, and optimize response delivery, use the `exclude` query parameter. The `exclude` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests for fields with nested object responses. When excluded, these fields will be replaced with an object that contains only an `href` field.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
)

// checks if the PaymentMethodUpdateInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodUpdateInput{}

// PaymentMethodUpdateInput struct for PaymentMethodUpdateInput
type PaymentMethodUpdateInput struct {
	BillingAddress       map[string]interface{} `json:"billing_address,omitempty"`
	CardholderName       *string                `json:"cardholder_name,omitempty"`
	Default              *bool                  `json:"default,omitempty"`
	ExpirationMonth      *string                `json:"expiration_month,omitempty"`
	ExpirationYear       *int32                 `json:"expiration_year,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaymentMethodUpdateInput PaymentMethodUpdateInput

// NewPaymentMethodUpdateInput instantiates a new PaymentMethodUpdateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodUpdateInput() *PaymentMethodUpdateInput {
	this := PaymentMethodUpdateInput{}
	return &this
}

// NewPaymentMethodUpdateInputWithDefaults instantiates a new PaymentMethodUpdateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodUpdateInputWithDefaults() *PaymentMethodUpdateInput {
	this := PaymentMethodUpdateInput{}
	return &this
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *PaymentMethodUpdateInput) GetBillingAddress() map[string]interface{} {
	if o == nil || IsNil(o.BillingAddress) {
		var ret map[string]interface{}
		return ret
	}
	return o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodUpdateInput) GetBillingAddressOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.BillingAddress) {
		return map[string]interface{}{}, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *PaymentMethodUpdateInput) HasBillingAddress() bool {
	if o != nil && !IsNil(o.BillingAddress) {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given map[string]interface{} and assigns it to the BillingAddress field.
func (o *PaymentMethodUpdateInput) SetBillingAddress(v map[string]interface{}) {
	o.BillingAddress = v
}

// GetCardholderName returns the CardholderName field value if set, zero value otherwise.
func (o *PaymentMethodUpdateInput) GetCardholderName() string {
	if o == nil || IsNil(o.CardholderName) {
		var ret string
		return ret
	}
	return *o.CardholderName
}

// GetCardholderNameOk returns a tuple with the CardholderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodUpdateInput) GetCardholderNameOk() (*string, bool) {
	if o == nil || IsNil(o.CardholderName) {
		return nil, false
	}
	return o.CardholderName, true
}

// HasCardholderName returns a boolean if a field has been set.
func (o *PaymentMethodUpdateInput) HasCardholderName() bool {
	if o != nil && !IsNil(o.CardholderName) {
		return true
	}

	return false
}

// SetCardholderName gets a reference to the given string and assigns it to the CardholderName field.
func (o *PaymentMethodUpdateInput) SetCardholderName(v string) {
	o.CardholderName = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *PaymentMethodUpdateInput) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodUpdateInput) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *PaymentMethodUpdateInput) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *PaymentMethodUpdateInput) SetDefault(v bool) {
	o.Default = &v
}

// GetExpirationMonth returns the ExpirationMonth field value if set, zero value otherwise.
func (o *PaymentMethodUpdateInput) GetExpirationMonth() string {
	if o == nil || IsNil(o.ExpirationMonth) {
		var ret string
		return ret
	}
	return *o.ExpirationMonth
}

// GetExpirationMonthOk returns a tuple with the ExpirationMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodUpdateInput) GetExpirationMonthOk() (*string, bool) {
	if o == nil || IsNil(o.ExpirationMonth) {
		return nil, false
	}
	return o.ExpirationMonth, true
}

// HasExpirationMonth returns a boolean if a field has been set.
func (o *PaymentMethodUpdateInput) HasExpirationMonth() bool {
	if o != nil && !IsNil(o.ExpirationMonth) {
		return true
	}

	return false
}

// SetExpirationMonth gets a reference to the given string and assigns it to the ExpirationMonth field.
func (o *PaymentMethodUpdateInput) SetExpirationMonth(v string) {
	o.ExpirationMonth = &v
}

// GetExpirationYear returns the ExpirationYear field value if set, zero value otherwise.
func (o *PaymentMethodUpdateInput) GetExpirationYear() int32 {
	if o == nil || IsNil(o.ExpirationYear) {
		var ret int32
		return ret
	}
	return *o.ExpirationYear
}

// GetExpirationYearOk returns a tuple with the ExpirationYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodUpdateInput) GetExpirationYearOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpirationYear) {
		return nil, false
	}
	return o.ExpirationYear, true
}

// HasExpirationYear returns a boolean if a field has been set.
func (o *PaymentMethodUpdateInput) HasExpirationYear() bool {
	if o != nil && !IsNil(o.ExpirationYear) {
		return true
	}

	return false
}

// SetExpirationYear gets a reference to the given int32 and assigns it to the ExpirationYear field.
func (o *PaymentMethodUpdateInput) SetExpirationYear(v int32) {
	o.ExpirationYear = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PaymentMethodUpdateInput) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodUpdateInput) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PaymentMethodUpdateInput) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PaymentMethodUpdateInput) SetName(v string) {
	o.Name = &v
}

func (o PaymentMethodUpdateInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodUpdateInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BillingAddress) {
		toSerialize["billing_address"] = o.BillingAddress
	}
	if !IsNil(o.CardholderName) {
		toSerialize["cardholder_name"] = o.CardholderName
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.ExpirationMonth) {
		toSerialize["expiration_month"] = o.ExpirationMonth
	}
	if !IsNil(o.ExpirationYear) {
		toSerialize["expiration_year"] = o.ExpirationYear
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaymentMethodUpdateInput) UnmarshalJSON(bytes []byte) (err error) {
	varPaymentMethodUpdateInput := _PaymentMethodUpdateInput{}

	err = json.Unmarshal(bytes, &varPaymentMethodUpdateInput)

	if err != nil {
		return err
	}

	*o = PaymentMethodUpdateInput(varPaymentMethodUpdateInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "billing_address")
		delete(additionalProperties, "cardholder_name")
		delete(additionalProperties, "default")
		delete(additionalProperties, "expiration_month")
		delete(additionalProperties, "expiration_year")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentMethodUpdateInput struct {
	value *PaymentMethodUpdateInput
	isSet bool
}

func (v NullablePaymentMethodUpdateInput) Get() *PaymentMethodUpdateInput {
	return v.value
}

func (v *NullablePaymentMethodUpdateInput) Set(val *PaymentMethodUpdateInput) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodUpdateInput) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodUpdateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodUpdateInput(val *PaymentMethodUpdateInput) *NullablePaymentMethodUpdateInput {
	return &NullablePaymentMethodUpdateInput{value: val, isSet: true}
}

func (v NullablePaymentMethodUpdateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodUpdateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
