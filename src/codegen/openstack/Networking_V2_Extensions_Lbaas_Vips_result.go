package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas/vips"
)


//extract response info from pager for ListNetworkingV2ExtensionsLbaasVips
func ExtractListNetworkingV2ExtensionsLbaasVipsResponse(response *ListNetworkingV2ExtensionsLbaasVipsResponse)([]vips.VirtualIP,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return vips.ExtractVIPs(page)
}