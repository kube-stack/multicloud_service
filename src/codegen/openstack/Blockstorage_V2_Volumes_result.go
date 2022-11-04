package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/blockstorage/v2/volumes"
)


//extract response info from pager for ListBlockstorageV2Volumes
func ExtractListBlockstorageV2VolumesResponse(response *ListBlockstorageV2VolumesResponse)([]volumes.Volume,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return volumes.ExtractVolumes(page)
}