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

// checks if the Instance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Instance{}

// Instance struct for Instance
type Instance struct {
	// The unique identifier of the instance
	Id string `json:"id"`
	Type TypeName `json:"type"`
	Resources Resources `json:"resources"`
	Region RegionName `json:"region"`
	// The identifying name set to the instance
	Reference NullableString `json:"reference"`
	// Date and time when the instance was started for the first time, right after launching it
	StartedAt NullableTime `json:"startedAt"`
	// Market App ID that must be installed into the instance
	MarketAppId NullableString `json:"marketAppId"`
	State State `json:"state"`
	// The product type
	ProductType string `json:"productType"`
	HasPublicIpV4 bool `json:"hasPublicIpV4"`
	IncludesPrivateNetwork bool `json:"hasPrivateNetwork"`
	HasUserData bool `json:"hasUserData"`
	// The root disk's size in GB. Must be at least 5 GB for Linux and FreeBSD instances and 50 GB for Windows instances
	RootDiskSize int32 `json:"rootDiskSize"`
	RootDiskStorageType StorageType `json:"rootDiskStorageType"`
	Contract Contract `json:"contract"`
	AutoScalingGroup NullableAutoScalingGroup `json:"autoScalingGroup"`
	Image Image `json:"image"`
	Ips []Ip `json:"ips"`
	AdditionalProperties map[string]interface{}
}

type _Instance Instance

// NewInstance instantiates a new Instance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstance(id string, type_ TypeName, resources Resources, region RegionName, reference NullableString, startedAt NullableTime, marketAppId NullableString, state State, productType string, hasPublicIpV4 bool, includesPrivateNetwork bool, hasUserData bool, rootDiskSize int32, rootDiskStorageType StorageType, contract Contract, autoScalingGroup NullableAutoScalingGroup, image Image, ips []Ip) *Instance {
	this := Instance{}
	this.Id = id
	this.Type = type_
	this.Resources = resources
	this.Region = region
	this.Reference = reference
	this.StartedAt = startedAt
	this.MarketAppId = marketAppId
	this.State = state
	this.ProductType = productType
	this.HasPublicIpV4 = hasPublicIpV4
	this.IncludesPrivateNetwork = includesPrivateNetwork
	this.HasUserData = hasUserData
	this.RootDiskSize = rootDiskSize
	this.RootDiskStorageType = rootDiskStorageType
	this.Contract = contract
	this.AutoScalingGroup = autoScalingGroup
	this.Image = image
	this.Ips = ips
	return &this
}

// NewInstanceWithDefaults instantiates a new Instance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceWithDefaults() *Instance {
	this := Instance{}
	return &this
}

// GetId returns the Id field value
func (o *Instance) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Instance) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Instance) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *Instance) GetType() TypeName {
	if o == nil {
		var ret TypeName
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Instance) GetTypeOk() (*TypeName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Instance) SetType(v TypeName) {
	o.Type = v
}

// GetResources returns the Resources field value
func (o *Instance) GetResources() Resources {
	if o == nil {
		var ret Resources
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *Instance) GetResourcesOk() (*Resources, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resources, true
}

// SetResources sets field value
func (o *Instance) SetResources(v Resources) {
	o.Resources = v
}

// GetRegion returns the Region field value
func (o *Instance) GetRegion() RegionName {
	if o == nil {
		var ret RegionName
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *Instance) GetRegionOk() (*RegionName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *Instance) SetRegion(v RegionName) {
	o.Region = v
}

// GetReference returns the Reference field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Instance) GetReference() string {
	if o == nil || o.Reference.Get() == nil {
		var ret string
		return ret
	}

	return *o.Reference.Get()
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reference.Get(), o.Reference.IsSet()
}

// SetReference sets field value
func (o *Instance) SetReference(v string) {
	o.Reference.Set(&v)
}

// GetStartedAt returns the StartedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Instance) GetStartedAt() time.Time {
	if o == nil || o.StartedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.StartedAt.Get()
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedAt.Get(), o.StartedAt.IsSet()
}

// SetStartedAt sets field value
func (o *Instance) SetStartedAt(v time.Time) {
	o.StartedAt.Set(&v)
}

// GetMarketAppId returns the MarketAppId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Instance) GetMarketAppId() string {
	if o == nil || o.MarketAppId.Get() == nil {
		var ret string
		return ret
	}

	return *o.MarketAppId.Get()
}

