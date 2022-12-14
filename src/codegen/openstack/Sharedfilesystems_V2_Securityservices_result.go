package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/sharedfilesystems/v2/securityservices"
)


// call result's extract function
func ExtractCreateSharedfilesystemsV2SecurityservicesResponse(response *CreateSharedfilesystemsV2SecurityservicesResponse)(interface{}, error){
    return response.CreateResult.Body, response.CreateResult.Err
}



// call result's extract function
func ExtractDeleteSharedfilesystemsV2SecurityservicesResponse(response *DeleteSharedfilesystemsV2SecurityservicesResponse)(interface{}, error){
    return response.DeleteResult.Body, response.DeleteResult.Err
}



//extract response info from pager for ListSharedfilesystemsV2Securityservices
func ExtractListSharedfilesystemsV2SecurityservicesResponse(response *ListSharedfilesystemsV2SecurityservicesResponse)([]securityservices.SecurityService,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return securityservices.ExtractSecurityServices(page)
}


// call result's extract function
func ExtractGetSharedfilesystemsV2SecurityservicesResponse(response *GetSharedfilesystemsV2SecurityservicesResponse)(interface{}, error){
    return response.GetResult.Body, response.GetResult.Err
}



// call result's extract function
func ExtractUpdateSharedfilesystemsV2SecurityservicesResponse(response *UpdateSharedfilesystemsV2SecurityservicesResponse)(interface{}, error){
    return response.UpdateResult.Body, response.UpdateResult.Err
}
