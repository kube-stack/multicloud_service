package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/l7policies"
)


//extract response info from pager for ListNetworkingV2ExtensionsLbaas_v2L7policies
func ExtractListNetworkingV2ExtensionsLbaas_v2L7policiesResponse(response *ListNetworkingV2ExtensionsLbaas_v2L7policiesResponse)([]l7policies.L7Policy,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return l7policies.ExtractL7Policies(page)
}


//extract response info from pager for ListRulesNetworkingV2ExtensionsLbaas_v2L7policies
func ExtractListRulesNetworkingV2ExtensionsLbaas_v2L7policiesResponse(response *ListRulesNetworkingV2ExtensionsLbaas_v2L7policiesResponse)([]l7policies.Rule,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return l7policies.ExtractRules(page)
}