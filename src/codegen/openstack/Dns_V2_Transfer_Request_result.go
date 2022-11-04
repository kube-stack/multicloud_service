package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/dns/v2/transfer/request"
)


//extract response info from pager for ListDnsV2TransferRequest
func ExtractListDnsV2TransferRequestResponse(response *ListDnsV2TransferRequestResponse)([]request.TransferRequest,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return request.ExtractTransferRequests(page)
}