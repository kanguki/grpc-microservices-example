package auth

import (
	"context"
	"testing"
	"time"
)

var service *Service

func TestMain(m *testing.M) {
	db := Db{Users: map[string]User{
		"mo":    {Password: "mo", Acls: map[ACL]bool{ACL_READ: true}},
		"admin": {Password: "admin", Acls: map[ACL]bool{ACL_READ: true, ACL_WRITE: true}},
	}}
	service = NewService(&db, 0)
	m.Run()
}

func TestNewService(t *testing.T) {
	s := NewService(nil, 10)
	s.rights["x"] = map[ACL]bool{ACL_READ: true}
	time.Sleep(15 * time.Millisecond)
	if s.rights["x"] != nil {
		t.Fatalf("cache doesn't clear by schedule. cache: %v", s.rights)
	}

}

func TestLogin(t *testing.T) {
	ctx := context.TODO()
	{
		nonExistentUserRequest := LoginRequest{Username: "hacker", Password: "crack"}
		res, err := service.Authenticate(ctx, &nonExistentUserRequest)
		if err == nil || res != nil {
			t.Fatalf("expect error User %v doesn't exist, got successful response", nonExistentUserRequest.Username)
		}
	}
	{
		wrongPasswordRequest := LoginRequest{Username: "mo", Password: "crack"}
		res, err := service.Authenticate(ctx, &wrongPasswordRequest)
		if err == nil || res != nil {
			t.Fatal("expect error wrong password, got successful response")
		}
	}
	{
		successRequest := LoginRequest{Username: "mo", Password: "mo"}
		_, err := service.Authenticate(ctx, &successRequest)
		if err != nil {
			t.Fatalf("expect success response but got error %v", err)
		}
	}
}

func TestAuthorize(t *testing.T) {
	ctx := context.TODO()
	{
		nonExistentTokenRequest := AuthorizeRequest{Token: "fake token", Acls: []ACL{ACL_READ}}
		res, _ := service.Authorize(ctx, &nonExistentTokenRequest)
		if len(res.UnauthorizedACLs) != 1 {
			t.Fatalf("expect fake token not having right to read but got successful response %v", *res)
		}
	}
	{
		loginResponse, _ := service.Authenticate(ctx, &LoginRequest{Username: "mo", Password: "mo"})
		token := loginResponse.Token
		badRequest := AuthorizeRequest{Token: token, Acls: []ACL{ACL_READ, ACL_WRITE}}
		res, _ := service.Authorize(ctx, &badRequest)
		if len(res.UnauthorizedACLs) != 1 {
			t.Fatalf("expect 1 unauthorizedACL but got response %v", *res)
		}
	}
	{
		loginResponse, _ := service.Authenticate(ctx, &LoginRequest{Username: "admin", Password: "admin"})
		token := loginResponse.Token
		successRequest := AuthorizeRequest{Token: token, Acls: []ACL{ACL_READ, ACL_WRITE}}
		res, _ := service.Authorize(ctx, &successRequest)
		if res.UnauthorizedACLs != nil {
			t.Fatalf("expect no ACL is unauthorized but ot response %v", *res)
		}
	}
}

func TestLogout(t *testing.T) {
	ctx := context.TODO()
	{
		service.Logout(ctx, &LogoutRequest{Token: "fake token but should be ok"})
	}
	{
		loginResponse, _ := service.Authenticate(ctx, &LoginRequest{Username: "admin", Password: "admin"})
		token := loginResponse.Token
		service.Logout(ctx, &LogoutRequest{Token: token})
		if service.rights[token] != nil { //not cleared yet
			t.Fatalf("token hasn't been cleared after logout. len of right-cache: %v", len(service.rights))
		}
	}
}
