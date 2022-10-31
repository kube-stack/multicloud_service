package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/loadbalancer/v2/amphorae"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListLoadbalancerV2Amphorae
type ListLoadbalancerV2AmphoraeRequest struct{
    Opts amphorae.ListOpts
}

func NewListLoadbalancerV2AmphoraeRequest()*ListLoadbalancerV2AmphoraeRequest{
    return &ListLoadbalancerV2AmphoraeRequest{}
}

//response struct for the ListLoadbalancerV2Amphorae
type ListLoadbalancerV2AmphoraeResponse struct{
    Pager pagination.Pager
}

func NewListLoadbalancerV2AmphoraeResponse(pager pagination.Pager,)*ListLoadbalancerV2AmphoraeResponse {
    return &ListLoadbalancerV2AmphoraeResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListLoadbalancerV2Amphorae(req *ListLoadbalancerV2AmphoraeRequest)(*ListLoadbalancerV2AmphoraeResponse){
    return NewListLoadbalancerV2AmphoraeResponse(amphorae.List(oc.Client,req.Opts, ))

}
//request struct for the GetLoadbalancerV2Amphorae
type GetLoadbalancerV2AmphoraeRequest struct{
    Id string
}

func NewGetLoadbalancerV2AmphoraeRequest()*GetLoadbalancerV2AmphoraeRequest{
    return &GetLoadbalancerV2AmphoraeRequest{}
}

//response struct for the GetLoadbalancerV2Amphorae
type GetLoadbalancerV2AmphoraeResponse struct{
    GetResult amphorae.GetResult
}

func NewGetLoadbalancerV2AmphoraeResponse(getResult amphorae.GetResult,)*GetLoadbalancerV2AmphoraeResponse {
    return &GetLoadbalancerV2AmphoraeResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetLoadbalancerV2Amphorae(req *GetLoadbalancerV2AmphoraeRequest)(*GetLoadbalancerV2AmphoraeResponse){
    return NewGetLoadbalancerV2AmphoraeResponse(amphorae.Get(oc.Client,req.Id, ))

}
//request struct for the FailoverLoadbalancerV2Amphorae
type FailoverLoadbalancerV2AmphoraeRequest struct{
    Id string
}

func NewFailoverLoadbalancerV2AmphoraeRequest()*FailoverLoadbalancerV2AmphoraeRequest{
    return &FailoverLoadbalancerV2AmphoraeRequest{}
}

//response struct for the FailoverLoadbalancerV2Amphorae
type FailoverLoadbalancerV2AmphoraeResponse struct{
    FailoverResult amphorae.FailoverResult
}

func NewFailoverLoadbalancerV2AmphoraeResponse(failoverResult amphorae.FailoverResult,)*FailoverLoadbalancerV2AmphoraeResponse {
    return &FailoverLoadbalancerV2AmphoraeResponse{
            FailoverResult:failoverResult,
    }
}

// action function
func (oc *OpenstackClient) FailoverLoadbalancerV2Amphorae(req *FailoverLoadbalancerV2AmphoraeRequest)(*FailoverLoadbalancerV2AmphoraeResponse){
    return NewFailoverLoadbalancerV2AmphoraeResponse(amphorae.Failover(oc.Client,req.Id, ))

}