package grpcx

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/kwilteam/kwil-db/internal/logx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func NewServer(logger logx.Logger) *grpc.Server {
	l := logger.WithOptions(zap.WithCaller(false))

	return grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_zap.StreamServerInterceptor(l),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(l),
		)),
	)
}
