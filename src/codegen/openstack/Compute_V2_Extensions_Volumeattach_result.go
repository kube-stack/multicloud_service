package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/volumeattach"
)


//extract response info from pager for ListComputeV2ExtensionsVolumeattach
func ExtractListComputeV2ExtensionsVolumeattachResponse(response *ListComputeV2ExtensionsVolumeattachResponse)([]volumeattach.VolumeAttachment,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return volumeattach.ExtractVolumeAttachments(page)
}