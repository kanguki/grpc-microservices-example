package apiServer

import (
	"context"
	"github.com/kanguki/grpc-microservices-example/api/auth"
	"github.com/kanguki/grpc-microservices-example/api/div"
	"github.com/kanguki/grpc-microservices-example/api/log"
	"github.com/kanguki/grpc-microservices-example/api/mul"
	"github.com/kanguki/grpc-microservices-example/api/sub"
	"github.com/kanguki/grpc-microservices-example/api/sum"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"os"
	"time"
)

type Server struct {
	Port            string
	authClient      auth.AuthClient
	sumClient       sum.SumClient
	subClient       sub.SubClient
	mulClient       mul.MulClient
	divClient       div.DivClient
	timeoutInSecond time.Duration
}

func NewServer() (*Server, error) {
	server := &Server{}
	server.setPort()
	err := server.connectGrpcClients()
	if err != nil {
		return nil, err
	}
	return server, nil
}
func (s *Server) setPort() {
	port := os.Getenv("API_PORT")
	if port == "" {
		port = ":3999"
	}
	s.Port = port
}
func (s *Server) connectGrpcClients() error {
	{
		sumUrl := os.Getenv("SUM_URL")
		sumConn, err := grpc.Dial(sumUrl,
			grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Log("failed to dial %v: %v", sumUrl, err)
			return err
		}
		s.sumClient = sum.NewSumClient(sumConn)
	}
	{
		subUrl := os.Getenv("SUB_URL")
		subConn, err := grpc.Dial(subUrl,
			grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Log("failed to dial %v: %v", subUrl, err)
			return err
		}
		s.subClient = sub.NewSubClient(subConn)
	}
	{
		mulUrl := os.Getenv("MUL_URL")
		mulConn, err := grpc.Dial(mulUrl,
			grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Log("failed to dial %v: %v", mulUrl, err)
			return err
		}
		s.mulClient = mul.NewMulClient(mulConn)
	}
	{
		divUrl := os.Getenv("DIV_URL")
		divConn, err := grpc.Dial(divUrl,
			grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Log("failed to dial %v: %v", divUrl, err)
			return err
		}
		s.divClient = div.NewDivClient(divConn)
	}
	return nil
}

func (s *Server) Auth(next func(http.ResponseWriter, *http.Request), acls []auth.ACL) func(
	http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")
		if err != nil {
			w.WriteHeader(401)
			w.Write([]byte("UNAUTHORIZED"))
			return
		}
		if time.Now().After(token.Expires) {
			w.WriteHeader(401)
			w.Write([]byte("UNAUTHORIZED"))
			return
		}
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		authResponse, err := s.authClient.Authorize(ctx, &auth.AuthorizeRequest{Token: token.Value, Acls: acls})
		if err != nil {
			w.WriteHeader(401)
			w.Write([]byte("UNAUTHORIZED"))
			return
		}
		ctx = context.WithValue(ctx, "acls", authResponse.UnauthorizedACLs)
		next(w, r.WithContext(ctx))
	}
}
