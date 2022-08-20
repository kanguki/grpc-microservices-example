package apiServer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kanguki/grpc-microservices-example/api/auth"
	"github.com/kanguki/grpc-microservices-example/api/div"
	"github.com/kanguki/grpc-microservices-example/api/mul"
	"github.com/kanguki/grpc-microservices-example/api/sub"
	"github.com/kanguki/grpc-microservices-example/api/sum"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	Port       string
	authClient auth.AuthClient
	sumClient  sum.SumClient
	subClient  sub.SubClient
	mulClient  mul.MulClient
	divClient  div.DivClient
	CookieTTL  time.Duration //minute
}

func NewServer() (*Server, error) {
	server := &Server{}
	server.setPort()
	server.setTime()
	server.connectGrpcClients()
	return server, nil
}
func (s *Server) setTime() {
	_cookieTTL := os.Getenv("API_COOKIE_TTL")
	cookieTTL, _ := strconv.ParseInt(_cookieTTL, 10, 64)
	if cookieTTL == 0 {
		cookieTTL = 60
	}
	s.CookieTTL = time.Duration(cookieTTL)
}
func (s *Server) setPort() {
	port := os.Getenv("API_PORT")
	if port == "" {
		port = ":3999"
	}
	s.Port = port
}
func (s *Server) connectGrpcClients() {
	{
		authUrl := os.Getenv("AUTH_URL")
		authConn, _ := grpc.Dial(authUrl,
			grpc.WithTransportCredentials(insecure.NewCredentials()))
		s.authClient = auth.NewAuthClient(authConn)
	}
	{
		sumUrl := os.Getenv("SUM_URL")
		sumConn, _ := grpc.Dial(sumUrl,
			grpc.WithTransportCredentials(insecure.NewCredentials()))
		s.sumClient = sum.NewSumClient(sumConn)
	}
	{
		subUrl := os.Getenv("SUB_URL")
		subConn, _ := grpc.Dial(subUrl,
			grpc.WithTransportCredentials(insecure.NewCredentials()))
		s.subClient = sub.NewSubClient(subConn)
	}
	{
		mulUrl := os.Getenv("MUL_URL")
		mulConn, _ := grpc.Dial(mulUrl,
			grpc.WithTransportCredentials(insecure.NewCredentials()))
		s.mulClient = mul.NewMulClient(mulConn)
	}
	{
		divUrl := os.Getenv("DIV_URL")
		divConn, _ := grpc.Dial(divUrl,
			grpc.WithTransportCredentials(insecure.NewCredentials()))
		s.divClient = div.NewDivClient(divConn)
	}
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	lReq := loginRequest{}
	err := json.NewDecoder(r.Body).Decode(&lReq)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("bad request: " + err.Error()))
		return
	}
	res, err := s.authClient.Authenticate(context.Background(), &auth.LoginRequest{Username: lReq.Username, Password: lReq.Password})
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("error authenticating user: " + err.Error()))
		return
	}
	w.Header().Add("Set-Cookie", fmt.Sprintf("token=%s; expires=%s",
		res.Token, time.Now().Add(s.CookieTTL*time.Minute).Format(http.TimeFormat)))
	w.Write([]byte("ok"))
}

func (s *Server) Logout(w http.ResponseWriter, r *http.Request) {
	token, _ := r.Cookie("token")
	s.authClient.Logout(context.Background(), &auth.LogoutRequest{Token: token.Value})
	w.Header().Add("Set-Cookie", "token=empty; path=/; expires=Thu, 01 Jan 1970 00:00:00 GMT")
	w.Write([]byte("ok"))
}

var unauthACLs = "unauthACLs"

func (s *Server) Auth(next func(http.ResponseWriter, *http.Request), acls []auth.ACL) func(
	http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")
		if err != nil {
			w.Header().Add("Set-Cookie", "empty; path=/; expires=Thu, 01 Jan 1970 00:00:00 GMT")
			w.WriteHeader(401)
			w.Write([]byte("UNAUTHORIZED"))
			return
		}
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		authResponse, _ := s.authClient.Authorize(ctx, &auth.AuthorizeRequest{Token: token.Value, Acls: acls})
		ctx = context.WithValue(ctx, unauthACLs, authResponse.UnauthorizedACLs)
		next(w, r.WithContext(ctx))
	}
}

