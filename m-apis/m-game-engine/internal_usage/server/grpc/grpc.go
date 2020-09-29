package grpc

import (
	"context"
	"github.com/fvukojevic/grpc_test/m-apis/m-game-engine/internal_usage/server/logic"
	v1gameengine "github.com/fvukojevic/grpc_test/m-apis/m-game-engine/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
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

func (g *Grpc) GetSize(context.Context, *v1gameengine.GetSizeRequest) (*v1gameengine.GetSizeResponse, error) {
	log.Info().Msg("GetSize in m-game-engine called")

	return &v1gameengine.GetSizeResponse{
		Size: logic.GetSize(),
	}, nil
}
func (g *Grpc) SetScore(context.Context, *v1gameengine.SetScoreRequest) (*v1gameengine.SetScoreResponse, error) {

}
