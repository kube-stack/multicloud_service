package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/sharedfilesystems/v2/sharetypes"
)


//extract response info from pager for ListSharedfilesystemsV2Sharetypes
func ExtractListSharedfilesystemsV2SharetypesResponse(response *ListSharedfilesystemsV2SharetypesResponse)([]sharetypes.ShareType,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return sharetypes.ExtractShareTypes(page)
}