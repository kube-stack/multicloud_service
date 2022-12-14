package openstack

// Code generated by cloud manager.

import (
    "github.com/gophercloud/gophercloud/openstack/blockstorage/v3/attachments"
)


// call result's extract function
func ExtractCreateBlockstorageV3AttachmentsResponse(response *CreateBlockstorageV3AttachmentsResponse)(interface{}, error){
    return response.CreateResult.Body, response.CreateResult.Err
}



// call result's extract function
func ExtractDeleteBlockstorageV3AttachmentsResponse(response *DeleteBlockstorageV3AttachmentsResponse)(interface{}, error){
    return response.DeleteResult.Body, response.DeleteResult.Err
}



// call result's extract function
func ExtractGetBlockstorageV3AttachmentsResponse(response *GetBlockstorageV3AttachmentsResponse)(interface{}, error){
    return response.GetResult.Body, response.GetResult.Err
}



//extract response info from pager for ListBlockstorageV3Attachments
func ExtractListBlockstorageV3AttachmentsResponse(response *ListBlockstorageV3AttachmentsResponse)([]attachments.Attachment,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return attachments.ExtractAttachments(page)
}


// call result's extract function
func ExtractUpdateBlockstorageV3AttachmentsResponse(response *UpdateBlockstorageV3AttachmentsResponse)(interface{}, error){
    return response.UpdateResult.Body, response.UpdateResult.Err
}



// call result's extract function
func ExtractCompleteBlockstorageV3AttachmentsResponse(response *CompleteBlockstorageV3AttachmentsResponse)(interface{}, error){
    return response.CompleteResult.Body, response.CompleteResult.Err
}
