package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/clustering/v1/profiles"
)


// call result's extract function
func ExtractCreateClusteringV1ProfilesResponse(response *CreateClusteringV1ProfilesResponse)(interface{}, error){
    return response.CreateResult.Body, response.CreateResult.Err
}



// call result's extract function
func ExtractGetClusteringV1ProfilesResponse(response *GetClusteringV1ProfilesResponse)(interface{}, error){
    return response.GetResult.Body, response.GetResult.Err
}



//extract response info from pager for ListClusteringV1Profiles
func ExtractListClusteringV1ProfilesResponse(response *ListClusteringV1ProfilesResponse)([]profiles.Profile,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return profiles.ExtractProfiles(page)
}


// call result's extract function
func ExtractUpdateClusteringV1ProfilesResponse(response *UpdateClusteringV1ProfilesResponse)(interface{}, error){
    return response.UpdateResult.Body, response.UpdateResult.Err
}



// call result's extract function
func ExtractDeleteClusteringV1ProfilesResponse(response *DeleteClusteringV1ProfilesResponse)(interface{}, error){
    return response.DeleteResult.Body, response.DeleteResult.Err
}



// call result's extract function
func ExtractValidateClusteringV1ProfilesResponse(response *ValidateClusteringV1ProfilesResponse)(interface{}, error){
    return response.ValidateResult.Body, response.ValidateResult.Err
}
