/*
LeaseWeb API for launching and managing Public Cloud instances

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publicCloud

import (
	"encoding/json"
	"fmt"
)

// checks if the LaunchLoadBalancerOpts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LaunchLoadBalancerOpts{}

// LaunchLoadBalancerOpts struct for LaunchLoadBalancerOpts
type LaunchLoadBalancerOpts struct {
	Region RegionName `json:"region"`
	Type TypeName `json:"type"`
	// An identifying name you can refer to the load balancer
	Reference *string `json:"reference,omitempty"`
	// The contract applicable for the load balancer
	ContractType string `json:"contractType"`
	// How often you wish to be charged. Used only when contract type is MONTHLY. '1' means every month, '3' every three months and so on.
	BillingFrequency int32 `json:"billingFrequency"`
	RootDiskStorageType StorageType `json:"rootDiskStorageType"`
	// The port that the registered instances listen to
	TargetPort int32 `json:"targetPort"`
	AdditionalProperties map[string]interface{}
}

type _LaunchLoadBalancerOpts LaunchLoadBalancerOpts

// NewLaunchLoadBalancerOpts instantiates a new LaunchLoadBalancerOpts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLaunchLoadBalancerOpts(region RegionName, type_ TypeName, contractType string, billingFrequency int32, rootDiskStorageType StorageType, targetPort int32) *LaunchLoadBalancerOpts {
	this := LaunchLoadBalancerOpts{}
	this.Region = region
	this.Type = type_
	this.ContractType = contractType
	this.BillingFrequency = billingFrequency
	this.RootDiskStorageType = rootDiskStorageType
	this.TargetPort = targetPort
	return &this
}

// NewLaunchLoadBalancerOptsWithDefaults instantiates a new LaunchLoadBalancerOpts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLaunchLoadBalancerOptsWithDefaults() *LaunchLoadBalancerOpts {
	this := LaunchLoadBalancerOpts{}
	return &this
}

// GetRegion returns the Region field value
func (o *LaunchLoadBalancerOpts) GetRegion() RegionName {
	if o == nil {
		var ret RegionName
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *LaunchLoadBalancerOpts) GetRegionOk() (*RegionName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *LaunchLoadBalancerOpts) SetRegion(v RegionName) {
	o.Region = v
}

// GetType returns the Type field value
func (o *LaunchLoadBalancerOpts) GetType() TypeName {
	if o == nil {
		var ret TypeName
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LaunchLoadBalancerOpts) GetTypeOk() (*TypeName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LaunchLoadBalancerOpts) SetType(v TypeName) {
	o.Type = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *LaunchLoadBalancerOpts) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchLoadBalancerOpts) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *LaunchLoadBalancerOpts) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *LaunchLoadBalancerOpts) SetReference(v string) {
	o.Reference = &v
}

// GetContractType returns the ContractType field value
func (o *LaunchLoadBalancerOpts) GetContractType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractType
}

// GetContractTypeOk returns a tuple with the ContractType field value
// and a boolean to check if the value has been set.
func (o *LaunchLoadBalancerOpts) GetContractTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractType, true
}

// SetContractType sets field value
func (o *LaunchLoadBalancerOpts) SetContractType(v string) {
	o.ContractType = v
}

// GetBillingFrequency returns the BillingFrequency field value
func (o *LaunchLoadBalancerOpts) GetBillingFrequency() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BillingFrequency
}

// GetBillingFrequencyOk returns a tuple with the BillingFrequency field value
// and a boolean to check if the value has been set.
func (o *LaunchLoadBalancerOpts) GetBillingFrequencyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingFrequency, true
}

// SetBillingFrequency sets field value
func (o *LaunchLoadBalancerOpts) SetBillingFrequency(v int32) {
	o.BillingFrequency = v
}

// GetRootDiskStorageType returns the RootDiskStorageType field value
func (o *LaunchLoadBalancerOpts) GetRootDiskStorageType() StorageType {
	if o == nil {
		var ret StorageType
		return ret
	}

	return o.RootDiskStorageType
}

// GetRootDiskStorageTypeOk returns a tuple with the RootDiskStorageType field value
// and a boolean to check if the value has been set.
func (o *LaunchLoadBalancerOpts) GetRootDiskStorageTypeOk() (*StorageType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RootDiskStorageType, true
}

// SetRootDiskStorageType sets field value
func (o *LaunchLoadBalancerOpts) SetRootDiskStorageType(v StorageType) {
	o.RootDiskStorageType = v
}

// GetTargetPort returns the TargetPort field value
func (o *LaunchLoadBalancerOpts) GetTargetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TargetPort
}

// GetTargetPortOk returns a tuple with the TargetPort field value
// and a boolean to check if the value has been set.
func (o *LaunchLoadBalancerOpts) GetTargetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetPort, true
}

// SetTargetPort sets field value
func (o *LaunchLoadBalancerOpts) SetTargetPort(v int32) {
	o.TargetPort = v
}

func (o LaunchLoadBalancerOpts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LaunchLoadBalancerOpts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["region"] = o.Region
	toSerialize["type"] = o.Type
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	toSerialize["contractType"] = o.ContractType
	toSerialize["billingFrequency"] = o.BillingFrequency
	toSerialize["rootDiskStorageType"] = o.RootDiskStorageType
	toSerialize["targetPort"] = o.TargetPort

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LaunchLoadBalancerOpts) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"region",
		"type",
		"contractType",
		"billingFrequency",
		"rootDiskStorageType",
		"targetPort",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varLaunchLoadBalancerOpts := _LaunchLoadBalancerOpts{}

	err = json.Unmarshal(data, &varLaunchLoadBalancerOpts)

	if err != nil {
		return err
	}

	*o = LaunchLoadBalancerOpts(varLaunchLoadBalancerOpts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "region")
		delete(additionalProperties, "type")
		delete(additionalProperties, "reference")
		delete(additionalProperties, "contractType")
		delete(additionalProperties, "billingFrequency")
		delete(additionalProperties, "rootDiskStorageType")
		delete(additionalProperties, "targetPort")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLaunchLoadBalancerOpts struct {
	value *LaunchLoadBalancerOpts
	isSet bool
}

func (v NullableLaunchLoadBalancerOpts) Get() *LaunchLoadBalancerOpts {
	return v.value
}

func (v *NullableLaunchLoadBalancerOpts) Set(val *LaunchLoadBalancerOpts) {
	v.value = val
	v.isSet = true
}

func (v NullableLaunchLoadBalancerOpts) IsSet() bool {
	return v.isSet
}

func (v *NullableLaunchLoadBalancerOpts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLaunchLoadBalancerOpts(val *LaunchLoadBalancerOpts) *NullableLaunchLoadBalancerOpts {
	return &NullableLaunchLoadBalancerOpts{value: val, isSet: true}
}

func (v NullableLaunchLoadBalancerOpts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLaunchLoadBalancerOpts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


