package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/secgroups"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListByServerComputeV2ExtensionsSecgroups
type ListByServerComputeV2ExtensionsSecgroupsRequest struct{
    ServerID string
}

func NewListByServerComputeV2ExtensionsSecgroupsRequest()*ListByServerComputeV2ExtensionsSecgroupsRequest{
    return &ListByServerComputeV2ExtensionsSecgroupsRequest{}
}

//response struct for the ListByServerComputeV2ExtensionsSecgroups
type ListByServerComputeV2ExtensionsSecgroupsResponse struct{
    Pager pagination.Pager
}

func NewListByServerComputeV2ExtensionsSecgroupsResponse(pager pagination.Pager,)*ListByServerComputeV2ExtensionsSecgroupsResponse {
    return &ListByServerComputeV2ExtensionsSecgroupsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListByServerComputeV2ExtensionsSecgroups(req *ListByServerComputeV2ExtensionsSecgroupsRequest)(*ListByServerComputeV2ExtensionsSecgroupsResponse){
    return NewListByServerComputeV2ExtensionsSecgroupsResponse(secgroups.ListByServer(oc.Client,req.ServerID, ))

}
//request struct for the CreateComputeV2ExtensionsSecgroups
type CreateComputeV2ExtensionsSecgroupsRequest struct{
    Opts secgroups.CreateOpts
}

func NewCreateComputeV2ExtensionsSecgroupsRequest()*CreateComputeV2ExtensionsSecgroupsRequest{
    return &CreateComputeV2ExtensionsSecgroupsRequest{}
}

//response struct for the CreateComputeV2ExtensionsSecgroups
type CreateComputeV2ExtensionsSecgroupsResponse struct{
    CreateResult secgroups.CreateResult
}

func NewCreateComputeV2ExtensionsSecgroupsResponse(createResult secgroups.CreateResult,)*CreateComputeV2ExtensionsSecgroupsResponse {
    return &CreateComputeV2ExtensionsSecgroupsResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateComputeV2ExtensionsSecgroups(req *CreateComputeV2ExtensionsSecgroupsRequest)(*CreateComputeV2ExtensionsSecgroupsResponse){
    return NewCreateComputeV2ExtensionsSecgroupsResponse(secgroups.Create(oc.Client,req.Opts, ))

}
//request struct for the UpdateComputeV2ExtensionsSecgroups
type UpdateComputeV2ExtensionsSecgroupsRequest struct{
    Id string
    Opts secgroups.UpdateOpts
}

func NewUpdateComputeV2ExtensionsSecgroupsRequest()*UpdateComputeV2ExtensionsSecgroupsRequest{
    return &UpdateComputeV2ExtensionsSecgroupsRequest{}
}

//response struct for the UpdateComputeV2ExtensionsSecgroups
type UpdateComputeV2ExtensionsSecgroupsResponse struct{
    UpdateResult secgroups.UpdateResult
}

func NewUpdateComputeV2ExtensionsSecgroupsResponse(updateResult secgroups.UpdateResult,)*UpdateComputeV2ExtensionsSecgroupsResponse {
    return &UpdateComputeV2ExtensionsSecgroupsResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateComputeV2ExtensionsSecgroups(req *UpdateComputeV2ExtensionsSecgroupsRequest)(*UpdateComputeV2ExtensionsSecgroupsResponse){
    return NewUpdateComputeV2ExtensionsSecgroupsResponse(secgroups.Update(oc.Client,req.Id,req.Opts, ))

}
//request struct for the GetComputeV2ExtensionsSecgroups
type GetComputeV2ExtensionsSecgroupsRequest struct{
    Id string
}

func NewGetComputeV2ExtensionsSecgroupsRequest()*GetComputeV2ExtensionsSecgroupsRequest{
    return &GetComputeV2ExtensionsSecgroupsRequest{}
}

//response struct for the GetComputeV2ExtensionsSecgroups
type GetComputeV2ExtensionsSecgroupsResponse struct{
    GetResult secgroups.GetResult
}

func NewGetComputeV2ExtensionsSecgroupsResponse(getResult secgroups.GetResult,)*GetComputeV2ExtensionsSecgroupsResponse {
    return &GetComputeV2ExtensionsSecgroupsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetComputeV2ExtensionsSecgroups(req *GetComputeV2ExtensionsSecgroupsRequest)(*GetComputeV2ExtensionsSecgroupsResponse){
    return NewGetComputeV2ExtensionsSecgroupsResponse(secgroups.Get(oc.Client,req.Id, ))

}
//request struct for the DeleteComputeV2ExtensionsSecgroups
type DeleteComputeV2ExtensionsSecgroupsRequest struct{
    Id string
}

func NewDeleteComputeV2ExtensionsSecgroupsRequest()*DeleteComputeV2ExtensionsSecgroupsRequest{
    return &DeleteComputeV2ExtensionsSecgroupsRequest{}
}

//response struct for the DeleteComputeV2ExtensionsSecgroups
type DeleteComputeV2ExtensionsSecgroupsResponse struct{
    DeleteResult secgroups.DeleteResult
}

func NewDeleteComputeV2ExtensionsSecgroupsResponse(deleteResult secgroups.DeleteResult,)*DeleteComputeV2ExtensionsSecgroupsResponse {
    return &DeleteComputeV2ExtensionsSecgroupsResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteComputeV2ExtensionsSecgroups(req *DeleteComputeV2ExtensionsSecgroupsRequest)(*DeleteComputeV2ExtensionsSecgroupsResponse){
    return NewDeleteComputeV2ExtensionsSecgroupsResponse(secgroups.Delete(oc.Client,req.Id, ))

}
//request struct for the CreateRuleComputeV2ExtensionsSecgroups
type CreateRuleComputeV2ExtensionsSecgroupsRequest struct{
    Opts secgroups.CreateRuleOpts
}

func NewCreateRuleComputeV2ExtensionsSecgroupsRequest()*CreateRuleComputeV2ExtensionsSecgroupsRequest{
    return &CreateRuleComputeV2ExtensionsSecgroupsRequest{}
}

//response struct for the CreateRuleComputeV2ExtensionsSecgroups
type CreateRuleComputeV2ExtensionsSecgroupsResponse struct{
    CreateRuleResult secgroups.CreateRuleResult
}

func NewCreateRuleComputeV2ExtensionsSecgroupsResponse(createRuleResult secgroups.CreateRuleResult,)*CreateRuleComputeV2ExtensionsSecgroupsResponse {
    return &CreateRuleComputeV2ExtensionsSecgroupsResponse{
            CreateRuleResult:createRuleResult,
    }
}

// action function
func (oc *OpenstackClient) CreateRuleComputeV2ExtensionsSecgroups(req *CreateRuleComputeV2ExtensionsSecgroupsRequest)(*CreateRuleComputeV2ExtensionsSecgroupsResponse){
    return NewCreateRuleComputeV2ExtensionsSecgroupsResponse(secgroups.CreateRule(oc.Client,req.Opts, ))

}
//request struct for the DeleteRuleComputeV2ExtensionsSecgroups
type DeleteRuleComputeV2ExtensionsSecgroupsRequest struct{
    Id string
}

func NewDeleteRuleComputeV2ExtensionsSecgroupsRequest()*DeleteRuleComputeV2ExtensionsSecgroupsRequest{
    return &DeleteRuleComputeV2ExtensionsSecgroupsRequest{}
}

//response struct for the DeleteRuleComputeV2ExtensionsSecgroups
type DeleteRuleComputeV2ExtensionsSecgroupsResponse struct{
    DeleteRuleResult secgroups.DeleteRuleResult
}

func NewDeleteRuleComputeV2ExtensionsSecgroupsResponse(deleteRuleResult secgroups.DeleteRuleResult,)*DeleteRuleComputeV2ExtensionsSecgroupsResponse {
    return &DeleteRuleComputeV2ExtensionsSecgroupsResponse{
            DeleteRuleResult:deleteRuleResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteRuleComputeV2ExtensionsSecgroups(req *DeleteRuleComputeV2ExtensionsSecgroupsRequest)(*DeleteRuleComputeV2ExtensionsSecgroupsResponse){
    return NewDeleteRuleComputeV2ExtensionsSecgroupsResponse(secgroups.DeleteRule(oc.Client,req.Id, ))

}
//request struct for the AddServerComputeV2ExtensionsSecgroups
type AddServerComputeV2ExtensionsSecgroupsRequest struct{
    ServerID string
    GroupName string
}

func NewAddServerComputeV2ExtensionsSecgroupsRequest()*AddServerComputeV2ExtensionsSecgroupsRequest{
    return &AddServerComputeV2ExtensionsSecgroupsRequest{}
}

//response struct for the AddServerComputeV2ExtensionsSecgroups
type AddServerComputeV2ExtensionsSecgroupsResponse struct{
    AddServerResult secgroups.AddServerResult
}

func NewAddServerComputeV2ExtensionsSecgroupsResponse(addServerResult secgroups.AddServerResult,)*AddServerComputeV2ExtensionsSecgroupsResponse {
    return &AddServerComputeV2ExtensionsSecgroupsResponse{
            AddServerResult:addServerResult,
    }
}

// action function
func (oc *OpenstackClient) AddServerComputeV2ExtensionsSecgroups(req *AddServerComputeV2ExtensionsSecgroupsRequest)(*AddServerComputeV2ExtensionsSecgroupsResponse){
    return NewAddServerComputeV2ExtensionsSecgroupsResponse(secgroups.AddServer(oc.Client,req.ServerID,req.GroupName, ))

}
//request struct for the RemoveServerComputeV2ExtensionsSecgroups
type RemoveServerComputeV2ExtensionsSecgroupsRequest struct{
    ServerID string
    GroupName string
}

func NewRemoveServerComputeV2ExtensionsSecgroupsRequest()*RemoveServerComputeV2ExtensionsSecgroupsRequest{
    return &RemoveServerComputeV2ExtensionsSecgroupsRequest{}
}

//response struct for the RemoveServerComputeV2ExtensionsSecgroups
type RemoveServerComputeV2ExtensionsSecgroupsResponse struct{
    RemoveServerResult secgroups.RemoveServerResult
}

func NewRemoveServerComputeV2ExtensionsSecgroupsResponse(removeServerResult secgroups.RemoveServerResult,)*RemoveServerComputeV2ExtensionsSecgroupsResponse {
    return &RemoveServerComputeV2ExtensionsSecgroupsResponse{
            RemoveServerResult:removeServerResult,
    }
}

// action function
func (oc *OpenstackClient) RemoveServerComputeV2ExtensionsSecgroups(req *RemoveServerComputeV2ExtensionsSecgroupsRequest)(*RemoveServerComputeV2ExtensionsSecgroupsResponse){
    return NewRemoveServerComputeV2ExtensionsSecgroupsResponse(secgroups.RemoveServer(oc.Client,req.ServerID,req.GroupName, ))

}