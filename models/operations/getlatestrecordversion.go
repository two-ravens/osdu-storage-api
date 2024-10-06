// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type GetLatestRecordVersionRequest struct {
	// x-collaboration
	XCollaboration *string `header:"style=simple,explode=false,name=x-collaboration"`
	// Record id
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Filter attributes to restrict the returned fields of the record.  Usage: data.{record-data-field-name}.
	Attribute []string `queryParam:"style=form,explode=true,name=attribute"`
	// Tenant Id
	DataPartitionID string `header:"style=simple,explode=false,name=data-partition-id"`
}

func (o *GetLatestRecordVersionRequest) GetXCollaboration() *string {
	if o == nil {
		return nil
	}
	return o.XCollaboration
}

func (o *GetLatestRecordVersionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetLatestRecordVersionRequest) GetAttribute() []string {
	if o == nil {
		return nil
	}
	return o.Attribute
}

func (o *GetLatestRecordVersionRequest) GetDataPartitionID() string {
	if o == nil {
		return ""
	}
	return o.DataPartitionID
}

type GetLatestRecordVersionResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Record retrieved successfully.
	String *string
}

func (o *GetLatestRecordVersionResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetLatestRecordVersionResponse) GetString() *string {
	if o == nil {
		return nil
	}
	return o.String
}