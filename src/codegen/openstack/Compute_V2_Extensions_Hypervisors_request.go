package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/hypervisors"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListComputeV2ExtensionsHypervisors
type ListComputeV2ExtensionsHypervisorsRequest struct{
    Opts hypervisors.ListOpts
}

func NewListComputeV2ExtensionsHypervisorsRequest()*ListComputeV2ExtensionsHypervisorsRequest{
    return &ListComputeV2ExtensionsHypervisorsRequest{}
}

//response struct for the ListComputeV2ExtensionsHypervisors
type ListComputeV2ExtensionsHypervisorsResponse struct{
    Pager pagination.Pager
}

func NewListComputeV2ExtensionsHypervisorsResponse(pager pagination.Pager,)*ListComputeV2ExtensionsHypervisorsResponse {
    return &ListComputeV2ExtensionsHypervisorsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListComputeV2ExtensionsHypervisors(req *ListComputeV2ExtensionsHypervisorsRequest)(*ListComputeV2ExtensionsHypervisorsResponse){
    return NewListComputeV2ExtensionsHypervisorsResponse(hypervisors.List(oc.Client,req.Opts, ))

}
//request struct for the GetComputeV2ExtensionsHypervisors
type GetComputeV2ExtensionsHypervisorsRequest struct{
    HypervisorID string
}

func NewGetComputeV2ExtensionsHypervisorsRequest()*GetComputeV2ExtensionsHypervisorsRequest{
    return &GetComputeV2ExtensionsHypervisorsRequest{}
}

//response struct for the GetComputeV2ExtensionsHypervisors
type GetComputeV2ExtensionsHypervisorsResponse struct{
    HypervisorResult hypervisors.HypervisorResult
}

func NewGetComputeV2ExtensionsHypervisorsResponse(hypervisorResult hypervisors.HypervisorResult,)*GetComputeV2ExtensionsHypervisorsResponse {
    return &GetComputeV2ExtensionsHypervisorsResponse{
            HypervisorResult:hypervisorResult,
    }
}

// action function
func (oc *OpenstackClient) GetComputeV2ExtensionsHypervisors(req *GetComputeV2ExtensionsHypervisorsRequest)(*GetComputeV2ExtensionsHypervisorsResponse){
    return NewGetComputeV2ExtensionsHypervisorsResponse(hypervisors.Get(oc.Client,req.HypervisorID, ))

}
//request struct for the GetUptimeComputeV2ExtensionsHypervisors
type GetUptimeComputeV2ExtensionsHypervisorsRequest struct{
    HypervisorID string
}

func NewGetUptimeComputeV2ExtensionsHypervisorsRequest()*GetUptimeComputeV2ExtensionsHypervisorsRequest{
    return &GetUptimeComputeV2ExtensionsHypervisorsRequest{}
}

//response struct for the GetUptimeComputeV2ExtensionsHypervisors
type GetUptimeComputeV2ExtensionsHypervisorsResponse struct{
    UptimeResult hypervisors.UptimeResult
}

func NewGetUptimeComputeV2ExtensionsHypervisorsResponse(uptimeResult hypervisors.UptimeResult,)*GetUptimeComputeV2ExtensionsHypervisorsResponse {
    return &GetUptimeComputeV2ExtensionsHypervisorsResponse{
            UptimeResult:uptimeResult,
    }
}

// action function
func (oc *OpenstackClient) GetUptimeComputeV2ExtensionsHypervisors(req *GetUptimeComputeV2ExtensionsHypervisorsRequest)(*GetUptimeComputeV2ExtensionsHypervisorsResponse){
    return NewGetUptimeComputeV2ExtensionsHypervisorsResponse(hypervisors.GetUptime(oc.Client,req.HypervisorID, ))

}