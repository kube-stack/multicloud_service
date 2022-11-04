package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/blockstorage/extensions/limits"
)
//request struct for the GetBlockstorageExtensionsLimits
type GetBlockstorageExtensionsLimitsRequest struct{
}

func NewGetBlockstorageExtensionsLimitsRequest()*GetBlockstorageExtensionsLimitsRequest{
    return &GetBlockstorageExtensionsLimitsRequest{}
}

//response struct for the GetBlockstorageExtensionsLimits
type GetBlockstorageExtensionsLimitsResponse struct{
    GetResult limits.GetResult
}

func NewGetBlockstorageExtensionsLimitsResponse(getResult limits.GetResult,)*GetBlockstorageExtensionsLimitsResponse {
    return &GetBlockstorageExtensionsLimitsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetBlockstorageExtensionsLimits(req *GetBlockstorageExtensionsLimitsRequest)(*GetBlockstorageExtensionsLimitsResponse){
    return NewGetBlockstorageExtensionsLimitsResponse(limits.Get(oc.Client, ))

}