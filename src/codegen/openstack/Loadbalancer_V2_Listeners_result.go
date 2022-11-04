package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/loadbalancer/v2/listeners"
)


//extract response info from pager for ListLoadbalancerV2Listeners
func ExtractListLoadbalancerV2ListenersResponse(response *ListLoadbalancerV2ListenersResponse)([]listeners.Listener,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return listeners.ExtractListeners(page)
}