// Code generated by goa v3.2.2, DO NOT EDIT.
//
// resource endpoints
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package resource

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "resource" service endpoints.
type Endpoints struct {
	Query             goa.Endpoint
	List              goa.Endpoint
	VersionsByID      goa.Endpoint
	ByKindNameVersion goa.Endpoint
	ByVersionID       goa.Endpoint
	ByKindName        goa.Endpoint
	ByID              goa.Endpoint
}

// NewEndpoints wraps the methods of the "resource" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Query:             NewQueryEndpoint(s),
		List:              NewListEndpoint(s),
		VersionsByID:      NewVersionsByIDEndpoint(s),
		ByKindNameVersion: NewByKindNameVersionEndpoint(s),
		ByVersionID:       NewByVersionIDEndpoint(s),
		ByKindName:        NewByKindNameEndpoint(s),
		ByID:              NewByIDEndpoint(s),
	}
}

// Use applies the given middleware to all the "resource" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Query = m(e.Query)
	e.List = m(e.List)
	e.VersionsByID = m(e.VersionsByID)
	e.ByKindNameVersion = m(e.ByKindNameVersion)
	e.ByVersionID = m(e.ByVersionID)
	e.ByKindName = m(e.ByKindName)
	e.ByID = m(e.ByID)
}

// NewQueryEndpoint returns an endpoint function that calls the method "Query"
// of service "resource".
func NewQueryEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*QueryPayload)
		res, err := s.Query(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedResourceCollection(res, "withoutVersion")
		return vres, nil
	}
}

// NewListEndpoint returns an endpoint function that calls the method "List" of
// service "resource".
func NewListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ListPayload)
		res, err := s.List(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedResourceCollection(res, "withoutVersion")
		return vres, nil
	}
}

// NewVersionsByIDEndpoint returns an endpoint function that calls the method
// "VersionsByID" of service "resource".
func NewVersionsByIDEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*VersionsByIDPayload)
		res, err := s.VersionsByID(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedVersions(res, "default")
		return vres, nil
	}
}

// NewByKindNameVersionEndpoint returns an endpoint function that calls the
// method "ByKindNameVersion" of service "resource".
func NewByKindNameVersionEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ByKindNameVersionPayload)
		res, err := s.ByKindNameVersion(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedVersion(res, "default")
		return vres, nil
	}
}

// NewByVersionIDEndpoint returns an endpoint function that calls the method
// "ByVersionId" of service "resource".
func NewByVersionIDEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ByVersionIDPayload)
		res, err := s.ByVersionID(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedVersion(res, "default")
		return vres, nil
	}
}

// NewByKindNameEndpoint returns an endpoint function that calls the method
// "ByKindName" of service "resource".
func NewByKindNameEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ByKindNamePayload)
		res, err := s.ByKindName(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedResourceCollection(res, "withoutVersion")
		return vres, nil
	}
}

// NewByIDEndpoint returns an endpoint function that calls the method "ById" of
// service "resource".
func NewByIDEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ByIDPayload)
		res, err := s.ByID(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedResource(res, "default")
		return vres, nil
	}
}
