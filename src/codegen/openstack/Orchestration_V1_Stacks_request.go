package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/orchestration/v1/stacks"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the CreateOrchestrationV1Stacks
type CreateOrchestrationV1StacksRequest struct{
    Opts stacks.CreateOpts
}

func NewCreateOrchestrationV1StacksRequest()*CreateOrchestrationV1StacksRequest{
    return &CreateOrchestrationV1StacksRequest{}
}

//response struct for the CreateOrchestrationV1Stacks
type CreateOrchestrationV1StacksResponse struct{
    CreateResult stacks.CreateResult
}

func NewCreateOrchestrationV1StacksResponse(createResult stacks.CreateResult,)*CreateOrchestrationV1StacksResponse {
    return &CreateOrchestrationV1StacksResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateOrchestrationV1Stacks(req *CreateOrchestrationV1StacksRequest)(*CreateOrchestrationV1StacksResponse){
    return NewCreateOrchestrationV1StacksResponse(stacks.Create(oc.Client,req.Opts, ))

}
//request struct for the AdoptOrchestrationV1Stacks
type AdoptOrchestrationV1StacksRequest struct{
    Opts stacks.AdoptOpts
}

func NewAdoptOrchestrationV1StacksRequest()*AdoptOrchestrationV1StacksRequest{
    return &AdoptOrchestrationV1StacksRequest{}
}

//response struct for the AdoptOrchestrationV1Stacks
type AdoptOrchestrationV1StacksResponse struct{
    AdoptResult stacks.AdoptResult
}

func NewAdoptOrchestrationV1StacksResponse(adoptResult stacks.AdoptResult,)*AdoptOrchestrationV1StacksResponse {
    return &AdoptOrchestrationV1StacksResponse{
            AdoptResult:adoptResult,
    }
}

// action function
func (oc *OpenstackClient) AdoptOrchestrationV1Stacks(req *AdoptOrchestrationV1StacksRequest)(*AdoptOrchestrationV1StacksResponse){
    return NewAdoptOrchestrationV1StacksResponse(stacks.Adopt(oc.Client,req.Opts, ))

}
//request struct for the ListOrchestrationV1Stacks
type ListOrchestrationV1StacksRequest struct{
    Opts stacks.ListOpts
}

func NewListOrchestrationV1StacksRequest()*ListOrchestrationV1StacksRequest{
    return &ListOrchestrationV1StacksRequest{}
}

//response struct for the ListOrchestrationV1Stacks
type ListOrchestrationV1StacksResponse struct{
    Pager pagination.Pager
}

func NewListOrchestrationV1StacksResponse(pager pagination.Pager,)*ListOrchestrationV1StacksResponse {
    return &ListOrchestrationV1StacksResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListOrchestrationV1Stacks(req *ListOrchestrationV1StacksRequest)(*ListOrchestrationV1StacksResponse){
    return NewListOrchestrationV1StacksResponse(stacks.List(oc.Client,req.Opts, ))

}
//request struct for the GetOrchestrationV1Stacks
type GetOrchestrationV1StacksRequest struct{
    StackName string
    StackID string
}

func NewGetOrchestrationV1StacksRequest()*GetOrchestrationV1StacksRequest{
    return &GetOrchestrationV1StacksRequest{}
}

//response struct for the GetOrchestrationV1Stacks
type GetOrchestrationV1StacksResponse struct{
    GetResult stacks.GetResult
}

func NewGetOrchestrationV1StacksResponse(getResult stacks.GetResult,)*GetOrchestrationV1StacksResponse {
    return &GetOrchestrationV1StacksResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetOrchestrationV1Stacks(req *GetOrchestrationV1StacksRequest)(*GetOrchestrationV1StacksResponse){
    return NewGetOrchestrationV1StacksResponse(stacks.Get(oc.Client,req.StackName,req.StackID, ))

}
//request struct for the FindOrchestrationV1Stacks
type FindOrchestrationV1StacksRequest struct{
    StackIdentity string
}

func NewFindOrchestrationV1StacksRequest()*FindOrchestrationV1StacksRequest{
    return &FindOrchestrationV1StacksRequest{}
}

//response struct for the FindOrchestrationV1Stacks
type FindOrchestrationV1StacksResponse struct{
    GetResult stacks.GetResult
}

func NewFindOrchestrationV1StacksResponse(getResult stacks.GetResult,)*FindOrchestrationV1StacksResponse {
    return &FindOrchestrationV1StacksResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) FindOrchestrationV1Stacks(req *FindOrchestrationV1StacksRequest)(*FindOrchestrationV1StacksResponse){
    return NewFindOrchestrationV1StacksResponse(stacks.Find(oc.Client,req.StackIdentity, ))

}
//request struct for the UpdateOrchestrationV1Stacks
type UpdateOrchestrationV1StacksRequest struct{
    StackName string
    StackID string
    Opts stacks.UpdateOpts
}

func NewUpdateOrchestrationV1StacksRequest()*UpdateOrchestrationV1StacksRequest{
    return &UpdateOrchestrationV1StacksRequest{}
}

//response struct for the UpdateOrchestrationV1Stacks
type UpdateOrchestrationV1StacksResponse struct{
    UpdateResult stacks.UpdateResult
}

func NewUpdateOrchestrationV1StacksResponse(updateResult stacks.UpdateResult,)*UpdateOrchestrationV1StacksResponse {
    return &UpdateOrchestrationV1StacksResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateOrchestrationV1Stacks(req *UpdateOrchestrationV1StacksRequest)(*UpdateOrchestrationV1StacksResponse){
    return NewUpdateOrchestrationV1StacksResponse(stacks.Update(oc.Client,req.StackName,req.StackID,req.Opts, ))

}
//request struct for the UpdatePatchOrchestrationV1Stacks
type UpdatePatchOrchestrationV1StacksRequest struct{
    StackName string
    StackID string
    Opts stacks.UpdateOpts
}

func NewUpdatePatchOrchestrationV1StacksRequest()*UpdatePatchOrchestrationV1StacksRequest{
    return &UpdatePatchOrchestrationV1StacksRequest{}
}

//response struct for the UpdatePatchOrchestrationV1Stacks
type UpdatePatchOrchestrationV1StacksResponse struct{
    UpdateResult stacks.UpdateResult
}

func NewUpdatePatchOrchestrationV1StacksResponse(updateResult stacks.UpdateResult,)*UpdatePatchOrchestrationV1StacksResponse {
    return &UpdatePatchOrchestrationV1StacksResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdatePatchOrchestrationV1Stacks(req *UpdatePatchOrchestrationV1StacksRequest)(*UpdatePatchOrchestrationV1StacksResponse){
    return NewUpdatePatchOrchestrationV1StacksResponse(stacks.UpdatePatch(oc.Client,req.StackName,req.StackID,req.Opts, ))

}
//request struct for the DeleteOrchestrationV1Stacks
type DeleteOrchestrationV1StacksRequest struct{
    StackName string
    StackID string
}

func NewDeleteOrchestrationV1StacksRequest()*DeleteOrchestrationV1StacksRequest{
    return &DeleteOrchestrationV1StacksRequest{}
}

//response struct for the DeleteOrchestrationV1Stacks
type DeleteOrchestrationV1StacksResponse struct{
    DeleteResult stacks.DeleteResult
}

func NewDeleteOrchestrationV1StacksResponse(deleteResult stacks.DeleteResult,)*DeleteOrchestrationV1StacksResponse {
    return &DeleteOrchestrationV1StacksResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteOrchestrationV1Stacks(req *DeleteOrchestrationV1StacksRequest)(*DeleteOrchestrationV1StacksResponse){
    return NewDeleteOrchestrationV1StacksResponse(stacks.Delete(oc.Client,req.StackName,req.StackID, ))

}
//request struct for the PreviewOrchestrationV1Stacks
type PreviewOrchestrationV1StacksRequest struct{
    Opts stacks.PreviewOpts
}

func NewPreviewOrchestrationV1StacksRequest()*PreviewOrchestrationV1StacksRequest{
    return &PreviewOrchestrationV1StacksRequest{}
}

//response struct for the PreviewOrchestrationV1Stacks
type PreviewOrchestrationV1StacksResponse struct{
    PreviewResult stacks.PreviewResult
}

func NewPreviewOrchestrationV1StacksResponse(previewResult stacks.PreviewResult,)*PreviewOrchestrationV1StacksResponse {
    return &PreviewOrchestrationV1StacksResponse{
            PreviewResult:previewResult,
    }
}

// action function
func (oc *OpenstackClient) PreviewOrchestrationV1Stacks(req *PreviewOrchestrationV1StacksRequest)(*PreviewOrchestrationV1StacksResponse){
    return NewPreviewOrchestrationV1StacksResponse(stacks.Preview(oc.Client,req.Opts, ))

}
//request struct for the AbandonOrchestrationV1Stacks
type AbandonOrchestrationV1StacksRequest struct{
    StackName string
    StackID string
}

func NewAbandonOrchestrationV1StacksRequest()*AbandonOrchestrationV1StacksRequest{
    return &AbandonOrchestrationV1StacksRequest{}
}

//response struct for the AbandonOrchestrationV1Stacks
type AbandonOrchestrationV1StacksResponse struct{
    AbandonResult stacks.AbandonResult
}

func NewAbandonOrchestrationV1StacksResponse(abandonResult stacks.AbandonResult,)*AbandonOrchestrationV1StacksResponse {
    return &AbandonOrchestrationV1StacksResponse{
            AbandonResult:abandonResult,
    }
}

// action function
func (oc *OpenstackClient) AbandonOrchestrationV1Stacks(req *AbandonOrchestrationV1StacksRequest)(*AbandonOrchestrationV1StacksResponse){
    return NewAbandonOrchestrationV1StacksResponse(stacks.Abandon(oc.Client,req.StackName,req.StackID, ))

}