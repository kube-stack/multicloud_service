package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/loadbalancer/v2/pools"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListLoadbalancerV2Pools
type ListLoadbalancerV2PoolsRequest struct{
    Opts pools.ListOpts
}

func NewListLoadbalancerV2PoolsRequest()*ListLoadbalancerV2PoolsRequest{
    return &ListLoadbalancerV2PoolsRequest{}
}

//response struct for the ListLoadbalancerV2Pools
type ListLoadbalancerV2PoolsResponse struct{
    Pager pagination.Pager
}

func NewListLoadbalancerV2PoolsResponse(pager pagination.Pager,)*ListLoadbalancerV2PoolsResponse {
    return &ListLoadbalancerV2PoolsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListLoadbalancerV2Pools(req *ListLoadbalancerV2PoolsRequest)(*ListLoadbalancerV2PoolsResponse){
    return NewListLoadbalancerV2PoolsResponse(pools.List(oc.Client,req.Opts, ))

}
//request struct for the CreateLoadbalancerV2Pools
type CreateLoadbalancerV2PoolsRequest struct{
    Opts pools.CreateOpts
}

func NewCreateLoadbalancerV2PoolsRequest()*CreateLoadbalancerV2PoolsRequest{
    return &CreateLoadbalancerV2PoolsRequest{}
}

//response struct for the CreateLoadbalancerV2Pools
type CreateLoadbalancerV2PoolsResponse struct{
    CreateResult pools.CreateResult
}

func NewCreateLoadbalancerV2PoolsResponse(createResult pools.CreateResult,)*CreateLoadbalancerV2PoolsResponse {
    return &CreateLoadbalancerV2PoolsResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateLoadbalancerV2Pools(req *CreateLoadbalancerV2PoolsRequest)(*CreateLoadbalancerV2PoolsResponse){
    return NewCreateLoadbalancerV2PoolsResponse(pools.Create(oc.Client,req.Opts, ))

}
//request struct for the GetLoadbalancerV2Pools
type GetLoadbalancerV2PoolsRequest struct{
    Id string
}

func NewGetLoadbalancerV2PoolsRequest()*GetLoadbalancerV2PoolsRequest{
    return &GetLoadbalancerV2PoolsRequest{}
}

//response struct for the GetLoadbalancerV2Pools
type GetLoadbalancerV2PoolsResponse struct{
    GetResult pools.GetResult
}

func NewGetLoadbalancerV2PoolsResponse(getResult pools.GetResult,)*GetLoadbalancerV2PoolsResponse {
    return &GetLoadbalancerV2PoolsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetLoadbalancerV2Pools(req *GetLoadbalancerV2PoolsRequest)(*GetLoadbalancerV2PoolsResponse){
    return NewGetLoadbalancerV2PoolsResponse(pools.Get(oc.Client,req.Id, ))

}
//request struct for the UpdateLoadbalancerV2Pools
type UpdateLoadbalancerV2PoolsRequest struct{
    Id string
    Opts pools.UpdateOpts
}

func NewUpdateLoadbalancerV2PoolsRequest()*UpdateLoadbalancerV2PoolsRequest{
    return &UpdateLoadbalancerV2PoolsRequest{}
}

//response struct for the UpdateLoadbalancerV2Pools
type UpdateLoadbalancerV2PoolsResponse struct{
    UpdateResult pools.UpdateResult
}

func NewUpdateLoadbalancerV2PoolsResponse(updateResult pools.UpdateResult,)*UpdateLoadbalancerV2PoolsResponse {
    return &UpdateLoadbalancerV2PoolsResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateLoadbalancerV2Pools(req *UpdateLoadbalancerV2PoolsRequest)(*UpdateLoadbalancerV2PoolsResponse){
    return NewUpdateLoadbalancerV2PoolsResponse(pools.Update(oc.Client,req.Id,req.Opts, ))

}
//request struct for the DeleteLoadbalancerV2Pools
type DeleteLoadbalancerV2PoolsRequest struct{
    Id string
}

func NewDeleteLoadbalancerV2PoolsRequest()*DeleteLoadbalancerV2PoolsRequest{
    return &DeleteLoadbalancerV2PoolsRequest{}
}

//response struct for the DeleteLoadbalancerV2Pools
type DeleteLoadbalancerV2PoolsResponse struct{
    DeleteResult pools.DeleteResult
}

func NewDeleteLoadbalancerV2PoolsResponse(deleteResult pools.DeleteResult,)*DeleteLoadbalancerV2PoolsResponse {
    return &DeleteLoadbalancerV2PoolsResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteLoadbalancerV2Pools(req *DeleteLoadbalancerV2PoolsRequest)(*DeleteLoadbalancerV2PoolsResponse){
    return NewDeleteLoadbalancerV2PoolsResponse(pools.Delete(oc.Client,req.Id, ))

}
//request struct for the ListMembersLoadbalancerV2Pools
type ListMembersLoadbalancerV2PoolsRequest struct{
    PoolID string
    Opts pools.ListMembersOpts
}

func NewListMembersLoadbalancerV2PoolsRequest()*ListMembersLoadbalancerV2PoolsRequest{
    return &ListMembersLoadbalancerV2PoolsRequest{}
}

//response struct for the ListMembersLoadbalancerV2Pools
type ListMembersLoadbalancerV2PoolsResponse struct{
    Pager pagination.Pager
}

func NewListMembersLoadbalancerV2PoolsResponse(pager pagination.Pager,)*ListMembersLoadbalancerV2PoolsResponse {
    return &ListMembersLoadbalancerV2PoolsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListMembersLoadbalancerV2Pools(req *ListMembersLoadbalancerV2PoolsRequest)(*ListMembersLoadbalancerV2PoolsResponse){
    return NewListMembersLoadbalancerV2PoolsResponse(pools.ListMembers(oc.Client,req.PoolID,req.Opts, ))

}
//request struct for the CreateMemberLoadbalancerV2Pools
type CreateMemberLoadbalancerV2PoolsRequest struct{
    PoolID string
    Opts pools.CreateMemberOpts
}

func NewCreateMemberLoadbalancerV2PoolsRequest()*CreateMemberLoadbalancerV2PoolsRequest{
    return &CreateMemberLoadbalancerV2PoolsRequest{}
}

//response struct for the CreateMemberLoadbalancerV2Pools
type CreateMemberLoadbalancerV2PoolsResponse struct{
    CreateMemberResult pools.CreateMemberResult
}

func NewCreateMemberLoadbalancerV2PoolsResponse(createMemberResult pools.CreateMemberResult,)*CreateMemberLoadbalancerV2PoolsResponse {
    return &CreateMemberLoadbalancerV2PoolsResponse{
            CreateMemberResult:createMemberResult,
    }
}

// action function
func (oc *OpenstackClient) CreateMemberLoadbalancerV2Pools(req *CreateMemberLoadbalancerV2PoolsRequest)(*CreateMemberLoadbalancerV2PoolsResponse){
    return NewCreateMemberLoadbalancerV2PoolsResponse(pools.CreateMember(oc.Client,req.PoolID,req.Opts, ))

}
//request struct for the GetMemberLoadbalancerV2Pools
type GetMemberLoadbalancerV2PoolsRequest struct{
    PoolID string
    MemberID string
}

func NewGetMemberLoadbalancerV2PoolsRequest()*GetMemberLoadbalancerV2PoolsRequest{
    return &GetMemberLoadbalancerV2PoolsRequest{}
}

//response struct for the GetMemberLoadbalancerV2Pools
type GetMemberLoadbalancerV2PoolsResponse struct{
    GetMemberResult pools.GetMemberResult
}

func NewGetMemberLoadbalancerV2PoolsResponse(getMemberResult pools.GetMemberResult,)*GetMemberLoadbalancerV2PoolsResponse {
    return &GetMemberLoadbalancerV2PoolsResponse{
            GetMemberResult:getMemberResult,
    }
}

// action function
func (oc *OpenstackClient) GetMemberLoadbalancerV2Pools(req *GetMemberLoadbalancerV2PoolsRequest)(*GetMemberLoadbalancerV2PoolsResponse){
    return NewGetMemberLoadbalancerV2PoolsResponse(pools.GetMember(oc.Client,req.PoolID,req.MemberID, ))

}
//request struct for the UpdateMemberLoadbalancerV2Pools
type UpdateMemberLoadbalancerV2PoolsRequest struct{
    PoolID string
    MemberID string
    Opts pools.UpdateMemberOpts
}

func NewUpdateMemberLoadbalancerV2PoolsRequest()*UpdateMemberLoadbalancerV2PoolsRequest{
    return &UpdateMemberLoadbalancerV2PoolsRequest{}
}

//response struct for the UpdateMemberLoadbalancerV2Pools
type UpdateMemberLoadbalancerV2PoolsResponse struct{
    UpdateMemberResult pools.UpdateMemberResult
}

func NewUpdateMemberLoadbalancerV2PoolsResponse(updateMemberResult pools.UpdateMemberResult,)*UpdateMemberLoadbalancerV2PoolsResponse {
    return &UpdateMemberLoadbalancerV2PoolsResponse{
            UpdateMemberResult:updateMemberResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateMemberLoadbalancerV2Pools(req *UpdateMemberLoadbalancerV2PoolsRequest)(*UpdateMemberLoadbalancerV2PoolsResponse){
    return NewUpdateMemberLoadbalancerV2PoolsResponse(pools.UpdateMember(oc.Client,req.PoolID,req.MemberID,req.Opts, ))

}
//request struct for the DeleteMemberLoadbalancerV2Pools
type DeleteMemberLoadbalancerV2PoolsRequest struct{
    PoolID string
    MemberID string
}

func NewDeleteMemberLoadbalancerV2PoolsRequest()*DeleteMemberLoadbalancerV2PoolsRequest{
    return &DeleteMemberLoadbalancerV2PoolsRequest{}
}

//response struct for the DeleteMemberLoadbalancerV2Pools
type DeleteMemberLoadbalancerV2PoolsResponse struct{
    DeleteMemberResult pools.DeleteMemberResult
}

func NewDeleteMemberLoadbalancerV2PoolsResponse(deleteMemberResult pools.DeleteMemberResult,)*DeleteMemberLoadbalancerV2PoolsResponse {
    return &DeleteMemberLoadbalancerV2PoolsResponse{
            DeleteMemberResult:deleteMemberResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteMemberLoadbalancerV2Pools(req *DeleteMemberLoadbalancerV2PoolsRequest)(*DeleteMemberLoadbalancerV2PoolsResponse){
    return NewDeleteMemberLoadbalancerV2PoolsResponse(pools.DeleteMember(oc.Client,req.PoolID,req.MemberID, ))

}