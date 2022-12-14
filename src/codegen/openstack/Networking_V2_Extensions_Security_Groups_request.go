package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/groups"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListNetworkingV2ExtensionsSecurityGroups
type ListNetworkingV2ExtensionsSecurityGroupsRequest struct{
    Opts groups.ListOpts
}

func NewListNetworkingV2ExtensionsSecurityGroupsRequest()*ListNetworkingV2ExtensionsSecurityGroupsRequest{
    return &ListNetworkingV2ExtensionsSecurityGroupsRequest{}
}

//response struct for the ListNetworkingV2ExtensionsSecurityGroups
type ListNetworkingV2ExtensionsSecurityGroupsResponse struct{
    Pager pagination.Pager
}

func NewListNetworkingV2ExtensionsSecurityGroupsResponse(pager pagination.Pager,)*ListNetworkingV2ExtensionsSecurityGroupsResponse {
    return &ListNetworkingV2ExtensionsSecurityGroupsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListNetworkingV2ExtensionsSecurityGroups(req *ListNetworkingV2ExtensionsSecurityGroupsRequest)(*ListNetworkingV2ExtensionsSecurityGroupsResponse){
    return NewListNetworkingV2ExtensionsSecurityGroupsResponse(groups.List(oc.Client,req.Opts, ))

}
//request struct for the CreateNetworkingV2ExtensionsSecurityGroups
type CreateNetworkingV2ExtensionsSecurityGroupsRequest struct{
    Opts groups.CreateOpts
}

func NewCreateNetworkingV2ExtensionsSecurityGroupsRequest()*CreateNetworkingV2ExtensionsSecurityGroupsRequest{
    return &CreateNetworkingV2ExtensionsSecurityGroupsRequest{}
}

//response struct for the CreateNetworkingV2ExtensionsSecurityGroups
type CreateNetworkingV2ExtensionsSecurityGroupsResponse struct{
    CreateResult groups.CreateResult
}

func NewCreateNetworkingV2ExtensionsSecurityGroupsResponse(createResult groups.CreateResult,)*CreateNetworkingV2ExtensionsSecurityGroupsResponse {
    return &CreateNetworkingV2ExtensionsSecurityGroupsResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateNetworkingV2ExtensionsSecurityGroups(req *CreateNetworkingV2ExtensionsSecurityGroupsRequest)(*CreateNetworkingV2ExtensionsSecurityGroupsResponse){
    return NewCreateNetworkingV2ExtensionsSecurityGroupsResponse(groups.Create(oc.Client,req.Opts, ))

}
//request struct for the UpdateNetworkingV2ExtensionsSecurityGroups
type UpdateNetworkingV2ExtensionsSecurityGroupsRequest struct{
    Id string
    Opts groups.UpdateOpts
}

func NewUpdateNetworkingV2ExtensionsSecurityGroupsRequest()*UpdateNetworkingV2ExtensionsSecurityGroupsRequest{
    return &UpdateNetworkingV2ExtensionsSecurityGroupsRequest{}
}

//response struct for the UpdateNetworkingV2ExtensionsSecurityGroups
type UpdateNetworkingV2ExtensionsSecurityGroupsResponse struct{
    UpdateResult groups.UpdateResult
}

func NewUpdateNetworkingV2ExtensionsSecurityGroupsResponse(updateResult groups.UpdateResult,)*UpdateNetworkingV2ExtensionsSecurityGroupsResponse {
    return &UpdateNetworkingV2ExtensionsSecurityGroupsResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateNetworkingV2ExtensionsSecurityGroups(req *UpdateNetworkingV2ExtensionsSecurityGroupsRequest)(*UpdateNetworkingV2ExtensionsSecurityGroupsResponse){
    return NewUpdateNetworkingV2ExtensionsSecurityGroupsResponse(groups.Update(oc.Client,req.Id,req.Opts, ))

}
//request struct for the GetNetworkingV2ExtensionsSecurityGroups
type GetNetworkingV2ExtensionsSecurityGroupsRequest struct{
    Id string
}

func NewGetNetworkingV2ExtensionsSecurityGroupsRequest()*GetNetworkingV2ExtensionsSecurityGroupsRequest{
    return &GetNetworkingV2ExtensionsSecurityGroupsRequest{}
}

//response struct for the GetNetworkingV2ExtensionsSecurityGroups
type GetNetworkingV2ExtensionsSecurityGroupsResponse struct{
    GetResult groups.GetResult
}

func NewGetNetworkingV2ExtensionsSecurityGroupsResponse(getResult groups.GetResult,)*GetNetworkingV2ExtensionsSecurityGroupsResponse {
    return &GetNetworkingV2ExtensionsSecurityGroupsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetNetworkingV2ExtensionsSecurityGroups(req *GetNetworkingV2ExtensionsSecurityGroupsRequest)(*GetNetworkingV2ExtensionsSecurityGroupsResponse){
    return NewGetNetworkingV2ExtensionsSecurityGroupsResponse(groups.Get(oc.Client,req.Id, ))

}
//request struct for the DeleteNetworkingV2ExtensionsSecurityGroups
type DeleteNetworkingV2ExtensionsSecurityGroupsRequest struct{
    Id string
}

func NewDeleteNetworkingV2ExtensionsSecurityGroupsRequest()*DeleteNetworkingV2ExtensionsSecurityGroupsRequest{
    return &DeleteNetworkingV2ExtensionsSecurityGroupsRequest{}
}

//response struct for the DeleteNetworkingV2ExtensionsSecurityGroups
type DeleteNetworkingV2ExtensionsSecurityGroupsResponse struct{
    DeleteResult groups.DeleteResult
}

func NewDeleteNetworkingV2ExtensionsSecurityGroupsResponse(deleteResult groups.DeleteResult,)*DeleteNetworkingV2ExtensionsSecurityGroupsResponse {
    return &DeleteNetworkingV2ExtensionsSecurityGroupsResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteNetworkingV2ExtensionsSecurityGroups(req *DeleteNetworkingV2ExtensionsSecurityGroupsRequest)(*DeleteNetworkingV2ExtensionsSecurityGroupsResponse){
    return NewDeleteNetworkingV2ExtensionsSecurityGroupsResponse(groups.Delete(oc.Client,req.Id, ))

}