package openstack

// Code generated by cloud manager.

import (
)


// call result's extract function
func ExtractGetOrchestrationV1StacktemplatesResponse(response *GetOrchestrationV1StacktemplatesResponse)(interface{}, error){
    return response.GetResult.Body, response.GetResult.Err
}



// call result's extract function
func ExtractValidateOrchestrationV1StacktemplatesResponse(response *ValidateOrchestrationV1StacktemplatesResponse)(interface{}, error){
    return response.ValidateResult.Body, response.ValidateResult.Err
}
