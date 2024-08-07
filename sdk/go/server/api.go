/*
 * Kowabunga API documentation
 *
 * Kvm Orchestrator With A BUNch of Goods Added
 *
 * API version: 0.42.0
 * Contact: csops@dalet.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

import (
	"context"
	"net/http"
)



// AdapterAPIRouter defines the required methods for binding the api requests to a responses for the AdapterAPI
// The AdapterAPIRouter implementation should parse necessary information from the http request,
// pass the data to a AdapterAPIServicer to perform the required actions, then write the service results to the http response.
type AdapterAPIRouter interface { 
	DeleteAdapter(http.ResponseWriter, *http.Request)
	ListAdapters(http.ResponseWriter, *http.Request)
	ReadAdapter(http.ResponseWriter, *http.Request)
	UpdateAdapter(http.ResponseWriter, *http.Request)
}
// AgentAPIRouter defines the required methods for binding the api requests to a responses for the AgentAPI
// The AgentAPIRouter implementation should parse necessary information from the http request,
// pass the data to a AgentAPIServicer to perform the required actions, then write the service results to the http response.
type AgentAPIRouter interface { 
	CreateAgent(http.ResponseWriter, *http.Request)
	DeleteAgent(http.ResponseWriter, *http.Request)
	ListAgents(http.ResponseWriter, *http.Request)
	ReadAgent(http.ResponseWriter, *http.Request)
	SetAgentApiToken(http.ResponseWriter, *http.Request)
	UpdateAgent(http.ResponseWriter, *http.Request)
}
// GroupAPIRouter defines the required methods for binding the api requests to a responses for the GroupAPI
// The GroupAPIRouter implementation should parse necessary information from the http request,
// pass the data to a GroupAPIServicer to perform the required actions, then write the service results to the http response.
type GroupAPIRouter interface { 
	CreateGroup(http.ResponseWriter, *http.Request)
	DeleteGroup(http.ResponseWriter, *http.Request)
	ListGroups(http.ResponseWriter, *http.Request)
	ReadGroup(http.ResponseWriter, *http.Request)
	UpdateGroup(http.ResponseWriter, *http.Request)
}
// HostAPIRouter defines the required methods for binding the api requests to a responses for the HostAPI
// The HostAPIRouter implementation should parse necessary information from the http request,
// pass the data to a HostAPIServicer to perform the required actions, then write the service results to the http response.
type HostAPIRouter interface { 
	DeleteHost(http.ResponseWriter, *http.Request)
	ListHostInstances(http.ResponseWriter, *http.Request)
	ListHosts(http.ResponseWriter, *http.Request)
	ReadHost(http.ResponseWriter, *http.Request)
	ReadHostCaps(http.ResponseWriter, *http.Request)
	UpdateHost(http.ResponseWriter, *http.Request)
}
// InstanceAPIRouter defines the required methods for binding the api requests to a responses for the InstanceAPI
// The InstanceAPIRouter implementation should parse necessary information from the http request,
// pass the data to a InstanceAPIServicer to perform the required actions, then write the service results to the http response.
type InstanceAPIRouter interface { 
	DeleteInstance(http.ResponseWriter, *http.Request)
	ListInstances(http.ResponseWriter, *http.Request)
	ReadInstance(http.ResponseWriter, *http.Request)
	ReadInstanceRemoteConnection(http.ResponseWriter, *http.Request)
	ReadInstanceState(http.ResponseWriter, *http.Request)
	RebootInstance(http.ResponseWriter, *http.Request)
	ResetInstance(http.ResponseWriter, *http.Request)
	ResumeInstance(http.ResponseWriter, *http.Request)
	ShutdownInstance(http.ResponseWriter, *http.Request)
	StartInstance(http.ResponseWriter, *http.Request)
	StopInstance(http.ResponseWriter, *http.Request)
	SuspendInstance(http.ResponseWriter, *http.Request)
	UpdateInstance(http.ResponseWriter, *http.Request)
}
// KceAPIRouter defines the required methods for binding the api requests to a responses for the KceAPI
// The KceAPIRouter implementation should parse necessary information from the http request,
// pass the data to a KceAPIServicer to perform the required actions, then write the service results to the http response.
type KceAPIRouter interface { 
	DeleteKCE(http.ResponseWriter, *http.Request)
	ListKCEs(http.ResponseWriter, *http.Request)
	ReadKCE(http.ResponseWriter, *http.Request)
	ReadKCEState(http.ResponseWriter, *http.Request)
	RebootKCE(http.ResponseWriter, *http.Request)
	ResetKCE(http.ResponseWriter, *http.Request)
	ResumeKCE(http.ResponseWriter, *http.Request)
	ShutdownKCE(http.ResponseWriter, *http.Request)
	StartKCE(http.ResponseWriter, *http.Request)
	StopKCE(http.ResponseWriter, *http.Request)
	SuspendKCE(http.ResponseWriter, *http.Request)
	UpdateKCE(http.ResponseWriter, *http.Request)
}
// KfsAPIRouter defines the required methods for binding the api requests to a responses for the KfsAPI
// The KfsAPIRouter implementation should parse necessary information from the http request,
// pass the data to a KfsAPIServicer to perform the required actions, then write the service results to the http response.
type KfsAPIRouter interface { 
	DeleteKFS(http.ResponseWriter, *http.Request)
	ListKFSs(http.ResponseWriter, *http.Request)
	ReadKFS(http.ResponseWriter, *http.Request)
	UpdateKFS(http.ResponseWriter, *http.Request)
}
// KgwAPIRouter defines the required methods for binding the api requests to a responses for the KgwAPI
// The KgwAPIRouter implementation should parse necessary information from the http request,
// pass the data to a KgwAPIServicer to perform the required actions, then write the service results to the http response.
type KgwAPIRouter interface { 
	DeleteKGW(http.ResponseWriter, *http.Request)
	ListKGWs(http.ResponseWriter, *http.Request)
	ReadKGW(http.ResponseWriter, *http.Request)
	UpdateKGW(http.ResponseWriter, *http.Request)
}
// KonveyAPIRouter defines the required methods for binding the api requests to a responses for the KonveyAPI
// The KonveyAPIRouter implementation should parse necessary information from the http request,
// pass the data to a KonveyAPIServicer to perform the required actions, then write the service results to the http response.
type KonveyAPIRouter interface { 
	DeleteKonvey(http.ResponseWriter, *http.Request)
	ListKonveys(http.ResponseWriter, *http.Request)
	ReadKonvey(http.ResponseWriter, *http.Request)
	UpdateKonvey(http.ResponseWriter, *http.Request)
}
// NetgwAPIRouter defines the required methods for binding the api requests to a responses for the NetgwAPI
// The NetgwAPIRouter implementation should parse necessary information from the http request,
// pass the data to a NetgwAPIServicer to perform the required actions, then write the service results to the http response.
type NetgwAPIRouter interface { 
	DeleteNetGW(http.ResponseWriter, *http.Request)
	ListNetGWs(http.ResponseWriter, *http.Request)
	ReadNetGW(http.ResponseWriter, *http.Request)
	UpdateNetGW(http.ResponseWriter, *http.Request)
}
// NfsAPIRouter defines the required methods for binding the api requests to a responses for the NfsAPI
// The NfsAPIRouter implementation should parse necessary information from the http request,
// pass the data to a NfsAPIServicer to perform the required actions, then write the service results to the http response.
type NfsAPIRouter interface { 
	DeleteStorageNFS(http.ResponseWriter, *http.Request)
	ListStorageNFSKFSs(http.ResponseWriter, *http.Request)
	ListStorageNFSs(http.ResponseWriter, *http.Request)
	ReadStorageNFS(http.ResponseWriter, *http.Request)
	UpdateStorageNFS(http.ResponseWriter, *http.Request)
}
// PoolAPIRouter defines the required methods for binding the api requests to a responses for the PoolAPI
// The PoolAPIRouter implementation should parse necessary information from the http request,
// pass the data to a PoolAPIServicer to perform the required actions, then write the service results to the http response.
type PoolAPIRouter interface { 
	CreateTemplate(http.ResponseWriter, *http.Request)
	DeleteStoragePool(http.ResponseWriter, *http.Request)
	ListStoragePoolTemplates(http.ResponseWriter, *http.Request)
	ListStoragePoolVolumes(http.ResponseWriter, *http.Request)
	ListStoragePools(http.ResponseWriter, *http.Request)
	ReadStoragePool(http.ResponseWriter, *http.Request)
	SetStoragePoolDefaultTemplate(http.ResponseWriter, *http.Request)
	UpdateStoragePool(http.ResponseWriter, *http.Request)
}
// ProjectAPIRouter defines the required methods for binding the api requests to a responses for the ProjectAPI
// The ProjectAPIRouter implementation should parse necessary information from the http request,
// pass the data to a ProjectAPIServicer to perform the required actions, then write the service results to the http response.
type ProjectAPIRouter interface { 
	CreateProject(http.ResponseWriter, *http.Request)
	CreateProjectDnsRecord(http.ResponseWriter, *http.Request)
	CreateProjectRegionKFS(http.ResponseWriter, *http.Request)
	CreateProjectRegionKGW(http.ResponseWriter, *http.Request)
	CreateProjectRegionKonvey(http.ResponseWriter, *http.Request)
	CreateProjectRegionVolume(http.ResponseWriter, *http.Request)
	CreateProjectZoneInstance(http.ResponseWriter, *http.Request)
	CreateProjectZoneKCE(http.ResponseWriter, *http.Request)
	CreateProjectZoneKonvey(http.ResponseWriter, *http.Request)
	DeleteProject(http.ResponseWriter, *http.Request)
	ListProjectDnsRecords(http.ResponseWriter, *http.Request)
	ListProjectRegionKFSs(http.ResponseWriter, *http.Request)
	ListProjectRegionKGWs(http.ResponseWriter, *http.Request)
	ListProjectRegionKonveys(http.ResponseWriter, *http.Request)
	ListProjectRegionVolumes(http.ResponseWriter, *http.Request)
	ListProjectZoneInstances(http.ResponseWriter, *http.Request)
	ListProjectZoneKCEs(http.ResponseWriter, *http.Request)
	ListProjectZoneKonveys(http.ResponseWriter, *http.Request)
	ListProjects(http.ResponseWriter, *http.Request)
	ReadProject(http.ResponseWriter, *http.Request)
	ReadProjectCost(http.ResponseWriter, *http.Request)
	ReadProjectUsage(http.ResponseWriter, *http.Request)
	UpdateProject(http.ResponseWriter, *http.Request)
}
// RecordAPIRouter defines the required methods for binding the api requests to a responses for the RecordAPI
// The RecordAPIRouter implementation should parse necessary information from the http request,
// pass the data to a RecordAPIServicer to perform the required actions, then write the service results to the http response.
type RecordAPIRouter interface { 
	DeleteDnsRecord(http.ResponseWriter, *http.Request)
	ReadDnsRecord(http.ResponseWriter, *http.Request)
	UpdateDnsRecord(http.ResponseWriter, *http.Request)
}
// RegionAPIRouter defines the required methods for binding the api requests to a responses for the RegionAPI
// The RegionAPIRouter implementation should parse necessary information from the http request,
// pass the data to a RegionAPIServicer to perform the required actions, then write the service results to the http response.
type RegionAPIRouter interface { 
	CreateNetGW(http.ResponseWriter, *http.Request)
	CreateRegion(http.ResponseWriter, *http.Request)
	CreateStorageNFS(http.ResponseWriter, *http.Request)
	CreateStoragePool(http.ResponseWriter, *http.Request)
	CreateVNet(http.ResponseWriter, *http.Request)
	CreateZone(http.ResponseWriter, *http.Request)
	DeleteRegion(http.ResponseWriter, *http.Request)
	ListRegionNetGWs(http.ResponseWriter, *http.Request)
	ListRegionStorageNFSs(http.ResponseWriter, *http.Request)
	ListRegionStoragePools(http.ResponseWriter, *http.Request)
	ListRegionVNets(http.ResponseWriter, *http.Request)
	ListRegionZones(http.ResponseWriter, *http.Request)
	ListRegions(http.ResponseWriter, *http.Request)
	ReadRegion(http.ResponseWriter, *http.Request)
	SetRegionDefaultStorageNFS(http.ResponseWriter, *http.Request)
	SetRegionDefaultStoragePool(http.ResponseWriter, *http.Request)
	UpdateRegion(http.ResponseWriter, *http.Request)
}
// SubnetAPIRouter defines the required methods for binding the api requests to a responses for the SubnetAPI
// The SubnetAPIRouter implementation should parse necessary information from the http request,
// pass the data to a SubnetAPIServicer to perform the required actions, then write the service results to the http response.
type SubnetAPIRouter interface { 
	CreateAdapter(http.ResponseWriter, *http.Request)
	DeleteSubnet(http.ResponseWriter, *http.Request)
	ListSubnetAdapters(http.ResponseWriter, *http.Request)
	ListSubnets(http.ResponseWriter, *http.Request)
	ReadSubnet(http.ResponseWriter, *http.Request)
	UpdateSubnet(http.ResponseWriter, *http.Request)
}
// TemplateAPIRouter defines the required methods for binding the api requests to a responses for the TemplateAPI
// The TemplateAPIRouter implementation should parse necessary information from the http request,
// pass the data to a TemplateAPIServicer to perform the required actions, then write the service results to the http response.
type TemplateAPIRouter interface { 
	DeleteTemplate(http.ResponseWriter, *http.Request)
	ListTemplates(http.ResponseWriter, *http.Request)
	ReadTemplate(http.ResponseWriter, *http.Request)
	UpdateTemplate(http.ResponseWriter, *http.Request)
}
// TokenAPIRouter defines the required methods for binding the api requests to a responses for the TokenAPI
// The TokenAPIRouter implementation should parse necessary information from the http request,
// pass the data to a TokenAPIServicer to perform the required actions, then write the service results to the http response.
type TokenAPIRouter interface { 
	DeleteApiToken(http.ResponseWriter, *http.Request)
	ListApiTokens(http.ResponseWriter, *http.Request)
	ReadApiToken(http.ResponseWriter, *http.Request)
	UpdateApiToken(http.ResponseWriter, *http.Request)
}
// UserAPIRouter defines the required methods for binding the api requests to a responses for the UserAPI
// The UserAPIRouter implementation should parse necessary information from the http request,
// pass the data to a UserAPIServicer to perform the required actions, then write the service results to the http response.
type UserAPIRouter interface { 
	CreateUser(http.ResponseWriter, *http.Request)
	DeleteUser(http.ResponseWriter, *http.Request)
	ListUsers(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)
	ReadUser(http.ResponseWriter, *http.Request)
	ResetUserPassword(http.ResponseWriter, *http.Request)
	SetUserApiToken(http.ResponseWriter, *http.Request)
	SetUserPassword(http.ResponseWriter, *http.Request)
	UpdateUser(http.ResponseWriter, *http.Request)
}
// VnetAPIRouter defines the required methods for binding the api requests to a responses for the VnetAPI
// The VnetAPIRouter implementation should parse necessary information from the http request,
// pass the data to a VnetAPIServicer to perform the required actions, then write the service results to the http response.
type VnetAPIRouter interface { 
	CreateSubnet(http.ResponseWriter, *http.Request)
	DeleteVNet(http.ResponseWriter, *http.Request)
	ListVNetSubnets(http.ResponseWriter, *http.Request)
	ListVNets(http.ResponseWriter, *http.Request)
	ReadVNet(http.ResponseWriter, *http.Request)
	SetVNetDefaultSubnet(http.ResponseWriter, *http.Request)
	UpdateVNet(http.ResponseWriter, *http.Request)
}
// VolumeAPIRouter defines the required methods for binding the api requests to a responses for the VolumeAPI
// The VolumeAPIRouter implementation should parse necessary information from the http request,
// pass the data to a VolumeAPIServicer to perform the required actions, then write the service results to the http response.
type VolumeAPIRouter interface { 
	DeleteVolume(http.ResponseWriter, *http.Request)
	ListVolumes(http.ResponseWriter, *http.Request)
	ReadVolume(http.ResponseWriter, *http.Request)
	UpdateVolume(http.ResponseWriter, *http.Request)
}
// ZoneAPIRouter defines the required methods for binding the api requests to a responses for the ZoneAPI
// The ZoneAPIRouter implementation should parse necessary information from the http request,
// pass the data to a ZoneAPIServicer to perform the required actions, then write the service results to the http response.
type ZoneAPIRouter interface { 
	CreateHost(http.ResponseWriter, *http.Request)
	DeleteZone(http.ResponseWriter, *http.Request)
	ListZoneHosts(http.ResponseWriter, *http.Request)
	ListZones(http.ResponseWriter, *http.Request)
	ReadZone(http.ResponseWriter, *http.Request)
	UpdateZone(http.ResponseWriter, *http.Request)
}


// AdapterAPIServicer defines the api actions for the AdapterAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type AdapterAPIServicer interface { 
	DeleteAdapter(context.Context, string) (ImplResponse, error)
	ListAdapters(context.Context) (ImplResponse, error)
	ReadAdapter(context.Context, string) (ImplResponse, error)
	UpdateAdapter(context.Context, string, Adapter) (ImplResponse, error)
}


// AgentAPIServicer defines the api actions for the AgentAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type AgentAPIServicer interface { 
	CreateAgent(context.Context, Agent) (ImplResponse, error)
	DeleteAgent(context.Context, string) (ImplResponse, error)
	ListAgents(context.Context) (ImplResponse, error)
	ReadAgent(context.Context, string) (ImplResponse, error)
	SetAgentApiToken(context.Context, string, bool, string) (ImplResponse, error)
	UpdateAgent(context.Context, string, Agent) (ImplResponse, error)
}


// GroupAPIServicer defines the api actions for the GroupAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type GroupAPIServicer interface { 
	CreateGroup(context.Context, Group) (ImplResponse, error)
	DeleteGroup(context.Context, string) (ImplResponse, error)
	ListGroups(context.Context) (ImplResponse, error)
	ReadGroup(context.Context, string) (ImplResponse, error)
	UpdateGroup(context.Context, string, Group) (ImplResponse, error)
}


// HostAPIServicer defines the api actions for the HostAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type HostAPIServicer interface { 
	DeleteHost(context.Context, string) (ImplResponse, error)
	ListHostInstances(context.Context, string) (ImplResponse, error)
	ListHosts(context.Context) (ImplResponse, error)
	ReadHost(context.Context, string) (ImplResponse, error)
	ReadHostCaps(context.Context, string) (ImplResponse, error)
	UpdateHost(context.Context, string, Host) (ImplResponse, error)
}


// InstanceAPIServicer defines the api actions for the InstanceAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type InstanceAPIServicer interface { 
	DeleteInstance(context.Context, string) (ImplResponse, error)
	ListInstances(context.Context) (ImplResponse, error)
	ReadInstance(context.Context, string) (ImplResponse, error)
	ReadInstanceRemoteConnection(context.Context, string) (ImplResponse, error)
	ReadInstanceState(context.Context, string) (ImplResponse, error)
	RebootInstance(context.Context, string) (ImplResponse, error)
	ResetInstance(context.Context, string) (ImplResponse, error)
	ResumeInstance(context.Context, string) (ImplResponse, error)
	ShutdownInstance(context.Context, string) (ImplResponse, error)
	StartInstance(context.Context, string) (ImplResponse, error)
	StopInstance(context.Context, string) (ImplResponse, error)
	SuspendInstance(context.Context, string) (ImplResponse, error)
	UpdateInstance(context.Context, string, Instance) (ImplResponse, error)
}


// KceAPIServicer defines the api actions for the KceAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type KceAPIServicer interface { 
	DeleteKCE(context.Context, string) (ImplResponse, error)
	ListKCEs(context.Context) (ImplResponse, error)
	ReadKCE(context.Context, string) (ImplResponse, error)
	ReadKCEState(context.Context, string) (ImplResponse, error)
	RebootKCE(context.Context, string) (ImplResponse, error)
	ResetKCE(context.Context, string) (ImplResponse, error)
	ResumeKCE(context.Context, string) (ImplResponse, error)
	ShutdownKCE(context.Context, string) (ImplResponse, error)
	StartKCE(context.Context, string) (ImplResponse, error)
	StopKCE(context.Context, string) (ImplResponse, error)
	SuspendKCE(context.Context, string) (ImplResponse, error)
	UpdateKCE(context.Context, string, Kce) (ImplResponse, error)
}


// KfsAPIServicer defines the api actions for the KfsAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type KfsAPIServicer interface { 
	DeleteKFS(context.Context, string) (ImplResponse, error)
	ListKFSs(context.Context) (ImplResponse, error)
	ReadKFS(context.Context, string) (ImplResponse, error)
	UpdateKFS(context.Context, string, Kfs) (ImplResponse, error)
}


// KgwAPIServicer defines the api actions for the KgwAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type KgwAPIServicer interface { 
	DeleteKGW(context.Context, string) (ImplResponse, error)
	ListKGWs(context.Context) (ImplResponse, error)
	ReadKGW(context.Context, string) (ImplResponse, error)
	UpdateKGW(context.Context, string, Kgw) (ImplResponse, error)
}


// KonveyAPIServicer defines the api actions for the KonveyAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type KonveyAPIServicer interface { 
	DeleteKonvey(context.Context, string) (ImplResponse, error)
	ListKonveys(context.Context) (ImplResponse, error)
	ReadKonvey(context.Context, string) (ImplResponse, error)
	UpdateKonvey(context.Context, string, Konvey) (ImplResponse, error)
}


// NetgwAPIServicer defines the api actions for the NetgwAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type NetgwAPIServicer interface { 
	DeleteNetGW(context.Context, string) (ImplResponse, error)
	ListNetGWs(context.Context) (ImplResponse, error)
	ReadNetGW(context.Context, string) (ImplResponse, error)
	UpdateNetGW(context.Context, string, NetGw) (ImplResponse, error)
}


// NfsAPIServicer defines the api actions for the NfsAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type NfsAPIServicer interface { 
	DeleteStorageNFS(context.Context, string) (ImplResponse, error)
	ListStorageNFSKFSs(context.Context, string) (ImplResponse, error)
	ListStorageNFSs(context.Context) (ImplResponse, error)
	ReadStorageNFS(context.Context, string) (ImplResponse, error)
	UpdateStorageNFS(context.Context, string, StorageNfs) (ImplResponse, error)
}


// PoolAPIServicer defines the api actions for the PoolAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type PoolAPIServicer interface { 
	CreateTemplate(context.Context, string, Template) (ImplResponse, error)
	DeleteStoragePool(context.Context, string) (ImplResponse, error)
	ListStoragePoolTemplates(context.Context, string) (ImplResponse, error)
	ListStoragePoolVolumes(context.Context, string) (ImplResponse, error)
	ListStoragePools(context.Context) (ImplResponse, error)
	ReadStoragePool(context.Context, string) (ImplResponse, error)
	SetStoragePoolDefaultTemplate(context.Context, string, string) (ImplResponse, error)
	UpdateStoragePool(context.Context, string, StoragePool) (ImplResponse, error)
}


// ProjectAPIServicer defines the api actions for the ProjectAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type ProjectAPIServicer interface { 
	CreateProject(context.Context, Project, int32) (ImplResponse, error)
	CreateProjectDnsRecord(context.Context, string, DnsRecord) (ImplResponse, error)
	CreateProjectRegionKFS(context.Context, string, string, Kfs, string) (ImplResponse, error)
	CreateProjectRegionKGW(context.Context, string, string, Kgw) (ImplResponse, error)
	CreateProjectRegionKonvey(context.Context, string, string, Konvey) (ImplResponse, error)
	CreateProjectRegionVolume(context.Context, string, string, Volume, string, string) (ImplResponse, error)
	CreateProjectZoneInstance(context.Context, string, string, Instance) (ImplResponse, error)
	CreateProjectZoneKCE(context.Context, string, string, Kce, string, string, bool) (ImplResponse, error)
	CreateProjectZoneKonvey(context.Context, string, string, Konvey) (ImplResponse, error)
	DeleteProject(context.Context, string) (ImplResponse, error)
	ListProjectDnsRecords(context.Context, string) (ImplResponse, error)
	ListProjectRegionKFSs(context.Context, string, string, string) (ImplResponse, error)
	ListProjectRegionKGWs(context.Context, string, string) (ImplResponse, error)
	ListProjectRegionKonveys(context.Context, string, string) (ImplResponse, error)
	ListProjectRegionVolumes(context.Context, string, string) (ImplResponse, error)
	ListProjectZoneInstances(context.Context, string, string) (ImplResponse, error)
	ListProjectZoneKCEs(context.Context, string, string) (ImplResponse, error)
	ListProjectZoneKonveys(context.Context, string, string) (ImplResponse, error)
	ListProjects(context.Context, int32) (ImplResponse, error)
	ReadProject(context.Context, string) (ImplResponse, error)
	ReadProjectCost(context.Context, string) (ImplResponse, error)
	ReadProjectUsage(context.Context, string) (ImplResponse, error)
	UpdateProject(context.Context, string, Project) (ImplResponse, error)
}


// RecordAPIServicer defines the api actions for the RecordAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type RecordAPIServicer interface { 
	DeleteDnsRecord(context.Context, string) (ImplResponse, error)
	ReadDnsRecord(context.Context, string) (ImplResponse, error)
	UpdateDnsRecord(context.Context, string, DnsRecord) (ImplResponse, error)
}


// RegionAPIServicer defines the api actions for the RegionAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type RegionAPIServicer interface { 
	CreateNetGW(context.Context, string, NetGw) (ImplResponse, error)
	CreateRegion(context.Context, Region) (ImplResponse, error)
	CreateStorageNFS(context.Context, string, StorageNfs, string) (ImplResponse, error)
	CreateStoragePool(context.Context, string, StoragePool) (ImplResponse, error)
	CreateVNet(context.Context, string, VNet) (ImplResponse, error)
	CreateZone(context.Context, string, Zone) (ImplResponse, error)
	DeleteRegion(context.Context, string) (ImplResponse, error)
	ListRegionNetGWs(context.Context, string) (ImplResponse, error)
	ListRegionStorageNFSs(context.Context, string, string) (ImplResponse, error)
	ListRegionStoragePools(context.Context, string) (ImplResponse, error)
	ListRegionVNets(context.Context, string) (ImplResponse, error)
	ListRegionZones(context.Context, string) (ImplResponse, error)
	ListRegions(context.Context) (ImplResponse, error)
	ReadRegion(context.Context, string) (ImplResponse, error)
	SetRegionDefaultStorageNFS(context.Context, string, string) (ImplResponse, error)
	SetRegionDefaultStoragePool(context.Context, string, string) (ImplResponse, error)
	UpdateRegion(context.Context, string, Region) (ImplResponse, error)
}


// SubnetAPIServicer defines the api actions for the SubnetAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type SubnetAPIServicer interface { 
	CreateAdapter(context.Context, string, Adapter, bool) (ImplResponse, error)
	DeleteSubnet(context.Context, string) (ImplResponse, error)
	ListSubnetAdapters(context.Context, string) (ImplResponse, error)
	ListSubnets(context.Context) (ImplResponse, error)
	ReadSubnet(context.Context, string) (ImplResponse, error)
	UpdateSubnet(context.Context, string, Subnet) (ImplResponse, error)
}


// TemplateAPIServicer defines the api actions for the TemplateAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type TemplateAPIServicer interface { 
	DeleteTemplate(context.Context, string) (ImplResponse, error)
	ListTemplates(context.Context) (ImplResponse, error)
	ReadTemplate(context.Context, string) (ImplResponse, error)
	UpdateTemplate(context.Context, string, Template) (ImplResponse, error)
}


// TokenAPIServicer defines the api actions for the TokenAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type TokenAPIServicer interface { 
	DeleteApiToken(context.Context, string) (ImplResponse, error)
	ListApiTokens(context.Context) (ImplResponse, error)
	ReadApiToken(context.Context, string) (ImplResponse, error)
	UpdateApiToken(context.Context, string, ApiToken) (ImplResponse, error)
}


// UserAPIServicer defines the api actions for the UserAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type UserAPIServicer interface { 
	CreateUser(context.Context, User) (ImplResponse, error)
	DeleteUser(context.Context, string) (ImplResponse, error)
	ListUsers(context.Context) (ImplResponse, error)
	Login(context.Context, UserCredentials) (ImplResponse, error)
	ReadUser(context.Context, string) (ImplResponse, error)
	ResetUserPassword(context.Context, string) (ImplResponse, error)
	SetUserApiToken(context.Context, string, bool, string) (ImplResponse, error)
	SetUserPassword(context.Context, string, Password) (ImplResponse, error)
	UpdateUser(context.Context, string, User) (ImplResponse, error)
}


// VnetAPIServicer defines the api actions for the VnetAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type VnetAPIServicer interface { 
	CreateSubnet(context.Context, string, Subnet) (ImplResponse, error)
	DeleteVNet(context.Context, string) (ImplResponse, error)
	ListVNetSubnets(context.Context, string) (ImplResponse, error)
	ListVNets(context.Context) (ImplResponse, error)
	ReadVNet(context.Context, string) (ImplResponse, error)
	SetVNetDefaultSubnet(context.Context, string, string) (ImplResponse, error)
	UpdateVNet(context.Context, string, VNet) (ImplResponse, error)
}


// VolumeAPIServicer defines the api actions for the VolumeAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type VolumeAPIServicer interface { 
	DeleteVolume(context.Context, string) (ImplResponse, error)
	ListVolumes(context.Context) (ImplResponse, error)
	ReadVolume(context.Context, string) (ImplResponse, error)
	UpdateVolume(context.Context, string, Volume) (ImplResponse, error)
}


// ZoneAPIServicer defines the api actions for the ZoneAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type ZoneAPIServicer interface { 
	CreateHost(context.Context, string, Host) (ImplResponse, error)
	DeleteZone(context.Context, string) (ImplResponse, error)
	ListZoneHosts(context.Context, string) (ImplResponse, error)
	ListZones(context.Context) (ImplResponse, error)
	ReadZone(context.Context, string) (ImplResponse, error)
	UpdateZone(context.Context, string, Zone) (ImplResponse, error)
}
