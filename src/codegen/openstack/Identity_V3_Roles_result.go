package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/identity/v3/roles"
)


//extract response info from pager for ListIdentityV3Roles
func ExtractListIdentityV3RolesResponse(response *ListIdentityV3RolesResponse)([]roles.Role,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return roles.ExtractRoles(page)
}


//extract response info from pager for ListAssignmentsIdentityV3Roles
func ExtractListAssignmentsIdentityV3RolesResponse(response *ListAssignmentsIdentityV3RolesResponse)([]roles.RoleAssignment,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return roles.ExtractRoleAssignments(page)
}


//extract response info from pager for ListAssignmentsOnResourceIdentityV3Roles
func ExtractListAssignmentsOnResourceIdentityV3RolesResponse(response *ListAssignmentsOnResourceIdentityV3RolesResponse)([]roles.Role,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return roles.ExtractRoles(page)
}