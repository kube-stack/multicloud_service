package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/sharedfilesystems/v2/availabilityzones"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListSharedfilesystemsV2Availabilityzones
type ListSharedfilesystemsV2AvailabilityzonesRequest struct{
}

func NewListSharedfilesystemsV2AvailabilityzonesRequest()*ListSharedfilesystemsV2AvailabilityzonesRequest{
    return &ListSharedfilesystemsV2AvailabilityzonesRequest{}
}

//response struct for the ListSharedfilesystemsV2Availabilityzones
type ListSharedfilesystemsV2AvailabilityzonesResponse struct{
    Pager pagination.Pager
}

func NewListSharedfilesystemsV2AvailabilityzonesResponse(pager pagination.Pager,)*ListSharedfilesystemsV2AvailabilityzonesResponse {
    return &ListSharedfilesystemsV2AvailabilityzonesResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListSharedfilesystemsV2Availabilityzones(req *ListSharedfilesystemsV2AvailabilityzonesRequest)(*ListSharedfilesystemsV2AvailabilityzonesResponse){
    return NewListSharedfilesystemsV2AvailabilityzonesResponse(availabilityzones.List(oc.Client, ))

}