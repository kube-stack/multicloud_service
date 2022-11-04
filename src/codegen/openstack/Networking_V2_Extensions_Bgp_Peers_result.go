package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/bgp/peers"
)


//extract response info from pager for ListNetworkingV2ExtensionsBgpPeers
func ExtractListNetworkingV2ExtensionsBgpPeersResponse(response *ListNetworkingV2ExtensionsBgpPeersResponse)([]peers.BGPPeer,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return peers.ExtractBGPPeers(page)
}