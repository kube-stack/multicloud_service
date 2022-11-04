package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/pools"
)


//extract response info from pager for ListNetworkingV2ExtensionsLbaas_v2Pools
func ExtractListNetworkingV2ExtensionsLbaas_v2PoolsResponse(response *ListNetworkingV2ExtensionsLbaas_v2PoolsResponse)([]pools.Pool,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return pools.ExtractPools(page)
}


//extract response info from pager for ListMembersNetworkingV2ExtensionsLbaas_v2Pools
func ExtractListMembersNetworkingV2ExtensionsLbaas_v2PoolsResponse(response *ListMembersNetworkingV2ExtensionsLbaas_v2PoolsResponse)([]pools.Member,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return pools.ExtractMembers(page)
}