// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type MultiRecordResponse struct {
	Records            []string           `json:"records,omitempty"`
	NotFound           []string           `json:"notFound,omitempty"`
	ConversionStatuses []ConversionStatus `json:"conversionStatuses,omitempty"`
}

func (o *MultiRecordResponse) GetRecords() []string {
	if o == nil {
		return nil
	}
	return o.Records
}

func (o *MultiRecordResponse) GetNotFound() []string {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *MultiRecordResponse) GetConversionStatuses() []ConversionStatus {
	if o == nil {
		return nil
	}
	return o.ConversionStatuses
}
