package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/clustering/v1/policies"
)


//extract response info from pager for ListClusteringV1Policies
func ExtractListClusteringV1PoliciesResponse(response *ListClusteringV1PoliciesResponse)([]policies.Policy,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return policies.ExtractPolicies(page)
}


// call result's extract function
func ExtractCreateClusteringV1PoliciesResponse(response *CreateClusteringV1PoliciesResponse)(interface{}, error){
    return response.CreateResult.Body, response.CreateResult.Err
}



// call result's extract function
func ExtractDeleteClusteringV1PoliciesResponse(response *DeleteClusteringV1PoliciesResponse)(interface{}, error){
    return response.DeleteResult.Body, response.DeleteResult.Err
}



// call result's extract function
func ExtractUpdateClusteringV1PoliciesResponse(response *UpdateClusteringV1PoliciesResponse)(interface{}, error){
    return response.UpdateResult.Body, response.UpdateResult.Err
}



// call result's extract function
func ExtractValidateClusteringV1PoliciesResponse(response *ValidateClusteringV1PoliciesResponse)(interface{}, error){
    return response.ValidateResult.Body, response.ValidateResult.Err
}



// call result's extract function
func ExtractGetClusteringV1PoliciesResponse(response *GetClusteringV1PoliciesResponse)(interface{}, error){
    return response.GetResult.Body, response.GetResult.Err
}
