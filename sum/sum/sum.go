package sum

import (
	"context"
	"fmt"
	"github.com/kanguki/grpc-microservices-example/sum/log"
)

type Service struct {
	Max, Min                 int64 //Max/Min number it can handle
	MaxFreeTier, MinFreeTier int64 //Max/Min for public users
}

func (s Service) Do(ctx context.Context, req *Request) (*Response, error) {
	log.Log("Do get called with request %v", *req)
	res := Response{}
	if req.Term1 > s.Max || req.Term2 > s.Max ||
		req.Term1 < s.Min || req.Term2 < s.Min {
		res.Error = &Error{
			Code:    Error_BAD_REQUEST,
			Message: fmt.Sprintf("cannot handle number smaller than %v or bigger than %v", s.Min, s.Max),
		}
		return &res, nil
	}
	if !req.IsAuthorized && (req.Term1 > s.MaxFreeTier || req.Term2 > s.MaxFreeTier ||
		req.Term1 < s.MinFreeTier || req.Term2 < s.MinFreeTier) {
		res.Error = &Error{
			Code:    Error_UNAUTHORIZED,
			Message: fmt.Sprintf("unauthenticated cannot use number smaller than %v or bigger than %v", s.MinFreeTier, s.MaxFreeTier),
		}
		return &res, nil
	}
	res.Sum = req.Term1 + req.Term2
	return &res, nil
}
