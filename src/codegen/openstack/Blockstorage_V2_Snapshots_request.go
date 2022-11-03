package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/blockstorage/v2/snapshots"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the CreateBlockstorageV2Snapshots
type CreateBlockstorageV2SnapshotsRequest struct{
    Opts snapshots.CreateOpts
}

func NewCreateBlockstorageV2SnapshotsRequest()*CreateBlockstorageV2SnapshotsRequest{
    return &CreateBlockstorageV2SnapshotsRequest{}
}

//response struct for the CreateBlockstorageV2Snapshots
type CreateBlockstorageV2SnapshotsResponse struct{
    CreateResult snapshots.CreateResult
}

func NewCreateBlockstorageV2SnapshotsResponse(createResult snapshots.CreateResult,)*CreateBlockstorageV2SnapshotsResponse {
    return &CreateBlockstorageV2SnapshotsResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateBlockstorageV2Snapshots(req *CreateBlockstorageV2SnapshotsRequest)(*CreateBlockstorageV2SnapshotsResponse){
    return NewCreateBlockstorageV2SnapshotsResponse(snapshots.Create(oc.Client,req.Opts, ))

}
//request struct for the DeleteBlockstorageV2Snapshots
type DeleteBlockstorageV2SnapshotsRequest struct{
    Id string
}

func NewDeleteBlockstorageV2SnapshotsRequest()*DeleteBlockstorageV2SnapshotsRequest{
    return &DeleteBlockstorageV2SnapshotsRequest{}
}

//response struct for the DeleteBlockstorageV2Snapshots
type DeleteBlockstorageV2SnapshotsResponse struct{
    DeleteResult snapshots.DeleteResult
}

func NewDeleteBlockstorageV2SnapshotsResponse(deleteResult snapshots.DeleteResult,)*DeleteBlockstorageV2SnapshotsResponse {
    return &DeleteBlockstorageV2SnapshotsResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteBlockstorageV2Snapshots(req *DeleteBlockstorageV2SnapshotsRequest)(*DeleteBlockstorageV2SnapshotsResponse){
    return NewDeleteBlockstorageV2SnapshotsResponse(snapshots.Delete(oc.Client,req.Id, ))

}
//request struct for the GetBlockstorageV2Snapshots
type GetBlockstorageV2SnapshotsRequest struct{
    Id string
}

func NewGetBlockstorageV2SnapshotsRequest()*GetBlockstorageV2SnapshotsRequest{
    return &GetBlockstorageV2SnapshotsRequest{}
}

//response struct for the GetBlockstorageV2Snapshots
type GetBlockstorageV2SnapshotsResponse struct{
    GetResult snapshots.GetResult
}

func NewGetBlockstorageV2SnapshotsResponse(getResult snapshots.GetResult,)*GetBlockstorageV2SnapshotsResponse {
    return &GetBlockstorageV2SnapshotsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetBlockstorageV2Snapshots(req *GetBlockstorageV2SnapshotsRequest)(*GetBlockstorageV2SnapshotsResponse){
    return NewGetBlockstorageV2SnapshotsResponse(snapshots.Get(oc.Client,req.Id, ))

}
//request struct for the ListBlockstorageV2Snapshots
type ListBlockstorageV2SnapshotsRequest struct{
    Opts snapshots.ListOpts
}

func NewListBlockstorageV2SnapshotsRequest()*ListBlockstorageV2SnapshotsRequest{
    return &ListBlockstorageV2SnapshotsRequest{}
}

//response struct for the ListBlockstorageV2Snapshots
type ListBlockstorageV2SnapshotsResponse struct{
    Pager pagination.Pager
}

func NewListBlockstorageV2SnapshotsResponse(pager pagination.Pager,)*ListBlockstorageV2SnapshotsResponse {
    return &ListBlockstorageV2SnapshotsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListBlockstorageV2Snapshots(req *ListBlockstorageV2SnapshotsRequest)(*ListBlockstorageV2SnapshotsResponse){
    return NewListBlockstorageV2SnapshotsResponse(snapshots.List(oc.Client,req.Opts, ))

}
//request struct for the UpdateMetadataBlockstorageV2Snapshots
type UpdateMetadataBlockstorageV2SnapshotsRequest struct{
    Id string
    Opts snapshots.UpdateMetadataOpts
}

func NewUpdateMetadataBlockstorageV2SnapshotsRequest()*UpdateMetadataBlockstorageV2SnapshotsRequest{
    return &UpdateMetadataBlockstorageV2SnapshotsRequest{}
}

//response struct for the UpdateMetadataBlockstorageV2Snapshots
type UpdateMetadataBlockstorageV2SnapshotsResponse struct{
    UpdateMetadataResult snapshots.UpdateMetadataResult
}

func NewUpdateMetadataBlockstorageV2SnapshotsResponse(updateMetadataResult snapshots.UpdateMetadataResult,)*UpdateMetadataBlockstorageV2SnapshotsResponse {
    return &UpdateMetadataBlockstorageV2SnapshotsResponse{
            UpdateMetadataResult:updateMetadataResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateMetadataBlockstorageV2Snapshots(req *UpdateMetadataBlockstorageV2SnapshotsRequest)(*UpdateMetadataBlockstorageV2SnapshotsResponse){
    return NewUpdateMetadataBlockstorageV2SnapshotsResponse(snapshots.UpdateMetadata(oc.Client,req.Id,req.Opts, ))

}