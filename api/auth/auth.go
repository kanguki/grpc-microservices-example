package auth

import (
	"context"
	"google.golang.org/grpc"
)

type MockClient struct {
	LogoutFunc       func(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	AuthenticateFunc func(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	AuthorizeFunc    func(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error)
}

func (c *MockClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	return c.LogoutFunc(ctx, in, opts...)
}

func (c *MockClient) Authenticate(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	return c.AuthenticateFunc(ctx, in, opts...)
}

func (c *MockClient) Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error) {
	return c.AuthorizeFunc(ctx, in, opts...)
}
