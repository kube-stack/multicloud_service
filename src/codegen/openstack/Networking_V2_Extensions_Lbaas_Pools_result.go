package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas/pools"
)


//extract response info from pager for ListNetworkingV2ExtensionsLbaasPools
func ExtractListNetworkingV2ExtensionsLbaasPoolsResponse(response *ListNetworkingV2ExtensionsLbaasPoolsResponse)([]pools.Pool,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return pools.ExtractPools(page)
}


// call result's extract function
func ExtractCreateNetworkingV2ExtensionsLbaasPoolsResponse(response *CreateNetworkingV2ExtensionsLbaasPoolsResponse)(interface{}, error){
    return response.CreateResult.Body, response.CreateResult.Err
}



// call result's extract function
func ExtractGetNetworkingV2ExtensionsLbaasPoolsResponse(response *GetNetworkingV2ExtensionsLbaasPoolsResponse)(interface{}, error){
    return response.GetResult.Body, response.GetResult.Err
}



// call result's extract function
func ExtractUpdateNetworkingV2ExtensionsLbaasPoolsResponse(response *UpdateNetworkingV2ExtensionsLbaasPoolsResponse)(interface{}, error){
    return response.UpdateResult.Body, response.UpdateResult.Err
}



// call result's extract function
func ExtractDeleteNetworkingV2ExtensionsLbaasPoolsResponse(response *DeleteNetworkingV2ExtensionsLbaasPoolsResponse)(interface{}, error){
    return response.DeleteResult.Body, response.DeleteResult.Err
}



// call result's extract function
func ExtractAssociateMonitorNetworkingV2ExtensionsLbaasPoolsResponse(response *AssociateMonitorNetworkingV2ExtensionsLbaasPoolsResponse)(interface{}, error){
    return response.AssociateResult.Body, response.AssociateResult.Err
}



// call result's extract function
func ExtractDisassociateMonitorNetworkingV2ExtensionsLbaasPoolsResponse(response *DisassociateMonitorNetworkingV2ExtensionsLbaasPoolsResponse)(interface{}, error){
    return response.AssociateResult.Body, response.AssociateResult.Err
}
