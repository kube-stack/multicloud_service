package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/identity/v3/limits"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListIdentityV3Limits
type ListIdentityV3LimitsRequest struct{
    Opts limits.ListOpts
}

func NewListIdentityV3LimitsRequest()*ListIdentityV3LimitsRequest{
    return &ListIdentityV3LimitsRequest{}
}

//response struct for the ListIdentityV3Limits
type ListIdentityV3LimitsResponse struct{
    Pager pagination.Pager
}

func NewListIdentityV3LimitsResponse(pager pagination.Pager,)*ListIdentityV3LimitsResponse {
    return &ListIdentityV3LimitsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListIdentityV3Limits(req *ListIdentityV3LimitsRequest)(*ListIdentityV3LimitsResponse){
    return NewListIdentityV3LimitsResponse(limits.List(oc.Client,req.Opts, ))

}