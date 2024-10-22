// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type DatastoreQueryResult struct {
	Cursor  *string  `json:"cursor,omitempty"`
	Results []string `json:"results,omitempty"`
}

func (o *DatastoreQueryResult) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *DatastoreQueryResult) GetResults() []string {
	if o == nil {
		return nil
	}
	return o.Results
}
