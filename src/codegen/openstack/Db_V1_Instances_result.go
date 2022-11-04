package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/db/v1/instances"
)


//extract response info from pager for ListDbV1Instances
func ExtractListDbV1InstancesResponse(response *ListDbV1InstancesResponse)([]instances.Instance,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return instances.ExtractInstances(page)
}