package apiServer

import (
	"context"
	"fmt"
	"github.com/kanguki/grpc-microservices-example/api/auth"
	"github.com/kanguki/grpc-microservices-example/api/div"
	"github.com/kanguki/grpc-microservices-example/api/mul"
	"github.com/kanguki/grpc-microservices-example/api/sub"
	"github.com/kanguki/grpc-microservices-example/api/sum"
	"google.golang.org/grpc"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var server *Server
var authClient = &auth.MockClient{
	LogoutFunc: func(ctx context.Context, in *auth.LogoutRequest, opts ...grpc.CallOption) (*auth.LogoutResponse, error) {
		return nil, nil
	},
	AuthorizeFunc: func(ctx context.Context, in *auth.AuthorizeRequest, opts ...grpc.CallOption) (*auth.AuthorizeResponse, error) {
		return &auth.AuthorizeResponse{UnauthorizedACLs: nil}, nil
	},
}

func TestMain(t *testing.M) {
	s := &Server{authClient: authClient}
	server = s
	t.Run()
}

func TestNewServer(t *testing.T) {
	_, err := NewServer()
	if err != nil {
		t.Fatalf("error creating api server")
	}
}

func TestServer_Auth(t *testing.T) {
	{
		emptyTokenReq := httptest.NewRequest(http.MethodGet, "/", nil)
		emptyTokenReq.Header.Add("Cookie", "tokenNotExist")
		w := httptest.NewRecorder()
		server.Auth(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))
		}, []auth.ACL{auth.ACL_READ})(w, emptyTokenReq)
		res := w.Result()
		if res.StatusCode != 401 {
			t.Errorf("expect response status 401 but got %v", res.StatusCode)
		}
	}
	{
		authClient.AuthorizeFunc = func(ctx context.Context, in *auth.AuthorizeRequest, opts ...grpc.CallOption) (*auth.AuthorizeResponse, error) {
			return &auth.AuthorizeResponse{UnauthorizedACLs: nil}, nil
		}
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Add("Cookie", "token=xxxx")
		w := httptest.NewRecorder()
		server.Auth(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))
		}, []auth.ACL{auth.ACL_READ})(w, req)
		res := w.Result()
		defer res.Body.Close()
		_, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("expected error to be nil got %v", err)
		}
	}
}

func TestServer_Login(t *testing.T) {
	{
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`can't be parsed by json'`))
		w := httptest.NewRecorder()
		server.Login(w, req)
		res := w.Result()
		if res.StatusCode != 400 {
			t.Errorf("expect response status 400 got %v", res.StatusCode)
		}
	}
	{
		authClient.AuthenticateFunc = func(ctx context.Context, in *auth.LoginRequest, opts ...grpc.CallOption) (*auth.LoginResponse, error) {
			return nil, fmt.Errorf("fake error")
		}
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"username":"mo","password":"mo"}`))
		w := httptest.NewRecorder()
		server.Login(w, req)
		res := w.Result()
		if res.StatusCode != 500 {
			t.Errorf("expect response status 500 got %v", res.StatusCode)
		}
	}
	{
		authClient.AuthenticateFunc = func(ctx context.Context, in *auth.LoginRequest, opts ...grpc.CallOption) (*auth.LoginResponse, error) {
			return &auth.LoginResponse{Token: "xxx"}, nil
		}
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"username":"mo","password":"mo"}`))
		w := httptest.NewRecorder()
		server.Login(w, req)
		res := w.Result()
		token := res.Header.Get("Set-Cookie")
		if !strings.Contains(token, "token=xxx") {
			t.Errorf("expected token to be xxx got %v", token)
		}
	}
}

func TestServer_Logout(t *testing.T) {
	authClient.LogoutFunc = func(ctx context.Context, in *auth.LogoutRequest, opts ...grpc.CallOption) (*auth.LogoutResponse, error) {
		return nil, nil
	}
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Add("Cookie", "token=xxx; expires=xxx")
	w := httptest.NewRecorder()
	server.Logout(w, req)
	res := w.Result()
	token := res.Header.Get("Set-Cookie")
	if !strings.Contains(token, "token=empty; path=/; expires=Thu, 01 Jan 1970 00:00:00 GMT") {
		t.Errorf("expected token to be empty got %v", token)
	}
}

