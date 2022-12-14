package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/identity/v3/credentials"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListIdentityV3Credentials
type ListIdentityV3CredentialsRequest struct{
    Opts credentials.ListOpts
}

func NewListIdentityV3CredentialsRequest()*ListIdentityV3CredentialsRequest{
    return &ListIdentityV3CredentialsRequest{}
}

//response struct for the ListIdentityV3Credentials
type ListIdentityV3CredentialsResponse struct{
    Pager pagination.Pager
}

func NewListIdentityV3CredentialsResponse(pager pagination.Pager,)*ListIdentityV3CredentialsResponse {
    return &ListIdentityV3CredentialsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListIdentityV3Credentials(req *ListIdentityV3CredentialsRequest)(*ListIdentityV3CredentialsResponse){
    return NewListIdentityV3CredentialsResponse(credentials.List(oc.Client,req.Opts, ))

}
//request struct for the GetIdentityV3Credentials
type GetIdentityV3CredentialsRequest struct{
    Id string
}

func NewGetIdentityV3CredentialsRequest()*GetIdentityV3CredentialsRequest{
    return &GetIdentityV3CredentialsRequest{}
}

//response struct for the GetIdentityV3Credentials
type GetIdentityV3CredentialsResponse struct{
    GetResult credentials.GetResult
}

func NewGetIdentityV3CredentialsResponse(getResult credentials.GetResult,)*GetIdentityV3CredentialsResponse {
    return &GetIdentityV3CredentialsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetIdentityV3Credentials(req *GetIdentityV3CredentialsRequest)(*GetIdentityV3CredentialsResponse){
    return NewGetIdentityV3CredentialsResponse(credentials.Get(oc.Client,req.Id, ))

}
//request struct for the CreateIdentityV3Credentials
type CreateIdentityV3CredentialsRequest struct{
    Opts credentials.CreateOpts
}

func NewCreateIdentityV3CredentialsRequest()*CreateIdentityV3CredentialsRequest{
    return &CreateIdentityV3CredentialsRequest{}
}

//response struct for the CreateIdentityV3Credentials
type CreateIdentityV3CredentialsResponse struct{
    CreateResult credentials.CreateResult
}

func NewCreateIdentityV3CredentialsResponse(createResult credentials.CreateResult,)*CreateIdentityV3CredentialsResponse {
    return &CreateIdentityV3CredentialsResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateIdentityV3Credentials(req *CreateIdentityV3CredentialsRequest)(*CreateIdentityV3CredentialsResponse){
    return NewCreateIdentityV3CredentialsResponse(credentials.Create(oc.Client,req.Opts, ))

}
//request struct for the DeleteIdentityV3Credentials
type DeleteIdentityV3CredentialsRequest struct{
    Id string
}

func NewDeleteIdentityV3CredentialsRequest()*DeleteIdentityV3CredentialsRequest{
    return &DeleteIdentityV3CredentialsRequest{}
}

//response struct for the DeleteIdentityV3Credentials
type DeleteIdentityV3CredentialsResponse struct{
    DeleteResult credentials.DeleteResult
}

func NewDeleteIdentityV3CredentialsResponse(deleteResult credentials.DeleteResult,)*DeleteIdentityV3CredentialsResponse {
    return &DeleteIdentityV3CredentialsResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteIdentityV3Credentials(req *DeleteIdentityV3CredentialsRequest)(*DeleteIdentityV3CredentialsResponse){
    return NewDeleteIdentityV3CredentialsResponse(credentials.Delete(oc.Client,req.Id, ))

}
//request struct for the UpdateIdentityV3Credentials
type UpdateIdentityV3CredentialsRequest struct{
    Id string
    Opts credentials.UpdateOpts
}

func NewUpdateIdentityV3CredentialsRequest()*UpdateIdentityV3CredentialsRequest{
    return &UpdateIdentityV3CredentialsRequest{}
}

//response struct for the UpdateIdentityV3Credentials
type UpdateIdentityV3CredentialsResponse struct{
    UpdateResult credentials.UpdateResult
}

func NewUpdateIdentityV3CredentialsResponse(updateResult credentials.UpdateResult,)*UpdateIdentityV3CredentialsResponse {
    return &UpdateIdentityV3CredentialsResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateIdentityV3Credentials(req *UpdateIdentityV3CredentialsRequest)(*UpdateIdentityV3CredentialsResponse){
    return NewUpdateIdentityV3CredentialsResponse(credentials.Update(oc.Client,req.Id,req.Opts, ))

}