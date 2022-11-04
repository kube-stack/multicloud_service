package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/identity/v3/services"
)


//extract response info from pager for ListIdentityV3Services
func ExtractListIdentityV3ServicesResponse(response *ListIdentityV3ServicesResponse)([]services.Service,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return services.ExtractServices(page)
}