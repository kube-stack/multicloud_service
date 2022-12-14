package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/blockstorage/v3/attachments"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the CreateBlockstorageV3Attachments
type CreateBlockstorageV3AttachmentsRequest struct{
    Opts attachments.CreateOpts
}

func NewCreateBlockstorageV3AttachmentsRequest()*CreateBlockstorageV3AttachmentsRequest{
    return &CreateBlockstorageV3AttachmentsRequest{}
}

//response struct for the CreateBlockstorageV3Attachments
type CreateBlockstorageV3AttachmentsResponse struct{
    CreateResult attachments.CreateResult
}

func NewCreateBlockstorageV3AttachmentsResponse(createResult attachments.CreateResult,)*CreateBlockstorageV3AttachmentsResponse {
    return &CreateBlockstorageV3AttachmentsResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateBlockstorageV3Attachments(req *CreateBlockstorageV3AttachmentsRequest)(*CreateBlockstorageV3AttachmentsResponse){
    return NewCreateBlockstorageV3AttachmentsResponse(attachments.Create(oc.Client,req.Opts, ))

}
//request struct for the DeleteBlockstorageV3Attachments
type DeleteBlockstorageV3AttachmentsRequest struct{
    Id string
}

func NewDeleteBlockstorageV3AttachmentsRequest()*DeleteBlockstorageV3AttachmentsRequest{
    return &DeleteBlockstorageV3AttachmentsRequest{}
}

//response struct for the DeleteBlockstorageV3Attachments
type DeleteBlockstorageV3AttachmentsResponse struct{
    DeleteResult attachments.DeleteResult
}

func NewDeleteBlockstorageV3AttachmentsResponse(deleteResult attachments.DeleteResult,)*DeleteBlockstorageV3AttachmentsResponse {
    return &DeleteBlockstorageV3AttachmentsResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteBlockstorageV3Attachments(req *DeleteBlockstorageV3AttachmentsRequest)(*DeleteBlockstorageV3AttachmentsResponse){
    return NewDeleteBlockstorageV3AttachmentsResponse(attachments.Delete(oc.Client,req.Id, ))

}
//request struct for the GetBlockstorageV3Attachments
type GetBlockstorageV3AttachmentsRequest struct{
    Id string
}

func NewGetBlockstorageV3AttachmentsRequest()*GetBlockstorageV3AttachmentsRequest{
    return &GetBlockstorageV3AttachmentsRequest{}
}

//response struct for the GetBlockstorageV3Attachments
type GetBlockstorageV3AttachmentsResponse struct{
    GetResult attachments.GetResult
}

func NewGetBlockstorageV3AttachmentsResponse(getResult attachments.GetResult,)*GetBlockstorageV3AttachmentsResponse {
    return &GetBlockstorageV3AttachmentsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetBlockstorageV3Attachments(req *GetBlockstorageV3AttachmentsRequest)(*GetBlockstorageV3AttachmentsResponse){
    return NewGetBlockstorageV3AttachmentsResponse(attachments.Get(oc.Client,req.Id, ))

}
//request struct for the ListBlockstorageV3Attachments
type ListBlockstorageV3AttachmentsRequest struct{
    Opts attachments.ListOpts
}

func NewListBlockstorageV3AttachmentsRequest()*ListBlockstorageV3AttachmentsRequest{
    return &ListBlockstorageV3AttachmentsRequest{}
}

//response struct for the ListBlockstorageV3Attachments
type ListBlockstorageV3AttachmentsResponse struct{
    Pager pagination.Pager
}

func NewListBlockstorageV3AttachmentsResponse(pager pagination.Pager,)*ListBlockstorageV3AttachmentsResponse {
    return &ListBlockstorageV3AttachmentsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListBlockstorageV3Attachments(req *ListBlockstorageV3AttachmentsRequest)(*ListBlockstorageV3AttachmentsResponse){
    return NewListBlockstorageV3AttachmentsResponse(attachments.List(oc.Client,req.Opts, ))

}
//request struct for the UpdateBlockstorageV3Attachments
type UpdateBlockstorageV3AttachmentsRequest struct{
    Id string
    Opts attachments.UpdateOpts
}

func NewUpdateBlockstorageV3AttachmentsRequest()*UpdateBlockstorageV3AttachmentsRequest{
    return &UpdateBlockstorageV3AttachmentsRequest{}
}

//response struct for the UpdateBlockstorageV3Attachments
type UpdateBlockstorageV3AttachmentsResponse struct{
    UpdateResult attachments.UpdateResult
}

func NewUpdateBlockstorageV3AttachmentsResponse(updateResult attachments.UpdateResult,)*UpdateBlockstorageV3AttachmentsResponse {
    return &UpdateBlockstorageV3AttachmentsResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateBlockstorageV3Attachments(req *UpdateBlockstorageV3AttachmentsRequest)(*UpdateBlockstorageV3AttachmentsResponse){
    return NewUpdateBlockstorageV3AttachmentsResponse(attachments.Update(oc.Client,req.Id,req.Opts, ))

}
//request struct for the CompleteBlockstorageV3Attachments
type CompleteBlockstorageV3AttachmentsRequest struct{
    Id string
}

func NewCompleteBlockstorageV3AttachmentsRequest()*CompleteBlockstorageV3AttachmentsRequest{
    return &CompleteBlockstorageV3AttachmentsRequest{}
}

//response struct for the CompleteBlockstorageV3Attachments
type CompleteBlockstorageV3AttachmentsResponse struct{
    CompleteResult attachments.CompleteResult
}

func NewCompleteBlockstorageV3AttachmentsResponse(completeResult attachments.CompleteResult,)*CompleteBlockstorageV3AttachmentsResponse {
    return &CompleteBlockstorageV3AttachmentsResponse{
            CompleteResult:completeResult,
    }
}

// action function
func (oc *OpenstackClient) CompleteBlockstorageV3Attachments(req *CompleteBlockstorageV3AttachmentsRequest)(*CompleteBlockstorageV3AttachmentsResponse){
    return NewCompleteBlockstorageV3AttachmentsResponse(attachments.Complete(oc.Client,req.Id, ))

}