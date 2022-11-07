package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/usage"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the SingleTenantComputeV2ExtensionsUsage
type SingleTenantComputeV2ExtensionsUsageRequest struct{
    TenantID string
    Opts usage.SingleTenantOpts
}

func NewSingleTenantComputeV2ExtensionsUsageRequest()*SingleTenantComputeV2ExtensionsUsageRequest{
    return &SingleTenantComputeV2ExtensionsUsageRequest{}
}

//response struct for the SingleTenantComputeV2ExtensionsUsage
type SingleTenantComputeV2ExtensionsUsageResponse struct{
    Pager pagination.Pager
}

func NewSingleTenantComputeV2ExtensionsUsageResponse(pager pagination.Pager,)*SingleTenantComputeV2ExtensionsUsageResponse {
    return &SingleTenantComputeV2ExtensionsUsageResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) SingleTenantComputeV2ExtensionsUsage(req *SingleTenantComputeV2ExtensionsUsageRequest)(*SingleTenantComputeV2ExtensionsUsageResponse){
    return NewSingleTenantComputeV2ExtensionsUsageResponse(usage.SingleTenant(oc.Client,req.TenantID,req.Opts, ))

}
//request struct for the AllTenantsComputeV2ExtensionsUsage
type AllTenantsComputeV2ExtensionsUsageRequest struct{
    Opts usage.AllTenantsOpts
}

func NewAllTenantsComputeV2ExtensionsUsageRequest()*AllTenantsComputeV2ExtensionsUsageRequest{
    return &AllTenantsComputeV2ExtensionsUsageRequest{}
}

//response struct for the AllTenantsComputeV2ExtensionsUsage
type AllTenantsComputeV2ExtensionsUsageResponse struct{
    Pager pagination.Pager
}

func NewAllTenantsComputeV2ExtensionsUsageResponse(pager pagination.Pager,)*AllTenantsComputeV2ExtensionsUsageResponse {
    return &AllTenantsComputeV2ExtensionsUsageResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) AllTenantsComputeV2ExtensionsUsage(req *AllTenantsComputeV2ExtensionsUsageRequest)(*AllTenantsComputeV2ExtensionsUsageResponse){
    return NewAllTenantsComputeV2ExtensionsUsageResponse(usage.AllTenants(oc.Client,req.Opts, ))

}