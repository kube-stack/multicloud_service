package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/clustering/v1/actions"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListClusteringV1Actions
type ListClusteringV1ActionsRequest struct{
    Opts actions.ListOpts
}

func NewListClusteringV1ActionsRequest()*ListClusteringV1ActionsRequest{
    return &ListClusteringV1ActionsRequest{}
}

//response struct for the ListClusteringV1Actions
type ListClusteringV1ActionsResponse struct{
    Pager pagination.Pager
}

func NewListClusteringV1ActionsResponse(pager pagination.Pager,)*ListClusteringV1ActionsResponse {
    return &ListClusteringV1ActionsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListClusteringV1Actions(req *ListClusteringV1ActionsRequest)(*ListClusteringV1ActionsResponse){
    return NewListClusteringV1ActionsResponse(actions.List(oc.Client,req.Opts, ))

}
//request struct for the GetClusteringV1Actions
type GetClusteringV1ActionsRequest struct{
    Id string
}

func NewGetClusteringV1ActionsRequest()*GetClusteringV1ActionsRequest{
    return &GetClusteringV1ActionsRequest{}
}

//response struct for the GetClusteringV1Actions
type GetClusteringV1ActionsResponse struct{
    GetResult actions.GetResult
}

func NewGetClusteringV1ActionsResponse(getResult actions.GetResult,)*GetClusteringV1ActionsResponse {
    return &GetClusteringV1ActionsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetClusteringV1Actions(req *GetClusteringV1ActionsRequest)(*GetClusteringV1ActionsResponse){
    return NewGetClusteringV1ActionsResponse(actions.Get(oc.Client,req.Id, ))

}