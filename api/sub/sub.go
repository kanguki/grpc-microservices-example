package sub

import (
	"context"
	"google.golang.org/grpc"
)

type MockClient struct {
	DoFunc func(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

func (c *MockClient) Do(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	return c.DoFunc(ctx, in, opts...)
}
