package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/containerinfra/v1/clusters"
)


//extract response info from pager for ListContainerinfraV1Clusters
func ExtractListContainerinfraV1ClustersResponse(response *ListContainerinfraV1ClustersResponse)([]clusters.Cluster,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return clusters.ExtractClusters(page)
}


//extract response info from pager for ListDetailContainerinfraV1Clusters
func ExtractListDetailContainerinfraV1ClustersResponse(response *ListDetailContainerinfraV1ClustersResponse)([]clusters.Cluster,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return clusters.ExtractClusters(page)
}