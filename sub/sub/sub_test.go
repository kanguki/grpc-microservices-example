package sub

import (
	"context"
	"testing"
)

func TestSum(t *testing.T) {
	ctx := context.TODO()
	service := Service{Max: 1000, Min: -1000, MaxFreeTier: 10, MinFreeTier: -10}
	{
		badRequest := Request{Term1: 1001, Term2: 0, IsAuthorized: true}
		res, _ := service.Do(ctx, &badRequest)
		if !(res.Error != nil && res.Error.Code == Error_BAD_REQUEST) {
			t.Fatalf("expect bad request error but got response %v", *res)
		}
	}
	{
		unauthorizedRequest := Request{Term1: 11, Term2: -3, IsAuthorized: false}
		res, _ := service.Do(ctx, &unauthorizedRequest)
		if !(res.Error != nil && res.Error.Code == Error_UNAUTHORIZED) {
			t.Fatalf("expect unauthorized error but got response %v", *res)
		}
	}
	{
		successRequest := Request{Term1: 5, Term2: -5, IsAuthorized: true}
		res, _ := service.Do(ctx, &successRequest)
		if res.Error != nil || res.Sub != 10 {
			t.Fatalf("expect success response but got response %v", *res)
		}
	}
}
