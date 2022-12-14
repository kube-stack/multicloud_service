package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/blockstorage/apiversions"
)


//extract response info from pager for ListBlockstorageApiversions
func ExtractListBlockstorageApiversionsResponse(response *ListBlockstorageApiversionsResponse)([]apiversions.APIVersion,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return apiversions.ExtractAPIVersions(page)
}