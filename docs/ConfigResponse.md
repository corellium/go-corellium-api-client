# ConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnSite** | Pointer to **NullableBool** | Denotes whether it&#39;s an on-site install | [optional] 
**ShowDomainSettings** | Pointer to **NullableBool** | Denotes whether to show domain settings | [optional] 
**Version** | Pointer to **NullableString** | Denotes the version | [optional] 
**InvalidBuild** | Pointer to **NullableBool** | Denotes whether the build is invalid | [optional] 
**SepSim** | Pointer to **NullableBool** | Denotes whether sepSim is enabled | [optional] 
**InstallerAvailable** | Pointer to **NullableBool** | Denotes whether installer is available | [optional] 
**InvoicedPublishableKey** | Pointer to **NullableString** | API publishable key to use for Invoiced | [optional] 
**StripePublicKey** | Pointer to **NullableString** | Stripe public key | [optional] 
**IntercomId** | Pointer to **NullableString** | Intercom app ID, also known as workspace ID | [optional] 
**AuxWebhook** | Pointer to **NullableString** | Webhook URL for aux | [optional] 
**GtmId** | Pointer to **NullableString** | Google Tag Manager | [optional] 
**ZapierFeedbackWebhook** | Pointer to **NullableString** | Webhook URL to send feedback into Productboard automatically | [optional] 
**ZapierBugsWebhook** | Pointer to **NullableString** | Webhook URL to send frontend errors to Jira automatically | [optional] 
**BillingBackend** | Pointer to **NullableString** | Default backend billing api name for new subscriptions (e.g. \&quot;stripe\&quot;) | [optional] 
**Maintenance** | Pointer to [**NullableConfigResponseMaintenance**](ConfigResponseMaintenance.md) |  | [optional] 
**Logging** | Pointer to [**Logging**](Logging.md) |  | [optional] 
**CloudAdmin** | Pointer to **NullableString** | URL for cloud admin login | [optional] 
**FilesAdmin** | Pointer to **NullableString** | URL for files admin login | [optional] 
**CloudDomain** | Pointer to **NullableString** | Cloud domain name (usually corellium.com or staging.corellium.com) | [optional] 
**BillingDomain** | Pointer to **NullableString** | Billing domain name | [optional] 
**AppDomain** | Pointer to **NullableString** | App domain name (usually app.corellium.com or app.staging.corellium.com) | [optional] 
**MaxNetworkMonitorFileSize** | Pointer to **NullableString** | Maximum network monitor file size | [optional] 
**EnableFirmwareImageUpload** | Pointer to **NullableBool** | Denotes whether users can upload firmware images | [optional] 
**AuthProviders** | Pointer to [**[]AuthProvider**](AuthProvider.md) | Auth providers | [optional] 
**SnapshotLimit** | Pointer to **NullableFloat32** | Maximum number of snapshots to allow | [optional] 
**MaxKernelSize** | Pointer to **NullableFloat32** | The maximum size, in bytes, (default: 100 MB) that an uploaded kernel should be | [optional] 
**MaxRamdiskSize** | Pointer to **NullableFloat32** | The maximum size, in bytes, (default: 500 MB) that an uploaded ramdisk should be | [optional] 
**CharmSDK** | Pointer to **NullableString** | Denotes whether charmSDK is enabled | [optional] 
**Trial** | Pointer to [**Trial**](Trial.md) |  | [optional] 
**SentryUrl** | Pointer to **NullableString** | Sentry URL | [optional] 
**DomainAuthenticationProviders** | Pointer to **NullableBool** | If enabled, adds the default providers in their current configuration | [optional] 

## Methods

### NewConfigResponse

`func NewConfigResponse() *ConfigResponse`

NewConfigResponse instantiates a new ConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigResponseWithDefaults

`func NewConfigResponseWithDefaults() *ConfigResponse`

NewConfigResponseWithDefaults instantiates a new ConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnSite

`func (o *ConfigResponse) GetOnSite() bool`

GetOnSite returns the OnSite field if non-nil, zero value otherwise.

### GetOnSiteOk

`func (o *ConfigResponse) GetOnSiteOk() (*bool, bool)`

GetOnSiteOk returns a tuple with the OnSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSite

`func (o *ConfigResponse) SetOnSite(v bool)`

SetOnSite sets OnSite field to given value.

### HasOnSite

`func (o *ConfigResponse) HasOnSite() bool`

HasOnSite returns a boolean if a field has been set.

### SetOnSiteNil

