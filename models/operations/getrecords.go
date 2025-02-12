// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type GetRecordsRequest struct {
	// x-collaboration
	XCollaboration *string `header:"style=simple,explode=false,name=x-collaboration"`
	// Tenant Id
	DataPartitionID string                    `header:"style=simple,explode=false,name=data-partition-id"`
	MultiRecordIds  components.MultiRecordIds `request:"mediaType=application/json"`
}

func (o *GetRecordsRequest) GetXCollaboration() *string {
	if o == nil {
		return nil
	}
	return o.XCollaboration
}

func (o *GetRecordsRequest) GetDataPartitionID() string {
	if o == nil {
		return ""
	}
	return o.DataPartitionID
}

func (o *GetRecordsRequest) GetMultiRecordIds() components.MultiRecordIds {
	if o == nil {
		return components.MultiRecordIds{}
	}
	return o.MultiRecordIds
}

type GetRecordsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Fetch multiple records successfully.
	MultiRecordInfo *components.MultiRecordInfo
}

func (o *GetRecordsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetRecordsResponse) GetMultiRecordInfo() *components.MultiRecordInfo {
	if o == nil {
		return nil
	}
	return o.MultiRecordInfo
}
