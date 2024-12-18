/*
LeaseWeb API for launching and managing Public Cloud instances

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publiccloud

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the VpsBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VpsBase{}

// VpsBase struct for VpsBase
type VpsBase struct {
	Id string `json:"id"`
	Pack VpsPackType `json:"pack"`
	Resources BaseResources `json:"resources"`
	Region RegionName `json:"region"`
	Datacenter Datacenter `json:"datacenter"`
	// The identifying name set to the instance
	Reference string `json:"reference"`
	Image Image `json:"image"`
	MarketAppId NullableMarketAppId `json:"marketAppId,omitempty"`
	State VpsState `json:"state"`
	HasPublicIpV4 bool `json:"hasPublicIpV4"`
	// The root disk's size in GB. Must be at least 5 GB for Linux and FreeBSD instances and 50 GB for Windows instances
	RootDiskSize int32 `json:"rootDiskSize"`
	// Date and time when the VPS was started for the first time, right after launching it
	StartedAt NullableTime `json:"startedAt"`
	Contract VpsContract `json:"contract"`
	AdditionalProperties map[string]interface{}
}

type _VpsBase VpsBase

// NewVpsBase instantiates a new VpsBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVpsBase(id string, pack VpsPackType, resources BaseResources, region RegionName, datacenter Datacenter, reference string, image Image, state VpsState, hasPublicIpV4 bool, rootDiskSize int32, startedAt NullableTime, contract VpsContract) *VpsBase {
	this := VpsBase{}
	this.Id = id
	this.Pack = pack
	this.Resources = resources
	this.Region = region
	this.Datacenter = datacenter
	this.Reference = reference
	this.Image = image
	this.State = state
	this.HasPublicIpV4 = hasPublicIpV4
	this.RootDiskSize = rootDiskSize
	this.StartedAt = startedAt
	this.Contract = contract
	return &this
}

// NewVpsBaseWithDefaults instantiates a new VpsBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVpsBaseWithDefaults() *VpsBase {
	this := VpsBase{}
	return &this
}

// GetId returns the Id field value
func (o *VpsBase) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VpsBase) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VpsBase) SetId(v string) {
	o.Id = v
}

// GetPack returns the Pack field value
func (o *VpsBase) GetPack() VpsPackType {
	if o == nil {
		var ret VpsPackType
		return ret
	}

	return o.Pack
}

// GetPackOk returns a tuple with the Pack field value
// and a boolean to check if the value has been set.
func (o *VpsBase) GetPackOk() (*VpsPackType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pack, true
}

// SetPack sets field value
func (o *VpsBase) SetPack(v VpsPackType) {
	o.Pack = v
}

// GetResources returns the Resources field value
func (o *VpsBase) GetResources() BaseResources {
	if o == nil {
		var ret BaseResources
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *VpsBase) GetResourcesOk() (*BaseResources, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resources, true
}

// SetResources sets field value
func (o *VpsBase) SetResources(v BaseResources) {
	o.Resources = v
}

// GetRegion returns the Region field value
func (o *VpsBase) GetRegion() RegionName {
	if o == nil {
		var ret RegionName
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *VpsBase) GetRegionOk() (*RegionName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *VpsBase) SetRegion(v RegionName) {
	o.Region = v
}

// GetDatacenter returns the Datacenter field value
func (o *VpsBase) GetDatacenter() Datacenter {
	if o == nil {
		var ret Datacenter
		return ret
	}

	return o.Datacenter
}

// GetDatacenterOk returns a tuple with the Datacenter field value
// and a boolean to check if the value has been set.
func (o *VpsBase) GetDatacenterOk() (*Datacenter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datacenter, true
}

// SetDatacenter sets field value
func (o *VpsBase) SetDatacenter(v Datacenter) {
	o.Datacenter = v
}

// GetReference returns the Reference field value
func (o *VpsBase) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *VpsBase) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *VpsBase) SetReference(v string) {
	o.Reference = v
}

// GetImage returns the Image field value
func (o *VpsBase) GetImage() Image {
	if o == nil {
		var ret Image
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *VpsBase) GetImageOk() (*Image, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *VpsBase) SetImage(v Image) {
	o.Image = v
}

// GetMarketAppId returns the MarketAppId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VpsBase) GetMarketAppId() MarketAppId {
	if o == nil || IsNil(o.MarketAppId.Get()) {
		var ret MarketAppId
		return ret
	}
	return *o.MarketAppId.Get()
}

// GetMarketAppIdOk returns a tuple with the MarketAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VpsBase) GetMarketAppIdOk() (*MarketAppId, bool) {
	if o == nil {
		return nil, false
	}
	return o.MarketAppId.Get(), o.MarketAppId.IsSet()
}

// HasMarketAppId returns a boolean if a field has been set.
func (o *VpsBase) HasMarketAppId() bool {
	if o != nil && o.MarketAppId.IsSet() {
		return true
	}

	return false
}

// SetMarketAppId gets a reference to the given NullableMarketAppId and assigns it to the MarketAppId field.
func (o *VpsBase) SetMarketAppId(v MarketAppId) {
	o.MarketAppId.Set(&v)
}
// SetMarketAppIdNil sets the value for MarketAppId to be an explicit nil
func (o *VpsBase) SetMarketAppIdNil() {
	o.MarketAppId.Set(nil)
}

// UnsetMarketAppId ensures that no value is present for MarketAppId, not even an explicit nil
func (o *VpsBase) UnsetMarketAppId() {
	o.MarketAppId.Unset()
}

// GetState returns the State field value
func (o *VpsBase) GetState() VpsState {
	if o == nil {
		var ret VpsState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *VpsBase) GetStateOk() (*VpsState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *VpsBase) SetState(v VpsState) {
	o.State = v
}

// GetHasPublicIpV4 returns the HasPublicIpV4 field value
func (o *VpsBase) GetHasPublicIpV4() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasPublicIpV4
}

// GetHasPublicIpV4Ok returns a tuple with the HasPublicIpV4 field value
// and a boolean to check if the value has been set.
func (o *VpsBase) GetHasPublicIpV4Ok() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasPublicIpV4, true
}

// SetHasPublicIpV4 sets field value
func (o *VpsBase) SetHasPublicIpV4(v bool) {
	o.HasPublicIpV4 = v
}

// GetRootDiskSize returns the RootDiskSize field value
func (o *VpsBase) GetRootDiskSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RootDiskSize
}

// GetRootDiskSizeOk returns a tuple with the RootDiskSize field value
// and a boolean to check if the value has been set.
func (o *VpsBase) GetRootDiskSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RootDiskSize, true
}

// SetRootDiskSize sets field value
func (o *VpsBase) SetRootDiskSize(v int32) {
	o.RootDiskSize = v
}

// GetStartedAt returns the StartedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *VpsBase) GetStartedAt() time.Time {
	if o == nil || o.StartedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.StartedAt.Get()
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VpsBase) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedAt.Get(), o.StartedAt.IsSet()
}

// SetStartedAt sets field value
func (o *VpsBase) SetStartedAt(v time.Time) {
	o.StartedAt.Set(&v)
}

// GetContract returns the Contract field value
func (o *VpsBase) GetContract() VpsContract {
	if o == nil {
		var ret VpsContract
		return ret
	}

	return o.Contract
}

// GetContractOk returns a tuple with the Contract field value
// and a boolean to check if the value has been set.
func (o *VpsBase) GetContractOk() (*VpsContract, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contract, true
}

// SetContract sets field value
func (o *VpsBase) SetContract(v VpsContract) {
	o.Contract = v
}

func (o VpsBase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VpsBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["pack"] = o.Pack
	toSerialize["resources"] = o.Resources
	toSerialize["region"] = o.Region
	toSerialize["datacenter"] = o.Datacenter
	toSerialize["reference"] = o.Reference
	toSerialize["image"] = o.Image
	if o.MarketAppId.IsSet() {
		toSerialize["marketAppId"] = o.MarketAppId.Get()
	}
	toSerialize["state"] = o.State
	toSerialize["hasPublicIpV4"] = o.HasPublicIpV4
	toSerialize["rootDiskSize"] = o.RootDiskSize
	toSerialize["startedAt"] = o.StartedAt.Get()
	toSerialize["contract"] = o.Contract

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VpsBase) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"pack",
		"resources",
		"region",
		"datacenter",
		"reference",
		"image",
		"state",
		"hasPublicIpV4",
		"rootDiskSize",
		"startedAt",
		"contract",
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

	varVpsBase := _VpsBase{}

	err = json.Unmarshal(data, &varVpsBase)

	if err != nil {
		return err
	}

	*o = VpsBase(varVpsBase)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "pack")
		delete(additionalProperties, "resources")
		delete(additionalProperties, "region")
		delete(additionalProperties, "datacenter")
		delete(additionalProperties, "reference")
		delete(additionalProperties, "image")
		delete(additionalProperties, "marketAppId")
		delete(additionalProperties, "state")
		delete(additionalProperties, "hasPublicIpV4")
		delete(additionalProperties, "rootDiskSize")
		delete(additionalProperties, "startedAt")
		delete(additionalProperties, "contract")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVpsBase struct {
	value *VpsBase
	isSet bool
}

func (v NullableVpsBase) Get() *VpsBase {
	return v.value
}

func (v *NullableVpsBase) Set(val *VpsBase) {
	v.value = val
	v.isSet = true
}

func (v NullableVpsBase) IsSet() bool {
	return v.isSet
}

func (v *NullableVpsBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVpsBase(val *VpsBase) *NullableVpsBase {
	return &NullableVpsBase{value: val, isSet: true}
}

func (v NullableVpsBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVpsBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

