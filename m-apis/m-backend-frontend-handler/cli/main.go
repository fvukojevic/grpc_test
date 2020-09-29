package main

import (
	"flag"
	"github.com/fvukojevic/grpc_test/m-apis/m-backend-frontend-handler/domain"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

var router = gin.Default()

func main() {
	grpcAddressHighScore := flag.String("address-m-highscore", "localhost:50051", "The gRPC server address for m-highscore microservice")
	grpcAddressGameEngine := flag.String("address-m-game-engine", "localhost:60051", "The gRPC server address for m-game-engine microservice")

	serverAddress := flag.String("address-http", ":8081", "HTTP server address")

	flag.Parse()

	gameHighscoreClient, err := domain.NewGRPCGameServiceClient(*grpcAddressHighScore)
	if err != nil {
		log.Error().Msg("error in creating a client for m-highscore")
	}
	gameEngineClient, err := domain.NewGRPCGameEngineServiceClient(*grpcAddressGameEngine)
	if err != nil {
		log.Error().Msg("error in creating a client for m-game-engine")
	}

	gameResource := domain.NewGameResource(gameEngineClient, gameHighscoreClient)

	router.GET("/geths", gameResource.GetHighScore)
	router.GET("/seths/:hs", gameResource.SetHighScore)
	router.GET("/getsize", gameResource.GetSize)
	router.GET("/setscore/:score", gameResource.SetScore)

	router.Run(*serverAddress)

	log.Info().Msgf("Started the microservice at: %v", *serverAddress)
}
