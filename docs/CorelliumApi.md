# \CorelliumApi

All URIs are relative to *https://app.corellium.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAssessment**](CorelliumApi.md#CreateAssessment) | **Post** /v1/services/matrix/{instanceId}/assessments | Create assessment
[**DeleteAssessment**](CorelliumApi.md#DeleteAssessment) | **Delete** /v1/services/matrix/{instanceId}/assessments/{assessmentId} | Delete assessment
[**DownloadReport**](CorelliumApi.md#DownloadReport) | **Get** /v1/services/matrix/{instanceId}/assessments/{assessmentId}/download | Download report
[**GetAssessmentById**](CorelliumApi.md#GetAssessmentById) | **Get** /v1/services/matrix/{instanceId}/assessments/{assessmentId} | Get assessment by ID
[**GetAssessmentsByInstanceId**](CorelliumApi.md#GetAssessmentsByInstanceId) | **Get** /v1/services/matrix/{instanceId}/instances/{instanceId}/assessments | Get assessments by instanceId
[**RunTests**](CorelliumApi.md#RunTests) | **Post** /v1/services/matrix/{instanceId}/assessments/{assessmentId}/test | Update assessment state and execute MATRIX tests
[**StartMonitoring**](CorelliumApi.md#StartMonitoring) | **Post** /v1/services/matrix/{instanceId}/assessments/{assessmentId}/start | Update assessment state and begin device monitoring
[**StopMonitoring**](CorelliumApi.md#StopMonitoring) | **Post** /v1/services/matrix/{instanceId}/assessments/{assessmentId}/stop | Update assessment state and stop device monitoring
[**V1AcceptSharedSnapshot**](CorelliumApi.md#V1AcceptSharedSnapshot) | **Post** /v1/snapshots/accept | Accept a snapshot shared with you
[**V1ActivityExport**](CorelliumApi.md#V1ActivityExport) | **Post** /v1/activity/export | Start activity export
[**V1ActivityList**](CorelliumApi.md#V1ActivityList) | **Get** /v1/activity | Get resource activities
[**V1AddProjectKey**](CorelliumApi.md#V1AddProjectKey) | **Post** /v1/projects/{projectId}/keys | Add Project Authorized Key
[**V1AddTeamRoleToProject**](CorelliumApi.md#V1AddTeamRoleToProject) | **Put** /v1/roles/projects/{projectId}/teams/{teamId}/roles/{roleId} | Add team role to project
[**V1AddUserRoleToProject**](CorelliumApi.md#V1AddUserRoleToProject) | **Put** /v1/roles/projects/{projectId}/users/{userId}/roles/{roleId} | Add user role to project
[**V1AddUserToTeam**](CorelliumApi.md#V1AddUserToTeam) | **Put** /v1/teams/{teamId}/users/{userId} | Add user to team
[**V1AgentAppReady**](CorelliumApi.md#V1AgentAppReady) | **Get** /v1/instances/{instanceId}/agent/v1/app/ready | Check if App subsystem is ready
[**V1AgentDeleteFile**](CorelliumApi.md#V1AgentDeleteFile) | **Delete** /v1/instances/{instanceId}/agent/v1/file/device/{filePath} | Delete a File on VM
[**V1AgentGetFile**](CorelliumApi.md#V1AgentGetFile) | **Get** /v1/instances/{instanceId}/agent/v1/file/device/{filePath} | Download a File from VM
[**V1AgentGetTempFilename**](CorelliumApi.md#V1AgentGetTempFilename) | **Post** /v1/instances/{instanceId}/agent/v1/file/temp | Get the path for a new temp file
[**V1AgentInstallApp**](CorelliumApi.md#V1AgentInstallApp) | **Post** /v1/instances/{instanceId}/agent/v1/app/install | Install App at path
[**V1AgentInstallProfile**](CorelliumApi.md#V1AgentInstallProfile) | **Post** /v1/instances/{instanceId}/agent/v1/profile/install | Install a Profile to VM
[**V1AgentKillApp**](CorelliumApi.md#V1AgentKillApp) | **Post** /v1/instances/{instanceId}/agent/v1/app/apps/{bundleId}/kill | Kill an App
[**V1AgentListAppIcons**](CorelliumApi.md#V1AgentListAppIcons) | **Get** /v1/instances/{instanceId}/agent/v1/app/icons | List App Icons
[**V1AgentListApps**](CorelliumApi.md#V1AgentListApps) | **Get** /v1/instances/{instanceId}/agent/v1/app/apps | List Apps
[**V1AgentListAppsStatus**](CorelliumApi.md#V1AgentListAppsStatus) | **Get** /v1/instances/{instanceId}/agent/v1/app/apps/update | List Apps Status
[**V1AgentListProfiles**](CorelliumApi.md#V1AgentListProfiles) | **Get** /v1/instances/{instanceId}/agent/v1/profile/profiles | List Profiles
[**V1AgentRunApp**](CorelliumApi.md#V1AgentRunApp) | **Post** /v1/instances/{instanceId}/agent/v1/app/apps/{bundleId}/run | Run an App
[**V1AgentSetFileAttributes**](CorelliumApi.md#V1AgentSetFileAttributes) | **Patch** /v1/instances/{instanceId}/agent/v1/file/device/{filePath} | Change Attrs of a File on VM
[**V1AgentSystemGetAdbAuth**](CorelliumApi.md#V1AgentSystemGetAdbAuth) | **Get** /v1/instances/{instanceId}/agent/v1/system/adbauth | Get ADB Auth Setting (AOSP only)
[**V1AgentSystemGetNetwork**](CorelliumApi.md#V1AgentSystemGetNetwork) | **Get** /v1/instances/{instanceId}/agent/v1/system/network | Get IP of eth0 (AOSP only)
[**V1AgentSystemGetProp**](CorelliumApi.md#V1AgentSystemGetProp) | **Post** /v1/instances/{instanceId}/agent/v1/system/getprop | Get System Property (AOSP only)
[**V1AgentSystemInstallOpenGApps**](CorelliumApi.md#V1AgentSystemInstallOpenGApps) | **Post** /v1/instances/{instanceId}/agent/v1/system/install-opengapps | Install OpenGApps (AOSP only)
[**V1AgentSystemLock**](CorelliumApi.md#V1AgentSystemLock) | **Post** /v1/instances/{instanceId}/agent/v1/system/lock | Lock Device (iOS Only)
[**V1AgentSystemSetAdbAuth**](CorelliumApi.md#V1AgentSystemSetAdbAuth) | **Put** /v1/instances/{instanceId}/agent/v1/system/adbauth | Set ADB Auth Setting (AOSP only)
[**V1AgentSystemSetHostname**](CorelliumApi.md#V1AgentSystemSetHostname) | **Post** /v1/instances/{instanceId}/agent/v1/system/setHostname | Set Hostname of instance
[**V1AgentSystemShutdown**](CorelliumApi.md#V1AgentSystemShutdown) | **Post** /v1/instances/{instanceId}/agent/v1/system/shutdown | Instruct VM to halt
[**V1AgentSystemUnlock**](CorelliumApi.md#V1AgentSystemUnlock) | **Post** /v1/instances/{instanceId}/agent/v1/system/unlock | Unlock Device (iOS Only)
[**V1AgentUninstallApp**](CorelliumApi.md#V1AgentUninstallApp) | **Post** /v1/instances/{instanceId}/agent/v1/app/apps/{bundleId}/uninstall | Uninstall an App
[**V1AgentUninstallProfile**](CorelliumApi.md#V1AgentUninstallProfile) | **Delete** /v1/instances/{instanceId}/agent/v1/profile/profiles/{profileId} | Uninstall a Profile from VM
[**V1AgentUploadFile**](CorelliumApi.md#V1AgentUploadFile) | **Put** /v1/instances/{instanceId}/agent/v1/file/device/{filePath} | Upload a file to VM
[**V1AuthLogin**](CorelliumApi.md#V1AuthLogin) | **Post** /v1/auth/login | Log In
[**V1BtracePreauthorize**](CorelliumApi.md#V1BtracePreauthorize) | **Get** /v1/instances/{instanceId}/btrace-authorize | Pre-authorize an btrace download
[**V1CheckSubdomainExistence**](CorelliumApi.md#V1CheckSubdomainExistence) | **Post** /v1/domain/check | Check the existence of a subdomain
[**V1ClearCoreTrace**](CorelliumApi.md#V1ClearCoreTrace) | **Delete** /v1/instances/{instanceId}/strace | Clear CoreTrace logs
[**V1ClearHyperTrace**](CorelliumApi.md#V1ClearHyperTrace) | **Delete** /v1/instances/{instanceId}/btrace | Clear HyperTrace logs
[**V1ClearHyperTraceHooks**](CorelliumApi.md#V1ClearHyperTraceHooks) | **Post** /v1/instances/{instanceId}/hooks/clear | Clear Hooks on an instance
[**V1ClearInstancePanics**](CorelliumApi.md#V1ClearInstancePanics) | **Delete** /v1/instances/{instanceId}/panics | Clear Panics
[**V1CreateDomainAuthProvider**](CorelliumApi.md#V1CreateDomainAuthProvider) | **Post** /v1/domain/{domainId}/auth | Create a new auth provider for a domain
[**V1CreateHook**](CorelliumApi.md#V1CreateHook) | **Post** /v1/instances/{instanceId}/hooks | Create hypervisor hook for Instance
[**V1CreateImage**](CorelliumApi.md#V1CreateImage) | **Post** /v1/images | Create a new Image
[**V1CreateInstance**](CorelliumApi.md#V1CreateInstance) | **Post** /v1/instances | Create Instance
[**V1CreateNetworkConnection**](CorelliumApi.md#V1CreateNetworkConnection) | **Post** /v1/network/connections | Create a new Network Connection
[**V1CreateProject**](CorelliumApi.md#V1CreateProject) | **Post** /v1/projects | Create a Project
[**V1CreateSnapshot**](CorelliumApi.md#V1CreateSnapshot) | **Post** /v1/instances/{instanceId}/snapshots | Create Instance Snapshot
[**V1CreateUser**](CorelliumApi.md#V1CreateUser) | **Post** /v1/users | Create User
[**V1DeleteDomainAuthProvider**](CorelliumApi.md#V1DeleteDomainAuthProvider) | **Delete** /v1/domain/{domainId}/auth/{providerId} | Delete an auth provider from a domain
[**V1DeleteExtension**](CorelliumApi.md#V1DeleteExtension) | **Delete** /v1/extensions/{extensionId} | Delete an existing extension
[**V1DeleteHook**](CorelliumApi.md#V1DeleteHook) | **Delete** /v1/hooks/{hookId} | Delete an existing hypervisor hook
[**V1DeleteImage**](CorelliumApi.md#V1DeleteImage) | **Delete** /v2/images/{imageId} | Delete Image
[**V1DeleteInstance**](CorelliumApi.md#V1DeleteInstance) | **Delete** /v1/instances/{instanceId} | Remove Instance
[**V1DeleteInstanceSnapshot**](CorelliumApi.md#V1DeleteInstanceSnapshot) | **Delete** /v1/instances/{instanceId}/snapshots/{snapshotId} | Delete an Instance Snapshot
[**V1DeleteNetworkConnection**](CorelliumApi.md#V1DeleteNetworkConnection) | **Delete** /v1/network/connections/{id} | Delete an existing Network Connection
[**V1DeleteProject**](CorelliumApi.md#V1DeleteProject) | **Delete** /v1/projects/{projectId} | Delete a Project
[**V1DeleteSnapshot**](CorelliumApi.md#V1DeleteSnapshot) | **Delete** /v1/snapshots/{snapshotId} | Delete a Snapshot
[**V1DeleteSnapshotPermissions**](CorelliumApi.md#V1DeleteSnapshotPermissions) | **Delete** /v1/snapshots/{snapshotId}/permissions | Delete member(s)
[**V1DeleteUser**](CorelliumApi.md#V1DeleteUser) | **Delete** /v1/users/{userId} | Delete User
[**V1DisableExposePort**](CorelliumApi.md#V1DisableExposePort) | **Post** /v1/instances/{instanceId}/exposeport/disable | Disable an exposed port on an instance for device access.
[**V1DownloadActivity**](CorelliumApi.md#V1DownloadActivity) | **Get** /v1/activity/export/{taskId}/download | Download activity
[**V1EnableExposePort**](CorelliumApi.md#V1EnableExposePort) | **Post** /v1/instances/{instanceId}/exposeport/enable | Enable an exposed port on an instance for device access.
[**V1ExecuteHyperTraceHooks**](CorelliumApi.md#V1ExecuteHyperTraceHooks) | **Post** /v1/instances/{instanceId}/hooks/execute | Execute Hooks on an instance
[**V1GetActivityExportStatus**](CorelliumApi.md#V1GetActivityExportStatus) | **Get** /v1/activity/export/{taskId} | Get export task status
[**V1GetActivityExportTasks**](CorelliumApi.md#V1GetActivityExportTasks) | **Get** /v1/activity/export | Get all export tasks for user
[**V1GetConfig**](CorelliumApi.md#V1GetConfig) | **Get** /v1/config | Get all configs
[**V1GetDomainAuthProviders**](CorelliumApi.md#V1GetDomainAuthProviders) | **Get** /v1/domain/{domainId}/auth | Return all configured auth providers for a domain (including globally configured providers)
[**V1GetExtensionById**](CorelliumApi.md#V1GetExtensionById) | **Get** /v1/extensions/{extensionId} | Get extension by id
[**V1GetExtensions**](CorelliumApi.md#V1GetExtensions) | **Get** /v1/extensions | Get all extensions
[**V1GetHookById**](CorelliumApi.md#V1GetHookById) | **Get** /v1/hooks/{hookId} | Get hypervisor hook by id
[**V1GetHooks**](CorelliumApi.md#V1GetHooks) | **Get** /v1/instances/{instanceId}/hooks | Get all hypervisor hooks for Instance
[**V1GetImage**](CorelliumApi.md#V1GetImage) | **Get** /v1/images/{imageId} | Get Image Metadata
[**V1GetImages**](CorelliumApi.md#V1GetImages) | **Get** /v1/images | Get all Images Metadata
[**V1GetInstance**](CorelliumApi.md#V1GetInstance) | **Get** /v1/instances/{instanceId} | Get Instance
[**V1GetInstanceConsole**](CorelliumApi.md#V1GetInstanceConsole) | **Get** /v1/instances/{instanceId}/console | Get console websocket URL
[**V1GetInstanceConsoleLog**](CorelliumApi.md#V1GetInstanceConsoleLog) | **Get** /v1/instances/{instanceId}/consoleLog | Get Console Log
[**V1GetInstanceGpios**](CorelliumApi.md#V1GetInstanceGpios) | **Get** /v1/instances/{instanceId}/gpios | Get Instance GPIOs
[**V1GetInstancePanics**](CorelliumApi.md#V1GetInstancePanics) | **Get** /v1/instances/{instanceId}/panics | Get Panics
[**V1GetInstancePeripherals**](CorelliumApi.md#V1GetInstancePeripherals) | **Get** /v1/instances/{instanceId}/peripherals | Get Instance Peripherals
[**V1GetInstanceScreenshot**](CorelliumApi.md#V1GetInstanceScreenshot) | **Get** /v1/instances/{instanceId}/screenshot.{format} | Get Instance Screenshot
[**V1GetInstanceSnapshot**](CorelliumApi.md#V1GetInstanceSnapshot) | **Get** /v1/instances/{instanceId}/snapshots/{snapshotId} | Get Instance Snapshot
[**V1GetInstanceSnapshots**](CorelliumApi.md#V1GetInstanceSnapshots) | **Get** /v1/instances/{instanceId}/snapshots | Get Instance Snapshots
[**V1GetInstances**](CorelliumApi.md#V1GetInstances) | **Get** /v1/instances | Get Instances
[**V1GetModelSoftware**](CorelliumApi.md#V1GetModelSoftware) | **Get** /v1/models/{model}/software | Get Software for Model
[**V1GetModels**](CorelliumApi.md#V1GetModels) | **Get** /v1/models | Get Supported Models
[**V1GetNetworkConnection**](CorelliumApi.md#V1GetNetworkConnection) | **Get** /v1/network/connections/{id} | Return the configuration and per project statuses for a single network provider.
[**V1GetProject**](CorelliumApi.md#V1GetProject) | **Get** /v1/projects/{projectId} | Get a Project
[**V1GetProjectInstances**](CorelliumApi.md#V1GetProjectInstances) | **Get** /v1/projects/{projectId}/instances | Get Instances in Project
[**V1GetProjectKeys**](CorelliumApi.md#V1GetProjectKeys) | **Get** /v1/projects/{projectId}/keys | Get Project Authorized Keys
[**V1GetProjectNetworkLog**](CorelliumApi.md#V1GetProjectNetworkLog) | **Get** /v1/projects/{projectId}/network/log | Retrieve the network connection log for a project
[**V1GetProjectNetworkStatus**](CorelliumApi.md#V1GetProjectNetworkStatus) | **Get** /v1/projects/{projectId}/network/status | Retrieve the network connection status for a project
[**V1GetProjectVpnConfig**](CorelliumApi.md#V1GetProjectVpnConfig) | **Get** /v1/projects/{projectId}/vpnconfig/{format} | Get Project VPN Configuration
[**V1GetProjects**](CorelliumApi.md#V1GetProjects) | **Get** /v1/projects | Get Projects
[**V1GetResetLinkInfo**](CorelliumApi.md#V1GetResetLinkInfo) | **Get** /v1/users/reset-link-info | Send Password Reset Link Info
[**V1GetSharedSnapshots**](CorelliumApi.md#V1GetSharedSnapshots) | **Get** /v1/snapshots/shared | Fetch shared snapshots
[**V1GetSnapshot**](CorelliumApi.md#V1GetSnapshot) | **Get** /v1/snapshots/{snapshotId} | Get Snapshot
[**V1InstancesInstanceIdMessagePost**](CorelliumApi.md#V1InstancesInstanceIdMessagePost) | **Post** /v1/instances/{instanceId}/message | Inject a message into an iOS VM
[**V1InstancesInstanceIdNetdumpPcapGet**](CorelliumApi.md#V1InstancesInstanceIdNetdumpPcapGet) | **Get** /v1/instances/{instanceId}/netdump.pcap | Download a netdump pcap file
[**V1InstancesInstanceIdNetworkMonitorPcapGet**](CorelliumApi.md#V1InstancesInstanceIdNetworkMonitorPcapGet) | **Get** /v1/instances/{instanceId}/networkMonitor.pcap | Download a Network Monitor pcap file
[**V1Kcrange**](CorelliumApi.md#V1Kcrange) | **Get** /v1/instances/{instanceId}/btrace-kcrange | Get Kernel extension ranges
[**V1ListNetworkConnections**](CorelliumApi.md#V1ListNetworkConnections) | **Get** /v1/network/connections | List available network connections
[**V1ListNetworkInterfaces**](CorelliumApi.md#V1ListNetworkInterfaces) | **Get** /v1/network/interfaces | List available physical network interfaces
[**V1ListNetworkProviders**](CorelliumApi.md#V1ListNetworkProviders) | **Get** /v1/network/providers | List available network providers
[**V1ListThreads**](CorelliumApi.md#V1ListThreads) | **Get** /v1/instances/{instanceId}/strace/thread-list | Get Running Threads (CoreTrace)
[**V1LoadExtension**](CorelliumApi.md#V1LoadExtension) | **Post** /v1/extensions | Load an extension
[**V1MediaPlay**](CorelliumApi.md#V1MediaPlay) | **Post** /v1/instances/{instanceId}/media/play | Start playing media
[**V1MediaStop**](CorelliumApi.md#V1MediaStop) | **Post** /v1/instances/{instanceId}/media/stop | Stop playing media
[**V1ParseExtensionJson**](CorelliumApi.md#V1ParseExtensionJson) | **Post** /v1/extensions/parse/extension.json | Validates extension.json
[**V1PartialUpdateNetworkConnection**](CorelliumApi.md#V1PartialUpdateNetworkConnection) | **Patch** /v1/network/connections/{id} | Update Network Connection (partial)
[**V1PatchInstance**](CorelliumApi.md#V1PatchInstance) | **Patch** /v1/instances/{instanceId} | Update Instance
[**V1PatchInstanceReadOnly**](CorelliumApi.md#V1PatchInstanceReadOnly) | **Patch** /v1/instances/{instanceId}/read-only | Update Instance Read Only
[**V1PauseInstance**](CorelliumApi.md#V1PauseInstance) | **Post** /v1/instances/{instanceId}/pause | Pause an Instance
[**V1PostInstanceInput**](CorelliumApi.md#V1PostInstanceInput) | **Post** /v1/instances/{instanceId}/input | Provide Instance Input
[**V1Ready**](CorelliumApi.md#V1Ready) | **Get** /v1/ready | API Status
[**V1RebootInstance**](CorelliumApi.md#V1RebootInstance) | **Post** /v1/instances/{instanceId}/reboot | Reboot an Instance
[**V1RemoveProjectKey**](CorelliumApi.md#V1RemoveProjectKey) | **Delete** /v1/projects/{projectId}/keys/{keyId} | Delete Project Authorized Key
[**V1RemoveTeamRoleFromProject**](CorelliumApi.md#V1RemoveTeamRoleFromProject) | **Delete** /v1/roles/projects/{projectId}/teams/{teamId}/roles/{roleId} | Remove team role from project
[**V1RemoveUserFromTeam**](CorelliumApi.md#V1RemoveUserFromTeam) | **Delete** /v1/teams/{teamId}/users/{userId} | Remove user from team
[**V1RemoveUserRoleFromProject**](CorelliumApi.md#V1RemoveUserRoleFromProject) | **Delete** /v1/roles/projects/{projectId}/users/{userId}/roles/{roleId} | Remove user role from project
[**V1RenameInstanceSnapshot**](CorelliumApi.md#V1RenameInstanceSnapshot) | **Patch** /v1/instances/{instanceId}/snapshots/{snapshotId} | Rename an Instance Snapshot
[**V1ResetUserPassword**](CorelliumApi.md#V1ResetUserPassword) | **Post** /v1/users/reset-password | Reset User Password
[**V1RestoreBackup**](CorelliumApi.md#V1RestoreBackup) | **Post** /v1/instances/{instanceId}/restoreBackup | Restore backup
[**V1RestoreInstanceSnapshot**](CorelliumApi.md#V1RestoreInstanceSnapshot) | **Post** /v1/instances/{instanceId}/snapshots/{snapshotId}/restore | Restore an Instance Snapshot
[**V1Roles**](CorelliumApi.md#V1Roles) | **Get** /v1/roles | Get all roles
[**V1RotateInstance**](CorelliumApi.md#V1RotateInstance) | **Post** /v1/instances/{instanceId}/rotate | Rotate device to specified orientation
[**V1SendUserResetLink**](CorelliumApi.md#V1SendUserResetLink) | **Post** /v1/users/send-reset-link | Send Password Reset Link
[**V1SetInstanceGpios**](CorelliumApi.md#V1SetInstanceGpios) | **Put** /v1/instances/{instanceId}/gpios | Set Instance GPIOs
[**V1SetInstancePeripherals**](CorelliumApi.md#V1SetInstancePeripherals) | **Put** /v1/instances/{instanceId}/peripherals | Set Instance Peripherals
[**V1SetInstanceState**](CorelliumApi.md#V1SetInstanceState) | **Put** /v1/instances/{instanceId}/state | Set state of Instance
[**V1SetSnapshotPermissions**](CorelliumApi.md#V1SetSnapshotPermissions) | **Post** /v1/snapshots/{snapshotId}/permissions | Set member list
[**V1ShareSnapshot**](CorelliumApi.md#V1ShareSnapshot) | **Post** /v1/snapshots/{snapshotId}/share | Share snapshot
[**V1SnapshotRename**](CorelliumApi.md#V1SnapshotRename) | **Patch** /v1/snapshots/{snapshotId} | Rename a Snapshot
[**V1StartCoreTrace**](CorelliumApi.md#V1StartCoreTrace) | **Post** /v1/instances/{instanceId}/strace/enable | Start CoreTrace on an instance
[**V1StartHyperTrace**](CorelliumApi.md#V1StartHyperTrace) | **Post** /v1/instances/{instanceId}/btrace/enable | Start HyperTrace on an instance
[**V1StartInstance**](CorelliumApi.md#V1StartInstance) | **Post** /v1/instances/{instanceId}/start | Start an Instance
[**V1StartNetdump**](CorelliumApi.md#V1StartNetdump) | **Post** /v1/instances/{instanceId}/netdump/enable | Start Enhanced Network Monitor on an instance.
[**V1StartNetworkMonitor**](CorelliumApi.md#V1StartNetworkMonitor) | **Post** /v1/instances/{instanceId}/sslsplit/enable | Start Network Monitor on an instance.
[**V1StopCoreTrace**](CorelliumApi.md#V1StopCoreTrace) | **Post** /v1/instances/{instanceId}/strace/disable | Stop CoreTrace on an instance.
[**V1StopHyperTrace**](CorelliumApi.md#V1StopHyperTrace) | **Post** /v1/instances/{instanceId}/btrace/disable | Stop HyperTrace on an instance.
[**V1StopInstance**](CorelliumApi.md#V1StopInstance) | **Post** /v1/instances/{instanceId}/stop | Stop an Instance
[**V1StopNetdump**](CorelliumApi.md#V1StopNetdump) | **Post** /v1/instances/{instanceId}/netdump/disable | Stop Enhanced Network Monitor on an instance.
[**V1StopNetworkMonitor**](CorelliumApi.md#V1StopNetworkMonitor) | **Post** /v1/instances/{instanceId}/sslsplit/disable | Stop Network Monitor on an instance.
[**V1TeamChange**](CorelliumApi.md#V1TeamChange) | **Patch** /v1/teams/{teamId} | Update team
[**V1TeamCreate**](CorelliumApi.md#V1TeamCreate) | **Post** /v1/teams | Create team
[**V1TeamDelete**](CorelliumApi.md#V1TeamDelete) | **Delete** /v1/teams/{teamId} | Delete team
[**V1Teams**](CorelliumApi.md#V1Teams) | **Get** /v1/teams | Get teams
[**V1UnpauseInstance**](CorelliumApi.md#V1UnpauseInstance) | **Post** /v1/instances/{instanceId}/unpause | Unpause an Instance
[**V1UpdateDomainAuthProvider**](CorelliumApi.md#V1UpdateDomainAuthProvider) | **Put** /v1/domain/{domainId}/auth/{providerId} | Update an auth provider for a domain
[**V1UpdateExtension**](CorelliumApi.md#V1UpdateExtension) | **Put** /v1/extensions/{extensionId} | Update an existing extension
[**V1UpdateHook**](CorelliumApi.md#V1UpdateHook) | **Put** /v1/hooks/{hookId} | Update an existing hypervisor hook
[**V1UpdateNetworkConnection**](CorelliumApi.md#V1UpdateNetworkConnection) | **Put** /v1/network/connections/{id} | Update Network Connection
[**V1UpdateProject**](CorelliumApi.md#V1UpdateProject) | **Patch** /v1/projects/{projectId} | Update a Project
[**V1UpdateProjectSettings**](CorelliumApi.md#V1UpdateProjectSettings) | **Patch** /v1/projects/{projectId}/settings | Change Project Settings
[**V1UpdateUser**](CorelliumApi.md#V1UpdateUser) | **Patch** /v1/users/{userId} | Update User
[**V1UpgradeInstance**](CorelliumApi.md#V1UpgradeInstance) | **Post** /v1/instances/{instanceId}/upgrade | Upgrade iOS version
[**V1UploadImageData**](CorelliumApi.md#V1UploadImageData) | **Post** /v1/images/{imageId} | Upload Image Data
[**V1UserAgreeTerms**](CorelliumApi.md#V1UserAgreeTerms) | **Post** /v1/users/agree | Consent to the current terms and conditions
[**V1UsersChangePasswordPost**](CorelliumApi.md#V1UsersChangePasswordPost) | **Post** /v1/users/change-password | Change User Password
[**V1UsersLogin**](CorelliumApi.md#V1UsersLogin) | **Post** /v1/users/login | Log In
[**V1WebPlayerAllowedDomains**](CorelliumApi.md#V1WebPlayerAllowedDomains) | **Get** /v1/webplayer/allowedDomains | Retrieve the list of allowed domains for all Webplayer sessions
[**V1WebPlayerCreateSession**](CorelliumApi.md#V1WebPlayerCreateSession) | **Post** /v1/webplayer | Create a new Webplayer Session
[**V1WebPlayerDestroySession**](CorelliumApi.md#V1WebPlayerDestroySession) | **Delete** /v1/webplayer/{sessionId} | Tear down a Webplayer Session
[**V1WebPlayerListSessions**](CorelliumApi.md#V1WebPlayerListSessions) | **Get** /v1/webplayer | List all Webplayer sessions
[**V1WebPlayerSessionInfo**](CorelliumApi.md#V1WebPlayerSessionInfo) | **Get** /v1/webplayer/{sessionId} | Retrieve Webplayer Session Information
[**V2GetInstanceQuickConnectCommand**](CorelliumApi.md#V2GetInstanceQuickConnectCommand) | **Get** /v2/instances/{instanceId}/quickConnectCommand | Recommended SSH Command for Quick Connect
[**V2GetInstanceState**](CorelliumApi.md#V2GetInstanceState) | **Get** /v2/instances/{instanceId}/state | Get state of Instance



## CreateAssessment

> AssessmentIdStatus CreateAssessment(ctx, instanceId).CreateAssessmentDto(createAssessmentDto).Execute()

Create assessment

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of instance
    createAssessmentDto := *openapiclient.NewCreateAssessmentDto("InstanceId_example", "com.android.egg") // CreateAssessmentDto | Create a new assessment

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.CreateAssessment(context.Background(), instanceId).CreateAssessmentDto(createAssessmentDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.CreateAssessment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAssessment`: AssessmentIdStatus
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.CreateAssessment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | ID of instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssessmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAssessmentDto** | [**CreateAssessmentDto**](CreateAssessmentDto.md) | Create a new assessment | 

### Return type

[**AssessmentIdStatus**](AssessmentIdStatus.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAssessment

> DeleteAssessment(ctx, instanceId, assessmentId).Execute()

Delete assessment

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of instance
    assessmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of assessment

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.DeleteAssessment(context.Background(), instanceId, assessmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.DeleteAssessment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | ID of instance | 
**assessmentId** | **string** | ID of assessment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAssessmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadReport

> string DownloadReport(ctx, instanceId, assessmentId).Format(format).Execute()

Download report

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of instance
    assessmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of assessment
    format := "format_example" // string | Assessment report download format (default to "html")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.DownloadReport(context.Background(), instanceId, assessmentId).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.DownloadReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadReport`: string
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.DownloadReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | ID of instance | 
**assessmentId** | **string** | ID of assessment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **string** | Assessment report download format | [default to &quot;html&quot;]

### Return type

**string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssessmentById

> Assessment GetAssessmentById(ctx, instanceId, assessmentId).Execute()

Get assessment by ID

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of instance
    assessmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of assessment

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.GetAssessmentById(context.Background(), instanceId, assessmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.GetAssessmentById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssessmentById`: Assessment
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.GetAssessmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | ID of instance | 
**assessmentId** | **string** | ID of assessment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssessmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Assessment**](Assessment.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssessmentsByInstanceId

> []Assessment GetAssessmentsByInstanceId(ctx, instanceId).Execute()

Get assessments by instanceId

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of instance

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.GetAssessmentsByInstanceId(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.GetAssessmentsByInstanceId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssessmentsByInstanceId`: []Assessment
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.GetAssessmentsByInstanceId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | ID of instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssessmentsByInstanceIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Assessment**](Assessment.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunTests

> AssessmentIdStatus RunTests(ctx, instanceId, assessmentId).TestAssessmentDto(testAssessmentDto).Execute()

Update assessment state and execute MATRIX tests

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of instance
    assessmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of assessment
    testAssessmentDto := *openapiclient.NewTestAssessmentDto() // TestAssessmentDto | Execute MATRIX tests (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.RunTests(context.Background(), instanceId, assessmentId).TestAssessmentDto(testAssessmentDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.RunTests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunTests`: AssessmentIdStatus
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.RunTests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | ID of instance | 
**assessmentId** | **string** | ID of assessment | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **testAssessmentDto** | [**TestAssessmentDto**](TestAssessmentDto.md) | Execute MATRIX tests | 

### Return type

[**AssessmentIdStatus**](AssessmentIdStatus.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartMonitoring

> StartMonitoring(ctx, instanceId, assessmentId).Execute()

Update assessment state and begin device monitoring

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of instance
    assessmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of assessment

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.StartMonitoring(context.Background(), instanceId, assessmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.StartMonitoring``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | ID of instance | 
**assessmentId** | **string** | ID of assessment | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartMonitoringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopMonitoring

> StopMonitoring(ctx, instanceId, assessmentId).Execute()

Update assessment state and stop device monitoring

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of instance
    assessmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of assessment

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.StopMonitoring(context.Background(), instanceId, assessmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.StopMonitoring``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | ID of instance | 
**assessmentId** | **string** | ID of assessment | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopMonitoringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AcceptSharedSnapshot

> Snapshot V1AcceptSharedSnapshot(ctx).PostShareSnapshotRequestPayload(postShareSnapshotRequestPayload).Execute()

Accept a snapshot shared with you



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    postShareSnapshotRequestPayload := *openapiclient.NewPostShareSnapshotRequestPayload("SharingType_example") // PostShareSnapshotRequestPayload | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1AcceptSharedSnapshot(context.Background()).PostShareSnapshotRequestPayload(postShareSnapshotRequestPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1AcceptSharedSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AcceptSharedSnapshot`: Snapshot
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1AcceptSharedSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AcceptSharedSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postShareSnapshotRequestPayload** | [**PostShareSnapshotRequestPayload**](PostShareSnapshotRequestPayload.md) |  | 

### Return type

[**Snapshot**](Snapshot.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ActivityExport

> ActivityExportResponse V1ActivityExport(ctx).ActivityExportDto(activityExportDto).Execute()

Start activity export

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    activityExportDto := *openapiclient.NewActivityExportDto() // ActivityExportDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1ActivityExport(context.Background()).ActivityExportDto(activityExportDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1ActivityExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ActivityExport`: ActivityExportResponse
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1ActivityExport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ActivityExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activityExportDto** | [**ActivityExportDto**](ActivityExportDto.md) |  | 

### Return type

[**ActivityExportResponse**](ActivityExportResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ActivityList

> ActivityList V1ActivityList(ctx).Request(request).Execute()

Get resource activities



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    request := map[string][]openapiclient.ActivityRequest{ ... } // ActivityRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1ActivityList(context.Background()).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1ActivityList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ActivityList`: ActivityList
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1ActivityList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ActivityListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**ActivityRequest**](ActivityRequest.md) |  | 

### Return type

[**ActivityList**](ActivityList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AddProjectKey

> ProjectKey V1AddProjectKey(ctx, projectId).ProjectKey(projectKey).Execute()

Add Project Authorized Key



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    projectId := "projectId_example" // string | Project ID - uuid
    projectKey := *openapiclient.NewProjectKey("Kind_example", "Key_example") // ProjectKey | Key to add

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1AddProjectKey(context.Background(), projectId).ProjectKey(projectKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1AddProjectKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AddProjectKey`: ProjectKey
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1AddProjectKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AddProjectKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectKey** | [**ProjectKey**](ProjectKey.md) | Key to add | 

### Return type

[**ProjectKey**](ProjectKey.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AddTeamRoleToProject

> V1AddTeamRoleToProject(ctx, projectId, teamId, roleId).Execute()

Add team role to project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    projectId := "projectId_example" // string | Project ID - uuid
    teamId := "teamId_example" // string | Team ID - uuid
    roleId := "roleId_example" // string | Role ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1AddTeamRoleToProject(context.Background(), projectId, teamId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1AddTeamRoleToProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID - uuid | 
**teamId** | **string** | Team ID - uuid | 
**roleId** | **string** | Role ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AddTeamRoleToProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AddUserRoleToProject

> V1AddUserRoleToProject(ctx, projectId, userId, roleId).Execute()

Add user role to project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    projectId := "projectId_example" // string | Project ID - uuid
    userId := "userId_example" // string | User ID - uuid
    roleId := "roleId_example" // string | Role ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1AddUserRoleToProject(context.Background(), projectId, userId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1AddUserRoleToProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID - uuid | 
**userId** | **string** | User ID - uuid | 
**roleId** | **string** | Role ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AddUserRoleToProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AddUserToTeam

> V1AddUserToTeam(ctx, teamId, userId).Execute()

Add user to team



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    teamId := "teamId_example" // string | Team ID - uuid
    userId := "userId_example" // string | User ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1AddUserToTeam(context.Background(), teamId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1AddUserToTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | Team ID - uuid | 
**userId** | **string** | User ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AddUserToTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AgentAppReady

> AgentAppReadyResponse V1AgentAppReady(ctx, instanceId).Execute()

Check if App subsystem is ready



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1AgentAppReady(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1AgentAppReady``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AgentAppReady`: AgentAppReadyResponse
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1AgentAppReady`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AgentAppReadyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AgentAppReadyResponse**](AgentAppReadyResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AgentDeleteFile

> V1AgentDeleteFile(ctx, instanceId, filePath).Execute()

Delete a File on VM



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    filePath := "filePath_example" // string | File Path on VM

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1AgentDeleteFile(context.Background(), instanceId, filePath).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1AgentDeleteFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 
**filePath** | **string** | File Path on VM | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AgentDeleteFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AgentGetFile

> *os.File V1AgentGetFile(ctx, instanceId, filePath).Execute()

Download a File from VM



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    filePath := "filePath_example" // string | File Path on VM

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1AgentGetFile(context.Background(), instanceId, filePath).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1AgentGetFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AgentGetFile`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1AgentGetFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 
**filePath** | **string** | File Path on VM | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AgentGetFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AgentGetTempFilename

> string V1AgentGetTempFilename(ctx, instanceId).Execute()

Get the path for a new temp file



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1AgentGetTempFilename(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1AgentGetTempFilename``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AgentGetTempFilename`: string
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1AgentGetTempFilename`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AgentGetTempFilenameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AgentInstallApp

> V1AgentInstallApp(ctx, instanceId).AgentInstallBody(agentInstallBody).Execute()

Install App at path



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    agentInstallBody := *openapiclient.NewAgentInstallBody() // AgentInstallBody | App parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1AgentInstallApp(context.Background(), instanceId).AgentInstallBody(agentInstallBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1AgentInstallApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AgentInstallAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentInstallBody** | [**AgentInstallBody**](AgentInstallBody.md) | App parameters | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AgentInstallProfile

> V1AgentInstallProfile(ctx, instanceId).Body(body).Execute()

Install a Profile to VM



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    body := os.NewFile(1234, "some_file") // *os.File | Profile to Install

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1AgentInstallProfile(context.Background(), instanceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1AgentInstallProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AgentInstallProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** | Profile to Install | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AgentKillApp

> map[string]interface{} V1AgentKillApp(ctx, instanceId, bundleId).Execute()

Kill an App



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    bundleId := "bundleId_example" // string | App Bundle ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1AgentKillApp(context.Background(), instanceId, bundleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1AgentKillApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AgentKillApp`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1AgentKillApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 
**bundleId** | **string** | App Bundle ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AgentKillAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AgentListAppIcons

> []AgentIcons V1AgentListAppIcons(ctx, instanceId).BundleID(bundleID).Execute()

List App Icons



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    bundleID := []string{"Inner_example"} // []string | App Bundle ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1AgentListAppIcons(context.Background(), instanceId).BundleID(bundleID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1AgentListAppIcons``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AgentListAppIcons`: []AgentIcons
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1AgentListAppIcons`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AgentListAppIconsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bundleID** | **[]string** | App Bundle ID | 

### Return type

[**[]AgentIcons**](AgentIcons.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AgentListApps

> AgentAppsList V1AgentListApps(ctx, instanceId).Execute()

List Apps



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1AgentListApps(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1AgentListApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AgentListApps`: AgentAppsList
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1AgentListApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AgentListAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AgentAppsList**](AgentAppsList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AgentListAppsStatus

> AgentAppsList V1AgentListAppsStatus(ctx, instanceId).Execute()

List Apps Status



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1AgentListAppsStatus(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1AgentListAppsStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AgentListAppsStatus`: AgentAppsList
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1AgentListAppsStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AgentListAppsStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AgentAppsList**](AgentAppsList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AgentListProfiles

> AgentProfilesReturn V1AgentListProfiles(ctx, instanceId).Execute()

List Profiles



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1AgentListProfiles(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1AgentListProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AgentListProfiles`: AgentProfilesReturn
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1AgentListProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AgentListProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AgentProfilesReturn**](AgentProfilesReturn.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AgentRunApp

> map[string]interface{} V1AgentRunApp(ctx, instanceId, bundleId).Execute()

Run an App



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    bundleId := "bundleId_example" // string | App Bundle ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1AgentRunApp(context.Background(), instanceId, bundleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1AgentRunApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AgentRunApp`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1AgentRunApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 
**bundleId** | **string** | App Bundle ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AgentRunAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AgentSetFileAttributes

> V1AgentSetFileAttributes(ctx, instanceId, filePath).FileChanges(fileChanges).Execute()

Change Attrs of a File on VM



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    filePath := "filePath_example" // string | File Path on VM
    fileChanges := *openapiclient.NewFileChanges() // FileChanges | New attrs

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1AgentSetFileAttributes(context.Background(), instanceId, filePath).FileChanges(fileChanges).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1AgentSetFileAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 
**filePath** | **string** | File Path on VM | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AgentSetFileAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fileChanges** | [**FileChanges**](FileChanges.md) | New attrs | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AgentSystemGetAdbAuth

> AgentSystemAdbAuth V1AgentSystemGetAdbAuth(ctx, instanceId).Execute()

Get ADB Auth Setting (AOSP only)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1AgentSystemGetAdbAuth(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1AgentSystemGetAdbAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AgentSystemGetAdbAuth`: AgentSystemAdbAuth
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1AgentSystemGetAdbAuth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AgentSystemGetAdbAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AgentSystemAdbAuth**](AgentSystemAdbAuth.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AgentSystemGetNetwork

> AgentValueReturn V1AgentSystemGetNetwork(ctx, instanceId).Execute()

Get IP of eth0 (AOSP only)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1AgentSystemGetNetwork(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1AgentSystemGetNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AgentSystemGetNetwork`: AgentValueReturn
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1AgentSystemGetNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AgentSystemGetNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AgentValueReturn**](AgentValueReturn.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AgentSystemGetProp

> AgentValueReturn V1AgentSystemGetProp(ctx, instanceId).AgentSystemGetPropBody(agentSystemGetPropBody).Execute()

Get System Property (AOSP only)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    agentSystemGetPropBody := *openapiclient.NewAgentSystemGetPropBody("Property_example") // AgentSystemGetPropBody | Parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1AgentSystemGetProp(context.Background(), instanceId).AgentSystemGetPropBody(agentSystemGetPropBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1AgentSystemGetProp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AgentSystemGetProp`: AgentValueReturn
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1AgentSystemGetProp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AgentSystemGetPropRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentSystemGetPropBody** | [**AgentSystemGetPropBody**](AgentSystemGetPropBody.md) | Parameters | 

### Return type

[**AgentValueReturn**](AgentValueReturn.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AgentSystemInstallOpenGApps

> V1AgentSystemInstallOpenGApps(ctx, instanceId).Body(body).Execute()

Install OpenGApps (AOSP only)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    body := map[string]interface{}{ ... } // map[string]interface{} | Installation parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1AgentSystemInstallOpenGApps(context.Background(), instanceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1AgentSystemInstallOpenGApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AgentSystemInstallOpenGAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Installation parameters | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AgentSystemLock

> V1AgentSystemLock(ctx, instanceId).Execute()

Lock Device (iOS Only)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1AgentSystemLock(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1AgentSystemLock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AgentSystemLockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AgentSystemSetAdbAuth

> V1AgentSystemSetAdbAuth(ctx, instanceId).AgentSystemAdbAuth(agentSystemAdbAuth).Execute()

Set ADB Auth Setting (AOSP only)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    agentSystemAdbAuth := *openapiclient.NewAgentSystemAdbAuth() // AgentSystemAdbAuth | Desired ADB Auth Setting

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1AgentSystemSetAdbAuth(context.Background(), instanceId).AgentSystemAdbAuth(agentSystemAdbAuth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1AgentSystemSetAdbAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AgentSystemSetAdbAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentSystemAdbAuth** | [**AgentSystemAdbAuth**](AgentSystemAdbAuth.md) | Desired ADB Auth Setting | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AgentSystemSetHostname

> V1AgentSystemSetHostname(ctx, instanceId).AgentSystemSetHostnameBody(agentSystemSetHostnameBody).Execute()

Set Hostname of instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    agentSystemSetHostnameBody := *openapiclient.NewAgentSystemSetHostnameBody("Hostname_example") // AgentSystemSetHostnameBody | New hostname

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1AgentSystemSetHostname(context.Background(), instanceId).AgentSystemSetHostnameBody(agentSystemSetHostnameBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1AgentSystemSetHostname``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AgentSystemSetHostnameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentSystemSetHostnameBody** | [**AgentSystemSetHostnameBody**](AgentSystemSetHostnameBody.md) | New hostname | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AgentSystemShutdown

> V1AgentSystemShutdown(ctx, instanceId).Execute()

Instruct VM to halt



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1AgentSystemShutdown(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1AgentSystemShutdown``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AgentSystemShutdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AgentSystemUnlock

> V1AgentSystemUnlock(ctx, instanceId).Execute()

Unlock Device (iOS Only)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1AgentSystemUnlock(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1AgentSystemUnlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AgentSystemUnlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AgentUninstallApp

> map[string]interface{} V1AgentUninstallApp(ctx, instanceId, bundleId).Execute()

Uninstall an App



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    bundleId := "bundleId_example" // string | App Bundle ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1AgentUninstallApp(context.Background(), instanceId, bundleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1AgentUninstallApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AgentUninstallApp`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1AgentUninstallApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 
**bundleId** | **string** | App Bundle ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AgentUninstallAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AgentUninstallProfile

> V1AgentUninstallProfile(ctx, instanceId, profileId).Execute()

Uninstall a Profile from VM



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    profileId := "profileId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1AgentUninstallProfile(context.Background(), instanceId, profileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1AgentUninstallProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 
**profileId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AgentUninstallProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AgentUploadFile

> V1AgentUploadFile(ctx, instanceId, filePath).Body(body).Execute()

Upload a file to VM



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    filePath := "filePath_example" // string | File Path on VM to write to
    body := os.NewFile(1234, "some_file") // *os.File | Uploaded File Contents

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1AgentUploadFile(context.Background(), instanceId, filePath).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1AgentUploadFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 
**filePath** | **string** | File Path on VM to write to | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AgentUploadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | ***os.File** | Uploaded File Contents | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AuthLogin

> Token V1AuthLogin(ctx).Body(body).Execute()

Log In

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    body := map[string]interface{}{ ... } // map[string]interface{} | Authorization data ( Credentials|ApiToken )

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1AuthLogin(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1AuthLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AuthLogin`: Token
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1AuthLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AuthLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Authorization data ( Credentials|ApiToken ) | 

### Return type

[**Token**](Token.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1BtracePreauthorize

> map[string]interface{} V1BtracePreauthorize(ctx, instanceId).Execute()

Pre-authorize an btrace download

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1BtracePreauthorize(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1BtracePreauthorize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1BtracePreauthorize`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1BtracePreauthorize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1BtracePreauthorizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CheckSubdomainExistence

> CheckSubdomainResponse V1CheckSubdomainExistence(ctx).V1CheckSubdomainExistenceParameters(v1CheckSubdomainExistenceParameters).Execute()

Check the existence of a subdomain

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    v1CheckSubdomainExistenceParameters := *openapiclient.NewV1CheckSubdomainExistenceParameters("Domain_example") // V1CheckSubdomainExistenceParameters | application/json

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1CheckSubdomainExistence(context.Background()).V1CheckSubdomainExistenceParameters(v1CheckSubdomainExistenceParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1CheckSubdomainExistence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CheckSubdomainExistence`: CheckSubdomainResponse
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1CheckSubdomainExistence`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CheckSubdomainExistenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1CheckSubdomainExistenceParameters** | [**V1CheckSubdomainExistenceParameters**](V1CheckSubdomainExistenceParameters.md) | application/json | 

### Return type

[**CheckSubdomainResponse**](CheckSubdomainResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ClearCoreTrace

> V1ClearCoreTrace(ctx, instanceId).Execute()

Clear CoreTrace logs

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1ClearCoreTrace(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1ClearCoreTrace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ClearCoreTraceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ClearHyperTrace

> V1ClearHyperTrace(ctx, instanceId).Execute()

Clear HyperTrace logs

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1ClearHyperTrace(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1ClearHyperTrace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ClearHyperTraceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ClearHyperTraceHooks

> V1ClearHyperTraceHooks(ctx, instanceId).Execute()

Clear Hooks on an instance

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1ClearHyperTraceHooks(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1ClearHyperTraceHooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ClearHyperTraceHooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ClearInstancePanics

> V1ClearInstancePanics(ctx, instanceId).Execute()

Clear Panics

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1ClearInstancePanics(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1ClearInstancePanics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ClearInstancePanicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CreateDomainAuthProvider

> DomainAuthProviderResponse V1CreateDomainAuthProvider(ctx, domainId).DomainAuthProviderRequest(domainAuthProviderRequest).Execute()

Create a new auth provider for a domain

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    domainId := "domainId_example" // string | Domain ID - uuid
    domainAuthProviderRequest := *openapiclient.NewDomainAuthProviderRequest("ProviderType_example", false) // DomainAuthProviderRequest | Request Data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1CreateDomainAuthProvider(context.Background(), domainId).DomainAuthProviderRequest(domainAuthProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1CreateDomainAuthProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CreateDomainAuthProvider`: DomainAuthProviderResponse
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1CreateDomainAuthProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | Domain ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CreateDomainAuthProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **domainAuthProviderRequest** | [**DomainAuthProviderRequest**](DomainAuthProviderRequest.md) | Request Data | 

### Return type

[**DomainAuthProviderResponse**](DomainAuthProviderResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CreateHook

> Hook V1CreateHook(ctx, instanceId).V1CreateHookParameters(v1CreateHookParameters).Execute()

Create hypervisor hook for Instance

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    v1CreateHookParameters := *openapiclient.NewV1CreateHookParameters("Label_example", "Address_example", "Patch_example", "PatchType_example") // V1CreateHookParameters | application/json

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1CreateHook(context.Background(), instanceId).V1CreateHookParameters(v1CreateHookParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1CreateHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CreateHook`: Hook
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1CreateHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CreateHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1CreateHookParameters** | [**V1CreateHookParameters**](V1CreateHookParameters.md) | application/json | 

### Return type

[**Hook**](Hook.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CreateImage

> Image V1CreateImage(ctx).Type_(type_).Encoding(encoding).Encapsulated(encapsulated).Name(name).Project(project).Instance(instance).File(file).Execute()

Create a new Image

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    type_ := "type__example" // string | Image type
    encoding := "encoding_example" // string | How the file is stored
    encapsulated := true // bool | set to false if the uploaded file is not encapsulated inside an outer zipfile (optional)
    name := "name_example" // string | Image name (optional)
    project := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project ID (optional)
    instance := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Instance ID (optional)
    file := os.NewFile(1234, "some_file") // *os.File | Optionally the actual file (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1CreateImage(context.Background()).Type_(type_).Encoding(encoding).Encapsulated(encapsulated).Name(name).Project(project).Instance(instance).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1CreateImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CreateImage`: Image
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1CreateImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CreateImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Image type | 
 **encoding** | **string** | How the file is stored | 
 **encapsulated** | **bool** | set to false if the uploaded file is not encapsulated inside an outer zipfile | 
 **name** | **string** | Image name | 
 **project** | **string** | Project ID | 
 **instance** | **string** | Instance ID | 
 **file** | ***os.File** | Optionally the actual file | 

### Return type

[**Image**](Image.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CreateInstance

> InstanceReturn V1CreateInstance(ctx).InstanceCreateOptions(instanceCreateOptions).Execute()

Create Instance

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceCreateOptions := *openapiclient.NewInstanceCreateOptions("Flavor_example", "Project_example", "Os_example") // InstanceCreateOptions | The vm definition to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1CreateInstance(context.Background()).InstanceCreateOptions(instanceCreateOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1CreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CreateInstance`: InstanceReturn
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1CreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instanceCreateOptions** | [**InstanceCreateOptions**](InstanceCreateOptions.md) | The vm definition to create | 

### Return type

[**InstanceReturn**](InstanceReturn.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CreateNetworkConnection

> NetworkConnection V1CreateNetworkConnection(ctx).CreateNetworkConnectionOptions(createNetworkConnectionOptions).Execute()

Create a new Network Connection



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    createNetworkConnectionOptions := *openapiclient.NewCreateNetworkConnectionOptions("Identifier_example", "Name_example", "Provider_example") // CreateNetworkConnectionOptions | Network Connection Options

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1CreateNetworkConnection(context.Background()).CreateNetworkConnectionOptions(createNetworkConnectionOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1CreateNetworkConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CreateNetworkConnection`: NetworkConnection
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1CreateNetworkConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CreateNetworkConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetworkConnectionOptions** | [**CreateNetworkConnectionOptions**](CreateNetworkConnectionOptions.md) | Network Connection Options | 

### Return type

[**NetworkConnection**](NetworkConnection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CreateProject

> Project V1CreateProject(ctx).Project(project).Execute()

Create a Project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    project := *openapiclient.NewProject("Id_example") // Project | Project

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1CreateProject(context.Background()).Project(project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1CreateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CreateProject`: Project
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1CreateProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | [**Project**](Project.md) | Project | 

### Return type

[**Project**](Project.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CreateSnapshot

> Snapshot V1CreateSnapshot(ctx, instanceId).SnapshotCreationOptions(snapshotCreationOptions).Execute()

Create Instance Snapshot



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    snapshotCreationOptions := *openapiclient.NewSnapshotCreationOptions("Name_example") // SnapshotCreationOptions | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1CreateSnapshot(context.Background(), instanceId).SnapshotCreationOptions(snapshotCreationOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1CreateSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CreateSnapshot`: Snapshot
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1CreateSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CreateSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **snapshotCreationOptions** | [**SnapshotCreationOptions**](SnapshotCreationOptions.md) |  | 

### Return type

[**Snapshot**](Snapshot.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CreateUser

> map[string]interface{} V1CreateUser(ctx).Body(body).Execute()

Create User

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    body := map[string]interface{}{ ... } // map[string]interface{} | User data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1CreateUser(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CreateUser`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | User data | 

### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeleteDomainAuthProvider

> map[string]interface{} V1DeleteDomainAuthProvider(ctx, domainId, providerId).Execute()

Delete an auth provider from a domain

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    domainId := "domainId_example" // string | Domain ID - uuid
    providerId := "providerId_example" // string | Provider ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1DeleteDomainAuthProvider(context.Background(), domainId, providerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1DeleteDomainAuthProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DeleteDomainAuthProvider`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1DeleteDomainAuthProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | Domain ID - uuid | 
**providerId** | **string** | Provider ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeleteDomainAuthProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeleteExtension

> V1DeleteExtension(ctx, extensionId).Execute()

Delete an existing extension

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    extensionId := "extensionId_example" // string | Extension ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1DeleteExtension(context.Background(), extensionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1DeleteExtension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **string** | Extension ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeleteExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeleteHook

> V1DeleteHook(ctx, hookId).Execute()

Delete an existing hypervisor hook

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    hookId := "hookId_example" // string | Hook ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1DeleteHook(context.Background(), hookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1DeleteHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hookId** | **string** | Hook ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeleteHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeleteImage

> V1DeleteImage(ctx, imageId).Execute()

Delete Image

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    imageId := "imageId_example" // string | Image ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1DeleteImage(context.Background(), imageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1DeleteImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | Image ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeleteImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeleteInstance

> V1DeleteInstance(ctx, instanceId).Execute()

Remove Instance

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1DeleteInstance(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1DeleteInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeleteInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeleteInstanceSnapshot

> V1DeleteInstanceSnapshot(ctx, instanceId, snapshotId).Execute()

Delete an Instance Snapshot



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    snapshotId := "snapshotId_example" // string | Snapshot ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1DeleteInstanceSnapshot(context.Background(), instanceId, snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1DeleteInstanceSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 
**snapshotId** | **string** | Snapshot ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeleteInstanceSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeleteNetworkConnection

> V1DeleteNetworkConnection(ctx, id).Force(force).Execute()

Delete an existing Network Connection



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    id := "id_example" // string | Network Connection Identifier - uuid
    force := true // bool | Force deletion (true only or not present) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1DeleteNetworkConnection(context.Background(), id).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1DeleteNetworkConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Network Connection Identifier - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeleteNetworkConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | Force deletion (true only or not present) | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeleteProject

> V1DeleteProject(ctx, projectId).Execute()

Delete a Project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    projectId := "projectId_example" // string | Project ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1DeleteProject(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1DeleteProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeleteProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeleteSnapshot

> V1DeleteSnapshot(ctx, snapshotId).Execute()

Delete a Snapshot



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    snapshotId := "snapshotId_example" // string | Snapshot ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1DeleteSnapshot(context.Background(), snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1DeleteSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | Snapshot ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeleteSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeleteSnapshotPermissions

> Snapshot V1DeleteSnapshotPermissions(ctx, snapshotId).SnapshotPermissionsRequestPayload(snapshotPermissionsRequestPayload).Execute()

Delete member(s)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    snapshotId := "snapshotId_example" // string | Snapshot ID - uuid
    snapshotPermissionsRequestPayload := *openapiclient.NewSnapshotPermissionsRequestPayload([]string{"Members_example"}) // SnapshotPermissionsRequestPayload | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1DeleteSnapshotPermissions(context.Background(), snapshotId).SnapshotPermissionsRequestPayload(snapshotPermissionsRequestPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1DeleteSnapshotPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DeleteSnapshotPermissions`: Snapshot
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1DeleteSnapshotPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | Snapshot ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeleteSnapshotPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **snapshotPermissionsRequestPayload** | [**SnapshotPermissionsRequestPayload**](SnapshotPermissionsRequestPayload.md) |  | 

### Return type

[**Snapshot**](Snapshot.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeleteUser

> map[string]interface{} V1DeleteUser(ctx, userId).Execute()

Delete User

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    userId := "userId_example" // string | userId - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1DeleteUser(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1DeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DeleteUser`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1DeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | userId - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DisableExposePort

> V1DisableExposePort(ctx, instanceId).Execute()

Disable an exposed port on an instance for device access.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1DisableExposePort(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1DisableExposePort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DisableExposePortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DownloadActivity

> ActivityList V1DownloadActivity(ctx, taskId).Execute()

Download activity

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    taskId := "taskId_example" // string | Export ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1DownloadActivity(context.Background(), taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1DownloadActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DownloadActivity`: ActivityList
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1DownloadActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | Export ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DownloadActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActivityList**](ActivityList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnableExposePort

> V1EnableExposePort(ctx, instanceId).Execute()

Enable an exposed port on an instance for device access.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1EnableExposePort(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1EnableExposePort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnableExposePortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ExecuteHyperTraceHooks

> V1ExecuteHyperTraceHooks(ctx, instanceId).Execute()

Execute Hooks on an instance

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1ExecuteHyperTraceHooks(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1ExecuteHyperTraceHooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ExecuteHyperTraceHooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetActivityExportStatus

> ActivityList V1GetActivityExportStatus(ctx, taskId).Execute()

Get export task status

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    taskId := "taskId_example" // string | Export ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetActivityExportStatus(context.Background(), taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetActivityExportStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetActivityExportStatus`: ActivityList
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetActivityExportStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | Export ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetActivityExportStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActivityList**](ActivityList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetActivityExportTasks

> ActivityList V1GetActivityExportTasks(ctx).Execute()

Get all export tasks for user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetActivityExportTasks(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetActivityExportTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetActivityExportTasks`: ActivityList
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetActivityExportTasks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetActivityExportTasksRequest struct via the builder pattern


### Return type

[**ActivityList**](ActivityList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetConfig

> ConfigResponse V1GetConfig(ctx).Execute()

Get all configs

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetConfig`: ConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetConfigRequest struct via the builder pattern


### Return type

[**ConfigResponse**](ConfigResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetDomainAuthProviders

> []DomainAuthProviderResponse V1GetDomainAuthProviders(ctx, domainId).Execute()

Return all configured auth providers for a domain (including globally configured providers)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    domainId := "domainId_example" // string | Domain ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetDomainAuthProviders(context.Background(), domainId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetDomainAuthProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetDomainAuthProviders`: []DomainAuthProviderResponse
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetDomainAuthProviders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | Domain ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetDomainAuthProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DomainAuthProviderResponse**](DomainAuthProviderResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetExtensionById

> Extension V1GetExtensionById(ctx, extensionId).Execute()

Get extension by id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    extensionId := "extensionId_example" // string | Extension Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetExtensionById(context.Background(), extensionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetExtensionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetExtensionById`: Extension
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetExtensionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **string** | Extension Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetExtensionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Extension**](Extension.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetExtensions

> []Extension V1GetExtensions(ctx).Limit(limit).Offset(offset).IfNoneMatch(ifNoneMatch).Execute()

Get all extensions

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    limit := float32(8.14) // float32 | limit for pagination results, defaults to 20 (optional)
    offset := float32(8.14) // float32 | offset for pagination results, defaults to 0 (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | sha256sum of the last response with the same parameters (optional) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetExtensions(context.Background()).Limit(limit).Offset(offset).IfNoneMatch(ifNoneMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetExtensions`: []Extension
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetExtensions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1GetExtensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** | limit for pagination results, defaults to 20 | 
 **offset** | **float32** | offset for pagination results, defaults to 0 | 
 **ifNoneMatch** | **string** | sha256sum of the last response with the same parameters (optional) | 

### Return type

[**[]Extension**](Extension.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetHookById

> Hook V1GetHookById(ctx, hookId).Execute()

Get hypervisor hook by id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    hookId := "hookId_example" // string | Hook Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetHookById(context.Background(), hookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetHookById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetHookById`: Hook
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetHookById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hookId** | **string** | Hook Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetHookByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Hook**](Hook.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetHooks

> []Hook V1GetHooks(ctx, instanceId).Limit(limit).Offset(offset).Sort(sort).Execute()

Get all hypervisor hooks for Instance

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    limit := float32(8.14) // float32 | limit for pagination results, defaults to 20 (optional)
    offset := float32(8.14) // float32 | offset for pagination results, defaults to 0 (optional)
    sort := "sort_example" // string | sort ASC or DESC, defaults to DESC (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetHooks(context.Background(), instanceId).Limit(limit).Offset(offset).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetHooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetHooks`: []Hook
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetHooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetHooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **float32** | limit for pagination results, defaults to 20 | 
 **offset** | **float32** | offset for pagination results, defaults to 0 | 
 **sort** | **string** | sort ASC or DESC, defaults to DESC | 

### Return type

[**[]Hook**](Hook.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetImage

> Image V1GetImage(ctx, imageId).Execute()

Get Image Metadata

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    imageId := "imageId_example" // string | Image ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetImage(context.Background(), imageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetImage`: Image
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | Image ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Image**](Image.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetImages

> []Image V1GetImages(ctx).Project(project).Execute()

Get all Images Metadata

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    project := "project_example" // string | Optionally filter by project - uuid (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetImages(context.Background()).Project(project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetImages`: []Image
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1GetImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **string** | Optionally filter by project - uuid | 

### Return type

[**[]Image**](Image.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetInstance

> Instance V1GetInstance(ctx, instanceId).ReturnAttr(returnAttr).Execute()

Get Instance

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID
    returnAttr := []string{"Inner_example"} // []string | Attributes to include in instance return (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetInstance(context.Background(), instanceId).ReturnAttr(returnAttr).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetInstance`: Instance
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnAttr** | **[]string** | Attributes to include in instance return | 

### Return type

[**Instance**](Instance.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetInstanceConsole

> InstanceConsoleEndpoint V1GetInstanceConsole(ctx, instanceId).Execute()

Get console websocket URL

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetInstanceConsole(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetInstanceConsole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetInstanceConsole`: InstanceConsoleEndpoint
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetInstanceConsole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetInstanceConsoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InstanceConsoleEndpoint**](InstanceConsoleEndpoint.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetInstanceConsoleLog

> string V1GetInstanceConsoleLog(ctx, instanceId).Execute()

Get Console Log

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetInstanceConsoleLog(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetInstanceConsoleLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetInstanceConsoleLog`: string
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetInstanceConsoleLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetInstanceConsoleLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetInstanceGpios

> GpiosState V1GetInstanceGpios(ctx, instanceId).Execute()

Get Instance GPIOs

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetInstanceGpios(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetInstanceGpios``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetInstanceGpios`: GpiosState
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetInstanceGpios`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetInstanceGpiosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GpiosState**](GpiosState.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetInstancePanics

> []map[string]interface{} V1GetInstancePanics(ctx, instanceId).Execute()

Get Panics

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetInstancePanics(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetInstancePanics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetInstancePanics`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetInstancePanics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetInstancePanicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetInstancePeripherals

> PeripheralsData V1GetInstancePeripherals(ctx, instanceId).Execute()

Get Instance Peripherals

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetInstancePeripherals(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetInstancePeripherals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetInstancePeripherals`: PeripheralsData
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetInstancePeripherals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetInstancePeripheralsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PeripheralsData**](PeripheralsData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetInstanceScreenshot

> *os.File V1GetInstanceScreenshot(ctx, instanceId, format).Scale(scale).Execute()

Get Instance Screenshot

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    format := "format_example" // string | New peripherals state
    scale := float32(56) // float32 | Screenshot scale 1:N (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetInstanceScreenshot(context.Background(), instanceId, format).Scale(scale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetInstanceScreenshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetInstanceScreenshot`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetInstanceScreenshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 
**format** | **string** | New peripherals state | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetInstanceScreenshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scale** | **float32** | Screenshot scale 1:N | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/png, image/jpeg, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetInstanceSnapshot

> Snapshot V1GetInstanceSnapshot(ctx, instanceId, snapshotId).Execute()

Get Instance Snapshot



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    snapshotId := "snapshotId_example" // string | Snapshot ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetInstanceSnapshot(context.Background(), instanceId, snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetInstanceSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetInstanceSnapshot`: Snapshot
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetInstanceSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 
**snapshotId** | **string** | Snapshot ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetInstanceSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Snapshot**](Snapshot.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetInstanceSnapshots

> []Snapshot V1GetInstanceSnapshots(ctx, instanceId).Execute()

Get Instance Snapshots



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetInstanceSnapshots(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetInstanceSnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetInstanceSnapshots`: []Snapshot
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetInstanceSnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetInstanceSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Snapshot**](Snapshot.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetInstances

> []Instance V1GetInstances(ctx).Name(name).ReturnAttr(returnAttr).Execute()

Get Instances

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    name := "name_example" // string | Optionally filter by instance name (optional)
    returnAttr := []string{"Inner_example"} // []string | Attributes to include in instance return (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetInstances(context.Background()).Name(name).ReturnAttr(returnAttr).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetInstances`: []Instance
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1GetInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Optionally filter by instance name | 
 **returnAttr** | **[]string** | Attributes to include in instance return | 

### Return type

[**[]Instance**](Instance.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetModelSoftware

> []Firmware V1GetModelSoftware(ctx, model).Execute()

Get Software for Model

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    model := "model_example" // string | Model to list available software for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetModelSoftware(context.Background(), model).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetModelSoftware``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetModelSoftware`: []Firmware
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetModelSoftware`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**model** | **string** | Model to list available software for | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetModelSoftwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Firmware**](Firmware.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetModels

> []Model V1GetModels(ctx).Execute()

Get Supported Models

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetModels(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetModels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetModels`: []Model
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetModels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetModelsRequest struct via the builder pattern


### Return type

[**[]Model**](Model.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetNetworkConnection

> NetworkConnection V1GetNetworkConnection(ctx, id).Execute()

Return the configuration and per project statuses for a single network provider.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    id := "id_example" // string | Network Connection Identifier - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetNetworkConnection(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetNetworkConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetNetworkConnection`: NetworkConnection
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetNetworkConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Network Connection Identifier - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetNetworkConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkConnection**](NetworkConnection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetProject

> Project V1GetProject(ctx, projectId).Execute()

Get a Project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    projectId := "projectId_example" // string | Project ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetProject(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetProject`: Project
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Project**](Project.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetProjectInstances

> []Instance V1GetProjectInstances(ctx, projectId).Name(name).ReturnAttr(returnAttr).Execute()

Get Instances in Project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    projectId := "projectId_example" // string | Project ID - uuid
    name := "name_example" // string | Filter by project name (optional)
    returnAttr := []string{"Inner_example"} // []string | Attributes to include in instance return (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetProjectInstances(context.Background(), projectId).Name(name).ReturnAttr(returnAttr).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetProjectInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetProjectInstances`: []Instance
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetProjectInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetProjectInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Filter by project name | 
 **returnAttr** | **[]string** | Attributes to include in instance return | 

### Return type

[**[]Instance**](Instance.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetProjectKeys

> []ProjectKey V1GetProjectKeys(ctx, projectId).Execute()

Get Project Authorized Keys



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    projectId := "projectId_example" // string | Project ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetProjectKeys(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetProjectKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetProjectKeys`: []ProjectKey
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetProjectKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetProjectKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ProjectKey**](ProjectKey.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetProjectNetworkLog

> string V1GetProjectNetworkLog(ctx, projectId).Execute()

Retrieve the network connection log for a project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    projectId := "projectId_example" // string | Project ID (must be a valid UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetProjectNetworkLog(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetProjectNetworkLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetProjectNetworkLog`: string
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetProjectNetworkLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID (must be a valid UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetProjectNetworkLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetProjectNetworkStatus

> NetworkStatusResponse V1GetProjectNetworkStatus(ctx, projectId).Execute()

Retrieve the network connection status for a project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    projectId := "projectId_example" // string | Project ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetProjectNetworkStatus(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetProjectNetworkStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetProjectNetworkStatus`: NetworkStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetProjectNetworkStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetProjectNetworkStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkStatusResponse**](NetworkStatusResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetProjectVpnConfig

> string V1GetProjectVpnConfig(ctx, projectId, format).Execute()

Get Project VPN Configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    projectId := "projectId_example" // string | Project ID - uuid
    format := "format_example" // string | VPN Config format

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetProjectVpnConfig(context.Background(), projectId, format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetProjectVpnConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetProjectVpnConfig`: string
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetProjectVpnConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID - uuid | 
**format** | **string** | VPN Config format | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetProjectVpnConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetProjects

> []Project V1GetProjects(ctx).Name(name).IdsOnly(idsOnly).Execute()

Get Projects



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    name := "name_example" // string | Filter by project name (optional)
    idsOnly := true // bool | Only include id of project in results (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetProjects(context.Background()).Name(name).IdsOnly(idsOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetProjects`: []Project
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1GetProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by project name | 
 **idsOnly** | **bool** | Only include id of project in results | 

### Return type

[**[]Project**](Project.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetResetLinkInfo

> map[string]interface{} V1GetResetLinkInfo(ctx).Token(token).Execute()

Send Password Reset Link Info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    token := "token_example" // string | Reset token

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetResetLinkInfo(context.Background()).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetResetLinkInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetResetLinkInfo`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetResetLinkInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1GetResetLinkInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | Reset token | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetSharedSnapshots

> UserSnapshots V1GetSharedSnapshots(ctx).Execute()

Fetch shared snapshots



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetSharedSnapshots(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetSharedSnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetSharedSnapshots`: UserSnapshots
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetSharedSnapshots`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetSharedSnapshotsRequest struct via the builder pattern


### Return type

[**UserSnapshots**](UserSnapshots.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetSnapshot

> Snapshot V1GetSnapshot(ctx, snapshotId).Execute()

Get Snapshot



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    snapshotId := "snapshotId_example" // string | Snapshot ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1GetSnapshot(context.Background(), snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1GetSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GetSnapshot`: Snapshot
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1GetSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | Snapshot ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Snapshot**](Snapshot.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesInstanceIdMessagePost

> V1InstancesInstanceIdMessagePost(ctx, instanceId).Body(body).Execute()

Inject a message into an iOS VM



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    body := map[string]interface{}{ ... } // map[string]interface{} | Message data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1InstancesInstanceIdMessagePost(context.Background(), instanceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1InstancesInstanceIdMessagePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesInstanceIdMessagePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Message data | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesInstanceIdNetdumpPcapGet

> *os.File V1InstancesInstanceIdNetdumpPcapGet(ctx, instanceId).Execute()

Download a netdump pcap file

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1InstancesInstanceIdNetdumpPcapGet(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1InstancesInstanceIdNetdumpPcapGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesInstanceIdNetdumpPcapGet`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1InstancesInstanceIdNetdumpPcapGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesInstanceIdNetdumpPcapGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.tcpdump.pcap, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesInstanceIdNetworkMonitorPcapGet

> *os.File V1InstancesInstanceIdNetworkMonitorPcapGet(ctx, instanceId).Execute()

Download a Network Monitor pcap file

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1InstancesInstanceIdNetworkMonitorPcapGet(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1InstancesInstanceIdNetworkMonitorPcapGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesInstanceIdNetworkMonitorPcapGet`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1InstancesInstanceIdNetworkMonitorPcapGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesInstanceIdNetworkMonitorPcapGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.tcpdump.pcap, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1Kcrange

> []Kcrange V1Kcrange(ctx, instanceId).Execute()

Get Kernel extension ranges

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1Kcrange(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1Kcrange``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1Kcrange`: []Kcrange
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1Kcrange`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1KcrangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Kcrange**](Kcrange.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ListNetworkConnections

> NetworkConnectionOffsetPaginationResult V1ListNetworkConnections(ctx).Limit(limit).Offset(offset).Execute()

List available network connections



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    limit := float32(8.14) // float32 | The maximum number of items to return. (optional)
    offset := float32(8.14) // float32 | The starting index of the items to return. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1ListNetworkConnections(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1ListNetworkConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ListNetworkConnections`: NetworkConnectionOffsetPaginationResult
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1ListNetworkConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ListNetworkConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** | The maximum number of items to return. | 
 **offset** | **float32** | The starting index of the items to return. | 

### Return type

[**NetworkConnectionOffsetPaginationResult**](NetworkConnectionOffsetPaginationResult.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ListNetworkInterfaces

> []string V1ListNetworkInterfaces(ctx).Execute()

List available physical network interfaces



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1ListNetworkInterfaces(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1ListNetworkInterfaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ListNetworkInterfaces`: []string
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1ListNetworkInterfaces`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1ListNetworkInterfacesRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ListNetworkProviders

> NetworkConnectionProviderOffsetPaginationResult V1ListNetworkProviders(ctx).Execute()

List available network providers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1ListNetworkProviders(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1ListNetworkProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ListNetworkProviders`: NetworkConnectionProviderOffsetPaginationResult
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1ListNetworkProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1ListNetworkProvidersRequest struct via the builder pattern


### Return type

[**NetworkConnectionProviderOffsetPaginationResult**](NetworkConnectionProviderOffsetPaginationResult.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ListThreads

> []KernelTask V1ListThreads(ctx, instanceId).Execute()

Get Running Threads (CoreTrace)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1ListThreads(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1ListThreads``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ListThreads`: []KernelTask
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1ListThreads`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ListThreadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]KernelTask**](KernelTask.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1LoadExtension

> Extension V1LoadExtension(ctx).V1LoadExtensionParameters(v1LoadExtensionParameters).Execute()

Load an extension

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    v1LoadExtensionParameters := *openapiclient.NewV1LoadExtensionParameters("ImageId_example") // V1LoadExtensionParameters | application/json

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1LoadExtension(context.Background()).V1LoadExtensionParameters(v1LoadExtensionParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1LoadExtension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1LoadExtension`: Extension
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1LoadExtension`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1LoadExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1LoadExtensionParameters** | [**V1LoadExtensionParameters**](V1LoadExtensionParameters.md) | application/json | 

### Return type

[**Extension**](Extension.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MediaPlay

> V1MediaPlay(ctx, instanceId).MediaPlayBody(mediaPlayBody).Execute()

Start playing media

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    mediaPlayBody := *openapiclient.NewMediaPlayBody() // MediaPlayBody | Request Body

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1MediaPlay(context.Background(), instanceId).MediaPlayBody(mediaPlayBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1MediaPlay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MediaPlayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mediaPlayBody** | [**MediaPlayBody**](MediaPlayBody.md) | Request Body | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MediaStop

> V1MediaStop(ctx, instanceId).Execute()

Stop playing media

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1MediaStop(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1MediaStop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MediaStopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ParseExtensionJson

> map[string]interface{} V1ParseExtensionJson(ctx).Extension(extension).Execute()

Validates extension.json

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    extension := *openapiclient.NewExtension() // Extension | extension.json contents

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1ParseExtensionJson(context.Background()).Extension(extension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1ParseExtensionJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ParseExtensionJson`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1ParseExtensionJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ParseExtensionJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **extension** | [**Extension**](Extension.md) | extension.json contents | 

### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PartialUpdateNetworkConnection

> V1PartialUpdateNetworkConnection(ctx, id).UpdateNetworkConnectionOptions(updateNetworkConnectionOptions).Force(force).Execute()

Update Network Connection (partial)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    id := "id_example" // string | Network Connection Identifier - uuid
    updateNetworkConnectionOptions := *openapiclient.NewUpdateNetworkConnectionOptions("Name_example") // UpdateNetworkConnectionOptions | Network Connection Options
    force := true // bool | Force deletion (true only or not present) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1PartialUpdateNetworkConnection(context.Background(), id).UpdateNetworkConnectionOptions(updateNetworkConnectionOptions).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1PartialUpdateNetworkConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Network Connection Identifier - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PartialUpdateNetworkConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkConnectionOptions** | [**UpdateNetworkConnectionOptions**](UpdateNetworkConnectionOptions.md) | Network Connection Options | 
 **force** | **bool** | Force deletion (true only or not present) | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PatchInstance

> Instance V1PatchInstance(ctx, instanceId).PatchInstanceOptions(patchInstanceOptions).Execute()

Update Instance

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID
    patchInstanceOptions := *openapiclient.NewPatchInstanceOptions() // PatchInstanceOptions | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1PatchInstance(context.Background(), instanceId).PatchInstanceOptions(patchInstanceOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1PatchInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1PatchInstance`: Instance
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1PatchInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PatchInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchInstanceOptions** | [**PatchInstanceOptions**](PatchInstanceOptions.md) |  | 

### Return type

[**Instance**](Instance.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PatchInstanceReadOnly

> Instance V1PatchInstanceReadOnly(ctx, instanceId).PatchInstanceReadOnly(patchInstanceReadOnly).Execute()

Update Instance Read Only

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID
    patchInstanceReadOnly := *openapiclient.NewPatchInstanceReadOnly() // PatchInstanceReadOnly | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1PatchInstanceReadOnly(context.Background(), instanceId).PatchInstanceReadOnly(patchInstanceReadOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1PatchInstanceReadOnly``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1PatchInstanceReadOnly`: Instance
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1PatchInstanceReadOnly`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PatchInstanceReadOnlyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchInstanceReadOnly** | [**PatchInstanceReadOnly**](PatchInstanceReadOnly.md) |  | 

### Return type

[**Instance**](Instance.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PauseInstance

> V1PauseInstance(ctx, instanceId).Execute()

Pause an Instance

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1PauseInstance(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1PauseInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PauseInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PostInstanceInput

> InputResponse V1PostInstanceInput(ctx, instanceId).InstanceInput(instanceInput).Execute()

Provide Instance Input



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    instanceInput := []openapiclient.InstanceInput{openapiclient.InstanceInput{TextInput: openapiclient.NewTextInput("Required_example")}} // []InstanceInput | The input to send to the VM

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1PostInstanceInput(context.Background(), instanceId).InstanceInput(instanceInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1PostInstanceInput``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1PostInstanceInput`: InputResponse
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1PostInstanceInput`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PostInstanceInputRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instanceInput** | [**[]InstanceInput**](InstanceInput.md) | The input to send to the VM | 

### Return type

[**InputResponse**](InputResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1Ready

> V1Ready(ctx).Execute()

API Status



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1Ready(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1Ready``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1ReadyRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1RebootInstance

> V1RebootInstance(ctx, instanceId).Execute()

Reboot an Instance

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1RebootInstance(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1RebootInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1RebootInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1RemoveProjectKey

> V1RemoveProjectKey(ctx, projectId, keyId).Execute()

Delete Project Authorized Key



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    projectId := "projectId_example" // string | Project ID - uuid
    keyId := "keyId_example" // string | Key ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1RemoveProjectKey(context.Background(), projectId, keyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1RemoveProjectKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID - uuid | 
**keyId** | **string** | Key ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1RemoveProjectKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1RemoveTeamRoleFromProject

> V1RemoveTeamRoleFromProject(ctx, projectId, teamId, roleId).Execute()

Remove team role from project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    projectId := "projectId_example" // string | Project ID - uuid
    teamId := "teamId_example" // string | Team ID - uuid
    roleId := "roleId_example" // string | Role ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1RemoveTeamRoleFromProject(context.Background(), projectId, teamId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1RemoveTeamRoleFromProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID - uuid | 
**teamId** | **string** | Team ID - uuid | 
**roleId** | **string** | Role ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1RemoveTeamRoleFromProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1RemoveUserFromTeam

> V1RemoveUserFromTeam(ctx, teamId, userId).Execute()

Remove user from team



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    teamId := "teamId_example" // string | Team ID - uuid
    userId := "userId_example" // string | User ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1RemoveUserFromTeam(context.Background(), teamId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1RemoveUserFromTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | Team ID - uuid | 
**userId** | **string** | User ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1RemoveUserFromTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1RemoveUserRoleFromProject

> V1RemoveUserRoleFromProject(ctx, projectId, userId, roleId).Execute()

Remove user role from project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    projectId := "projectId_example" // string | Project ID - uuid
    userId := "userId_example" // string | User ID - uuid
    roleId := "roleId_example" // string | Role ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1RemoveUserRoleFromProject(context.Background(), projectId, userId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1RemoveUserRoleFromProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID - uuid | 
**userId** | **string** | User ID - uuid | 
**roleId** | **string** | Role ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1RemoveUserRoleFromProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1RenameInstanceSnapshot

> Snapshot V1RenameInstanceSnapshot(ctx, instanceId, snapshotId).SnapshotCreationOptions(snapshotCreationOptions).Execute()

Rename an Instance Snapshot



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    snapshotId := "snapshotId_example" // string | Snapshot ID - uuid
    snapshotCreationOptions := *openapiclient.NewSnapshotCreationOptions("Name_example") // SnapshotCreationOptions | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1RenameInstanceSnapshot(context.Background(), instanceId, snapshotId).SnapshotCreationOptions(snapshotCreationOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1RenameInstanceSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1RenameInstanceSnapshot`: Snapshot
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1RenameInstanceSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 
**snapshotId** | **string** | Snapshot ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1RenameInstanceSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **snapshotCreationOptions** | [**SnapshotCreationOptions**](SnapshotCreationOptions.md) |  | 

### Return type

[**Snapshot**](Snapshot.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ResetUserPassword

> V1ResetUserPassword(ctx).PasswordResetBody(passwordResetBody).Execute()

Reset User Password

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    passwordResetBody := *openapiclient.NewPasswordResetBody("Token_example", "TotpToken_example", "NewPassword_example") // PasswordResetBody | application/json

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1ResetUserPassword(context.Background()).PasswordResetBody(passwordResetBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1ResetUserPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ResetUserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordResetBody** | [**PasswordResetBody**](PasswordResetBody.md) | application/json | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1RestoreBackup

> V1RestoreBackup(ctx, instanceId).Body(body).Execute()

Restore backup

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    body := map[string]interface{}{ ... } // map[string]interface{} | Restore backup data (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1RestoreBackup(context.Background(), instanceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1RestoreBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1RestoreBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Restore backup data | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1RestoreInstanceSnapshot

> V1RestoreInstanceSnapshot(ctx, instanceId, snapshotId).Execute()

Restore an Instance Snapshot



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    snapshotId := "snapshotId_example" // string | Snapshot ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1RestoreInstanceSnapshot(context.Background(), instanceId, snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1RestoreInstanceSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 
**snapshotId** | **string** | Snapshot ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1RestoreInstanceSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1Roles

> []Role V1Roles(ctx).Execute()

Get all roles



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1Roles(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1Roles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1Roles`: []Role
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1Roles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1RolesRequest struct via the builder pattern


### Return type

[**[]Role**](Role.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1RotateInstance

> V1RotateInstance(ctx, instanceId).RotateBody(rotateBody).Execute()

Rotate device to specified orientation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | 
    rotateBody := *openapiclient.NewRotateBody(float32(123)) // RotateBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1RotateInstance(context.Background(), instanceId).RotateBody(rotateBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1RotateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1RotateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rotateBody** | [**RotateBody**](RotateBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SendUserResetLink

> V1SendUserResetLink(ctx).ResetLinkBody(resetLinkBody).Execute()

Send Password Reset Link

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    resetLinkBody := *openapiclient.NewResetLinkBody("Email_example") // ResetLinkBody | application/json

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1SendUserResetLink(context.Background()).ResetLinkBody(resetLinkBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1SendUserResetLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SendUserResetLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetLinkBody** | [**ResetLinkBody**](ResetLinkBody.md) | application/json | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SetInstanceGpios

> GpiosState V1SetInstanceGpios(ctx, instanceId).GpiosState(gpiosState).Execute()

Set Instance GPIOs

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    gpiosState := *openapiclient.NewGpiosState() // GpiosState | New GPIO state

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1SetInstanceGpios(context.Background(), instanceId).GpiosState(gpiosState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1SetInstanceGpios``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1SetInstanceGpios`: GpiosState
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1SetInstanceGpios`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SetInstanceGpiosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gpiosState** | [**GpiosState**](GpiosState.md) | New GPIO state | 

### Return type

[**GpiosState**](GpiosState.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SetInstancePeripherals

> PeripheralsData V1SetInstancePeripherals(ctx, instanceId).PeripheralsData(peripheralsData).Execute()

Set Instance Peripherals

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    peripheralsData := *openapiclient.NewPeripheralsData() // PeripheralsData | New peripherals state

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1SetInstancePeripherals(context.Background(), instanceId).PeripheralsData(peripheralsData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1SetInstancePeripherals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1SetInstancePeripherals`: PeripheralsData
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1SetInstancePeripherals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SetInstancePeripheralsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **peripheralsData** | [**PeripheralsData**](PeripheralsData.md) | New peripherals state | 

### Return type

[**PeripheralsData**](PeripheralsData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SetInstanceState

> V1SetInstanceState(ctx, instanceId).V1SetStateBody(v1SetStateBody).Execute()

Set state of Instance

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    v1SetStateBody := *openapiclient.NewV1SetStateBody(openapiclient.InstanceState("on")) // V1SetStateBody | Desired State

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1SetInstanceState(context.Background(), instanceId).V1SetStateBody(v1SetStateBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1SetInstanceState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SetInstanceStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1SetStateBody** | [**V1SetStateBody**](V1SetStateBody.md) | Desired State | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SetSnapshotPermissions

> Snapshot V1SetSnapshotPermissions(ctx, snapshotId).SnapshotPermissionsRequestPayload(snapshotPermissionsRequestPayload).Execute()

Set member list



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    snapshotId := "snapshotId_example" // string | Snapshot ID - uuid
    snapshotPermissionsRequestPayload := *openapiclient.NewSnapshotPermissionsRequestPayload([]string{"Members_example"}) // SnapshotPermissionsRequestPayload | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1SetSnapshotPermissions(context.Background(), snapshotId).SnapshotPermissionsRequestPayload(snapshotPermissionsRequestPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1SetSnapshotPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1SetSnapshotPermissions`: Snapshot
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1SetSnapshotPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | Snapshot ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SetSnapshotPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **snapshotPermissionsRequestPayload** | [**SnapshotPermissionsRequestPayload**](SnapshotPermissionsRequestPayload.md) |  | 

### Return type

[**Snapshot**](Snapshot.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ShareSnapshot

> Snapshot V1ShareSnapshot(ctx, snapshotId).PostShareSnapshotRequestPayload(postShareSnapshotRequestPayload).Execute()

Share snapshot



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    snapshotId := "snapshotId_example" // string | Snapshot ID - uuid
    postShareSnapshotRequestPayload := *openapiclient.NewPostShareSnapshotRequestPayload("SharingType_example") // PostShareSnapshotRequestPayload | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1ShareSnapshot(context.Background(), snapshotId).PostShareSnapshotRequestPayload(postShareSnapshotRequestPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1ShareSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ShareSnapshot`: Snapshot
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1ShareSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | Snapshot ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ShareSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postShareSnapshotRequestPayload** | [**PostShareSnapshotRequestPayload**](PostShareSnapshotRequestPayload.md) |  | 

### Return type

[**Snapshot**](Snapshot.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SnapshotRename

> Snapshot V1SnapshotRename(ctx, snapshotId).SnapshotCreationOptions(snapshotCreationOptions).Execute()

Rename a Snapshot



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    snapshotId := "snapshotId_example" // string | Snapshot ID - uuid
    snapshotCreationOptions := *openapiclient.NewSnapshotCreationOptions("Name_example") // SnapshotCreationOptions | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1SnapshotRename(context.Background(), snapshotId).SnapshotCreationOptions(snapshotCreationOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1SnapshotRename``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1SnapshotRename`: Snapshot
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1SnapshotRename`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | Snapshot ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SnapshotRenameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **snapshotCreationOptions** | [**SnapshotCreationOptions**](SnapshotCreationOptions.md) |  | 

### Return type

[**Snapshot**](Snapshot.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1StartCoreTrace

> V1StartCoreTrace(ctx, instanceId).Execute()

Start CoreTrace on an instance

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1StartCoreTrace(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1StartCoreTrace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1StartCoreTraceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1StartHyperTrace

> V1StartHyperTrace(ctx, instanceId).BtraceEnableOptions(btraceEnableOptions).Execute()

Start HyperTrace on an instance

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    btraceEnableOptions := *openapiclient.NewBtraceEnableOptions() // BtraceEnableOptions | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1StartHyperTrace(context.Background(), instanceId).BtraceEnableOptions(btraceEnableOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1StartHyperTrace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1StartHyperTraceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **btraceEnableOptions** | [**BtraceEnableOptions**](BtraceEnableOptions.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1StartInstance

> V1StartInstance(ctx, instanceId).InstanceStartOptions(instanceStartOptions).Execute()

Start an Instance

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    instanceStartOptions := *openapiclient.NewInstanceStartOptions() // InstanceStartOptions | Start options (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1StartInstance(context.Background(), instanceId).InstanceStartOptions(instanceStartOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1StartInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1StartInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instanceStartOptions** | [**InstanceStartOptions**](InstanceStartOptions.md) | Start options | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1StartNetdump

> V1StartNetdump(ctx, instanceId).NetdumpFilter(netdumpFilter).Execute()

Start Enhanced Network Monitor on an instance.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    netdumpFilter := *openapiclient.NewNetdumpFilter() // NetdumpFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1StartNetdump(context.Background(), instanceId).NetdumpFilter(netdumpFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1StartNetdump``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1StartNetdumpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **netdumpFilter** | [**NetdumpFilter**](NetdumpFilter.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1StartNetworkMonitor

> V1StartNetworkMonitor(ctx, instanceId).SslsplitFilter(sslsplitFilter).Execute()

Start Network Monitor on an instance.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    sslsplitFilter := *openapiclient.NewSslsplitFilter() // SslsplitFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1StartNetworkMonitor(context.Background(), instanceId).SslsplitFilter(sslsplitFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1StartNetworkMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1StartNetworkMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sslsplitFilter** | [**SslsplitFilter**](SslsplitFilter.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1StopCoreTrace

> V1StopCoreTrace(ctx, instanceId).Execute()

Stop CoreTrace on an instance.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1StopCoreTrace(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1StopCoreTrace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1StopCoreTraceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1StopHyperTrace

> V1StopHyperTrace(ctx, instanceId).Execute()

Stop HyperTrace on an instance.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1StopHyperTrace(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1StopHyperTrace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1StopHyperTraceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1StopInstance

> V1StopInstance(ctx, instanceId).InstanceStopOptions(instanceStopOptions).Execute()

Stop an Instance

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    instanceStopOptions := *openapiclient.NewInstanceStopOptions() // InstanceStopOptions | Stop options (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1StopInstance(context.Background(), instanceId).InstanceStopOptions(instanceStopOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1StopInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1StopInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instanceStopOptions** | [**InstanceStopOptions**](InstanceStopOptions.md) | Stop options | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1StopNetdump

> V1StopNetdump(ctx, instanceId).Execute()

Stop Enhanced Network Monitor on an instance.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1StopNetdump(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1StopNetdump``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1StopNetdumpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1StopNetworkMonitor

> V1StopNetworkMonitor(ctx, instanceId).Execute()

Stop Network Monitor on an instance.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1StopNetworkMonitor(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1StopNetworkMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1StopNetworkMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1TeamChange

> V1TeamChange(ctx, teamId).CreateTeam(createTeam).Execute()

Update team



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    teamId := "teamId_example" // string | Team ID - uuid
    createTeam := *openapiclient.NewCreateTeam("Name_example") // CreateTeam | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1TeamChange(context.Background(), teamId).CreateTeam(createTeam).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1TeamChange``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | Team ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1TeamChangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTeam** | [**CreateTeam**](CreateTeam.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1TeamCreate

> TeamCreate V1TeamCreate(ctx).CreateTeam(createTeam).Execute()

Create team



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    createTeam := *openapiclient.NewCreateTeam("Name_example") // CreateTeam | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1TeamCreate(context.Background()).CreateTeam(createTeam).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1TeamCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1TeamCreate`: TeamCreate
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1TeamCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TeamCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTeam** | [**CreateTeam**](CreateTeam.md) |  | 

### Return type

[**TeamCreate**](TeamCreate.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1TeamDelete

> V1TeamDelete(ctx, teamId).Execute()

Delete team



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    teamId := "teamId_example" // string | Team ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1TeamDelete(context.Background(), teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1TeamDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | Team ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1TeamDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1Teams

> []Team V1Teams(ctx).Execute()

Get teams



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1Teams(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1Teams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1Teams`: []Team
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1Teams`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1TeamsRequest struct via the builder pattern


### Return type

[**[]Team**](Team.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UnpauseInstance

> V1UnpauseInstance(ctx, instanceId).Execute()

Unpause an Instance

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1UnpauseInstance(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1UnpauseInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UnpauseInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UpdateDomainAuthProvider

> DomainAuthProviderResponse V1UpdateDomainAuthProvider(ctx, domainId, providerId).DomainAuthProviderRequest(domainAuthProviderRequest).Execute()

Update an auth provider for a domain

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    domainId := "domainId_example" // string | Domain ID - uuid
    providerId := "providerId_example" // string | Provider ID - uuid
    domainAuthProviderRequest := *openapiclient.NewDomainAuthProviderRequest("ProviderType_example", false) // DomainAuthProviderRequest | Request Data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1UpdateDomainAuthProvider(context.Background(), domainId, providerId).DomainAuthProviderRequest(domainAuthProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1UpdateDomainAuthProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1UpdateDomainAuthProvider`: DomainAuthProviderResponse
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1UpdateDomainAuthProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | Domain ID - uuid | 
**providerId** | **string** | Provider ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UpdateDomainAuthProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **domainAuthProviderRequest** | [**DomainAuthProviderRequest**](DomainAuthProviderRequest.md) | Request Data | 

### Return type

[**DomainAuthProviderResponse**](DomainAuthProviderResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UpdateExtension

> V1UpdateExtension(ctx, extensionId).UpdateExtension(updateExtension).Execute()

Update an existing extension

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    extensionId := "extensionId_example" // string | Extension ID
    updateExtension := *openapiclient.NewUpdateExtension() // UpdateExtension | application/json

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1UpdateExtension(context.Background(), extensionId).UpdateExtension(updateExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1UpdateExtension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **string** | Extension ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UpdateExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateExtension** | [**UpdateExtension**](UpdateExtension.md) | application/json | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UpdateHook

> Hook V1UpdateHook(ctx, hookId).V1CreateHookParameters(v1CreateHookParameters).Execute()

Update an existing hypervisor hook

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    hookId := "hookId_example" // string | Hook ID
    v1CreateHookParameters := *openapiclient.NewV1CreateHookParameters("Label_example", "Address_example", "Patch_example", "PatchType_example") // V1CreateHookParameters | application/json

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1UpdateHook(context.Background(), hookId).V1CreateHookParameters(v1CreateHookParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1UpdateHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1UpdateHook`: Hook
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1UpdateHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hookId** | **string** | Hook ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UpdateHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1CreateHookParameters** | [**V1CreateHookParameters**](V1CreateHookParameters.md) | application/json | 

### Return type

[**Hook**](Hook.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UpdateNetworkConnection

> NetworkConnection V1UpdateNetworkConnection(ctx, id).UpdateNetworkConnectionOptions(updateNetworkConnectionOptions).Force(force).Execute()

Update Network Connection



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    id := "id_example" // string | Network Connection Identifier - uuid
    updateNetworkConnectionOptions := *openapiclient.NewUpdateNetworkConnectionOptions("Name_example") // UpdateNetworkConnectionOptions | Network Connection Options
    force := true // bool | Force deletion (true only or not present) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1UpdateNetworkConnection(context.Background(), id).UpdateNetworkConnectionOptions(updateNetworkConnectionOptions).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1UpdateNetworkConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1UpdateNetworkConnection`: NetworkConnection
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1UpdateNetworkConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Network Connection Identifier - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UpdateNetworkConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkConnectionOptions** | [**UpdateNetworkConnectionOptions**](UpdateNetworkConnectionOptions.md) | Network Connection Options | 
 **force** | **bool** | Force deletion (true only or not present) | 

### Return type

[**NetworkConnection**](NetworkConnection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UpdateProject

> Project V1UpdateProject(ctx, projectId).Project(project).Execute()

Update a Project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    projectId := "projectId_example" // string | Project ID - uuid
    project := *openapiclient.NewProject("Id_example") // Project | Updated Project Settings

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1UpdateProject(context.Background(), projectId).Project(project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1UpdateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1UpdateProject`: Project
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1UpdateProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UpdateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **project** | [**Project**](Project.md) | Updated Project Settings | 

### Return type

[**Project**](Project.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UpdateProjectSettings

> V1UpdateProjectSettings(ctx, projectId).ProjectSettings(projectSettings).Execute()

Change Project Settings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    projectId := "projectId_example" // string | Project ID - uuid
    projectSettings := *openapiclient.NewProjectSettings(false, false) // ProjectSettings | New settings

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1UpdateProjectSettings(context.Background(), projectId).ProjectSettings(projectSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1UpdateProjectSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UpdateProjectSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectSettings** | [**ProjectSettings**](ProjectSettings.md) | New settings | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UpdateUser

> map[string]interface{} V1UpdateUser(ctx, userId).Body(body).Execute()

Update User

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    userId := "userId_example" // string | userId - uuid
    body := map[string]interface{}{ ... } // map[string]interface{} | User data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1UpdateUser(context.Background(), userId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1UpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1UpdateUser`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | userId - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | User data | 

### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UpgradeInstance

> V1UpgradeInstance(ctx, instanceId).InstanceUpgradeBody(instanceUpgradeBody).Execute()

Upgrade iOS version

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | 
    instanceUpgradeBody := *openapiclient.NewInstanceUpgradeBody("Os_example") // InstanceUpgradeBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1UpgradeInstance(context.Background(), instanceId).InstanceUpgradeBody(instanceUpgradeBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1UpgradeInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UpgradeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instanceUpgradeBody** | [**InstanceUpgradeBody**](InstanceUpgradeBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UploadImageData

> Image V1UploadImageData(ctx, imageId).Body(body).Execute()

Upload Image Data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    imageId := "imageId_example" // string | Image ID - uuid
    body := "body_example" // string | Uploaded Image

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1UploadImageData(context.Background(), imageId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1UploadImageData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1UploadImageData`: Image
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1UploadImageData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | Image ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UploadImageDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | Uploaded Image | 

### Return type

[**Image**](Image.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: binary
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UserAgreeTerms

> AgreedAck V1UserAgreeTerms(ctx).Execute()

Consent to the current terms and conditions

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1UserAgreeTerms(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1UserAgreeTerms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1UserAgreeTerms`: AgreedAck
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1UserAgreeTerms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1UserAgreeTermsRequest struct via the builder pattern


### Return type

[**AgreedAck**](AgreedAck.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UsersChangePasswordPost

> V1UsersChangePasswordPost(ctx).PasswordChangeBody(passwordChangeBody).Execute()

Change User Password



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    passwordChangeBody := *openapiclient.NewPasswordChangeBody("User_example", "OldPassword_example", "NewPassword_example") // PasswordChangeBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1UsersChangePasswordPost(context.Background()).PasswordChangeBody(passwordChangeBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1UsersChangePasswordPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1UsersChangePasswordPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordChangeBody** | [**PasswordChangeBody**](PasswordChangeBody.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UsersLogin

> Token V1UsersLogin(ctx).Credentials(credentials).Execute()

Log In

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    credentials := *openapiclient.NewCredentials("Username_example", "Password_example") // Credentials | Authorization data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1UsersLogin(context.Background()).Credentials(credentials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1UsersLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1UsersLogin`: Token
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1UsersLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1UsersLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentials** | [**Credentials**](Credentials.md) | Authorization data | 

### Return type

[**Token**](Token.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1WebPlayerAllowedDomains

> V1WebPlayerAllowedDomains(ctx).Execute()

Retrieve the list of allowed domains for all Webplayer sessions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorelliumApi.V1WebPlayerAllowedDomains(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1WebPlayerAllowedDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1WebPlayerAllowedDomainsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1WebPlayerCreateSession

> WebPlayerSession V1WebPlayerCreateSession(ctx).WebPlayerCreateSessionRequest(webPlayerCreateSessionRequest).Execute()

Create a new Webplayer Session



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    webPlayerCreateSessionRequest := *openapiclient.NewWebPlayerCreateSessionRequest("ProjectId_example", "InstanceId_example", float32(123), *openapiclient.NewFeatures()) // WebPlayerCreateSessionRequest | Request Data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1WebPlayerCreateSession(context.Background()).WebPlayerCreateSessionRequest(webPlayerCreateSessionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1WebPlayerCreateSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1WebPlayerCreateSession`: WebPlayerSession
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1WebPlayerCreateSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1WebPlayerCreateSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webPlayerCreateSessionRequest** | [**WebPlayerCreateSessionRequest**](WebPlayerCreateSessionRequest.md) | Request Data | 

### Return type

[**WebPlayerSession**](WebPlayerSession.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1WebPlayerDestroySession

> []string V1WebPlayerDestroySession(ctx, sessionId).Execute()

Tear down a Webplayer Session



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    sessionId := "sessionId_example" // string | Webplayer Session identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1WebPlayerDestroySession(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1WebPlayerDestroySession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1WebPlayerDestroySession`: []string
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1WebPlayerDestroySession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Webplayer Session identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1WebPlayerDestroySessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1WebPlayerListSessions

> []WebPlayerSession V1WebPlayerListSessions(ctx).Execute()

List all Webplayer sessions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1WebPlayerListSessions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1WebPlayerListSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1WebPlayerListSessions`: []WebPlayerSession
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1WebPlayerListSessions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1WebPlayerListSessionsRequest struct via the builder pattern


### Return type

[**[]WebPlayerSession**](WebPlayerSession.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1WebPlayerSessionInfo

> WebPlayerSession V1WebPlayerSessionInfo(ctx, sessionId).Execute()

Retrieve Webplayer Session Information



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    sessionId := "sessionId_example" // string | Webplayer Session identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V1WebPlayerSessionInfo(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V1WebPlayerSessionInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1WebPlayerSessionInfo`: WebPlayerSession
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V1WebPlayerSessionInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Webplayer Session identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1WebPlayerSessionInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebPlayerSession**](WebPlayerSession.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2GetInstanceQuickConnectCommand

> string V2GetInstanceQuickConnectCommand(ctx, instanceId).Execute()

Recommended SSH Command for Quick Connect

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V2GetInstanceQuickConnectCommand(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V2GetInstanceQuickConnectCommand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2GetInstanceQuickConnectCommand`: string
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V2GetInstanceQuickConnectCommand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2GetInstanceQuickConnectCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2GetInstanceState

> InstanceState V2GetInstanceState(ctx, instanceId).ReturnAttr(returnAttr).Execute()

Get state of Instance

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/corellium/go-corellium-api-client"
)

func main() {
    instanceId := "instanceId_example" // string | Instance ID - uuid
    returnAttr := []string{"Inner_example"} // []string | The attributes to return. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorelliumApi.V2GetInstanceState(context.Background(), instanceId).ReturnAttr(returnAttr).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorelliumApi.V2GetInstanceState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2GetInstanceState`: InstanceState
    fmt.Fprintf(os.Stdout, "Response from `CorelliumApi.V2GetInstanceState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | Instance ID - uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2GetInstanceStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnAttr** | **[]string** | The attributes to return. | 

### Return type

[**InstanceState**](InstanceState.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

