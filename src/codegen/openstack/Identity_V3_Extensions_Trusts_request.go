package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/identity/v3/extensions/trusts"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the CreateIdentityV3ExtensionsTrusts
type CreateIdentityV3ExtensionsTrustsRequest struct{
    Opts trusts.CreateOpts
}

func NewCreateIdentityV3ExtensionsTrustsRequest()*CreateIdentityV3ExtensionsTrustsRequest{
    return &CreateIdentityV3ExtensionsTrustsRequest{}
}

//response struct for the CreateIdentityV3ExtensionsTrusts
type CreateIdentityV3ExtensionsTrustsResponse struct{
    CreateResult trusts.CreateResult
}

func NewCreateIdentityV3ExtensionsTrustsResponse(createResult trusts.CreateResult,)*CreateIdentityV3ExtensionsTrustsResponse {
    return &CreateIdentityV3ExtensionsTrustsResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateIdentityV3ExtensionsTrusts(req *CreateIdentityV3ExtensionsTrustsRequest)(*CreateIdentityV3ExtensionsTrustsResponse){
    return NewCreateIdentityV3ExtensionsTrustsResponse(trusts.Create(oc.Client,req.Opts, ))

}
//request struct for the DeleteIdentityV3ExtensionsTrusts
type DeleteIdentityV3ExtensionsTrustsRequest struct{
    TrustID string
}

func NewDeleteIdentityV3ExtensionsTrustsRequest()*DeleteIdentityV3ExtensionsTrustsRequest{
    return &DeleteIdentityV3ExtensionsTrustsRequest{}
}

//response struct for the DeleteIdentityV3ExtensionsTrusts
type DeleteIdentityV3ExtensionsTrustsResponse struct{
    DeleteResult trusts.DeleteResult
}

func NewDeleteIdentityV3ExtensionsTrustsResponse(deleteResult trusts.DeleteResult,)*DeleteIdentityV3ExtensionsTrustsResponse {
    return &DeleteIdentityV3ExtensionsTrustsResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteIdentityV3ExtensionsTrusts(req *DeleteIdentityV3ExtensionsTrustsRequest)(*DeleteIdentityV3ExtensionsTrustsResponse){
    return NewDeleteIdentityV3ExtensionsTrustsResponse(trusts.Delete(oc.Client,req.TrustID, ))

}
//request struct for the ListIdentityV3ExtensionsTrusts
type ListIdentityV3ExtensionsTrustsRequest struct{
    Opts trusts.ListOpts
}

func NewListIdentityV3ExtensionsTrustsRequest()*ListIdentityV3ExtensionsTrustsRequest{
    return &ListIdentityV3ExtensionsTrustsRequest{}
}

//response struct for the ListIdentityV3ExtensionsTrusts
type ListIdentityV3ExtensionsTrustsResponse struct{
    Pager pagination.Pager
}

func NewListIdentityV3ExtensionsTrustsResponse(pager pagination.Pager,)*ListIdentityV3ExtensionsTrustsResponse {
    return &ListIdentityV3ExtensionsTrustsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListIdentityV3ExtensionsTrusts(req *ListIdentityV3ExtensionsTrustsRequest)(*ListIdentityV3ExtensionsTrustsResponse){
    return NewListIdentityV3ExtensionsTrustsResponse(trusts.List(oc.Client,req.Opts, ))

}
//request struct for the GetIdentityV3ExtensionsTrusts
type GetIdentityV3ExtensionsTrustsRequest struct{
    Id string
}

func NewGetIdentityV3ExtensionsTrustsRequest()*GetIdentityV3ExtensionsTrustsRequest{
    return &GetIdentityV3ExtensionsTrustsRequest{}
}

//response struct for the GetIdentityV3ExtensionsTrusts
type GetIdentityV3ExtensionsTrustsResponse struct{
    GetResult trusts.GetResult
}

func NewGetIdentityV3ExtensionsTrustsResponse(getResult trusts.GetResult,)*GetIdentityV3ExtensionsTrustsResponse {
    return &GetIdentityV3ExtensionsTrustsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetIdentityV3ExtensionsTrusts(req *GetIdentityV3ExtensionsTrustsRequest)(*GetIdentityV3ExtensionsTrustsResponse){
    return NewGetIdentityV3ExtensionsTrustsResponse(trusts.Get(oc.Client,req.Id, ))

}
//request struct for the ListRolesIdentityV3ExtensionsTrusts
type ListRolesIdentityV3ExtensionsTrustsRequest struct{
    Id string
}

func NewListRolesIdentityV3ExtensionsTrustsRequest()*ListRolesIdentityV3ExtensionsTrustsRequest{
    return &ListRolesIdentityV3ExtensionsTrustsRequest{}
}

//response struct for the ListRolesIdentityV3ExtensionsTrusts
type ListRolesIdentityV3ExtensionsTrustsResponse struct{
    Pager pagination.Pager
}

func NewListRolesIdentityV3ExtensionsTrustsResponse(pager pagination.Pager,)*ListRolesIdentityV3ExtensionsTrustsResponse {
    return &ListRolesIdentityV3ExtensionsTrustsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListRolesIdentityV3ExtensionsTrusts(req *ListRolesIdentityV3ExtensionsTrustsRequest)(*ListRolesIdentityV3ExtensionsTrustsResponse){
    return NewListRolesIdentityV3ExtensionsTrustsResponse(trusts.ListRoles(oc.Client,req.Id, ))

}
//request struct for the GetRoleIdentityV3ExtensionsTrusts
type GetRoleIdentityV3ExtensionsTrustsRequest struct{
    Id string
    RoleID string
}

func NewGetRoleIdentityV3ExtensionsTrustsRequest()*GetRoleIdentityV3ExtensionsTrustsRequest{
    return &GetRoleIdentityV3ExtensionsTrustsRequest{}
}

//response struct for the GetRoleIdentityV3ExtensionsTrusts
type GetRoleIdentityV3ExtensionsTrustsResponse struct{
    GetRoleResult trusts.GetRoleResult
}

func NewGetRoleIdentityV3ExtensionsTrustsResponse(getRoleResult trusts.GetRoleResult,)*GetRoleIdentityV3ExtensionsTrustsResponse {
    return &GetRoleIdentityV3ExtensionsTrustsResponse{
            GetRoleResult:getRoleResult,
    }
}

// action function
func (oc *OpenstackClient) GetRoleIdentityV3ExtensionsTrusts(req *GetRoleIdentityV3ExtensionsTrustsRequest)(*GetRoleIdentityV3ExtensionsTrustsResponse){
    return NewGetRoleIdentityV3ExtensionsTrustsResponse(trusts.GetRole(oc.Client,req.Id,req.RoleID, ))

}
//request struct for the CheckRoleIdentityV3ExtensionsTrusts
type CheckRoleIdentityV3ExtensionsTrustsRequest struct{
    Id string
    RoleID string
}

func NewCheckRoleIdentityV3ExtensionsTrustsRequest()*CheckRoleIdentityV3ExtensionsTrustsRequest{
    return &CheckRoleIdentityV3ExtensionsTrustsRequest{}
}

//response struct for the CheckRoleIdentityV3ExtensionsTrusts
type CheckRoleIdentityV3ExtensionsTrustsResponse struct{
    CheckRoleResult trusts.CheckRoleResult
}

func NewCheckRoleIdentityV3ExtensionsTrustsResponse(checkRoleResult trusts.CheckRoleResult,)*CheckRoleIdentityV3ExtensionsTrustsResponse {
    return &CheckRoleIdentityV3ExtensionsTrustsResponse{
            CheckRoleResult:checkRoleResult,
    }
}

// action function
func (oc *OpenstackClient) CheckRoleIdentityV3ExtensionsTrusts(req *CheckRoleIdentityV3ExtensionsTrustsRequest)(*CheckRoleIdentityV3ExtensionsTrustsResponse){
    return NewCheckRoleIdentityV3ExtensionsTrustsResponse(trusts.CheckRole(oc.Client,req.Id,req.RoleID, ))

}