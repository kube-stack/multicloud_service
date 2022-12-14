package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/qos/ruletypes"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListRuleTypesNetworkingV2ExtensionsQosRuletypes
type ListRuleTypesNetworkingV2ExtensionsQosRuletypesRequest struct{
}

func NewListRuleTypesNetworkingV2ExtensionsQosRuletypesRequest()*ListRuleTypesNetworkingV2ExtensionsQosRuletypesRequest{
    return &ListRuleTypesNetworkingV2ExtensionsQosRuletypesRequest{}
}

//response struct for the ListRuleTypesNetworkingV2ExtensionsQosRuletypes
type ListRuleTypesNetworkingV2ExtensionsQosRuletypesResponse struct{
    Pager pagination.Pager
}

func NewListRuleTypesNetworkingV2ExtensionsQosRuletypesResponse(pager pagination.Pager,)*ListRuleTypesNetworkingV2ExtensionsQosRuletypesResponse {
    return &ListRuleTypesNetworkingV2ExtensionsQosRuletypesResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListRuleTypesNetworkingV2ExtensionsQosRuletypes(req *ListRuleTypesNetworkingV2ExtensionsQosRuletypesRequest)(*ListRuleTypesNetworkingV2ExtensionsQosRuletypesResponse){
    return NewListRuleTypesNetworkingV2ExtensionsQosRuletypesResponse(ruletypes.ListRuleTypes(oc.Client, ))

}
//request struct for the GetRuleTypeNetworkingV2ExtensionsQosRuletypes
type GetRuleTypeNetworkingV2ExtensionsQosRuletypesRequest struct{
    Name string
}

func NewGetRuleTypeNetworkingV2ExtensionsQosRuletypesRequest()*GetRuleTypeNetworkingV2ExtensionsQosRuletypesRequest{
    return &GetRuleTypeNetworkingV2ExtensionsQosRuletypesRequest{}
}

//response struct for the GetRuleTypeNetworkingV2ExtensionsQosRuletypes
type GetRuleTypeNetworkingV2ExtensionsQosRuletypesResponse struct{
    GetResult ruletypes.GetResult
}

func NewGetRuleTypeNetworkingV2ExtensionsQosRuletypesResponse(getResult ruletypes.GetResult,)*GetRuleTypeNetworkingV2ExtensionsQosRuletypesResponse {
    return &GetRuleTypeNetworkingV2ExtensionsQosRuletypesResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetRuleTypeNetworkingV2ExtensionsQosRuletypes(req *GetRuleTypeNetworkingV2ExtensionsQosRuletypesRequest)(*GetRuleTypeNetworkingV2ExtensionsQosRuletypesResponse){
    return NewGetRuleTypeNetworkingV2ExtensionsQosRuletypesResponse(ruletypes.GetRuleType(oc.Client,req.Name, ))

}