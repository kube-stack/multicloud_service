package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/networking/v2/ports"
)


//extract response info from pager for ListNetworkingV2Ports
func ExtractListNetworkingV2PortsResponse(response *ListNetworkingV2PortsResponse)([]ports.Port,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return ports.ExtractPorts(page)
}