package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/listeners"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListNetworkingV2ExtensionsLbaas_v2Listeners
type ListNetworkingV2ExtensionsLbaas_v2ListenersRequest struct{
    Opts listeners.ListOpts
}

func NewListNetworkingV2ExtensionsLbaas_v2ListenersRequest()*ListNetworkingV2ExtensionsLbaas_v2ListenersRequest{
    return &ListNetworkingV2ExtensionsLbaas_v2ListenersRequest{}
}

//response struct for the ListNetworkingV2ExtensionsLbaas_v2Listeners
type ListNetworkingV2ExtensionsLbaas_v2ListenersResponse struct{
    Pager pagination.Pager
}

func NewListNetworkingV2ExtensionsLbaas_v2ListenersResponse(pager pagination.Pager,)*ListNetworkingV2ExtensionsLbaas_v2ListenersResponse {
    return &ListNetworkingV2ExtensionsLbaas_v2ListenersResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListNetworkingV2ExtensionsLbaas_v2Listeners(req *ListNetworkingV2ExtensionsLbaas_v2ListenersRequest)(*ListNetworkingV2ExtensionsLbaas_v2ListenersResponse){
    return NewListNetworkingV2ExtensionsLbaas_v2ListenersResponse(listeners.List(oc.Client,req.Opts, ))

}
//request struct for the CreateNetworkingV2ExtensionsLbaas_v2Listeners
type CreateNetworkingV2ExtensionsLbaas_v2ListenersRequest struct{
    Opts listeners.CreateOpts
}

func NewCreateNetworkingV2ExtensionsLbaas_v2ListenersRequest()*CreateNetworkingV2ExtensionsLbaas_v2ListenersRequest{
    return &CreateNetworkingV2ExtensionsLbaas_v2ListenersRequest{}
}

//response struct for the CreateNetworkingV2ExtensionsLbaas_v2Listeners
type CreateNetworkingV2ExtensionsLbaas_v2ListenersResponse struct{
    CreateResult listeners.CreateResult
}

func NewCreateNetworkingV2ExtensionsLbaas_v2ListenersResponse(createResult listeners.CreateResult,)*CreateNetworkingV2ExtensionsLbaas_v2ListenersResponse {
    return &CreateNetworkingV2ExtensionsLbaas_v2ListenersResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateNetworkingV2ExtensionsLbaas_v2Listeners(req *CreateNetworkingV2ExtensionsLbaas_v2ListenersRequest)(*CreateNetworkingV2ExtensionsLbaas_v2ListenersResponse){
    return NewCreateNetworkingV2ExtensionsLbaas_v2ListenersResponse(listeners.Create(oc.Client,req.Opts, ))

}
//request struct for the GetNetworkingV2ExtensionsLbaas_v2Listeners
type GetNetworkingV2ExtensionsLbaas_v2ListenersRequest struct{
    Id string
}

func NewGetNetworkingV2ExtensionsLbaas_v2ListenersRequest()*GetNetworkingV2ExtensionsLbaas_v2ListenersRequest{
    return &GetNetworkingV2ExtensionsLbaas_v2ListenersRequest{}
}

//response struct for the GetNetworkingV2ExtensionsLbaas_v2Listeners
type GetNetworkingV2ExtensionsLbaas_v2ListenersResponse struct{
    GetResult listeners.GetResult
}

func NewGetNetworkingV2ExtensionsLbaas_v2ListenersResponse(getResult listeners.GetResult,)*GetNetworkingV2ExtensionsLbaas_v2ListenersResponse {
    return &GetNetworkingV2ExtensionsLbaas_v2ListenersResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetNetworkingV2ExtensionsLbaas_v2Listeners(req *GetNetworkingV2ExtensionsLbaas_v2ListenersRequest)(*GetNetworkingV2ExtensionsLbaas_v2ListenersResponse){
    return NewGetNetworkingV2ExtensionsLbaas_v2ListenersResponse(listeners.Get(oc.Client,req.Id, ))

}
//request struct for the UpdateNetworkingV2ExtensionsLbaas_v2Listeners
type UpdateNetworkingV2ExtensionsLbaas_v2ListenersRequest struct{
    Id string
    Opts listeners.UpdateOpts
}

func NewUpdateNetworkingV2ExtensionsLbaas_v2ListenersRequest()*UpdateNetworkingV2ExtensionsLbaas_v2ListenersRequest{
    return &UpdateNetworkingV2ExtensionsLbaas_v2ListenersRequest{}
}

//response struct for the UpdateNetworkingV2ExtensionsLbaas_v2Listeners
type UpdateNetworkingV2ExtensionsLbaas_v2ListenersResponse struct{
    UpdateResult listeners.UpdateResult
}

func NewUpdateNetworkingV2ExtensionsLbaas_v2ListenersResponse(updateResult listeners.UpdateResult,)*UpdateNetworkingV2ExtensionsLbaas_v2ListenersResponse {
    return &UpdateNetworkingV2ExtensionsLbaas_v2ListenersResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateNetworkingV2ExtensionsLbaas_v2Listeners(req *UpdateNetworkingV2ExtensionsLbaas_v2ListenersRequest)(*UpdateNetworkingV2ExtensionsLbaas_v2ListenersResponse){
    return NewUpdateNetworkingV2ExtensionsLbaas_v2ListenersResponse(listeners.Update(oc.Client,req.Id,req.Opts, ))

}
//request struct for the DeleteNetworkingV2ExtensionsLbaas_v2Listeners
type DeleteNetworkingV2ExtensionsLbaas_v2ListenersRequest struct{
    Id string
}

func NewDeleteNetworkingV2ExtensionsLbaas_v2ListenersRequest()*DeleteNetworkingV2ExtensionsLbaas_v2ListenersRequest{
    return &DeleteNetworkingV2ExtensionsLbaas_v2ListenersRequest{}
}

//response struct for the DeleteNetworkingV2ExtensionsLbaas_v2Listeners
type DeleteNetworkingV2ExtensionsLbaas_v2ListenersResponse struct{
    DeleteResult listeners.DeleteResult
}

func NewDeleteNetworkingV2ExtensionsLbaas_v2ListenersResponse(deleteResult listeners.DeleteResult,)*DeleteNetworkingV2ExtensionsLbaas_v2ListenersResponse {
    return &DeleteNetworkingV2ExtensionsLbaas_v2ListenersResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteNetworkingV2ExtensionsLbaas_v2Listeners(req *DeleteNetworkingV2ExtensionsLbaas_v2ListenersRequest)(*DeleteNetworkingV2ExtensionsLbaas_v2ListenersResponse){
    return NewDeleteNetworkingV2ExtensionsLbaas_v2ListenersResponse(listeners.Delete(oc.Client,req.Id, ))

}