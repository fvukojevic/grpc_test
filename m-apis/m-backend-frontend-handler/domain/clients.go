package domain

import (
	v1gameengine "github.com/fvukojevic/grpc_test/m-apis/m-game-engine/v1"
	v1highscore "github.com/fvukojevic/grpc_test/m-apis/m-highscore/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func NewGRPCGameServiceClient(serverAddr string) (v1highscore.GameClient, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())

	if err != nil {
		log.Fatal().Msgf("Failed to dial: %v", err)
		return nil, err
	}

	log.Info().Msgf("Successfully connected to [%s]", serverAddr)

	if conn == nil {
		log.Info().Msg("m-highscore connection is nil")
	}

	client := v1highscore.NewGameClient(conn)

	return client, nil
}

func NewGRPCGameEngineServiceClient(serverAddr string) (v1gameengine.GameEngineClient, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())

	if err != nil {
		log.Fatal().Msgf("Failed to dial: %v", err)
		return nil, err
	}

	log.Info().Msgf("Successfully connected to [%s]", serverAddr)

	if conn == nil {
		log.Info().Msg("m-game-engine connection is nil")
	}

	client := v1gameengine.NewGameEngineClient(conn)

	return client, nil
}
