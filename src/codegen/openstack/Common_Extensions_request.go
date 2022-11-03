package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/common/extensions"
)
//request struct for the GetCommonExtensions
type GetCommonExtensionsRequest struct{
    Alias string
}

func NewGetCommonExtensionsRequest()*GetCommonExtensionsRequest{
    return &GetCommonExtensionsRequest{}
}

//response struct for the GetCommonExtensions
type GetCommonExtensionsResponse struct{
    GetResult extensions.GetResult
}

func NewGetCommonExtensionsResponse(getResult extensions.GetResult,)*GetCommonExtensionsResponse {
    return &GetCommonExtensionsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetCommonExtensions(req *GetCommonExtensionsRequest)(*GetCommonExtensionsResponse){
    return NewGetCommonExtensionsResponse(extensions.Get(oc.Client,req.Alias, ))

}