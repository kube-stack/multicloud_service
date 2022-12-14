package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/messaging/v2/queues"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListMessagingV2Queues
type ListMessagingV2QueuesRequest struct{
    Opts queues.ListOpts
}

func NewListMessagingV2QueuesRequest()*ListMessagingV2QueuesRequest{
    return &ListMessagingV2QueuesRequest{}
}

//response struct for the ListMessagingV2Queues
type ListMessagingV2QueuesResponse struct{
    Pager pagination.Pager
}

func NewListMessagingV2QueuesResponse(pager pagination.Pager,)*ListMessagingV2QueuesResponse {
    return &ListMessagingV2QueuesResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListMessagingV2Queues(req *ListMessagingV2QueuesRequest)(*ListMessagingV2QueuesResponse){
    return NewListMessagingV2QueuesResponse(queues.List(oc.Client,req.Opts, ))

}
//request struct for the CreateMessagingV2Queues
type CreateMessagingV2QueuesRequest struct{
    Opts queues.CreateOpts
}

func NewCreateMessagingV2QueuesRequest()*CreateMessagingV2QueuesRequest{
    return &CreateMessagingV2QueuesRequest{}
}

//response struct for the CreateMessagingV2Queues
type CreateMessagingV2QueuesResponse struct{
    CreateResult queues.CreateResult
}

func NewCreateMessagingV2QueuesResponse(createResult queues.CreateResult,)*CreateMessagingV2QueuesResponse {
    return &CreateMessagingV2QueuesResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateMessagingV2Queues(req *CreateMessagingV2QueuesRequest)(*CreateMessagingV2QueuesResponse){
    return NewCreateMessagingV2QueuesResponse(queues.Create(oc.Client,req.Opts, ))

}
//request struct for the UpdateMessagingV2Queues
type UpdateMessagingV2QueuesRequest struct{
    QueueName string
    Opts queues.BatchUpdateOpts
}

func NewUpdateMessagingV2QueuesRequest()*UpdateMessagingV2QueuesRequest{
    return &UpdateMessagingV2QueuesRequest{}
}

//response struct for the UpdateMessagingV2Queues
type UpdateMessagingV2QueuesResponse struct{
    UpdateResult queues.UpdateResult
}

func NewUpdateMessagingV2QueuesResponse(updateResult queues.UpdateResult,)*UpdateMessagingV2QueuesResponse {
    return &UpdateMessagingV2QueuesResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateMessagingV2Queues(req *UpdateMessagingV2QueuesRequest)(*UpdateMessagingV2QueuesResponse){
    return NewUpdateMessagingV2QueuesResponse(queues.Update(oc.Client,req.QueueName,req.Opts, ))

}
//request struct for the GetMessagingV2Queues
type GetMessagingV2QueuesRequest struct{
    QueueName string
}

func NewGetMessagingV2QueuesRequest()*GetMessagingV2QueuesRequest{
    return &GetMessagingV2QueuesRequest{}
}

//response struct for the GetMessagingV2Queues
type GetMessagingV2QueuesResponse struct{
    GetResult queues.GetResult
}

func NewGetMessagingV2QueuesResponse(getResult queues.GetResult,)*GetMessagingV2QueuesResponse {
    return &GetMessagingV2QueuesResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetMessagingV2Queues(req *GetMessagingV2QueuesRequest)(*GetMessagingV2QueuesResponse){
    return NewGetMessagingV2QueuesResponse(queues.Get(oc.Client,req.QueueName, ))

}
//request struct for the DeleteMessagingV2Queues
type DeleteMessagingV2QueuesRequest struct{
    QueueName string
}

func NewDeleteMessagingV2QueuesRequest()*DeleteMessagingV2QueuesRequest{
    return &DeleteMessagingV2QueuesRequest{}
}

//response struct for the DeleteMessagingV2Queues
type DeleteMessagingV2QueuesResponse struct{
    DeleteResult queues.DeleteResult
}

func NewDeleteMessagingV2QueuesResponse(deleteResult queues.DeleteResult,)*DeleteMessagingV2QueuesResponse {
    return &DeleteMessagingV2QueuesResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteMessagingV2Queues(req *DeleteMessagingV2QueuesRequest)(*DeleteMessagingV2QueuesResponse){
    return NewDeleteMessagingV2QueuesResponse(queues.Delete(oc.Client,req.QueueName, ))

}
//request struct for the GetStatsMessagingV2Queues
type GetStatsMessagingV2QueuesRequest struct{
    QueueName string
}

func NewGetStatsMessagingV2QueuesRequest()*GetStatsMessagingV2QueuesRequest{
    return &GetStatsMessagingV2QueuesRequest{}
}

//response struct for the GetStatsMessagingV2Queues
type GetStatsMessagingV2QueuesResponse struct{
    StatResult queues.StatResult
}

func NewGetStatsMessagingV2QueuesResponse(statResult queues.StatResult,)*GetStatsMessagingV2QueuesResponse {
    return &GetStatsMessagingV2QueuesResponse{
            StatResult:statResult,
    }
}

// action function
func (oc *OpenstackClient) GetStatsMessagingV2Queues(req *GetStatsMessagingV2QueuesRequest)(*GetStatsMessagingV2QueuesResponse){
    return NewGetStatsMessagingV2QueuesResponse(queues.GetStats(oc.Client,req.QueueName, ))

}
//request struct for the ShareMessagingV2Queues
type ShareMessagingV2QueuesRequest struct{
    QueueName string
    Opts queues.ShareOpts
}

func NewShareMessagingV2QueuesRequest()*ShareMessagingV2QueuesRequest{
    return &ShareMessagingV2QueuesRequest{}
}

//response struct for the ShareMessagingV2Queues
type ShareMessagingV2QueuesResponse struct{
    ShareResult queues.ShareResult
}

func NewShareMessagingV2QueuesResponse(shareResult queues.ShareResult,)*ShareMessagingV2QueuesResponse {
    return &ShareMessagingV2QueuesResponse{
            ShareResult:shareResult,
    }
}

// action function
func (oc *OpenstackClient) ShareMessagingV2Queues(req *ShareMessagingV2QueuesRequest)(*ShareMessagingV2QueuesResponse){
    return NewShareMessagingV2QueuesResponse(queues.Share(oc.Client,req.QueueName,req.Opts, ))

}
//request struct for the PurgeMessagingV2Queues
type PurgeMessagingV2QueuesRequest struct{
    QueueName string
    Opts queues.PurgeOpts
}

func NewPurgeMessagingV2QueuesRequest()*PurgeMessagingV2QueuesRequest{
    return &PurgeMessagingV2QueuesRequest{}
}

//response struct for the PurgeMessagingV2Queues
type PurgeMessagingV2QueuesResponse struct{
    PurgeResult queues.PurgeResult
}

func NewPurgeMessagingV2QueuesResponse(purgeResult queues.PurgeResult,)*PurgeMessagingV2QueuesResponse {
    return &PurgeMessagingV2QueuesResponse{
            PurgeResult:purgeResult,
    }
}

// action function
func (oc *OpenstackClient) PurgeMessagingV2Queues(req *PurgeMessagingV2QueuesRequest)(*PurgeMessagingV2QueuesResponse){
    return NewPurgeMessagingV2QueuesResponse(queues.Purge(oc.Client,req.QueueName,req.Opts, ))

}