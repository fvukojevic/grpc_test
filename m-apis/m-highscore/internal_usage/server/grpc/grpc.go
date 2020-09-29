package grpc

import (
	"context"
	v1highscore "github.com/fvukojevic/grpc_test/m-apis/m-highscore/v1"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
)

type Grpc struct {
	address string //address where the gRPC will listen at
	srv     *grpc.Server
}

func NewServer(address string) *Grpc {
	return &Grpc{
		address: address,
	}
}

var HighScore = 999999.0

func (g *Grpc) SetHighScore(ctx context.Context, input *v1highscore.SetHighScoreRequest) (*v1highscore.SetHighScoreResponse, error) {
	log.Info().Msg("SetHighscore in m-highscore is called")

	HighScore = input.HighScore
	return &v1highscore.SetHighScoreResponse{
		Set: true,
	}, nil
}

func (g *Grpc) GetHighScore(ctx context.Context, input *v1highscore.GetHighScoreRequest) (*v1highscore.GetHighScoreResponse, error) {
	log.Info().Msg("GetHighscore in m-highscore is called")

	return &v1highscore.GetHighScoreResponse{
		HighScore: HighScore,
	}, nil
}

func (g *Grpc) ListenAndServe() error {
	listener, err := net.Listen("tcp", g.address)

	if err != nil {
		return errors.Wrap(err, "Failed to open tcp port")
	}

	var serverOpts []grpc.ServerOption

	g.srv = grpc.NewServer(serverOpts...)

	v1highscore.RegisterGameServer(g.srv, g)

	log.Info().Str("Address", g.address).Msg("Starting gRPC server for highscore microservice")

	if err := g.srv.Serve(listener); err != nil {
		return errors.Wrap(err, "Failed to start gRPC server")
	}

	return nil
}
