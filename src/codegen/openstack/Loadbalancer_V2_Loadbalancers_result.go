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