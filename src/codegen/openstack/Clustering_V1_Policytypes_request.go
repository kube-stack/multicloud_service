package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/clustering/v1/policytypes"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListClusteringV1Policytypes
type ListClusteringV1PolicytypesRequest struct{
}

func NewListClusteringV1PolicytypesRequest()*ListClusteringV1PolicytypesRequest{
    return &ListClusteringV1PolicytypesRequest{}
}

//response struct for the ListClusteringV1Policytypes
type ListClusteringV1PolicytypesResponse struct{
    Pager pagination.Pager
}

func NewListClusteringV1PolicytypesResponse(pager pagination.Pager,)*ListClusteringV1PolicytypesResponse {
    return &ListClusteringV1PolicytypesResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListClusteringV1Policytypes(req *ListClusteringV1PolicytypesRequest)(*ListClusteringV1PolicytypesResponse){
    return NewListClusteringV1PolicytypesResponse(policytypes.List(oc.Client, ))

}
//request struct for the GetClusteringV1Policytypes
type GetClusteringV1PolicytypesRequest struct{
    PolicyTypeName string
}

func NewGetClusteringV1PolicytypesRequest()*GetClusteringV1PolicytypesRequest{
    return &GetClusteringV1PolicytypesRequest{}
}

//response struct for the GetClusteringV1Policytypes
type GetClusteringV1PolicytypesResponse struct{
    GetResult policytypes.GetResult
}

func NewGetClusteringV1PolicytypesResponse(getResult policytypes.GetResult,)*GetClusteringV1PolicytypesResponse {
    return &GetClusteringV1PolicytypesResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetClusteringV1Policytypes(req *GetClusteringV1PolicytypesRequest)(*GetClusteringV1PolicytypesResponse){
    return NewGetClusteringV1PolicytypesResponse(policytypes.Get(oc.Client,req.PolicyTypeName, ))

}