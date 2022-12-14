package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/vpnaas/ikepolicies"
)


// call result's extract function
func ExtractCreateNetworkingV2ExtensionsVpnaasIkepoliciesResponse(response *CreateNetworkingV2ExtensionsVpnaasIkepoliciesResponse)(interface{}, error){
    return response.CreateResult.Body, response.CreateResult.Err
}



// call result's extract function
func ExtractGetNetworkingV2ExtensionsVpnaasIkepoliciesResponse(response *GetNetworkingV2ExtensionsVpnaasIkepoliciesResponse)(interface{}, error){
    return response.GetResult.Body, response.GetResult.Err
}



// call result's extract function
func ExtractDeleteNetworkingV2ExtensionsVpnaasIkepoliciesResponse(response *DeleteNetworkingV2ExtensionsVpnaasIkepoliciesResponse)(interface{}, error){
    return response.DeleteResult.Body, response.DeleteResult.Err
}



//extract response info from pager for ListNetworkingV2ExtensionsVpnaasIkepolicies
func ExtractListNetworkingV2ExtensionsVpnaasIkepoliciesResponse(response *ListNetworkingV2ExtensionsVpnaasIkepoliciesResponse)([]ikepolicies.Policy,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return ikepolicies.ExtractPolicies(page)
}


// call result's extract function
func ExtractUpdateNetworkingV2ExtensionsVpnaasIkepoliciesResponse(response *UpdateNetworkingV2ExtensionsVpnaasIkepoliciesResponse)(interface{}, error){
    return response.UpdateResult.Body, response.UpdateResult.Err
}
