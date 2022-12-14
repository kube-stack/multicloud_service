package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/tenantnetworks"
)


//extract response info from pager for ListComputeV2ExtensionsTenantnetworks
func ExtractListComputeV2ExtensionsTenantnetworksResponse(response *ListComputeV2ExtensionsTenantnetworksResponse)([]tenantnetworks.Network,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return tenantnetworks.ExtractNetworks(page)
}


// call result's extract function
func ExtractGetComputeV2ExtensionsTenantnetworksResponse(response *GetComputeV2ExtensionsTenantnetworksResponse)(interface{}, error){
    return response.GetResult.Body, response.GetResult.Err
}
