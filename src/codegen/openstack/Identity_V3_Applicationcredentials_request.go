package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/identity/v3/applicationcredentials"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListIdentityV3Applicationcredentials
type ListIdentityV3ApplicationcredentialsRequest struct{
    UserID string
    Opts applicationcredentials.ListOpts
}

func NewListIdentityV3ApplicationcredentialsRequest()*ListIdentityV3ApplicationcredentialsRequest{
    return &ListIdentityV3ApplicationcredentialsRequest{}
}

//response struct for the ListIdentityV3Applicationcredentials
type ListIdentityV3ApplicationcredentialsResponse struct{
    Pager pagination.Pager
}

func NewListIdentityV3ApplicationcredentialsResponse(pager pagination.Pager,)*ListIdentityV3ApplicationcredentialsResponse {
    return &ListIdentityV3ApplicationcredentialsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListIdentityV3Applicationcredentials(req *ListIdentityV3ApplicationcredentialsRequest)(*ListIdentityV3ApplicationcredentialsResponse){
    return NewListIdentityV3ApplicationcredentialsResponse(applicationcredentials.List(oc.Client,req.UserID,req.Opts, ))

}
//request struct for the GetIdentityV3Applicationcredentials
type GetIdentityV3ApplicationcredentialsRequest struct{
    UserID string
    Id string
}

func NewGetIdentityV3ApplicationcredentialsRequest()*GetIdentityV3ApplicationcredentialsRequest{
    return &GetIdentityV3ApplicationcredentialsRequest{}
}

//response struct for the GetIdentityV3Applicationcredentials
type GetIdentityV3ApplicationcredentialsResponse struct{
    GetResult applicationcredentials.GetResult
}

func NewGetIdentityV3ApplicationcredentialsResponse(getResult applicationcredentials.GetResult,)*GetIdentityV3ApplicationcredentialsResponse {
    return &GetIdentityV3ApplicationcredentialsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetIdentityV3Applicationcredentials(req *GetIdentityV3ApplicationcredentialsRequest)(*GetIdentityV3ApplicationcredentialsResponse){
    return NewGetIdentityV3ApplicationcredentialsResponse(applicationcredentials.Get(oc.Client,req.UserID,req.Id, ))

}
//request struct for the CreateIdentityV3Applicationcredentials
type CreateIdentityV3ApplicationcredentialsRequest struct{
    UserID string
    Opts applicationcredentials.CreateOpts
}

func NewCreateIdentityV3ApplicationcredentialsRequest()*CreateIdentityV3ApplicationcredentialsRequest{
    return &CreateIdentityV3ApplicationcredentialsRequest{}
}

//response struct for the CreateIdentityV3Applicationcredentials
type CreateIdentityV3ApplicationcredentialsResponse struct{
    CreateResult applicationcredentials.CreateResult
}

func NewCreateIdentityV3ApplicationcredentialsResponse(createResult applicationcredentials.CreateResult,)*CreateIdentityV3ApplicationcredentialsResponse {
    return &CreateIdentityV3ApplicationcredentialsResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateIdentityV3Applicationcredentials(req *CreateIdentityV3ApplicationcredentialsRequest)(*CreateIdentityV3ApplicationcredentialsResponse){
    return NewCreateIdentityV3ApplicationcredentialsResponse(applicationcredentials.Create(oc.Client,req.UserID,req.Opts, ))

}
//request struct for the DeleteIdentityV3Applicationcredentials
type DeleteIdentityV3ApplicationcredentialsRequest struct{
    UserID string
    Id string
}

func NewDeleteIdentityV3ApplicationcredentialsRequest()*DeleteIdentityV3ApplicationcredentialsRequest{
    return &DeleteIdentityV3ApplicationcredentialsRequest{}
}

//response struct for the DeleteIdentityV3Applicationcredentials
type DeleteIdentityV3ApplicationcredentialsResponse struct{
    DeleteResult applicationcredentials.DeleteResult
}

func NewDeleteIdentityV3ApplicationcredentialsResponse(deleteResult applicationcredentials.DeleteResult,)*DeleteIdentityV3ApplicationcredentialsResponse {
    return &DeleteIdentityV3ApplicationcredentialsResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteIdentityV3Applicationcredentials(req *DeleteIdentityV3ApplicationcredentialsRequest)(*DeleteIdentityV3ApplicationcredentialsResponse){
    return NewDeleteIdentityV3ApplicationcredentialsResponse(applicationcredentials.Delete(oc.Client,req.UserID,req.Id, ))

}
//request struct for the ListAccessRulesIdentityV3Applicationcredentials
type ListAccessRulesIdentityV3ApplicationcredentialsRequest struct{
    UserID string
}

func NewListAccessRulesIdentityV3ApplicationcredentialsRequest()*ListAccessRulesIdentityV3ApplicationcredentialsRequest{
    return &ListAccessRulesIdentityV3ApplicationcredentialsRequest{}
}

//response struct for the ListAccessRulesIdentityV3Applicationcredentials
type ListAccessRulesIdentityV3ApplicationcredentialsResponse struct{
    Pager pagination.Pager
}

func NewListAccessRulesIdentityV3ApplicationcredentialsResponse(pager pagination.Pager,)*ListAccessRulesIdentityV3ApplicationcredentialsResponse {
    return &ListAccessRulesIdentityV3ApplicationcredentialsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListAccessRulesIdentityV3Applicationcredentials(req *ListAccessRulesIdentityV3ApplicationcredentialsRequest)(*ListAccessRulesIdentityV3ApplicationcredentialsResponse){
    return NewListAccessRulesIdentityV3ApplicationcredentialsResponse(applicationcredentials.ListAccessRules(oc.Client,req.UserID, ))

}
//request struct for the GetAccessRuleIdentityV3Applicationcredentials
type GetAccessRuleIdentityV3ApplicationcredentialsRequest struct{
    UserID string
    Id string
}

func NewGetAccessRuleIdentityV3ApplicationcredentialsRequest()*GetAccessRuleIdentityV3ApplicationcredentialsRequest{
    return &GetAccessRuleIdentityV3ApplicationcredentialsRequest{}
}

//response struct for the GetAccessRuleIdentityV3Applicationcredentials
type GetAccessRuleIdentityV3ApplicationcredentialsResponse struct{
    GetAccessRuleResult applicationcredentials.GetAccessRuleResult
}

func NewGetAccessRuleIdentityV3ApplicationcredentialsResponse(getAccessRuleResult applicationcredentials.GetAccessRuleResult,)*GetAccessRuleIdentityV3ApplicationcredentialsResponse {
    return &GetAccessRuleIdentityV3ApplicationcredentialsResponse{
            GetAccessRuleResult:getAccessRuleResult,
    }
}

// action function
func (oc *OpenstackClient) GetAccessRuleIdentityV3Applicationcredentials(req *GetAccessRuleIdentityV3ApplicationcredentialsRequest)(*GetAccessRuleIdentityV3ApplicationcredentialsResponse){
    return NewGetAccessRuleIdentityV3ApplicationcredentialsResponse(applicationcredentials.GetAccessRule(oc.Client,req.UserID,req.Id, ))

}
//request struct for the DeleteAccessRuleIdentityV3Applicationcredentials
type DeleteAccessRuleIdentityV3ApplicationcredentialsRequest struct{
    UserID string
    Id string
}

func NewDeleteAccessRuleIdentityV3ApplicationcredentialsRequest()*DeleteAccessRuleIdentityV3ApplicationcredentialsRequest{
    return &DeleteAccessRuleIdentityV3ApplicationcredentialsRequest{}
}

//response struct for the DeleteAccessRuleIdentityV3Applicationcredentials
type DeleteAccessRuleIdentityV3ApplicationcredentialsResponse struct{
    DeleteResult applicationcredentials.DeleteResult
}

func NewDeleteAccessRuleIdentityV3ApplicationcredentialsResponse(deleteResult applicationcredentials.DeleteResult,)*DeleteAccessRuleIdentityV3ApplicationcredentialsResponse {
    return &DeleteAccessRuleIdentityV3ApplicationcredentialsResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteAccessRuleIdentityV3Applicationcredentials(req *DeleteAccessRuleIdentityV3ApplicationcredentialsRequest)(*DeleteAccessRuleIdentityV3ApplicationcredentialsResponse){
    return NewDeleteAccessRuleIdentityV3ApplicationcredentialsResponse(applicationcredentials.DeleteAccessRule(oc.Client,req.UserID,req.Id, ))

}