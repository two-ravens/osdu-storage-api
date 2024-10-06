// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type FetchRecordsRequest struct {
	// x-collaboration
	XCollaboration *string `header:"style=simple,explode=false,name=x-collaboration"`
	// Tenant Id
	DataPartitionID string `header:"style=simple,explode=false,name=data-partition-id"`
	// This value indicates whether normalization applies, should be either `none` or `units=SI;crs=wgs84;elevation=msl;azimuth=true north;dates=utc;`
	FrameOfReference   string                        `header:"style=simple,explode=false,name=frame-of-reference"`
	MultiRecordRequest components.MultiRecordRequest `request:"mediaType=application/json"`
}

func (o *FetchRecordsRequest) GetXCollaboration() *string {
	if o == nil {
		return nil
	}
	return o.XCollaboration
}

func (o *FetchRecordsRequest) GetDataPartitionID() string {
	if o == nil {
		return ""
	}
	return o.DataPartitionID
}

func (o *FetchRecordsRequest) GetFrameOfReference() string {
	if o == nil {
		return ""
	}
	return o.FrameOfReference
}

func (o *FetchRecordsRequest) GetMultiRecordRequest() components.MultiRecordRequest {
	if o == nil {
		return components.MultiRecordRequest{}
	}
	return o.MultiRecordRequest
}

type FetchRecordsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Fetch multiple records successfully.
	MultiRecordResponse *components.MultiRecordResponse
}

func (o *FetchRecordsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *FetchRecordsResponse) GetMultiRecordResponse() *components.MultiRecordResponse {
	if o == nil {
		return nil
	}
	return o.MultiRecordResponse
}