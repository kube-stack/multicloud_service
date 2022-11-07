package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/networking/v2/subnets"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListNetworkingV2Subnets
type ListNetworkingV2SubnetsRequest struct{
    Opts subnets.ListOpts
}

func NewListNetworkingV2SubnetsRequest()*ListNetworkingV2SubnetsRequest{
    return &ListNetworkingV2SubnetsRequest{}
}

//response struct for the ListNetworkingV2Subnets
type ListNetworkingV2SubnetsResponse struct{
    Pager pagination.Pager
}

func NewListNetworkingV2SubnetsResponse(pager pagination.Pager,)*ListNetworkingV2SubnetsResponse {
    return &ListNetworkingV2SubnetsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListNetworkingV2Subnets(req *ListNetworkingV2SubnetsRequest)(*ListNetworkingV2SubnetsResponse){
    return NewListNetworkingV2SubnetsResponse(subnets.List(oc.Client,req.Opts, ))

}
//request struct for the GetNetworkingV2Subnets
type GetNetworkingV2SubnetsRequest struct{
    Id string
}

func NewGetNetworkingV2SubnetsRequest()*GetNetworkingV2SubnetsRequest{
    return &GetNetworkingV2SubnetsRequest{}
}

//response struct for the GetNetworkingV2Subnets
type GetNetworkingV2SubnetsResponse struct{
    GetResult subnets.GetResult
}

func NewGetNetworkingV2SubnetsResponse(getResult subnets.GetResult,)*GetNetworkingV2SubnetsResponse {
    return &GetNetworkingV2SubnetsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetNetworkingV2Subnets(req *GetNetworkingV2SubnetsRequest)(*GetNetworkingV2SubnetsResponse){
    return NewGetNetworkingV2SubnetsResponse(subnets.Get(oc.Client,req.Id, ))

}
//request struct for the CreateNetworkingV2Subnets
type CreateNetworkingV2SubnetsRequest struct{
    Opts subnets.CreateOpts
}

func NewCreateNetworkingV2SubnetsRequest()*CreateNetworkingV2SubnetsRequest{
    return &CreateNetworkingV2SubnetsRequest{}
}

//response struct for the CreateNetworkingV2Subnets
type CreateNetworkingV2SubnetsResponse struct{
    CreateResult subnets.CreateResult
}

func NewCreateNetworkingV2SubnetsResponse(createResult subnets.CreateResult,)*CreateNetworkingV2SubnetsResponse {
    return &CreateNetworkingV2SubnetsResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateNetworkingV2Subnets(req *CreateNetworkingV2SubnetsRequest)(*CreateNetworkingV2SubnetsResponse){
    return NewCreateNetworkingV2SubnetsResponse(subnets.Create(oc.Client,req.Opts, ))

}
//request struct for the UpdateNetworkingV2Subnets
type UpdateNetworkingV2SubnetsRequest struct{
    Id string
    Opts subnets.UpdateOpts
}

func NewUpdateNetworkingV2SubnetsRequest()*UpdateNetworkingV2SubnetsRequest{
    return &UpdateNetworkingV2SubnetsRequest{}
}

//response struct for the UpdateNetworkingV2Subnets
type UpdateNetworkingV2SubnetsResponse struct{
    UpdateResult subnets.UpdateResult
}

func NewUpdateNetworkingV2SubnetsResponse(updateResult subnets.UpdateResult,)*UpdateNetworkingV2SubnetsResponse {
    return &UpdateNetworkingV2SubnetsResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateNetworkingV2Subnets(req *UpdateNetworkingV2SubnetsRequest)(*UpdateNetworkingV2SubnetsResponse){
    return NewUpdateNetworkingV2SubnetsResponse(subnets.Update(oc.Client,req.Id,req.Opts, ))

}
//request struct for the DeleteNetworkingV2Subnets
type DeleteNetworkingV2SubnetsRequest struct{
    Id string
}

func NewDeleteNetworkingV2SubnetsRequest()*DeleteNetworkingV2SubnetsRequest{
    return &DeleteNetworkingV2SubnetsRequest{}
}

//response struct for the DeleteNetworkingV2Subnets
type DeleteNetworkingV2SubnetsResponse struct{
    DeleteResult subnets.DeleteResult
}

func NewDeleteNetworkingV2SubnetsResponse(deleteResult subnets.DeleteResult,)*DeleteNetworkingV2SubnetsResponse {
    return &DeleteNetworkingV2SubnetsResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteNetworkingV2Subnets(req *DeleteNetworkingV2SubnetsRequest)(*DeleteNetworkingV2SubnetsResponse){
    return NewDeleteNetworkingV2SubnetsResponse(subnets.Delete(oc.Client,req.Id, ))

}