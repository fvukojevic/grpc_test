package grpc

import (
	"context"
	"github.com/fvukojevic/grpc_test/m-apis/m-game-engine/internal_usage/server/logic"
	v1gameengine "github.com/fvukojevic/grpc_test/m-apis/m-game-engine/v1"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
)

type Grpc struct {
	address string
	srv     *grpc.Server
}

func NewServer(address string) *Grpc {
	return &Grpc{
		address: address,
	}
}

func (g *Grpc) GetSize(ctx context.Context, input *v1gameengine.GetSizeRequest) (*v1gameengine.GetSizeResponse, error) {
	log.Info().Msg("GetSize in m-game-engine called")

	return &v1gameengine.GetSizeResponse{
		Size: logic.GetSize(),
	}, nil
}
func (g *Grpc) SetScore(ctx context.Context, input *v1gameengine.SetScoreRequest) (*v1gameengine.SetScoreResponse, error) {
	log.Info().Msg("SetScore in m-game-engine called")

	logic.SetScore(input.Score)
	return &v1gameengine.SetScoreResponse{
		Set: true,
	}, nil
}

func (g *Grpc) ListenAndServe() error {
	listener, err := net.Listen("tcp", g.address)

	if err != nil {
		return errors.Wrap(err, "Failed to open tcp port")
	}

	var serverOpts []grpc.ServerOption

	g.srv = grpc.NewServer(serverOpts...)

	v1gameengine.RegisterGameEngineServer(g.srv, g)

	log.Info().Str("Address", g.address).Msg("Starting gRPC server for game engine microservice")

	if err := g.srv.Serve(listener); err != nil {
		return errors.Wrap(err, "Failed to start gRPC server")
	}

	return nil
}
