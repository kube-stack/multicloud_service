package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/fwaas/policies"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListNetworkingV2ExtensionsFwaasPolicies
type ListNetworkingV2ExtensionsFwaasPoliciesRequest struct{
    Opts policies.ListOpts
}

func NewListNetworkingV2ExtensionsFwaasPoliciesRequest()*ListNetworkingV2ExtensionsFwaasPoliciesRequest{
    return &ListNetworkingV2ExtensionsFwaasPoliciesRequest{}
}

//response struct for the ListNetworkingV2ExtensionsFwaasPolicies
type ListNetworkingV2ExtensionsFwaasPoliciesResponse struct{
    Pager pagination.Pager
}

func NewListNetworkingV2ExtensionsFwaasPoliciesResponse(pager pagination.Pager,)*ListNetworkingV2ExtensionsFwaasPoliciesResponse {
    return &ListNetworkingV2ExtensionsFwaasPoliciesResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListNetworkingV2ExtensionsFwaasPolicies(req *ListNetworkingV2ExtensionsFwaasPoliciesRequest)(*ListNetworkingV2ExtensionsFwaasPoliciesResponse){
    return NewListNetworkingV2ExtensionsFwaasPoliciesResponse(policies.List(oc.Client,req.Opts, ))

}
//request struct for the CreateNetworkingV2ExtensionsFwaasPolicies
type CreateNetworkingV2ExtensionsFwaasPoliciesRequest struct{
    Opts policies.CreateOpts
}

func NewCreateNetworkingV2ExtensionsFwaasPoliciesRequest()*CreateNetworkingV2ExtensionsFwaasPoliciesRequest{
    return &CreateNetworkingV2ExtensionsFwaasPoliciesRequest{}
}

//response struct for the CreateNetworkingV2ExtensionsFwaasPolicies
type CreateNetworkingV2ExtensionsFwaasPoliciesResponse struct{
    CreateResult policies.CreateResult
}

func NewCreateNetworkingV2ExtensionsFwaasPoliciesResponse(createResult policies.CreateResult,)*CreateNetworkingV2ExtensionsFwaasPoliciesResponse {
    return &CreateNetworkingV2ExtensionsFwaasPoliciesResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateNetworkingV2ExtensionsFwaasPolicies(req *CreateNetworkingV2ExtensionsFwaasPoliciesRequest)(*CreateNetworkingV2ExtensionsFwaasPoliciesResponse){
    return NewCreateNetworkingV2ExtensionsFwaasPoliciesResponse(policies.Create(oc.Client,req.Opts, ))

}
//request struct for the GetNetworkingV2ExtensionsFwaasPolicies
type GetNetworkingV2ExtensionsFwaasPoliciesRequest struct{
    Id string
}

func NewGetNetworkingV2ExtensionsFwaasPoliciesRequest()*GetNetworkingV2ExtensionsFwaasPoliciesRequest{
    return &GetNetworkingV2ExtensionsFwaasPoliciesRequest{}
}

//response struct for the GetNetworkingV2ExtensionsFwaasPolicies
type GetNetworkingV2ExtensionsFwaasPoliciesResponse struct{
    GetResult policies.GetResult
}

func NewGetNetworkingV2ExtensionsFwaasPoliciesResponse(getResult policies.GetResult,)*GetNetworkingV2ExtensionsFwaasPoliciesResponse {
    return &GetNetworkingV2ExtensionsFwaasPoliciesResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetNetworkingV2ExtensionsFwaasPolicies(req *GetNetworkingV2ExtensionsFwaasPoliciesRequest)(*GetNetworkingV2ExtensionsFwaasPoliciesResponse){
    return NewGetNetworkingV2ExtensionsFwaasPoliciesResponse(policies.Get(oc.Client,req.Id, ))

}
//request struct for the UpdateNetworkingV2ExtensionsFwaasPolicies
type UpdateNetworkingV2ExtensionsFwaasPoliciesRequest struct{
    Id string
    Opts policies.UpdateOpts
}

func NewUpdateNetworkingV2ExtensionsFwaasPoliciesRequest()*UpdateNetworkingV2ExtensionsFwaasPoliciesRequest{
    return &UpdateNetworkingV2ExtensionsFwaasPoliciesRequest{}
}

//response struct for the UpdateNetworkingV2ExtensionsFwaasPolicies
type UpdateNetworkingV2ExtensionsFwaasPoliciesResponse struct{
    UpdateResult policies.UpdateResult
}

func NewUpdateNetworkingV2ExtensionsFwaasPoliciesResponse(updateResult policies.UpdateResult,)*UpdateNetworkingV2ExtensionsFwaasPoliciesResponse {
    return &UpdateNetworkingV2ExtensionsFwaasPoliciesResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateNetworkingV2ExtensionsFwaasPolicies(req *UpdateNetworkingV2ExtensionsFwaasPoliciesRequest)(*UpdateNetworkingV2ExtensionsFwaasPoliciesResponse){
    return NewUpdateNetworkingV2ExtensionsFwaasPoliciesResponse(policies.Update(oc.Client,req.Id,req.Opts, ))

}
//request struct for the DeleteNetworkingV2ExtensionsFwaasPolicies
type DeleteNetworkingV2ExtensionsFwaasPoliciesRequest struct{
    Id string
}

func NewDeleteNetworkingV2ExtensionsFwaasPoliciesRequest()*DeleteNetworkingV2ExtensionsFwaasPoliciesRequest{
    return &DeleteNetworkingV2ExtensionsFwaasPoliciesRequest{}
}

//response struct for the DeleteNetworkingV2ExtensionsFwaasPolicies
type DeleteNetworkingV2ExtensionsFwaasPoliciesResponse struct{
    DeleteResult policies.DeleteResult
}

func NewDeleteNetworkingV2ExtensionsFwaasPoliciesResponse(deleteResult policies.DeleteResult,)*DeleteNetworkingV2ExtensionsFwaasPoliciesResponse {
    return &DeleteNetworkingV2ExtensionsFwaasPoliciesResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteNetworkingV2ExtensionsFwaasPolicies(req *DeleteNetworkingV2ExtensionsFwaasPoliciesRequest)(*DeleteNetworkingV2ExtensionsFwaasPoliciesResponse){
    return NewDeleteNetworkingV2ExtensionsFwaasPoliciesResponse(policies.Delete(oc.Client,req.Id, ))

}
//request struct for the AddRuleNetworkingV2ExtensionsFwaasPolicies
type AddRuleNetworkingV2ExtensionsFwaasPoliciesRequest struct{
    Id string
    Opts policies.InsertRuleOpts
}

func NewAddRuleNetworkingV2ExtensionsFwaasPoliciesRequest()*AddRuleNetworkingV2ExtensionsFwaasPoliciesRequest{
    return &AddRuleNetworkingV2ExtensionsFwaasPoliciesRequest{}
}

//response struct for the AddRuleNetworkingV2ExtensionsFwaasPolicies
type AddRuleNetworkingV2ExtensionsFwaasPoliciesResponse struct{
    InsertRuleResult policies.InsertRuleResult
}

func NewAddRuleNetworkingV2ExtensionsFwaasPoliciesResponse(insertRuleResult policies.InsertRuleResult,)*AddRuleNetworkingV2ExtensionsFwaasPoliciesResponse {
    return &AddRuleNetworkingV2ExtensionsFwaasPoliciesResponse{
            InsertRuleResult:insertRuleResult,
    }
}

// action function
func (oc *OpenstackClient) AddRuleNetworkingV2ExtensionsFwaasPolicies(req *AddRuleNetworkingV2ExtensionsFwaasPoliciesRequest)(*AddRuleNetworkingV2ExtensionsFwaasPoliciesResponse){
    return NewAddRuleNetworkingV2ExtensionsFwaasPoliciesResponse(policies.AddRule(oc.Client,req.Id,req.Opts, ))

}
//request struct for the RemoveRuleNetworkingV2ExtensionsFwaasPolicies
type RemoveRuleNetworkingV2ExtensionsFwaasPoliciesRequest struct{
    Id string
    RuleID string
}

func NewRemoveRuleNetworkingV2ExtensionsFwaasPoliciesRequest()*RemoveRuleNetworkingV2ExtensionsFwaasPoliciesRequest{
    return &RemoveRuleNetworkingV2ExtensionsFwaasPoliciesRequest{}
}

//response struct for the RemoveRuleNetworkingV2ExtensionsFwaasPolicies
type RemoveRuleNetworkingV2ExtensionsFwaasPoliciesResponse struct{
    RemoveRuleResult policies.RemoveRuleResult
}

func NewRemoveRuleNetworkingV2ExtensionsFwaasPoliciesResponse(removeRuleResult policies.RemoveRuleResult,)*RemoveRuleNetworkingV2ExtensionsFwaasPoliciesResponse {
    return &RemoveRuleNetworkingV2ExtensionsFwaasPoliciesResponse{
            RemoveRuleResult:removeRuleResult,
    }
}

// action function
func (oc *OpenstackClient) RemoveRuleNetworkingV2ExtensionsFwaasPolicies(req *RemoveRuleNetworkingV2ExtensionsFwaasPoliciesRequest)(*RemoveRuleNetworkingV2ExtensionsFwaasPoliciesResponse){
    return NewRemoveRuleNetworkingV2ExtensionsFwaasPoliciesResponse(policies.RemoveRule(oc.Client,req.Id,req.RuleID, ))

}