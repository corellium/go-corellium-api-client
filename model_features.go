/*
Corellium API

REST API to manage your virtual devices.

API version: 4.3.1-16664
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
)

// checks if the Features type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Features{}

// Features 
type Features struct {
	// 
	Apps NullableBool `json:"apps,omitempty"`
	// 
	Connect NullableBool `json:"connect,omitempty"`
	// 
	Console NullableBool `json:"console,omitempty"`
	// 
	Coretrace NullableBool `json:"coretrace,omitempty"`
	// 
	DeviceControl NullableBool `json:"deviceControl,omitempty"`
	// 
	DeviceDelete NullableBool `json:"deviceDelete,omitempty"`
	// 
	Files NullableBool `json:"files,omitempty"`
	// 
	Frida NullableBool `json:"frida,omitempty"`
	// 
	Images NullableBool `json:"images,omitempty"`
	// 
	Messaging NullableBool `json:"messaging,omitempty"`
	// 
	Netmon NullableBool `json:"netmon,omitempty"`
	// 
	Network NullableBool `json:"network,omitempty"`
	// 
	PortForwarding NullableBool `json:"portForwarding,omitempty"`
	// 
	PowerManagement NullableBool `json:"powerManagement,omitempty"`
	// 
	Profile NullableBool `json:"profile,omitempty"`
	// 
	Sensors NullableBool `json:"sensors,omitempty"`
	// 
	Settings NullableBool `json:"settings,omitempty"`
	// 
	Snapshots NullableBool `json:"snapshots,omitempty"`
	// 
	Strace NullableBool `json:"strace,omitempty"`
	// 
	System NullableBool `json:"system,omitempty"`
}

// NewFeatures instantiates a new Features object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatures() *Features {
	this := Features{}
	return &this
}

// NewFeaturesWithDefaults instantiates a new Features object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeaturesWithDefaults() *Features {
	this := Features{}
	return &this
}

// GetApps returns the Apps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Features) GetApps() bool {
	if o == nil || IsNil(o.Apps.Get()) {
		var ret bool
		return ret
	}
	return *o.Apps.Get()
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Features) GetAppsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Apps.Get(), o.Apps.IsSet()
}

// HasApps returns a boolean if a field has been set.
func (o *Features) HasApps() bool {
	if o != nil && o.Apps.IsSet() {
		return true
	}

	return false
}

// SetApps gets a reference to the given NullableBool and assigns it to the Apps field.
func (o *Features) SetApps(v bool) {
	o.Apps.Set(&v)
}
// SetAppsNil sets the value for Apps to be an explicit nil
func (o *Features) SetAppsNil() {
	o.Apps.Set(nil)
}

// UnsetApps ensures that no value is present for Apps, not even an explicit nil
func (o *Features) UnsetApps() {
	o.Apps.Unset()
}

// GetConnect returns the Connect field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Features) GetConnect() bool {
	if o == nil || IsNil(o.Connect.Get()) {
		var ret bool
		return ret
	}
	return *o.Connect.Get()
}

// GetConnectOk returns a tuple with the Connect field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Features) GetConnectOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Connect.Get(), o.Connect.IsSet()
}

// HasConnect returns a boolean if a field has been set.
func (o *Features) HasConnect() bool {
	if o != nil && o.Connect.IsSet() {
		return true
	}

	return false
}

// SetConnect gets a reference to the given NullableBool and assigns it to the Connect field.
func (o *Features) SetConnect(v bool) {
	o.Connect.Set(&v)
}
// SetConnectNil sets the value for Connect to be an explicit nil
func (o *Features) SetConnectNil() {
	o.Connect.Set(nil)
}

// UnsetConnect ensures that no value is present for Connect, not even an explicit nil
func (o *Features) UnsetConnect() {
	o.Connect.Unset()
}

// GetConsole returns the Console field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Features) GetConsole() bool {
	if o == nil || IsNil(o.Console.Get()) {
		var ret bool
		return ret
	}
	return *o.Console.Get()
}

// GetConsoleOk returns a tuple with the Console field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Features) GetConsoleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Console.Get(), o.Console.IsSet()
}

// HasConsole returns a boolean if a field has been set.
func (o *Features) HasConsole() bool {
	if o != nil && o.Console.IsSet() {
		return true
	}

	return false
}

// SetConsole gets a reference to the given NullableBool and assigns it to the Console field.
func (o *Features) SetConsole(v bool) {
	o.Console.Set(&v)
}
// SetConsoleNil sets the value for Console to be an explicit nil
func (o *Features) SetConsoleNil() {
	o.Console.Set(nil)
}

// UnsetConsole ensures that no value is present for Console, not even an explicit nil
func (o *Features) UnsetConsole() {
	o.Console.Unset()
}

// GetCoretrace returns the Coretrace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Features) GetCoretrace() bool {
	if o == nil || IsNil(o.Coretrace.Get()) {
		var ret bool
		return ret
	}
	return *o.Coretrace.Get()
}

// GetCoretraceOk returns a tuple with the Coretrace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Features) GetCoretraceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Coretrace.Get(), o.Coretrace.IsSet()
}

// HasCoretrace returns a boolean if a field has been set.
func (o *Features) HasCoretrace() bool {
	if o != nil && o.Coretrace.IsSet() {
		return true
	}

	return false
}

// SetCoretrace gets a reference to the given NullableBool and assigns it to the Coretrace field.
func (o *Features) SetCoretrace(v bool) {
	o.Coretrace.Set(&v)
}
// SetCoretraceNil sets the value for Coretrace to be an explicit nil
func (o *Features) SetCoretraceNil() {
	o.Coretrace.Set(nil)
}

// UnsetCoretrace ensures that no value is present for Coretrace, not even an explicit nil
func (o *Features) UnsetCoretrace() {
	o.Coretrace.Unset()
}

// GetDeviceControl returns the DeviceControl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Features) GetDeviceControl() bool {
	if o == nil || IsNil(o.DeviceControl.Get()) {
		var ret bool
		return ret
	}
	return *o.DeviceControl.Get()
}

// GetDeviceControlOk returns a tuple with the DeviceControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Features) GetDeviceControlOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceControl.Get(), o.DeviceControl.IsSet()
}

// HasDeviceControl returns a boolean if a field has been set.
func (o *Features) HasDeviceControl() bool {
	if o != nil && o.DeviceControl.IsSet() {
		return true
	}

	return false
}

// SetDeviceControl gets a reference to the given NullableBool and assigns it to the DeviceControl field.
func (o *Features) SetDeviceControl(v bool) {
	o.DeviceControl.Set(&v)
}
// SetDeviceControlNil sets the value for DeviceControl to be an explicit nil
func (o *Features) SetDeviceControlNil() {
	o.DeviceControl.Set(nil)
}

// UnsetDeviceControl ensures that no value is present for DeviceControl, not even an explicit nil
func (o *Features) UnsetDeviceControl() {
	o.DeviceControl.Unset()
}

// GetDeviceDelete returns the DeviceDelete field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Features) GetDeviceDelete() bool {
	if o == nil || IsNil(o.DeviceDelete.Get()) {
		var ret bool
		return ret
	}
	return *o.DeviceDelete.Get()
}

// GetDeviceDeleteOk returns a tuple with the DeviceDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Features) GetDeviceDeleteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceDelete.Get(), o.DeviceDelete.IsSet()
}

// HasDeviceDelete returns a boolean if a field has been set.
func (o *Features) HasDeviceDelete() bool {
	if o != nil && o.DeviceDelete.IsSet() {
		return true
	}

	return false
}

// SetDeviceDelete gets a reference to the given NullableBool and assigns it to the DeviceDelete field.
func (o *Features) SetDeviceDelete(v bool) {
	o.DeviceDelete.Set(&v)
}
// SetDeviceDeleteNil sets the value for DeviceDelete to be an explicit nil
func (o *Features) SetDeviceDeleteNil() {
	o.DeviceDelete.Set(nil)
}

// UnsetDeviceDelete ensures that no value is present for DeviceDelete, not even an explicit nil
func (o *Features) UnsetDeviceDelete() {
	o.DeviceDelete.Unset()
}

// GetFiles returns the Files field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Features) GetFiles() bool {
	if o == nil || IsNil(o.Files.Get()) {
		var ret bool
		return ret
	}
	return *o.Files.Get()
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Features) GetFilesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Files.Get(), o.Files.IsSet()
}

// HasFiles returns a boolean if a field has been set.
func (o *Features) HasFiles() bool {
	if o != nil && o.Files.IsSet() {
		return true
	}

	return false
}

// SetFiles gets a reference to the given NullableBool and assigns it to the Files field.
func (o *Features) SetFiles(v bool) {
	o.Files.Set(&v)
}
// SetFilesNil sets the value for Files to be an explicit nil
func (o *Features) SetFilesNil() {
	o.Files.Set(nil)
}

// UnsetFiles ensures that no value is present for Files, not even an explicit nil
func (o *Features) UnsetFiles() {
	o.Files.Unset()
}

// GetFrida returns the Frida field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Features) GetFrida() bool {
	if o == nil || IsNil(o.Frida.Get()) {
		var ret bool
		return ret
	}
	return *o.Frida.Get()
}

// GetFridaOk returns a tuple with the Frida field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Features) GetFridaOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Frida.Get(), o.Frida.IsSet()
}

// HasFrida returns a boolean if a field has been set.
func (o *Features) HasFrida() bool {
	if o != nil && o.Frida.IsSet() {
		return true
	}

	return false
}

// SetFrida gets a reference to the given NullableBool and assigns it to the Frida field.
func (o *Features) SetFrida(v bool) {
	o.Frida.Set(&v)
}
// SetFridaNil sets the value for Frida to be an explicit nil
func (o *Features) SetFridaNil() {
	o.Frida.Set(nil)
}

// UnsetFrida ensures that no value is present for Frida, not even an explicit nil
func (o *Features) UnsetFrida() {
	o.Frida.Unset()
}

// GetImages returns the Images field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Features) GetImages() bool {
	if o == nil || IsNil(o.Images.Get()) {
		var ret bool
		return ret
	}
	return *o.Images.Get()
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Features) GetImagesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Images.Get(), o.Images.IsSet()
}

// HasImages returns a boolean if a field has been set.
func (o *Features) HasImages() bool {
	if o != nil && o.Images.IsSet() {
		return true
	}

	return false
}

// SetImages gets a reference to the given NullableBool and assigns it to the Images field.
func (o *Features) SetImages(v bool) {
	o.Images.Set(&v)
}
// SetImagesNil sets the value for Images to be an explicit nil
func (o *Features) SetImagesNil() {
	o.Images.Set(nil)
}

// UnsetImages ensures that no value is present for Images, not even an explicit nil
func (o *Features) UnsetImages() {
	o.Images.Unset()
}

// GetMessaging returns the Messaging field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Features) GetMessaging() bool {
	if o == nil || IsNil(o.Messaging.Get()) {
		var ret bool
		return ret
	}
	return *o.Messaging.Get()
}

// GetMessagingOk returns a tuple with the Messaging field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Features) GetMessagingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Messaging.Get(), o.Messaging.IsSet()
}

// HasMessaging returns a boolean if a field has been set.
func (o *Features) HasMessaging() bool {
	if o != nil && o.Messaging.IsSet() {
		return true
	}

	return false
}

// SetMessaging gets a reference to the given NullableBool and assigns it to the Messaging field.
func (o *Features) SetMessaging(v bool) {
	o.Messaging.Set(&v)
}
// SetMessagingNil sets the value for Messaging to be an explicit nil
func (o *Features) SetMessagingNil() {
	o.Messaging.Set(nil)
}

// UnsetMessaging ensures that no value is present for Messaging, not even an explicit nil
func (o *Features) UnsetMessaging() {
	o.Messaging.Unset()
}

// GetNetmon returns the Netmon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Features) GetNetmon() bool {
	if o == nil || IsNil(o.Netmon.Get()) {
		var ret bool
		return ret
	}
	return *o.Netmon.Get()
}

// GetNetmonOk returns a tuple with the Netmon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Features) GetNetmonOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Netmon.Get(), o.Netmon.IsSet()
}

// HasNetmon returns a boolean if a field has been set.
func (o *Features) HasNetmon() bool {
	if o != nil && o.Netmon.IsSet() {
		return true
	}

	return false
}

// SetNetmon gets a reference to the given NullableBool and assigns it to the Netmon field.
func (o *Features) SetNetmon(v bool) {
	o.Netmon.Set(&v)
}
// SetNetmonNil sets the value for Netmon to be an explicit nil
func (o *Features) SetNetmonNil() {
	o.Netmon.Set(nil)
}

// UnsetNetmon ensures that no value is present for Netmon, not even an explicit nil
func (o *Features) UnsetNetmon() {
	o.Netmon.Unset()
}

// GetNetwork returns the Network field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Features) GetNetwork() bool {
	if o == nil || IsNil(o.Network.Get()) {
		var ret bool
		return ret
	}
	return *o.Network.Get()
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Features) GetNetworkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Network.Get(), o.Network.IsSet()
}

// HasNetwork returns a boolean if a field has been set.
func (o *Features) HasNetwork() bool {
	if o != nil && o.Network.IsSet() {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given NullableBool and assigns it to the Network field.
func (o *Features) SetNetwork(v bool) {
	o.Network.Set(&v)
}
// SetNetworkNil sets the value for Network to be an explicit nil
func (o *Features) SetNetworkNil() {
	o.Network.Set(nil)
}

// UnsetNetwork ensures that no value is present for Network, not even an explicit nil
func (o *Features) UnsetNetwork() {
	o.Network.Unset()
}

// GetPortForwarding returns the PortForwarding field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Features) GetPortForwarding() bool {
	if o == nil || IsNil(o.PortForwarding.Get()) {
		var ret bool
		return ret
	}
	return *o.PortForwarding.Get()
}

// GetPortForwardingOk returns a tuple with the PortForwarding field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Features) GetPortForwardingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PortForwarding.Get(), o.PortForwarding.IsSet()
}

// HasPortForwarding returns a boolean if a field has been set.
func (o *Features) HasPortForwarding() bool {
	if o != nil && o.PortForwarding.IsSet() {
		return true
	}

	return false
}

// SetPortForwarding gets a reference to the given NullableBool and assigns it to the PortForwarding field.
func (o *Features) SetPortForwarding(v bool) {
	o.PortForwarding.Set(&v)
}
// SetPortForwardingNil sets the value for PortForwarding to be an explicit nil
func (o *Features) SetPortForwardingNil() {
	o.PortForwarding.Set(nil)
}

// UnsetPortForwarding ensures that no value is present for PortForwarding, not even an explicit nil
func (o *Features) UnsetPortForwarding() {
	o.PortForwarding.Unset()
}

// GetPowerManagement returns the PowerManagement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Features) GetPowerManagement() bool {
	if o == nil || IsNil(o.PowerManagement.Get()) {
		var ret bool
		return ret
	}
	return *o.PowerManagement.Get()
}

// GetPowerManagementOk returns a tuple with the PowerManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Features) GetPowerManagementOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PowerManagement.Get(), o.PowerManagement.IsSet()
}

// HasPowerManagement returns a boolean if a field has been set.
func (o *Features) HasPowerManagement() bool {
	if o != nil && o.PowerManagement.IsSet() {
		return true
	}

	return false
}

// SetPowerManagement gets a reference to the given NullableBool and assigns it to the PowerManagement field.
func (o *Features) SetPowerManagement(v bool) {
	o.PowerManagement.Set(&v)
}
// SetPowerManagementNil sets the value for PowerManagement to be an explicit nil
func (o *Features) SetPowerManagementNil() {
	o.PowerManagement.Set(nil)
}

// UnsetPowerManagement ensures that no value is present for PowerManagement, not even an explicit nil
func (o *Features) UnsetPowerManagement() {
	o.PowerManagement.Unset()
}

// GetProfile returns the Profile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Features) GetProfile() bool {
	if o == nil || IsNil(o.Profile.Get()) {
		var ret bool
		return ret
	}
	return *o.Profile.Get()
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Features) GetProfileOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Profile.Get(), o.Profile.IsSet()
}

// HasProfile returns a boolean if a field has been set.
func (o *Features) HasProfile() bool {
	if o != nil && o.Profile.IsSet() {
		return true
	}

	return false
}

// SetProfile gets a reference to the given NullableBool and assigns it to the Profile field.
func (o *Features) SetProfile(v bool) {
	o.Profile.Set(&v)
}
// SetProfileNil sets the value for Profile to be an explicit nil
func (o *Features) SetProfileNil() {
	o.Profile.Set(nil)
}

// UnsetProfile ensures that no value is present for Profile, not even an explicit nil
func (o *Features) UnsetProfile() {
	o.Profile.Unset()
}

// GetSensors returns the Sensors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Features) GetSensors() bool {
	if o == nil || IsNil(o.Sensors.Get()) {
		var ret bool
		return ret
	}
	return *o.Sensors.Get()
}

// GetSensorsOk returns a tuple with the Sensors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Features) GetSensorsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sensors.Get(), o.Sensors.IsSet()
}

// HasSensors returns a boolean if a field has been set.
func (o *Features) HasSensors() bool {
	if o != nil && o.Sensors.IsSet() {
		return true
	}

	return false
}

// SetSensors gets a reference to the given NullableBool and assigns it to the Sensors field.
func (o *Features) SetSensors(v bool) {
	o.Sensors.Set(&v)
}
// SetSensorsNil sets the value for Sensors to be an explicit nil
func (o *Features) SetSensorsNil() {
	o.Sensors.Set(nil)
}

// UnsetSensors ensures that no value is present for Sensors, not even an explicit nil
func (o *Features) UnsetSensors() {
	o.Sensors.Unset()
}

// GetSettings returns the Settings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Features) GetSettings() bool {
	if o == nil || IsNil(o.Settings.Get()) {
		var ret bool
		return ret
	}
	return *o.Settings.Get()
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Features) GetSettingsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Settings.Get(), o.Settings.IsSet()
}

// HasSettings returns a boolean if a field has been set.
func (o *Features) HasSettings() bool {
	if o != nil && o.Settings.IsSet() {
		return true
	}

	return false
}

// SetSettings gets a reference to the given NullableBool and assigns it to the Settings field.
func (o *Features) SetSettings(v bool) {
	o.Settings.Set(&v)
}
// SetSettingsNil sets the value for Settings to be an explicit nil
func (o *Features) SetSettingsNil() {
	o.Settings.Set(nil)
}

// UnsetSettings ensures that no value is present for Settings, not even an explicit nil
func (o *Features) UnsetSettings() {
	o.Settings.Unset()
}

// GetSnapshots returns the Snapshots field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Features) GetSnapshots() bool {
	if o == nil || IsNil(o.Snapshots.Get()) {
		var ret bool
		return ret
	}
	return *o.Snapshots.Get()
}

// GetSnapshotsOk returns a tuple with the Snapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Features) GetSnapshotsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Snapshots.Get(), o.Snapshots.IsSet()
}

// HasSnapshots returns a boolean if a field has been set.
func (o *Features) HasSnapshots() bool {
	if o != nil && o.Snapshots.IsSet() {
		return true
	}

	return false
}

// SetSnapshots gets a reference to the given NullableBool and assigns it to the Snapshots field.
func (o *Features) SetSnapshots(v bool) {
	o.Snapshots.Set(&v)
}
// SetSnapshotsNil sets the value for Snapshots to be an explicit nil
func (o *Features) SetSnapshotsNil() {
	o.Snapshots.Set(nil)
}

// UnsetSnapshots ensures that no value is present for Snapshots, not even an explicit nil
func (o *Features) UnsetSnapshots() {
	o.Snapshots.Unset()
}

// GetStrace returns the Strace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Features) GetStrace() bool {
	if o == nil || IsNil(o.Strace.Get()) {
		var ret bool
		return ret
	}
	return *o.Strace.Get()
}

// GetStraceOk returns a tuple with the Strace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Features) GetStraceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Strace.Get(), o.Strace.IsSet()
}

// HasStrace returns a boolean if a field has been set.
func (o *Features) HasStrace() bool {
	if o != nil && o.Strace.IsSet() {
		return true
	}

	return false
}

// SetStrace gets a reference to the given NullableBool and assigns it to the Strace field.
func (o *Features) SetStrace(v bool) {
	o.Strace.Set(&v)
}
// SetStraceNil sets the value for Strace to be an explicit nil
func (o *Features) SetStraceNil() {
	o.Strace.Set(nil)
}

// UnsetStrace ensures that no value is present for Strace, not even an explicit nil
func (o *Features) UnsetStrace() {
	o.Strace.Unset()
}

// GetSystem returns the System field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Features) GetSystem() bool {
	if o == nil || IsNil(o.System.Get()) {
		var ret bool
		return ret
	}
	return *o.System.Get()
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Features) GetSystemOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.System.Get(), o.System.IsSet()
}

// HasSystem returns a boolean if a field has been set.
func (o *Features) HasSystem() bool {
	if o != nil && o.System.IsSet() {
		return true
	}

	return false
}

// SetSystem gets a reference to the given NullableBool and assigns it to the System field.
func (o *Features) SetSystem(v bool) {
	o.System.Set(&v)
}
// SetSystemNil sets the value for System to be an explicit nil
func (o *Features) SetSystemNil() {
	o.System.Set(nil)
}

// UnsetSystem ensures that no value is present for System, not even an explicit nil
func (o *Features) UnsetSystem() {
	o.System.Unset()
}

func (o Features) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Features) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Apps.IsSet() {
		toSerialize["apps"] = o.Apps.Get()
	}
	if o.Connect.IsSet() {
		toSerialize["connect"] = o.Connect.Get()
	}
	if o.Console.IsSet() {
		toSerialize["console"] = o.Console.Get()
	}
	if o.Coretrace.IsSet() {
		toSerialize["coretrace"] = o.Coretrace.Get()
	}
	if o.DeviceControl.IsSet() {
		toSerialize["deviceControl"] = o.DeviceControl.Get()
	}
	if o.DeviceDelete.IsSet() {
		toSerialize["deviceDelete"] = o.DeviceDelete.Get()
	}
	if o.Files.IsSet() {
		toSerialize["files"] = o.Files.Get()
	}
	if o.Frida.IsSet() {
		toSerialize["frida"] = o.Frida.Get()
	}
	if o.Images.IsSet() {
		toSerialize["images"] = o.Images.Get()
	}
	if o.Messaging.IsSet() {
		toSerialize["messaging"] = o.Messaging.Get()
	}
	if o.Netmon.IsSet() {
		toSerialize["netmon"] = o.Netmon.Get()
	}
	if o.Network.IsSet() {
		toSerialize["network"] = o.Network.Get()
	}
	if o.PortForwarding.IsSet() {
		toSerialize["portForwarding"] = o.PortForwarding.Get()
	}
	if o.PowerManagement.IsSet() {
		toSerialize["powerManagement"] = o.PowerManagement.Get()
	}
	if o.Profile.IsSet() {
		toSerialize["profile"] = o.Profile.Get()
	}
	if o.Sensors.IsSet() {
		toSerialize["sensors"] = o.Sensors.Get()
	}
	if o.Settings.IsSet() {
		toSerialize["settings"] = o.Settings.Get()
	}
	if o.Snapshots.IsSet() {
		toSerialize["snapshots"] = o.Snapshots.Get()
	}
	if o.Strace.IsSet() {
		toSerialize["strace"] = o.Strace.Get()
	}
	if o.System.IsSet() {
		toSerialize["system"] = o.System.Get()
	}
	return toSerialize, nil
}

type NullableFeatures struct {
	value *Features
	isSet bool
}

func (v NullableFeatures) Get() *Features {
	return v.value
}

func (v *NullableFeatures) Set(val *Features) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatures) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatures(val *Features) *NullableFeatures {
	return &NullableFeatures{value: val, isSet: true}
}

func (v NullableFeatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


