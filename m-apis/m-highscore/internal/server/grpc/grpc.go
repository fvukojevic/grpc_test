package grpc

import (
	pbhighscore "github.com/fvukojevic/grpc_test/"
	"context"
	"google.golang.org/grpc"
)

type Grpc struct {
	address string //where the grpc will listen at
	srv *grpc.Server
}

func (g *Grpc) SetHighScore(ctx context.Context, input *SetHighScoreRequest) (*SetHighScoreResponse, error) {

}