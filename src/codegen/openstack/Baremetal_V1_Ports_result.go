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


// call result's extract function
func ExtractGetBaremetalV1PortsResponse(response *GetBaremetalV1PortsResponse)(interface{}, error){
    return response.GetResult.Body, response.GetResult.Err
}



// call result's extract function
func ExtractCreateBaremetalV1PortsResponse(response *CreateBaremetalV1PortsResponse)(interface{}, error){
    return response.CreateResult.Body, response.CreateResult.Err
}



// call result's extract function
func ExtractUpdateBaremetalV1PortsResponse(response *UpdateBaremetalV1PortsResponse)(interface{}, error){
    return response.UpdateResult.Body, response.UpdateResult.Err
}



// call result's extract function
func ExtractDeleteBaremetalV1PortsResponse(response *DeleteBaremetalV1PortsResponse)(interface{}, error){
    return response.DeleteResult.Body, response.DeleteResult.Err
}
