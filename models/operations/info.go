// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type InfoRequest struct {
	// Tenant Id
	DataPartitionID string `header:"style=simple,explode=false,name=data-partition-id"`
}

func (o *InfoRequest) GetDataPartitionID() string {
	if o == nil {
		return ""
	}
	return o.DataPartitionID
}

type InfoResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	VersionInfo *components.VersionInfo
}

func (o *InfoResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *InfoResponse) GetVersionInfo() *components.VersionInfo {
	if o == nil {
		return nil
	}
	return o.VersionInfo
}
