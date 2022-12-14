package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/compute/v2/images"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListDetailComputeV2Images
type ListDetailComputeV2ImagesRequest struct{
    Opts images.ListOpts
}

func NewListDetailComputeV2ImagesRequest()*ListDetailComputeV2ImagesRequest{
    return &ListDetailComputeV2ImagesRequest{}
}

//response struct for the ListDetailComputeV2Images
type ListDetailComputeV2ImagesResponse struct{
    Pager pagination.Pager
}

func NewListDetailComputeV2ImagesResponse(pager pagination.Pager,)*ListDetailComputeV2ImagesResponse {
    return &ListDetailComputeV2ImagesResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListDetailComputeV2Images(req *ListDetailComputeV2ImagesRequest)(*ListDetailComputeV2ImagesResponse){
    return NewListDetailComputeV2ImagesResponse(images.ListDetail(oc.Client,req.Opts, ))

}
//request struct for the GetComputeV2Images
type GetComputeV2ImagesRequest struct{
    Id string
}

func NewGetComputeV2ImagesRequest()*GetComputeV2ImagesRequest{
    return &GetComputeV2ImagesRequest{}
}

//response struct for the GetComputeV2Images
type GetComputeV2ImagesResponse struct{
    GetResult images.GetResult
}

func NewGetComputeV2ImagesResponse(getResult images.GetResult,)*GetComputeV2ImagesResponse {
    return &GetComputeV2ImagesResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetComputeV2Images(req *GetComputeV2ImagesRequest)(*GetComputeV2ImagesResponse){
    return NewGetComputeV2ImagesResponse(images.Get(oc.Client,req.Id, ))

}
//request struct for the DeleteComputeV2Images
type DeleteComputeV2ImagesRequest struct{
    Id string
}

func NewDeleteComputeV2ImagesRequest()*DeleteComputeV2ImagesRequest{
    return &DeleteComputeV2ImagesRequest{}
}

//response struct for the DeleteComputeV2Images
type DeleteComputeV2ImagesResponse struct{
    DeleteResult images.DeleteResult
}

func NewDeleteComputeV2ImagesResponse(deleteResult images.DeleteResult,)*DeleteComputeV2ImagesResponse {
    return &DeleteComputeV2ImagesResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteComputeV2Images(req *DeleteComputeV2ImagesRequest)(*DeleteComputeV2ImagesResponse){
    return NewDeleteComputeV2ImagesResponse(images.Delete(oc.Client,req.Id, ))

}