// GetMarketAppIdOk returns a tuple with the MarketAppId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetMarketAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MarketAppId.Get(), o.MarketAppId.IsSet()
}

// SetMarketAppId sets field value
func (o *Instance) SetMarketAppId(v string) {
	o.MarketAppId.Set(&v)
}

// GetState returns the State field value
func (o *Instance) GetState() State {
	if o == nil {
		var ret State
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *Instance) GetStateOk() (*State, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *Instance) SetState(v State) {
	o.State = v
}

// GetProductType returns the ProductType field value
func (o *Instance) GetProductType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value
// and a boolean to check if the value has been set.
func (o *Instance) GetProductTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductType, true
}

// SetProductType sets field value
func (o *Instance) SetProductType(v string) {
	o.ProductType = v
}

// GetHasPublicIpV4 returns the HasPublicIpV4 field value
func (o *Instance) GetHasPublicIpV4() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasPublicIpV4
}

// GetHasPublicIpV4Ok returns a tuple with the HasPublicIpV4 field value
// and a boolean to check if the value has been set.
func (o *Instance) GetHasPublicIpV4Ok() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasPublicIpV4, true
}

// SetHasPublicIpV4 sets field value
func (o *Instance) SetHasPublicIpV4(v bool) {
	o.HasPublicIpV4 = v
}

// GetIncludesPrivateNetwork returns the IncludesPrivateNetwork field value
func (o *Instance) GetIncludesPrivateNetwork() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IncludesPrivateNetwork
}

// GetIncludesPrivateNetworkOk returns a tuple with the IncludesPrivateNetwork field value
// and a boolean to check if the value has been set.
func (o *Instance) GetIncludesPrivateNetworkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludesPrivateNetwork, true
}

// SetIncludesPrivateNetwork sets field value
func (o *Instance) SetIncludesPrivateNetwork(v bool) {
	o.IncludesPrivateNetwork = v
}

// GetHasUserData returns the HasUserData field value
func (o *Instance) GetHasUserData() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasUserData
}

// GetHasUserDataOk returns a tuple with the HasUserData field value
// and a boolean to check if the value has been set.
func (o *Instance) GetHasUserDataOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasUserData, true
}

// SetHasUserData sets field value
func (o *Instance) SetHasUserData(v bool) {
	o.HasUserData = v
}

// GetRootDiskSize returns the RootDiskSize field value
func (o *Instance) GetRootDiskSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RootDiskSize
}

// GetRootDiskSizeOk returns a tuple with the RootDiskSize field value
// and a boolean to check if the value has been set.
func (o *Instance) GetRootDiskSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RootDiskSize, true
}

// SetRootDiskSize sets field value
func (o *Instance) SetRootDiskSize(v int32) {
	o.RootDiskSize = v
}

// GetRootDiskStorageType returns the RootDiskStorageType field value
func (o *Instance) GetRootDiskStorageType() StorageType {
	if o == nil {
		var ret StorageType
		return ret
	}

	return o.RootDiskStorageType
}

// GetRootDiskStorageTypeOk returns a tuple with the RootDiskStorageType field value
// and a boolean to check if the value has been set.
func (o *Instance) GetRootDiskStorageTypeOk() (*StorageType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RootDiskStorageType, true
}

// SetRootDiskStorageType sets field value
func (o *Instance) SetRootDiskStorageType(v StorageType) {
	o.RootDiskStorageType = v
}

// GetContract returns the Contract field value
func (o *Instance) GetContract() Contract {
	if o == nil {
		var ret Contract
		return ret
	}

	return o.Contract
}

// GetContractOk returns a tuple with the Contract field value
// and a boolean to check if the value has been set.
func (o *Instance) GetContractOk() (*Contract, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contract, true
}

// SetContract sets field value
func (o *Instance) SetContract(v Contract) {
	o.Contract = v
}

