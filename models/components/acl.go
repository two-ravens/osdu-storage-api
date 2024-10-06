// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ACL struct {
	Viewers []string `json:"viewers,omitempty"`
	Owners  []string `json:"owners,omitempty"`
}

func (o *ACL) GetViewers() []string {
	if o == nil {
		return nil
	}
	return o.Viewers
}

func (o *ACL) GetOwners() []string {
	if o == nil {
		return nil
	}
	return o.Owners
}