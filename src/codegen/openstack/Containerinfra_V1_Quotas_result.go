package openstack

// Code generated by cloud manager.

import (
)


// call result's extract function
func ExtractCreateContainerinfraV1QuotasResponse(response *CreateContainerinfraV1QuotasResponse)(interface{}, error){
    return response.CreateResult.Body, response.CreateResult.Err
}