func (s *Server) Calculate(w http.ResponseWriter, r *http.Request) {
	var isAuthorized bool
	if _unauthACLs, ok := r.Context().Value(unauthACLs).([]auth.ACL); ok && _unauthACLs == nil {
		isAuthorized = true
	}
	queries := r.URL.Query()
	_term1, _term2 := queries["term1"][0], queries["term2"][0]
	term1, _ := strconv.ParseInt(_term1, 10, 64)
	term2, _ := strconv.ParseInt(_term2, 10, 64)
	operationType := strings.TrimPrefix(r.URL.Path, "/")
	switch operationType {
	case "sum":
		res, _ := s.sumClient.Do(context.Background(), &sum.Request{Term1: term1, Term2: term2, IsAuthorized: isAuthorized})
		if err := res.Error; err != nil {
			switch err.Code {
			case sum.Error_UNAUTHORIZED:
				w.WriteHeader(401)
				w.Write([]byte(err.Message))
				return
			case sum.Error_INTERNAL_SERVER_ERROR:
				w.WriteHeader(500)
				w.Write([]byte(err.Message))
				return
			case sum.Error_BAD_REQUEST:
				w.WriteHeader(400)
				w.Write([]byte(err.Message))
				return
			}
		}
		w.Write([]byte(fmt.Sprintf("%d", res.Sum)))
		break
	case "sub":
		res, _ := s.subClient.Do(context.Background(), &sub.Request{Term1: term1, Term2: term2, IsAuthorized: isAuthorized})
		if err := res.Error; err != nil {
			switch err.Code {
			case sub.Error_UNAUTHORIZED:
				w.WriteHeader(401)
				w.Write([]byte(err.Message))
				return
			case sub.Error_INTERNAL_SERVER_ERROR:
				w.WriteHeader(500)
				w.Write([]byte(err.Message))
				return
			case sub.Error_BAD_REQUEST:
				w.WriteHeader(400)
				w.Write([]byte(err.Message))
				return
			}
		}
		w.Write([]byte(fmt.Sprintf("%d", res.Sub)))
		break
	case "mul":
		res, _ := s.mulClient.Do(context.Background(), &mul.Request{Term1: term1, Term2: term2, IsAuthorized: isAuthorized})
		if err := res.Error; err != nil {
			switch err.Code {
			case mul.Error_UNAUTHORIZED:
				w.WriteHeader(401)
				w.Write([]byte(err.Message))
				return
			case mul.Error_INTERNAL_SERVER_ERROR:
				w.WriteHeader(500)
				w.Write([]byte(err.Message))
				return
			case mul.Error_BAD_REQUEST:
				w.WriteHeader(400)
				w.Write([]byte(err.Message))
				return
			}
		}
		w.Write([]byte(fmt.Sprintf("%d", res.Mul)))
		break
	case "div":
		res, _ := s.divClient.Do(context.Background(), &div.Request{Term1: term1, Term2: term2, IsAuthorized: isAuthorized})
		if err := res.Error; err != nil {
			switch err.Code {
			case div.Error_UNAUTHORIZED:
				w.WriteHeader(401)
				w.Write([]byte(err.Message))
				return
			case div.Error_INTERNAL_SERVER_ERROR:
				w.WriteHeader(500)
				w.Write([]byte(err.Message))
				return
			case div.Error_BAD_REQUEST:
				w.WriteHeader(400)
				w.Write([]byte(err.Message))
				return
			}
		}
		w.Write([]byte(fmt.Sprintf("%d", res.Div)))
		break
	default:
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf("operation %s is not supported", operationType)))
	}

}

// func (s *Server) Cors(next func(http.ResponseWriter, *http.Request)) func(
// 	http.ResponseWriter, *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		allowOrigins := os.Getenv("API_ALLOW_ORIGINS")
// 		if allowOrigins != "" {
// 			w.Header().Set("Access-Control-Allow-Origin", allowOrigins)
// 		}
// 		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 		w.Header().Set("Access-Control-Allow-Headers", "*")
// 		w.Header().Set("Access-Control-Allow-Credentials","true")
// 		if r.Method == http.MethodOptions {
// 			return
// 		}
// 		next(w, r)
// 	}
// }
