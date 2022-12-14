package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/availabilityzones"
)


//extract response info from pager for ListComputeV2ExtensionsAvailabilityzones
func ExtractListComputeV2ExtensionsAvailabilityzonesResponse(response *ListComputeV2ExtensionsAvailabilityzonesResponse)([]availabilityzones.AvailabilityZone,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return availabilityzones.ExtractAvailabilityZones(page)
}


//extract response info from pager for ListDetailComputeV2ExtensionsAvailabilityzones
func ExtractListDetailComputeV2ExtensionsAvailabilityzonesResponse(response *ListDetailComputeV2ExtensionsAvailabilityzonesResponse)([]availabilityzones.AvailabilityZone,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return availabilityzones.ExtractAvailabilityZones(page)
}