// Code generated by goa v3.2.2, DO NOT EDIT.
//
// resource client HTTP transport
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the resource service endpoint HTTP clients.
type Client struct {
	// Query Doer is the HTTP client used to make requests to the Query endpoint.
	QueryDoer goahttp.Doer

	// List Doer is the HTTP client used to make requests to the List endpoint.
	ListDoer goahttp.Doer

	// VersionsByID Doer is the HTTP client used to make requests to the
	// VersionsByID endpoint.
	VersionsByIDDoer goahttp.Doer

	// ByKindNameVersion Doer is the HTTP client used to make requests to the
	// ByKindNameVersion endpoint.
	ByKindNameVersionDoer goahttp.Doer

	// ByVersionID Doer is the HTTP client used to make requests to the ByVersionId
	// endpoint.
	ByVersionIDDoer goahttp.Doer

	// ByKindName Doer is the HTTP client used to make requests to the ByKindName
	// endpoint.
	ByKindNameDoer goahttp.Doer

	// ByID Doer is the HTTP client used to make requests to the ById endpoint.
	ByIDDoer goahttp.Doer

	// CORS Doer is the HTTP client used to make requests to the  endpoint.
	CORSDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the resource service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		QueryDoer:             doer,
		ListDoer:              doer,
		VersionsByIDDoer:      doer,
		ByKindNameVersionDoer: doer,
		ByVersionIDDoer:       doer,
		ByKindNameDoer:        doer,
		ByIDDoer:              doer,
		CORSDoer:              doer,
		RestoreResponseBody:   restoreBody,
		scheme:                scheme,
		host:                  host,
		decoder:               dec,
		encoder:               enc,
	}
}

// Query returns an endpoint that makes HTTP requests to the resource service
// Query server.
func (c *Client) Query() goa.Endpoint {
	var (
		encodeRequest  = EncodeQueryRequest(c.encoder)
		decodeResponse = DecodeQueryResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildQueryRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.QueryDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("resource", "Query", err)
		}
		return decodeResponse(resp)
	}
}

// List returns an endpoint that makes HTTP requests to the resource service
// List server.
func (c *Client) List() goa.Endpoint {
	var (
		encodeRequest  = EncodeListRequest(c.encoder)
		decodeResponse = DecodeListResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildListRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ListDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("resource", "List", err)
		}
		return decodeResponse(resp)
	}
}

// VersionsByID returns an endpoint that makes HTTP requests to the resource
// service VersionsByID server.
func (c *Client) VersionsByID() goa.Endpoint {
	var (
		decodeResponse = DecodeVersionsByIDResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildVersionsByIDRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.VersionsByIDDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("resource", "VersionsByID", err)
		}
		return decodeResponse(resp)
	}
}

// ByKindNameVersion returns an endpoint that makes HTTP requests to the
// resource service ByKindNameVersion server.
func (c *Client) ByKindNameVersion() goa.Endpoint {
	var (
		decodeResponse = DecodeByKindNameVersionResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildByKindNameVersionRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ByKindNameVersionDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("resource", "ByKindNameVersion", err)
		}
		return decodeResponse(resp)
	}
}

// ByVersionID returns an endpoint that makes HTTP requests to the resource
// service ByVersionId server.
func (c *Client) ByVersionID() goa.Endpoint {
	var (
		decodeResponse = DecodeByVersionIDResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildByVersionIDRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ByVersionIDDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("resource", "ByVersionId", err)
		}
		return decodeResponse(resp)
	}
}

// ByKindName returns an endpoint that makes HTTP requests to the resource
// service ByKindName server.
func (c *Client) ByKindName() goa.Endpoint {
	var (
		decodeResponse = DecodeByKindNameResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildByKindNameRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ByKindNameDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("resource", "ByKindName", err)
		}
		return decodeResponse(resp)
	}
}

// ByID returns an endpoint that makes HTTP requests to the resource service
// ById server.
func (c *Client) ByID() goa.Endpoint {
	var (
		decodeResponse = DecodeByIDResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildByIDRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ByIDDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("resource", "ById", err)
		}
		return decodeResponse(resp)
	}
}
