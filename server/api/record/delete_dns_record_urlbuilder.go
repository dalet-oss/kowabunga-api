// Code generated by go-swagger; DO NOT EDIT.

package record

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// DeleteDNSRecordURL generates an URL for the delete Dns record operation
type DeleteDNSRecordURL struct {
	RecordID string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DeleteDNSRecordURL) WithBasePath(bp string) *DeleteDNSRecordURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DeleteDNSRecordURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *DeleteDNSRecordURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/record/{recordId}"

	recordID := o.RecordID
	if recordID != "" {
		_path = strings.Replace(_path, "{recordId}", recordID, -1)
	} else {
		return nil, errors.New("recordId is required on DeleteDNSRecordURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *DeleteDNSRecordURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *DeleteDNSRecordURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *DeleteDNSRecordURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on DeleteDNSRecordURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on DeleteDNSRecordURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *DeleteDNSRecordURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}