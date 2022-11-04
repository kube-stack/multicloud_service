package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/fwaas/rules"
)


//extract response info from pager for ListNetworkingV2ExtensionsFwaasRules
func ExtractListNetworkingV2ExtensionsFwaasRulesResponse(response *ListNetworkingV2ExtensionsFwaasRulesResponse)([]rules.Rule,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return rules.ExtractRules(page)
}