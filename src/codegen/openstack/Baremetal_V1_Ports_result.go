package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/baremetal/v1/ports"
)


//extract response info from pager for ListBaremetalV1Ports
func ExtractListBaremetalV1PortsResponse(response *ListBaremetalV1PortsResponse)([]ports.Port,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return ports.ExtractPorts(page)
}


//extract response info from pager for ListDetailBaremetalV1Ports
func ExtractListDetailBaremetalV1PortsResponse(response *ListDetailBaremetalV1PortsResponse)([]ports.Port,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return ports.ExtractPorts(page)
}