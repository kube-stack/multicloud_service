package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/identity/v3/policies"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListIdentityV3Policies
type ListIdentityV3PoliciesRequest struct{
    Opts policies.ListOpts
}

func NewListIdentityV3PoliciesRequest()*ListIdentityV3PoliciesRequest{
    return &ListIdentityV3PoliciesRequest{}
}

//response struct for the ListIdentityV3Policies
type ListIdentityV3PoliciesResponse struct{
    Pager pagination.Pager
}

func NewListIdentityV3PoliciesResponse(pager pagination.Pager,)*ListIdentityV3PoliciesResponse {
    return &ListIdentityV3PoliciesResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListIdentityV3Policies(req *ListIdentityV3PoliciesRequest)(*ListIdentityV3PoliciesResponse){
    return NewListIdentityV3PoliciesResponse(policies.List(oc.Client,req.Opts, ))

}
//request struct for the CreateIdentityV3Policies
type CreateIdentityV3PoliciesRequest struct{
    Opts policies.CreateOpts
}

func NewCreateIdentityV3PoliciesRequest()*CreateIdentityV3PoliciesRequest{
    return &CreateIdentityV3PoliciesRequest{}
}

//response struct for the CreateIdentityV3Policies
type CreateIdentityV3PoliciesResponse struct{
    CreateResult policies.CreateResult
}

func NewCreateIdentityV3PoliciesResponse(createResult policies.CreateResult,)*CreateIdentityV3PoliciesResponse {
    return &CreateIdentityV3PoliciesResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateIdentityV3Policies(req *CreateIdentityV3PoliciesRequest)(*CreateIdentityV3PoliciesResponse){
    return NewCreateIdentityV3PoliciesResponse(policies.Create(oc.Client,req.Opts, ))

}
//request struct for the GetIdentityV3Policies
type GetIdentityV3PoliciesRequest struct{
    PolicyID string
}

func NewGetIdentityV3PoliciesRequest()*GetIdentityV3PoliciesRequest{
    return &GetIdentityV3PoliciesRequest{}
}

//response struct for the GetIdentityV3Policies
type GetIdentityV3PoliciesResponse struct{
    GetResult policies.GetResult
}

func NewGetIdentityV3PoliciesResponse(getResult policies.GetResult,)*GetIdentityV3PoliciesResponse {
    return &GetIdentityV3PoliciesResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetIdentityV3Policies(req *GetIdentityV3PoliciesRequest)(*GetIdentityV3PoliciesResponse){
    return NewGetIdentityV3PoliciesResponse(policies.Get(oc.Client,req.PolicyID, ))

}
//request struct for the UpdateIdentityV3Policies
type UpdateIdentityV3PoliciesRequest struct{
    PolicyID string
    Opts policies.UpdateOpts
}

func NewUpdateIdentityV3PoliciesRequest()*UpdateIdentityV3PoliciesRequest{
    return &UpdateIdentityV3PoliciesRequest{}
}

//response struct for the UpdateIdentityV3Policies
type UpdateIdentityV3PoliciesResponse struct{
    UpdateResult policies.UpdateResult
}

func NewUpdateIdentityV3PoliciesResponse(updateResult policies.UpdateResult,)*UpdateIdentityV3PoliciesResponse {
    return &UpdateIdentityV3PoliciesResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateIdentityV3Policies(req *UpdateIdentityV3PoliciesRequest)(*UpdateIdentityV3PoliciesResponse){
    return NewUpdateIdentityV3PoliciesResponse(policies.Update(oc.Client,req.PolicyID,req.Opts, ))

}
//request struct for the DeleteIdentityV3Policies
type DeleteIdentityV3PoliciesRequest struct{
    PolicyID string
}

func NewDeleteIdentityV3PoliciesRequest()*DeleteIdentityV3PoliciesRequest{
    return &DeleteIdentityV3PoliciesRequest{}
}

//response struct for the DeleteIdentityV3Policies
type DeleteIdentityV3PoliciesResponse struct{
    DeleteResult policies.DeleteResult
}

func NewDeleteIdentityV3PoliciesResponse(deleteResult policies.DeleteResult,)*DeleteIdentityV3PoliciesResponse {
    return &DeleteIdentityV3PoliciesResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteIdentityV3Policies(req *DeleteIdentityV3PoliciesRequest)(*DeleteIdentityV3PoliciesResponse){
    return NewDeleteIdentityV3PoliciesResponse(policies.Delete(oc.Client,req.PolicyID, ))

}