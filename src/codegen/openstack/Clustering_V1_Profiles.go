package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/clustering/v1/profiles"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the CreateClusteringV1Profiles
type CreateClusteringV1ProfilesRequest struct{
    Opts profiles.CreateOpts
}

func NewCreateClusteringV1ProfilesRequest()*CreateClusteringV1ProfilesRequest{
    return &CreateClusteringV1ProfilesRequest{}
}

//response struct for the CreateClusteringV1Profiles
type CreateClusteringV1ProfilesResponse struct{
    CreateResult profiles.CreateResult
}

func NewCreateClusteringV1ProfilesResponse(createResult profiles.CreateResult,)*CreateClusteringV1ProfilesResponse {
    return &CreateClusteringV1ProfilesResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateClusteringV1Profiles(req *CreateClusteringV1ProfilesRequest)(*CreateClusteringV1ProfilesResponse){
    return NewCreateClusteringV1ProfilesResponse(profiles.Create(oc.Client,req.Opts, ))

}
//request struct for the GetClusteringV1Profiles
type GetClusteringV1ProfilesRequest struct{
    Id string
}

func NewGetClusteringV1ProfilesRequest()*GetClusteringV1ProfilesRequest{
    return &GetClusteringV1ProfilesRequest{}
}

//response struct for the GetClusteringV1Profiles
type GetClusteringV1ProfilesResponse struct{
    GetResult profiles.GetResult
}

func NewGetClusteringV1ProfilesResponse(getResult profiles.GetResult,)*GetClusteringV1ProfilesResponse {
    return &GetClusteringV1ProfilesResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetClusteringV1Profiles(req *GetClusteringV1ProfilesRequest)(*GetClusteringV1ProfilesResponse){
    return NewGetClusteringV1ProfilesResponse(profiles.Get(oc.Client,req.Id, ))

}
//request struct for the ListClusteringV1Profiles
type ListClusteringV1ProfilesRequest struct{
    Opts profiles.ListOpts
}

func NewListClusteringV1ProfilesRequest()*ListClusteringV1ProfilesRequest{
    return &ListClusteringV1ProfilesRequest{}
}

//response struct for the ListClusteringV1Profiles
type ListClusteringV1ProfilesResponse struct{
    Pager pagination.Pager
}

func NewListClusteringV1ProfilesResponse(pager pagination.Pager,)*ListClusteringV1ProfilesResponse {
    return &ListClusteringV1ProfilesResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListClusteringV1Profiles(req *ListClusteringV1ProfilesRequest)(*ListClusteringV1ProfilesResponse){
    return NewListClusteringV1ProfilesResponse(profiles.List(oc.Client,req.Opts, ))

}
//request struct for the UpdateClusteringV1Profiles
type UpdateClusteringV1ProfilesRequest struct{
    Id string
    Opts profiles.UpdateOpts
}

func NewUpdateClusteringV1ProfilesRequest()*UpdateClusteringV1ProfilesRequest{
    return &UpdateClusteringV1ProfilesRequest{}
}

//response struct for the UpdateClusteringV1Profiles
type UpdateClusteringV1ProfilesResponse struct{
    UpdateResult profiles.UpdateResult
}

func NewUpdateClusteringV1ProfilesResponse(updateResult profiles.UpdateResult,)*UpdateClusteringV1ProfilesResponse {
    return &UpdateClusteringV1ProfilesResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateClusteringV1Profiles(req *UpdateClusteringV1ProfilesRequest)(*UpdateClusteringV1ProfilesResponse){
    return NewUpdateClusteringV1ProfilesResponse(profiles.Update(oc.Client,req.Id,req.Opts, ))

}
//request struct for the DeleteClusteringV1Profiles
type DeleteClusteringV1ProfilesRequest struct{
    Id string
}

func NewDeleteClusteringV1ProfilesRequest()*DeleteClusteringV1ProfilesRequest{
    return &DeleteClusteringV1ProfilesRequest{}
}

//response struct for the DeleteClusteringV1Profiles
type DeleteClusteringV1ProfilesResponse struct{
    DeleteResult profiles.DeleteResult
}

func NewDeleteClusteringV1ProfilesResponse(deleteResult profiles.DeleteResult,)*DeleteClusteringV1ProfilesResponse {
    return &DeleteClusteringV1ProfilesResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteClusteringV1Profiles(req *DeleteClusteringV1ProfilesRequest)(*DeleteClusteringV1ProfilesResponse){
    return NewDeleteClusteringV1ProfilesResponse(profiles.Delete(oc.Client,req.Id, ))

}
//request struct for the ValidateClusteringV1Profiles
type ValidateClusteringV1ProfilesRequest struct{
    Opts profiles.ValidateOpts
}

func NewValidateClusteringV1ProfilesRequest()*ValidateClusteringV1ProfilesRequest{
    return &ValidateClusteringV1ProfilesRequest{}
}

//response struct for the ValidateClusteringV1Profiles
type ValidateClusteringV1ProfilesResponse struct{
    ValidateResult profiles.ValidateResult
}

func NewValidateClusteringV1ProfilesResponse(validateResult profiles.ValidateResult,)*ValidateClusteringV1ProfilesResponse {
    return &ValidateClusteringV1ProfilesResponse{
            ValidateResult:validateResult,
    }
}

// action function
func (oc *OpenstackClient) ValidateClusteringV1Profiles(req *ValidateClusteringV1ProfilesRequest)(*ValidateClusteringV1ProfilesResponse){
    return NewValidateClusteringV1ProfilesResponse(profiles.Validate(oc.Client,req.Opts, ))

}