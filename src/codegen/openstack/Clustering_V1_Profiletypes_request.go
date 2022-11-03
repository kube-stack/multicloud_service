package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/clustering/v1/profiletypes"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the GetClusteringV1Profiletypes
type GetClusteringV1ProfiletypesRequest struct{
    Id string
}

func NewGetClusteringV1ProfiletypesRequest()*GetClusteringV1ProfiletypesRequest{
    return &GetClusteringV1ProfiletypesRequest{}
}

//response struct for the GetClusteringV1Profiletypes
type GetClusteringV1ProfiletypesResponse struct{
    GetResult profiletypes.GetResult
}

func NewGetClusteringV1ProfiletypesResponse(getResult profiletypes.GetResult,)*GetClusteringV1ProfiletypesResponse {
    return &GetClusteringV1ProfiletypesResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetClusteringV1Profiletypes(req *GetClusteringV1ProfiletypesRequest)(*GetClusteringV1ProfiletypesResponse){
    return NewGetClusteringV1ProfiletypesResponse(profiletypes.Get(oc.Client,req.Id, ))

}
//request struct for the ListOpsClusteringV1Profiletypes
type ListOpsClusteringV1ProfiletypesRequest struct{
    Id string
}

func NewListOpsClusteringV1ProfiletypesRequest()*ListOpsClusteringV1ProfiletypesRequest{
    return &ListOpsClusteringV1ProfiletypesRequest{}
}

//response struct for the ListOpsClusteringV1Profiletypes
type ListOpsClusteringV1ProfiletypesResponse struct{
    Pager pagination.Pager
}

func NewListOpsClusteringV1ProfiletypesResponse(pager pagination.Pager,)*ListOpsClusteringV1ProfiletypesResponse {
    return &ListOpsClusteringV1ProfiletypesResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListOpsClusteringV1Profiletypes(req *ListOpsClusteringV1ProfiletypesRequest)(*ListOpsClusteringV1ProfiletypesResponse){
    return NewListOpsClusteringV1ProfiletypesResponse(profiletypes.ListOps(oc.Client,req.Id, ))

}