`func (o *ConfigResponse) SetOnSiteNil(b bool)`

 SetOnSiteNil sets the value for OnSite to be an explicit nil

### UnsetOnSite
`func (o *ConfigResponse) UnsetOnSite()`

UnsetOnSite ensures that no value is present for OnSite, not even an explicit nil
### GetShowDomainSettings

`func (o *ConfigResponse) GetShowDomainSettings() bool`

GetShowDomainSettings returns the ShowDomainSettings field if non-nil, zero value otherwise.

### GetShowDomainSettingsOk

`func (o *ConfigResponse) GetShowDomainSettingsOk() (*bool, bool)`

GetShowDomainSettingsOk returns a tuple with the ShowDomainSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDomainSettings

`func (o *ConfigResponse) SetShowDomainSettings(v bool)`

SetShowDomainSettings sets ShowDomainSettings field to given value.

### HasShowDomainSettings

`func (o *ConfigResponse) HasShowDomainSettings() bool`

HasShowDomainSettings returns a boolean if a field has been set.

### SetShowDomainSettingsNil

`func (o *ConfigResponse) SetShowDomainSettingsNil(b bool)`

 SetShowDomainSettingsNil sets the value for ShowDomainSettings to be an explicit nil

### UnsetShowDomainSettings
`func (o *ConfigResponse) UnsetShowDomainSettings()`

UnsetShowDomainSettings ensures that no value is present for ShowDomainSettings, not even an explicit nil
### GetVersion

