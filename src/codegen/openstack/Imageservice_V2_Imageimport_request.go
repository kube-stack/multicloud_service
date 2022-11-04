package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/imageservice/v2/imageimport"
)
//request struct for the GetImageserviceV2Imageimport
type GetImageserviceV2ImageimportRequest struct{
}

func NewGetImageserviceV2ImageimportRequest()*GetImageserviceV2ImageimportRequest{
    return &GetImageserviceV2ImageimportRequest{}
}

//response struct for the GetImageserviceV2Imageimport
type GetImageserviceV2ImageimportResponse struct{
    GetResult imageimport.GetResult
}

func NewGetImageserviceV2ImageimportResponse(getResult imageimport.GetResult,)*GetImageserviceV2ImageimportResponse {
    return &GetImageserviceV2ImageimportResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetImageserviceV2Imageimport(req *GetImageserviceV2ImageimportRequest)(*GetImageserviceV2ImageimportResponse){
    return NewGetImageserviceV2ImageimportResponse(imageimport.Get(oc.Client, ))

}
//request struct for the CreateImageserviceV2Imageimport
type CreateImageserviceV2ImageimportRequest struct{
    ImageID string
    Opts imageimport.CreateOpts
}

func NewCreateImageserviceV2ImageimportRequest()*CreateImageserviceV2ImageimportRequest{
    return &CreateImageserviceV2ImageimportRequest{}
}

//response struct for the CreateImageserviceV2Imageimport
type CreateImageserviceV2ImageimportResponse struct{
    CreateResult imageimport.CreateResult
}

func NewCreateImageserviceV2ImageimportResponse(createResult imageimport.CreateResult,)*CreateImageserviceV2ImageimportResponse {
    return &CreateImageserviceV2ImageimportResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateImageserviceV2Imageimport(req *CreateImageserviceV2ImageimportRequest)(*CreateImageserviceV2ImageimportResponse){
    return NewCreateImageserviceV2ImageimportResponse(imageimport.Create(oc.Client,req.ImageID,req.Opts, ))

}