package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/loadbalancer/v2/pools"
)


//extract response info from pager for ListLoadbalancerV2Pools
func ExtractListLoadbalancerV2PoolsResponse(response *ListLoadbalancerV2PoolsResponse)([]pools.Pool,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return pools.ExtractPools(page)
}


//extract response info from pager for ListMembersLoadbalancerV2Pools
func ExtractListMembersLoadbalancerV2PoolsResponse(response *ListMembersLoadbalancerV2PoolsResponse)([]pools.Member,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return pools.ExtractMembers(page)
}