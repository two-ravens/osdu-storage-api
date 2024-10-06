// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ConversionStatus struct {
	ID     *string  `json:"id,omitempty"`
	Status *string  `json:"status,omitempty"`
	Errors []string `json:"errors,omitempty"`
}

func (o *ConversionStatus) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ConversionStatus) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *ConversionStatus) GetErrors() []string {
	if o == nil {
		return nil
	}
	return o.Errors
}
