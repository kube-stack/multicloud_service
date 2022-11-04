package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/defsecrules"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListComputeV2ExtensionsDefsecrules
type ListComputeV2ExtensionsDefsecrulesRequest struct{
}

func NewListComputeV2ExtensionsDefsecrulesRequest()*ListComputeV2ExtensionsDefsecrulesRequest{
    return &ListComputeV2ExtensionsDefsecrulesRequest{}
}

//response struct for the ListComputeV2ExtensionsDefsecrules
type ListComputeV2ExtensionsDefsecrulesResponse struct{
    Pager pagination.Pager
}

func NewListComputeV2ExtensionsDefsecrulesResponse(pager pagination.Pager,)*ListComputeV2ExtensionsDefsecrulesResponse {
    return &ListComputeV2ExtensionsDefsecrulesResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListComputeV2ExtensionsDefsecrules(req *ListComputeV2ExtensionsDefsecrulesRequest)(*ListComputeV2ExtensionsDefsecrulesResponse){
    return NewListComputeV2ExtensionsDefsecrulesResponse(defsecrules.List(oc.Client, ))

}
//request struct for the CreateComputeV2ExtensionsDefsecrules
type CreateComputeV2ExtensionsDefsecrulesRequest struct{
    Opts defsecrules.CreateOpts
}

func NewCreateComputeV2ExtensionsDefsecrulesRequest()*CreateComputeV2ExtensionsDefsecrulesRequest{
    return &CreateComputeV2ExtensionsDefsecrulesRequest{}
}

//response struct for the CreateComputeV2ExtensionsDefsecrules
type CreateComputeV2ExtensionsDefsecrulesResponse struct{
    CreateResult defsecrules.CreateResult
}

func NewCreateComputeV2ExtensionsDefsecrulesResponse(createResult defsecrules.CreateResult,)*CreateComputeV2ExtensionsDefsecrulesResponse {
    return &CreateComputeV2ExtensionsDefsecrulesResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateComputeV2ExtensionsDefsecrules(req *CreateComputeV2ExtensionsDefsecrulesRequest)(*CreateComputeV2ExtensionsDefsecrulesResponse){
    return NewCreateComputeV2ExtensionsDefsecrulesResponse(defsecrules.Create(oc.Client,req.Opts, ))

}
//request struct for the GetComputeV2ExtensionsDefsecrules
type GetComputeV2ExtensionsDefsecrulesRequest struct{
    Id string
}

func NewGetComputeV2ExtensionsDefsecrulesRequest()*GetComputeV2ExtensionsDefsecrulesRequest{
    return &GetComputeV2ExtensionsDefsecrulesRequest{}
}

//response struct for the GetComputeV2ExtensionsDefsecrules
type GetComputeV2ExtensionsDefsecrulesResponse struct{
    GetResult defsecrules.GetResult
}

func NewGetComputeV2ExtensionsDefsecrulesResponse(getResult defsecrules.GetResult,)*GetComputeV2ExtensionsDefsecrulesResponse {
    return &GetComputeV2ExtensionsDefsecrulesResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetComputeV2ExtensionsDefsecrules(req *GetComputeV2ExtensionsDefsecrulesRequest)(*GetComputeV2ExtensionsDefsecrulesResponse){
    return NewGetComputeV2ExtensionsDefsecrulesResponse(defsecrules.Get(oc.Client,req.Id, ))

}
//request struct for the DeleteComputeV2ExtensionsDefsecrules
type DeleteComputeV2ExtensionsDefsecrulesRequest struct{
    Id string
}

func NewDeleteComputeV2ExtensionsDefsecrulesRequest()*DeleteComputeV2ExtensionsDefsecrulesRequest{
    return &DeleteComputeV2ExtensionsDefsecrulesRequest{}
}

//response struct for the DeleteComputeV2ExtensionsDefsecrules
type DeleteComputeV2ExtensionsDefsecrulesResponse struct{
    DeleteResult defsecrules.DeleteResult
}

func NewDeleteComputeV2ExtensionsDefsecrulesResponse(deleteResult defsecrules.DeleteResult,)*DeleteComputeV2ExtensionsDefsecrulesResponse {
    return &DeleteComputeV2ExtensionsDefsecrulesResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteComputeV2ExtensionsDefsecrules(req *DeleteComputeV2ExtensionsDefsecrulesRequest)(*DeleteComputeV2ExtensionsDefsecrulesResponse){
    return NewDeleteComputeV2ExtensionsDefsecrulesResponse(defsecrules.Delete(oc.Client,req.Id, ))

}