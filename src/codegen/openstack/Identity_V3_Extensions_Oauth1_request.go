package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/identity/v3/extensions/oauth1"
    "github.com/gophercloud/gophercloud/openstack/identity/v3/tokens"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the CreateIdentityV3ExtensionsOauth1
type CreateIdentityV3ExtensionsOauth1Request struct{
    Opts oauth1.AuthOptions
}

func NewCreateIdentityV3ExtensionsOauth1Request()*CreateIdentityV3ExtensionsOauth1Request{
    return &CreateIdentityV3ExtensionsOauth1Request{}
}

//response struct for the CreateIdentityV3ExtensionsOauth1
type CreateIdentityV3ExtensionsOauth1Response struct{
    CreateResult tokens.CreateResult
}

func NewCreateIdentityV3ExtensionsOauth1Response(createResult tokens.CreateResult,)*CreateIdentityV3ExtensionsOauth1Response {
    return &CreateIdentityV3ExtensionsOauth1Response{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateIdentityV3ExtensionsOauth1(req *CreateIdentityV3ExtensionsOauth1Request)(*CreateIdentityV3ExtensionsOauth1Response){
    return NewCreateIdentityV3ExtensionsOauth1Response(oauth1.Create(oc.Client,req.Opts, ))

}
//request struct for the CreateConsumerIdentityV3ExtensionsOauth1
type CreateConsumerIdentityV3ExtensionsOauth1Request struct{
    Opts oauth1.CreateConsumerOpts
}

func NewCreateConsumerIdentityV3ExtensionsOauth1Request()*CreateConsumerIdentityV3ExtensionsOauth1Request{
    return &CreateConsumerIdentityV3ExtensionsOauth1Request{}
}

//response struct for the CreateConsumerIdentityV3ExtensionsOauth1
type CreateConsumerIdentityV3ExtensionsOauth1Response struct{
    CreateConsumerResult oauth1.CreateConsumerResult
}

func NewCreateConsumerIdentityV3ExtensionsOauth1Response(createConsumerResult oauth1.CreateConsumerResult,)*CreateConsumerIdentityV3ExtensionsOauth1Response {
    return &CreateConsumerIdentityV3ExtensionsOauth1Response{
            CreateConsumerResult:createConsumerResult,
    }
}

// action function
func (oc *OpenstackClient) CreateConsumerIdentityV3ExtensionsOauth1(req *CreateConsumerIdentityV3ExtensionsOauth1Request)(*CreateConsumerIdentityV3ExtensionsOauth1Response){
    return NewCreateConsumerIdentityV3ExtensionsOauth1Response(oauth1.CreateConsumer(oc.Client,req.Opts, ))

}
//request struct for the DeleteConsumerIdentityV3ExtensionsOauth1
type DeleteConsumerIdentityV3ExtensionsOauth1Request struct{
    Id string
}

func NewDeleteConsumerIdentityV3ExtensionsOauth1Request()*DeleteConsumerIdentityV3ExtensionsOauth1Request{
    return &DeleteConsumerIdentityV3ExtensionsOauth1Request{}
}

//response struct for the DeleteConsumerIdentityV3ExtensionsOauth1
type DeleteConsumerIdentityV3ExtensionsOauth1Response struct{
    DeleteConsumerResult oauth1.DeleteConsumerResult
}

func NewDeleteConsumerIdentityV3ExtensionsOauth1Response(deleteConsumerResult oauth1.DeleteConsumerResult,)*DeleteConsumerIdentityV3ExtensionsOauth1Response {
    return &DeleteConsumerIdentityV3ExtensionsOauth1Response{
            DeleteConsumerResult:deleteConsumerResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteConsumerIdentityV3ExtensionsOauth1(req *DeleteConsumerIdentityV3ExtensionsOauth1Request)(*DeleteConsumerIdentityV3ExtensionsOauth1Response){
    return NewDeleteConsumerIdentityV3ExtensionsOauth1Response(oauth1.DeleteConsumer(oc.Client,req.Id, ))

}
//request struct for the ListConsumersIdentityV3ExtensionsOauth1
type ListConsumersIdentityV3ExtensionsOauth1Request struct{
}

func NewListConsumersIdentityV3ExtensionsOauth1Request()*ListConsumersIdentityV3ExtensionsOauth1Request{
    return &ListConsumersIdentityV3ExtensionsOauth1Request{}
}

//response struct for the ListConsumersIdentityV3ExtensionsOauth1
type ListConsumersIdentityV3ExtensionsOauth1Response struct{
    Pager pagination.Pager
}

func NewListConsumersIdentityV3ExtensionsOauth1Response(pager pagination.Pager,)*ListConsumersIdentityV3ExtensionsOauth1Response {
    return &ListConsumersIdentityV3ExtensionsOauth1Response{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListConsumersIdentityV3ExtensionsOauth1(req *ListConsumersIdentityV3ExtensionsOauth1Request)(*ListConsumersIdentityV3ExtensionsOauth1Response){
    return NewListConsumersIdentityV3ExtensionsOauth1Response(oauth1.ListConsumers(oc.Client, ))

}
//request struct for the GetConsumerIdentityV3ExtensionsOauth1
type GetConsumerIdentityV3ExtensionsOauth1Request struct{
    Id string
}

func NewGetConsumerIdentityV3ExtensionsOauth1Request()*GetConsumerIdentityV3ExtensionsOauth1Request{
    return &GetConsumerIdentityV3ExtensionsOauth1Request{}
}

//response struct for the GetConsumerIdentityV3ExtensionsOauth1
type GetConsumerIdentityV3ExtensionsOauth1Response struct{
    GetConsumerResult oauth1.GetConsumerResult
}

func NewGetConsumerIdentityV3ExtensionsOauth1Response(getConsumerResult oauth1.GetConsumerResult,)*GetConsumerIdentityV3ExtensionsOauth1Response {
    return &GetConsumerIdentityV3ExtensionsOauth1Response{
            GetConsumerResult:getConsumerResult,
    }
}

// action function
func (oc *OpenstackClient) GetConsumerIdentityV3ExtensionsOauth1(req *GetConsumerIdentityV3ExtensionsOauth1Request)(*GetConsumerIdentityV3ExtensionsOauth1Response){
    return NewGetConsumerIdentityV3ExtensionsOauth1Response(oauth1.GetConsumer(oc.Client,req.Id, ))

}
//request struct for the UpdateConsumerIdentityV3ExtensionsOauth1
type UpdateConsumerIdentityV3ExtensionsOauth1Request struct{
    Id string
    Opts oauth1.UpdateConsumerOpts
}

func NewUpdateConsumerIdentityV3ExtensionsOauth1Request()*UpdateConsumerIdentityV3ExtensionsOauth1Request{
    return &UpdateConsumerIdentityV3ExtensionsOauth1Request{}
}

//response struct for the UpdateConsumerIdentityV3ExtensionsOauth1
type UpdateConsumerIdentityV3ExtensionsOauth1Response struct{
    UpdateConsumerResult oauth1.UpdateConsumerResult
}

func NewUpdateConsumerIdentityV3ExtensionsOauth1Response(updateConsumerResult oauth1.UpdateConsumerResult,)*UpdateConsumerIdentityV3ExtensionsOauth1Response {
    return &UpdateConsumerIdentityV3ExtensionsOauth1Response{
            UpdateConsumerResult:updateConsumerResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateConsumerIdentityV3ExtensionsOauth1(req *UpdateConsumerIdentityV3ExtensionsOauth1Request)(*UpdateConsumerIdentityV3ExtensionsOauth1Response){
    return NewUpdateConsumerIdentityV3ExtensionsOauth1Response(oauth1.UpdateConsumer(oc.Client,req.Id,req.Opts, ))

}
//request struct for the RequestTokenIdentityV3ExtensionsOauth1
type RequestTokenIdentityV3ExtensionsOauth1Request struct{
    Opts oauth1.RequestTokenOpts
}

func NewRequestTokenIdentityV3ExtensionsOauth1Request()*RequestTokenIdentityV3ExtensionsOauth1Request{
    return &RequestTokenIdentityV3ExtensionsOauth1Request{}
}

//response struct for the RequestTokenIdentityV3ExtensionsOauth1
type RequestTokenIdentityV3ExtensionsOauth1Response struct{
    TokenResult oauth1.TokenResult
}

func NewRequestTokenIdentityV3ExtensionsOauth1Response(tokenResult oauth1.TokenResult,)*RequestTokenIdentityV3ExtensionsOauth1Response {
    return &RequestTokenIdentityV3ExtensionsOauth1Response{
            TokenResult:tokenResult,
    }
}

// action function
func (oc *OpenstackClient) RequestTokenIdentityV3ExtensionsOauth1(req *RequestTokenIdentityV3ExtensionsOauth1Request)(*RequestTokenIdentityV3ExtensionsOauth1Response){
    return NewRequestTokenIdentityV3ExtensionsOauth1Response(oauth1.RequestToken(oc.Client,req.Opts, ))

}
//request struct for the AuthorizeTokenIdentityV3ExtensionsOauth1
type AuthorizeTokenIdentityV3ExtensionsOauth1Request struct{
    Id string
    Opts oauth1.AuthorizeTokenOpts
}

func NewAuthorizeTokenIdentityV3ExtensionsOauth1Request()*AuthorizeTokenIdentityV3ExtensionsOauth1Request{
    return &AuthorizeTokenIdentityV3ExtensionsOauth1Request{}
}

//response struct for the AuthorizeTokenIdentityV3ExtensionsOauth1
type AuthorizeTokenIdentityV3ExtensionsOauth1Response struct{
    AuthorizeTokenResult oauth1.AuthorizeTokenResult
}

func NewAuthorizeTokenIdentityV3ExtensionsOauth1Response(authorizeTokenResult oauth1.AuthorizeTokenResult,)*AuthorizeTokenIdentityV3ExtensionsOauth1Response {
    return &AuthorizeTokenIdentityV3ExtensionsOauth1Response{
            AuthorizeTokenResult:authorizeTokenResult,
    }
}

// action function
func (oc *OpenstackClient) AuthorizeTokenIdentityV3ExtensionsOauth1(req *AuthorizeTokenIdentityV3ExtensionsOauth1Request)(*AuthorizeTokenIdentityV3ExtensionsOauth1Response){
    return NewAuthorizeTokenIdentityV3ExtensionsOauth1Response(oauth1.AuthorizeToken(oc.Client,req.Id,req.Opts, ))

}
//request struct for the CreateAccessTokenIdentityV3ExtensionsOauth1
type CreateAccessTokenIdentityV3ExtensionsOauth1Request struct{
    Opts oauth1.CreateAccessTokenOpts
}

func NewCreateAccessTokenIdentityV3ExtensionsOauth1Request()*CreateAccessTokenIdentityV3ExtensionsOauth1Request{
    return &CreateAccessTokenIdentityV3ExtensionsOauth1Request{}
}

//response struct for the CreateAccessTokenIdentityV3ExtensionsOauth1
type CreateAccessTokenIdentityV3ExtensionsOauth1Response struct{
    TokenResult oauth1.TokenResult
}

func NewCreateAccessTokenIdentityV3ExtensionsOauth1Response(tokenResult oauth1.TokenResult,)*CreateAccessTokenIdentityV3ExtensionsOauth1Response {
    return &CreateAccessTokenIdentityV3ExtensionsOauth1Response{
            TokenResult:tokenResult,
    }
}

// action function
func (oc *OpenstackClient) CreateAccessTokenIdentityV3ExtensionsOauth1(req *CreateAccessTokenIdentityV3ExtensionsOauth1Request)(*CreateAccessTokenIdentityV3ExtensionsOauth1Response){
    return NewCreateAccessTokenIdentityV3ExtensionsOauth1Response(oauth1.CreateAccessToken(oc.Client,req.Opts, ))

}
//request struct for the GetAccessTokenIdentityV3ExtensionsOauth1
type GetAccessTokenIdentityV3ExtensionsOauth1Request struct{
    UserID string
    Id string
}

func NewGetAccessTokenIdentityV3ExtensionsOauth1Request()*GetAccessTokenIdentityV3ExtensionsOauth1Request{
    return &GetAccessTokenIdentityV3ExtensionsOauth1Request{}
}

//response struct for the GetAccessTokenIdentityV3ExtensionsOauth1
type GetAccessTokenIdentityV3ExtensionsOauth1Response struct{
    GetAccessTokenResult oauth1.GetAccessTokenResult
}

func NewGetAccessTokenIdentityV3ExtensionsOauth1Response(getAccessTokenResult oauth1.GetAccessTokenResult,)*GetAccessTokenIdentityV3ExtensionsOauth1Response {
    return &GetAccessTokenIdentityV3ExtensionsOauth1Response{
            GetAccessTokenResult:getAccessTokenResult,
    }
}

// action function
func (oc *OpenstackClient) GetAccessTokenIdentityV3ExtensionsOauth1(req *GetAccessTokenIdentityV3ExtensionsOauth1Request)(*GetAccessTokenIdentityV3ExtensionsOauth1Response){
    return NewGetAccessTokenIdentityV3ExtensionsOauth1Response(oauth1.GetAccessToken(oc.Client,req.UserID,req.Id, ))

}
//request struct for the RevokeAccessTokenIdentityV3ExtensionsOauth1
type RevokeAccessTokenIdentityV3ExtensionsOauth1Request struct{
    UserID string
    Id string
}

func NewRevokeAccessTokenIdentityV3ExtensionsOauth1Request()*RevokeAccessTokenIdentityV3ExtensionsOauth1Request{
    return &RevokeAccessTokenIdentityV3ExtensionsOauth1Request{}
}

//response struct for the RevokeAccessTokenIdentityV3ExtensionsOauth1
type RevokeAccessTokenIdentityV3ExtensionsOauth1Response struct{
    RevokeAccessTokenResult oauth1.RevokeAccessTokenResult
}

func NewRevokeAccessTokenIdentityV3ExtensionsOauth1Response(revokeAccessTokenResult oauth1.RevokeAccessTokenResult,)*RevokeAccessTokenIdentityV3ExtensionsOauth1Response {
    return &RevokeAccessTokenIdentityV3ExtensionsOauth1Response{
            RevokeAccessTokenResult:revokeAccessTokenResult,
    }
}

// action function
func (oc *OpenstackClient) RevokeAccessTokenIdentityV3ExtensionsOauth1(req *RevokeAccessTokenIdentityV3ExtensionsOauth1Request)(*RevokeAccessTokenIdentityV3ExtensionsOauth1Response){
    return NewRevokeAccessTokenIdentityV3ExtensionsOauth1Response(oauth1.RevokeAccessToken(oc.Client,req.UserID,req.Id, ))

}
//request struct for the ListAccessTokensIdentityV3ExtensionsOauth1
type ListAccessTokensIdentityV3ExtensionsOauth1Request struct{
    UserID string
}

func NewListAccessTokensIdentityV3ExtensionsOauth1Request()*ListAccessTokensIdentityV3ExtensionsOauth1Request{
    return &ListAccessTokensIdentityV3ExtensionsOauth1Request{}
}

//response struct for the ListAccessTokensIdentityV3ExtensionsOauth1
type ListAccessTokensIdentityV3ExtensionsOauth1Response struct{
    Pager pagination.Pager
}

func NewListAccessTokensIdentityV3ExtensionsOauth1Response(pager pagination.Pager,)*ListAccessTokensIdentityV3ExtensionsOauth1Response {
    return &ListAccessTokensIdentityV3ExtensionsOauth1Response{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListAccessTokensIdentityV3ExtensionsOauth1(req *ListAccessTokensIdentityV3ExtensionsOauth1Request)(*ListAccessTokensIdentityV3ExtensionsOauth1Response){
    return NewListAccessTokensIdentityV3ExtensionsOauth1Response(oauth1.ListAccessTokens(oc.Client,req.UserID, ))

}
//request struct for the ListAccessTokenRolesIdentityV3ExtensionsOauth1
type ListAccessTokenRolesIdentityV3ExtensionsOauth1Request struct{
    UserID string
    Id string
}

func NewListAccessTokenRolesIdentityV3ExtensionsOauth1Request()*ListAccessTokenRolesIdentityV3ExtensionsOauth1Request{
    return &ListAccessTokenRolesIdentityV3ExtensionsOauth1Request{}
}

//response struct for the ListAccessTokenRolesIdentityV3ExtensionsOauth1
type ListAccessTokenRolesIdentityV3ExtensionsOauth1Response struct{
    Pager pagination.Pager
}

func NewListAccessTokenRolesIdentityV3ExtensionsOauth1Response(pager pagination.Pager,)*ListAccessTokenRolesIdentityV3ExtensionsOauth1Response {
    return &ListAccessTokenRolesIdentityV3ExtensionsOauth1Response{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListAccessTokenRolesIdentityV3ExtensionsOauth1(req *ListAccessTokenRolesIdentityV3ExtensionsOauth1Request)(*ListAccessTokenRolesIdentityV3ExtensionsOauth1Response){
    return NewListAccessTokenRolesIdentityV3ExtensionsOauth1Response(oauth1.ListAccessTokenRoles(oc.Client,req.UserID,req.Id, ))

}
//request struct for the GetAccessTokenRoleIdentityV3ExtensionsOauth1
type GetAccessTokenRoleIdentityV3ExtensionsOauth1Request struct{
    UserID string
    Id string
    RoleID string
}

func NewGetAccessTokenRoleIdentityV3ExtensionsOauth1Request()*GetAccessTokenRoleIdentityV3ExtensionsOauth1Request{
    return &GetAccessTokenRoleIdentityV3ExtensionsOauth1Request{}
}

//response struct for the GetAccessTokenRoleIdentityV3ExtensionsOauth1
type GetAccessTokenRoleIdentityV3ExtensionsOauth1Response struct{
    GetAccessTokenRoleResult oauth1.GetAccessTokenRoleResult
}

func NewGetAccessTokenRoleIdentityV3ExtensionsOauth1Response(getAccessTokenRoleResult oauth1.GetAccessTokenRoleResult,)*GetAccessTokenRoleIdentityV3ExtensionsOauth1Response {
    return &GetAccessTokenRoleIdentityV3ExtensionsOauth1Response{
            GetAccessTokenRoleResult:getAccessTokenRoleResult,
    }
}

// action function
func (oc *OpenstackClient) GetAccessTokenRoleIdentityV3ExtensionsOauth1(req *GetAccessTokenRoleIdentityV3ExtensionsOauth1Request)(*GetAccessTokenRoleIdentityV3ExtensionsOauth1Response){
    return NewGetAccessTokenRoleIdentityV3ExtensionsOauth1Response(oauth1.GetAccessTokenRole(oc.Client,req.UserID,req.Id,req.RoleID, ))

}