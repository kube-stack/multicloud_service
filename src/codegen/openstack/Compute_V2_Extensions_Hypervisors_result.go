package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/hypervisors"
)


//extract response info from pager for ListComputeV2ExtensionsHypervisors
func ExtractListComputeV2ExtensionsHypervisorsResponse(response *ListComputeV2ExtensionsHypervisorsResponse)([]hypervisors.Hypervisor,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return hypervisors.ExtractHypervisors(page)
}


// call result's extract function
func ExtractGetStatisticsComputeV2ExtensionsHypervisorsResponse(response *GetStatisticsComputeV2ExtensionsHypervisorsResponse)(interface{}, error){
    return response.StatisticsResult.Body, response.StatisticsResult.Err
}



// call result's extract function
func ExtractGetComputeV2ExtensionsHypervisorsResponse(response *GetComputeV2ExtensionsHypervisorsResponse)(interface{}, error){
    return response.HypervisorResult.Body, response.HypervisorResult.Err
}



// call result's extract function
func ExtractGetUptimeComputeV2ExtensionsHypervisorsResponse(response *GetUptimeComputeV2ExtensionsHypervisorsResponse)(interface{}, error){
    return response.UptimeResult.Body, response.UptimeResult.Err
}
