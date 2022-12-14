package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/loadbalancer/v2/loadbalancers"
)


//extract response info from pager for ListLoadbalancerV2Loadbalancers
func ExtractListLoadbalancerV2LoadbalancersResponse(response *ListLoadbalancerV2LoadbalancersResponse)([]loadbalancers.LoadBalancer,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return loadbalancers.ExtractLoadBalancers(page)
}


// call result's extract function
func ExtractCreateLoadbalancerV2LoadbalancersResponse(response *CreateLoadbalancerV2LoadbalancersResponse)(interface{}, error){
    return response.CreateResult.Body, response.CreateResult.Err
}



// call result's extract function
func ExtractGetLoadbalancerV2LoadbalancersResponse(response *GetLoadbalancerV2LoadbalancersResponse)(interface{}, error){
    return response.GetResult.Body, response.GetResult.Err
}



// call result's extract function
func ExtractUpdateLoadbalancerV2LoadbalancersResponse(response *UpdateLoadbalancerV2LoadbalancersResponse)(interface{}, error){
    return response.UpdateResult.Body, response.UpdateResult.Err
}



// call result's extract function
func ExtractDeleteLoadbalancerV2LoadbalancersResponse(response *DeleteLoadbalancerV2LoadbalancersResponse)(interface{}, error){
    return response.DeleteResult.Body, response.DeleteResult.Err
}



// call result's extract function
func ExtractGetStatusesLoadbalancerV2LoadbalancersResponse(response *GetStatusesLoadbalancerV2LoadbalancersResponse)(interface{}, error){
    return response.GetStatusesResult.Body, response.GetStatusesResult.Err
}



// call result's extract function
func ExtractGetStatsLoadbalancerV2LoadbalancersResponse(response *GetStatsLoadbalancerV2LoadbalancersResponse)(interface{}, error){
    return response.StatsResult.Body, response.StatsResult.Err
}



// call result's extract function
func ExtractFailoverLoadbalancerV2LoadbalancersResponse(response *FailoverLoadbalancerV2LoadbalancersResponse)(interface{}, error){
    return response.FailoverResult.Body, response.FailoverResult.Err
}