func TestServer_Calculate(t *testing.T) {
	{
		sumClient := &sum.MockClient{}
		server.sumClient = sumClient
		{
			badCodes := []sum.Error_Code{sum.Error_BAD_REQUEST, sum.Error_UNAUTHORIZED, sum.Error_INTERNAL_SERVER_ERROR}
			for _, code := range badCodes {
				sumClient.DoFunc = func(ctx context.Context, in *sum.Request, opts ...grpc.CallOption) (*sum.Response, error) {
					return &sum.Response{Error: &sum.Error{Code: code}}, nil
				}
				req := httptest.NewRequest(http.MethodGet, "/sum?term1=1&term2=1", nil)
				w := httptest.NewRecorder()
				server.Calculate(w, req)
				res := w.Result()
				if res.StatusCode == 200 {
					t.Fatal("expect error status but got 200")
				}
			}
		}
		{
			sumClient.DoFunc = func(ctx context.Context, in *sum.Request, opts ...grpc.CallOption) (*sum.Response, error) {
				return &sum.Response{}, nil
			}
			req := httptest.NewRequest(http.MethodGet, "/sum?term1=1&term2=1", nil)
			w := httptest.NewRecorder()
			var emptyUnauthACLs []auth.ACL
			server.Calculate(w, req.WithContext(context.WithValue(context.Background(), unauthACLs, emptyUnauthACLs)))
			res := w.Result()
			if res.StatusCode != 200 {
				t.Fatalf("expect status 200 but got %v", res.StatusCode)
			}
		}
	}
	{
		subClient := &sub.MockClient{}
		server.subClient = subClient
		{
			badCodes := []sub.Error_Code{sub.Error_BAD_REQUEST, sub.Error_UNAUTHORIZED, sub.Error_INTERNAL_SERVER_ERROR}
			for _, code := range badCodes {
				subClient.DoFunc = func(ctx context.Context, in *sub.Request, opts ...grpc.CallOption) (*sub.Response, error) {
					return &sub.Response{Error: &sub.Error{Code: code}}, nil
				}
				req := httptest.NewRequest(http.MethodGet, "/sub?term1=1&term2=1", nil)
				w := httptest.NewRecorder()
				server.Calculate(w, req)
				res := w.Result()
				if res.StatusCode == 200 {
					t.Fatal("expect error status but got 200")
				}
			}
		}
		{
			subClient.DoFunc = func(ctx context.Context, in *sub.Request, opts ...grpc.CallOption) (*sub.Response, error) {
				return &sub.Response{}, nil
			}
			req := httptest.NewRequest(http.MethodGet, "/sub?term1=1&term2=1", nil)
			w := httptest.NewRecorder()
			server.Calculate(w, req)
			res := w.Result()
			if res.StatusCode != 200 {
				t.Fatalf("expect status 200 but got %v", res.StatusCode)
			}
		}
	}
	{
		mulClient := &mul.MockClient{}
		server.mulClient = mulClient
		{
			badCodes := []mul.Error_Code{mul.Error_BAD_REQUEST, mul.Error_UNAUTHORIZED, mul.Error_INTERNAL_SERVER_ERROR}
			for _, code := range badCodes {
				mulClient.DoFunc = func(ctx context.Context, in *mul.Request, opts ...grpc.CallOption) (*mul.Response, error) {
					return &mul.Response{Error: &mul.Error{Code: code}}, nil
				}
				req := httptest.NewRequest(http.MethodGet, "/mul?term1=1&term2=1", nil)
				w := httptest.NewRecorder()
				server.Calculate(w, req)
				res := w.Result()
				if res.StatusCode == 200 {
					t.Fatal("expect error status but got 200")
				}
			}
		}
		{
			mulClient.DoFunc = func(ctx context.Context, in *mul.Request, opts ...grpc.CallOption) (*mul.Response, error) {
				return &mul.Response{}, nil
			}
			req := httptest.NewRequest(http.MethodGet, "/mul?term1=1&term2=1", nil)
			w := httptest.NewRecorder()
			server.Calculate(w, req)
			res := w.Result()
			if res.StatusCode != 200 {
				t.Fatalf("expect status 200 but got %v", res.StatusCode)
			}
		}
	}
	{
		divClient := &div.MockClient{}
		server.divClient = divClient
		{
			badCodes := []div.Error_Code{div.Error_BAD_REQUEST, div.Error_UNAUTHORIZED, div.Error_INTERNAL_SERVER_ERROR}
			for _, code := range badCodes {
				divClient.DoFunc = func(ctx context.Context, in *div.Request, opts ...grpc.CallOption) (*div.Response, error) {
					return &div.Response{Error: &div.Error{Code: code}}, nil
				}
				req := httptest.NewRequest(http.MethodGet, "/div?term1=1&term2=1", nil)
				w := httptest.NewRecorder()
				server.Calculate(w, req)
				res := w.Result()
				if res.StatusCode == 200 {
					t.Fatal("expect error status but got 200")
				}
			}
		}
		{
			divClient.DoFunc = func(ctx context.Context, in *div.Request, opts ...grpc.CallOption) (*div.Response, error) {
				return &div.Response{}, nil
			}
			req := httptest.NewRequest(http.MethodGet, "/div?term1=1&term2=1", nil)
			w := httptest.NewRecorder()
			server.Calculate(w, req)
			res := w.Result()
			if res.StatusCode != 200 {
				t.Fatalf("expect status 200 but got %v", res.StatusCode)
			}
		}
	}
	{
		req := httptest.NewRequest(http.MethodGet, "/nonexist?term1=1&term2=1", nil)
		w := httptest.NewRecorder()
		server.Calculate(w, req)
		res := w.Result()
		if res.StatusCode != 400 {
			t.Fatalf("expect status 400 but got %v", res.StatusCode)
		}
	}
}
