package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/networking/v2/apiversions"
)


//extract response info from pager for ListVersionsNetworkingV2Apiversions
func ExtractListVersionsNetworkingV2ApiversionsResponse(response *ListVersionsNetworkingV2ApiversionsResponse)([]apiversions.APIVersion,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return apiversions.ExtractAPIVersions(page)
}


//extract response info from pager for ListVersionResourcesNetworkingV2Apiversions
func ExtractListVersionResourcesNetworkingV2ApiversionsResponse(response *ListVersionResourcesNetworkingV2ApiversionsResponse)([]apiversions.APIVersionResource,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return apiversions.ExtractVersionResources(page)
}