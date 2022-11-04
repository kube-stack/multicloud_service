package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/secgroups"
)


//extract response info from pager for ListComputeV2ExtensionsSecgroups
func ExtractListComputeV2ExtensionsSecgroupsResponse(response *ListComputeV2ExtensionsSecgroupsResponse)([]secgroups.SecurityGroup,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return secgroups.ExtractSecurityGroups(page)
}


//extract response info from pager for ListByServerComputeV2ExtensionsSecgroups
func ExtractListByServerComputeV2ExtensionsSecgroupsResponse(response *ListByServerComputeV2ExtensionsSecgroupsResponse)([]secgroups.SecurityGroup,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return secgroups.ExtractSecurityGroups(page)
}