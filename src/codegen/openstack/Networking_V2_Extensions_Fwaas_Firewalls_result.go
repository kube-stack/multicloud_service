package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/fwaas/firewalls"
)


//extract response info from pager for ListNetworkingV2ExtensionsFwaasFirewalls
func ExtractListNetworkingV2ExtensionsFwaasFirewallsResponse(response *ListNetworkingV2ExtensionsFwaasFirewallsResponse)([]firewalls.Firewall,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return firewalls.ExtractFirewalls(page)
}