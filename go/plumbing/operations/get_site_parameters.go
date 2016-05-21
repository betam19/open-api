package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSiteParams creates a new GetSiteParams object
// with the default values initialized.
func NewGetSiteParams() *GetSiteParams {
	var ()
	return &GetSiteParams{}
}

/*GetSiteParams contains all the parameters to send to the API endpoint
for the get site operation typically these are written to a http.Request
*/
type GetSiteParams struct {

	/*SiteID*/
	SiteID string
}

// WithSiteID adds the siteId to the get site params
func (o *GetSiteParams) WithSiteID(SiteID string) *GetSiteParams {
	o.SiteID = SiteID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetSiteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}