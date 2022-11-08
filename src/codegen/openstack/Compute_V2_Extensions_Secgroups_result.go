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


// call result's extract function
func ExtractCreateComputeV2ExtensionsSecgroupsResponse(response *CreateComputeV2ExtensionsSecgroupsResponse)(interface{}, error){
    return response.CreateResult.Body, response.CreateResult.Err
}



// call result's extract function
func ExtractUpdateComputeV2ExtensionsSecgroupsResponse(response *UpdateComputeV2ExtensionsSecgroupsResponse)(interface{}, error){
    return response.UpdateResult.Body, response.UpdateResult.Err
}



// call result's extract function
func ExtractGetComputeV2ExtensionsSecgroupsResponse(response *GetComputeV2ExtensionsSecgroupsResponse)(interface{}, error){
    return response.GetResult.Body, response.GetResult.Err
}



// call result's extract function
func ExtractDeleteComputeV2ExtensionsSecgroupsResponse(response *DeleteComputeV2ExtensionsSecgroupsResponse)(interface{}, error){
    return response.DeleteResult.Body, response.DeleteResult.Err
}



// call result's extract function
func ExtractCreateRuleComputeV2ExtensionsSecgroupsResponse(response *CreateRuleComputeV2ExtensionsSecgroupsResponse)(interface{}, error){
    return response.CreateRuleResult.Body, response.CreateRuleResult.Err
}



// call result's extract function
func ExtractDeleteRuleComputeV2ExtensionsSecgroupsResponse(response *DeleteRuleComputeV2ExtensionsSecgroupsResponse)(interface{}, error){
    return response.DeleteRuleResult.Body, response.DeleteRuleResult.Err
}



// call result's extract function
func ExtractAddServerComputeV2ExtensionsSecgroupsResponse(response *AddServerComputeV2ExtensionsSecgroupsResponse)(interface{}, error){
    return response.AddServerResult.Body, response.AddServerResult.Err
}



// call result's extract function
func ExtractRemoveServerComputeV2ExtensionsSecgroupsResponse(response *RemoveServerComputeV2ExtensionsSecgroupsResponse)(interface{}, error){
    return response.RemoveServerResult.Body, response.RemoveServerResult.Err
}