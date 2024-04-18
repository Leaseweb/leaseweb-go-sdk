/*
LeaseWeb API for launching and managing Public Cloud instances

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publicCloud

import (
	"encoding/json"
	"time"
)

// checks if the UpdateInstanceResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateInstanceResult{}

// UpdateInstanceResult struct for UpdateInstanceResult
type UpdateInstanceResult struct {
	// The customer's equipment ID
	EquipmentId *string `json:"equipmentId,omitempty"`
	// The customer's ID
	CustomerId *string `json:"customerId,omitempty"`
	SalesOrgId *string `json:"salesOrgId,omitempty"`
	// The instance's identifier
	Id *string `json:"id,omitempty"`
	// Instance type
	Type *string `json:"type,omitempty"`
	// The region where the instance has been launched into
	Region *string `json:"region,omitempty"`
	// The identifying name set to the instance
	Reference *string `json:"reference,omitempty"`
	OperatingSystem *OperatingSystem `json:"operatingSystem,omitempty"`
	Resources *InstanceResources `json:"resources,omitempty"`
	State *string `json:"state,omitempty"`
	PublicIpV4 *bool `json:"publicIpV4,omitempty"`
	PrivateNetwork *bool `json:"privateNetwork,omitempty"`
	StartedAt NullableTime `json:"startedAt,omitempty"`
	RootDiskSize *int32 `json:"rootDiskSize,omitempty"`
	Ips []Ip `json:"ips,omitempty"`
	BillingFrequency *int32 `json:"billingFrequency,omitempty"`
	ContractTerm *int32 `json:"contractTerm,omitempty"`
	ContractType *ContractType `json:"contractType,omitempty"`
	ContractEndsAt NullableTime `json:"contractEndsAt,omitempty"`
	ContractRenewalsAt *time.Time `json:"contractRenewalsAt,omitempty"`
	ContractCreatedAt *time.Time `json:"contractCreatedAt,omitempty"`
	Iso NullableString `json:"iso,omitempty"`
}

// NewUpdateInstanceResult instantiates a new UpdateInstanceResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateInstanceResult() *UpdateInstanceResult {
	this := UpdateInstanceResult{}
	return &this
}

// NewUpdateInstanceResultWithDefaults instantiates a new UpdateInstanceResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateInstanceResultWithDefaults() *UpdateInstanceResult {
	this := UpdateInstanceResult{}
	return &this
}

// GetEquipmentId returns the EquipmentId field value if set, zero value otherwise.
func (o *UpdateInstanceResult) GetEquipmentId() string {
	if o == nil || IsNil(o.EquipmentId) {
		var ret string
		return ret
	}
	return *o.EquipmentId
}

// GetEquipmentIdOk returns a tuple with the EquipmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceResult) GetEquipmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.EquipmentId) {
		return nil, false
	}
	return o.EquipmentId, true
}

// HasEquipmentId returns a boolean if a field has been set.
func (o *UpdateInstanceResult) HasEquipmentId() bool {
	if o != nil && !IsNil(o.EquipmentId) {
		return true
	}

	return false
}

// SetEquipmentId gets a reference to the given string and assigns it to the EquipmentId field.
func (o *UpdateInstanceResult) SetEquipmentId(v string) {
	o.EquipmentId = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *UpdateInstanceResult) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceResult) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *UpdateInstanceResult) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *UpdateInstanceResult) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetSalesOrgId returns the SalesOrgId field value if set, zero value otherwise.
func (o *UpdateInstanceResult) GetSalesOrgId() string {
	if o == nil || IsNil(o.SalesOrgId) {
		var ret string
		return ret
	}
	return *o.SalesOrgId
}

// GetSalesOrgIdOk returns a tuple with the SalesOrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceResult) GetSalesOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.SalesOrgId) {
		return nil, false
	}
	return o.SalesOrgId, true
}

// HasSalesOrgId returns a boolean if a field has been set.
func (o *UpdateInstanceResult) HasSalesOrgId() bool {
	if o != nil && !IsNil(o.SalesOrgId) {
		return true
	}

	return false
}

// SetSalesOrgId gets a reference to the given string and assigns it to the SalesOrgId field.
func (o *UpdateInstanceResult) SetSalesOrgId(v string) {
	o.SalesOrgId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateInstanceResult) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceResult) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateInstanceResult) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UpdateInstanceResult) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UpdateInstanceResult) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceResult) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UpdateInstanceResult) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UpdateInstanceResult) SetType(v string) {
	o.Type = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *UpdateInstanceResult) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceResult) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *UpdateInstanceResult) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *UpdateInstanceResult) SetRegion(v string) {
	o.Region = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *UpdateInstanceResult) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceResult) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *UpdateInstanceResult) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *UpdateInstanceResult) SetReference(v string) {
	o.Reference = &v
}

// GetOperatingSystem returns the OperatingSystem field value if set, zero value otherwise.
func (o *UpdateInstanceResult) GetOperatingSystem() OperatingSystem {
	if o == nil || IsNil(o.OperatingSystem) {
		var ret OperatingSystem
		return ret
	}
	return *o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceResult) GetOperatingSystemOk() (*OperatingSystem, bool) {
	if o == nil || IsNil(o.OperatingSystem) {
		return nil, false
	}
	return o.OperatingSystem, true
}

// HasOperatingSystem returns a boolean if a field has been set.
func (o *UpdateInstanceResult) HasOperatingSystem() bool {
	if o != nil && !IsNil(o.OperatingSystem) {
		return true
	}

	return false
}

// SetOperatingSystem gets a reference to the given OperatingSystem and assigns it to the OperatingSystem field.
func (o *UpdateInstanceResult) SetOperatingSystem(v OperatingSystem) {
	o.OperatingSystem = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *UpdateInstanceResult) GetResources() InstanceResources {
	if o == nil || IsNil(o.Resources) {
		var ret InstanceResources
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceResult) GetResourcesOk() (*InstanceResources, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *UpdateInstanceResult) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given InstanceResources and assigns it to the Resources field.
func (o *UpdateInstanceResult) SetResources(v InstanceResources) {
	o.Resources = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *UpdateInstanceResult) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceResult) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *UpdateInstanceResult) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *UpdateInstanceResult) SetState(v string) {
	o.State = &v
}

// GetPublicIpV4 returns the PublicIpV4 field value if set, zero value otherwise.
func (o *UpdateInstanceResult) GetPublicIpV4() bool {
	if o == nil || IsNil(o.PublicIpV4) {
		var ret bool
		return ret
	}
	return *o.PublicIpV4
}

// GetPublicIpV4Ok returns a tuple with the PublicIpV4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceResult) GetPublicIpV4Ok() (*bool, bool) {
	if o == nil || IsNil(o.PublicIpV4) {
		return nil, false
	}
	return o.PublicIpV4, true
}

// HasPublicIpV4 returns a boolean if a field has been set.
func (o *UpdateInstanceResult) HasPublicIpV4() bool {
	if o != nil && !IsNil(o.PublicIpV4) {
		return true
	}

	return false
}

// SetPublicIpV4 gets a reference to the given bool and assigns it to the PublicIpV4 field.
func (o *UpdateInstanceResult) SetPublicIpV4(v bool) {
	o.PublicIpV4 = &v
}

// GetPrivateNetwork returns the PrivateNetwork field value if set, zero value otherwise.
func (o *UpdateInstanceResult) GetPrivateNetwork() bool {
	if o == nil || IsNil(o.PrivateNetwork) {
		var ret bool
		return ret
	}
	return *o.PrivateNetwork
}

// GetPrivateNetworkOk returns a tuple with the PrivateNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceResult) GetPrivateNetworkOk() (*bool, bool) {
	if o == nil || IsNil(o.PrivateNetwork) {
		return nil, false
	}
	return o.PrivateNetwork, true
}

// HasPrivateNetwork returns a boolean if a field has been set.
func (o *UpdateInstanceResult) HasPrivateNetwork() bool {
	if o != nil && !IsNil(o.PrivateNetwork) {
		return true
	}

	return false
}

// SetPrivateNetwork gets a reference to the given bool and assigns it to the PrivateNetwork field.
func (o *UpdateInstanceResult) SetPrivateNetwork(v bool) {
	o.PrivateNetwork = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateInstanceResult) GetStartedAt() time.Time {
	if o == nil || IsNil(o.StartedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.StartedAt.Get()
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateInstanceResult) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedAt.Get(), o.StartedAt.IsSet()
}

// HasStartedAt returns a boolean if a field has been set.
func (o *UpdateInstanceResult) HasStartedAt() bool {
	if o != nil && o.StartedAt.IsSet() {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given NullableTime and assigns it to the StartedAt field.
func (o *UpdateInstanceResult) SetStartedAt(v time.Time) {
	o.StartedAt.Set(&v)
}
// SetStartedAtNil sets the value for StartedAt to be an explicit nil
func (o *UpdateInstanceResult) SetStartedAtNil() {
	o.StartedAt.Set(nil)
}

// UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
func (o *UpdateInstanceResult) UnsetStartedAt() {
	o.StartedAt.Unset()
}

// GetRootDiskSize returns the RootDiskSize field value if set, zero value otherwise.
func (o *UpdateInstanceResult) GetRootDiskSize() int32 {
	if o == nil || IsNil(o.RootDiskSize) {
		var ret int32
		return ret
	}
	return *o.RootDiskSize
}

// GetRootDiskSizeOk returns a tuple with the RootDiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceResult) GetRootDiskSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.RootDiskSize) {
		return nil, false
	}
	return o.RootDiskSize, true
}

// HasRootDiskSize returns a boolean if a field has been set.
func (o *UpdateInstanceResult) HasRootDiskSize() bool {
	if o != nil && !IsNil(o.RootDiskSize) {
		return true
	}

	return false
}

// SetRootDiskSize gets a reference to the given int32 and assigns it to the RootDiskSize field.
func (o *UpdateInstanceResult) SetRootDiskSize(v int32) {
	o.RootDiskSize = &v
}

// GetIps returns the Ips field value if set, zero value otherwise.
func (o *UpdateInstanceResult) GetIps() []Ip {
	if o == nil || IsNil(o.Ips) {
		var ret []Ip
		return ret
	}
	return o.Ips
}

// GetIpsOk returns a tuple with the Ips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceResult) GetIpsOk() ([]Ip, bool) {
	if o == nil || IsNil(o.Ips) {
		return nil, false
	}
	return o.Ips, true
}

// HasIps returns a boolean if a field has been set.
func (o *UpdateInstanceResult) HasIps() bool {
	if o != nil && !IsNil(o.Ips) {
		return true
	}

	return false
}

// SetIps gets a reference to the given []Ip and assigns it to the Ips field.
func (o *UpdateInstanceResult) SetIps(v []Ip) {
	o.Ips = v
}

// GetBillingFrequency returns the BillingFrequency field value if set, zero value otherwise.
func (o *UpdateInstanceResult) GetBillingFrequency() int32 {
	if o == nil || IsNil(o.BillingFrequency) {
		var ret int32
		return ret
	}
	return *o.BillingFrequency
}

// GetBillingFrequencyOk returns a tuple with the BillingFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceResult) GetBillingFrequencyOk() (*int32, bool) {
	if o == nil || IsNil(o.BillingFrequency) {
		return nil, false
	}
	return o.BillingFrequency, true
}

// HasBillingFrequency returns a boolean if a field has been set.
func (o *UpdateInstanceResult) HasBillingFrequency() bool {
	if o != nil && !IsNil(o.BillingFrequency) {
		return true
	}

	return false
}

// SetBillingFrequency gets a reference to the given int32 and assigns it to the BillingFrequency field.
func (o *UpdateInstanceResult) SetBillingFrequency(v int32) {
	o.BillingFrequency = &v
}

// GetContractTerm returns the ContractTerm field value if set, zero value otherwise.
func (o *UpdateInstanceResult) GetContractTerm() int32 {
	if o == nil || IsNil(o.ContractTerm) {
		var ret int32
		return ret
	}
	return *o.ContractTerm
}

// GetContractTermOk returns a tuple with the ContractTerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceResult) GetContractTermOk() (*int32, bool) {
	if o == nil || IsNil(o.ContractTerm) {
		return nil, false
	}
	return o.ContractTerm, true
}

// HasContractTerm returns a boolean if a field has been set.
func (o *UpdateInstanceResult) HasContractTerm() bool {
	if o != nil && !IsNil(o.ContractTerm) {
		return true
	}

	return false
}

// SetContractTerm gets a reference to the given int32 and assigns it to the ContractTerm field.
func (o *UpdateInstanceResult) SetContractTerm(v int32) {
	o.ContractTerm = &v
}

// GetContractType returns the ContractType field value if set, zero value otherwise.
func (o *UpdateInstanceResult) GetContractType() ContractType {
	if o == nil || IsNil(o.ContractType) {
		var ret ContractType
		return ret
	}
	return *o.ContractType
}

// GetContractTypeOk returns a tuple with the ContractType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceResult) GetContractTypeOk() (*ContractType, bool) {
	if o == nil || IsNil(o.ContractType) {
		return nil, false
	}
	return o.ContractType, true
}

// HasContractType returns a boolean if a field has been set.
func (o *UpdateInstanceResult) HasContractType() bool {
	if o != nil && !IsNil(o.ContractType) {
		return true
	}

	return false
}

// SetContractType gets a reference to the given ContractType and assigns it to the ContractType field.
func (o *UpdateInstanceResult) SetContractType(v ContractType) {
	o.ContractType = &v
}

// GetContractEndsAt returns the ContractEndsAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateInstanceResult) GetContractEndsAt() time.Time {
	if o == nil || IsNil(o.ContractEndsAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ContractEndsAt.Get()
}

// GetContractEndsAtOk returns a tuple with the ContractEndsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateInstanceResult) GetContractEndsAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContractEndsAt.Get(), o.ContractEndsAt.IsSet()
}

// HasContractEndsAt returns a boolean if a field has been set.
func (o *UpdateInstanceResult) HasContractEndsAt() bool {
	if o != nil && o.ContractEndsAt.IsSet() {
		return true
	}

	return false
}

// SetContractEndsAt gets a reference to the given NullableTime and assigns it to the ContractEndsAt field.
func (o *UpdateInstanceResult) SetContractEndsAt(v time.Time) {
	o.ContractEndsAt.Set(&v)
}
// SetContractEndsAtNil sets the value for ContractEndsAt to be an explicit nil
func (o *UpdateInstanceResult) SetContractEndsAtNil() {
	o.ContractEndsAt.Set(nil)
}

// UnsetContractEndsAt ensures that no value is present for ContractEndsAt, not even an explicit nil
func (o *UpdateInstanceResult) UnsetContractEndsAt() {
	o.ContractEndsAt.Unset()
}

// GetContractRenewalsAt returns the ContractRenewalsAt field value if set, zero value otherwise.
func (o *UpdateInstanceResult) GetContractRenewalsAt() time.Time {
	if o == nil || IsNil(o.ContractRenewalsAt) {
		var ret time.Time
		return ret
	}
	return *o.ContractRenewalsAt
}

// GetContractRenewalsAtOk returns a tuple with the ContractRenewalsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceResult) GetContractRenewalsAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ContractRenewalsAt) {
		return nil, false
	}
	return o.ContractRenewalsAt, true
}

// HasContractRenewalsAt returns a boolean if a field has been set.
func (o *UpdateInstanceResult) HasContractRenewalsAt() bool {
	if o != nil && !IsNil(o.ContractRenewalsAt) {
		return true
	}

	return false
}

// SetContractRenewalsAt gets a reference to the given time.Time and assigns it to the ContractRenewalsAt field.
func (o *UpdateInstanceResult) SetContractRenewalsAt(v time.Time) {
	o.ContractRenewalsAt = &v
}

// GetContractCreatedAt returns the ContractCreatedAt field value if set, zero value otherwise.
func (o *UpdateInstanceResult) GetContractCreatedAt() time.Time {
	if o == nil || IsNil(o.ContractCreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.ContractCreatedAt
}

// GetContractCreatedAtOk returns a tuple with the ContractCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceResult) GetContractCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ContractCreatedAt) {
		return nil, false
	}
	return o.ContractCreatedAt, true
}

// HasContractCreatedAt returns a boolean if a field has been set.
func (o *UpdateInstanceResult) HasContractCreatedAt() bool {
	if o != nil && !IsNil(o.ContractCreatedAt) {
		return true
	}

	return false
}

// SetContractCreatedAt gets a reference to the given time.Time and assigns it to the ContractCreatedAt field.
func (o *UpdateInstanceResult) SetContractCreatedAt(v time.Time) {
	o.ContractCreatedAt = &v
}

// GetIso returns the Iso field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateInstanceResult) GetIso() string {
	if o == nil || IsNil(o.Iso.Get()) {
		var ret string
		return ret
	}
	return *o.Iso.Get()
}

// GetIsoOk returns a tuple with the Iso field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateInstanceResult) GetIsoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Iso.Get(), o.Iso.IsSet()
}

// HasIso returns a boolean if a field has been set.
func (o *UpdateInstanceResult) HasIso() bool {
	if o != nil && o.Iso.IsSet() {
		return true
	}

	return false
}

// SetIso gets a reference to the given NullableString and assigns it to the Iso field.
func (o *UpdateInstanceResult) SetIso(v string) {
	o.Iso.Set(&v)
}
// SetIsoNil sets the value for Iso to be an explicit nil
func (o *UpdateInstanceResult) SetIsoNil() {
	o.Iso.Set(nil)
}

// UnsetIso ensures that no value is present for Iso, not even an explicit nil
func (o *UpdateInstanceResult) UnsetIso() {
	o.Iso.Unset()
}

func (o UpdateInstanceResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateInstanceResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EquipmentId) {
		toSerialize["equipmentId"] = o.EquipmentId
	}
	if !IsNil(o.CustomerId) {
		toSerialize["customerId"] = o.CustomerId
	}
	if !IsNil(o.SalesOrgId) {
		toSerialize["salesOrgId"] = o.SalesOrgId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !IsNil(o.OperatingSystem) {
		toSerialize["operatingSystem"] = o.OperatingSystem
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.PublicIpV4) {
		toSerialize["publicIpV4"] = o.PublicIpV4
	}
	if !IsNil(o.PrivateNetwork) {
		toSerialize["privateNetwork"] = o.PrivateNetwork
	}
	if o.StartedAt.IsSet() {
		toSerialize["startedAt"] = o.StartedAt.Get()
	}
	if !IsNil(o.RootDiskSize) {
		toSerialize["rootDiskSize"] = o.RootDiskSize
	}
	if !IsNil(o.Ips) {
		toSerialize["ips"] = o.Ips
	}
	if !IsNil(o.BillingFrequency) {
		toSerialize["billingFrequency"] = o.BillingFrequency
	}
	if !IsNil(o.ContractTerm) {
		toSerialize["contractTerm"] = o.ContractTerm
	}
	if !IsNil(o.ContractType) {
		toSerialize["contractType"] = o.ContractType
	}
	if o.ContractEndsAt.IsSet() {
		toSerialize["contractEndsAt"] = o.ContractEndsAt.Get()
	}
	if !IsNil(o.ContractRenewalsAt) {
		toSerialize["contractRenewalsAt"] = o.ContractRenewalsAt
	}
	if !IsNil(o.ContractCreatedAt) {
		toSerialize["contractCreatedAt"] = o.ContractCreatedAt
	}
	if o.Iso.IsSet() {
		toSerialize["iso"] = o.Iso.Get()
	}
	return toSerialize, nil
}

type NullableUpdateInstanceResult struct {
	value *UpdateInstanceResult
	isSet bool
}

func (v NullableUpdateInstanceResult) Get() *UpdateInstanceResult {
	return v.value
}

func (v *NullableUpdateInstanceResult) Set(val *UpdateInstanceResult) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateInstanceResult) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateInstanceResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateInstanceResult(val *UpdateInstanceResult) *NullableUpdateInstanceResult {
	return &NullableUpdateInstanceResult{value: val, isSet: true}
}

func (v NullableUpdateInstanceResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateInstanceResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

