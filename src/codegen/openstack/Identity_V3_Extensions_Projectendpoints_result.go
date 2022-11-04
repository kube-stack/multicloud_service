package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/identity/v3/extensions/projectendpoints"
)


//extract response info from pager for ListIdentityV3ExtensionsProjectendpoints
func ExtractListIdentityV3ExtensionsProjectendpointsResponse(response *ListIdentityV3ExtensionsProjectendpointsResponse)([]projectendpoints.Endpoint,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return projectendpoints.ExtractEndpoints(page)
}