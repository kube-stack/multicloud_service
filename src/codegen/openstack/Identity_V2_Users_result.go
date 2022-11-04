package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/identity/v2/users"
)


//extract response info from pager for ListIdentityV2Users
func ExtractListIdentityV2UsersResponse(response *ListIdentityV2UsersResponse)([]users.User,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return users.ExtractUsers(page)
}


//extract response info from pager for ListRolesIdentityV2Users
func ExtractListRolesIdentityV2UsersResponse(response *ListRolesIdentityV2UsersResponse)([]users.Role,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return users.ExtractRoles(page)
}