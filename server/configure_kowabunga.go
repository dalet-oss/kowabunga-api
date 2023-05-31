// This file is safe to edit. Once it exists it will not be overwritten

package server

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	swaggerapi "github.com/dalet-oss/kowabunga-api/server/api"
	"github.com/dalet-oss/kowabunga-api/server/api/host"
	"github.com/dalet-oss/kowabunga-api/server/api/instance"
	"github.com/dalet-oss/kowabunga-api/server/api/netgw"
	"github.com/dalet-oss/kowabunga-api/server/api/pool"
	"github.com/dalet-oss/kowabunga-api/server/api/project"
	"github.com/dalet-oss/kowabunga-api/server/api/region"
	"github.com/dalet-oss/kowabunga-api/server/api/vnet"
	"github.com/dalet-oss/kowabunga-api/server/api/zone"
)

//go:generate swagger generate server --target ../../kowabunga-api --name Kowabunga --spec ../swagger.generated.yml --api-package api --server-package server --principal interface{} --exclude-main

func configureFlags(api *swaggerapi.KowabungaAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *swaggerapi.KowabungaAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()
	api.TxtProducer = runtime.TextProducer()

	// Applies when the "x-token" header is set
	if api.KeyAuth == nil {
		api.KeyAuth = func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (key) x-token from header param [x-token] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.ZoneCreateHostHandler == nil {
		api.ZoneCreateHostHandler = zone.CreateHostHandlerFunc(func(params zone.CreateHostParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation zone.CreateHost has not yet been implemented")
		})
	}
	if api.ZoneCreateNetGWHandler == nil {
		api.ZoneCreateNetGWHandler = zone.CreateNetGWHandlerFunc(func(params zone.CreateNetGWParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation zone.CreateNetGW has not yet been implemented")
		})
	}
	if api.ZoneCreatePoolHandler == nil {
		api.ZoneCreatePoolHandler = zone.CreatePoolHandlerFunc(func(params zone.CreatePoolParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation zone.CreatePool has not yet been implemented")
		})
	}
	if api.ProjectCreateProjectHandler == nil {
		api.ProjectCreateProjectHandler = project.CreateProjectHandlerFunc(func(params project.CreateProjectParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation project.CreateProject has not yet been implemented")
		})
	}
	if api.RegionCreateRegionHandler == nil {
		api.RegionCreateRegionHandler = region.CreateRegionHandlerFunc(func(params region.CreateRegionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation region.CreateRegion has not yet been implemented")
		})
	}
	if api.ZoneCreateVNetHandler == nil {
		api.ZoneCreateVNetHandler = zone.CreateVNetHandlerFunc(func(params zone.CreateVNetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation zone.CreateVNet has not yet been implemented")
		})
	}
	if api.RegionCreateZoneHandler == nil {
		api.RegionCreateZoneHandler = region.CreateZoneHandlerFunc(func(params region.CreateZoneParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation region.CreateZone has not yet been implemented")
		})
	}
	if api.HostDeleteHostHandler == nil {
		api.HostDeleteHostHandler = host.DeleteHostHandlerFunc(func(params host.DeleteHostParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation host.DeleteHost has not yet been implemented")
		})
	}
	if api.NetgwDeleteNetGWHandler == nil {
		api.NetgwDeleteNetGWHandler = netgw.DeleteNetGWHandlerFunc(func(params netgw.DeleteNetGWParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation netgw.DeleteNetGW has not yet been implemented")
		})
	}
	if api.PoolDeletePoolHandler == nil {
		api.PoolDeletePoolHandler = pool.DeletePoolHandlerFunc(func(params pool.DeletePoolParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pool.DeletePool has not yet been implemented")
		})
	}
	if api.ProjectDeleteProjectHandler == nil {
		api.ProjectDeleteProjectHandler = project.DeleteProjectHandlerFunc(func(params project.DeleteProjectParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation project.DeleteProject has not yet been implemented")
		})
	}
	if api.RegionDeleteRegionHandler == nil {
		api.RegionDeleteRegionHandler = region.DeleteRegionHandlerFunc(func(params region.DeleteRegionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation region.DeleteRegion has not yet been implemented")
		})
	}
	if api.VnetDeleteVNetHandler == nil {
		api.VnetDeleteVNetHandler = vnet.DeleteVNetHandlerFunc(func(params vnet.DeleteVNetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation vnet.DeleteVNet has not yet been implemented")
		})
	}
	if api.ZoneDeleteZoneHandler == nil {
		api.ZoneDeleteZoneHandler = zone.DeleteZoneHandlerFunc(func(params zone.DeleteZoneParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation zone.DeleteZone has not yet been implemented")
		})
	}
	if api.HostGetAllHostsHandler == nil {
		api.HostGetAllHostsHandler = host.GetAllHostsHandlerFunc(func(params host.GetAllHostsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation host.GetAllHosts has not yet been implemented")
		})
	}
	if api.InstanceGetAllInstancesHandler == nil {
		api.InstanceGetAllInstancesHandler = instance.GetAllInstancesHandlerFunc(func(params instance.GetAllInstancesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation instance.GetAllInstances has not yet been implemented")
		})
	}
	if api.NetgwGetAllNetGWsHandler == nil {
		api.NetgwGetAllNetGWsHandler = netgw.GetAllNetGWsHandlerFunc(func(params netgw.GetAllNetGWsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation netgw.GetAllNetGWs has not yet been implemented")
		})
	}
	if api.PoolGetAllPoolsHandler == nil {
		api.PoolGetAllPoolsHandler = pool.GetAllPoolsHandlerFunc(func(params pool.GetAllPoolsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pool.GetAllPools has not yet been implemented")
		})
	}
	if api.ProjectGetAllProjectsHandler == nil {
		api.ProjectGetAllProjectsHandler = project.GetAllProjectsHandlerFunc(func(params project.GetAllProjectsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation project.GetAllProjects has not yet been implemented")
		})
	}
	if api.RegionGetAllRegionsHandler == nil {
		api.RegionGetAllRegionsHandler = region.GetAllRegionsHandlerFunc(func(params region.GetAllRegionsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation region.GetAllRegions has not yet been implemented")
		})
	}
	if api.VnetGetAllVNetsHandler == nil {
		api.VnetGetAllVNetsHandler = vnet.GetAllVNetsHandlerFunc(func(params vnet.GetAllVNetsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation vnet.GetAllVNets has not yet been implemented")
		})
	}
	if api.ZoneGetAllZonesHandler == nil {
		api.ZoneGetAllZonesHandler = zone.GetAllZonesHandlerFunc(func(params zone.GetAllZonesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation zone.GetAllZones has not yet been implemented")
		})
	}
	if api.HostGetHostHandler == nil {
		api.HostGetHostHandler = host.GetHostHandlerFunc(func(params host.GetHostParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation host.GetHost has not yet been implemented")
		})
	}
	if api.HostGetHostCapsHandler == nil {
		api.HostGetHostCapsHandler = host.GetHostCapsHandlerFunc(func(params host.GetHostCapsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation host.GetHostCaps has not yet been implemented")
		})
	}
	if api.HostGetHostInstancesHandler == nil {
		api.HostGetHostInstancesHandler = host.GetHostInstancesHandlerFunc(func(params host.GetHostInstancesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation host.GetHostInstances has not yet been implemented")
		})
	}
	if api.InstanceGetInstanceHandler == nil {
		api.InstanceGetInstanceHandler = instance.GetInstanceHandlerFunc(func(params instance.GetInstanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation instance.GetInstance has not yet been implemented")
		})
	}
	if api.InstanceGetInstanceStateHandler == nil {
		api.InstanceGetInstanceStateHandler = instance.GetInstanceStateHandlerFunc(func(params instance.GetInstanceStateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation instance.GetInstanceState has not yet been implemented")
		})
	}
	if api.NetgwGetNetGWHandler == nil {
		api.NetgwGetNetGWHandler = netgw.GetNetGWHandlerFunc(func(params netgw.GetNetGWParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation netgw.GetNetGW has not yet been implemented")
		})
	}
	if api.PoolGetPoolHandler == nil {
		api.PoolGetPoolHandler = pool.GetPoolHandlerFunc(func(params pool.GetPoolParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pool.GetPool has not yet been implemented")
		})
	}
	if api.ProjectGetProjectHandler == nil {
		api.ProjectGetProjectHandler = project.GetProjectHandlerFunc(func(params project.GetProjectParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation project.GetProject has not yet been implemented")
		})
	}
	if api.ProjectGetProjectQuotasHandler == nil {
		api.ProjectGetProjectQuotasHandler = project.GetProjectQuotasHandlerFunc(func(params project.GetProjectQuotasParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation project.GetProjectQuotas has not yet been implemented")
		})
	}
	if api.ProjectGetProjectUsageHandler == nil {
		api.ProjectGetProjectUsageHandler = project.GetProjectUsageHandlerFunc(func(params project.GetProjectUsageParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation project.GetProjectUsage has not yet been implemented")
		})
	}
	if api.RegionGetRegionHandler == nil {
		api.RegionGetRegionHandler = region.GetRegionHandlerFunc(func(params region.GetRegionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation region.GetRegion has not yet been implemented")
		})
	}
	if api.RegionGetRegionZonesHandler == nil {
		api.RegionGetRegionZonesHandler = region.GetRegionZonesHandlerFunc(func(params region.GetRegionZonesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation region.GetRegionZones has not yet been implemented")
		})
	}
	if api.VnetGetVNetHandler == nil {
		api.VnetGetVNetHandler = vnet.GetVNetHandlerFunc(func(params vnet.GetVNetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation vnet.GetVNet has not yet been implemented")
		})
	}
	if api.ZoneGetZoneHandler == nil {
		api.ZoneGetZoneHandler = zone.GetZoneHandlerFunc(func(params zone.GetZoneParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation zone.GetZone has not yet been implemented")
		})
	}
	if api.ZoneGetZoneHostsHandler == nil {
		api.ZoneGetZoneHostsHandler = zone.GetZoneHostsHandlerFunc(func(params zone.GetZoneHostsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation zone.GetZoneHosts has not yet been implemented")
		})
	}
	if api.ZoneGetZoneNetGWsHandler == nil {
		api.ZoneGetZoneNetGWsHandler = zone.GetZoneNetGWsHandlerFunc(func(params zone.GetZoneNetGWsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation zone.GetZoneNetGWs has not yet been implemented")
		})
	}
	if api.ZoneGetZonePoolsHandler == nil {
		api.ZoneGetZonePoolsHandler = zone.GetZonePoolsHandlerFunc(func(params zone.GetZonePoolsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation zone.GetZonePools has not yet been implemented")
		})
	}
	if api.ZoneGetZoneVNetsHandler == nil {
		api.ZoneGetZoneVNetsHandler = zone.GetZoneVNetsHandlerFunc(func(params zone.GetZoneVNetsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation zone.GetZoneVNets has not yet been implemented")
		})
	}
	if api.HealthzHandler == nil {
		api.HealthzHandler = swaggerapi.HealthzHandlerFunc(func(params swaggerapi.HealthzParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation api.Healthz has not yet been implemented")
		})
	}
	if api.InstanceRebootInstanceHandler == nil {
		api.InstanceRebootInstanceHandler = instance.RebootInstanceHandlerFunc(func(params instance.RebootInstanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation instance.RebootInstance has not yet been implemented")
		})
	}
	if api.InstanceResetInstanceHandler == nil {
		api.InstanceResetInstanceHandler = instance.ResetInstanceHandlerFunc(func(params instance.ResetInstanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation instance.ResetInstance has not yet been implemented")
		})
	}
	if api.ProjectResetProjectQuotasHandler == nil {
		api.ProjectResetProjectQuotasHandler = project.ResetProjectQuotasHandlerFunc(func(params project.ResetProjectQuotasParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation project.ResetProjectQuotas has not yet been implemented")
		})
	}
	if api.InstanceResumeInstanceHandler == nil {
		api.InstanceResumeInstanceHandler = instance.ResumeInstanceHandlerFunc(func(params instance.ResumeInstanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation instance.ResumeInstance has not yet been implemented")
		})
	}
	if api.InstanceShutdownInstanceHandler == nil {
		api.InstanceShutdownInstanceHandler = instance.ShutdownInstanceHandlerFunc(func(params instance.ShutdownInstanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation instance.ShutdownInstance has not yet been implemented")
		})
	}
	if api.InstanceStartInstanceHandler == nil {
		api.InstanceStartInstanceHandler = instance.StartInstanceHandlerFunc(func(params instance.StartInstanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation instance.StartInstance has not yet been implemented")
		})
	}
	if api.InstanceStopInstanceHandler == nil {
		api.InstanceStopInstanceHandler = instance.StopInstanceHandlerFunc(func(params instance.StopInstanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation instance.StopInstance has not yet been implemented")
		})
	}
	if api.InstanceSuspendInstanceHandler == nil {
		api.InstanceSuspendInstanceHandler = instance.SuspendInstanceHandlerFunc(func(params instance.SuspendInstanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation instance.SuspendInstance has not yet been implemented")
		})
	}
	if api.HostUpdateHostHandler == nil {
		api.HostUpdateHostHandler = host.UpdateHostHandlerFunc(func(params host.UpdateHostParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation host.UpdateHost has not yet been implemented")
		})
	}
	if api.NetgwUpdateNetGWHandler == nil {
		api.NetgwUpdateNetGWHandler = netgw.UpdateNetGWHandlerFunc(func(params netgw.UpdateNetGWParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation netgw.UpdateNetGW has not yet been implemented")
		})
	}
	if api.PoolUpdatePoolHandler == nil {
		api.PoolUpdatePoolHandler = pool.UpdatePoolHandlerFunc(func(params pool.UpdatePoolParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pool.UpdatePool has not yet been implemented")
		})
	}
	if api.ProjectUpdateProjectHandler == nil {
		api.ProjectUpdateProjectHandler = project.UpdateProjectHandlerFunc(func(params project.UpdateProjectParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation project.UpdateProject has not yet been implemented")
		})
	}
	if api.ProjectUpdateProjectQuotasHandler == nil {
		api.ProjectUpdateProjectQuotasHandler = project.UpdateProjectQuotasHandlerFunc(func(params project.UpdateProjectQuotasParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation project.UpdateProjectQuotas has not yet been implemented")
		})
	}
	if api.RegionUpdateRegionHandler == nil {
		api.RegionUpdateRegionHandler = region.UpdateRegionHandlerFunc(func(params region.UpdateRegionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation region.UpdateRegion has not yet been implemented")
		})
	}
	if api.VnetUpdateVNetHandler == nil {
		api.VnetUpdateVNetHandler = vnet.UpdateVNetHandlerFunc(func(params vnet.UpdateVNetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation vnet.UpdateVNet has not yet been implemented")
		})
	}
	if api.ZoneUpdateZoneHandler == nil {
		api.ZoneUpdateZoneHandler = zone.UpdateZoneHandlerFunc(func(params zone.UpdateZoneParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation zone.UpdateZone has not yet been implemented")
		})
	}
	if api.ZoneUpdateZoneDefaultVNetHandler == nil {
		api.ZoneUpdateZoneDefaultVNetHandler = zone.UpdateZoneDefaultVNetHandlerFunc(func(params zone.UpdateZoneDefaultVNetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation zone.UpdateZoneDefaultVNet has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
