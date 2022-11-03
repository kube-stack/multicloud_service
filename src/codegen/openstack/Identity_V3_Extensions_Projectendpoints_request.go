package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/identity/v3/extensions/projectendpoints"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the CreateIdentityV3ExtensionsProjectendpoints
type CreateIdentityV3ExtensionsProjectendpointsRequest struct{
    ProjectID string
    EndpointID string
}

func NewCreateIdentityV3ExtensionsProjectendpointsRequest()*CreateIdentityV3ExtensionsProjectendpointsRequest{
    return &CreateIdentityV3ExtensionsProjectendpointsRequest{}
}

//response struct for the CreateIdentityV3ExtensionsProjectendpoints
type CreateIdentityV3ExtensionsProjectendpointsResponse struct{
    CreateResult projectendpoints.CreateResult
}

func NewCreateIdentityV3ExtensionsProjectendpointsResponse(createResult projectendpoints.CreateResult,)*CreateIdentityV3ExtensionsProjectendpointsResponse {
    return &CreateIdentityV3ExtensionsProjectendpointsResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateIdentityV3ExtensionsProjectendpoints(req *CreateIdentityV3ExtensionsProjectendpointsRequest)(*CreateIdentityV3ExtensionsProjectendpointsResponse){
    return NewCreateIdentityV3ExtensionsProjectendpointsResponse(projectendpoints.Create(oc.Client,req.ProjectID,req.EndpointID, ))

}
//request struct for the ListIdentityV3ExtensionsProjectendpoints
type ListIdentityV3ExtensionsProjectendpointsRequest struct{
    ProjectID string
}

func NewListIdentityV3ExtensionsProjectendpointsRequest()*ListIdentityV3ExtensionsProjectendpointsRequest{
    return &ListIdentityV3ExtensionsProjectendpointsRequest{}
}

//response struct for the ListIdentityV3ExtensionsProjectendpoints
type ListIdentityV3ExtensionsProjectendpointsResponse struct{
    Pager pagination.Pager
}

func NewListIdentityV3ExtensionsProjectendpointsResponse(pager pagination.Pager,)*ListIdentityV3ExtensionsProjectendpointsResponse {
    return &ListIdentityV3ExtensionsProjectendpointsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListIdentityV3ExtensionsProjectendpoints(req *ListIdentityV3ExtensionsProjectendpointsRequest)(*ListIdentityV3ExtensionsProjectendpointsResponse){
    return NewListIdentityV3ExtensionsProjectendpointsResponse(projectendpoints.List(oc.Client,req.ProjectID, ))

}
//request struct for the DeleteIdentityV3ExtensionsProjectendpoints
type DeleteIdentityV3ExtensionsProjectendpointsRequest struct{
    ProjectID string
    EndpointID string
}

func NewDeleteIdentityV3ExtensionsProjectendpointsRequest()*DeleteIdentityV3ExtensionsProjectendpointsRequest{
    return &DeleteIdentityV3ExtensionsProjectendpointsRequest{}
}

//response struct for the DeleteIdentityV3ExtensionsProjectendpoints
type DeleteIdentityV3ExtensionsProjectendpointsResponse struct{
    DeleteResult projectendpoints.DeleteResult
}

func NewDeleteIdentityV3ExtensionsProjectendpointsResponse(deleteResult projectendpoints.DeleteResult,)*DeleteIdentityV3ExtensionsProjectendpointsResponse {
    return &DeleteIdentityV3ExtensionsProjectendpointsResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteIdentityV3ExtensionsProjectendpoints(req *DeleteIdentityV3ExtensionsProjectendpointsRequest)(*DeleteIdentityV3ExtensionsProjectendpointsResponse){
    return NewDeleteIdentityV3ExtensionsProjectendpointsResponse(projectendpoints.Delete(oc.Client,req.ProjectID,req.EndpointID, ))

}