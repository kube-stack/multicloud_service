package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/blockstorage/v1/volumes"
)


//extract response info from pager for ListBlockstorageV1Volumes
func ExtractListBlockstorageV1VolumesResponse(response *ListBlockstorageV1VolumesResponse)([]volumes.Volume,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return volumes.ExtractVolumes(page)
}