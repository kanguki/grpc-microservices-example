package auth

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/kanguki/grpc-microservices-example/auth/log"
)

type User struct {
	Password string
	Acls     map[ACL]bool
}

type Db struct {
	Users map[string]User
}

func (db *Db) findOne(username, Password string) (User, error) {
	if user, ok := db.Users[username]; !ok {
		return User{}, fmt.Errorf("user %v doesn't exist", username)
	} else if user.Password != Password {
		return User{}, fmt.Errorf("password doesn't match")
	} else {
		return user, nil
	}
}

type Service struct {
	db     *Db
	rights map[string]map[ACL]bool
	mu     *sync.RWMutex
}

func NewService(db *Db, tokenClearDurationInMillisecond int64) *Service {
	rights := make(map[string]map[ACL]bool)
	mu := &sync.RWMutex{}
	if tokenClearDurationInMillisecond == 0 {
		tokenClearDurationInMillisecond = 24 * 60 * 60 * 1000 //24 hours
	}
	go func() {
		for {
			select {
			case <-time.After(time.Millisecond * time.Duration(tokenClearDurationInMillisecond)):
				log.Log("Cleaning rights cache by schedule.")
				for k := range rights {
					mu.Lock()
					delete(rights, k)
					mu.Unlock()
				}
			}
		}
	}()
	return &Service{
		db:     db,
		rights: rights,
		mu:     mu,
	}
}

func (s *Service) Authenticate(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	//TODO: send log to mq
	log.Log("authenticate with request: %v", *req)
	user, err := s.db.findOne(req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	token := uuid.New().String()
	s.mu.Lock()
	s.rights[token] = user.Acls
	s.mu.Unlock()
	return &LoginResponse{Token: token}, nil
}

func (s *Service) Authorize(ctx context.Context, req *AuthorizeRequest) (*AuthorizeResponse, error) {
	unauthorizedACLs := []ACL{}
	s.mu.RLock()
	rights := s.rights[req.Token]
	s.mu.RUnlock()
	if rights != nil {
		for _, v := range req.Acls {
			if rights[v] == false {
				unauthorizedACLs = append(unauthorizedACLs, v)
			}
		}
	} else {
		return &AuthorizeResponse{UnauthorizedACLs: req.Acls}, nil
	}
	if len(unauthorizedACLs) == 0 {
		unauthorizedACLs = nil
	}
	return &AuthorizeResponse{UnauthorizedACLs: unauthorizedACLs}, nil
}

func (s *Service) Logout(ctx context.Context, req *LogoutRequest) (*LogoutResponse, error) {
	s.mu.Lock()
	delete(s.rights, req.Token)
	s.mu.Unlock()
	return nil, nil
}
