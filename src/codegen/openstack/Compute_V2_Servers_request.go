package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListComputeV2Servers
type ListComputeV2ServersRequest struct{
    Opts servers.ListOpts
}

func NewListComputeV2ServersRequest()*ListComputeV2ServersRequest{
    return &ListComputeV2ServersRequest{}
}

//response struct for the ListComputeV2Servers
type ListComputeV2ServersResponse struct{
    Pager pagination.Pager
}

func NewListComputeV2ServersResponse(pager pagination.Pager,)*ListComputeV2ServersResponse {
    return &ListComputeV2ServersResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListComputeV2Servers(req *ListComputeV2ServersRequest)(*ListComputeV2ServersResponse){
    return NewListComputeV2ServersResponse(servers.List(oc.Client,req.Opts, ))

}
//request struct for the CreateComputeV2Servers
type CreateComputeV2ServersRequest struct{
    Opts servers.CreateOpts
}

func NewCreateComputeV2ServersRequest()*CreateComputeV2ServersRequest{
    return &CreateComputeV2ServersRequest{}
}

//response struct for the CreateComputeV2Servers
type CreateComputeV2ServersResponse struct{
    CreateResult servers.CreateResult
}

func NewCreateComputeV2ServersResponse(createResult servers.CreateResult,)*CreateComputeV2ServersResponse {
    return &CreateComputeV2ServersResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateComputeV2Servers(req *CreateComputeV2ServersRequest)(*CreateComputeV2ServersResponse){
    return NewCreateComputeV2ServersResponse(servers.Create(oc.Client,req.Opts, ))

}
//request struct for the DeleteComputeV2Servers
type DeleteComputeV2ServersRequest struct{
    Id string
}

func NewDeleteComputeV2ServersRequest()*DeleteComputeV2ServersRequest{
    return &DeleteComputeV2ServersRequest{}
}

//response struct for the DeleteComputeV2Servers
type DeleteComputeV2ServersResponse struct{
    DeleteResult servers.DeleteResult
}

func NewDeleteComputeV2ServersResponse(deleteResult servers.DeleteResult,)*DeleteComputeV2ServersResponse {
    return &DeleteComputeV2ServersResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteComputeV2Servers(req *DeleteComputeV2ServersRequest)(*DeleteComputeV2ServersResponse){
    return NewDeleteComputeV2ServersResponse(servers.Delete(oc.Client,req.Id, ))

}
//request struct for the ForceDeleteComputeV2Servers
type ForceDeleteComputeV2ServersRequest struct{
    Id string
}

func NewForceDeleteComputeV2ServersRequest()*ForceDeleteComputeV2ServersRequest{
    return &ForceDeleteComputeV2ServersRequest{}
}

//response struct for the ForceDeleteComputeV2Servers
type ForceDeleteComputeV2ServersResponse struct{
    ActionResult servers.ActionResult
}

func NewForceDeleteComputeV2ServersResponse(actionResult servers.ActionResult,)*ForceDeleteComputeV2ServersResponse {
    return &ForceDeleteComputeV2ServersResponse{
            ActionResult:actionResult,
    }
}

// action function
func (oc *OpenstackClient) ForceDeleteComputeV2Servers(req *ForceDeleteComputeV2ServersRequest)(*ForceDeleteComputeV2ServersResponse){
    return NewForceDeleteComputeV2ServersResponse(servers.ForceDelete(oc.Client,req.Id, ))

}
//request struct for the GetComputeV2Servers
type GetComputeV2ServersRequest struct{
    Id string
}

func NewGetComputeV2ServersRequest()*GetComputeV2ServersRequest{
    return &GetComputeV2ServersRequest{}
}

//response struct for the GetComputeV2Servers
type GetComputeV2ServersResponse struct{
    GetResult servers.GetResult
}

func NewGetComputeV2ServersResponse(getResult servers.GetResult,)*GetComputeV2ServersResponse {
    return &GetComputeV2ServersResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetComputeV2Servers(req *GetComputeV2ServersRequest)(*GetComputeV2ServersResponse){
    return NewGetComputeV2ServersResponse(servers.Get(oc.Client,req.Id, ))

}
//request struct for the UpdateComputeV2Servers
type UpdateComputeV2ServersRequest struct{
    Id string
    Opts servers.UpdateOpts
}

func NewUpdateComputeV2ServersRequest()*UpdateComputeV2ServersRequest{
    return &UpdateComputeV2ServersRequest{}
}

//response struct for the UpdateComputeV2Servers
type UpdateComputeV2ServersResponse struct{
    UpdateResult servers.UpdateResult
}

func NewUpdateComputeV2ServersResponse(updateResult servers.UpdateResult,)*UpdateComputeV2ServersResponse {
    return &UpdateComputeV2ServersResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateComputeV2Servers(req *UpdateComputeV2ServersRequest)(*UpdateComputeV2ServersResponse){
    return NewUpdateComputeV2ServersResponse(servers.Update(oc.Client,req.Id,req.Opts, ))

}
//request struct for the ChangeAdminPasswordComputeV2Servers
type ChangeAdminPasswordComputeV2ServersRequest struct{
    Id string
    NewPassword string
}

func NewChangeAdminPasswordComputeV2ServersRequest()*ChangeAdminPasswordComputeV2ServersRequest{
    return &ChangeAdminPasswordComputeV2ServersRequest{}
}

//response struct for the ChangeAdminPasswordComputeV2Servers
type ChangeAdminPasswordComputeV2ServersResponse struct{
    ActionResult servers.ActionResult
}

func NewChangeAdminPasswordComputeV2ServersResponse(actionResult servers.ActionResult,)*ChangeAdminPasswordComputeV2ServersResponse {
    return &ChangeAdminPasswordComputeV2ServersResponse{
            ActionResult:actionResult,
    }
}

// action function
func (oc *OpenstackClient) ChangeAdminPasswordComputeV2Servers(req *ChangeAdminPasswordComputeV2ServersRequest)(*ChangeAdminPasswordComputeV2ServersResponse){
    return NewChangeAdminPasswordComputeV2ServersResponse(servers.ChangeAdminPassword(oc.Client,req.Id,req.NewPassword, ))

}
//request struct for the RebootComputeV2Servers
type RebootComputeV2ServersRequest struct{
    Id string
    Opts servers.RebootOpts
}

func NewRebootComputeV2ServersRequest()*RebootComputeV2ServersRequest{
    return &RebootComputeV2ServersRequest{}
}

//response struct for the RebootComputeV2Servers
type RebootComputeV2ServersResponse struct{
    ActionResult servers.ActionResult
}

func NewRebootComputeV2ServersResponse(actionResult servers.ActionResult,)*RebootComputeV2ServersResponse {
    return &RebootComputeV2ServersResponse{
            ActionResult:actionResult,
    }
}

// action function
func (oc *OpenstackClient) RebootComputeV2Servers(req *RebootComputeV2ServersRequest)(*RebootComputeV2ServersResponse){
    return NewRebootComputeV2ServersResponse(servers.Reboot(oc.Client,req.Id,req.Opts, ))

}
//request struct for the RebuildComputeV2Servers
type RebuildComputeV2ServersRequest struct{
    Id string
    Opts servers.RebuildOpts
}

func NewRebuildComputeV2ServersRequest()*RebuildComputeV2ServersRequest{
    return &RebuildComputeV2ServersRequest{}
}

//response struct for the RebuildComputeV2Servers
type RebuildComputeV2ServersResponse struct{
    RebuildResult servers.RebuildResult
}

func NewRebuildComputeV2ServersResponse(rebuildResult servers.RebuildResult,)*RebuildComputeV2ServersResponse {
    return &RebuildComputeV2ServersResponse{
            RebuildResult:rebuildResult,
    }
}

// action function
func (oc *OpenstackClient) RebuildComputeV2Servers(req *RebuildComputeV2ServersRequest)(*RebuildComputeV2ServersResponse){
    return NewRebuildComputeV2ServersResponse(servers.Rebuild(oc.Client,req.Id,req.Opts, ))

}
//request struct for the ResizeComputeV2Servers
type ResizeComputeV2ServersRequest struct{
    Id string
    Opts servers.ResizeOpts
}

func NewResizeComputeV2ServersRequest()*ResizeComputeV2ServersRequest{
    return &ResizeComputeV2ServersRequest{}
}

//response struct for the ResizeComputeV2Servers
type ResizeComputeV2ServersResponse struct{
    ActionResult servers.ActionResult
}

func NewResizeComputeV2ServersResponse(actionResult servers.ActionResult,)*ResizeComputeV2ServersResponse {
    return &ResizeComputeV2ServersResponse{
            ActionResult:actionResult,
    }
}

// action function
func (oc *OpenstackClient) ResizeComputeV2Servers(req *ResizeComputeV2ServersRequest)(*ResizeComputeV2ServersResponse){
    return NewResizeComputeV2ServersResponse(servers.Resize(oc.Client,req.Id,req.Opts, ))

}
//request struct for the ConfirmResizeComputeV2Servers
type ConfirmResizeComputeV2ServersRequest struct{
    Id string
}

func NewConfirmResizeComputeV2ServersRequest()*ConfirmResizeComputeV2ServersRequest{
    return &ConfirmResizeComputeV2ServersRequest{}
}

//response struct for the ConfirmResizeComputeV2Servers
type ConfirmResizeComputeV2ServersResponse struct{
    ActionResult servers.ActionResult
}

func NewConfirmResizeComputeV2ServersResponse(actionResult servers.ActionResult,)*ConfirmResizeComputeV2ServersResponse {
    return &ConfirmResizeComputeV2ServersResponse{
            ActionResult:actionResult,
    }
}

// action function
func (oc *OpenstackClient) ConfirmResizeComputeV2Servers(req *ConfirmResizeComputeV2ServersRequest)(*ConfirmResizeComputeV2ServersResponse){
    return NewConfirmResizeComputeV2ServersResponse(servers.ConfirmResize(oc.Client,req.Id, ))

}
//request struct for the RevertResizeComputeV2Servers
type RevertResizeComputeV2ServersRequest struct{
    Id string
}

func NewRevertResizeComputeV2ServersRequest()*RevertResizeComputeV2ServersRequest{
    return &RevertResizeComputeV2ServersRequest{}
}

//response struct for the RevertResizeComputeV2Servers
type RevertResizeComputeV2ServersResponse struct{
    ActionResult servers.ActionResult
}

func NewRevertResizeComputeV2ServersResponse(actionResult servers.ActionResult,)*RevertResizeComputeV2ServersResponse {
    return &RevertResizeComputeV2ServersResponse{
            ActionResult:actionResult,
    }
}

// action function
func (oc *OpenstackClient) RevertResizeComputeV2Servers(req *RevertResizeComputeV2ServersRequest)(*RevertResizeComputeV2ServersResponse){
    return NewRevertResizeComputeV2ServersResponse(servers.RevertResize(oc.Client,req.Id, ))

}
//request struct for the ResetMetadataComputeV2Servers
type ResetMetadataComputeV2ServersRequest struct{
    Id string
    Opts servers.MetadataOpts
}

func NewResetMetadataComputeV2ServersRequest()*ResetMetadataComputeV2ServersRequest{
    return &ResetMetadataComputeV2ServersRequest{}
}

//response struct for the ResetMetadataComputeV2Servers
type ResetMetadataComputeV2ServersResponse struct{
    ResetMetadataResult servers.ResetMetadataResult
}

func NewResetMetadataComputeV2ServersResponse(resetMetadataResult servers.ResetMetadataResult,)*ResetMetadataComputeV2ServersResponse {
    return &ResetMetadataComputeV2ServersResponse{
            ResetMetadataResult:resetMetadataResult,
    }
}

// action function
func (oc *OpenstackClient) ResetMetadataComputeV2Servers(req *ResetMetadataComputeV2ServersRequest)(*ResetMetadataComputeV2ServersResponse){
    return NewResetMetadataComputeV2ServersResponse(servers.ResetMetadata(oc.Client,req.Id,req.Opts, ))

}
//request struct for the MetadataComputeV2Servers
type MetadataComputeV2ServersRequest struct{
    Id string
}

func NewMetadataComputeV2ServersRequest()*MetadataComputeV2ServersRequest{
    return &MetadataComputeV2ServersRequest{}
}

//response struct for the MetadataComputeV2Servers
type MetadataComputeV2ServersResponse struct{
    GetMetadataResult servers.GetMetadataResult
}

func NewMetadataComputeV2ServersResponse(getMetadataResult servers.GetMetadataResult,)*MetadataComputeV2ServersResponse {
    return &MetadataComputeV2ServersResponse{
            GetMetadataResult:getMetadataResult,
    }
}

// action function
func (oc *OpenstackClient) MetadataComputeV2Servers(req *MetadataComputeV2ServersRequest)(*MetadataComputeV2ServersResponse){
    return NewMetadataComputeV2ServersResponse(servers.Metadata(oc.Client,req.Id, ))

}
//request struct for the UpdateMetadataComputeV2Servers
type UpdateMetadataComputeV2ServersRequest struct{
    Id string
    Opts servers.MetadataOpts
}

func NewUpdateMetadataComputeV2ServersRequest()*UpdateMetadataComputeV2ServersRequest{
    return &UpdateMetadataComputeV2ServersRequest{}
}

//response struct for the UpdateMetadataComputeV2Servers
type UpdateMetadataComputeV2ServersResponse struct{
    UpdateMetadataResult servers.UpdateMetadataResult
}

func NewUpdateMetadataComputeV2ServersResponse(updateMetadataResult servers.UpdateMetadataResult,)*UpdateMetadataComputeV2ServersResponse {
    return &UpdateMetadataComputeV2ServersResponse{
            UpdateMetadataResult:updateMetadataResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateMetadataComputeV2Servers(req *UpdateMetadataComputeV2ServersRequest)(*UpdateMetadataComputeV2ServersResponse){
    return NewUpdateMetadataComputeV2ServersResponse(servers.UpdateMetadata(oc.Client,req.Id,req.Opts, ))

}
//request struct for the CreateMetadatumComputeV2Servers
type CreateMetadatumComputeV2ServersRequest struct{
    Id string
    Opts servers.MetadatumOpts
}

func NewCreateMetadatumComputeV2ServersRequest()*CreateMetadatumComputeV2ServersRequest{
    return &CreateMetadatumComputeV2ServersRequest{}
}

//response struct for the CreateMetadatumComputeV2Servers
type CreateMetadatumComputeV2ServersResponse struct{
    CreateMetadatumResult servers.CreateMetadatumResult
}

func NewCreateMetadatumComputeV2ServersResponse(createMetadatumResult servers.CreateMetadatumResult,)*CreateMetadatumComputeV2ServersResponse {
    return &CreateMetadatumComputeV2ServersResponse{
            CreateMetadatumResult:createMetadatumResult,
    }
}

// action function
func (oc *OpenstackClient) CreateMetadatumComputeV2Servers(req *CreateMetadatumComputeV2ServersRequest)(*CreateMetadatumComputeV2ServersResponse){
    return NewCreateMetadatumComputeV2ServersResponse(servers.CreateMetadatum(oc.Client,req.Id,req.Opts, ))

}
//request struct for the MetadatumComputeV2Servers
type MetadatumComputeV2ServersRequest struct{
    Id string
    Key string
}

func NewMetadatumComputeV2ServersRequest()*MetadatumComputeV2ServersRequest{
    return &MetadatumComputeV2ServersRequest{}
}

//response struct for the MetadatumComputeV2Servers
type MetadatumComputeV2ServersResponse struct{
    GetMetadatumResult servers.GetMetadatumResult
}

func NewMetadatumComputeV2ServersResponse(getMetadatumResult servers.GetMetadatumResult,)*MetadatumComputeV2ServersResponse {
    return &MetadatumComputeV2ServersResponse{
            GetMetadatumResult:getMetadatumResult,
    }
}

// action function
func (oc *OpenstackClient) MetadatumComputeV2Servers(req *MetadatumComputeV2ServersRequest)(*MetadatumComputeV2ServersResponse){
    return NewMetadatumComputeV2ServersResponse(servers.Metadatum(oc.Client,req.Id,req.Key, ))

}
//request struct for the DeleteMetadatumComputeV2Servers
type DeleteMetadatumComputeV2ServersRequest struct{
    Id string
    Key string
}

func NewDeleteMetadatumComputeV2ServersRequest()*DeleteMetadatumComputeV2ServersRequest{
    return &DeleteMetadatumComputeV2ServersRequest{}
}

//response struct for the DeleteMetadatumComputeV2Servers
type DeleteMetadatumComputeV2ServersResponse struct{
    DeleteMetadatumResult servers.DeleteMetadatumResult
}

func NewDeleteMetadatumComputeV2ServersResponse(deleteMetadatumResult servers.DeleteMetadatumResult,)*DeleteMetadatumComputeV2ServersResponse {
    return &DeleteMetadatumComputeV2ServersResponse{
            DeleteMetadatumResult:deleteMetadatumResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteMetadatumComputeV2Servers(req *DeleteMetadatumComputeV2ServersRequest)(*DeleteMetadatumComputeV2ServersResponse){
    return NewDeleteMetadatumComputeV2ServersResponse(servers.DeleteMetadatum(oc.Client,req.Id,req.Key, ))

}
//request struct for the ListAddressesComputeV2Servers
type ListAddressesComputeV2ServersRequest struct{
    Id string
}

func NewListAddressesComputeV2ServersRequest()*ListAddressesComputeV2ServersRequest{
    return &ListAddressesComputeV2ServersRequest{}
}

//response struct for the ListAddressesComputeV2Servers
type ListAddressesComputeV2ServersResponse struct{
    Pager pagination.Pager
}

func NewListAddressesComputeV2ServersResponse(pager pagination.Pager,)*ListAddressesComputeV2ServersResponse {
    return &ListAddressesComputeV2ServersResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListAddressesComputeV2Servers(req *ListAddressesComputeV2ServersRequest)(*ListAddressesComputeV2ServersResponse){
    return NewListAddressesComputeV2ServersResponse(servers.ListAddresses(oc.Client,req.Id, ))

}
//request struct for the ListAddressesByNetworkComputeV2Servers
type ListAddressesByNetworkComputeV2ServersRequest struct{
    Id string
    Network string
}

func NewListAddressesByNetworkComputeV2ServersRequest()*ListAddressesByNetworkComputeV2ServersRequest{
    return &ListAddressesByNetworkComputeV2ServersRequest{}
}

//response struct for the ListAddressesByNetworkComputeV2Servers
type ListAddressesByNetworkComputeV2ServersResponse struct{
    Pager pagination.Pager
}

func NewListAddressesByNetworkComputeV2ServersResponse(pager pagination.Pager,)*ListAddressesByNetworkComputeV2ServersResponse {
    return &ListAddressesByNetworkComputeV2ServersResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListAddressesByNetworkComputeV2Servers(req *ListAddressesByNetworkComputeV2ServersRequest)(*ListAddressesByNetworkComputeV2ServersResponse){
    return NewListAddressesByNetworkComputeV2ServersResponse(servers.ListAddressesByNetwork(oc.Client,req.Id,req.Network, ))

}
//request struct for the CreateImageComputeV2Servers
type CreateImageComputeV2ServersRequest struct{
    Id string
    Opts servers.CreateImageOpts
}

func NewCreateImageComputeV2ServersRequest()*CreateImageComputeV2ServersRequest{
    return &CreateImageComputeV2ServersRequest{}
}

//response struct for the CreateImageComputeV2Servers
type CreateImageComputeV2ServersResponse struct{
    CreateImageResult servers.CreateImageResult
}

func NewCreateImageComputeV2ServersResponse(createImageResult servers.CreateImageResult,)*CreateImageComputeV2ServersResponse {
    return &CreateImageComputeV2ServersResponse{
            CreateImageResult:createImageResult,
    }
}

// action function
func (oc *OpenstackClient) CreateImageComputeV2Servers(req *CreateImageComputeV2ServersRequest)(*CreateImageComputeV2ServersResponse){
    return NewCreateImageComputeV2ServersResponse(servers.CreateImage(oc.Client,req.Id,req.Opts, ))

}
//request struct for the GetPasswordComputeV2Servers
type GetPasswordComputeV2ServersRequest struct{
    ServerId string
}

func NewGetPasswordComputeV2ServersRequest()*GetPasswordComputeV2ServersRequest{
    return &GetPasswordComputeV2ServersRequest{}
}

//response struct for the GetPasswordComputeV2Servers
type GetPasswordComputeV2ServersResponse struct{
    GetPasswordResult servers.GetPasswordResult
}

func NewGetPasswordComputeV2ServersResponse(getPasswordResult servers.GetPasswordResult,)*GetPasswordComputeV2ServersResponse {
    return &GetPasswordComputeV2ServersResponse{
            GetPasswordResult:getPasswordResult,
    }
}

// action function
func (oc *OpenstackClient) GetPasswordComputeV2Servers(req *GetPasswordComputeV2ServersRequest)(*GetPasswordComputeV2ServersResponse){
    return NewGetPasswordComputeV2ServersResponse(servers.GetPassword(oc.Client,req.ServerId, ))

}
//request struct for the ShowConsoleOutputComputeV2Servers
type ShowConsoleOutputComputeV2ServersRequest struct{
    Id string
    Opts servers.ShowConsoleOutputOpts
}

func NewShowConsoleOutputComputeV2ServersRequest()*ShowConsoleOutputComputeV2ServersRequest{
    return &ShowConsoleOutputComputeV2ServersRequest{}
}

//response struct for the ShowConsoleOutputComputeV2Servers
type ShowConsoleOutputComputeV2ServersResponse struct{
    ShowConsoleOutputResult servers.ShowConsoleOutputResult
}

func NewShowConsoleOutputComputeV2ServersResponse(showConsoleOutputResult servers.ShowConsoleOutputResult,)*ShowConsoleOutputComputeV2ServersResponse {
    return &ShowConsoleOutputComputeV2ServersResponse{
            ShowConsoleOutputResult:showConsoleOutputResult,
    }
}

// action function
func (oc *OpenstackClient) ShowConsoleOutputComputeV2Servers(req *ShowConsoleOutputComputeV2ServersRequest)(*ShowConsoleOutputComputeV2ServersResponse){
    return NewShowConsoleOutputComputeV2ServersResponse(servers.ShowConsoleOutput(oc.Client,req.Id,req.Opts, ))

}