// GetAutoScalingGroup returns the AutoScalingGroup field value
// If the value is explicit nil, the zero value for AutoScalingGroup will be returned
func (o *Instance) GetAutoScalingGroup() AutoScalingGroup {
	if o == nil || o.AutoScalingGroup.Get() == nil {
		var ret AutoScalingGroup
		return ret
	}

	return *o.AutoScalingGroup.Get()
}

// GetAutoScalingGroupOk returns a tuple with the AutoScalingGroup field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetAutoScalingGroupOk() (*AutoScalingGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoScalingGroup.Get(), o.AutoScalingGroup.IsSet()
}

// SetAutoScalingGroup sets field value
func (o *Instance) SetAutoScalingGroup(v AutoScalingGroup) {
	o.AutoScalingGroup.Set(&v)
}

// GetImage returns the Image field value
func (o *Instance) GetImage() Image {
	if o == nil {
		var ret Image
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *Instance) GetImageOk() (*Image, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *Instance) SetImage(v Image) {
	o.Image = v
}

// GetIps returns the Ips field value
func (o *Instance) GetIps() []Ip {
	if o == nil {
		var ret []Ip
		return ret
	}

	return o.Ips
}

// GetIpsOk returns a tuple with the Ips field value
// and a boolean to check if the value has been set.
func (o *Instance) GetIpsOk() ([]Ip, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ips, true
}

// SetIps sets field value
func (o *Instance) SetIps(v []Ip) {
	o.Ips = v
}

func (o Instance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Instance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["resources"] = o.Resources
	toSerialize["region"] = o.Region
	toSerialize["reference"] = o.Reference.Get()
	toSerialize["startedAt"] = o.StartedAt.Get()
	toSerialize["marketAppId"] = o.MarketAppId.Get()
	toSerialize["state"] = o.State
	toSerialize["productType"] = o.ProductType
	toSerialize["hasPublicIpV4"] = o.HasPublicIpV4
	toSerialize["hasPrivateNetwork"] = o.IncludesPrivateNetwork
	toSerialize["hasUserData"] = o.HasUserData
	toSerialize["rootDiskSize"] = o.RootDiskSize
	toSerialize["rootDiskStorageType"] = o.RootDiskStorageType
	toSerialize["contract"] = o.Contract
	toSerialize["autoScalingGroup"] = o.AutoScalingGroup.Get()
	toSerialize["image"] = o.Image
	toSerialize["ips"] = o.Ips

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Instance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"resources",
		"region",
		"reference",
		"startedAt",
		"marketAppId",
		"state",
		"productType",
		"hasPublicIpV4",
		"hasPrivateNetwork",
		"hasUserData",
		"rootDiskSize",
		"rootDiskStorageType",
		"contract",
		"autoScalingGroup",
		"image",
		"ips",
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

	varInstance := _Instance{}

	err = json.Unmarshal(data, &varInstance)

	if err != nil {
		return err
	}

	*o = Instance(varInstance)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "resources")
		delete(additionalProperties, "region")
		delete(additionalProperties, "reference")
		delete(additionalProperties, "startedAt")
		delete(additionalProperties, "marketAppId")
		delete(additionalProperties, "state")
		delete(additionalProperties, "productType")
		delete(additionalProperties, "hasPublicIpV4")
		delete(additionalProperties, "hasPrivateNetwork")
		delete(additionalProperties, "hasUserData")
		delete(additionalProperties, "rootDiskSize")
		delete(additionalProperties, "rootDiskStorageType")
		delete(additionalProperties, "contract")
		delete(additionalProperties, "autoScalingGroup")
		delete(additionalProperties, "image")
		delete(additionalProperties, "ips")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInstance struct {
	value *Instance
	isSet bool
}

func (v NullableInstance) Get() *Instance {
	return v.value
}

func (v *NullableInstance) Set(val *Instance) {
	v.value = val
	v.isSet = true
}

func (v NullableInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstance(val *Instance) *NullableInstance {
	return &NullableInstance{value: val, isSet: true}
}

func (v NullableInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