`func (o *ConfigResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ConfigResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ConfigResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ConfigResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ConfigResponse) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ConfigResponse) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetInvalidBuild

`func (o *ConfigResponse) GetInvalidBuild() bool`

GetInvalidBuild returns the InvalidBuild field if non-nil, zero value otherwise.

### GetInvalidBuildOk

`func (o *ConfigResponse) GetInvalidBuildOk() (*bool, bool)`

GetInvalidBuildOk returns a tuple with the InvalidBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidBuild

`func (o *ConfigResponse) SetInvalidBuild(v bool)`

SetInvalidBuild sets InvalidBuild field to given value.

### HasInvalidBuild

`func (o *ConfigResponse) HasInvalidBuild() bool`

HasInvalidBuild returns a boolean if a field has been set.

### SetInvalidBuildNil

`func (o *ConfigResponse) SetInvalidBuildNil(b bool)`

 SetInvalidBuildNil sets the value for InvalidBuild to be an explicit nil

### UnsetInvalidBuild
`func (o *ConfigResponse) UnsetInvalidBuild()`

UnsetInvalidBuild ensures that no value is present for InvalidBuild, not even an explicit nil
### GetSepSim

`func (o *ConfigResponse) GetSepSim() bool`

GetSepSim returns the SepSim field if non-nil, zero value otherwise.

### GetSepSimOk

`func (o *ConfigResponse) GetSepSimOk() (*bool, bool)`

GetSepSimOk returns a tuple with the SepSim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepSim

`func (o *ConfigResponse) SetSepSim(v bool)`

SetSepSim sets SepSim field to given value.

### HasSepSim

`func (o *ConfigResponse) HasSepSim() bool`

HasSepSim returns a boolean if a field has been set.

### SetSepSimNil

`func (o *ConfigResponse) SetSepSimNil(b bool)`

 SetSepSimNil sets the value for SepSim to be an explicit nil

### UnsetSepSim
`func (o *ConfigResponse) UnsetSepSim()`

UnsetSepSim ensures that no value is present for SepSim, not even an explicit nil
### GetInstallerAvailable

`func (o *ConfigResponse) GetInstallerAvailable() bool`

GetInstallerAvailable returns the InstallerAvailable field if non-nil, zero value otherwise.

### GetInstallerAvailableOk

`func (o *ConfigResponse) GetInstallerAvailableOk() (*bool, bool)`

GetInstallerAvailableOk returns a tuple with the InstallerAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallerAvailable

`func (o *ConfigResponse) SetInstallerAvailable(v bool)`

SetInstallerAvailable sets InstallerAvailable field to given value.

### HasInstallerAvailable

`func (o *ConfigResponse) HasInstallerAvailable() bool`

HasInstallerAvailable returns a boolean if a field has been set.

### SetInstallerAvailableNil

`func (o *ConfigResponse) SetInstallerAvailableNil(b bool)`

 SetInstallerAvailableNil sets the value for InstallerAvailable to be an explicit nil

### UnsetInstallerAvailable
`func (o *ConfigResponse) UnsetInstallerAvailable()`

UnsetInstallerAvailable ensures that no value is present for InstallerAvailable, not even an explicit nil
### GetInvoicedPublishableKey

`func (o *ConfigResponse) GetInvoicedPublishableKey() string`

GetInvoicedPublishableKey returns the InvoicedPublishableKey field if non-nil, zero value otherwise.

### GetInvoicedPublishableKeyOk

`func (o *ConfigResponse) GetInvoicedPublishableKeyOk() (*string, bool)`

GetInvoicedPublishableKeyOk returns a tuple with the InvoicedPublishableKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicedPublishableKey

`func (o *ConfigResponse) SetInvoicedPublishableKey(v string)`

SetInvoicedPublishableKey sets InvoicedPublishableKey field to given value.

### HasInvoicedPublishableKey

`func (o *ConfigResponse) HasInvoicedPublishableKey() bool`

HasInvoicedPublishableKey returns a boolean if a field has been set.

### SetInvoicedPublishableKeyNil

`func (o *ConfigResponse) SetInvoicedPublishableKeyNil(b bool)`

 SetInvoicedPublishableKeyNil sets the value for InvoicedPublishableKey to be an explicit nil

### UnsetInvoicedPublishableKey
`func (o *ConfigResponse) UnsetInvoicedPublishableKey()`

UnsetInvoicedPublishableKey ensures that no value is present for InvoicedPublishableKey, not even an explicit nil
### GetStripePublicKey

`func (o *ConfigResponse) GetStripePublicKey() string`

GetStripePublicKey returns the StripePublicKey field if non-nil, zero value otherwise.

### GetStripePublicKeyOk

`func (o *ConfigResponse) GetStripePublicKeyOk() (*string, bool)`

GetStripePublicKeyOk returns a tuple with the StripePublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripePublicKey

`func (o *ConfigResponse) SetStripePublicKey(v string)`

SetStripePublicKey sets StripePublicKey field to given value.

### HasStripePublicKey

`func (o *ConfigResponse) HasStripePublicKey() bool`

HasStripePublicKey returns a boolean if a field has been set.

### SetStripePublicKeyNil

`func (o *ConfigResponse) SetStripePublicKeyNil(b bool)`

 SetStripePublicKeyNil sets the value for StripePublicKey to be an explicit nil

### UnsetStripePublicKey
`func (o *ConfigResponse) UnsetStripePublicKey()`

UnsetStripePublicKey ensures that no value is present for StripePublicKey, not even an explicit nil
### GetIntercomId

`func (o *ConfigResponse) GetIntercomId() string`

GetIntercomId returns the IntercomId field if non-nil, zero value otherwise.

### GetIntercomIdOk

`func (o *ConfigResponse) GetIntercomIdOk() (*string, bool)`

GetIntercomIdOk returns a tuple with the IntercomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntercomId

`func (o *ConfigResponse) SetIntercomId(v string)`

SetIntercomId sets IntercomId field to given value.

### HasIntercomId

`func (o *ConfigResponse) HasIntercomId() bool`

HasIntercomId returns a boolean if a field has been set.

### SetIntercomIdNil

`func (o *ConfigResponse) SetIntercomIdNil(b bool)`

 SetIntercomIdNil sets the value for IntercomId to be an explicit nil

### UnsetIntercomId
`func (o *ConfigResponse) UnsetIntercomId()`

UnsetIntercomId ensures that no value is present for IntercomId, not even an explicit nil
### GetAuxWebhook

`func (o *ConfigResponse) GetAuxWebhook() string`

GetAuxWebhook returns the AuxWebhook field if non-nil, zero value otherwise.

### GetAuxWebhookOk

`func (o *ConfigResponse) GetAuxWebhookOk() (*string, bool)`

GetAuxWebhookOk returns a tuple with the AuxWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxWebhook

`func (o *ConfigResponse) SetAuxWebhook(v string)`

SetAuxWebhook sets AuxWebhook field to given value.

### HasAuxWebhook

`func (o *ConfigResponse) HasAuxWebhook() bool`

HasAuxWebhook returns a boolean if a field has been set.

### SetAuxWebhookNil

`func (o *ConfigResponse) SetAuxWebhookNil(b bool)`

 SetAuxWebhookNil sets the value for AuxWebhook to be an explicit nil

### UnsetAuxWebhook
`func (o *ConfigResponse) UnsetAuxWebhook()`

UnsetAuxWebhook ensures that no value is present for AuxWebhook, not even an explicit nil
### GetGtmId

`func (o *ConfigResponse) GetGtmId() string`

GetGtmId returns the GtmId field if non-nil, zero value otherwise.

### GetGtmIdOk

`func (o *ConfigResponse) GetGtmIdOk() (*string, bool)`

GetGtmIdOk returns a tuple with the GtmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGtmId

`func (o *ConfigResponse) SetGtmId(v string)`

SetGtmId sets GtmId field to given value.

### HasGtmId

`func (o *ConfigResponse) HasGtmId() bool`

HasGtmId returns a boolean if a field has been set.

### SetGtmIdNil

`func (o *ConfigResponse) SetGtmIdNil(b bool)`

 SetGtmIdNil sets the value for GtmId to be an explicit nil

### UnsetGtmId
`func (o *ConfigResponse) UnsetGtmId()`

UnsetGtmId ensures that no value is present for GtmId, not even an explicit nil
### GetZapierFeedbackWebhook

`func (o *ConfigResponse) GetZapierFeedbackWebhook() string`

GetZapierFeedbackWebhook returns the ZapierFeedbackWebhook field if non-nil, zero value otherwise.

### GetZapierFeedbackWebhookOk

`func (o *ConfigResponse) GetZapierFeedbackWebhookOk() (*string, bool)`

GetZapierFeedbackWebhookOk returns a tuple with the ZapierFeedbackWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZapierFeedbackWebhook

`func (o *ConfigResponse) SetZapierFeedbackWebhook(v string)`

SetZapierFeedbackWebhook sets ZapierFeedbackWebhook field to given value.

### HasZapierFeedbackWebhook

`func (o *ConfigResponse) HasZapierFeedbackWebhook() bool`

HasZapierFeedbackWebhook returns a boolean if a field has been set.

### SetZapierFeedbackWebhookNil

`func (o *ConfigResponse) SetZapierFeedbackWebhookNil(b bool)`

 SetZapierFeedbackWebhookNil sets the value for ZapierFeedbackWebhook to be an explicit nil

### UnsetZapierFeedbackWebhook
`func (o *ConfigResponse) UnsetZapierFeedbackWebhook()`

UnsetZapierFeedbackWebhook ensures that no value is present for ZapierFeedbackWebhook, not even an explicit nil
### GetZapierBugsWebhook

`func (o *ConfigResponse) GetZapierBugsWebhook() string`

GetZapierBugsWebhook returns the ZapierBugsWebhook field if non-nil, zero value otherwise.

### GetZapierBugsWebhookOk

`func (o *ConfigResponse) GetZapierBugsWebhookOk() (*string, bool)`

GetZapierBugsWebhookOk returns a tuple with the ZapierBugsWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZapierBugsWebhook

`func (o *ConfigResponse) SetZapierBugsWebhook(v string)`

SetZapierBugsWebhook sets ZapierBugsWebhook field to given value.

### HasZapierBugsWebhook

`func (o *ConfigResponse) HasZapierBugsWebhook() bool`

HasZapierBugsWebhook returns a boolean if a field has been set.

### SetZapierBugsWebhookNil

`func (o *ConfigResponse) SetZapierBugsWebhookNil(b bool)`

 SetZapierBugsWebhookNil sets the value for ZapierBugsWebhook to be an explicit nil

### UnsetZapierBugsWebhook
`func (o *ConfigResponse) UnsetZapierBugsWebhook()`

UnsetZapierBugsWebhook ensures that no value is present for ZapierBugsWebhook, not even an explicit nil
### GetBillingBackend

`func (o *ConfigResponse) GetBillingBackend() string`

GetBillingBackend returns the BillingBackend field if non-nil, zero value otherwise.

### GetBillingBackendOk

`func (o *ConfigResponse) GetBillingBackendOk() (*string, bool)`

GetBillingBackendOk returns a tuple with the BillingBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingBackend

`func (o *ConfigResponse) SetBillingBackend(v string)`

SetBillingBackend sets BillingBackend field to given value.

### HasBillingBackend

`func (o *ConfigResponse) HasBillingBackend() bool`

HasBillingBackend returns a boolean if a field has been set.

### SetBillingBackendNil

`func (o *ConfigResponse) SetBillingBackendNil(b bool)`

 SetBillingBackendNil sets the value for BillingBackend to be an explicit nil

### UnsetBillingBackend
`func (o *ConfigResponse) UnsetBillingBackend()`

UnsetBillingBackend ensures that no value is present for BillingBackend, not even an explicit nil
### GetMaintenance

`func (o *ConfigResponse) GetMaintenance() ConfigResponseMaintenance`

GetMaintenance returns the Maintenance field if non-nil, zero value otherwise.

### GetMaintenanceOk

`func (o *ConfigResponse) GetMaintenanceOk() (*ConfigResponseMaintenance, bool)`

GetMaintenanceOk returns a tuple with the Maintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenance

`func (o *ConfigResponse) SetMaintenance(v ConfigResponseMaintenance)`

SetMaintenance sets Maintenance field to given value.

### HasMaintenance

`func (o *ConfigResponse) HasMaintenance() bool`

HasMaintenance returns a boolean if a field has been set.

### SetMaintenanceNil

`func (o *ConfigResponse) SetMaintenanceNil(b bool)`

 SetMaintenanceNil sets the value for Maintenance to be an explicit nil

### UnsetMaintenance
`func (o *ConfigResponse) UnsetMaintenance()`

UnsetMaintenance ensures that no value is present for Maintenance, not even an explicit nil
### GetLogging

`func (o *ConfigResponse) GetLogging() Logging`

GetLogging returns the Logging field if non-nil, zero value otherwise.

### GetLoggingOk

`func (o *ConfigResponse) GetLoggingOk() (*Logging, bool)`

GetLoggingOk returns a tuple with the Logging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogging

`func (o *ConfigResponse) SetLogging(v Logging)`

SetLogging sets Logging field to given value.

### HasLogging

`func (o *ConfigResponse) HasLogging() bool`

HasLogging returns a boolean if a field has been set.

### GetCloudAdmin

`func (o *ConfigResponse) GetCloudAdmin() string`

GetCloudAdmin returns the CloudAdmin field if non-nil, zero value otherwise.

### GetCloudAdminOk

`func (o *ConfigResponse) GetCloudAdminOk() (*string, bool)`

GetCloudAdminOk returns a tuple with the CloudAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAdmin

`func (o *ConfigResponse) SetCloudAdmin(v string)`

SetCloudAdmin sets CloudAdmin field to given value.

### HasCloudAdmin

`func (o *ConfigResponse) HasCloudAdmin() bool`

HasCloudAdmin returns a boolean if a field has been set.

### SetCloudAdminNil

`func (o *ConfigResponse) SetCloudAdminNil(b bool)`

 SetCloudAdminNil sets the value for CloudAdmin to be an explicit nil

### UnsetCloudAdmin
`func (o *ConfigResponse) UnsetCloudAdmin()`

UnsetCloudAdmin ensures that no value is present for CloudAdmin, not even an explicit nil
### GetFilesAdmin

`func (o *ConfigResponse) GetFilesAdmin() string`

GetFilesAdmin returns the FilesAdmin field if non-nil, zero value otherwise.

### GetFilesAdminOk

`func (o *ConfigResponse) GetFilesAdminOk() (*string, bool)`

GetFilesAdminOk returns a tuple with the FilesAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesAdmin

`func (o *ConfigResponse) SetFilesAdmin(v string)`

SetFilesAdmin sets FilesAdmin field to given value.

### HasFilesAdmin

`func (o *ConfigResponse) HasFilesAdmin() bool`

HasFilesAdmin returns a boolean if a field has been set.

### SetFilesAdminNil

`func (o *ConfigResponse) SetFilesAdminNil(b bool)`

 SetFilesAdminNil sets the value for FilesAdmin to be an explicit nil

### UnsetFilesAdmin
`func (o *ConfigResponse) UnsetFilesAdmin()`

UnsetFilesAdmin ensures that no value is present for FilesAdmin, not even an explicit nil
### GetCloudDomain

`func (o *ConfigResponse) GetCloudDomain() string`

GetCloudDomain returns the CloudDomain field if non-nil, zero value otherwise.

### GetCloudDomainOk

`func (o *ConfigResponse) GetCloudDomainOk() (*string, bool)`

GetCloudDomainOk returns a tuple with the CloudDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudDomain

`func (o *ConfigResponse) SetCloudDomain(v string)`

SetCloudDomain sets CloudDomain field to given value.

### HasCloudDomain

`func (o *ConfigResponse) HasCloudDomain() bool`

HasCloudDomain returns a boolean if a field has been set.

### SetCloudDomainNil

`func (o *ConfigResponse) SetCloudDomainNil(b bool)`

 SetCloudDomainNil sets the value for CloudDomain to be an explicit nil

### UnsetCloudDomain
`func (o *ConfigResponse) UnsetCloudDomain()`

UnsetCloudDomain ensures that no value is present for CloudDomain, not even an explicit nil
### GetBillingDomain

`func (o *ConfigResponse) GetBillingDomain() string`

GetBillingDomain returns the BillingDomain field if non-nil, zero value otherwise.

### GetBillingDomainOk

`func (o *ConfigResponse) GetBillingDomainOk() (*string, bool)`

GetBillingDomainOk returns a tuple with the BillingDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingDomain

`func (o *ConfigResponse) SetBillingDomain(v string)`

SetBillingDomain sets BillingDomain field to given value.

### HasBillingDomain

`func (o *ConfigResponse) HasBillingDomain() bool`

HasBillingDomain returns a boolean if a field has been set.

### SetBillingDomainNil

`func (o *ConfigResponse) SetBillingDomainNil(b bool)`

 SetBillingDomainNil sets the value for BillingDomain to be an explicit nil

### UnsetBillingDomain
`func (o *ConfigResponse) UnsetBillingDomain()`

UnsetBillingDomain ensures that no value is present for BillingDomain, not even an explicit nil
### GetAppDomain

`func (o *ConfigResponse) GetAppDomain() string`

GetAppDomain returns the AppDomain field if non-nil, zero value otherwise.

### GetAppDomainOk

`func (o *ConfigResponse) GetAppDomainOk() (*string, bool)`

GetAppDomainOk returns a tuple with the AppDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDomain

`func (o *ConfigResponse) SetAppDomain(v string)`

SetAppDomain sets AppDomain field to given value.

### HasAppDomain

`func (o *ConfigResponse) HasAppDomain() bool`

HasAppDomain returns a boolean if a field has been set.

### SetAppDomainNil

`func (o *ConfigResponse) SetAppDomainNil(b bool)`

 SetAppDomainNil sets the value for AppDomain to be an explicit nil

### UnsetAppDomain
`func (o *ConfigResponse) UnsetAppDomain()`

UnsetAppDomain ensures that no value is present for AppDomain, not even an explicit nil
### GetMaxNetworkMonitorFileSize

`func (o *ConfigResponse) GetMaxNetworkMonitorFileSize() string`

GetMaxNetworkMonitorFileSize returns the MaxNetworkMonitorFileSize field if non-nil, zero value otherwise.

### GetMaxNetworkMonitorFileSizeOk

`func (o *ConfigResponse) GetMaxNetworkMonitorFileSizeOk() (*string, bool)`

GetMaxNetworkMonitorFileSizeOk returns a tuple with the MaxNetworkMonitorFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNetworkMonitorFileSize

`func (o *ConfigResponse) SetMaxNetworkMonitorFileSize(v string)`

SetMaxNetworkMonitorFileSize sets MaxNetworkMonitorFileSize field to given value.

### HasMaxNetworkMonitorFileSize

`func (o *ConfigResponse) HasMaxNetworkMonitorFileSize() bool`

HasMaxNetworkMonitorFileSize returns a boolean if a field has been set.

### SetMaxNetworkMonitorFileSizeNil

`func (o *ConfigResponse) SetMaxNetworkMonitorFileSizeNil(b bool)`

 SetMaxNetworkMonitorFileSizeNil sets the value for MaxNetworkMonitorFileSize to be an explicit nil

### UnsetMaxNetworkMonitorFileSize
`func (o *ConfigResponse) UnsetMaxNetworkMonitorFileSize()`

UnsetMaxNetworkMonitorFileSize ensures that no value is present for MaxNetworkMonitorFileSize, not even an explicit nil
### GetEnableFirmwareImageUpload

`func (o *ConfigResponse) GetEnableFirmwareImageUpload() bool`

GetEnableFirmwareImageUpload returns the EnableFirmwareImageUpload field if non-nil, zero value otherwise.

### GetEnableFirmwareImageUploadOk

`func (o *ConfigResponse) GetEnableFirmwareImageUploadOk() (*bool, bool)`

GetEnableFirmwareImageUploadOk returns a tuple with the EnableFirmwareImageUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFirmwareImageUpload

`func (o *ConfigResponse) SetEnableFirmwareImageUpload(v bool)`

SetEnableFirmwareImageUpload sets EnableFirmwareImageUpload field to given value.

### HasEnableFirmwareImageUpload

`func (o *ConfigResponse) HasEnableFirmwareImageUpload() bool`

HasEnableFirmwareImageUpload returns a boolean if a field has been set.

### SetEnableFirmwareImageUploadNil

`func (o *ConfigResponse) SetEnableFirmwareImageUploadNil(b bool)`

 SetEnableFirmwareImageUploadNil sets the value for EnableFirmwareImageUpload to be an explicit nil

### UnsetEnableFirmwareImageUpload
`func (o *ConfigResponse) UnsetEnableFirmwareImageUpload()`

UnsetEnableFirmwareImageUpload ensures that no value is present for EnableFirmwareImageUpload, not even an explicit nil
### GetAuthProviders

`func (o *ConfigResponse) GetAuthProviders() []AuthProvider`

GetAuthProviders returns the AuthProviders field if non-nil, zero value otherwise.

### GetAuthProvidersOk

`func (o *ConfigResponse) GetAuthProvidersOk() (*[]AuthProvider, bool)`

GetAuthProvidersOk returns a tuple with the AuthProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProviders

`func (o *ConfigResponse) SetAuthProviders(v []AuthProvider)`

SetAuthProviders sets AuthProviders field to given value.

### HasAuthProviders

`func (o *ConfigResponse) HasAuthProviders() bool`

HasAuthProviders returns a boolean if a field has been set.

### SetAuthProvidersNil

`func (o *ConfigResponse) SetAuthProvidersNil(b bool)`

 SetAuthProvidersNil sets the value for AuthProviders to be an explicit nil

### UnsetAuthProviders
`func (o *ConfigResponse) UnsetAuthProviders()`

UnsetAuthProviders ensures that no value is present for AuthProviders, not even an explicit nil
### GetSnapshotLimit

`func (o *ConfigResponse) GetSnapshotLimit() float32`

GetSnapshotLimit returns the SnapshotLimit field if non-nil, zero value otherwise.

### GetSnapshotLimitOk

`func (o *ConfigResponse) GetSnapshotLimitOk() (*float32, bool)`

GetSnapshotLimitOk returns a tuple with the SnapshotLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotLimit

`func (o *ConfigResponse) SetSnapshotLimit(v float32)`

SetSnapshotLimit sets SnapshotLimit field to given value.

### HasSnapshotLimit

`func (o *ConfigResponse) HasSnapshotLimit() bool`

HasSnapshotLimit returns a boolean if a field has been set.

### SetSnapshotLimitNil

`func (o *ConfigResponse) SetSnapshotLimitNil(b bool)`

 SetSnapshotLimitNil sets the value for SnapshotLimit to be an explicit nil

### UnsetSnapshotLimit
`func (o *ConfigResponse) UnsetSnapshotLimit()`

UnsetSnapshotLimit ensures that no value is present for SnapshotLimit, not even an explicit nil
### GetMaxKernelSize

`func (o *ConfigResponse) GetMaxKernelSize() float32`

GetMaxKernelSize returns the MaxKernelSize field if non-nil, zero value otherwise.

### GetMaxKernelSizeOk

`func (o *ConfigResponse) GetMaxKernelSizeOk() (*float32, bool)`

GetMaxKernelSizeOk returns a tuple with the MaxKernelSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxKernelSize

`func (o *ConfigResponse) SetMaxKernelSize(v float32)`

SetMaxKernelSize sets MaxKernelSize field to given value.

### HasMaxKernelSize

`func (o *ConfigResponse) HasMaxKernelSize() bool`

HasMaxKernelSize returns a boolean if a field has been set.

### SetMaxKernelSizeNil

`func (o *ConfigResponse) SetMaxKernelSizeNil(b bool)`

 SetMaxKernelSizeNil sets the value for MaxKernelSize to be an explicit nil

### UnsetMaxKernelSize
`func (o *ConfigResponse) UnsetMaxKernelSize()`

UnsetMaxKernelSize ensures that no value is present for MaxKernelSize, not even an explicit nil
### GetMaxRamdiskSize

`func (o *ConfigResponse) GetMaxRamdiskSize() float32`

GetMaxRamdiskSize returns the MaxRamdiskSize field if non-nil, zero value otherwise.

### GetMaxRamdiskSizeOk

`func (o *ConfigResponse) GetMaxRamdiskSizeOk() (*float32, bool)`

GetMaxRamdiskSizeOk returns a tuple with the MaxRamdiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRamdiskSize

`func (o *ConfigResponse) SetMaxRamdiskSize(v float32)`

SetMaxRamdiskSize sets MaxRamdiskSize field to given value.

### HasMaxRamdiskSize

`func (o *ConfigResponse) HasMaxRamdiskSize() bool`

HasMaxRamdiskSize returns a boolean if a field has been set.

### SetMaxRamdiskSizeNil

`func (o *ConfigResponse) SetMaxRamdiskSizeNil(b bool)`

 SetMaxRamdiskSizeNil sets the value for MaxRamdiskSize to be an explicit nil

### UnsetMaxRamdiskSize
`func (o *ConfigResponse) UnsetMaxRamdiskSize()`

UnsetMaxRamdiskSize ensures that no value is present for MaxRamdiskSize, not even an explicit nil
### GetCharmSDK

`func (o *ConfigResponse) GetCharmSDK() string`

GetCharmSDK returns the CharmSDK field if non-nil, zero value otherwise.

### GetCharmSDKOk

`func (o *ConfigResponse) GetCharmSDKOk() (*string, bool)`

GetCharmSDKOk returns a tuple with the CharmSDK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharmSDK

`func (o *ConfigResponse) SetCharmSDK(v string)`

SetCharmSDK sets CharmSDK field to given value.

### HasCharmSDK

`func (o *ConfigResponse) HasCharmSDK() bool`

HasCharmSDK returns a boolean if a field has been set.

### SetCharmSDKNil

`func (o *ConfigResponse) SetCharmSDKNil(b bool)`

 SetCharmSDKNil sets the value for CharmSDK to be an explicit nil

### UnsetCharmSDK
`func (o *ConfigResponse) UnsetCharmSDK()`

UnsetCharmSDK ensures that no value is present for CharmSDK, not even an explicit nil
### GetTrial

`func (o *ConfigResponse) GetTrial() Trial`

GetTrial returns the Trial field if non-nil, zero value otherwise.

### GetTrialOk

`func (o *ConfigResponse) GetTrialOk() (*Trial, bool)`

GetTrialOk returns a tuple with the Trial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrial

`func (o *ConfigResponse) SetTrial(v Trial)`

SetTrial sets Trial field to given value.

### HasTrial

`func (o *ConfigResponse) HasTrial() bool`

HasTrial returns a boolean if a field has been set.

### GetSentryUrl

`func (o *ConfigResponse) GetSentryUrl() string`

GetSentryUrl returns the SentryUrl field if non-nil, zero value otherwise.

### GetSentryUrlOk

`func (o *ConfigResponse) GetSentryUrlOk() (*string, bool)`

GetSentryUrlOk returns a tuple with the SentryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentryUrl

`func (o *ConfigResponse) SetSentryUrl(v string)`

SetSentryUrl sets SentryUrl field to given value.

### HasSentryUrl

`func (o *ConfigResponse) HasSentryUrl() bool`

HasSentryUrl returns a boolean if a field has been set.

### SetSentryUrlNil

`func (o *ConfigResponse) SetSentryUrlNil(b bool)`

 SetSentryUrlNil sets the value for SentryUrl to be an explicit nil

### UnsetSentryUrl
`func (o *ConfigResponse) UnsetSentryUrl()`

UnsetSentryUrl ensures that no value is present for SentryUrl, not even an explicit nil
### GetDomainAuthenticationProviders

`func (o *ConfigResponse) GetDomainAuthenticationProviders() bool`

GetDomainAuthenticationProviders returns the DomainAuthenticationProviders field if non-nil, zero value otherwise.

### GetDomainAuthenticationProvidersOk

`func (o *ConfigResponse) GetDomainAuthenticationProvidersOk() (*bool, bool)`

GetDomainAuthenticationProvidersOk returns a tuple with the DomainAuthenticationProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainAuthenticationProviders

`func (o *ConfigResponse) SetDomainAuthenticationProviders(v bool)`

SetDomainAuthenticationProviders sets DomainAuthenticationProviders field to given value.

### HasDomainAuthenticationProviders

`func (o *ConfigResponse) HasDomainAuthenticationProviders() bool`

HasDomainAuthenticationProviders returns a boolean if a field has been set.

### SetDomainAuthenticationProvidersNil

`func (o *ConfigResponse) SetDomainAuthenticationProvidersNil(b bool)`

 SetDomainAuthenticationProvidersNil sets the value for DomainAuthenticationProviders to be an explicit nil

### UnsetDomainAuthenticationProviders
`func (o *ConfigResponse) UnsetDomainAuthenticationProviders()`

UnsetDomainAuthenticationProviders ensures that no value is present for DomainAuthenticationProviders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


