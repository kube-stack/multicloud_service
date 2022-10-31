package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/cdn/v1/base"
)
//request struct for the GetCdnV1Base
type GetCdnV1BaseRequest struct{
}

func NewGetCdnV1BaseRequest()*GetCdnV1BaseRequest{
    return &GetCdnV1BaseRequest{}
}

//response struct for the GetCdnV1Base
type GetCdnV1BaseResponse struct{
    GetResult base.GetResult
}

func NewGetCdnV1BaseResponse(getResult base.GetResult,)*GetCdnV1BaseResponse {
    return &GetCdnV1BaseResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetCdnV1Base(req *GetCdnV1BaseRequest)(*GetCdnV1BaseResponse){
    return NewGetCdnV1BaseResponse(base.Get(oc.Client, ))

}
//request struct for the PingCdnV1Base
type PingCdnV1BaseRequest struct{
}

func NewPingCdnV1BaseRequest()*PingCdnV1BaseRequest{
    return &PingCdnV1BaseRequest{}
}

//response struct for the PingCdnV1Base
type PingCdnV1BaseResponse struct{
    PingResult base.PingResult
}

func NewPingCdnV1BaseResponse(pingResult base.PingResult,)*PingCdnV1BaseResponse {
    return &PingCdnV1BaseResponse{
            PingResult:pingResult,
    }
}

// action function
func (oc *OpenstackClient) PingCdnV1Base(req *PingCdnV1BaseRequest)(*PingCdnV1BaseResponse){
    return NewPingCdnV1BaseResponse(base.Ping(oc.Client, ))

}