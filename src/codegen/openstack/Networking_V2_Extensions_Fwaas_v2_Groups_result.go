package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/fwaas_v2/groups"
)


//extract response info from pager for ListNetworkingV2ExtensionsFwaas_v2Groups
func ExtractListNetworkingV2ExtensionsFwaas_v2GroupsResponse(response *ListNetworkingV2ExtensionsFwaas_v2GroupsResponse)([]groups.Group,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return groups.ExtractGroups(page)
}