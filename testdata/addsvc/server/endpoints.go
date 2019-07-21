package server

import (
	"context"

	"github.com/bongnv/kitgen/testdata/addsvc/pb"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints ...
type Endpoints struct {
	SumEndpoint    endpoint.Endpoint
	ConcatEndpoint endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct where each endpoint invokes
// the corresponding method on the provided service.
func MakeServerEndpoints(s pb.AddServer) Endpoints {
	return Endpoints{
		SumEndpoint:    makeSumEndpoint(s),
		ConcatEndpoint: makeConcatEndpoint(s),
	}
}

// makeSumEndpoint returns an endpoint via the passed service.
func makeSumEndpoint(s pb.AddServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.SumRequest)
		resp, e := s.Sum(ctx, req)
		return resp, e
	}
}

// makeConcatEndpoint returns an endpoint via the passed service.
func makeConcatEndpoint(s pb.AddServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.ConcatRequest)
		resp, e := s.Concat(ctx, req)
		return resp, e
	}
}
